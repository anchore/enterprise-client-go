# ECSInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Tasks** | Pointer to [**[]ECSInventoryTasks**](ECSInventoryTasks.md) |  | [optional] 
**Containers** | Pointer to [**[]ECSInventoryContainers**](ECSInventoryContainers.md) |  | [optional] 

## Methods

### NewECSInventory

`func NewECSInventory(clusterName string, timestamp time.Time, ) *ECSInventory`

NewECSInventory instantiates a new ECSInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSInventoryWithDefaults

`func NewECSInventoryWithDefaults() *ECSInventory`

NewECSInventoryWithDefaults instantiates a new ECSInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ECSInventory) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ECSInventory) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ECSInventory) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetTimestamp

`func (o *ECSInventory) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ECSInventory) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ECSInventory) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTasks

`func (o *ECSInventory) GetTasks() []ECSInventoryTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ECSInventory) GetTasksOk() (*[]ECSInventoryTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ECSInventory) SetTasks(v []ECSInventoryTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ECSInventory) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *ECSInventory) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *ECSInventory) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil
### GetContainers

`func (o *ECSInventory) GetContainers() []ECSInventoryContainers`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ECSInventory) GetContainersOk() (*[]ECSInventoryContainers, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ECSInventory) SetContainers(v []ECSInventoryContainers)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *ECSInventory) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


