# UserGroupRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForAccount** | **string** | The account for this role | 
**Roles** | [**[]UserGroupRoleRolesInner**](UserGroupRoleRolesInner.md) |  | 

## Methods

### NewUserGroupRole

`func NewUserGroupRole(forAccount string, roles []UserGroupRoleRolesInner, ) *UserGroupRole`

NewUserGroupRole instantiates a new UserGroupRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupRoleWithDefaults

`func NewUserGroupRoleWithDefaults() *UserGroupRole`

NewUserGroupRoleWithDefaults instantiates a new UserGroupRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForAccount

`func (o *UserGroupRole) GetForAccount() string`

GetForAccount returns the ForAccount field if non-nil, zero value otherwise.

### GetForAccountOk

`func (o *UserGroupRole) GetForAccountOk() (*string, bool)`

GetForAccountOk returns a tuple with the ForAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAccount

`func (o *UserGroupRole) SetForAccount(v string)`

SetForAccount sets ForAccount field to given value.


### GetRoles

`func (o *UserGroupRole) GetRoles() []UserGroupRoleRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserGroupRole) GetRolesOk() (*[]UserGroupRoleRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserGroupRole) SetRoles(v []UserGroupRoleRolesInner)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


