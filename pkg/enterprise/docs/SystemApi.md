# \SystemApi

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFeed**](SystemApi.md#DeleteFeed) | **Delete** /system/feeds/{feed} | 
[**DeleteService**](SystemApi.md#DeleteService) | **Delete** /system/services/{service_name}/{host_id} | Delete the service config
[**DescribeErrorCodes**](SystemApi.md#DescribeErrorCodes) | **Get** /system/error-codes | Describe anchore engine error codes.
[**DescribePolicy**](SystemApi.md#DescribePolicy) | **Get** /system/policy-spec | Describe the policy language spec implemented by this service.
[**GetServiceDetail**](SystemApi.md#GetServiceDetail) | **Get** /system | System status
[**GetServicesByName**](SystemApi.md#GetServicesByName) | **Get** /system/services/{service_name} | Get a service configuration and state
[**GetServicesByNameAndHost**](SystemApi.md#GetServicesByNameAndHost) | **Get** /system/services/{service_name}/{host_id} | Get service config for a specific host
[**GetStatus**](SystemApi.md#GetStatus) | **Get** /status | Service status
[**GetSystemFeeds**](SystemApi.md#GetSystemFeeds) | **Get** /system/feeds | list feeds operations and information
[**ListServices**](SystemApi.md#ListServices) | **Get** /system/services | List system services
[**PostSystemFeeds**](SystemApi.md#PostSystemFeeds) | **Post** /system/feeds | trigger feeds operations
[**TestWebhook**](SystemApi.md#TestWebhook) | **Post** /system/webhooks/{webhook_type}/test | Adds the capabilities to test a webhook delivery for the given notification type
[**ToggleFeedEnabled**](SystemApi.md#ToggleFeedEnabled) | **Put** /system/feeds/{feed} | 



## DeleteFeed

> DeleteFeed(ctx, feed).Execute()





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
    feed := "feed_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteFeed(context.Background(), feed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feed** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteService

> DeleteService(ctx, serviceName, hostId).Execute()

Delete the service config

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
    serviceName := "serviceName_example" // string | 
    hostId := "hostId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.DeleteService(context.Background(), serviceName, hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DeleteService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** |  | 
**hostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeErrorCodes

> []AnchoreErrorCode DescribeErrorCodes(ctx).Execute()

Describe anchore engine error codes.



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
    resp, r, err := apiClient.SystemApi.DescribeErrorCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DescribeErrorCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeErrorCodes`: []AnchoreErrorCode
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.DescribeErrorCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeErrorCodesRequest struct via the builder pattern


### Return type

[**[]AnchoreErrorCode**](AnchoreErrorCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePolicy

> []GateSpec DescribePolicy(ctx).Execute()

Describe the policy language spec implemented by this service.



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
    resp, r, err := apiClient.SystemApi.DescribePolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.DescribePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribePolicy`: []GateSpec
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.DescribePolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDescribePolicyRequest struct via the builder pattern


### Return type

[**[]GateSpec**](GateSpec.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDetail

> SystemStatusResponse GetServiceDetail(ctx).Execute()

System status



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
    resp, r, err := apiClient.SystemApi.GetServiceDetail(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetServiceDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDetail`: SystemStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetServiceDetail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceDetailRequest struct via the builder pattern


### Return type

[**SystemStatusResponse**](SystemStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesByName

> []Service GetServicesByName(ctx, serviceName).Execute()

Get a service configuration and state

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
    serviceName := "serviceName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetServicesByName(context.Background(), serviceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetServicesByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicesByName`: []Service
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetServicesByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesByNameAndHost

> []Service GetServicesByNameAndHost(ctx, serviceName, hostId).Execute()

Get service config for a specific host

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
    serviceName := "serviceName_example" // string | 
    hostId := "hostId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetServicesByNameAndHost(context.Background(), serviceName, hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetServicesByNameAndHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicesByNameAndHost`: []Service
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetServicesByNameAndHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** |  | 
**hostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesByNameAndHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> StatusResponse GetStatus(ctx).Execute()

Service status



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
    resp, r, err := apiClient.SystemApi.GetStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatus`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemFeeds

> []FeedMetadata GetSystemFeeds(ctx).Execute()

list feeds operations and information



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
    resp, r, err := apiClient.SystemApi.GetSystemFeeds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSystemFeeds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemFeeds`: []FeedMetadata
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetSystemFeeds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemFeedsRequest struct via the builder pattern


### Return type

[**[]FeedMetadata**](FeedMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServices

> []Service ListServices(ctx).Execute()

List system services

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
    resp, r, err := apiClient.SystemApi.ListServices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.ListServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServices`: []Service
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.ListServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServicesRequest struct via the builder pattern


### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemFeeds

> []FeedSyncResult PostSystemFeeds(ctx).Execute()

trigger feeds operations



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
    resp, r, err := apiClient.SystemApi.PostSystemFeeds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PostSystemFeeds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSystemFeeds`: []FeedSyncResult
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.PostSystemFeeds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemFeedsRequest struct via the builder pattern


### Return type

[**[]FeedSyncResult**](FeedSyncResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestWebhook

> TestWebhook(ctx, webhookType).NotificationType(notificationType).Execute()

Adds the capabilities to test a webhook delivery for the given notification type



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
    webhookType := "webhookType_example" // string | The Webhook Type that we should test
    notificationType := "notificationType_example" // string | What kind of Notification to send (optional) (default to "tag_update")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.TestWebhook(context.Background(), webhookType).NotificationType(notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.TestWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookType** | **string** | The Webhook Type that we should test | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationType** | **string** | What kind of Notification to send | [default to &quot;tag_update&quot;]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleFeedEnabled

> FeedMetadata ToggleFeedEnabled(ctx, feed).Enabled(enabled).Execute()





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
    feed := "feed_example" // string | 
    enabled := true // bool | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.ToggleFeedEnabled(context.Background(), feed).Enabled(enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.ToggleFeedEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleFeedEnabled`: FeedMetadata
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.ToggleFeedEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feed** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleFeedEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enabled** | **bool** |  | 

### Return type

[**FeedMetadata**](FeedMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

