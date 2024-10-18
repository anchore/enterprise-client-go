# \SystemConfigurationAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration**](SystemConfigurationAPI.md#DeleteConfiguration) | **Delete** /system/configurations/{config_key} | Unset a configuration key
[**GetConfiguration**](SystemConfigurationAPI.md#GetConfiguration) | **Get** /system/configurations/{config_key} | Get an individual configuration
[**GetConfigurationSchema**](SystemConfigurationAPI.md#GetConfigurationSchema) | **Get** /system/configurations/schema | Get the system configuration schema
[**ListConfigurations**](SystemConfigurationAPI.md#ListConfigurations) | **Get** /system/configurations | List all API exposed configurations and their values
[**PatchConfiguration**](SystemConfigurationAPI.md#PatchConfiguration) | **Patch** /system/configurations | Apply a patch to the system configuration
[**PutConfiguration**](SystemConfigurationAPI.md#PutConfiguration) | **Put** /system/configurations/{config_key} | Put a single configuration value



## DeleteConfiguration

> DeleteConfiguration(ctx, configKey).Execute()

Unset a configuration key



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
	configKey := "configKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemConfigurationAPI.DeleteConfiguration(context.Background(), configKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemConfigurationAPI.DeleteConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationRequest struct via the builder pattern


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


## GetConfiguration

> SystemConfiguration GetConfiguration(ctx, configKey).Execute()

Get an individual configuration



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
	configKey := "configKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemConfigurationAPI.GetConfiguration(context.Background(), configKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemConfigurationAPI.GetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration`: SystemConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SystemConfigurationAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemConfiguration**](SystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationSchema

> interface{} GetConfigurationSchema(ctx).Execute()

Get the system configuration schema



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
	resp, r, err := apiClient.SystemConfigurationAPI.GetConfigurationSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemConfigurationAPI.GetConfigurationSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationSchema`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SystemConfigurationAPI.GetConfigurationSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationSchemaRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigurations

> SystemConfigurationList ListConfigurations(ctx).Execute()

List all API exposed configurations and their values



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
	resp, r, err := apiClient.SystemConfigurationAPI.ListConfigurations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemConfigurationAPI.ListConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConfigurations`: SystemConfigurationList
	fmt.Fprintf(os.Stdout, "Response from `SystemConfigurationAPI.ListConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConfigurationsRequest struct via the builder pattern


### Return type

[**SystemConfigurationList**](SystemConfigurationList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConfiguration

> SystemConfigurationList PatchConfiguration(ctx).SystemConfigurationPatchInner(systemConfigurationPatchInner).Execute()

Apply a patch to the system configuration



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
	systemConfigurationPatchInner := []openapiclient.SystemConfigurationPatchInner{*openapiclient.NewSystemConfigurationPatchInner("Key_example", openapiclient.SystemConfiguration_value{ArrayOfAny: new([]Any)})} // []SystemConfigurationPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemConfigurationAPI.PatchConfiguration(context.Background()).SystemConfigurationPatchInner(systemConfigurationPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemConfigurationAPI.PatchConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchConfiguration`: SystemConfigurationList
	fmt.Fprintf(os.Stdout, "Response from `SystemConfigurationAPI.PatchConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemConfigurationPatchInner** | [**[]SystemConfigurationPatchInner**](SystemConfigurationPatchInner.md) |  | 

### Return type

[**SystemConfigurationList**](SystemConfigurationList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConfiguration

> SystemConfiguration PutConfiguration(ctx, configKey).SystemConfigurationPut(systemConfigurationPut).Execute()

Put a single configuration value



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
	configKey := "configKey_example" // string | 
	systemConfigurationPut := *openapiclient.NewSystemConfigurationPut("Key_example", openapiclient.SystemConfiguration_value{ArrayOfAny: new([]Any)}) // SystemConfigurationPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemConfigurationAPI.PutConfiguration(context.Background(), configKey).SystemConfigurationPut(systemConfigurationPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemConfigurationAPI.PutConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutConfiguration`: SystemConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SystemConfigurationAPI.PutConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemConfigurationPut** | [**SystemConfigurationPut**](SystemConfigurationPut.md) |  | 

### Return type

[**SystemConfiguration**](SystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

