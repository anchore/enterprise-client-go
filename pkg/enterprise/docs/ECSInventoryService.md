# ECSInventoryService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** |  | 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewECSInventoryService

`func NewECSInventoryService(arn string, ) *ECSInventoryService`

NewECSInventoryService instantiates a new ECSInventoryService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSInventoryServiceWithDefaults

`func NewECSInventoryServiceWithDefaults() *ECSInventoryService`

NewECSInventoryServiceWithDefaults instantiates a new ECSInventoryService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ECSInventoryService) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ECSInventoryService) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ECSInventoryService) SetArn(v string)`

SetArn sets Arn field to given value.


### GetTags

`func (o *ECSInventoryService) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ECSInventoryService) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ECSInventoryService) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ECSInventoryService) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


