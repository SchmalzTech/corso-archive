package groups

import (
    "context"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder provides operations to call the restore method.
type ItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderInternal instantiates a new RestoreRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder) {
    m := &ItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/documentSetVersions/{documentSetVersion%2Did}/microsoft.graph.restore";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder instantiates a new RestoreRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation restore a document set version.
func (m *ItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post restore a document set version.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/documentsetversion-restore?view=graph-rest-1.0
func (m *ItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
