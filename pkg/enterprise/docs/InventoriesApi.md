# \InventoriesApi

All URIs are relative to *http://localhost/enterprise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInventory**](InventoriesApi.md#DeleteInventory) | **Delete** /inventories | Delete runtime inventory by type and context
[**GetImageInventory**](InventoriesApi.md#GetImageInventory) | **Get** /inventories | Return a list of the images in inventories for this account
[**PostEcsInventory**](InventoriesApi.md#PostEcsInventory) | **Post** /ecs-inventory | Add container metadata from Amazon ECS
[**PostKubernetesInventory**](InventoriesApi.md#PostKubernetesInventory) | **Post** /kubernetes-inventory | Add container metadata from a Kubernetes deployment
[**SyncImageInventory**](InventoriesApi.md#SyncImageInventory) | **Post** /inventories | synchronizes the list of the images in a given cluster for the inventory



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
    inventory := *openapiclient.NewECSInventory("ClusterName_example", time.Now()) // ECSInventory | 

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


## SyncImageInventory

> []InventoryItem SyncImageInventory(ctx).Inventory(inventory).XAnchoreAccount(xAnchoreAccount).Execute()

synchronizes the list of the images in a given cluster for the inventory



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
    inventory := *openapiclient.NewInventoryReport() // InventoryReport | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.SyncImageInventory(context.Background()).Inventory(inventory).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.SyncImageInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncImageInventory`: []InventoryItem
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.SyncImageInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncImageInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventory** | [**InventoryReport**](InventoryReport.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]InventoryItem**](InventoryItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

