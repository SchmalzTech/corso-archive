package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models/odataerrors"
    i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models/termstore"
)

// ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
type ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters to indicate which terms are related to the current term as either pinned or reused.
type ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters
}
// ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderInternal instantiates a new RelationItemRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) {
    m := &ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStore/groups/{group%2Did}/sets/{set%2Did}/terms/{term%2Did}/children/{term%2Did1}/relations/{relation%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder instantiates a new RelationItemRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property relations for sites
func (m *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation to indicate which terms are related to the current term as either pinned or reused.
func (m *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property relations in sites
func (m *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Relationable, requestConfiguration *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property relations for sites
func (m *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// FromTerm provides operations to manage the fromTerm property of the microsoft.graph.termStore.relation entity.
func (m *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) FromTerm()(*ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get to indicate which terms are related to the current term as either pinned or reused.
func (m *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration)(i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Relationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.CreateRelationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Relationable), nil
}
// Patch update the navigation property relations in sites
func (m *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) Patch(ctx context.Context, body i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Relationable, requestConfiguration *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration)(i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Relationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.CreateRelationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i9f80f9c244f49392da487c12fb13b03692a695949f8ff3e2d6cdb7662a064016.Relationable), nil
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.relation entity.
func (m *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) Set()(*ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemSetRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToTerm provides operations to manage the toTerm property of the microsoft.graph.termStore.relation entity.
func (m *ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsRelationItemRequestBuilder) ToTerm()(*ItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemToTermRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsItemTermsItemChildrenItemRelationsItemToTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
