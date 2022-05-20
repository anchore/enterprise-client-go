# \InventoriesApi

All URIs are relative to *http://localhost/enterprise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInventoryCluster**](InventoriesApi.md#AddInventoryCluster) | **Post** /inventories/clusters | Create a cluster inventory
[**DelInventoryClusterByName**](InventoriesApi.md#DelInventoryClusterByName) | **Delete** /inventories/clusters/{cluster_name} | Delete a configured inventory clusters by cluster_name
[**GetImageInventory**](InventoriesApi.md#GetImageInventory) | **Get** /inventories | Return a list of the images in inventories for this account
[**GetInventoryClusterByName**](InventoriesApi.md#GetInventoryClusterByName) | **Get** /inventories/clusters/{cluster_name} | Return a configured inventory cluster
[**ListInventoryClusters**](InventoriesApi.md#ListInventoryClusters) | **Get** /inventories/clusters | Return a list of the configured inventory clusters
[**SyncImageInventory**](InventoriesApi.md#SyncImageInventory) | **Post** /inventories | synchronizes the list of the images in a given cluster for the inventory



## AddInventoryCluster

> InventoryCluster AddInventoryCluster(ctx).Cluster(cluster).XAnchoreAccount(xAnchoreAccount).Execute()

Create a cluster inventory



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
    cluster := *openapiclient.NewInventoryCluster() // InventoryCluster | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.AddInventoryCluster(context.Background()).Cluster(cluster).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.AddInventoryCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInventoryCluster`: InventoryCluster
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.AddInventoryCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInventoryClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | [**InventoryCluster**](InventoryCluster.md) |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**InventoryCluster**](InventoryCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelInventoryClusterByName

> DelInventoryClusterByName(ctx, clusterName).XAnchoreAccount(xAnchoreAccount).Execute()

Delete a configured inventory clusters by cluster_name



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
    clusterName := "clusterName_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.DelInventoryClusterByName(context.Background(), clusterName).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.DelInventoryClusterByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelInventoryClusterByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetInventoryClusterByName

> InventoryCluster GetInventoryClusterByName(ctx, clusterName).XAnchoreAccount(xAnchoreAccount).Execute()

Return a configured inventory cluster



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
    clusterName := "clusterName_example" // string | 
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.GetInventoryClusterByName(context.Background(), clusterName).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.GetInventoryClusterByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInventoryClusterByName`: InventoryCluster
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.GetInventoryClusterByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoryClusterByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**InventoryCluster**](InventoryCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInventoryClusters

> []InventoryCluster ListInventoryClusters(ctx).InventoryType(inventoryType).XAnchoreAccount(xAnchoreAccount).Execute()

Return a list of the configured inventory clusters



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
    xAnchoreAccount := "xAnchoreAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoriesApi.ListInventoryClusters(context.Background()).InventoryType(inventoryType).XAnchoreAccount(xAnchoreAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoriesApi.ListInventoryClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInventoryClusters`: []InventoryCluster
    fmt.Fprintf(os.Stdout, "Response from `InventoriesApi.ListInventoryClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInventoryClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryType** | **string** |  | 
 **xAnchoreAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]InventoryCluster**](InventoryCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

