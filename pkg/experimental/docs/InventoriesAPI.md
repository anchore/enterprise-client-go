# \InventoriesAPI

All URIs are relative to */exp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKubernetesClusters**](InventoriesAPI.md#GetKubernetesClusters) | **Get** /kubernetes-clusters | Return a summary of Kubernetes clusters and namespaces for this account
[**GetKubernetesVulnerabilitiesSummary**](InventoriesAPI.md#GetKubernetesVulnerabilitiesSummary) | **Get** /kubernetes-vulnerabilities-summary | Return a summary of vulnerabilities found within the Kubernetes inventory for this account



## GetKubernetesClusters

> AnchoreSummaryKubernetesClusters GetKubernetesClusters(ctx).Execute()

Return a summary of Kubernetes clusters and namespaces for this account



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
	resp, r, err := apiClient.InventoriesAPI.GetKubernetesClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetKubernetesClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesClusters`: AnchoreSummaryKubernetesClusters
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClustersRequest struct via the builder pattern


### Return type

[**AnchoreSummaryKubernetesClusters**](AnchoreSummaryKubernetesClusters.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesVulnerabilitiesSummary

> AnchoreSummaryKubernetesVulnerabilities GetKubernetesVulnerabilitiesSummary(ctx).Context(context).Execute()

Return a summary of vulnerabilities found within the Kubernetes inventory for this account



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
	context := "context_example" // string | Matches on an exact cluster name or exact cluster/namespace combination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetKubernetesVulnerabilitiesSummary(context.Background()).Context(context).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetKubernetesVulnerabilitiesSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesVulnerabilitiesSummary`: AnchoreSummaryKubernetesVulnerabilities
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetKubernetesVulnerabilitiesSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesVulnerabilitiesSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string** | Matches on an exact cluster name or exact cluster/namespace combination | 

### Return type

[**AnchoreSummaryKubernetesVulnerabilities**](AnchoreSummaryKubernetesVulnerabilities.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

