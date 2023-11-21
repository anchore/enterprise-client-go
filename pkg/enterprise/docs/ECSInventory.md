# ECSInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterArn** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Tasks** | Pointer to [**[]ECSInventoryTasksInner**](ECSInventoryTasksInner.md) |  | [optional] 
**Containers** | Pointer to [**[]ECSInventoryContainersInner**](ECSInventoryContainersInner.md) |  | [optional] 
**Services** | Pointer to [**[]ECSInventoryServicesInner**](ECSInventoryServicesInner.md) |  | [optional] 

## Methods

### NewECSInventory

`func NewECSInventory(clusterArn string, timestamp time.Time, ) *ECSInventory`

NewECSInventory instantiates a new ECSInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSInventoryWithDefaults

`func NewECSInventoryWithDefaults() *ECSInventory`

NewECSInventoryWithDefaults instantiates a new ECSInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterArn

`func (o *ECSInventory) GetClusterArn() string`

GetClusterArn returns the ClusterArn field if non-nil, zero value otherwise.

### GetClusterArnOk

`func (o *ECSInventory) GetClusterArnOk() (*string, bool)`

GetClusterArnOk returns a tuple with the ClusterArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArn

`func (o *ECSInventory) SetClusterArn(v string)`

SetClusterArn sets ClusterArn field to given value.


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

`func (o *ECSInventory) GetTasks() []ECSInventoryTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ECSInventory) GetTasksOk() (*[]ECSInventoryTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ECSInventory) SetTasks(v []ECSInventoryTasksInner)`

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

`func (o *ECSInventory) GetContainers() []ECSInventoryContainersInner`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ECSInventory) GetContainersOk() (*[]ECSInventoryContainersInner, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ECSInventory) SetContainers(v []ECSInventoryContainersInner)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *ECSInventory) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetServices

`func (o *ECSInventory) GetServices() []ECSInventoryServicesInner`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ECSInventory) GetServicesOk() (*[]ECSInventoryServicesInner, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ECSInventory) SetServices(v []ECSInventoryServicesInner)`

SetServices sets Services field to given value.

### HasServices

`func (o *ECSInventory) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *ECSInventory) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *ECSInventory) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


