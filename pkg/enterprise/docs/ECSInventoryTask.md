# ECSInventoryTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** |  | 
**ServiceArn** | Pointer to **string** |  | [optional] 
**TaskDefinitionArn** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewECSInventoryTask

`func NewECSInventoryTask(arn string, ) *ECSInventoryTask`

NewECSInventoryTask instantiates a new ECSInventoryTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSInventoryTaskWithDefaults

`func NewECSInventoryTaskWithDefaults() *ECSInventoryTask`

NewECSInventoryTaskWithDefaults instantiates a new ECSInventoryTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ECSInventoryTask) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ECSInventoryTask) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ECSInventoryTask) SetArn(v string)`

SetArn sets Arn field to given value.


### GetServiceArn

`func (o *ECSInventoryTask) GetServiceArn() string`

GetServiceArn returns the ServiceArn field if non-nil, zero value otherwise.

### GetServiceArnOk

`func (o *ECSInventoryTask) GetServiceArnOk() (*string, bool)`

GetServiceArnOk returns a tuple with the ServiceArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceArn

`func (o *ECSInventoryTask) SetServiceArn(v string)`

SetServiceArn sets ServiceArn field to given value.

### HasServiceArn

`func (o *ECSInventoryTask) HasServiceArn() bool`

HasServiceArn returns a boolean if a field has been set.

### GetTaskDefinitionArn

`func (o *ECSInventoryTask) GetTaskDefinitionArn() string`

GetTaskDefinitionArn returns the TaskDefinitionArn field if non-nil, zero value otherwise.

### GetTaskDefinitionArnOk

`func (o *ECSInventoryTask) GetTaskDefinitionArnOk() (*string, bool)`

GetTaskDefinitionArnOk returns a tuple with the TaskDefinitionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionArn

`func (o *ECSInventoryTask) SetTaskDefinitionArn(v string)`

SetTaskDefinitionArn sets TaskDefinitionArn field to given value.

### HasTaskDefinitionArn

`func (o *ECSInventoryTask) HasTaskDefinitionArn() bool`

HasTaskDefinitionArn returns a boolean if a field has been set.

### GetTags

`func (o *ECSInventoryTask) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ECSInventoryTask) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ECSInventoryTask) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ECSInventoryTask) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


