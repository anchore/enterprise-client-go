# ECSInventoryTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** |  | 
**ServiceArn** | Pointer to **string** |  | [optional] 
**TaskDefinitionArn** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewECSInventoryTasks

`func NewECSInventoryTasks(arn string, ) *ECSInventoryTasks`

NewECSInventoryTasks instantiates a new ECSInventoryTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSInventoryTasksWithDefaults

`func NewECSInventoryTasksWithDefaults() *ECSInventoryTasks`

NewECSInventoryTasksWithDefaults instantiates a new ECSInventoryTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ECSInventoryTasks) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ECSInventoryTasks) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ECSInventoryTasks) SetArn(v string)`

SetArn sets Arn field to given value.


### GetServiceArn

`func (o *ECSInventoryTasks) GetServiceArn() string`

GetServiceArn returns the ServiceArn field if non-nil, zero value otherwise.

### GetServiceArnOk

`func (o *ECSInventoryTasks) GetServiceArnOk() (*string, bool)`

GetServiceArnOk returns a tuple with the ServiceArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceArn

`func (o *ECSInventoryTasks) SetServiceArn(v string)`

SetServiceArn sets ServiceArn field to given value.

### HasServiceArn

`func (o *ECSInventoryTasks) HasServiceArn() bool`

HasServiceArn returns a boolean if a field has been set.

### GetTaskDefinitionArn

`func (o *ECSInventoryTasks) GetTaskDefinitionArn() string`

GetTaskDefinitionArn returns the TaskDefinitionArn field if non-nil, zero value otherwise.

### GetTaskDefinitionArnOk

`func (o *ECSInventoryTasks) GetTaskDefinitionArnOk() (*string, bool)`

GetTaskDefinitionArnOk returns a tuple with the TaskDefinitionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionArn

`func (o *ECSInventoryTasks) SetTaskDefinitionArn(v string)`

SetTaskDefinitionArn sets TaskDefinitionArn field to given value.

### HasTaskDefinitionArn

`func (o *ECSInventoryTasks) HasTaskDefinitionArn() bool`

HasTaskDefinitionArn returns a boolean if a field has been set.

### GetTags

`func (o *ECSInventoryTasks) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ECSInventoryTasks) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ECSInventoryTasks) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ECSInventoryTasks) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


