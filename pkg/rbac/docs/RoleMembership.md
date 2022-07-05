# RoleMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | The name of the role the user has permissions for | [optional] 
**ForAccount** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRoleMembership

`func NewRoleMembership() *RoleMembership`

NewRoleMembership instantiates a new RoleMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMembershipWithDefaults

`func NewRoleMembershipWithDefaults() *RoleMembership`

NewRoleMembershipWithDefaults instantiates a new RoleMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *RoleMembership) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleMembership) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleMembership) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleMembership) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetForAccount

`func (o *RoleMembership) GetForAccount() []string`

GetForAccount returns the ForAccount field if non-nil, zero value otherwise.

### GetForAccountOk

`func (o *RoleMembership) GetForAccountOk() (*[]string, bool)`

GetForAccountOk returns a tuple with the ForAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAccount

`func (o *RoleMembership) SetForAccount(v []string)`

SetForAccount sets ForAccount field to given value.

### HasForAccount

`func (o *RoleMembership) HasForAccount() bool`

HasForAccount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RoleMembership) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoleMembership) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoleMembership) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RoleMembership) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


