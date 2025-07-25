# \ArtifactLifecycleAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArtifactLifecyclePolicy**](ArtifactLifecycleAPI.md#CreateArtifactLifecyclePolicy) | **Post** /system/artifact-lifecycle-policies | Create new artifact lifecycle policy
[**DeleteArtifactLifecyclePolicy**](ArtifactLifecycleAPI.md#DeleteArtifactLifecyclePolicy) | **Delete** /system/artifact-lifecycle-policies/{policy_uuid} | Delete lifecycle policy
[**GetArtifactLifecyclePolicy**](ArtifactLifecycleAPI.md#GetArtifactLifecyclePolicy) | **Get** /system/artifact-lifecycle-policies/{policy_uuid} | Get single artifact lifecycle policy
[**GetArtifactLifecyclePolicyByVersion**](ArtifactLifecycleAPI.md#GetArtifactLifecyclePolicyByVersion) | **Get** /system/artifact-lifecycle-policies/{policy_uuid}/versions | Get single artifact lifecycle policy by its version
[**ListArtifactLifecyclePolicies**](ArtifactLifecycleAPI.md#ListArtifactLifecyclePolicies) | **Get** /system/artifact-lifecycle-policies | List all artifact lifecycle policies
[**UpdateArtifactLifecyclePolicy**](ArtifactLifecycleAPI.md#UpdateArtifactLifecyclePolicy) | **Put** /system/artifact-lifecycle-policies/{policy_uuid} | Update a single artifact lifecycle policy



## CreateArtifactLifecyclePolicy

> ArtifactLifecyclePolicyResponse CreateArtifactLifecyclePolicy(ctx).ArtifactLifecyclePolicy(artifactLifecyclePolicy).Execute()

Create new artifact lifecycle policy

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
	artifactLifecyclePolicy := *openapiclient.NewArtifactLifecyclePolicy("Action_example", "Name_example", *openapiclient.NewArtifactLifecyclePolicyConditions(false, int32(123), "ArtifactType_example")) // ArtifactLifecyclePolicy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactLifecycleAPI.CreateArtifactLifecyclePolicy(context.Background()).ArtifactLifecyclePolicy(artifactLifecyclePolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecycleAPI.CreateArtifactLifecyclePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateArtifactLifecyclePolicy`: ArtifactLifecyclePolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ArtifactLifecycleAPI.CreateArtifactLifecyclePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtifactLifecyclePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifactLifecyclePolicy** | [**ArtifactLifecyclePolicy**](ArtifactLifecyclePolicy.md) |  | 

### Return type

[**ArtifactLifecyclePolicyResponse**](ArtifactLifecyclePolicyResponse.md)

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
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	policyUuid := "policyUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArtifactLifecycleAPI.DeleteArtifactLifecyclePolicy(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecycleAPI.DeleteArtifactLifecyclePolicy``: %v\n", err)
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactLifecyclePolicy

> ArtifactLifecyclePolicyResponse GetArtifactLifecyclePolicy(ctx, policyUuid).Execute()

Get single artifact lifecycle policy

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
	policyUuid := "policyUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactLifecycleAPI.GetArtifactLifecyclePolicy(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecycleAPI.GetArtifactLifecyclePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactLifecyclePolicy`: ArtifactLifecyclePolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ArtifactLifecycleAPI.GetArtifactLifecyclePolicy`: %v\n", resp)
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

[**ArtifactLifecyclePolicyResponse**](ArtifactLifecyclePolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactLifecyclePolicyByVersion

> ArtifactLifecyclePolicyResponse GetArtifactLifecyclePolicyByVersion(ctx, policyUuid).Version(version).Latest(latest).Execute()

Get single artifact lifecycle policy by its version

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
	policyUuid := "policyUuid_example" // string | 
	version := int32(56) // int32 | Request a specific version number (optional)
	latest := true // bool | Request the latest version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactLifecycleAPI.GetArtifactLifecyclePolicyByVersion(context.Background(), policyUuid).Version(version).Latest(latest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecycleAPI.GetArtifactLifecyclePolicyByVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactLifecyclePolicyByVersion`: ArtifactLifecyclePolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ArtifactLifecycleAPI.GetArtifactLifecyclePolicyByVersion`: %v\n", resp)
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

[**ArtifactLifecyclePolicyResponse**](ArtifactLifecyclePolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifactLifecyclePolicies

> ArtifactLifecyclePolicyList ListArtifactLifecyclePolicies(ctx).Execute()

List all artifact lifecycle policies

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
	resp, r, err := apiClient.ArtifactLifecycleAPI.ListArtifactLifecyclePolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecycleAPI.ListArtifactLifecyclePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListArtifactLifecyclePolicies`: ArtifactLifecyclePolicyList
	fmt.Fprintf(os.Stdout, "Response from `ArtifactLifecycleAPI.ListArtifactLifecyclePolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactLifecyclePoliciesRequest struct via the builder pattern


### Return type

[**ArtifactLifecyclePolicyList**](ArtifactLifecyclePolicyList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactLifecyclePolicy

> ArtifactLifecyclePolicyResponse UpdateArtifactLifecyclePolicy(ctx, policyUuid).ArtifactLifecyclePolicy(artifactLifecyclePolicy).Execute()

Update a single artifact lifecycle policy

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
	policyUuid := "policyUuid_example" // string | 
	artifactLifecyclePolicy := *openapiclient.NewArtifactLifecyclePolicy("Action_example", "Name_example", *openapiclient.NewArtifactLifecyclePolicyConditions(false, int32(123), "ArtifactType_example")) // ArtifactLifecyclePolicy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactLifecycleAPI.UpdateArtifactLifecyclePolicy(context.Background(), policyUuid).ArtifactLifecyclePolicy(artifactLifecyclePolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactLifecycleAPI.UpdateArtifactLifecyclePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateArtifactLifecyclePolicy`: ArtifactLifecyclePolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ArtifactLifecycleAPI.UpdateArtifactLifecyclePolicy`: %v\n", resp)
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

 **artifactLifecyclePolicy** | [**ArtifactLifecyclePolicy**](ArtifactLifecyclePolicy.md) |  | 

### Return type

[**ArtifactLifecyclePolicyResponse**](ArtifactLifecyclePolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

