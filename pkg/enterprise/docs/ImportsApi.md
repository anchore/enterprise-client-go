# \ImportsApi

All URIs are relative to *http://localhost/enterprise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOperation**](ImportsApi.md#CreateOperation) | **Post** /imports/sources | Begin the import of a source code repository analyzed by Syft into the system
[**FinalizeOperation**](ImportsApi.md#FinalizeOperation) | **Post** /imports/sources/{operation_id}/finalize | Add source records to catalog db
[**GetImportSourcesSbom**](ImportsApi.md#GetImportSourcesSbom) | **Get** /imports/sources/{operation_id}/sbom | list the packages of an imported source code repository
[**GetOperation**](ImportsApi.md#GetOperation) | **Get** /imports/sources/{operation_id} | Get detail on a single import
[**InvalidateOperation**](ImportsApi.md#InvalidateOperation) | **Delete** /imports/sources/{operation_id} | Invalidate operation ID so it can be garbage collected
[**ListOperations**](ImportsApi.md#ListOperations) | **Get** /imports/sources | Lists in-progress imports
[**UploadImportSourcesSbom**](ImportsApi.md#UploadImportSourcesSbom) | **Post** /imports/sources/{operation_id}/sbom | Begin the import of a source code repository analyzed by Syft into the system



## CreateOperation

> SourceImportOperation CreateOperation(ctx).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportsApi.CreateOperation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.CreateOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOperation`: SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.CreateOperation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationRequest struct via the builder pattern


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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportsApi.FinalizeOperation(context.Background(), operationId).Metadata(metadata).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportsApi.GetImportSourcesSbom(context.Background(), operationId).Execute()
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

> SourceImportOperation GetOperation(ctx, operationId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportsApi.GetOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.GetOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperation`: SourceImportOperation
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

[**SourceImportOperation**](SourceImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvalidateOperation

> SourceImportOperation InvalidateOperation(ctx, operationId).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportsApi.InvalidateOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.InvalidateOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvalidateOperation`: SourceImportOperation
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

[**SourceImportOperation**](SourceImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperations

> []SourceImportOperation ListOperations(ctx).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportsApi.ListOperations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.ListOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOperations`: []SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.ListOperations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsRequest struct via the builder pattern


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
    sbom := *openapiclient.NewNativeSBOM([]openapiclient.NativeSBOMPackage{*openapiclient.NewNativeSBOMPackage("Name_example", "Version_example", "Type_example", []openapiclient.NativeSBOMPackageLocation{*openapiclient.NewNativeSBOMPackageLocation("Path_example")}, []string{"Licenses_example"}, "Language_example", []string{"Cpes_example"})}, *openapiclient.NewNativeSBOMSource("Type_example", interface{}(123)), *openapiclient.NewNativeSBOMDistribution()) // NativeSBOM | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportsApi.UploadImportSourcesSbom(context.Background(), operationId).Sbom(sbom).Execute()
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

 **sbom** | [**NativeSBOM**](NativeSBOM.md) |  | 

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

