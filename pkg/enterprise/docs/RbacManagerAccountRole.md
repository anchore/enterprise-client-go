# RbacManagerAccountRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForAccount** | Pointer to **string** | The account scope that applies to the set of roles | [optional] 
**Roles** | Pointer to [**RbacManagerRole**](RbacManagerRole.md) |  | [optional] 
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 

## Methods

### NewRbacManagerAccountRole

`func NewRbacManagerAccountRole() *RbacManagerAccountRole`

NewRbacManagerAccountRole instantiates a new RbacManagerAccountRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacManagerAccountRoleWithDefaults

`func NewRbacManagerAccountRoleWithDefaults() *RbacManagerAccountRole`

NewRbacManagerAccountRoleWithDefaults instantiates a new RbacManagerAccountRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForAccount

`func (o *RbacManagerAccountRole) GetForAccount() string`

GetForAccount returns the ForAccount field if non-nil, zero value otherwise.

### GetForAccountOk

`func (o *RbacManagerAccountRole) GetForAccountOk() (*string, bool)`

GetForAccountOk returns a tuple with the ForAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAccount

`func (o *RbacManagerAccountRole) SetForAccount(v string)`

SetForAccount sets ForAccount field to given value.

### HasForAccount

`func (o *RbacManagerAccountRole) HasForAccount() bool`

HasForAccount returns a boolean if a field has been set.

### GetRoles

`func (o *RbacManagerAccountRole) GetRoles() RbacManagerRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RbacManagerAccountRole) GetRolesOk() (*RbacManagerRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RbacManagerAccountRole) SetRoles(v RbacManagerRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *RbacManagerAccountRole) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAccount

`func (o *RbacManagerAccountRole) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RbacManagerAccountRole) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RbacManagerAccountRole) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *RbacManagerAccountRole) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


