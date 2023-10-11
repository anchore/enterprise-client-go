# \ImportsAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOperation**](ImportsAPI.md#CreateOperation) | **Post** /imports/images | Begin the import of an image SBOM into the system
[**CreateSourcesOperation**](ImportsAPI.md#CreateSourcesOperation) | **Post** /imports/sources | Begin the import of a source code repository analyzed by Syft into the system
[**FinalizeOperation**](ImportsAPI.md#FinalizeOperation) | **Post** /imports/sources/{operation_id}/finalize | Add source records to catalog db
[**GetImportSourcesSbom**](ImportsAPI.md#GetImportSourcesSbom) | **Get** /imports/sources/{operation_id}/sbom | list the packages of an imported source code repository
[**GetOperation**](ImportsAPI.md#GetOperation) | **Get** /imports/images/{operation_id} | Get detail on a single import
[**GetSourcesOperation**](ImportsAPI.md#GetSourcesOperation) | **Get** /imports/sources/{operation_id} | Get detail on a single import
[**ImportContentSearches**](ImportsAPI.md#ImportContentSearches) | **Post** /imports/images/{operation_id}/content-searches | Import a content search analysis catalog
[**ImportFileContents**](ImportsAPI.md#ImportFileContents) | **Post** /imports/images/{operation_id}/file-contents | Import a file contents analysis catalog
[**ImportImageConfig**](ImportsAPI.md#ImportImageConfig) | **Post** /imports/images/{operation_id}/image-config | Import a docker or OCI image config to associate with the image
[**ImportImageDockerfile**](ImportsAPI.md#ImportImageDockerfile) | **Post** /imports/images/{operation_id}/dockerfile | Begin the import of an image analyzed by Syft into the system
[**ImportImageManifest**](ImportsAPI.md#ImportImageManifest) | **Post** /imports/images/{operation_id}/manifest | Import a docker or OCI distribution manifest to associate with the image
[**ImportImagePackages**](ImportsAPI.md#ImportImagePackages) | **Post** /imports/images/{operation_id}/packages | Begin the import of an image analyzed by Syft into the system
[**ImportImageParentManifest**](ImportsAPI.md#ImportImageParentManifest) | **Post** /imports/images/{operation_id}/parent-manifest | Import a docker or OCI distribution manifest list to associate with the image
[**ImportSecretSearches**](ImportsAPI.md#ImportSecretSearches) | **Post** /imports/images/{operation_id}/secret-searches | Import a secret search analysis catalog
[**InvalidateOperation**](ImportsAPI.md#InvalidateOperation) | **Delete** /imports/images/{operation_id} | Invalidate operation ID so it can be garbage collected
[**InvalidateSourcesOperation**](ImportsAPI.md#InvalidateSourcesOperation) | **Delete** /imports/sources/{operation_id} | Invalidate operation ID so it can be garbage collected
[**ListImportContentSearches**](ImportsAPI.md#ListImportContentSearches) | **Get** /imports/images/{operation_id}/content-searches | List uploaded content search results
[**ListImportDockerfiles**](ImportsAPI.md#ListImportDockerfiles) | **Get** /imports/images/{operation_id}/dockerfile | List uploaded dockerfiles
[**ListImportFileContents**](ImportsAPI.md#ListImportFileContents) | **Get** /imports/images/{operation_id}/file-contents | List uploaded file contents
[**ListImportImageConfigs**](ImportsAPI.md#ListImportImageConfigs) | **Get** /imports/images/{operation_id}/image-config | List uploaded image configs
[**ListImportImageManifests**](ImportsAPI.md#ListImportImageManifests) | **Get** /imports/images/{operation_id}/manifest | List uploaded image manifests
[**ListImportPackages**](ImportsAPI.md#ListImportPackages) | **Get** /imports/images/{operation_id}/packages | List uploaded package manifests
[**ListImportParentManifests**](ImportsAPI.md#ListImportParentManifests) | **Get** /imports/images/{operation_id}/parent-manifest | List uploaded parent manifests (manifest lists for a tag)
[**ListImportSecretSearches**](ImportsAPI.md#ListImportSecretSearches) | **Get** /imports/images/{operation_id}/secret-searches | List uploaded secret search results
[**ListOperations**](ImportsAPI.md#ListOperations) | **Get** /imports/images | Lists in-progress imports
[**ListSourcesOperations**](ImportsAPI.md#ListSourcesOperations) | **Get** /imports/sources | Lists in-progress imports
[**UploadImportSourcesSbom**](ImportsAPI.md#UploadImportSourcesSbom) | **Post** /imports/sources/{operation_id}/sbom | Begin the import of a source code repository analyzed by Syft into the system



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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.CreateOperation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.CreateOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOperation`: ImageImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.CreateOperation`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.CreateSourcesOperation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.CreateSourcesOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSourcesOperation`: SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.CreateSourcesOperation`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 
    metadata := *openapiclient.NewSourceImportMetadata("Host_example", "RepositoryName_example", "Revision_example", *openapiclient.NewSourceImportMetadataContents("Sbom_example")) // SourceImportMetadata | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.FinalizeOperation(context.Background(), operationId).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.FinalizeOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinalizeOperation`: SourceManifest
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.FinalizeOperation`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.GetImportSourcesSbom(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.GetImportSourcesSbom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImportSourcesSbom`: SourceImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.GetImportSourcesSbom`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.GetOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.GetOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperation`: ImageImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.GetOperation`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.GetSourcesOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.GetSourcesOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourcesOperation`: SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.GetSourcesOperation`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 
    contents := []openapiclient.ImageImportContentSearch{*openapiclient.NewImageImportContentSearch(*openapiclient.NewImportPackageLocation("Path_example"), []openapiclient.ImportContentSearchElement{*openapiclient.NewImportContentSearchElement("Classification_example", int32(123), int32(123), int32(123), int32(123))})} // []ImageImportContentSearch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ImportContentSearches(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ImportContentSearches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportContentSearches`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ImportContentSearches`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 
    contents := []openapiclient.ImageImportFileContent{*openapiclient.NewImageImportFileContent(*openapiclient.NewImportPackageLocation("Path_example"), "Contents_example")} // []ImageImportFileContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ImportFileContents(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ImportFileContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportFileContents`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ImportFileContents`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 
    contents := interface{}{ ... } // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ImportImageConfig(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ImportImageConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImageConfig`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ImportImageConfig`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 
    contents := "contents_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ImportImageDockerfile(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ImportImageDockerfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImageDockerfile`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ImportImageDockerfile`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 
    contents := interface{}{ ... } // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ImportImageManifest(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ImportImageManifest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImageManifest`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ImportImageManifest`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 
    sbom := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ImportImagePackages(context.Background(), operationId).Sbom(sbom).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ImportImagePackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImagePackages`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ImportImagePackages`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 
    contents := interface{}{ ... } // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ImportImageParentManifest(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ImportImageParentManifest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportImageParentManifest`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ImportImageParentManifest`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 
    contents := []openapiclient.ImageImportContentSearch{*openapiclient.NewImageImportContentSearch(*openapiclient.NewImportPackageLocation("Path_example"), []openapiclient.ImportContentSearchElement{*openapiclient.NewImportContentSearchElement("Classification_example", int32(123), int32(123), int32(123), int32(123))})} // []ImageImportContentSearch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ImportSecretSearches(context.Background(), operationId).Contents(contents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ImportSecretSearches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSecretSearches`: ImageImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ImportSecretSearches`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.InvalidateOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.InvalidateOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvalidateOperation`: ImageImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.InvalidateOperation`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.InvalidateSourcesOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.InvalidateSourcesOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvalidateSourcesOperation`: SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.InvalidateSourcesOperation`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ListImportContentSearches(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ListImportContentSearches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportContentSearches`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ListImportContentSearches`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ListImportDockerfiles(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ListImportDockerfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportDockerfiles`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ListImportDockerfiles`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ListImportFileContents(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ListImportFileContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportFileContents`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ListImportFileContents`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ListImportImageConfigs(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ListImportImageConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportImageConfigs`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ListImportImageConfigs`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ListImportImageManifests(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ListImportImageManifests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportImageManifests`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ListImportImageManifests`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ListImportPackages(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ListImportPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportPackages`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ListImportPackages`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ListImportParentManifests(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ListImportParentManifests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportParentManifests`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ListImportParentManifests`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ListImportSecretSearches(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ListImportSecretSearches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImportSecretSearches`: []string
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ListImportSecretSearches`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ListOperations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ListOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOperations`: []ImageImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ListOperations`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.ListSourcesOperations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.ListSourcesOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourcesOperations`: []SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.ListSourcesOperations`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    operationId := "operationId_example" // string | 
    sbom := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsAPI.UploadImportSourcesSbom(context.Background(), operationId).Sbom(sbom).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsAPI.UploadImportSourcesSbom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadImportSourcesSbom`: SourceImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportsAPI.UploadImportSourcesSbom`: %v\n", resp)
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

