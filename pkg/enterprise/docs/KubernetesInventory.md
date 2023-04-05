# KubernetesInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Namespaces** | Pointer to [**[]KubernetesInventoryNamespaces**](KubernetesInventoryNamespaces.md) |  | [optional] 
**Nodes** | Pointer to [**[]KubernetesInventoryNodes**](KubernetesInventoryNodes.md) |  | [optional] 
**Pods** | Pointer to [**[]KubernetesInventoryPods**](KubernetesInventoryPods.md) |  | [optional] 
**Containers** | Pointer to [**[]KubernetesInventoryContainers**](KubernetesInventoryContainers.md) |  | [optional] 

## Methods

### NewKubernetesInventory

`func NewKubernetesInventory(clusterName string, timestamp time.Time, ) *KubernetesInventory`

NewKubernetesInventory instantiates a new KubernetesInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInventoryWithDefaults

`func NewKubernetesInventoryWithDefaults() *KubernetesInventory`

NewKubernetesInventoryWithDefaults instantiates a new KubernetesInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *KubernetesInventory) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *KubernetesInventory) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *KubernetesInventory) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetTimestamp

`func (o *KubernetesInventory) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *KubernetesInventory) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *KubernetesInventory) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetNamespaces

`func (o *KubernetesInventory) GetNamespaces() []KubernetesInventoryNamespaces`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *KubernetesInventory) GetNamespacesOk() (*[]KubernetesInventoryNamespaces, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *KubernetesInventory) SetNamespaces(v []KubernetesInventoryNamespaces)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *KubernetesInventory) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetNodes

`func (o *KubernetesInventory) GetNodes() []KubernetesInventoryNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *KubernetesInventory) GetNodesOk() (*[]KubernetesInventoryNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *KubernetesInventory) SetNodes(v []KubernetesInventoryNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *KubernetesInventory) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetPods

`func (o *KubernetesInventory) GetPods() []KubernetesInventoryPods`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *KubernetesInventory) GetPodsOk() (*[]KubernetesInventoryPods, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *KubernetesInventory) SetPods(v []KubernetesInventoryPods)`

SetPods sets Pods field to given value.

### HasPods

`func (o *KubernetesInventory) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetContainers

`func (o *KubernetesInventory) GetContainers() []KubernetesInventoryContainers`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *KubernetesInventory) GetContainersOk() (*[]KubernetesInventoryContainers, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *KubernetesInventory) SetContainers(v []KubernetesInventoryContainers)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *KubernetesInventory) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


