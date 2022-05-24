# \SourcesApi

All URIs are relative to *http://localhost/enterprise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSource**](SourcesApi.md#DeleteSource) | **Delete** /sources/{source_id} | Delete source record from DB
[**GetSource**](SourcesApi.md#GetSource) | **Get** /sources/{source_id} | Get a detailed source repository analysis metadata record
[**GetSourceContentByType**](SourcesApi.md#GetSourceContentByType) | **Get** /sources/{source_id}/content/{content_type} | Get the content of an analyzed source repository
[**GetSourceContentTypes**](SourcesApi.md#GetSourceContentTypes) | **Get** /sources/{source_id}/content | Get a detailed source repository analysis metadata record
[**GetSourcePolicyCheck**](SourcesApi.md#GetSourcePolicyCheck) | **Get** /sources/{source_id}/check | Fetch or calculate policy evaluation for a source
[**GetSourceSbomNative**](SourcesApi.md#GetSourceSbomNative) | **Get** /sources/{source_id}/sbom/native | 
[**GetSourceSbomTypes**](SourcesApi.md#GetSourceSbomTypes) | **Get** /sources/{source_id}/sbom | Get a detailed source repository analysis metadata record
[**GetSourceVulnerabilities**](SourcesApi.md#GetSourceVulnerabilities) | **Get** /sources/{source_id}/vuln/{vtype} | Get vulnerabilities for the source by type
[**GetSourceVulnerabilityTypes**](SourcesApi.md#GetSourceVulnerabilityTypes) | **Get** /sources/{source_id}/vuln | Get the available vulnerability types for source
[**ListSources**](SourcesApi.md#ListSources) | **Get** /sources | List the source repository analysis records



## DeleteSource

> SourceManifest DeleteSource(ctx, sourceId).Force(force).Execute()

Delete source record from DB

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
    sourceId := "sourceId_example" // string | UUID of source to delete
    force := true // bool | force delete (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourcesApi.DeleteSource(context.Background(), sourceId).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.DeleteSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSource`: SourceManifest
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.DeleteSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | UUID of source to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete | 

### Return type

[**SourceManifest**](SourceManifest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSource

> SourceManifest GetSource(ctx, sourceId).Execute()

Get a detailed source repository analysis metadata record

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
    sourceId := "sourceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourcesApi.GetSource(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.GetSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSource`: SourceManifest
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.GetSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceManifest**](SourceManifest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceContentByType

> SourceContentPackageResponse GetSourceContentByType(ctx, sourceId, contentType).Execute()

Get the content of an analyzed source repository

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
    sourceId := "sourceId_example" // string | 
    contentType := "contentType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourcesApi.GetSourceContentByType(context.Background(), sourceId, contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.GetSourceContentByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceContentByType`: SourceContentPackageResponse
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.GetSourceContentByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 
**contentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceContentByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SourceContentPackageResponse**](SourceContentPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceContentTypes

> []string GetSourceContentTypes(ctx, sourceId).Execute()

Get a detailed source repository analysis metadata record

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
    sourceId := "sourceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourcesApi.GetSourceContentTypes(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.GetSourceContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceContentTypes`: []string
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.GetSourceContentTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceContentTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourcePolicyCheck

> []PolicyEvaluationResult GetSourcePolicyCheck(ctx, sourceId).PolicyId(policyId).Execute()

Fetch or calculate policy evaluation for a source

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
    sourceId := "sourceId_example" // string | UUID of source to get
    policyId := "policyId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourcesApi.GetSourcePolicyCheck(context.Background(), sourceId).PolicyId(policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.GetSourcePolicyCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourcePolicyCheck`: []PolicyEvaluationResult
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.GetSourcePolicyCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | UUID of source to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourcePolicyCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyId** | **string** |  | 

### Return type

[**[]PolicyEvaluationResult**](PolicyEvaluationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceSbomNative

> *os.File GetSourceSbomNative(ctx, sourceId).Execute()



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
    sourceId := "sourceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourcesApi.GetSourceSbomNative(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.GetSourceSbomNative``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceSbomNative`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.GetSourceSbomNative`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceSbomNativeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceSbomTypes

> []string GetSourceSbomTypes(ctx, sourceId).Execute()

Get a detailed source repository analysis metadata record

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
    sourceId := "sourceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourcesApi.GetSourceSbomTypes(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.GetSourceSbomTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceSbomTypes`: []string
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.GetSourceSbomTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceSbomTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceVulnerabilities

> SourceVulnerabilitiesResponse GetSourceVulnerabilities(ctx, sourceId, vtype).ForceRefresh(forceRefresh).WillNotFix(willNotFix).XAnchoreAccount(xAnchoreAccount).Execute()

Get vulnerabilities for the source by type

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
    sourceId := "sourceId_example" // string | 
    vtype := "vtype_example" // string | 
    forceRefresh := true // bool |  (optional)
    willNotFix := true // bool | Vulnerability data publishers explicitly won't fix some vulnerabilities. This is captured by will_not_fix attribute of each result. If the query parameter is set, results matching it's value will be filtered. Results are not filtered if the query parameter is unset (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourcesApi.GetSourceVulnerabilities(context.Background(), sourceId, vtype).ForceRefresh(forceRefresh).WillNotFix(willNotFix).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.GetSourceVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceVulnerabilities`: SourceVulnerabilitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.GetSourceVulnerabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 
**vtype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceRefresh** | **bool** |  | 
 **willNotFix** | **bool** | Vulnerability data publishers explicitly won&#39;t fix some vulnerabilities. This is captured by will_not_fix attribute of each result. If the query parameter is set, results matching it&#39;s value will be filtered. Results are not filtered if the query parameter is unset | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**SourceVulnerabilitiesResponse**](SourceVulnerabilitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceVulnerabilityTypes

> []string GetSourceVulnerabilityTypes(ctx, sourceId).XAnchoreAccount(xAnchoreAccount).Execute()

Get the available vulnerability types for source

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
    sourceId := "sourceId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourcesApi.GetSourceVulnerabilityTypes(context.Background(), sourceId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.GetSourceVulnerabilityTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceVulnerabilityTypes`: []string
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.GetSourceVulnerabilityTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceVulnerabilityTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSources

> []Source ListSources(ctx).Execute()

List the source repository analysis records

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
    resp, r, err := api_client.SourcesApi.ListSources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.ListSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSources`: []Source
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.ListSources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSourcesRequest struct via the builder pattern


### Return type

[**[]Source**](Source.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

