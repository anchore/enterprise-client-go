# UserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the user group | 
**Description** | Pointer to **string** | The description of the user group | [optional] 
**GroupUuid** | Pointer to **string** | The unique identifier for the user group | [optional] 

## Methods

### NewUserGroup

`func NewUserGroup(name string, ) *UserGroup`

NewUserGroup instantiates a new UserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupWithDefaults

`func NewUserGroupWithDefaults() *UserGroup`

NewUserGroupWithDefaults instantiates a new UserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupUuid

`func (o *UserGroup) GetGroupUuid() string`

GetGroupUuid returns the GroupUuid field if non-nil, zero value otherwise.

### GetGroupUuidOk

`func (o *UserGroup) GetGroupUuidOk() (*string, bool)`

GetGroupUuidOk returns a tuple with the GroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUuid

`func (o *UserGroup) SetGroupUuid(v string)`

SetGroupUuid sets GroupUuid field to given value.

### HasGroupUuid

`func (o *UserGroup) HasGroupUuid() bool`

HasGroupUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


