# \RepositoryCredentialsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRepository**](RepositoryCredentialsApi.md#AddRepository) | **Post** /repositories | Add repository to watch



## AddRepository

> []Subscription AddRepository(ctx).Repository(repository).Autosubscribe(autosubscribe).Dryrun(dryrun).XAnchoreAccount(xAnchoreAccount).Execute()

Add repository to watch

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
    repository := "repository_example" // string | full repository to add e.g. docker.io/library/alpine
    autosubscribe := true // bool | flag to enable/disable auto tag_update activation when new images from a repo are added (optional)
    dryrun := true // bool | flag to return tags in the repository without actually watching the repository, default is false (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoryCredentialsApi.AddRepository(context.Background()).Repository(repository).Autosubscribe(autosubscribe).Dryrun(dryrun).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryCredentialsApi.AddRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRepository`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `RepositoryCredentialsApi.AddRepository`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repository** | **string** | full repository to add e.g. docker.io/library/alpine | 
 **autosubscribe** | **bool** | flag to enable/disable auto tag_update activation when new images from a repo are added | 
 **dryrun** | **bool** | flag to return tags in the repository without actually watching the repository, default is false | 
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

