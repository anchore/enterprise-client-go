# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKubernetesNamespaces**](DefaultApi.md#DeleteKubernetesNamespaces) | **Delete** /kubernetes-namespaces | Delete Kubernetes namespaces for a given criteria
[**GetKubernetesContainers**](DefaultApi.md#GetKubernetesContainers) | **Get** /kubernetes-containers | Return a list of Kubernetes containers that have been inventoried for this account
[**GetKubernetesNamespaces**](DefaultApi.md#GetKubernetesNamespaces) | **Get** /kubernetes-namespaces | Return a list of Kubernetes namespaces that have been inventoried for this account
[**GetKubernetesNodes**](DefaultApi.md#GetKubernetesNodes) | **Get** /kubernetes-nodes | Return a list of Kubernetes nodes that have been inventoried for this account
[**GetOauthToken**](DefaultApi.md#GetOauthToken) | **Post** /oauth/token | 
[**HealthCheck**](DefaultApi.md#HealthCheck) | **Get** /health | 
[**ListFileContentSearchResults**](DefaultApi.md#ListFileContentSearchResults) | **Get** /images/{image_digest}/artifacts/file_content_search | Return a list of analyzer artifacts of the specified type
[**ListRetrievedFiles**](DefaultApi.md#ListRetrievedFiles) | **Get** /images/{image_digest}/artifacts/retrieved_files | Return a list of analyzer artifacts of the specified type
[**ListSecretSearchResults**](DefaultApi.md#ListSecretSearchResults) | **Get** /images/{image_digest}/artifacts/secret_search | Return a list of analyzer artifacts of the specified type
[**Ping**](DefaultApi.md#Ping) | **Get** / | 
[**RevokeOauthToken**](DefaultApi.md#RevokeOauthToken) | **Post** /oauth/revoke | 
[**VersionCheck**](DefaultApi.md#VersionCheck) | **Get** /version | 



## DeleteKubernetesNamespaces

> DeleteKubernetesNamespaces(ctx).ClusterName(clusterName).Execute()

Delete Kubernetes namespaces for a given criteria



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
    clusterName := "clusterName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteKubernetesNamespaces(context.Background()).ClusterName(clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteKubernetesNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKubernetesNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterName** | **string** |  | 

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


## GetKubernetesContainers

> KubernetesContainers GetKubernetesContainers(ctx).Execute()

Return a list of Kubernetes containers that have been inventoried for this account



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetKubernetesContainers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetKubernetesContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesContainers`: KubernetesContainers
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetKubernetesContainers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesContainersRequest struct via the builder pattern


### Return type

[**KubernetesContainers**](KubernetesContainers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesNamespaces

> KubernetesNamespaces GetKubernetesNamespaces(ctx).Execute()

Return a list of Kubernetes namespaces that have been inventoried for this account



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetKubernetesNamespaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetKubernetesNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesNamespaces`: KubernetesNamespaces
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetKubernetesNamespaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesNamespacesRequest struct via the builder pattern


### Return type

[**KubernetesNamespaces**](KubernetesNamespaces.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesNodes

> KubernetesNodes GetKubernetesNodes(ctx).Execute()

Return a list of Kubernetes nodes that have been inventoried for this account



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetKubernetesNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetKubernetesNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesNodes`: KubernetesNodes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetKubernetesNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesNodesRequest struct via the builder pattern


### Return type

[**KubernetesNodes**](KubernetesNodes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthToken

> TokenResponse GetOauthToken(ctx).GrantType(grantType).Username(username).Password(password).ClientId(clientId).RefreshToken(refreshToken).Execute()





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
    grantType := "grantType_example" // string | OAuth Grant type for token (optional) (default to "password")
    username := "username_example" // string | User to assign OAuth token to (optional)
    password := "password_example" // string | Password for corresponding user (optional)
    clientId := "clientId_example" // string | The type of client used for the OAuth token (optional) (default to "anonymous")
    refreshToken := "refreshToken_example" // string | The refresh token from a previous password grant request, used to get a new access_token (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetOauthToken(context.Background()).GrantType(grantType).Username(username).Password(password).ClientId(clientId).RefreshToken(refreshToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOauthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthToken`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOauthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | OAuth Grant type for token | [default to &quot;password&quot;]
 **username** | **string** | User to assign OAuth token to | 
 **password** | **string** | Password for corresponding user | 
 **clientId** | **string** | The type of client used for the OAuth token | [default to &quot;anonymous&quot;]
 **refreshToken** | **string** | The refresh token from a previous password grant request, used to get a new access_token | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.HealthCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.HealthCheck``: %v\n", err)
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


## ListFileContentSearchResults

> []FileContentSearchResult ListFileContentSearchResults(ctx, imageDigest).Execute()

Return a list of analyzer artifacts of the specified type

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
    imageDigest := "imageDigest_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListFileContentSearchResults(context.Background(), imageDigest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFileContentSearchResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFileContentSearchResults`: []FileContentSearchResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFileContentSearchResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFileContentSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FileContentSearchResult**](FileContentSearchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRetrievedFiles

> []RetrievedFile ListRetrievedFiles(ctx, imageDigest).Execute()

Return a list of analyzer artifacts of the specified type

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
    imageDigest := "imageDigest_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRetrievedFiles(context.Background(), imageDigest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRetrievedFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRetrievedFiles`: []RetrievedFile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRetrievedFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRetrievedFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RetrievedFile**](RetrievedFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecretSearchResults

> []SecretSearchResult ListSecretSearchResults(ctx, imageDigest).Execute()

Return a list of analyzer artifacts of the specified type

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
    imageDigest := "imageDigest_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSecretSearchResults(context.Background(), imageDigest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSecretSearchResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecretSearchResults`: []SecretSearchResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSecretSearchResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecretSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SecretSearchResult**](SecretSearchResult.md)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Ping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Ping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Ping`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Ping`: %v\n", resp)
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


## RevokeOauthToken

> RevokeOauthToken(ctx).Token(token).TokenTypeHint(tokenTypeHint).Execute()





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
    token := "token_example" // string | The token to be revoked (optional)
    tokenTypeHint := "tokenTypeHint_example" // string | A hint about the type of token to be revoked (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RevokeOauthToken(context.Background()).Token(token).TokenTypeHint(tokenTypeHint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RevokeOauthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeOauthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The token to be revoked | 
 **tokenTypeHint** | **string** | A hint about the type of token to be revoked | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.VersionCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VersionCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VersionCheck`: ServiceVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VersionCheck`: %v\n", resp)
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

