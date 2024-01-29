# \StatisticsApi

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemStatistics**](StatisticsApi.md#GetSystemStatistics) | **Get** /system/statistics | List System Statistics



## GetSystemStatistics

> SystemStatisticsList GetSystemStatistics(ctx).Execute()

List System Statistics



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticsApi.GetSystemStatistics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.GetSystemStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemStatistics`: SystemStatisticsList
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.GetSystemStatistics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemStatisticsRequest struct via the builder pattern


### Return type

[**SystemStatisticsList**](SystemStatisticsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

