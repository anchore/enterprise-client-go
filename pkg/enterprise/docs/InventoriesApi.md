# \InventoriesAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInventory**](InventoriesAPI.md#DeleteInventory) | **Delete** /inventories | Delete runtime inventory by type and context
[**DeleteKubernetesNamespaces**](InventoriesAPI.md#DeleteKubernetesNamespaces) | **Delete** /kubernetes-namespaces | Delete Kubernetes namespaces for a given criteria
[**GetEcsContainers**](InventoriesAPI.md#GetEcsContainers) | **Get** /ecs-containers | Return a list of ECS containers that have been inventoried for this account
[**GetEcsServices**](InventoriesAPI.md#GetEcsServices) | **Get** /ecs-services | Return a list of ECS services that have been inventoried for this account
[**GetEcsTasks**](InventoriesAPI.md#GetEcsTasks) | **Get** /ecs-tasks | Return a list of ECS tasks that have been inventoried for this account
[**GetImageInventory**](InventoriesAPI.md#GetImageInventory) | **Get** /inventories | Return a list of the images in inventories for this account
[**GetKubernetesContainers**](InventoriesAPI.md#GetKubernetesContainers) | **Get** /kubernetes-containers | Return a list of Kubernetes containers that have been inventoried for this account
[**GetKubernetesNamespace**](InventoriesAPI.md#GetKubernetesNamespace) | **Get** /kubernetes-namespaces/{namespace_id} | Return a Kubernetes namespace that has been inventoried for this account
[**GetKubernetesNamespaces**](InventoriesAPI.md#GetKubernetesNamespaces) | **Get** /kubernetes-namespaces | Return a list of Kubernetes namespaces that have been inventoried for this account
[**GetKubernetesNode**](InventoriesAPI.md#GetKubernetesNode) | **Get** /kubernetes-nodes/{node_id} | Return a Kubernetes node that has been inventoried for this account
[**GetKubernetesNodes**](InventoriesAPI.md#GetKubernetesNodes) | **Get** /kubernetes-nodes | Return a list of Kubernetes nodes that have been inventoried for this account
[**GetKubernetesPod**](InventoriesAPI.md#GetKubernetesPod) | **Get** /kubernetes-pods/{pod_id} | Return a Kubernetes pod that has been inventoried for this account
[**GetKubernetesPods**](InventoriesAPI.md#GetKubernetesPods) | **Get** /kubernetes-pods | Return a list of Kubernetes pods that have been inventoried for this account
[**PostEcsInventory**](InventoriesAPI.md#PostEcsInventory) | **Post** /ecs-inventory | Add container metadata from Amazon ECS
[**PostKubernetesInventory**](InventoriesAPI.md#PostKubernetesInventory) | **Post** /kubernetes-inventory | Add container metadata from a Kubernetes deployment



## DeleteInventory

> DeleteInventory(ctx).InventoryType(inventoryType).Context(context).ImageDigest(imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()

Delete runtime inventory by type and context



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
	inventoryType := "inventoryType_example" // string | 
	context := "context_example" // string | 
	imageDigest := "imageDigest_example" // string |  (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoriesAPI.DeleteInventory(context.Background()).InventoryType(inventoryType).Context(context).ImageDigest(imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.DeleteInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryType** | **string** |  | 
 **context** | **string** |  | 
 **imageDigest** | **string** |  | 
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


## DeleteKubernetesNamespaces

> DeleteKubernetesNamespaces(ctx).ClusterName(clusterName).Execute()

Delete Kubernetes namespaces for a given criteria



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
	clusterName := "clusterName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoriesAPI.DeleteKubernetesNamespaces(context.Background()).ClusterName(clusterName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.DeleteKubernetesNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKubernetesNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterName** | **string** |  | 

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


## GetEcsContainers

> ECSContainers GetEcsContainers(ctx).Page(page).PageSize(pageSize).Execute()

Return a list of ECS containers that have been inventoried for this account



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
	page := int32(56) // int32 | 
	pageSize := int32(56) // int32 |  (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetEcsContainers(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetEcsContainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcsContainers`: ECSContainers
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetEcsContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEcsContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | [default to 1000]

### Return type

[**ECSContainers**](ECSContainers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcsServices

> ECSServices GetEcsServices(ctx).Page(page).PageSize(pageSize).Execute()

Return a list of ECS services that have been inventoried for this account



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
	page := int32(56) // int32 | 
	pageSize := int32(56) // int32 |  (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetEcsServices(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetEcsServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcsServices`: ECSServices
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetEcsServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEcsServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | [default to 1000]

### Return type

[**ECSServices**](ECSServices.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcsTasks

> ECSTasks GetEcsTasks(ctx).Page(page).PageSize(pageSize).Execute()

Return a list of ECS tasks that have been inventoried for this account



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
	page := int32(56) // int32 | 
	pageSize := int32(56) // int32 |  (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetEcsTasks(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetEcsTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcsTasks`: ECSTasks
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetEcsTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEcsTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | [default to 1000]

### Return type

[**ECSTasks**](ECSTasks.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageInventory

> InventoryItems GetImageInventory(ctx).InventoryType(inventoryType).ImageDigest(imageDigest).Context(context).XAnchoreAccount(xAnchoreAccount).Execute()

Return a list of the images in inventories for this account



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
	inventoryType := "inventoryType_example" // string |  (optional)
	imageDigest := "imageDigest_example" // string |  (optional)
	context := "context_example" // string |  (optional)
	xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetImageInventory(context.Background()).InventoryType(inventoryType).ImageDigest(imageDigest).Context(context).XAnchoreAccount(xAnchoreAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetImageInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageInventory`: InventoryItems
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetImageInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryType** | **string** |  | 
 **imageDigest** | **string** |  | 
 **context** | **string** |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**InventoryItems**](InventoryItems.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesContainers

> KubernetesContainers GetKubernetesContainers(ctx).Page(page).PageSize(pageSize).Execute()

Return a list of Kubernetes containers that have been inventoried for this account



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
	page := int32(56) // int32 | 
	pageSize := int32(56) // int32 |  (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetKubernetesContainers(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetKubernetesContainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesContainers`: KubernetesContainers
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetKubernetesContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | [default to 1000]

### Return type

[**KubernetesContainers**](KubernetesContainers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesNamespace

> KubernetesNamespace GetKubernetesNamespace(ctx, namespaceId).Execute()

Return a Kubernetes namespace that has been inventoried for this account



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
	namespaceId := "namespaceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetKubernetesNamespace(context.Background(), namespaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetKubernetesNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesNamespace`: KubernetesNamespace
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetKubernetesNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesNamespace**](KubernetesNamespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesNamespaces

> KubernetesNamespaces GetKubernetesNamespaces(ctx).Page(page).PageSize(pageSize).Execute()

Return a list of Kubernetes namespaces that have been inventoried for this account



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
	page := int32(56) // int32 | 
	pageSize := int32(56) // int32 |  (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetKubernetesNamespaces(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetKubernetesNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesNamespaces`: KubernetesNamespaces
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetKubernetesNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | [default to 1000]

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


## GetKubernetesNode

> KubernetesNode GetKubernetesNode(ctx, nodeId).Execute()

Return a Kubernetes node that has been inventoried for this account



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
	nodeId := "nodeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetKubernetesNode(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetKubernetesNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesNode`: KubernetesNode
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetKubernetesNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesNode**](KubernetesNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesNodes

> KubernetesNodes GetKubernetesNodes(ctx).Page(page).PageSize(pageSize).Execute()

Return a list of Kubernetes nodes that have been inventoried for this account



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
	page := int32(56) // int32 | 
	pageSize := int32(56) // int32 |  (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetKubernetesNodes(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetKubernetesNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesNodes`: KubernetesNodes
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetKubernetesNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | [default to 1000]

### Return type

[**KubernetesNodes**](KubernetesNodes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesPod

> KubernetesPod GetKubernetesPod(ctx, podId).Execute()

Return a Kubernetes pod that has been inventoried for this account



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
	podId := "podId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetKubernetesPod(context.Background(), podId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetKubernetesPod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesPod`: KubernetesPod
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetKubernetesPod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesPod**](KubernetesPod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesPods

> KubernetesPods GetKubernetesPods(ctx).Page(page).PageSize(pageSize).Execute()

Return a list of Kubernetes pods that have been inventoried for this account



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
	page := int32(56) // int32 | 
	pageSize := int32(56) // int32 |  (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.GetKubernetesPods(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.GetKubernetesPods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesPods`: KubernetesPods
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.GetKubernetesPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | [default to 1000]

### Return type

[**KubernetesPods**](KubernetesPods.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEcsInventory

> PostEcsInventory(ctx).Inventory(inventory).Execute()

Add container metadata from Amazon ECS



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	inventory := *openapiclient.NewECSInventory("ClusterArn_example", time.Now()) // ECSInventory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoriesAPI.PostEcsInventory(context.Background()).Inventory(inventory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.PostEcsInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEcsInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventory** | [**ECSInventory**](ECSInventory.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKubernetesInventory

> PostKubernetesInventory(ctx).Inventory(inventory).Execute()

Add container metadata from a Kubernetes deployment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/anchore/enterprise-client-go"
)

func main() {
	inventory := *openapiclient.NewKubernetesInventory("ClusterName_example", time.Now()) // KubernetesInventory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoriesAPI.PostKubernetesInventory(context.Background()).Inventory(inventory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.PostKubernetesInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKubernetesInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventory** | [**KubernetesInventory**](KubernetesInventory.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

