package graph

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/jwt"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
)

// ---------------------------------------------------------------------------
// Error Interpretation Helpers
// ---------------------------------------------------------------------------

type errorCode string

const (
	applicationThrottled errorCode = "ApplicationThrottled"
	// this auth error is a catch-all used by graph in a variety of cases:
	// users without licenses, bad jwts, missing account permissions, etc.
	AuthenticationError errorCode = "AuthenticationError"
	// cannotOpenFileAttachment happen when an attachment is
	// inaccessible. The error message is usually "OLE conversion
	// failed for an attachment."
	cannotOpenFileAttachment errorCode = "ErrorCannotOpenFileAttachment"
	emailFolderNotFound      errorCode = "ErrorSyncFolderNotFound"
	ErrorAccessDenied        errorCode = "ErrorAccessDenied"
	errorItemNotFound        errorCode = "ErrorItemNotFound"
	// This error occurs when an attempt is made to create a folder that has
	// the same name as another folder in the same parent. Such duplicate folder
	// names are not allowed by graph.
	folderExists errorCode = "ErrorFolderExists"
	// Some datacenters are returning this when we try to get the inbox of a user
	// that doesn't exist.
	invalidUser                 errorCode = "ErrorInvalidUser"
	invalidAuthenticationToken  errorCode = "InvalidAuthenticationToken"
	itemNotFound                errorCode = "itemNotFound"
	MailboxNotEnabledForRESTAPI errorCode = "MailboxNotEnabledForRESTAPI"
	malwareDetected             errorCode = "malwareDetected"
	// nameAlreadyExists occurs when a request with
	// @microsoft.graph.conflictBehavior=fail finds a conflicting file.
	nameAlreadyExists       errorCode = "nameAlreadyExists"
	NotAllowed              errorCode = "notAllowed"
	noResolvedUsers         errorCode = "noResolvedUsers"
	QuotaExceeded           errorCode = "ErrorQuotaExceeded"
	RequestResourceNotFound errorCode = "Request_ResourceNotFound"
	// Returned when we try to get the inbox of a user that doesn't exist.
	ResourceNotFound   errorCode = "ResourceNotFound"
	resyncRequired     errorCode = "ResyncRequired"
	syncFolderNotFound errorCode = "ErrorSyncFolderNotFound"
	syncStateInvalid   errorCode = "SyncStateInvalid"
	syncStateNotFound  errorCode = "SyncStateNotFound"
)

// inner error codes
const (
	ResourceLocked errorCode = "resourceLocked"
)

type errorMessage string

const (
	IOErrDuringRead                 errorMessage = "IO error during request payload read"
	MysiteURLNotFound               errorMessage = "unable to retrieve user's mysite url"
	MysiteNotFound                  errorMessage = "user's mysite not found"
	NoSPLicense                     errorMessage = "Tenant does not have a SPO license"
	ParameterDeltaTokenNotSupported errorMessage = "Parameter 'DeltaToken' not supported for this request"
	usersCannotBeResolved           errorMessage = "One or more users could not be resolved"
	requestedSiteCouldNotBeFound    errorMessage = "Requested site could not be found"
)

const (
	LabelsMalware             = "malware_detected"
	LabelsMysiteNotFound      = "mysite_not_found"
	LabelsNoSharePointLicense = "no_sharepoint_license"

	// LabelsSkippable is used to determine if an error is skippable
	LabelsSkippable = "skippable_errors"
)

var (
	// ErrApplicationThrottled occurs if throttling retries are exhausted and completely
	// fails out.
	ErrApplicationThrottled = clues.New("application throttled")

	// The folder or item was deleted between the time we identified
	// it and when we tried to fetch data for it.
	ErrDeletedInFlight = clues.New("deleted in flight")

	// ErrItemAlreadyExistsConflict denotes that a post or put attempted to create
	// an item which already exists by some unique identifier.  The identifier is
	// not always the id.  For example, in onedrive, this error can be produced
	// when filenames collide in a @microsoft.graph.conflictBehavior=fail request.
	ErrItemAlreadyExistsConflict = clues.New("item already exists")

	// ErrMultipleResultsMatchIdentifier describes a situation where we're doing a lookup
	// in some way other than by canonical url ID (ex: filtering, searching, etc).
	// This error should only be returned if a unique result is an expected constraint
	// of the call results.  If it's possible to opportunistically select one of the many
	// replies, no error should get returned.
	ErrMultipleResultsMatchIdentifier = clues.New("multiple results match the identifier")

	// ErrResourceLocked occurs when a resource has had its access locked.
	// Example case: https://learn.microsoft.com/en-us/sharepoint/manage-lock-status
	// This makes the resource inaccessible for any Corso operations.
	ErrResourceLocked = clues.New("resource has been locked and must be unlocked by an administrator")

	// ErrServiceNotEnabled identifies that a resource owner does not have
	// access to a given service.
	ErrServiceNotEnabled = clues.New("service is not enabled for that resource owner")

	ErrResourceOwnerNotFound = clues.New("resource owner not found in tenant")

	ErrTokenExpired = clues.New("jwt token expired")
)

func IsErrApplicationThrottled(err error) bool {
	return errors.Is(err, ErrApplicationThrottled) ||
		hasErrorCode(err, applicationThrottled)
}

func IsErrAuthenticationError(err error) bool {
	return hasErrorCode(err, AuthenticationError)
}

func IsErrDeletedInFlight(err error) bool {
	if errors.Is(err, ErrDeletedInFlight) {
		return true
	}

	if hasErrorCode(
		err,
		errorItemNotFound,
		itemNotFound,
		syncFolderNotFound) {
		return true
	}

	return false
}

func IsErrItemNotFound(err error) bool {
	return hasErrorCode(err, itemNotFound)
}

func IsErrInvalidDelta(err error) bool {
	return hasErrorCode(err, syncStateNotFound, resyncRequired, syncStateInvalid)
}

func IsErrDeltaNotSupported(err error) bool {
	return hasErrorMessage(err, ParameterDeltaTokenNotSupported)
}

func IsErrQuotaExceeded(err error) bool {
	return hasErrorCode(err, QuotaExceeded)
}

func IsErrExchangeMailFolderNotFound(err error) bool {
	// Not sure if we can actually see a resourceNotFound error here. I've only
	// seen the latter two.
	return hasErrorCode(err, ResourceNotFound, errorItemNotFound, MailboxNotEnabledForRESTAPI)
}

func IsErrUserNotFound(err error) bool {
	if hasErrorCode(err, RequestResourceNotFound, invalidUser) {
		return true
	}

	if hasErrorCode(err, ResourceNotFound) {
		var odErr odataerrors.ODataErrorable
		if !errors.As(err, &odErr) {
			return false
		}

		mainMsg, _, _ := errData(odErr)

		return strings.Contains(strings.ToLower(mainMsg), "user")
	}

	return false
}

func IsErrCannotOpenFileAttachment(err error) bool {
	return hasErrorCode(err, cannotOpenFileAttachment)
}

func IsErrAccessDenied(err error) bool {
	return hasErrorCode(err, ErrorAccessDenied) ||
		clues.HasLabel(err, LabelStatus(http.StatusForbidden))
}

func IsErrTimeout(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		return err.Timeout()
	}

	return errors.Is(err, context.Canceled) ||
		errors.Is(err, context.DeadlineExceeded) ||
		errors.Is(err, http.ErrHandlerTimeout) ||
		os.IsTimeout(err)
}

func IsErrConnectionReset(err error) bool {
	return errors.Is(err, syscall.ECONNRESET)
}

func IsErrUnauthorizedOrBadToken(err error) bool {
	return clues.HasLabel(err, LabelStatus(http.StatusUnauthorized)) ||
		hasErrorCode(err, invalidAuthenticationToken) ||
		errors.Is(err, ErrTokenExpired)
}

func IsErrBadJWTToken(err error) bool {
	return hasErrorCode(err, invalidAuthenticationToken)
}

func IsErrItemAlreadyExistsConflict(err error) bool {
	return errors.Is(err, ErrItemAlreadyExistsConflict) ||
		hasErrorCode(err, nameAlreadyExists)
}

// LabelStatus transforms the provided statusCode into
// a standard label that can be attached to a clues error
// and later reviewed when checking error statuses.
func LabelStatus(statusCode int) string {
	return fmt.Sprintf("status_code_%d", statusCode)
}

// IsMalware is true if the graphAPI returns a "malware detected" error code.
func IsMalware(err error) bool {
	return hasErrorCode(err, malwareDetected)
}

func IsMalwareResp(ctx context.Context, resp *http.Response) bool {
	// https://learn.microsoft.com/en-us/openspecs/sharepoint_protocols/ms-wsshp/ba4ee7a8-704c-4e9c-ab14-fa44c574bdf4
	// https://learn.microsoft.com/en-us/openspecs/sharepoint_protocols/ms-wdvmoduu/6fa6d4a9-ac18-4cd7-b696-8a3b14a98291
	return resp != nil &&
		len(resp.Header) > 0 &&
		resp.Header.Get("X-Virus-Infected") == "true"
}

func IsErrFolderExists(err error) bool {
	return hasErrorCode(err, folderExists)
}

func IsErrUsersCannotBeResolved(err error) bool {
	return hasErrorCode(err, noResolvedUsers) || hasErrorMessage(err, usersCannotBeResolved)
}

func IsErrSiteNotFound(err error) bool {
	return hasErrorMessage(err, requestedSiteCouldNotBeFound)
}

func IsErrResourceLocked(err error) bool {
	return errors.Is(err, ErrResourceLocked) ||
		hasInnerErrorCode(err, ResourceLocked) ||
		hasErrorCode(err, NotAllowed)
}

// ---------------------------------------------------------------------------
// error parsers
// ---------------------------------------------------------------------------

func hasErrorCode(err error, codes ...errorCode) bool {
	if err == nil {
		return false
	}

	var oDataError odataerrors.ODataErrorable
	if !errors.As(err, &oDataError) {
		return false
	}

	code, ok := ptr.ValOK(oDataError.GetErrorEscaped().GetCode())
	if !ok {
		return false
	}

	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = string(c)
	}

	return filters.Equal(cs).Compare(code)
}

func hasInnerErrorCode(err error, codes ...errorCode) bool {
	if err == nil {
		return false
	}

	var oDataError odataerrors.ODataErrorable
	if !errors.As(err, &oDataError) {
		return false
	}

	inner := oDataError.GetErrorEscaped().GetInnerError()
	if inner == nil {
		return false
	}

	code, err := str.AnyValueToString("code", inner.GetAdditionalData())
	if err != nil {
		return false
	}

	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = string(c)
	}

	return filters.Equal(cs).Compare(code)
}

// only use this as a last resort.  Prefer the code or statuscode if possible.
func hasErrorMessage(err error, msgs ...errorMessage) bool {
	if err == nil {
		return false
	}

	var oDataError odataerrors.ODataErrorable
	if !errors.As(err, &oDataError) {
		return false
	}

	msg, ok := ptr.ValOK(oDataError.GetErrorEscaped().GetMessage())
	if !ok {
		return false
	}

	cs := make([]string, len(msgs))
	for i, c := range msgs {
		cs[i] = string(c)
	}

	return filters.In(cs).Compare(msg)
}

// Wrap is a helper function that extracts ODataError metadata from
// the error.  If the error is not an ODataError type, returns the error.
func Wrap(ctx context.Context, e error, msg string) *clues.Err {
	if e == nil {
		return nil
	}

	var oDataError odataerrors.ODataErrorable
	if !errors.As(e, &oDataError) {
		return clues.WrapWC(ctx, e, msg).WithTrace(1)
	}

	mainMsg, data, innerMsg := errData(oDataError)

	if len(mainMsg) > 0 {
		e = clues.Stack(e, clues.New(mainMsg))
	}

	ce := clues.WrapWC(ctx, e, msg).With(data...).WithTrace(1)

	return setLabels(ce, innerMsg)
}

// Stack is a helper function that extracts ODataError metadata from
// the error.  If the error is not an ODataError type, returns the error.
func Stack(ctx context.Context, e error) *clues.Err {
	if e == nil {
		return nil
	}

	var oDataError *odataerrors.ODataError
	if !errors.As(e, &oDataError) {
		return clues.StackWC(ctx, e).WithTrace(1)
	}

	mainMsg, data, innerMsg := errData(oDataError)

	if len(mainMsg) > 0 {
		e = clues.Stack(e, clues.New(mainMsg))
	}

	ce := clues.StackWC(ctx, e).With(data...).WithTrace(1)

	return setLabels(ce, innerMsg)
}

// stackReq is a helper function that extracts ODataError metadata from
// the error, plus http req/resp data.  If the error is not an ODataError
// type, returns the error with only the req/resp values.
func stackReq(
	ctx context.Context,
	req *http.Request,
	resp *http.Response,
	e error,
) *clues.Err {
	if e == nil {
		return nil
	}

	se := Stack(ctx, e).
		WithMap(reqData(req)).
		WithMap(respData(resp))

	return se
}

// Checks for the following conditions and labels the error accordingly:
// * mysiteNotFound | mysiteURLNotFound
// * malware
func setLabels(err *clues.Err, msg string) *clues.Err {
	if err == nil {
		return nil
	}

	f := filters.Contains([]string{msg})

	if f.Compare(string(MysiteNotFound)) ||
		f.Compare(string(MysiteURLNotFound)) {
		err = err.Label(LabelsMysiteNotFound)
	}

	if f.Compare(string(NoSPLicense)) {
		err = err.Label(LabelsNoSharePointLicense)
	}

	if IsMalware(err) {
		err = err.Label(LabelsMalware)
	}

	return err
}

func errData(err odataerrors.ODataErrorable) (string, []any, string) {
	data := make([]any, 0)

	// Get MainError
	mainErr := err.GetErrorEscaped()
	mainMsg := ptr.Val(mainErr.GetMessage())

	data = appendIf(data, "odataerror_code", mainErr.GetCode())
	data = appendIf(data, "odataerror_message", mainErr.GetMessage())
	data = appendIf(data, "odataerror_target", mainErr.GetTarget())
	msgConcat := ptr.Val(mainErr.GetMessage()) + ptr.Val(mainErr.GetCode())

	for i, d := range mainErr.GetDetails() {
		pfx := fmt.Sprintf("odataerror_details_%d_", i)
		data = appendIf(data, pfx+"code", d.GetCode())
		data = appendIf(data, pfx+"message", d.GetMessage())
		data = appendIf(data, pfx+"target", d.GetTarget())
		msgConcat += ptr.Val(d.GetMessage())
	}

	inner := mainErr.GetInnerError()
	if inner != nil {
		data = appendIf(data, "odataerror_inner_cli_req_id", inner.GetClientRequestId())
		data = appendIf(data, "odataerror_inner_req_id", inner.GetRequestId())
	}

	return mainMsg, data, strings.ToLower(msgConcat)
}

func reqData(req *http.Request) map[string]any {
	if req == nil {
		return nil
	}

	r := map[string]any{}
	r["req_method"] = req.Method
	r["req_len"] = req.ContentLength

	if req.URL != nil {
		r["req_url"] = LoggableURL(req.URL.String())
	}

	return r
}

func respData(resp *http.Response) map[string]any {
	if resp == nil {
		return nil
	}

	r := map[string]any{}
	r["resp_status"] = resp.Status
	r["resp_len"] = resp.ContentLength

	return r
}

func appendIf(a []any, k string, v *string) []any {
	if v == nil {
		return a
	}

	return append(a, k, *v)
}

// ItemInfo gathers potentially useful information about a drive item,
// and aggregates that data into a map.
func ItemInfo(item custom.LiteDriveItemable) map[string]any {
	m := map[string]any{}

	creator := item.GetCreatedByUser()
	if creator != nil {
		m[fault.AddtlCreatedBy] = ptr.Val(creator.GetId())
	}

	lastmodder := item.GetLastModifiedByUser()
	if lastmodder != nil {
		m[fault.AddtlLastModBy] = ptr.Val(lastmodder.GetId())
	}

	parent := item.GetParentReference()
	if parent != nil {
		m[fault.AddtlContainerID] = ptr.Val(parent.GetId())
		m[fault.AddtlContainerName] = ptr.Val(parent.GetName())
		containerPath := ""

		// Remove the "/drives/b!vF-sdsdsds-sdsdsa-sdsd/root:" prefix
		splitPath := strings.SplitN(ptr.Val(parent.GetPath()), ":", 2)
		if len(splitPath) > 1 {
			containerPath = splitPath[1]
		}

		m[fault.AddtlContainerPath] = containerPath
	}

	malware := item.GetMalware()
	if malware != nil {
		m[fault.AddtlMalwareDesc] = ptr.Val(malware.GetDescription())
	}

	return m
}

// ---------------------------------------------------------------------------
// other helpers
// ---------------------------------------------------------------------------

// JWTQueryParam is a query param embed in graph download URLs which holds
// JWT token.
const JWTQueryParam = "tempauth"

// IsURLExpired inspects the jwt token embed in the item download url
// and returns true if it is expired.
func IsURLExpired(
	ctx context.Context,
	urlStr string,
) (
	expiredErr error,
	err error,
) {
	// Extract the raw JWT string from the download url.
	rawJWT, err := common.GetQueryParamFromURL(urlStr, JWTQueryParam)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "jwt query param not found")
	}

	expired, err := jwt.IsJWTExpired(rawJWT)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "checking jwt expiry")
	}

	if expired {
		return clues.StackWC(ctx, ErrTokenExpired), nil
	}

	return nil, nil
}
