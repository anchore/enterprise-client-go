# ECSInventoryServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** |  | 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewECSInventoryServices

`func NewECSInventoryServices(arn string, ) *ECSInventoryServices`

NewECSInventoryServices instantiates a new ECSInventoryServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSInventoryServicesWithDefaults

`func NewECSInventoryServicesWithDefaults() *ECSInventoryServices`

NewECSInventoryServicesWithDefaults instantiates a new ECSInventoryServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ECSInventoryServices) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ECSInventoryServices) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ECSInventoryServices) SetArn(v string)`

SetArn sets Arn field to given value.


### GetTags

`func (o *ECSInventoryServices) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ECSInventoryServices) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ECSInventoryServices) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ECSInventoryServices) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


