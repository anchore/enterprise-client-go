# \DefaultApi

All URIs are relative to *http://localhost/enterprise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKubernetesNamespaces**](DefaultApi.md#GetKubernetesNamespaces) | **Get** /kubernetes-namespaces | Return a list of Kubernetes namespaces that have been inventoried for this account



## GetKubernetesNamespaces

> KubernetesNamespaces GetKubernetesNamespaces(ctx).Execute()

Return a list of Kubernetes namespaces that have been inventoried for this account



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
    resp, r, err := api_client.DefaultApi.GetKubernetesNamespaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetKubernetesNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesNamespaces`: KubernetesNamespaces
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetKubernetesNamespaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesNamespacesRequest struct via the builder pattern


### Return type

[**KubernetesNamespaces**](KubernetesNamespaces.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

