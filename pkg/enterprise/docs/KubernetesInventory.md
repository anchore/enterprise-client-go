# KubernetesInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** |  | 
**Timestamp** | **time.Time** |  | 
<<<<<<< HEAD
**Namespaces** | Pointer to [**[]KubernetesInventoryNamespacesInner**](KubernetesInventoryNamespacesInner.md) |  | [optional] 
**Nodes** | Pointer to [**[]KubernetesInventoryNodesInner**](KubernetesInventoryNodesInner.md) |  | [optional] 
**Pods** | Pointer to [**[]KubernetesInventoryPodsInner**](KubernetesInventoryPodsInner.md) |  | [optional] 
**Containers** | Pointer to [**[]KubernetesInventoryContainersInner**](KubernetesInventoryContainersInner.md) |  | [optional] 
=======
**Namespaces** | Pointer to [**[]KubernetesInventoryNamespaces**](KubernetesInventoryNamespaces.md) |  | [optional] 
**Nodes** | Pointer to [**[]KubernetesInventoryNodes**](KubernetesInventoryNodes.md) |  | [optional] 
**Pods** | Pointer to [**[]KubernetesInventoryPods**](KubernetesInventoryPods.md) |  | [optional] 
**Containers** | Pointer to [**[]KubernetesInventoryContainers**](KubernetesInventoryContainers.md) |  | [optional] 
>>>>>>> main

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

<<<<<<< HEAD
`func (o *KubernetesInventory) GetNamespaces() []KubernetesInventoryNamespacesInner`
=======
`func (o *KubernetesInventory) GetNamespaces() []KubernetesInventoryNamespaces`
>>>>>>> main

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

<<<<<<< HEAD
`func (o *KubernetesInventory) GetNamespacesOk() (*[]KubernetesInventoryNamespacesInner, bool)`
=======
`func (o *KubernetesInventory) GetNamespacesOk() (*[]KubernetesInventoryNamespaces, bool)`
>>>>>>> main

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

<<<<<<< HEAD
`func (o *KubernetesInventory) SetNamespaces(v []KubernetesInventoryNamespacesInner)`
=======
`func (o *KubernetesInventory) SetNamespaces(v []KubernetesInventoryNamespaces)`
>>>>>>> main

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *KubernetesInventory) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetNodes

<<<<<<< HEAD
`func (o *KubernetesInventory) GetNodes() []KubernetesInventoryNodesInner`
=======
`func (o *KubernetesInventory) GetNodes() []KubernetesInventoryNodes`
>>>>>>> main

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

<<<<<<< HEAD
`func (o *KubernetesInventory) GetNodesOk() (*[]KubernetesInventoryNodesInner, bool)`
=======
`func (o *KubernetesInventory) GetNodesOk() (*[]KubernetesInventoryNodes, bool)`
>>>>>>> main

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

<<<<<<< HEAD
`func (o *KubernetesInventory) SetNodes(v []KubernetesInventoryNodesInner)`
=======
`func (o *KubernetesInventory) SetNodes(v []KubernetesInventoryNodes)`
>>>>>>> main

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *KubernetesInventory) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetPods

<<<<<<< HEAD
`func (o *KubernetesInventory) GetPods() []KubernetesInventoryPodsInner`
=======
`func (o *KubernetesInventory) GetPods() []KubernetesInventoryPods`
>>>>>>> main

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

<<<<<<< HEAD
`func (o *KubernetesInventory) GetPodsOk() (*[]KubernetesInventoryPodsInner, bool)`
=======
`func (o *KubernetesInventory) GetPodsOk() (*[]KubernetesInventoryPods, bool)`
>>>>>>> main

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

<<<<<<< HEAD
`func (o *KubernetesInventory) SetPods(v []KubernetesInventoryPodsInner)`
=======
`func (o *KubernetesInventory) SetPods(v []KubernetesInventoryPods)`
>>>>>>> main

SetPods sets Pods field to given value.

### HasPods

`func (o *KubernetesInventory) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetContainers

<<<<<<< HEAD
`func (o *KubernetesInventory) GetContainers() []KubernetesInventoryContainersInner`
=======
`func (o *KubernetesInventory) GetContainers() []KubernetesInventoryContainers`
>>>>>>> main

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

<<<<<<< HEAD
`func (o *KubernetesInventory) GetContainersOk() (*[]KubernetesInventoryContainersInner, bool)`
=======
`func (o *KubernetesInventory) GetContainersOk() (*[]KubernetesInventoryContainers, bool)`
>>>>>>> main

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

<<<<<<< HEAD
`func (o *KubernetesInventory) SetContainers(v []KubernetesInventoryContainersInner)`
=======
`func (o *KubernetesInventory) SetContainers(v []KubernetesInventoryContainers)`
>>>>>>> main

SetContainers sets Containers field to given value.

### HasContainers

`func (o *KubernetesInventory) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


