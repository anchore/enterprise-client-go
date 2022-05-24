# \SummariesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListImagetags**](SummariesApi.md#ListImagetags) | **Get** /summaries/imagetags | List all visible image digests and tags



## ListImagetags

> []AnchoreImageTagSummary ListImagetags(ctx).ImageStatus(imageStatus).XAnchoreAccount(xAnchoreAccount).Execute()

List all visible image digests and tags



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    imageStatus := []string{"ImageStatus_example"} // []string | Filter images in one or more states such as active, deleting. Defaults to active images only if unspecified (optional) (default to ["active"])
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SummariesApi.ListImagetags(context.Background()).ImageStatus(imageStatus).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SummariesApi.ListImagetags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImagetags`: []AnchoreImageTagSummary
    fmt.Fprintf(os.Stdout, "Response from `SummariesApi.ListImagetags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImagetagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageStatus** | **[]string** | Filter images in one or more states such as active, deleting. Defaults to active images only if unspecified | [default to [&quot;active&quot;]]
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]AnchoreImageTagSummary**](AnchoreImageTagSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

