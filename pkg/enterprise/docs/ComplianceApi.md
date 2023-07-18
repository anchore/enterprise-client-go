# \ComplianceApi

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRuntimeComplianceCheck**](ComplianceApi.md#AddRuntimeComplianceCheck) | **Post** /runtime-compliance | Post a runtime compliance check
[**GetRuntimeComplianceChecks**](ComplianceApi.md#GetRuntimeComplianceChecks) | **Get** /runtime-compliance | Get all runtime compliance checks or just those for a given image digest
[**GetRuntimeComplianceResult**](ComplianceApi.md#GetRuntimeComplianceResult) | **Get** /runtime-compliance/result/{compliance_file_id} | Check the results of a a specific runtime compliance check



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComplianceApi.AddRuntimeComplianceCheck(context.Background()).CheckType(checkType).ImageDigest(imageDigest).XAnchoreAccount(xAnchoreAccount).Result(result).Pod(pod).Namespace(namespace).ImageTag(imageTag).StartTime(startTime).EndTime(endTime).ResultFile(resultFile).ReportFile(reportFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComplianceApi.AddRuntimeComplianceCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRuntimeComplianceCheck`: RuntimeComplianceCheck
    fmt.Fprintf(os.Stdout, "Response from `ComplianceApi.AddRuntimeComplianceCheck`: %v\n", resp)
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComplianceApi.GetRuntimeComplianceChecks(context.Background()).ImageDigest(imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComplianceApi.GetRuntimeComplianceChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuntimeComplianceChecks`: []RuntimeComplianceCheck
    fmt.Fprintf(os.Stdout, "Response from `ComplianceApi.GetRuntimeComplianceChecks`: %v\n", resp)
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComplianceApi.GetRuntimeComplianceResult(context.Background(), complianceFileId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComplianceApi.GetRuntimeComplianceResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuntimeComplianceResult`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ComplianceApi.GetRuntimeComplianceResult`: %v\n", resp)
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

