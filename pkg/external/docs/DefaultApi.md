# \DefaultApi

All URIs are relative to *http://localhost/enterprise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddActionPlan**](DefaultApi.md#AddActionPlan) | **Post** /actions | Submits an Action Plan
[**AddCorrection**](DefaultApi.md#AddCorrection) | **Post** /corrections | Create a correction record
[**AddInventoryCluster**](DefaultApi.md#AddInventoryCluster) | **Post** /inventories/clusters | Create a cluster inventory
[**AddRuntimeComplianceCheck**](DefaultApi.md#AddRuntimeComplianceCheck) | **Post** /runtime_compliance | Post a runtime compliance check
[**DelInventoryClusterByName**](DefaultApi.md#DelInventoryClusterByName) | **Delete** /inventories/clusters/{cluster_name} | Delete a configured inventory clusters by cluster_name
[**DeleteCorrectionByUuid**](DefaultApi.md#DeleteCorrectionByUuid) | **Delete** /corrections/{uuid} | Delete a correction by UUID
[**GetActionPlans**](DefaultApi.md#GetActionPlans) | **Get** /actions | Gets a list of submitted action (remediation) plans
[**GetAlertSummaries**](DefaultApi.md#GetAlertSummaries) | **Get** /alerts/summaries | List all alert summaries scoped to the account
[**GetComplianceViolationAlert**](DefaultApi.md#GetComplianceViolationAlert) | **Get** /alerts/compliance_violations/{uuid} | Get compliance violation alert by id
[**GetComplianceViolationAlerts**](DefaultApi.md#GetComplianceViolationAlerts) | **Get** /alerts/compliance_violations | List all compliance violation alerts scoped to the account
[**GetCorrectionByUuid**](DefaultApi.md#GetCorrectionByUuid) | **Get** /corrections/{uuid} | Retrieve a correction by UUID
[**GetCorrections**](DefaultApi.md#GetCorrections) | **Get** /corrections | Retrieve a list of corrections
[**GetImageAncestors**](DefaultApi.md#GetImageAncestors) | **Get** /images/{image_digest}/ancestors | Return the list of ancestor images for the given image
[**GetImageInventory**](DefaultApi.md#GetImageInventory) | **Get** /inventories | Return a list of the images in inventories for this account
[**GetImagePolicyCheckByDigest**](DefaultApi.md#GetImagePolicyCheckByDigest) | **Get** /images/{imageDigest}/check | Check policy evaluation status for image
[**GetImageVulnerabilitiesByDigest**](DefaultApi.md#GetImageVulnerabilitiesByDigest) | **Get** /images/{imageDigest}/vuln/{vtype} | Get vulnerabilities by type
[**GetInventoryClusterByName**](DefaultApi.md#GetInventoryClusterByName) | **Get** /inventories/clusters/{cluster_name} | Return a configured inventory cluster
[**GetRuntimeComplianceChecks**](DefaultApi.md#GetRuntimeComplianceChecks) | **Get** /runtime_compliance | Get all runtime compliance checks or just those for a given image digest
[**GetRuntimeComplianceResult**](DefaultApi.md#GetRuntimeComplianceResult) | **Get** /runtime_compliance/result/{compliance_file_id} | Check the results of a a specific runtime compliance check
[**ListInventoryClusters**](DefaultApi.md#ListInventoryClusters) | **Get** /inventories/clusters | Return a list of the configured inventory clusters
[**SyncImageInventory**](DefaultApi.md#SyncImageInventory) | **Post** /inventories | synchronizes the list of the images in a given cluster for the inventory
[**UpdateComplianceViolationAlertState**](DefaultApi.md#UpdateComplianceViolationAlertState) | **Put** /alerts/compliance_violations/{uuid}/{state} | Open or close a compliance violation alert
[**UpdateCorrectionByUuid**](DefaultApi.md#UpdateCorrectionByUuid) | **Put** /corrections/{uuid} | Update a correction by UUID



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

