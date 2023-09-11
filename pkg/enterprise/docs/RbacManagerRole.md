# RbacManagerRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the role | 
**Description** | Pointer to **string** | A role description for humans | [optional] 
**Permissions** | Pointer to [**[]RbacManagerPermission**](RbacManagerPermission.md) |  | [optional] 
**Immutable** | Pointer to **bool** | Are the permissions of this role modifiable by users (including admin users) | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the role was created | [optional] 
**LastUpdated** | Pointer to **time.Time** | The timestamp of the last update to the role metadata itself | [optional] 

## Methods

### NewRbacManagerRole

`func NewRbacManagerRole(name string, ) *RbacManagerRole`

NewRbacManagerRole instantiates a new RbacManagerRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacManagerRoleWithDefaults

`func NewRbacManagerRoleWithDefaults() *RbacManagerRole`

NewRbacManagerRoleWithDefaults instantiates a new RbacManagerRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RbacManagerRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RbacManagerRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RbacManagerRole) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RbacManagerRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RbacManagerRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RbacManagerRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RbacManagerRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *RbacManagerRole) GetPermissions() []RbacManagerPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RbacManagerRole) GetPermissionsOk() (*[]RbacManagerPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RbacManagerRole) SetPermissions(v []RbacManagerPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RbacManagerRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetImmutable

`func (o *RbacManagerRole) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *RbacManagerRole) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *RbacManagerRole) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *RbacManagerRole) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RbacManagerRole) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RbacManagerRole) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RbacManagerRole) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RbacManagerRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RbacManagerRole) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RbacManagerRole) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RbacManagerRole) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *RbacManagerRole) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


