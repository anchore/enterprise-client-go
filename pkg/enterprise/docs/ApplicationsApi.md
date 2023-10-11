# \ApplicationsAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplication**](ApplicationsAPI.md#AddApplication) | **Post** /applications | Create an application
[**AddApplicationVersion**](ApplicationsAPI.md#AddApplicationVersion) | **Post** /applications/{application_id}/versions | Create an application version
[**AddArtifactToApplicationVersion**](ApplicationsAPI.md#AddArtifactToApplicationVersion) | **Post** /applications/{application_id}/versions/{application_version_id}/artifacts | Add an artifact to an application version
[**DeleteApplication**](ApplicationsAPI.md#DeleteApplication) | **Delete** /applications/{application_id} | Delete an application by application_id
[**DeleteApplicationVersion**](ApplicationsAPI.md#DeleteApplicationVersion) | **Delete** /applications/{application_id}/versions/{application_version_id} | Delete an application version by application_id and application_version_id
[**GetApplication**](ApplicationsAPI.md#GetApplication) | **Get** /applications/{application_id} | Get an application by application_id
[**GetApplicationVersion**](ApplicationsAPI.md#GetApplicationVersion) | **Get** /applications/{application_id}/versions/{application_version_id} | Get an application version
[**GetApplicationVersionSbom**](ApplicationsAPI.md#GetApplicationVersionSbom) | **Get** /applications/{application_id}/versions/{application_version_id}/sboms/native-json | Get the combined sbom for the given application version, optionally filtered by artifact type
[**GetApplicationVersionVulnerabilities**](ApplicationsAPI.md#GetApplicationVersionVulnerabilities) | **Get** /applications/{application_id}/versions/{application_version_id}/vulnerabilities | Get the vulnerabilities for a given application version
[**GetApplicationVersions**](ApplicationsAPI.md#GetApplicationVersions) | **Get** /applications/{application_id}/versions | List all application verions
[**GetApplications**](ApplicationsAPI.md#GetApplications) | **Get** /applications | List all applications
[**ListArtifacts**](ApplicationsAPI.md#ListArtifacts) | **Get** /applications/{application_id}/versions/{application_version_id}/artifacts | List artifacts present on a given application version
[**RemoveArtifactFromApplicationVersion**](ApplicationsAPI.md#RemoveArtifactFromApplicationVersion) | **Delete** /applications/{application_id}/versions/{application_version_id}/artifacts/{association_id} | Delete an artifact from specified application version
[**UpdateApplication**](ApplicationsAPI.md#UpdateApplication) | **Put** /applications/{application_id} | Update application details
[**UpdateApplicationVersion**](ApplicationsAPI.md#UpdateApplicationVersion) | **Put** /applications/{application_id}/versions/{application_version_id} | Update application version details



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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    application := *openapiclient.NewApplication() // Application | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.AddApplication(context.Background()).Application(application).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.AddApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.AddApplication`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    applicationVersion := *openapiclient.NewApplicationVersion("VersionName_example") // ApplicationVersion | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.AddApplicationVersion(context.Background(), applicationId).ApplicationVersion(applicationVersion).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.AddApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApplicationVersion`: ApplicationVersion
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.AddApplicationVersion`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    artifactRequest := *openapiclient.NewArtifactAssociationRequest("ArtifactType_example", interface{}(123)) // ArtifactAssociationRequest | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.AddArtifactToApplicationVersion(context.Background(), applicationId, applicationVersionId).ArtifactRequest(artifactRequest).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.AddArtifactToApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddArtifactToApplicationVersion`: ArtifactAssociationResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.AddArtifactToApplicationVersion`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsAPI.DeleteApplication(context.Background(), applicationId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplication``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsAPI.DeleteApplicationVersion(context.Background(), applicationId, applicationVersionId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplicationVersion``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    includeVersions := true // bool |  (optional) (default to false)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetApplication(context.Background(), applicationId).IncludeVersions(includeVersions).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplication`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetApplicationVersion(context.Background(), applicationId, applicationVersionId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationVersion`: ApplicationVersion
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationVersion`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    artifactTypes := []string{"ArtifactTypes_example"} // []string |  (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetApplicationVersionSbom(context.Background(), applicationId, applicationVersionId).ArtifactTypes(artifactTypes).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationVersionSbom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationVersionSbom`: ApplicationVersionSbom
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationVersionSbom`: %v\n", resp)
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


## GetApplicationVersionVulnerabilities

> ApplicationVersionVulnerabilityReport GetApplicationVersionVulnerabilities(ctx, applicationId, applicationVersionId).WillNotFix(willNotFix).XAnchoreAccount(xAnchoreAccount).Execute()

Get the vulnerabilities for a given application version



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
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    willNotFix := true // bool | If true, include vulnerabilities that the vendor of an image distribution either disagrees with or does not intend to prioritize for remediation (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetApplicationVersionVulnerabilities(context.Background(), applicationId, applicationVersionId).WillNotFix(willNotFix).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationVersionVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationVersionVulnerabilities`: ApplicationVersionVulnerabilityReport
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationVersionVulnerabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**applicationVersionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationVersionVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **willNotFix** | **bool** | If true, include vulnerabilities that the vendor of an image distribution either disagrees with or does not intend to prioritize for remediation | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ApplicationVersionVulnerabilityReport**](ApplicationVersionVulnerabilityReport.md)

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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetApplicationVersions(context.Background(), applicationId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationVersions`: []ApplicationVersion
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationVersions`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    includeVersions := true // bool |  (optional) (default to false)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetApplications(context.Background()).IncludeVersions(includeVersions).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: []Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplications`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    artifactTypes := []string{"ArtifactTypes_example"} // []string |  (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.ListArtifacts(context.Background(), applicationId, applicationVersionId).ArtifactTypes(artifactTypes).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ListArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifacts`: ArtifactListResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.ListArtifacts`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    associationId := "associationId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsAPI.RemoveArtifactFromApplicationVersion(context.Background(), applicationId, applicationVersionId, associationId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.RemoveArtifactFromApplicationVersion``: %v\n", err)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    application := *openapiclient.NewApplication() // Application | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.UpdateApplication(context.Background(), applicationId).Application(application).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.UpdateApplication`: %v\n", resp)
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
    openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
    applicationId := "applicationId_example" // string | 
    applicationVersionId := "applicationVersionId_example" // string | 
    applicationVersion := *openapiclient.NewApplicationVersion("VersionName_example") // ApplicationVersion | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.UpdateApplicationVersion(context.Background(), applicationId, applicationVersionId).ApplicationVersion(applicationVersion).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.UpdateApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationVersion`: ApplicationVersion
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.UpdateApplicationVersion`: %v\n", resp)
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

