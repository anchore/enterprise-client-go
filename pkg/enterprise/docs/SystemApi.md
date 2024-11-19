# \SystemAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFeed**](SystemAPI.md#DeleteFeed) | **Delete** /system/feeds/{feed} | 
[**DeleteService**](SystemAPI.md#DeleteService) | **Delete** /system/services/{service_name}/{host_id} | Delete the service config
[**DescribeErrorCodes**](SystemAPI.md#DescribeErrorCodes) | **Get** /system/error-codes | Describe anchore engine error codes.
[**DescribePolicy**](SystemAPI.md#DescribePolicy) | **Get** /system/policy-spec | Describe the policy language spec implemented by this service.
[**GetAnchorectl**](SystemAPI.md#GetAnchorectl) | **Get** /system/anchorectl | Get an anchorectl binary compatible with this version of Anchore Enterprise
[**GetServiceDetail**](SystemAPI.md#GetServiceDetail) | **Get** /system | System status
[**GetServicesByName**](SystemAPI.md#GetServicesByName) | **Get** /system/services/{service_name} | Get a service configuration and state
[**GetServicesByNameAndHost**](SystemAPI.md#GetServicesByNameAndHost) | **Get** /system/services/{service_name}/{host_id} | Get service config for a specific host
[**GetStatus**](SystemAPI.md#GetStatus) | **Get** /status | Service status
[**GetSystemFeed**](SystemAPI.md#GetSystemFeed) | **Get** /system/feeds/{feed} | 
[**GetSystemFeeds**](SystemAPI.md#GetSystemFeeds) | **Get** /system/feeds | list feeds operations and information
[**HealthCheck**](SystemAPI.md#HealthCheck) | **Get** /health | 
[**ListServices**](SystemAPI.md#ListServices) | **Get** /system/services | List system services
[**Ping**](SystemAPI.md#Ping) | **Get** / | 
[**PostSystemFeeds**](SystemAPI.md#PostSystemFeeds) | **Post** /system/feeds | trigger feeds operations
[**SetNewLogLevel**](SystemAPI.md#SetNewLogLevel) | **Post** /system/logging | Change logging level for a running service
[**TestWebhook**](SystemAPI.md#TestWebhook) | **Post** /system/webhooks/{webhook_type}/test | Adds the capabilities to test a webhook delivery for the given notification type
[**ToggleFeedEnabled**](SystemAPI.md#ToggleFeedEnabled) | **Put** /system/feeds/{feed} | 
[**UploadSystemFeed**](SystemAPI.md#UploadSystemFeed) | **Post** /system/feeds/{feed} | 
[**VersionCheck**](SystemAPI.md#VersionCheck) | **Get** /version | 



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


## GetAnchorectl

> GetAnchorectl(ctx).OperatingSystem(operatingSystem).Architecture(architecture).Execute()

Get an anchorectl binary compatible with this version of Anchore Enterprise

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
	operatingSystem := "operatingSystem_example" // string | The operating system (platform) of the binary to retrieve
	architecture := "architecture_example" // string | The architecture of the binary to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemAPI.GetAnchorectl(context.Background()).OperatingSystem(operatingSystem).Architecture(architecture).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetAnchorectl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnchorectlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operatingSystem** | **string** | The operating system (platform) of the binary to retrieve | 
 **architecture** | **string** | The architecture of the binary to retrieve | 

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


## GetSystemFeed

> FeedDataRecord GetSystemFeed(ctx, feed).Version(version).Execute()





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
	feed := "feed_example" // string | The data feed to query
	version := "version_example" // string | The version of the feed to query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.GetSystemFeed(context.Background(), feed).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetSystemFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemFeed`: FeedDataRecord
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetSystemFeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feed** | **string** | The data feed to query | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The version of the feed to query | 

### Return type

[**FeedDataRecord**](FeedDataRecord.md)

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


## HealthCheck

> HealthCheck(ctx).Execute()





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
	r, err := apiClient.SystemAPI.HealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.HealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthCheckRequest struct via the builder pattern


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


## Ping

> string Ping(ctx).Execute()





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
	resp, r, err := apiClient.SystemAPI.Ping(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.Ping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Ping`: string
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemFeeds

> PostSystemFeeds200Response PostSystemFeeds(ctx).ForceSync(forceSync).Execute()

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
	forceSync := true // bool | If set this will cause the data-syncer to ignore checksum matches and force a resync (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.PostSystemFeeds(context.Background()).ForceSync(forceSync).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.PostSystemFeeds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemFeeds`: PostSystemFeeds200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.PostSystemFeeds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemFeedsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceSync** | **bool** | If set this will cause the data-syncer to ignore checksum matches and force a resync | 

### Return type

[**PostSystemFeeds200Response**](PostSystemFeeds200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNewLogLevel

> []LoggingLevel SetNewLogLevel(ctx).Logging(logging).Execute()

Change logging level for a running service



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
	logging := *openapiclient.NewLoggingLevel("ServiceName_example", "LoggingLevel_example") // LoggingLevel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.SetNewLogLevel(context.Background()).Logging(logging).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.SetNewLogLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNewLogLevel`: []LoggingLevel
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.SetNewLogLevel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetNewLogLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logging** | [**LoggingLevel**](LoggingLevel.md) |  | 

### Return type

[**[]LoggingLevel**](LoggingLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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


## UploadSystemFeed

> UploadSystemFeed(ctx, feed).Version(version).Built(built).Checksum(checksum).File(file).Execute()





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
	feed := "feed_example" // string | The data feed to upload
	version := "version_example" // string | The version of the data set to upload
	built := "built_example" // string | Build timestsamp
	checksum := "checksum_example" // string | Digest of the data record
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemAPI.UploadSystemFeed(context.Background(), feed).Version(version).Built(built).Checksum(checksum).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.UploadSystemFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feed** | **string** | The data feed to upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSystemFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The version of the data set to upload | 
 **built** | **string** | Build timestsamp | 
 **checksum** | **string** | Digest of the data record | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionCheck

> ServiceVersion VersionCheck(ctx).Execute()





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
	resp, r, err := apiClient.SystemAPI.VersionCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.VersionCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VersionCheck`: ServiceVersion
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.VersionCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionCheckRequest struct via the builder pattern


### Return type

[**ServiceVersion**](ServiceVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

