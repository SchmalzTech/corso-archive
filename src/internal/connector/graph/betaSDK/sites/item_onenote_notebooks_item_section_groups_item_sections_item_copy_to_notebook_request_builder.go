package sites

import (
    "context"
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6 "betasdk/models/odataerrors"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilder provides operations to call the copyToNotebook method.
type ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilderInternal instantiates a new CopyToNotebookRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilder) {
    m := &ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/microsoft.graph.copyToNotebook";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilder instantiates a new CopyToNotebookRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation copies a section to a specific notebook. For Copy operations, you follow an asynchronous calling pattern:  First call the Copy action, and then poll the operation endpoint for the result.
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBodyable, requestConfiguration *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post copies a section to a specific notebook. For Copy operations, you follow an asynchronous calling pattern:  First call the Copy action, and then poll the operation endpoint for the result.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/section-copytonotebook?view=graph-rest-1.0
func (m *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilder) Post(ctx context.Context, body ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBodyable, requestConfiguration *ItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilderPostRequestConfiguration)(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.OnenoteOperationable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
        "5XX": i2611c67443a66a7e6664535e21478294fb96d6d29c44551db4d04f63a0af61d6.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateOnenoteOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.OnenoteOperationable), nil
}
