# UserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the user group | 
**Description** | Pointer to **string** | The description of the user group | [optional] 
**GroupUuid** | **string** | The unique identifier for the user group | 
**CreatedAt** | Pointer to **time.Time** | The timestamp of when the user group was created | [optional] 
**LastUpdated** | Pointer to **time.Time** | The timestamp of the last update to this user group | [optional] 
**AccountRoles** | Pointer to [**UserGroupRoles**](UserGroupRoles.md) |  | [optional] 

## Methods

### NewUserGroup

`func NewUserGroup(name string, groupUuid string, ) *UserGroup`

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


### GetCreatedAt

`func (o *UserGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UserGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UserGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UserGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UserGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccountRoles

`func (o *UserGroup) GetAccountRoles() UserGroupRoles`

GetAccountRoles returns the AccountRoles field if non-nil, zero value otherwise.

### GetAccountRolesOk

`func (o *UserGroup) GetAccountRolesOk() (*UserGroupRoles, bool)`

GetAccountRolesOk returns a tuple with the AccountRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRoles

`func (o *UserGroup) SetAccountRoles(v UserGroupRoles)`

SetAccountRoles sets AccountRoles field to given value.

### HasAccountRoles

`func (o *UserGroup) HasAccountRoles() bool`

HasAccountRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


