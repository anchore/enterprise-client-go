# \ReportsApi

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalQueryResult**](ReportsApi.md#GetGlobalQueryResult) | **Get** /reporting/reports/global/scheduled-query-results/{result_uuid} | 
[**GetQueryResult**](ReportsApi.md#GetQueryResult) | **Get** /reporting/scheduled-query-results/{result_uuid} | 



## GetGlobalQueryResult

> GetGlobalQueryResult(ctx, resultUuid).Page(page).Execute()





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
    resultUuid := "resultUuid_example" // string | 
    page := int32(56) // int32 | Page number to fetch. If omitted, '1' is default. Page numbers start at 1 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.GetGlobalQueryResult(context.Background(), resultUuid).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetGlobalQueryResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalQueryResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number to fetch. If omitted, &#39;1&#39; is default. Page numbers start at 1 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueryResult

> GetQueryResult(ctx, resultUuid).Page(page).Execute()





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
    resultUuid := "resultUuid_example" // string | 
    page := int32(56) // int32 | Page number to fetch. If omitted, '1' is default. Page numbers start at 1 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.GetQueryResult(context.Background(), resultUuid).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetQueryResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number to fetch. If omitted, &#39;1&#39; is default. Page numbers start at 1 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

