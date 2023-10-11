# \SystemAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFeed**](SystemAPI.md#DeleteFeed) | **Delete** /system/feeds/{feed} | 
[**DeleteService**](SystemAPI.md#DeleteService) | **Delete** /system/services/{service_name}/{host_id} | Delete the service config
[**DescribeErrorCodes**](SystemAPI.md#DescribeErrorCodes) | **Get** /system/error-codes | Describe anchore engine error codes.
[**DescribePolicy**](SystemAPI.md#DescribePolicy) | **Get** /system/policy-spec | Describe the policy language spec implemented by this service.
[**GetServiceDetail**](SystemAPI.md#GetServiceDetail) | **Get** /system | System status
[**GetServicesByName**](SystemAPI.md#GetServicesByName) | **Get** /system/services/{service_name} | Get a service configuration and state
[**GetServicesByNameAndHost**](SystemAPI.md#GetServicesByNameAndHost) | **Get** /system/services/{service_name}/{host_id} | Get service config for a specific host
[**GetStatus**](SystemAPI.md#GetStatus) | **Get** /status | Service status
[**GetSystemFeeds**](SystemAPI.md#GetSystemFeeds) | **Get** /system/feeds | list feeds operations and information
[**ListServices**](SystemAPI.md#ListServices) | **Get** /system/services | List system services
[**PostSystemFeeds**](SystemAPI.md#PostSystemFeeds) | **Post** /system/feeds | trigger feeds operations
[**TestWebhook**](SystemAPI.md#TestWebhook) | **Post** /system/webhooks/{webhook_type}/test | Adds the capabilities to test a webhook delivery for the given notification type
[**ToggleFeedEnabled**](SystemAPI.md#ToggleFeedEnabled) | **Put** /system/feeds/{feed} | 



## DeleteFeed

> DeleteFeed(ctx, feed).Execute()





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
    feed := "feed_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SystemAPI.DeleteFeed(context.Background(), feed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.DeleteFeed``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    serviceName := "serviceName_example" // string | 
    hostId := "hostId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SystemAPI.DeleteService(context.Background(), serviceName, hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.DeleteService``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.DescribeErrorCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.DescribeErrorCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeErrorCodes`: []AnchoreErrorCode
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.DescribeErrorCodes`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.DescribePolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.DescribePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribePolicy`: []GateSpec
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.DescribePolicy`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.GetServiceDetail(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetServiceDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDetail`: SystemStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetServiceDetail`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    serviceName := "serviceName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.GetServicesByName(context.Background(), serviceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetServicesByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicesByName`: []Service
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetServicesByName`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    serviceName := "serviceName_example" // string | 
    hostId := "hostId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.GetServicesByNameAndHost(context.Background(), serviceName, hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetServicesByNameAndHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicesByNameAndHost`: []Service
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetServicesByNameAndHost`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.GetStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatus`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetStatus`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.GetSystemFeeds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetSystemFeeds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemFeeds`: []FeedMetadata
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetSystemFeeds`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.ListServices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServices`: []Service
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListServices`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.PostSystemFeeds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.PostSystemFeeds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSystemFeeds`: []FeedSyncResult
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.PostSystemFeeds`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    webhookType := "webhookType_example" // string | The Webhook Type that we should test
    notificationType := "notificationType_example" // string | What kind of Notification to send (optional) (default to "tag_update")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SystemAPI.TestWebhook(context.Background(), webhookType).NotificationType(notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.TestWebhook``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    feed := "feed_example" // string | 
    enabled := true // bool | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.ToggleFeedEnabled(context.Background(), feed).Enabled(enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ToggleFeedEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleFeedEnabled`: FeedMetadata
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ToggleFeedEnabled`: %v\n", resp)
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

