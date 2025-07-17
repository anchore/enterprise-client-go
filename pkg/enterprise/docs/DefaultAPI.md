# \DefaultAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportImageSbom**](DefaultAPI.md#ImportImageSbom) | **Post** /imports/images/{operation_id}/sbom | Import an AnchoreCTL generated SBOM document into the system



## ImportImageSbom

> ImageImportContentResponse ImportImageSbom(ctx, operationId).Execute()

Import an AnchoreCTL generated SBOM document into the system

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
	operationId := "operationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ImportImageSbom(context.Background(), operationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ImportImageSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportImageSbom`: ImageImportContentResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ImportImageSbom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportImageSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

