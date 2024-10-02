# \IntegrationsApi

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIntegration**](IntegrationsApi.md#DeleteIntegration) | **Delete** /system/integrations/{integration_uuid} | Delete an integration instance
[**GetIntegrationById**](IntegrationsApi.md#GetIntegrationById) | **Get** /system/integrations/{integration_uuid} | Get information about an integration instance
[**HandleHealthReport**](IntegrationsApi.md#HandleHealthReport) | **Post** /system/integrations/{integration_uuid}/health-report | Report health status for an integration
[**ListIntegrations**](IntegrationsApi.md#ListIntegrations) | **Get** /system/integrations | List known integration instances
[**RegisterIntegration**](IntegrationsApi.md#RegisterIntegration) | **Post** /system/integrations/registration | Register an integration instance



## DeleteIntegration

> DeleteIntegration(ctx, integrationUuid).Force(force).Execute()

Delete an integration instance

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
    integrationUuid := "84993c1f-863f-41f3-9bd8-dfcc821b1c8a" // string | The integration instance's universally unique identifier
    force := true // bool | Force deletion of the integration instance regardless of its state (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.DeleteIntegration(context.Background(), integrationUuid).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.DeleteIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationUuid** | **string** | The integration instance&#39;s universally unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force deletion of the integration instance regardless of its state | [default to false]

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


## GetIntegrationById

> Integration GetIntegrationById(ctx, integrationUuid).Execute()

Get information about an integration instance

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
    integrationUuid := "84993c1f-863f-41f3-9bd8-dfcc821b1c8a" // string | The integration instance's universally unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.GetIntegrationById(context.Background(), integrationUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetIntegrationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationById`: Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetIntegrationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationUuid** | **string** | The integration instance&#39;s universally unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleHealthReport

> HandleHealthReport(ctx, integrationUuid).HealthReport(healthReport).Execute()

Report health status for an integration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    integrationUuid := "84993c1f-863f-41f3-9bd8-dfcc821b1c8a" // string | The integration's universally unique identifier
    healthReport := *openapiclient.NewHealthReport("740c06a3-4c69-4b91-8e85-154cd53e9764", int32(1), time.Now(), float32(200.02312), int32(60), *openapiclient.NewHealthData(openapiclient.IntegrationType("anchore_k8s_inventory_agent"), int32(1))) // HealthReport | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.HandleHealthReport(context.Background(), integrationUuid).HealthReport(healthReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.HandleHealthReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationUuid** | **string** | The integration&#39;s universally unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleHealthReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **healthReport** | [**HealthReport**](HealthReport.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrations

> IntegrationListResponse ListIntegrations(ctx).Execute()

List known integration instances



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
    resp, r, err := apiClient.IntegrationsApi.ListIntegrations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrations`: IntegrationListResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationsRequest struct via the builder pattern


### Return type

[**IntegrationListResponse**](IntegrationListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterIntegration

> Integration RegisterIntegration(ctx).IntegrationRegister(integrationRegister).Execute()

Register an integration instance



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
    integrationRegister := *openapiclient.NewIntegrationRegister("84993c1f-863f-41f3-9bd8-dfcc821b1c8a", "67479449b7-zdzlr", openapiclient.IntegrationType("anchore_k8s_inventory_agent"), "admin", int32(60)) // IntegrationRegister | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.RegisterIntegration(context.Background()).IntegrationRegister(integrationRegister).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.RegisterIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterIntegration`: Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.RegisterIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationRegister** | [**IntegrationRegister**](IntegrationRegister.md) |  | 

### Return type

[**Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

