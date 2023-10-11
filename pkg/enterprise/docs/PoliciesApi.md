# \PoliciesAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPolicy**](PoliciesAPI.md#AddPolicy) | **Post** /policies | Add a new policy
[**DeletePolicy**](PoliciesAPI.md#DeletePolicy) | **Delete** /policies/{policy_id} | Delete policy
[**GetPolicy**](PoliciesAPI.md#GetPolicy) | **Get** /policies/{policy_id} | Get specific policy
[**ListPolicies**](PoliciesAPI.md#ListPolicies) | **Get** /policies | List policies
[**UpdatePolicy**](PoliciesAPI.md#UpdatePolicy) | **Put** /policies/{policy_id} | Update policy



## AddPolicy

> PolicyRecord AddPolicy(ctx).Policy(policy).XAnchoreAccount(xAnchoreAccount).Execute()

Add a new policy



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
    policy := *openapiclient.NewPolicy("Id_example", "Name_example", "Version_example", []openapiclient.RuleSet{*openapiclient.NewRuleSet("Id_example", "Name_example", "Version_example", []openapiclient.PolicyRule{*openapiclient.NewPolicyRule("Id_example", "Gate_example", "Trigger_example", "Action_example", []openapiclient.PolicyRuleParamsInner{*openapiclient.NewPolicyRuleParamsInner("Name_example", "Value_example")})})}, []openapiclient.MappingRule{*openapiclient.NewMappingRule("Id_example", "Name_example", []string{"AllowlistIds_example"}, []string{"RuleSetIds_example"}, "Registry_example", "Repository_example", *openapiclient.NewImageRef("Type_example", "Value_example"))}) // Policy | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesAPI.AddPolicy(context.Background()).Policy(policy).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AddPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPolicy`: PolicyRecord
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.AddPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policy** | [**Policy**](Policy.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**PolicyRecord**](PolicyRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> DeletePolicy(ctx, policyId).XAnchoreAccount(xAnchoreAccount).Execute()

Delete policy



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
    policyId := "policyId_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PoliciesAPI.DeletePolicy(context.Background(), policyId).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.DeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

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


## GetPolicy

> PolicyRecord GetPolicy(ctx, policyId).Detail(detail).XAnchoreAccount(xAnchoreAccount).Execute()

Get specific policy



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
    policyId := "policyId_example" // string | 
    detail := true // bool | Include policy detail in the form of the full policy content for each entry (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesAPI.GetPolicy(context.Background(), policyId).Detail(detail).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicy`: PolicyRecord
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detail** | **bool** | Include policy detail in the form of the full policy content for each entry | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**PolicyRecord**](PolicyRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> []PolicyRecord ListPolicies(ctx).Detail(detail).XAnchoreAccount(xAnchoreAccount).Execute()

List policies



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
    detail := true // bool | Include policy detail in the form of the full policy content for each entry (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesAPI.ListPolicies(context.Background()).Detail(detail).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.ListPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicies`: []PolicyRecord
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.ListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detail** | **bool** | Include policy detail in the form of the full policy content for each entry | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]PolicyRecord**](PolicyRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> PolicyRecord UpdatePolicy(ctx, policyId).Policy(policy).Active(active).XAnchoreAccount(xAnchoreAccount).Execute()

Update policy



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
    policyId := "policyId_example" // string | 
    policy := *openapiclient.NewPolicyRecord("PolicyId_example", false, "AccountName_example", "PolicySource_example", "Name_example") // PolicyRecord | 
    active := true // bool | Mark policy as active (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesAPI.UpdatePolicy(context.Background(), policyId).Policy(policy).Active(active).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.UpdatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicy`: PolicyRecord
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.UpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policy** | [**PolicyRecord**](PolicyRecord.md) |  | 
 **active** | **bool** | Mark policy as active | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**PolicyRecord**](PolicyRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

