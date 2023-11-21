# RbacManagerRoleMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | The name of the role the user has permissions for | [optional] 
**ForAccount** | Pointer to **string** | The account for which the user has the role permission | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRbacManagerRoleMembership

`func NewRbacManagerRoleMembership() *RbacManagerRoleMembership`

NewRbacManagerRoleMembership instantiates a new RbacManagerRoleMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacManagerRoleMembershipWithDefaults

`func NewRbacManagerRoleMembershipWithDefaults() *RbacManagerRoleMembership`

NewRbacManagerRoleMembershipWithDefaults instantiates a new RbacManagerRoleMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *RbacManagerRoleMembership) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RbacManagerRoleMembership) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RbacManagerRoleMembership) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *RbacManagerRoleMembership) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetForAccount

`func (o *RbacManagerRoleMembership) GetForAccount() string`

GetForAccount returns the ForAccount field if non-nil, zero value otherwise.

### GetForAccountOk

`func (o *RbacManagerRoleMembership) GetForAccountOk() (*string, bool)`

GetForAccountOk returns a tuple with the ForAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAccount

`func (o *RbacManagerRoleMembership) SetForAccount(v string)`

SetForAccount sets ForAccount field to given value.

### HasForAccount

`func (o *RbacManagerRoleMembership) HasForAccount() bool`

HasForAccount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RbacManagerRoleMembership) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RbacManagerRoleMembership) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RbacManagerRoleMembership) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RbacManagerRoleMembership) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


