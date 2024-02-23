# UserGroupPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the user group | 
**Description** | Pointer to **string** | The description of the user group | [optional] 

## Methods

### NewUserGroupPost

`func NewUserGroupPost(name string, ) *UserGroupPost`

NewUserGroupPost instantiates a new UserGroupPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupPostWithDefaults

`func NewUserGroupPostWithDefaults() *UserGroupPost`

NewUserGroupPostWithDefaults instantiates a new UserGroupPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserGroupPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroupPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroupPost) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UserGroupPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserGroupPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserGroupPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserGroupPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


