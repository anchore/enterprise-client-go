# UserGroupRoleRolesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | The role name | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp of when the role membership was created | [optional] 
**MembershipId** | Pointer to **string** | The unique identifier for the role membership | [optional] 

## Methods

### NewUserGroupRoleRolesInner

`func NewUserGroupRoleRolesInner() *UserGroupRoleRolesInner`

NewUserGroupRoleRolesInner instantiates a new UserGroupRoleRolesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupRoleRolesInnerWithDefaults

`func NewUserGroupRoleRolesInnerWithDefaults() *UserGroupRoleRolesInner`

NewUserGroupRoleRolesInnerWithDefaults instantiates a new UserGroupRoleRolesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *UserGroupRoleRolesInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserGroupRoleRolesInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserGroupRoleRolesInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserGroupRoleRolesInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserGroupRoleRolesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserGroupRoleRolesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserGroupRoleRolesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserGroupRoleRolesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMembershipId

`func (o *UserGroupRoleRolesInner) GetMembershipId() string`

GetMembershipId returns the MembershipId field if non-nil, zero value otherwise.

### GetMembershipIdOk

`func (o *UserGroupRoleRolesInner) GetMembershipIdOk() (*string, bool)`

GetMembershipIdOk returns a tuple with the MembershipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipId

`func (o *UserGroupRoleRolesInner) SetMembershipId(v string)`

SetMembershipId sets MembershipId field to given value.

### HasMembershipId

`func (o *UserGroupRoleRolesInner) HasMembershipId() bool`

HasMembershipId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


