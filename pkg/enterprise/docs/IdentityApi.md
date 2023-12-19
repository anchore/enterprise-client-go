# \IdentityApi

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCredential**](IdentityApi.md#AddCredential) | **Post** /user/credentials | add/replace credential
[**CreateApiKey**](IdentityApi.md#CreateApiKey) | **Post** /user/api-keys | Add a new API key
[**DeleteApiKey**](IdentityApi.md#DeleteApiKey) | **Delete** /user/api-keys/{key_name} | Delete a user API key
[**GetApiKey**](IdentityApi.md#GetApiKey) | **Get** /user/api-keys/{key_name} | Get a user API key
[**GetCredentials**](IdentityApi.md#GetCredentials) | **Get** /user/credentials | Get current credential summary
[**GetUser**](IdentityApi.md#GetUser) | **Get** /user | List authenticated user info
[**GetUsersAccount**](IdentityApi.md#GetUsersAccount) | **Get** /account | List the account for the authenticated user
[**ListApiKeys**](IdentityApi.md#ListApiKeys) | **Get** /user/api-keys | Get a list of API keys
[**PatchApiKey**](IdentityApi.md#PatchApiKey) | **Patch** /user/api-keys/{key_name} | Patch a user API key



## AddCredential

> AccessCredential AddCredential(ctx).Credential(credential).Execute()

add/replace credential

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
    credential := *openapiclient.NewAccessCredential("Type_example", "Value_example") // AccessCredential | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityApi.AddCredential(context.Background()).Credential(credential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.AddCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCredential`: AccessCredential
    fmt.Fprintf(os.Stdout, "Response from `IdentityApi.AddCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credential** | [**AccessCredential**](AccessCredential.md) |  | 

### Return type

[**AccessCredential**](AccessCredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiKey

> UserApiKey CreateApiKey(ctx).UserApiKey(userApiKey).Execute()

Add a new API key

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
    userApiKey := *openapiclient.NewUserApiKey("Name_example", time.Now()) // UserApiKey | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityApi.CreateApiKey(context.Background()).UserApiKey(userApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.CreateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiKey`: UserApiKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityApi.CreateApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userApiKey** | [**UserApiKey**](UserApiKey.md) |  | 

### Return type

[**UserApiKey**](UserApiKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKey

> DeleteApiKey(ctx, keyName).Execute()

Delete a user API key

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
    keyName := "keyName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityApi.DeleteApiKey(context.Background(), keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyRequest struct via the builder pattern


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


## GetApiKey

> UserApiKey GetApiKey(ctx, keyName).Execute()

Get a user API key

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
    keyName := "keyName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityApi.GetApiKey(context.Background(), keyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKey`: UserApiKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityApi.GetApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserApiKey**](UserApiKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentials

> []AccessCredential GetCredentials(ctx).Execute()

Get current credential summary

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
    resp, r, err := apiClient.IdentityApi.GetCredentials(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCredentials`: []AccessCredential
    fmt.Fprintf(os.Stdout, "Response from `IdentityApi.GetCredentials`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialsRequest struct via the builder pattern


### Return type

[**[]AccessCredential**](AccessCredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx).Execute()

List authenticated user info

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
    resp, r, err := apiClient.IdentityApi.GetUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: User
    fmt.Fprintf(os.Stdout, "Response from `IdentityApi.GetUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersAccount

> Account GetUsersAccount(ctx).Execute()

List the account for the authenticated user

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
    resp, r, err := apiClient.IdentityApi.GetUsersAccount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetUsersAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `IdentityApi.GetUsersAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersAccountRequest struct via the builder pattern


### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiKeys

> ApiKeyList ListApiKeys(ctx).Execute()

Get a list of API keys

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
    resp, r, err := apiClient.IdentityApi.ListApiKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.ListApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiKeys`: ApiKeyList
    fmt.Fprintf(os.Stdout, "Response from `IdentityApi.ListApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiKeysRequest struct via the builder pattern


### Return type

[**ApiKeyList**](ApiKeyList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiKey

> UserApiKey PatchApiKey(ctx, keyName).PatchUserApiKeyRequest(patchUserApiKeyRequest).Execute()

Patch a user API key

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
    keyName := "keyName_example" // string | 
    patchUserApiKeyRequest := *openapiclient.NewPatchUserApiKeyRequest() // PatchUserApiKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityApi.PatchApiKey(context.Background(), keyName).PatchUserApiKeyRequest(patchUserApiKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PatchApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchApiKey`: UserApiKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityApi.PatchApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchUserApiKeyRequest** | [**PatchUserApiKeyRequest**](PatchUserApiKeyRequest.md) |  | 

### Return type

[**UserApiKey**](UserApiKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

