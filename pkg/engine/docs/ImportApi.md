# \ImportApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportImageArchive**](ImportApi.md#ImportImageArchive) | **Post** /import/images | Import an anchore image tar.gz archive file. This is a deprecated API replaced by the \&quot;/imports/images\&quot; route



## ImportImageArchive

> []AnchoreImage ImportImageArchive(ctx).ArchiveFile(archiveFile).Execute()

Import an anchore image tar.gz archive file. This is a deprecated API replaced by the \"/imports/images\" route

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
    archiveFile := os.NewFile(1234, "some_file") // *os.File | anchore image tar archive.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportApi.ImportImageArchive(context.Background()).ArchiveFile(archiveFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportApi.ImportImageArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImageArchive`: []AnchoreImage
    fmt.Fprintf(os.Stdout, "Response from `ImportApi.ImportImageArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportImageArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archiveFile** | ***os.File** | anchore image tar archive. | 

### Return type

[**[]AnchoreImage**](AnchoreImage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

