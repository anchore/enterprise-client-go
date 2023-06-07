# \InventoriesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInventory**](InventoriesApi.md#DeleteInventory) | **Delete** /inventories | Delete runtime inventory by type and context
[**DeleteKubernetesNamespaces**](InventoriesApi.md#DeleteKubernetesNamespaces) | **Delete** /kubernetes-namespaces | Delete Kubernetes namespaces for a given criteria
[**GetEcsContainers**](InventoriesApi.md#GetEcsContainers) | **Get** /ecs-containers | Return a list of ECS containers that have been inventoried for this account
[**GetEcsServices**](InventoriesApi.md#GetEcsServices) | **Get** /ecs-services | Return a list of ECS services that have been inventoried for this account
[**GetEcsTasks**](InventoriesApi.md#GetEcsTasks) | **Get** /ecs-tasks | Return a list of ECS tasks that have been inventoried for this account
[**GetImageInventory**](InventoriesApi.md#GetImageInventory) | **Get** /inventories | Return a list of the images in inventories for this account
[**GetKubernetesContainers**](InventoriesApi.md#GetKubernetesContainers) | **Get** /kubernetes-containers | Return a list of Kubernetes containers that have been inventoried for this account
[**GetKubernetesNamespace**](InventoriesApi.md#GetKubernetesNamespace) | **Get** /kubernetes-namespaces/{namespace_id} | Return a Kubernetes namespace that has been inventoried for this account
[**GetKubernetesNamespaces**](InventoriesApi.md#GetKubernetesNamespaces) | **Get** /kubernetes-namespaces | Return a list of Kubernetes namespaces that have been inventoried for this account
[**GetKubernetesNode**](InventoriesApi.md#GetKubernetesNode) | **Get** /kubernetes-nodes/{node_id} | Return a Kubernetes node that has been inventoried for this account
[**GetKubernetesNodes**](InventoriesApi.md#GetKubernetesNodes) | **Get** /kubernetes-nodes | Return a list of Kubernetes nodes that have been inventoried for this account
[**GetKubernetesPod**](InventoriesApi.md#GetKubernetesPod) | **Get** /kubernetes-pods/{pod_id} | Return a Kubernetes pod that has been inventoried for this account
[**GetKubernetesPods**](InventoriesApi.md#GetKubernetesPods) | **Get** /kubernetes-pods | Return a list of Kubernetes pods that have been inventoried for this account
[**PostEcsInventory**](InventoriesApi.md#PostEcsInventory) | **Post** /ecs-inventory | Add container metadata from Amazon ECS
[**PostKubernetesInventory**](InventoriesApi.md#PostKubernetesInventory) | **Post** /kubernetes-inventory | Add container metadata from a Kubernetes deployment



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
    openapiclient "./openapi"
)

func main() {
    inventoryType := "inventoryType_example" // string | 
    context := "context_example" // string | 
    imageDigest := "imageDigest_example" // string |  (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.DeleteInventory(context.Background()).InventoryType(inventoryType).Context(context).ImageDigest(imageDigest).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.DeleteInventory``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    clusterName := "clusterName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.DeleteKubernetesNamespaces(context.Background()).ClusterName(clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.DeleteKubernetesNamespaces``: %v\n", err)
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

> ECSContainers GetEcsContainers(ctx).Execute()

Return a list of ECS containers that have been inventoried for this account



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
    resp, r, err := api_client.InventoriesApi.GetEcsContainers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetEcsContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEcsContainers`: ECSContainers
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetEcsContainers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcsContainersRequest struct via the builder pattern


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

> ECSServices GetEcsServices(ctx).Execute()

Return a list of ECS services that have been inventoried for this account



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
    resp, r, err := api_client.InventoriesApi.GetEcsServices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetEcsServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEcsServices`: ECSServices
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetEcsServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcsServicesRequest struct via the builder pattern


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

> ECSTasks GetEcsTasks(ctx).Execute()

Return a list of ECS tasks that have been inventoried for this account



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
    resp, r, err := api_client.InventoriesApi.GetEcsTasks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetEcsTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEcsTasks`: ECSTasks
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetEcsTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcsTasksRequest struct via the builder pattern


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

> []InventoryItem GetImageInventory(ctx).InventoryType(inventoryType).ImageDigest(imageDigest).Context(context).XAnchoreAccount(xAnchoreAccount).Execute()

Return a list of the images in inventories for this account



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
    inventoryType := "inventoryType_example" // string |  (optional)
    imageDigest := "imageDigest_example" // string |  (optional)
    context := "context_example" // string |  (optional)
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.GetImageInventory(context.Background()).InventoryType(inventoryType).ImageDigest(imageDigest).Context(context).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetImageInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageInventory`: []InventoryItem
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetImageInventory`: %v\n", resp)
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

[**[]InventoryItem**](InventoryItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesContainers

> KubernetesContainers GetKubernetesContainers(ctx).Execute()

Return a list of Kubernetes containers that have been inventoried for this account



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
    resp, r, err := api_client.InventoriesApi.GetKubernetesContainers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetKubernetesContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesContainers`: KubernetesContainers
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetKubernetesContainers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesContainersRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    namespaceId := "namespaceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.GetKubernetesNamespace(context.Background(), namespaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetKubernetesNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesNamespace`: KubernetesNamespace
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetKubernetesNamespace`: %v\n", resp)
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
    resp, r, err := api_client.InventoriesApi.GetKubernetesNamespaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetKubernetesNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesNamespaces`: KubernetesNamespaces
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetKubernetesNamespaces`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    nodeId := "nodeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.GetKubernetesNode(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetKubernetesNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesNode`: KubernetesNode
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetKubernetesNode`: %v\n", resp)
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

> KubernetesNodes GetKubernetesNodes(ctx).Execute()

Return a list of Kubernetes nodes that have been inventoried for this account



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
    resp, r, err := api_client.InventoriesApi.GetKubernetesNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetKubernetesNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesNodes`: KubernetesNodes
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetKubernetesNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesNodesRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    podId := "podId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.GetKubernetesPod(context.Background(), podId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetKubernetesPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesPod`: KubernetesPod
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetKubernetesPod`: %v\n", resp)
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

> KubernetesPods GetKubernetesPods(ctx).Execute()

Return a list of Kubernetes pods that have been inventoried for this account



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
    resp, r, err := api_client.InventoriesApi.GetKubernetesPods(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetKubernetesPods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesPods`: KubernetesPods
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetKubernetesPods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesPodsRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    inventory := *openapiclient.NewECSInventory("ClusterArn_example", time.Now()) // ECSInventory | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.PostEcsInventory(context.Background()).Inventory(inventory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.PostEcsInventory``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    inventory := *openapiclient.NewKubernetesInventory("ClusterName_example", time.Now()) // KubernetesInventory | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.PostKubernetesInventory(context.Background()).Inventory(inventory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.PostKubernetesInventory``: %v\n", err)
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

