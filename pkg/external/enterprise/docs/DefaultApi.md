# \DefaultApi

All URIs are relative to *http://localhost/enterprise*

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
[**GetSourcePolicyCheck**](DefaultApi.md#GetSourcePolicyCheck) | **Get** /sources/{source_id}/check | Fetch or calculate policy evaluation for a source
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

> ActionPlan AddActionPlan(ctx, actionPlan)

Submits an Action Plan

Submits an Action Plan and saves upon completion

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionPlan** | [**ActionPlan**](ActionPlan.md)|  | 

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

> Application AddApplication(ctx, application, optional)

Create an application

Create an application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | [**Application**](Application.md)|  | 
 **optional** | ***AddApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> ApplicationVersion AddApplicationVersion(ctx, applicationId, applicationVersion, optional)

Create an application version

Create an application version

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
**applicationVersion** | [**ApplicationVersion**](ApplicationVersion.md)|  | 
 **optional** | ***AddApplicationVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddApplicationVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> ArtifactAssociationResponse AddArtifactToApplicationVersion(ctx, applicationId, applicationVersionId, artifactRequest, optional)

Add an artifact to an application version

Add artifact to given application_id and application_version_id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
**applicationVersionId** | **string**|  | 
**artifactRequest** | [**ArtifactAssociationRequest**](ArtifactAssociationRequest.md)|  | 
 **optional** | ***AddArtifactToApplicationVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddArtifactToApplicationVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []Correction AddCorrection(ctx, correction, optional)

Create a correction record

Add a correction record that will be used to fix

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**correction** | [**Correction**](Correction.md)|  | 
 **optional** | ***AddCorrectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCorrectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> InventoryCluster AddInventoryCluster(ctx, cluster, optional)

Create a cluster inventory

Create a new cluster inventory with the provided configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | [**InventoryCluster**](InventoryCluster.md)|  | 
 **optional** | ***AddInventoryClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddInventoryClusterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> RuntimeComplianceCheck AddRuntimeComplianceCheck(ctx, checkType, imageDigest, optional)

Post a runtime compliance check

Post a runtime compliance check

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkType** | **string**| The type of runtime compliance check | 
**imageDigest** | **string**| The digest of the pod the check was run against | 
 **optional** | ***AddRuntimeComplianceCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddRuntimeComplianceCheckOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 
 **result** | **optional.String**| The result of the runtime compliance check | 
 **pod** | **optional.String**| The pod the check was run against | 
 **namespace** | **optional.String**| The namespace of the pod the check was run against | 
 **imageTag** | **optional.String**| The tag of the image in the pod the check was run against | 
 **startTime** | **optional.Time**| The type of runtime compliance check | 
 **endTime** | **optional.Time**| The type of runtime compliance check | 
 **resultFile** | **optional.Interface of *os.File****optional.*os.File**| The file with the check results | 
 **reportFile** | **optional.Interface of *os.File****optional.*os.File**| The file with the check port | 

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

> SourceImportOperation CreateOperation(ctx, )

Begin the import of a source code repository analyzed by Syft into the system

### Required Parameters

This endpoint does not need any parameter.

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

> DelInventoryClusterByName(ctx, clusterName, optional)

Delete a configured inventory clusters by cluster_name

Removes a configured cluster for reporting image inventory by cluster_name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string**|  | 
 **optional** | ***DelInventoryClusterByNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DelInventoryClusterByNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> DeleteApplication(ctx, applicationId, optional)

Delete an application by application_id

Delete an application by application_id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
 **optional** | ***DeleteApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> DeleteApplicationVersion(ctx, applicationId, applicationVersionId, optional)

Delete an application version by application_id and application_version_id

Delete an application version by application_id and application_version_id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
**applicationVersionId** | **string**|  | 
 **optional** | ***DeleteApplicationVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteApplicationVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> DeleteCorrectionByUuid(ctx, uuid, optional)

Delete a correction by UUID

Delete a single correction, looked up via it's uuid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***DeleteCorrectionByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteCorrectionByUuidOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> SourceManifest DeleteSource(ctx, sourceId, optional)

Delete source record from DB

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string**| UUID of source to delete | 
 **optional** | ***DeleteSourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| force delete | 

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

> SourceManifest FinalizeOperation(ctx, operationId, metadata)

Add source records to catalog db

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 
**metadata** | [**SourceImportMetadata**](SourceImportMetadata.md)|  | 

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

> []ActionPlan GetActionPlans(ctx, optional)

Gets a list of submitted action (remediation) plans

Retrieves a list of action plans that have been completed

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetActionPlansOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetActionPlansOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageTag** | **optional.String**|  | 
 **imageDigest** | **optional.String**|  | 
 **createdAfter** | **optional.Time**| RFC 3339 formatted UTC timestamp to filter out action plans that were only created after this date | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []AlertSummary GetAlertSummaries(ctx, optional)

List all alert summaries scoped to the account

Returns a paginated list of alert summaries in chronological order from the most to least recently generated alerts. Return alerts in the open state by default. Use query parameters for filtering

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAlertSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAlertSummariesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | [default to 1]
 **limit** | **optional.Int32**|  | [default to 100]
 **type_** | **optional.String**| Filter for alerts based on the type such as compliance violation | [default to all]
 **state** | **optional.String**| Filter for alerts by current state, defaults to open alerts unless specified | [default to open]
 **createdAfter** | **optional.Time**| Filter for alerts generated after the timestamp | 
 **createdBefore** | **optional.Time**| Filter for alerts generated before the timestamp | 
 **resourceLabel** | [**optional.Interface of []string**](string.md)| Filter for alerts associated with a resource where the label in key&#x3D;value format such as tag&#x3D;docker.io/library/alpine:latest or repository&#x3D;library/alpine | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> Application GetApplication(ctx, applicationId, optional)

Get an application by application_id

Get an application by application_id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
 **optional** | ***GetApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeVersions** | **optional.Bool**|  | [default to false]
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> ApplicationVersion GetApplicationVersion(ctx, applicationId, applicationVersionId, optional)

Get an application version

Get an application version by application_id and application_version_id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
**applicationVersionId** | **string**|  | 
 **optional** | ***GetApplicationVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> ApplicationVersionSbom GetApplicationVersionSbom(ctx, applicationId, applicationVersionId, optional)

Get the combined sbom for the given application version, optionally filtered by artifact type

Get the combined sbom for the given application version, optionally filtered by artifact type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
**applicationVersionId** | **string**|  | 
 **optional** | ***GetApplicationVersionSbomOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationVersionSbomOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **artifactTypes** | [**optional.Interface of []string**](string.md)|  | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []ApplicationVersion GetApplicationVersions(ctx, applicationId, optional)

List all application verions

List all application verions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
 **optional** | ***GetApplicationVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []Application GetApplications(ctx, optional)

List all applications

List all applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeVersions** | **optional.Bool**|  | [default to false]
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> ComplianceViolationAlert GetComplianceViolationAlert(ctx, uuid, optional)

Get compliance violation alert by id

Returns a single compliance violation alert object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Identifier for the alert | 
 **optional** | ***GetComplianceViolationAlertOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetComplianceViolationAlertOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []ComplianceViolationAlert GetComplianceViolationAlerts(ctx, optional)

List all compliance violation alerts scoped to the account

Returns a paginated list of compliance violation alerts in chronological order from the most to least recently generated alerts. Return alerts in the open state by default. Use query parameters for filtering

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetComplianceViolationAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetComplianceViolationAlertsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | [default to 1]
 **limit** | **optional.Int32**|  | [default to 100]
 **state** | **optional.String**| Filter for alerts by current state, defaults to open alerts unless specified | [default to open]
 **createdAfter** | **optional.Time**| Filter for alerts generated after the timestamp | 
 **createdBefore** | **optional.Time**| Filter for alerts generated before the timestamp | 
 **resourceImageDigest** | **optional.String**| Filter for alerts associated with image digest | 
 **resourceImageTag** | **optional.String**| Filter for alerts generated for the tag | 
 **resourceRegistry** | **optional.String**| Filter for alerts associated with registry | 
 **resourceRepository** | **optional.String**| Filter for alerts associated with repository | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> Correction GetCorrectionByUuid(ctx, uuid, optional)

Retrieve a correction by UUID

Returns a single correction, looked up via it's uuid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***GetCorrectionByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCorrectionByUuidOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []Correction GetCorrections(ctx, optional)

Retrieve a list of corrections

Returns a list of corrections

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCorrectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCorrectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **correctionType** | **optional.String**|  | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []ImageAncestor GetImageAncestors(ctx, imageDigest, optional)

Return the list of ancestor images for the given image

Returns list of ancestor images, which are the images that form the base layers of the image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
 **optional** | ***GetImageAncestorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageAncestorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []InventoryItem GetImageInventory(ctx, optional)

Return a list of the images in inventories for this account

Returns a list of the images that are in use

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetImageInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageInventoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryType** | **optional.String**|  | 
 **imageDigest** | **optional.String**|  | 
 **context** | **optional.String**|  | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []interface{} GetImagePolicyCheckByDigest(ctx, imageDigest, tag, optional)

Check policy evaluation status for image

Get the policy evaluation for the given image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
**tag** | **string**|  | 
 **optional** | ***GetImagePolicyCheckByDigestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImagePolicyCheckByDigestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyId** | **optional.String**|  | 
 **detail** | **optional.Bool**|  | [default to true]
 **history** | **optional.Bool**|  | [default to false]
 **interactive** | **optional.Bool**|  | [default to false]
 **baseDigest** | **optional.String**| Digest of a base image. If specified the evaluation will indicate results inherited from the base image | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> EnterpriseVulnerabilityResponse GetImageVulnerabilitiesByDigest(ctx, imageDigest, vtype, optional)

Get vulnerabilities by type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
**vtype** | **string**|  | 
 **optional** | ***GetImageVulnerabilitiesByDigestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageVulnerabilitiesByDigestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceRefresh** | **optional.Bool**|  | [default to false]
 **vendorOnly** | **optional.Bool**| Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data. When set to true, it will filter out all vulnerabilities where &#x60;will_not_fix&#x60; is False. If false all vulnerabilities are returned regardless of &#x60;will_not_fix&#x60; | [default to true]
 **baseDigest** | **optional.String**| Digest of a base image. If specified the vulnerabilities will indicate inheritance from the base image | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> SourceImportContentResponse GetImportSourcesSbom(ctx, operationId)

list the packages of an imported source code repository

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 

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

> InventoryCluster GetInventoryClusterByName(ctx, clusterName, optional)

Return a configured inventory cluster

Returns a cluster that is configured for reporting image inventory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string**|  | 
 **optional** | ***GetInventoryClusterByNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInventoryClusterByNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> SourceImportOperation GetOperation(ctx, operationId)

Get detail on a single import

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 

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

> []RuntimeComplianceCheck GetRuntimeComplianceChecks(ctx, optional)

Get all runtime compliance checks or just those for a given image digest

Get all runtime compliance checks or just those for a given image digest

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetRuntimeComplianceChecksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRuntimeComplianceChecksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageDigest** | **optional.String**|  | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> *os.File GetRuntimeComplianceResult(ctx, complianceFileId, optional)

Check the results of a a specific runtime compliance check

Get the results of a specific runtime compliance check

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**complianceFileId** | **string**|  | 
 **optional** | ***GetRuntimeComplianceResultOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRuntimeComplianceResultOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> SourceManifest GetSource(ctx, sourceId)

Get a detailed source repository analysis metadata record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string**|  | 

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

> SourceContentPackageResponse GetSourceContentByType(ctx, sourceId, contentType)

Get the content of an analyzed source repository

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string**|  | 
**contentType** | **string**|  | 

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

> []string GetSourceContentTypes(ctx, sourceId)

Get a detailed source repository analysis metadata record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string**|  | 

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

> []interface{} GetSourcePolicyCheck(ctx, sourceId, optional)

Fetch or calculate policy evaluation for a source

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string**| UUID of source to get | 
 **optional** | ***GetSourcePolicyCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSourcePolicyCheckOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyId** | **optional.String**|  | 
 **newestOnly** | **optional.Bool**|  | 
 **interactive** | **optional.Bool**|  | 
 **historyOnly** | **optional.Bool**|  | 

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


## GetSourceSbomNative

> *os.File GetSourceSbomNative(ctx, sourceId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string**|  | 

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

> []string GetSourceSbomTypes(ctx, sourceId)

Get a detailed source repository analysis metadata record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string**|  | 

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

> SourceVulnerabilitiesResponse GetSourceVulnerabilities(ctx, sourceId, vtype, optional)

Get vulnerabilities for the source by type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string**|  | 
**vtype** | **string**|  | 
 **optional** | ***GetSourceVulnerabilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSourceVulnerabilitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceRefresh** | **optional.Bool**|  | 
 **willNotFix** | **optional.Bool**| Vulnerability data publishers explicitly won&#39;t fix some vulnerabilities. This is captured by will_not_fix attribute of each result. If the query parameter is set, results matching it&#39;s value will be filtered. Results are not filtered if the query parameter is unset | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []string GetSourceVulnerabilityTypes(ctx, sourceId, optional)

Get the available vulnerability types for source

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string**|  | 
 **optional** | ***GetSourceVulnerabilityTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSourceVulnerabilityTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> SourceImportOperation InvalidateOperation(ctx, operationId)

Invalidate operation ID so it can be garbage collected

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 

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

> ArtifactListResponse ListArtifacts(ctx, applicationId, applicationVersionId, optional)

List artifacts present on a given application version

List artifacts present on a given application version

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
**applicationVersionId** | **string**|  | 
 **optional** | ***ListArtifactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListArtifactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **artifactTypes** | [**optional.Interface of []string**](string.md)|  | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []InventoryCluster ListInventoryClusters(ctx, optional)

Return a list of the configured inventory clusters

Returns a filterable list of the clusters that are configured for reporting image inventory

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListInventoryClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInventoryClustersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryType** | **optional.String**|  | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []SourceImportOperation ListOperations(ctx, )

Lists in-progress imports

### Required Parameters

This endpoint does not need any parameter.

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

> []Source ListSources(ctx, )

List the source repository analysis records

### Required Parameters

This endpoint does not need any parameter.

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

> RemoveArtifactFromApplicationVersion(ctx, applicationId, applicationVersionId, associationId, optional)

Delete an artifact from specified application version

Delete an artifact from specified application version

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
**applicationVersionId** | **string**|  | 
**associationId** | **string**|  | 
 **optional** | ***RemoveArtifactFromApplicationVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveArtifactFromApplicationVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> []InventoryItem SyncImageInventory(ctx, inventory, optional)

synchronizes the list of the images in a given cluster for the inventory

synchronizes the list of the images that are in use

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventory** | [**InventoryReport**](InventoryReport.md)|  | 
 **optional** | ***SyncImageInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncImageInventoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> Application UpdateApplication(ctx, applicationId, application, optional)

Update application details

Updates application details for given application_id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
**application** | [**Application**](Application.md)|  | 
 **optional** | ***UpdateApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> ApplicationVersion UpdateApplicationVersion(ctx, applicationId, applicationVersionId, applicationVersion, optional)

Update application version details

Updates application version details for given application_id and application_version_id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string**|  | 
**applicationVersionId** | **string**|  | 
**applicationVersion** | [**ApplicationVersion**](ApplicationVersion.md)|  | 
 **optional** | ***UpdateApplicationVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApplicationVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> ComplianceViolationAlert UpdateComplianceViolationAlertState(ctx, uuid, state, optional)

Open or close a compliance violation alert

Idempotent op for changing the alert state to open or closed

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Identifier for the alert | 
**state** | **string**|  | 
 **optional** | ***UpdateComplianceViolationAlertStateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateComplianceViolationAlertStateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> Correction UpdateCorrectionByUuid(ctx, uuid, correction, optional)

Update a correction by UUID

Updates a single correction, looked up via it's uuid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
**correction** | [**Correction**](Correction.md)|  | 
 **optional** | ***UpdateCorrectionByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCorrectionByUuidOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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

> SourceImportContentResponse UploadImportSourcesSbom(ctx, operationId, sbom)

Begin the import of a source code repository analyzed by Syft into the system

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 
**sbom** | [**NativeSbom**](NativeSbom.md)|  | 

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
