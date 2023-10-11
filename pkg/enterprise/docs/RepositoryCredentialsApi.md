# \RepositoryCredentialsAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRepository**](RepositoryCredentialsAPI.md#AddRepository) | **Post** /repositories | Add repository to watch



## AddRepository

> []Subscription AddRepository(ctx).Repository(repository).AutoSubscribe(autoSubscribe).DryRun(dryRun).XAnchoreAccount(xAnchoreAccount).Execute()

Add repository to watch

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
    repository := "repository_example" // string | full repository to add e.g. docker.io/library/alpine
    autoSubscribe := true // bool | flag to enable/disable auto tag_update activation when new images from a repo are added (optional)
    dryRun := true // bool | flag to return tags in the repository without actually watching the repository, default is false (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryCredentialsAPI.AddRepository(context.Background()).Repository(repository).AutoSubscribe(autoSubscribe).DryRun(dryRun).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryCredentialsAPI.AddRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRepository`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `RepositoryCredentialsAPI.AddRepository`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repository** | **string** | full repository to add e.g. docker.io/library/alpine | 
 **autoSubscribe** | **bool** | flag to enable/disable auto tag_update activation when new images from a repo are added | 
 **dryRun** | **bool** | flag to return tags in the repository without actually watching the repository, default is false | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

