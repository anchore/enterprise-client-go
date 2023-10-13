# \ImportsApi

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOperation**](ImportsApi.md#CreateOperation) | **Post** /imports/images | Begin the import of an image SBOM into the system
[**CreateSourcesOperation**](ImportsApi.md#CreateSourcesOperation) | **Post** /imports/sources | Begin the import of a source code repository analyzed by Syft into the system
[**FinalizeOperation**](ImportsApi.md#FinalizeOperation) | **Post** /imports/sources/{operation_id}/finalize | Add source records to catalog db
[**GetImportSourcesSbom**](ImportsApi.md#GetImportSourcesSbom) | **Get** /imports/sources/{operation_id}/sbom | list the packages of an imported source code repository
[**GetOperation**](ImportsApi.md#GetOperation) | **Get** /imports/images/{operation_id} | Get detail on a single import
[**GetSourcesOperation**](ImportsApi.md#GetSourcesOperation) | **Get** /imports/sources/{operation_id} | Get detail on a single import
[**ImportContentSearches**](ImportsApi.md#ImportContentSearches) | **Post** /imports/images/{operation_id}/content-searches | Import a content search analysis catalog
[**ImportFileContents**](ImportsApi.md#ImportFileContents) | **Post** /imports/images/{operation_id}/file-contents | Import a file contents analysis catalog
[**ImportImageConfig**](ImportsApi.md#ImportImageConfig) | **Post** /imports/images/{operation_id}/image-config | Import a docker or OCI image config to associate with the image
[**ImportImageDockerfile**](ImportsApi.md#ImportImageDockerfile) | **Post** /imports/images/{operation_id}/dockerfile | Begin the import of an image analyzed by Syft into the system
[**ImportImageManifest**](ImportsApi.md#ImportImageManifest) | **Post** /imports/images/{operation_id}/manifest | Import a docker or OCI distribution manifest to associate with the image
[**ImportImagePackages**](ImportsApi.md#ImportImagePackages) | **Post** /imports/images/{operation_id}/packages | Begin the import of an image analyzed by Syft into the system
[**ImportImageParentManifest**](ImportsApi.md#ImportImageParentManifest) | **Post** /imports/images/{operation_id}/parent-manifest | Import a docker or OCI distribution manifest list to associate with the image
[**ImportSecretSearches**](ImportsApi.md#ImportSecretSearches) | **Post** /imports/images/{operation_id}/secret-searches | Import a secret search analysis catalog
[**InvalidateOperation**](ImportsApi.md#InvalidateOperation) | **Delete** /imports/images/{operation_id} | Invalidate operation ID so it can be garbage collected
[**InvalidateSourcesOperation**](ImportsApi.md#InvalidateSourcesOperation) | **Delete** /imports/sources/{operation_id} | Invalidate operation ID so it can be garbage collected
[**ListImportContentSearches**](ImportsApi.md#ListImportContentSearches) | **Get** /imports/images/{operation_id}/content-searches | List uploaded content search results
[**ListImportDockerfiles**](ImportsApi.md#ListImportDockerfiles) | **Get** /imports/images/{operation_id}/dockerfile | List uploaded dockerfiles
[**ListImportFileContents**](ImportsApi.md#ListImportFileContents) | **Get** /imports/images/{operation_id}/file-contents | List uploaded file contents
[**ListImportImageConfigs**](ImportsApi.md#ListImportImageConfigs) | **Get** /imports/images/{operation_id}/image-config | List uploaded image configs
[**ListImportImageManifests**](ImportsApi.md#ListImportImageManifests) | **Get** /imports/images/{operation_id}/manifest | List uploaded image manifests
[**ListImportPackages**](ImportsApi.md#ListImportPackages) | **Get** /imports/images/{operation_id}/packages | List uploaded package manifests
[**ListImportParentManifests**](ImportsApi.md#ListImportParentManifests) | **Get** /imports/images/{operation_id}/parent-manifest | List uploaded parent manifests (manifest lists for a tag)
[**ListImportSecretSearches**](ImportsApi.md#ListImportSecretSearches) | **Get** /imports/images/{operation_id}/secret-searches | List uploaded secret search results
[**ListOperations**](ImportsApi.md#ListOperations) | **Get** /imports/images | Lists in-progress imports
[**ListSourcesOperations**](ImportsApi.md#ListSourcesOperations) | **Get** /imports/sources | Lists in-progress imports
[**UploadImportSourcesSbom**](ImportsApi.md#UploadImportSourcesSbom) | **Post** /imports/sources/{operation_id}/sbom | Begin the import of a source code repository analyzed by Syft into the system



## CreateOperation

> ImageImportOperation CreateOperation(ctx).Execute()

Begin the import of an image SBOM into the system

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
    resp, r, err := apiClient.ImportsApi.CreateOperation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.CreateOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOperation`: ImageImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.CreateOperation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationRequest struct via the builder pattern


### Return type

[**ImageImportOperation**](ImageImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSourcesOperation

> SourceImportOperation CreateSourcesOperation(ctx).Execute()

Begin the import of a source code repository analyzed by Syft into the system

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
    resp, r, err := apiClient.ImportsApi.CreateSourcesOperation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.CreateSourcesOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSourcesOperation`: SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.CreateSourcesOperation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourcesOperationRequest struct via the builder pattern


### Return type

[**SourceImportOperation**](SourceImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinalizeOperation

> SourceManifest FinalizeOperation(ctx, operationId).Metadata(metadata).Execute()

Add source records to catalog db

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
    operationId := "operationId_example" // string | 
    metadata := *openapiclient.NewSourceImportMetadata("Host_example", "RepositoryName_example", "Revision_example", *openapiclient.NewSourceImportMetadataContents("Sbom_example")) // SourceImportMetadata | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.FinalizeOperation(context.Background(), operationId).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.FinalizeOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinalizeOperation`: SourceManifest
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.FinalizeOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinalizeOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadata** | [**SourceImportMetadata**](SourceImportMetadata.md) |  | 

### Return type

[**SourceManifest**](SourceManifest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportSourcesSbom

> SourceImportContentResponse GetImportSourcesSbom(ctx, operationId).Execute()

list the packages of an imported source code repository

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.GetImportSourcesSbom(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.GetImportSourcesSbom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImportSourcesSbom`: SourceImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.GetImportSourcesSbom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportSourcesSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceImportContentResponse**](SourceImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperation

> ImageImportOperation GetOperation(ctx, operationId).Execute()

Get detail on a single import

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.GetOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.GetOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperation`: ImageImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.GetOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImageImportOperation**](ImageImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourcesOperation

> SourceImportOperation GetSourcesOperation(ctx, operationId).Execute()

Get detail on a single import

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.GetSourcesOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.GetSourcesOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourcesOperation`: SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.GetSourcesOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourcesOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceImportOperation**](SourceImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportContentSearches

> ImageImportContentResponse ImportContentSearches(ctx, operationId).Contents(contents).Execute()

Import a content search analysis catalog

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
    operationId := "operationId_example" // string | 
    contents := []openapiclient.ImageImportContentSearch{*openapiclient.NewImageImportContentSearch(*openapiclient.NewImportPackageLocation("Path_example"), []openapiclient.ImportContentSearchElement{*openapiclient.NewImportContentSearchElement("Classification_example", int32(123), int32(123), int32(123), int32(123))})} // []ImageImportContentSearch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ImportContentSearches(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ImportContentSearches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportContentSearches`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ImportContentSearches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportContentSearchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contents** | [**[]ImageImportContentSearch**](ImageImportContentSearch.md) |  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportFileContents

> ImageImportContentResponse ImportFileContents(ctx, operationId).Contents(contents).Execute()

Import a file contents analysis catalog

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
    operationId := "operationId_example" // string | 
    contents := []openapiclient.ImageImportFileContent{*openapiclient.NewImageImportFileContent(*openapiclient.NewImportPackageLocation("Path_example"), "Contents_example")} // []ImageImportFileContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ImportFileContents(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ImportFileContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportFileContents`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ImportFileContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportFileContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contents** | [**[]ImageImportFileContent**](ImageImportFileContent.md) |  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportImageConfig

> ImageImportContentResponse ImportImageConfig(ctx, operationId).Contents(contents).Execute()

Import a docker or OCI image config to associate with the image

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
    operationId := "operationId_example" // string | 
    contents := interface{}{ ... } // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ImportImageConfig(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ImportImageConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImageConfig`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ImportImageConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportImageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contents** | **interface{}** |  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportImageDockerfile

> ImageImportContentResponse ImportImageDockerfile(ctx, operationId).Contents(contents).Execute()

Begin the import of an image analyzed by Syft into the system

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
    operationId := "operationId_example" // string | 
    contents := "contents_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ImportImageDockerfile(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ImportImageDockerfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImageDockerfile`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ImportImageDockerfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportImageDockerfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contents** | **string** |  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain; utf-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportImageManifest

> ImageImportContentResponse ImportImageManifest(ctx, operationId).Contents(contents).Execute()

Import a docker or OCI distribution manifest to associate with the image

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
    operationId := "operationId_example" // string | 
    contents := interface{}{ ... } // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ImportImageManifest(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ImportImageManifest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImageManifest`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ImportImageManifest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportImageManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contents** | **interface{}** |  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.oci.image.manifest.v1+json, application/vnd.docker.distribution.manifest.v2+json, application/vnd.docker.distribution.manifest.v1+json, application/vnd.docker.distribution.manifest.v1+prettyjws
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportImagePackages

> ImageImportContentResponse ImportImagePackages(ctx, operationId).Sbom(sbom).Execute()

Begin the import of an image analyzed by Syft into the system

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
    operationId := "operationId_example" // string | 
    sbom := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ImportImagePackages(context.Background(), operationId).Sbom(sbom).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ImportImagePackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImagePackages`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ImportImagePackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportImagePackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sbom** | **map[string]interface{}** |  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportImageParentManifest

> ImageImportContentResponse ImportImageParentManifest(ctx, operationId).Contents(contents).Execute()

Import a docker or OCI distribution manifest list to associate with the image

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
    operationId := "operationId_example" // string | 
    contents := interface{}{ ... } // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ImportImageParentManifest(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ImportImageParentManifest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImageParentManifest`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ImportImageParentManifest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportImageParentManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contents** | **interface{}** |  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.docker.distribution.manifest.list.v2+json, application/vnd.oci.image.index.v1+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSecretSearches

> ImageImportContentResponse ImportSecretSearches(ctx, operationId).Contents(contents).Execute()

Import a secret search analysis catalog

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
    operationId := "operationId_example" // string | 
    contents := []openapiclient.ImageImportContentSearch{*openapiclient.NewImageImportContentSearch(*openapiclient.NewImportPackageLocation("Path_example"), []openapiclient.ImportContentSearchElement{*openapiclient.NewImportContentSearchElement("Classification_example", int32(123), int32(123), int32(123), int32(123))})} // []ImageImportContentSearch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ImportSecretSearches(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ImportSecretSearches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSecretSearches`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ImportSecretSearches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSecretSearchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contents** | [**[]ImageImportContentSearch**](ImageImportContentSearch.md) |  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvalidateOperation

> ImageImportOperation InvalidateOperation(ctx, operationId).Execute()

Invalidate operation ID so it can be garbage collected

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.InvalidateOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.InvalidateOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvalidateOperation`: ImageImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.InvalidateOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImageImportOperation**](ImageImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvalidateSourcesOperation

> SourceImportOperation InvalidateSourcesOperation(ctx, operationId).Execute()

Invalidate operation ID so it can be garbage collected

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.InvalidateSourcesOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.InvalidateSourcesOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvalidateSourcesOperation`: SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.InvalidateSourcesOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateSourcesOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceImportOperation**](SourceImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImportContentSearches

> []string ListImportContentSearches(ctx, operationId).Execute()

List uploaded content search results

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ListImportContentSearches(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ListImportContentSearches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportContentSearches`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ListImportContentSearches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImportContentSearchesRequest struct via the builder pattern


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


## ListImportDockerfiles

> []string ListImportDockerfiles(ctx, operationId).Execute()

List uploaded dockerfiles

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ListImportDockerfiles(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ListImportDockerfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportDockerfiles`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ListImportDockerfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImportDockerfilesRequest struct via the builder pattern


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


## ListImportFileContents

> []string ListImportFileContents(ctx, operationId).Execute()

List uploaded file contents

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ListImportFileContents(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ListImportFileContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportFileContents`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ListImportFileContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImportFileContentsRequest struct via the builder pattern


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


## ListImportImageConfigs

> []string ListImportImageConfigs(ctx, operationId).Execute()

List uploaded image configs

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ListImportImageConfigs(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ListImportImageConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportImageConfigs`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ListImportImageConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImportImageConfigsRequest struct via the builder pattern


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


## ListImportImageManifests

> []string ListImportImageManifests(ctx, operationId).Execute()

List uploaded image manifests

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ListImportImageManifests(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ListImportImageManifests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportImageManifests`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ListImportImageManifests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImportImageManifestsRequest struct via the builder pattern


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


## ListImportPackages

> []string ListImportPackages(ctx, operationId).Execute()

List uploaded package manifests

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ListImportPackages(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ListImportPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportPackages`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ListImportPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImportPackagesRequest struct via the builder pattern


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


## ListImportParentManifests

> []string ListImportParentManifests(ctx, operationId).Execute()

List uploaded parent manifests (manifest lists for a tag)

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ListImportParentManifests(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ListImportParentManifests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportParentManifests`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ListImportParentManifests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImportParentManifestsRequest struct via the builder pattern


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


## ListImportSecretSearches

> []string ListImportSecretSearches(ctx, operationId).Execute()

List uploaded secret search results

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
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.ListImportSecretSearches(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ListImportSecretSearches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportSecretSearches`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ListImportSecretSearches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImportSecretSearchesRequest struct via the builder pattern


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


## ListOperations

> []ImageImportOperation ListOperations(ctx).Execute()

Lists in-progress imports

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
    resp, r, err := apiClient.ImportsApi.ListOperations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ListOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOperations`: []ImageImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ListOperations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsRequest struct via the builder pattern


### Return type

[**[]ImageImportOperation**](ImageImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourcesOperations

> []SourceImportOperation ListSourcesOperations(ctx).Execute()

Lists in-progress imports

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
    resp, r, err := apiClient.ImportsApi.ListSourcesOperations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ListSourcesOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourcesOperations`: []SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ListSourcesOperations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSourcesOperationsRequest struct via the builder pattern


### Return type

[**[]SourceImportOperation**](SourceImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadImportSourcesSbom

> SourceImportContentResponse UploadImportSourcesSbom(ctx, operationId).Sbom(sbom).Execute()

Begin the import of a source code repository analyzed by Syft into the system

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
    operationId := "operationId_example" // string | 
    sbom := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.UploadImportSourcesSbom(context.Background(), operationId).Sbom(sbom).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.UploadImportSourcesSbom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadImportSourcesSbom`: SourceImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.UploadImportSourcesSbom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadImportSourcesSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sbom** | **map[string]interface{}** |  | 

### Return type

[**SourceImportContentResponse**](SourceImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

