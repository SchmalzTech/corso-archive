package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
    i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models/odataerrors"
)

// ItemListsItemContentTypesRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ItemListsItemContentTypesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemListsItemContentTypesRequestBuilderGetQueryParameters get the collection of [contentType][contentType] resources in a [list][].
type ItemListsItemContentTypesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemListsItemContentTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContentTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListsItemContentTypesRequestBuilderGetQueryParameters
}
// ItemListsItemContentTypesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContentTypesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddCopy provides operations to call the addCopy method.
func (m *ItemListsItemContentTypesRequestBuilder) AddCopy()(*ItemListsItemContentTypesAddCopyRequestBuilder) {
    return NewItemListsItemContentTypesAddCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddCopyFromContentTypeHub provides operations to call the addCopyFromContentTypeHub method.
func (m *ItemListsItemContentTypesRequestBuilder) AddCopyFromContentTypeHub()(*ItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder) {
    return NewItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemListsItemContentTypesRequestBuilderInternal instantiates a new ContentTypesRequestBuilder and sets the default values.
func NewItemListsItemContentTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContentTypesRequestBuilder) {
    m := &ItemListsItemContentTypesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/contentTypes{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemListsItemContentTypesRequestBuilder instantiates a new ContentTypesRequestBuilder and sets the default values.
func NewItemListsItemContentTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContentTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemContentTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemListsItemContentTypesRequestBuilder) Count()(*ItemListsItemContentTypesCountRequestBuilder) {
    return NewItemListsItemContentTypesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get the collection of [contentType][contentType] resources in a [list][].
func (m *ItemListsItemContentTypesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemContentTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to contentTypes for sites
func (m *ItemListsItemContentTypesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ContentTypeable, requestConfiguration *ItemListsItemContentTypesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get the collection of [contentType][contentType] resources in a [list][].
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/list-list-contenttypes?view=graph-rest-1.0
func (m *ItemListsItemContentTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemContentTypesRequestBuilderGetRequestConfiguration)(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ContentTypeCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateContentTypeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ContentTypeCollectionResponseable), nil
}
// GetCompatibleHubContentTypes provides operations to call the getCompatibleHubContentTypes method.
func (m *ItemListsItemContentTypesRequestBuilder) GetCompatibleHubContentTypes()(*ItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilder) {
    return NewItemListsItemContentTypesGetCompatibleHubContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to contentTypes for sites
func (m *ItemListsItemContentTypesRequestBuilder) Post(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ContentTypeable, requestConfiguration *ItemListsItemContentTypesRequestBuilderPostRequestConfiguration)(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ContentTypeable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ContentTypeable), nil
}
