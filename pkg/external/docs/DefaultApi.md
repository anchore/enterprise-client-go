# \DefaultApi

All URIs are relative to */enterprise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddActionPlan**](DefaultApi.md#AddActionPlan) | **Post** /actions | Submits an Action Plan
[**AddApplication**](DefaultApi.md#AddApplication) | **Post** /applications | Create an application
[**AddApplicationVersion**](DefaultApi.md#AddApplicationVersion) | **Post** /applications/{application_id}/versions | Create an application version
[**AddArtifactToApplicationVersion**](DefaultApi.md#AddArtifactToApplicationVersion) | **Post** /applications/{application_id}/versions/{application_version_id}/artifacts | Add an artifact to an application version
[**AddCorrection**](DefaultApi.md#AddCorrection) | **Post** /corrections | Create a correction record
[**AddInventoryCluster**](DefaultApi.md#AddInventoryCluster) | **Post** /inventories/clusters | Create a cluster inventory
[**AddRuntimeComplianceCheck**](DefaultApi.md#AddRuntimeComplianceCheck) | **Post** /runtime_compliance | Post a runtime compliance check
[**CreateOperation**](DefaultApi.md#CreateOperation) | **Post** /imports/sources | Begin the import of a source code repository analyzed by Syft into the system
[**DelInventoryClusterByName**](DefaultApi.md#DelInventoryClusterByName) | **Delete** /inventories/clusters/{cluster_name} | Delete a configured inventory clusters by cluster_name
[**DeleteApplication**](DefaultApi.md#DeleteApplication) | **Delete** /applications/{application_id} | Delete an application by application_id
[**DeleteApplicationVersion**](DefaultApi.md#DeleteApplicationVersion) | **Delete** /applications/{application_id}/versions/{application_version_id} | Delete an application version by application_id and application_version_id
[**DeleteCorrectionByUuid**](DefaultApi.md#DeleteCorrectionByUuid) | **Delete** /corrections/{uuid} | Delete a correction by UUID
[**DeleteSource**](DefaultApi.md#DeleteSource) | **Delete** /sources/{source_id} | Delete source record from DB
[**FinalizeOperation**](DefaultApi.md#FinalizeOperation) | **Post** /imports/sources/{operation_id}/finalize | Add source records to catalog db
[**GetActionPlans**](DefaultApi.md#GetActionPlans) | **Get** /actions | Gets a list of submitted action (remediation) plans
[**GetAlertSummaries**](DefaultApi.md#GetAlertSummaries) | **Get** /alerts/summaries | List all alert summaries scoped to the account
[**GetApplication**](DefaultApi.md#GetApplication) | **Get** /applications/{application_id} | Get an application by application_id
[**GetApplicationVersion**](DefaultApi.md#GetApplicationVersion) | **Get** /applications/{application_id}/versions/{application_version_id} | Get an application version
[**GetApplicationVersionSbom**](DefaultApi.md#GetApplicationVersionSbom) | **Get** /applications/{application_id}/versions/{application_version_id}/sboms/json | Get the combined sbom for the given application version, optionally filtered by artifact type
[**GetApplicationVersions**](DefaultApi.md#GetApplicationVersions) | **Get** /applications/{application_id}/versions | List all application verions
[**GetApplications**](DefaultApi.md#GetApplications) | **Get** /applications | List all applications
[**GetComplianceViolationAlert**](DefaultApi.md#GetComplianceViolationAlert) | **Get** /alerts/compliance_violations/{uuid} | Get compliance violation alert by id
[**GetComplianceViolationAlerts**](DefaultApi.md#GetComplianceViolationAlerts) | **Get** /alerts/compliance_violations | List all compliance violation alerts scoped to the account
[**GetCorrectionByUuid**](DefaultApi.md#GetCorrectionByUuid) | **Get** /corrections/{uuid} | Retrieve a correction by UUID
[**GetCorrections**](DefaultApi.md#GetCorrections) | **Get** /corrections | Retrieve a list of corrections
[**GetImageAncestors**](DefaultApi.md#GetImageAncestors) | **Get** /images/{image_digest}/ancestors | Return the list of ancestor images for the given image
[**GetImageInventory**](DefaultApi.md#GetImageInventory) | **Get** /inventories | Return a list of the images in inventories for this account
[**GetImagePolicyCheckByDigest**](DefaultApi.md#GetImagePolicyCheckByDigest) | **Get** /images/{imageDigest}/check | Check policy evaluation status for image
[**GetImageVulnerabilitiesByDigest**](DefaultApi.md#GetImageVulnerabilitiesByDigest) | **Get** /images/{imageDigest}/vuln/{vtype} | Get vulnerabilities by type
[**GetImportSourcesSbom**](DefaultApi.md#GetImportSourcesSbom) | **Get** /imports/sources/{operation_id}/sbom | list the packages of an imported source code repository
[**GetInventoryClusterByName**](DefaultApi.md#GetInventoryClusterByName) | **Get** /inventories/clusters/{cluster_name} | Return a configured inventory cluster
[**GetOperation**](DefaultApi.md#GetOperation) | **Get** /imports/sources/{operation_id} | Get detail on a single import
[**GetRuntimeComplianceChecks**](DefaultApi.md#GetRuntimeComplianceChecks) | **Get** /runtime_compliance | Get all runtime compliance checks or just those for a given image digest
[**GetRuntimeComplianceResult**](DefaultApi.md#GetRuntimeComplianceResult) | **Get** /runtime_compliance/result/{compliance_file_id} | Check the results of a a specific runtime compliance check
[**GetSource**](DefaultApi.md#GetSource) | **Get** /sources/{source_id} | Get a detailed source repository analysis metadata record
[**GetSourceContentByType**](DefaultApi.md#GetSourceContentByType) | **Get** /sources/{source_id}/content/{content_type} | Get the content of an analyzed source repository
[**GetSourceContentTypes**](DefaultApi.md#GetSourceContentTypes) | **Get** /sources/{source_id}/content | Get a detailed source repository analysis metadata record
[**GetSourceSbomNative**](DefaultApi.md#GetSourceSbomNative) | **Get** /sources/{source_id}/sbom/native | 
[**GetSourceSbomTypes**](DefaultApi.md#GetSourceSbomTypes) | **Get** /sources/{source_id}/sbom | Get a detailed source repository analysis metadata record
[**GetSourceVulnerabilities**](DefaultApi.md#GetSourceVulnerabilities) | **Get** /sources/{source_id}/vuln/{vtype} | Get vulnerabilities for the source by type
[**GetSourceVulnerabilityTypes**](DefaultApi.md#GetSourceVulnerabilityTypes) | **Get** /sources/{source_id}/vuln | Get the available vulnerability types for source
[**InvalidateOperation**](DefaultApi.md#InvalidateOperation) | **Delete** /imports/sources/{operation_id} | Invalidate operation ID so it can be garbage collected
[**ListArtifacts**](DefaultApi.md#ListArtifacts) | **Get** /applications/{application_id}/versions/{application_version_id}/artifacts | List artifacts present on a given application version
[**ListInventoryClusters**](DefaultApi.md#ListInventoryClusters) | **Get** /inventories/clusters | Return a list of the configured inventory clusters
[**ListOperations**](DefaultApi.md#ListOperations) | **Get** /imports/sources | Lists in-progress imports
[**ListSources**](DefaultApi.md#ListSources) | **Get** /sources | List the source repository analysis records
[**RemoveArtifactFromApplicationVersion**](DefaultApi.md#RemoveArtifactFromApplicationVersion) | **Delete** /applications/{application_id}/versions/{application_version_id}/artifacts/{association_id} | Delete an artifact from specified application version
[**SyncImageInventory**](DefaultApi.md#SyncImageInventory) | **Post** /inventories | synchronizes the list of the images in a given cluster for the inventory
[**UpdateApplication**](DefaultApi.md#UpdateApplication) | **Put** /applications/{application_id} | Update application details
[**UpdateApplicationVersion**](DefaultApi.md#UpdateApplicationVersion) | **Put** /applications/{application_id}/versions/{application_version_id} | Update application version details
[**UpdateComplianceViolationAlertState**](DefaultApi.md#UpdateComplianceViolationAlertState) | **Put** /alerts/compliance_violations/{uuid}/{state} | Open or close a compliance violation alert
[**UpdateCorrectionByUuid**](DefaultApi.md#UpdateCorrectionByUuid) | **Put** /corrections/{uuid} | Update a correction by UUID
[**UploadImportSourcesSbom**](DefaultApi.md#UploadImportSourcesSbom) | **Post** /imports/sources/{operation_id}/sbom | Begin the import of a source code repository analyzed by Syft into the system



## AddActionPlan

> ActionPlan AddActionPlan(ctx).ActionPlan(actionPlan).Execute()

Submits an Action Plan



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
    actionPlan := *openapiclient.NewActionPlan() // ActionPlan | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddActionPlan(context.Background()).ActionPlan(actionPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddActionPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddActionPlan`: ActionPlan
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddActionPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddActionPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionPlan** | [**ActionPlan**](ActionPlan.md) |  | 

### Return type

[**ActionPlan**](ActionPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddApplication

> Application AddApplication(ctx).Application(application).XAnchoreAccount(xAnchoreAccount).Execute()

Create an application



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
    application := *openapiclient.NewApplication() // Application | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddApplication(context.Background()).Application(application).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | [**Application**](Application.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddApplicationVersion

> ApplicationVersion AddApplicationVersion(ctx, applicationId).ApplicationVersion(applicationVersion).XAnchoreAccount(xAnchoreAccount).Execute()

Create an application version



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
    applicationId := "applicationId_example" // string | 
    applicationVersion := *openapiclient.NewApplicationVersion("VersionName_example") // ApplicationVersion | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddApplicationVersion(context.Background(), applicationId).ApplicationVersion(applicationVersion).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApplicationVersion`: ApplicationVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddApplicationVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationVersion** | [**ApplicationVersion**](ApplicationVersion.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ApplicationVersion**](ApplicationVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddArtifactToApplicationVersion

> ArtifactAssociationResponse AddArtifactToApplicationVersion(ctx, applicationId, applicationVersionId).ArtifactRequest(artifactRequest).XAnchoreAccount(xAnchoreAccount).Execute()

Add an artifact to an application version



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
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    artifactRequest := *openapiclient.NewArtifactAssociationRequest("ArtifactType_example", interface{}(123)) // ArtifactAssociationRequest | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddArtifactToApplicationVersion(context.Background(), applicationId, applicationVersionId).ArtifactRequest(artifactRequest).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddArtifactToApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddArtifactToApplicationVersion`: ArtifactAssociationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddArtifactToApplicationVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**applicationVersionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddArtifactToApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **artifactRequest** | [**ArtifactAssociationRequest**](ArtifactAssociationRequest.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ArtifactAssociationResponse**](ArtifactAssociationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCorrection

> []Correction AddCorrection(ctx).Correction(correction).XAnchoreAccount(xAnchoreAccount).Execute()

Create a correction record



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
    correction := *openapiclient.NewCorrection("Type_example", *openapiclient.NewCorrectionMatch("npm"), []openapiclient.CorrectionFieldMatch{*openapiclient.NewCorrectionFieldMatch("name", "FieldValue_example")}) // Correction | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddCorrection(context.Background()).Correction(correction).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddCorrection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCorrection`: []Correction
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddCorrection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCorrectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **correction** | [**Correction**](Correction.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]Correction**](Correction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddInventoryCluster

> InventoryCluster AddInventoryCluster(ctx).Cluster(cluster).XAnchoreAccount(xAnchoreAccount).Execute()

Create a cluster inventory



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
    cluster := *openapiclient.NewInventoryCluster() // InventoryCluster | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddInventoryCluster(context.Background()).Cluster(cluster).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddInventoryCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInventoryCluster`: InventoryCluster
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddInventoryCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInventoryClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | [**InventoryCluster**](InventoryCluster.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**InventoryCluster**](InventoryCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRuntimeComplianceCheck

> RuntimeComplianceCheck AddRuntimeComplianceCheck(ctx).CheckType(checkType).ImageDigest(imageDigest).XAnchoreAccount(xAnchoreAccount).Result(result).Pod(pod).Namespace(namespace).ImageTag(imageTag).StartTime(startTime).EndTime(endTime).ResultFile(resultFile).ReportFile(reportFile).Execute()

Post a runtime compliance check



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
    checkType := "checkType_example" // string | The type of runtime compliance check
    imageDigest := "imageDigest_example" // string | The digest of the pod the check was run against
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)
    result := "result_example" // string | The result of the runtime compliance check (optional)
    pod := "pod_example" // string | The pod the check was run against (optional)
    namespace := "namespace_example" // string | The namespace of the pod the check was run against (optional)
    imageTag := "imageTag_example" // string | The tag of the image in the pod the check was run against (optional)
    startTime := time.Now() // time.Time | The type of runtime compliance check (optional)
    endTime := time.Now() // time.Time | The type of runtime compliance check (optional)
    resultFile := os.NewFile(1234, "some_file") // *os.File | The file with the check results (optional)
    reportFile := os.NewFile(1234, "some_file") // *os.File | The file with the check port (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddRuntimeComplianceCheck(context.Background()).CheckType(checkType).ImageDigest(imageDigest).XAnchoreAccount(xAnchoreAccount).Result(result).Pod(pod).Namespace(namespace).ImageTag(imageTag).StartTime(startTime).EndTime(endTime).ResultFile(resultFile).ReportFile(reportFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddRuntimeComplianceCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRuntimeComplianceCheck`: RuntimeComplianceCheck
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddRuntimeComplianceCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRuntimeComplianceCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkType** | **string** | The type of runtime compliance check | 
 **imageDigest** | **string** | The digest of the pod the check was run against | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 
 **result** | **string** | The result of the runtime compliance check | 
 **pod** | **string** | The pod the check was run against | 
 **namespace** | **string** | The namespace of the pod the check was run against | 
 **imageTag** | **string** | The tag of the image in the pod the check was run against | 
 **startTime** | **time.Time** | The type of runtime compliance check | 
 **endTime** | **time.Time** | The type of runtime compliance check | 
 **resultFile** | ***os.File** | The file with the check results | 
 **reportFile** | ***os.File** | The file with the check port | 

### Return type

[**RuntimeComplianceCheck**](RuntimeComplianceCheck.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateOperation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOperation`: SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOperation`: %v\n", resp)
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


## DelInventoryClusterByName

> DelInventoryClusterByName(ctx, clusterName).XAnchoreAccount(xAnchoreAccount).Execute()

Delete a configured inventory clusters by cluster_name



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
    clusterName := "clusterName_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DelInventoryClusterByName(context.Background(), clusterName).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DelInventoryClusterByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelInventoryClusterByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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


## DeleteApplication

> DeleteApplication(ctx, applicationId).XAnchoreAccount(xAnchoreAccount).Execute()

Delete an application by application_id



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
    applicationId := "applicationId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteApplication(context.Background(), applicationId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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


## DeleteApplicationVersion

> DeleteApplicationVersion(ctx, applicationId, applicationVersionId).XAnchoreAccount(xAnchoreAccount).Execute()

Delete an application version by application_id and application_version_id



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
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteApplicationVersion(context.Background(), applicationId, applicationVersionId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**applicationVersionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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


## DeleteCorrectionByUuid

> DeleteCorrectionByUuid(ctx, uuid).XAnchoreAccount(xAnchoreAccount).Execute()

Delete a correction by UUID



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
    uuid := "uuid_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteCorrectionByUuid(context.Background(), uuid).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCorrectionByUuid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCorrectionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSource(context.Background(), sourceId).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSource`: SourceManifest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteSource`: %v\n", resp)
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
    metadata := *openapiclient.NewSourceImportMetadata() // SourceImportMetadata | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FinalizeOperation(context.Background(), operationId).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FinalizeOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinalizeOperation`: SourceManifest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FinalizeOperation`: %v\n", resp)
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


## GetActionPlans

> []ActionPlan GetActionPlans(ctx).ImageTag(imageTag).ImageDigest(imageDigest).CreatedAfter(createdAfter).XAnchoreAccount(xAnchoreAccount).Execute()

Gets a list of submitted action (remediation) plans



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
    imageTag := "imageTag_example" // string |  (optional)
    imageDigest := "imageDigest_example" // string |  (optional)
    createdAfter := time.Now() // time.Time | RFC 3339 formatted UTC timestamp to filter out action plans that were only created after this date (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetActionPlans(context.Background()).ImageTag(imageTag).ImageDigest(imageDigest).CreatedAfter(createdAfter).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetActionPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActionPlans`: []ActionPlan
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetActionPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActionPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageTag** | **string** |  | 
 **imageDigest** | **string** |  | 
 **createdAfter** | **time.Time** | RFC 3339 formatted UTC timestamp to filter out action plans that were only created after this date | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]ActionPlan**](ActionPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertSummaries

> []AlertSummary GetAlertSummaries(ctx).Page(page).Limit(limit).Type_(type_).State(state).CreatedAfter(createdAfter).CreatedBefore(createdBefore).ResourceLabel(resourceLabel).XAnchoreAccount(xAnchoreAccount).Execute()

List all alert summaries scoped to the account



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
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 100)
    type_ := "type__example" // string | Filter for alerts based on the type such as compliance violation (optional) (default to "all")
    state := "state_example" // string | Filter for alerts by current state, defaults to open alerts unless specified (optional) (default to "open")
    createdAfter := time.Now() // time.Time | Filter for alerts generated after the timestamp (optional)
    createdBefore := time.Now() // time.Time | Filter for alerts generated before the timestamp (optional)
    resourceLabel := []string{"Inner_example"} // []string | Filter for alerts associated with a resource where the label in key=value format such as tag=docker.io/library/alpine:latest or repository=library/alpine (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAlertSummaries(context.Background()).Page(page).Limit(limit).Type_(type_).State(state).CreatedAfter(createdAfter).CreatedBefore(createdBefore).ResourceLabel(resourceLabel).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAlertSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertSummaries`: []AlertSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAlertSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **type_** | **string** | Filter for alerts based on the type such as compliance violation | [default to &quot;all&quot;]
 **state** | **string** | Filter for alerts by current state, defaults to open alerts unless specified | [default to &quot;open&quot;]
 **createdAfter** | **time.Time** | Filter for alerts generated after the timestamp | 
 **createdBefore** | **time.Time** | Filter for alerts generated before the timestamp | 
 **resourceLabel** | **[]string** | Filter for alerts associated with a resource where the label in key&#x3D;value format such as tag&#x3D;docker.io/library/alpine:latest or repository&#x3D;library/alpine | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]AlertSummary**](AlertSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> Application GetApplication(ctx, applicationId).IncludeVersions(includeVersions).XAnchoreAccount(xAnchoreAccount).Execute()

Get an application by application_id



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
    applicationId := "applicationId_example" // string | 
    includeVersions := true // bool |  (optional) (default to false)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetApplication(context.Background(), applicationId).IncludeVersions(includeVersions).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeVersions** | **bool** |  | [default to false]
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationVersion

> ApplicationVersion GetApplicationVersion(ctx, applicationId, applicationVersionId).XAnchoreAccount(xAnchoreAccount).Execute()

Get an application version



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
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetApplicationVersion(context.Background(), applicationId, applicationVersionId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationVersion`: ApplicationVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplicationVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**applicationVersionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ApplicationVersion**](ApplicationVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationVersionSbom

> ApplicationVersionSbom GetApplicationVersionSbom(ctx, applicationId, applicationVersionId).ArtifactTypes(artifactTypes).XAnchoreAccount(xAnchoreAccount).Execute()

Get the combined sbom for the given application version, optionally filtered by artifact type



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
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    artifactTypes := []string{"ArtifactTypes_example"} // []string |  (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetApplicationVersionSbom(context.Background(), applicationId, applicationVersionId).ArtifactTypes(artifactTypes).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplicationVersionSbom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationVersionSbom`: ApplicationVersionSbom
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplicationVersionSbom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**applicationVersionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationVersionSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **artifactTypes** | **[]string** |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ApplicationVersionSbom**](ApplicationVersionSbom.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationVersions

> []ApplicationVersion GetApplicationVersions(ctx, applicationId).XAnchoreAccount(xAnchoreAccount).Execute()

List all application verions



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
    applicationId := "applicationId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetApplicationVersions(context.Background(), applicationId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplicationVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationVersions`: []ApplicationVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplicationVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]ApplicationVersion**](ApplicationVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> []Application GetApplications(ctx).IncludeVersions(includeVersions).XAnchoreAccount(xAnchoreAccount).Execute()

List all applications



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
    includeVersions := true // bool |  (optional) (default to false)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetApplications(context.Background()).IncludeVersions(includeVersions).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: []Application
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeVersions** | **bool** |  | [default to false]
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceViolationAlert

> ComplianceViolationAlert GetComplianceViolationAlert(ctx, uuid).XAnchoreAccount(xAnchoreAccount).Execute()

Get compliance violation alert by id



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
    uuid := "uuid_example" // string | Identifier for the alert
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetComplianceViolationAlert(context.Background(), uuid).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetComplianceViolationAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComplianceViolationAlert`: ComplianceViolationAlert
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetComplianceViolationAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Identifier for the alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceViolationAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ComplianceViolationAlert**](ComplianceViolationAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceViolationAlerts

> []ComplianceViolationAlert GetComplianceViolationAlerts(ctx).Page(page).Limit(limit).State(state).CreatedAfter(createdAfter).CreatedBefore(createdBefore).ResourceImageDigest(resourceImageDigest).ResourceImageTag(resourceImageTag).ResourceRegistry(resourceRegistry).ResourceRepository(resourceRepository).XAnchoreAccount(xAnchoreAccount).Execute()

List all compliance violation alerts scoped to the account



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
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 100)
    state := "state_example" // string | Filter for alerts by current state, defaults to open alerts unless specified (optional) (default to "open")
    createdAfter := time.Now() // time.Time | Filter for alerts generated after the timestamp (optional)
    createdBefore := time.Now() // time.Time | Filter for alerts generated before the timestamp (optional)
    resourceImageDigest := "resourceImageDigest_example" // string | Filter for alerts associated with image digest (optional)
    resourceImageTag := "resourceImageTag_example" // string | Filter for alerts generated for the tag (optional)
    resourceRegistry := "resourceRegistry_example" // string | Filter for alerts associated with registry (optional)
    resourceRepository := "resourceRepository_example" // string | Filter for alerts associated with repository (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetComplianceViolationAlerts(context.Background()).Page(page).Limit(limit).State(state).CreatedAfter(createdAfter).CreatedBefore(createdBefore).ResourceImageDigest(resourceImageDigest).ResourceImageTag(resourceImageTag).ResourceRegistry(resourceRegistry).ResourceRepository(resourceRepository).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetComplianceViolationAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComplianceViolationAlerts`: []ComplianceViolationAlert
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetComplianceViolationAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceViolationAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **state** | **string** | Filter for alerts by current state, defaults to open alerts unless specified | [default to &quot;open&quot;]
 **createdAfter** | **time.Time** | Filter for alerts generated after the timestamp | 
 **createdBefore** | **time.Time** | Filter for alerts generated before the timestamp | 
 **resourceImageDigest** | **string** | Filter for alerts associated with image digest | 
 **resourceImageTag** | **string** | Filter for alerts generated for the tag | 
 **resourceRegistry** | **string** | Filter for alerts associated with registry | 
 **resourceRepository** | **string** | Filter for alerts associated with repository | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]ComplianceViolationAlert**](ComplianceViolationAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorrectionByUuid

> Correction GetCorrectionByUuid(ctx, uuid).XAnchoreAccount(xAnchoreAccount).Execute()

Retrieve a correction by UUID



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
    uuid := "uuid_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCorrectionByUuid(context.Background(), uuid).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCorrectionByUuid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCorrectionByUuid`: Correction
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCorrectionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorrectionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**Correction**](Correction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorrections

> []Correction GetCorrections(ctx).CorrectionType(correctionType).XAnchoreAccount(xAnchoreAccount).Execute()

Retrieve a list of corrections



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
    correctionType := "correctionType_example" // string |  (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCorrections(context.Background()).CorrectionType(correctionType).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCorrections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCorrections`: []Correction
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCorrections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCorrectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **correctionType** | **string** |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]Correction**](Correction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageAncestors

> []ImageAncestor GetImageAncestors(ctx, imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Return the list of ancestor images for the given image



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
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetImageAncestors(context.Background(), imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetImageAncestors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageAncestors`: []ImageAncestor
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetImageAncestors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageAncestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]ImageAncestor**](ImageAncestor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageInventory

> []InventoryItem GetImageInventory(ctx).InventoryType(inventoryType).ImageDigest(imageDigest).Context(context).XAnchoreAccount(xAnchoreAccount).Execute()

Return a list of the images in inventories for this account



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
    inventoryType := "inventoryType_example" // string |  (optional)
    imageDigest := "imageDigest_example" // string |  (optional)
    context := "context_example" // string |  (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetImageInventory(context.Background()).InventoryType(inventoryType).ImageDigest(imageDigest).Context(context).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetImageInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageInventory`: []InventoryItem
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetImageInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryType** | **string** |  | 
 **imageDigest** | **string** |  | 
 **context** | **string** |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]InventoryItem**](InventoryItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImagePolicyCheckByDigest

> []interface{} GetImagePolicyCheckByDigest(ctx, imageDigest).Tag(tag).PolicyId(policyId).Detail(detail).History(history).Interactive(interactive).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Check policy evaluation status for image



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
    tag := "tag_example" // string | 
    policyId := "policyId_example" // string |  (optional)
    detail := true // bool |  (optional) (default to true)
    history := true // bool |  (optional) (default to false)
    interactive := true // bool |  (optional) (default to false)
    baseDigest := "baseDigest_example" // string | Digest of a base image. If specified the evaluation will indicate results inherited from the base image (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetImagePolicyCheckByDigest(context.Background(), imageDigest).Tag(tag).PolicyId(policyId).Detail(detail).History(history).Interactive(interactive).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetImagePolicyCheckByDigest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImagePolicyCheckByDigest`: []interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetImagePolicyCheckByDigest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagePolicyCheckByDigestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **string** |  | 
 **policyId** | **string** |  | 
 **detail** | **bool** |  | [default to true]
 **history** | **bool** |  | [default to false]
 **interactive** | **bool** |  | [default to false]
 **baseDigest** | **string** | Digest of a base image. If specified the evaluation will indicate results inherited from the base image | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageVulnerabilitiesByDigest

> EnterpriseVulnerabilityResponse GetImageVulnerabilitiesByDigest(ctx, imageDigest, vtype).ForceRefresh(forceRefresh).VendorOnly(vendorOnly).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get vulnerabilities by type

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
    vtype := "vtype_example" // string | 
    forceRefresh := true // bool |  (optional) (default to false)
    vendorOnly := true // bool | Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data. When set to true, it will filter out all vulnerabilities where `will_not_fix` is False. If false all vulnerabilities are returned regardless of `will_not_fix` (optional) (default to true)
    baseDigest := "baseDigest_example" // string | Digest of a base image. If specified the vulnerabilities will indicate inheritance from the base image (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetImageVulnerabilitiesByDigest(context.Background(), imageDigest, vtype).ForceRefresh(forceRefresh).VendorOnly(vendorOnly).BaseDigest(baseDigest).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetImageVulnerabilitiesByDigest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageVulnerabilitiesByDigest`: EnterpriseVulnerabilityResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetImageVulnerabilitiesByDigest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string** |  | 
**vtype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageVulnerabilitiesByDigestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceRefresh** | **bool** |  | [default to false]
 **vendorOnly** | **bool** | Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data. When set to true, it will filter out all vulnerabilities where &#x60;will_not_fix&#x60; is False. If false all vulnerabilities are returned regardless of &#x60;will_not_fix&#x60; | [default to true]
 **baseDigest** | **string** | Digest of a base image. If specified the vulnerabilities will indicate inheritance from the base image | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**EnterpriseVulnerabilityResponse**](EnterpriseVulnerabilityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
    resp, r, err := apiClient.DefaultApi.GetImportSourcesSbom(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetImportSourcesSbom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImportSourcesSbom`: SourceImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetImportSourcesSbom`: %v\n", resp)
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


## GetInventoryClusterByName

> InventoryCluster GetInventoryClusterByName(ctx, clusterName).XAnchoreAccount(xAnchoreAccount).Execute()

Return a configured inventory cluster



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
    clusterName := "clusterName_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetInventoryClusterByName(context.Background(), clusterName).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetInventoryClusterByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInventoryClusterByName`: InventoryCluster
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetInventoryClusterByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoryClusterByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**InventoryCluster**](InventoryCluster.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperation`: SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOperation`: %v\n", resp)
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


## GetRuntimeComplianceChecks

> []RuntimeComplianceCheck GetRuntimeComplianceChecks(ctx).ImageDigest(imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Get all runtime compliance checks or just those for a given image digest



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
    imageDigest := "imageDigest_example" // string |  (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRuntimeComplianceChecks(context.Background()).ImageDigest(imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRuntimeComplianceChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuntimeComplianceChecks`: []RuntimeComplianceCheck
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRuntimeComplianceChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRuntimeComplianceChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageDigest** | **string** |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RuntimeComplianceCheck**](RuntimeComplianceCheck.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuntimeComplianceResult

> *os.File GetRuntimeComplianceResult(ctx, complianceFileId).XAnchoreAccount(xAnchoreAccount).Execute()

Check the results of a a specific runtime compliance check



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
    complianceFileId := "complianceFileId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRuntimeComplianceResult(context.Background(), complianceFileId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRuntimeComplianceResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuntimeComplianceResult`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRuntimeComplianceResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**complianceFileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuntimeComplianceResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSource(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSource`: SourceManifest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSource`: %v\n", resp)
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

> ContentPackageResponse GetSourceContentByType(ctx, sourceId, contentType).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSourceContentByType(context.Background(), sourceId, contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSourceContentByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceContentByType`: ContentPackageResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSourceContentByType`: %v\n", resp)
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

[**ContentPackageResponse**](ContentPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceContentTypes

> GetSourceContentTypes(ctx, sourceId).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSourceContentTypes(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSourceContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSourceSbomNative(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSourceSbomNative``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceSbomNative`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSourceSbomNative`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSourceSbomTypes(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSourceSbomTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceSbomTypes`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSourceSbomTypes`: %v\n", resp)
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

> GetSourceVulnerabilities(ctx, sourceId, vtype).ForceRefresh(forceRefresh).WillNotFix(willNotFix).XAnchoreAccount(xAnchoreAccount).Execute()

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSourceVulnerabilities(context.Background(), sourceId, vtype).ForceRefresh(forceRefresh).WillNotFix(willNotFix).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSourceVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSourceVulnerabilityTypes(context.Background(), sourceId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSourceVulnerabilityTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceVulnerabilityTypes`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSourceVulnerabilityTypes`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.InvalidateOperation(context.Background(), operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InvalidateOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvalidateOperation`: SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InvalidateOperation`: %v\n", resp)
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


## ListArtifacts

> ArtifactListResponse ListArtifacts(ctx, applicationId, applicationVersionId).ArtifactTypes(artifactTypes).XAnchoreAccount(xAnchoreAccount).Execute()

List artifacts present on a given application version



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
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    artifactTypes := []string{"ArtifactTypes_example"} // []string |  (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListArtifacts(context.Background(), applicationId, applicationVersionId).ArtifactTypes(artifactTypes).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifacts`: ArtifactListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**applicationVersionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **artifactTypes** | **[]string** |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ArtifactListResponse**](ArtifactListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInventoryClusters

> []InventoryCluster ListInventoryClusters(ctx).InventoryType(inventoryType).XAnchoreAccount(xAnchoreAccount).Execute()

Return a list of the configured inventory clusters



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
    inventoryType := "inventoryType_example" // string |  (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListInventoryClusters(context.Background()).InventoryType(inventoryType).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListInventoryClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInventoryClusters`: []InventoryCluster
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListInventoryClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInventoryClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryType** | **string** |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]InventoryCluster**](InventoryCluster.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListOperations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOperations`: []SourceImportOperation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListOperations`: %v\n", resp)
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSources`: []Source
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSources`: %v\n", resp)
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


## RemoveArtifactFromApplicationVersion

> RemoveArtifactFromApplicationVersion(ctx, applicationId, applicationVersionId, associationId).XAnchoreAccount(xAnchoreAccount).Execute()

Delete an artifact from specified application version



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
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    associationId := "associationId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RemoveArtifactFromApplicationVersion(context.Background(), applicationId, applicationVersionId, associationId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RemoveArtifactFromApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**applicationVersionId** | **string** |  | 
**associationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveArtifactFromApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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


## SyncImageInventory

> []InventoryItem SyncImageInventory(ctx).Inventory(inventory).XAnchoreAccount(xAnchoreAccount).Execute()

synchronizes the list of the images in a given cluster for the inventory



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
    inventory := *openapiclient.NewInventoryReport() // InventoryReport | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SyncImageInventory(context.Background()).Inventory(inventory).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SyncImageInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncImageInventory`: []InventoryItem
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SyncImageInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncImageInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventory** | [**InventoryReport**](InventoryReport.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]InventoryItem**](InventoryItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> Application UpdateApplication(ctx, applicationId).Application(application).XAnchoreAccount(xAnchoreAccount).Execute()

Update application details



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
    applicationId := "applicationId_example" // string | 
    application := *openapiclient.NewApplication() // Application | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateApplication(context.Background(), applicationId).Application(application).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **application** | [**Application**](Application.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationVersion

> ApplicationVersion UpdateApplicationVersion(ctx, applicationId, applicationVersionId).ApplicationVersion(applicationVersion).XAnchoreAccount(xAnchoreAccount).Execute()

Update application version details



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
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    applicationVersion := *openapiclient.NewApplicationVersion("VersionName_example") // ApplicationVersion | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateApplicationVersion(context.Background(), applicationId, applicationVersionId).ApplicationVersion(applicationVersion).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationVersion`: ApplicationVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateApplicationVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**applicationVersionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationVersion** | [**ApplicationVersion**](ApplicationVersion.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ApplicationVersion**](ApplicationVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComplianceViolationAlertState

> ComplianceViolationAlert UpdateComplianceViolationAlertState(ctx, uuid, state).XAnchoreAccount(xAnchoreAccount).Execute()

Open or close a compliance violation alert



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
    uuid := "uuid_example" // string | Identifier for the alert
    state := "state_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateComplianceViolationAlertState(context.Background(), uuid, state).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateComplianceViolationAlertState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateComplianceViolationAlertState`: ComplianceViolationAlert
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateComplianceViolationAlertState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Identifier for the alert | 
**state** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComplianceViolationAlertStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ComplianceViolationAlert**](ComplianceViolationAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCorrectionByUuid

> Correction UpdateCorrectionByUuid(ctx, uuid).Correction(correction).XAnchoreAccount(xAnchoreAccount).Execute()

Update a correction by UUID



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
    uuid := "uuid_example" // string | 
    correction := *openapiclient.NewCorrection("Type_example", *openapiclient.NewCorrectionMatch("npm"), []openapiclient.CorrectionFieldMatch{*openapiclient.NewCorrectionFieldMatch("name", "FieldValue_example")}) // Correction | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateCorrectionByUuid(context.Background(), uuid).Correction(correction).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCorrectionByUuid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCorrectionByUuid`: Correction
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCorrectionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCorrectionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **correction** | [**Correction**](Correction.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**Correction**](Correction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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
    sbom := *openapiclient.NewNativeSBOM([]openapiclient.NativeSBOMPackage{*openapiclient.NewNativeSBOMPackage("Name_example", "Version_example", "Type_example", []openapiclient.NativeSBOMPackageLocation{*openapiclient.NewNativeSBOMPackageLocation("Path_example")}, []string{"Licenses_example"}, "Language_example", []string{"Cpes_example"})}, *openapiclient.NewNativeSBOMSource("Type_example", interface{}(123)), *openapiclient.NewNativeSBOMDistribution("Name_example", "Version_example", "IdLike_example")) // NativeSBOM | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UploadImportSourcesSbom(context.Background(), operationId).Sbom(sbom).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UploadImportSourcesSbom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadImportSourcesSbom`: SourceImportContentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UploadImportSourcesSbom`: %v\n", resp)
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

