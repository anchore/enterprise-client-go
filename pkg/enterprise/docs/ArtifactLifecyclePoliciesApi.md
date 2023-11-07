# \ArtifactLifecyclePoliciesApi

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArtifactLifecyclePolicy**](ArtifactLifecyclePoliciesApi.md#CreateArtifactLifecyclePolicy) | **Post** /system/artifact-lifecycle-policies | Create new artifact lifecycle policy
[**DeleteArtifactLifecyclePolicy**](ArtifactLifecyclePoliciesApi.md#DeleteArtifactLifecyclePolicy) | **Delete** /system/artifact-lifecycle-policies/{policy_uuid} | Delete lifecycle policy
[**GetArtifactLifecyclePolicy**](ArtifactLifecyclePoliciesApi.md#GetArtifactLifecyclePolicy) | **Get** /system/artifact-lifecycle-policies/{policy_uuid} | Get single artifact lifecycle policy
[**GetArtifactLifecyclePolicyByVersion**](ArtifactLifecyclePoliciesApi.md#GetArtifactLifecyclePolicyByVersion) | **Get** /system/artifact-lifecycle-policies/{policy_uuid}/versions | Get single artifact lifecycle policy by its version
[**ListArtifactLifecyclePolicies**](ArtifactLifecyclePoliciesApi.md#ListArtifactLifecyclePolicies) | **Get** /system/artifact-lifecycle-policies | List all artifact lifecycle policies
[**UpdateArtifactLifecyclePolicy**](ArtifactLifecyclePoliciesApi.md#UpdateArtifactLifecyclePolicy) | **Put** /system/artifact-lifecycle-policies/{policy_uuid} | Update a single artifact lifecycle policy



## CreateArtifactLifecyclePolicy

> ArtifactLifeCyclePolicy CreateArtifactLifecyclePolicy(ctx).ArtifactLifeCyclePolicy(artifactLifeCyclePolicy).Execute()

Create new artifact lifecycle policy

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
    artifactLifeCyclePolicy := *openapiclient.NewArtifactLifeCyclePolicy("Action_example", "Name_example", *openapiclient.NewArtifactLifeCyclePolicyConditions(int32(123), int32(123), "ArtifactType_example")) // ArtifactLifeCyclePolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactLifecyclePoliciesApi.CreateArtifactLifecyclePolicy(context.Background()).ArtifactLifeCyclePolicy(artifactLifeCyclePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecyclePoliciesApi.CreateArtifactLifecyclePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArtifactLifecyclePolicy`: ArtifactLifeCyclePolicy
    fmt.Fprintf(os.Stdout, "Response from `ArtifactLifecyclePoliciesApi.CreateArtifactLifecyclePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtifactLifecyclePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifactLifeCyclePolicy** | [**ArtifactLifeCyclePolicy**](ArtifactLifeCyclePolicy.md) |  | 

### Return type

[**ArtifactLifeCyclePolicy**](ArtifactLifeCyclePolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifactLifecyclePolicy

> DeleteArtifactLifecyclePolicy(ctx, policyUuid).Execute()

Delete lifecycle policy

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
    policyUuid := "policyUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactLifecyclePoliciesApi.DeleteArtifactLifecyclePolicy(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecyclePoliciesApi.DeleteArtifactLifecyclePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactLifecyclePolicyRequest struct via the builder pattern


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


## GetArtifactLifecyclePolicy

> ArtifactLifeCyclePolicy GetArtifactLifecyclePolicy(ctx, policyUuid).Execute()

Get single artifact lifecycle policy

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
    policyUuid := "policyUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactLifecyclePoliciesApi.GetArtifactLifecyclePolicy(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecyclePoliciesApi.GetArtifactLifecyclePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactLifecyclePolicy`: ArtifactLifeCyclePolicy
    fmt.Fprintf(os.Stdout, "Response from `ArtifactLifecyclePoliciesApi.GetArtifactLifecyclePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactLifecyclePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArtifactLifeCyclePolicy**](ArtifactLifeCyclePolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactLifecyclePolicyByVersion

> ArtifactLifeCyclePolicy GetArtifactLifecyclePolicyByVersion(ctx, policyUuid).Version(version).Latest(latest).Execute()

Get single artifact lifecycle policy by its version

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
    policyUuid := "policyUuid_example" // string | 
    version := int32(56) // int32 | Request a specific version number (optional)
    latest := true // bool | Request the latest version (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactLifecyclePoliciesApi.GetArtifactLifecyclePolicyByVersion(context.Background(), policyUuid).Version(version).Latest(latest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecyclePoliciesApi.GetArtifactLifecyclePolicyByVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactLifecyclePolicyByVersion`: ArtifactLifeCyclePolicy
    fmt.Fprintf(os.Stdout, "Response from `ArtifactLifecyclePoliciesApi.GetArtifactLifecyclePolicyByVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactLifecyclePolicyByVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | Request a specific version number | 
 **latest** | **bool** | Request the latest version | 

### Return type

[**ArtifactLifeCyclePolicy**](ArtifactLifeCyclePolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifactLifecyclePolicies

> ArtifactLifeCyclePolicyList ListArtifactLifecyclePolicies(ctx).Execute()

List all artifact lifecycle policies

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
    resp, r, err := api_client.ArtifactLifecyclePoliciesApi.ListArtifactLifecyclePolicies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecyclePoliciesApi.ListArtifactLifecyclePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifactLifecyclePolicies`: ArtifactLifeCyclePolicyList
    fmt.Fprintf(os.Stdout, "Response from `ArtifactLifecyclePoliciesApi.ListArtifactLifecyclePolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactLifecyclePoliciesRequest struct via the builder pattern


### Return type

[**ArtifactLifeCyclePolicyList**](ArtifactLifeCyclePolicyList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactLifecyclePolicy

> ArtifactLifeCyclePolicy UpdateArtifactLifecyclePolicy(ctx, policyUuid).ArtifactLifeCyclePolicy(artifactLifeCyclePolicy).Execute()

Update a single artifact lifecycle policy

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
    policyUuid := "policyUuid_example" // string | 
    artifactLifeCyclePolicy := *openapiclient.NewArtifactLifeCyclePolicy("Action_example", "Name_example", *openapiclient.NewArtifactLifeCyclePolicyConditions(int32(123), int32(123), "ArtifactType_example")) // ArtifactLifeCyclePolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactLifecyclePoliciesApi.UpdateArtifactLifecyclePolicy(context.Background(), policyUuid).ArtifactLifeCyclePolicy(artifactLifeCyclePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecyclePoliciesApi.UpdateArtifactLifecyclePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateArtifactLifecyclePolicy`: ArtifactLifeCyclePolicy
    fmt.Fprintf(os.Stdout, "Response from `ArtifactLifecyclePoliciesApi.UpdateArtifactLifecyclePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactLifecyclePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artifactLifeCyclePolicy** | [**ArtifactLifeCyclePolicy**](ArtifactLifeCyclePolicy.md) |  | 

### Return type

[**ArtifactLifeCyclePolicy**](ArtifactLifeCyclePolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

