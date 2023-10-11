# \RegistriesAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegistry**](RegistriesAPI.md#CreateRegistry) | **Post** /registries | Add a new registry
[**DeleteRegistry**](RegistriesAPI.md#DeleteRegistry) | **Delete** /registries/{registry} | Delete a registry configuration
[**GetRegistry**](RegistriesAPI.md#GetRegistry) | **Get** /registries/{registry} | Get a specific registry configuration
[**ListRegistries**](RegistriesAPI.md#ListRegistries) | **Get** /registries | List configured registries
[**UpdateRegistry**](RegistriesAPI.md#UpdateRegistry) | **Put** /registries/{registry} | Update/replace a registry configuration



## CreateRegistry

> []RegistryConfiguration CreateRegistry(ctx).RegistryData(registryData).Validate(validate).XAnchoreAccount(xAnchoreAccount).Execute()

Add a new registry



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
    registryData := *openapiclient.NewRegistryConfigurationRequest() // RegistryConfigurationRequest | 
    validate := true // bool | flag to determine whether or not to validate registry/credential at registry add time (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistriesAPI.CreateRegistry(context.Background()).RegistryData(registryData).Validate(validate).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesAPI.CreateRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegistry`: []RegistryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RegistriesAPI.CreateRegistry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryData** | [**RegistryConfigurationRequest**](RegistryConfigurationRequest.md) |  | 
 **validate** | **bool** | flag to determine whether or not to validate registry/credential at registry add time | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegistry

> DeleteRegistry(ctx, registry).XAnchoreAccount(xAnchoreAccount).Execute()

Delete a registry configuration



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
    registry := "registry_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RegistriesAPI.DeleteRegistry(context.Background(), registry).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesAPI.DeleteRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registry** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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


## GetRegistry

> []RegistryConfiguration GetRegistry(ctx, registry).XAnchoreAccount(xAnchoreAccount).Execute()

Get a specific registry configuration



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
    registry := "registry_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistriesAPI.GetRegistry(context.Background(), registry).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesAPI.GetRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistry`: []RegistryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RegistriesAPI.GetRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registry** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegistries

> []RegistryConfiguration ListRegistries(ctx).XAnchoreAccount(xAnchoreAccount).Execute()

List configured registries



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
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistriesAPI.ListRegistries(context.Background()).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesAPI.ListRegistries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRegistries`: []RegistryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RegistriesAPI.ListRegistries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRegistriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegistry

> []RegistryConfiguration UpdateRegistry(ctx, registry).RegistryData(registryData).Validate(validate).XAnchoreAccount(xAnchoreAccount).Execute()

Update/replace a registry configuration



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
    registry := "registry_example" // string | 
    registryData := *openapiclient.NewRegistryConfigurationRequest() // RegistryConfigurationRequest | 
    validate := true // bool | flag to determine whether or not to validate registry/credential at registry update time (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistriesAPI.UpdateRegistry(context.Background(), registry).RegistryData(registryData).Validate(validate).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesAPI.UpdateRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRegistry`: []RegistryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RegistriesAPI.UpdateRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registry** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registryData** | [**RegistryConfigurationRequest**](RegistryConfigurationRequest.md) |  | 
 **validate** | **bool** | flag to determine whether or not to validate registry/credential at registry update time | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

