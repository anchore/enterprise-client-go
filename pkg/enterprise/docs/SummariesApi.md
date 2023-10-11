# \SummariesAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListImageTags**](SummariesAPI.md#ListImageTags) | **Get** /summaries/image-tags | List all visible image digests and tags



## ListImageTags

> AnchoreImageTagSummaryList ListImageTags(ctx).ImageStatus(imageStatus).XAnchoreAccount(xAnchoreAccount).Execute()

List all visible image digests and tags



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    imageStatus := []string{"ImageStatus_example"} // []string | Filter images in one or more states such as active, deleting. Defaults to active images only if unspecified (optional) (default to ["active"])
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SummariesAPI.ListImageTags(context.Background()).ImageStatus(imageStatus).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SummariesAPI.ListImageTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImageTags`: AnchoreImageTagSummaryList
    fmt.Fprintf(os.Stdout, "Response from `SummariesAPI.ListImageTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImageTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageStatus** | **[]string** | Filter images in one or more states such as active, deleting. Defaults to active images only if unspecified | [default to [&quot;active&quot;]]
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**AnchoreImageTagSummaryList**](AnchoreImageTagSummaryList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

