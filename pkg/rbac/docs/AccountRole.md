# AccountRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForAccount** | Pointer to **string** | The account scope that applies to the set of roles | [optional] 
**Roles** | Pointer to [**Role**](Role.md) |  | [optional] 

## Methods

### NewAccountRole

`func NewAccountRole() *AccountRole`

NewAccountRole instantiates a new AccountRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRoleWithDefaults

`func NewAccountRoleWithDefaults() *AccountRole`

NewAccountRoleWithDefaults instantiates a new AccountRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForAccount

`func (o *AccountRole) GetForAccount() string`

GetForAccount returns the ForAccount field if non-nil, zero value otherwise.

### GetForAccountOk

`func (o *AccountRole) GetForAccountOk() (*string, bool)`

GetForAccountOk returns a tuple with the ForAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAccount

`func (o *AccountRole) SetForAccount(v string)`

SetForAccount sets ForAccount field to given value.

### HasForAccount

`func (o *AccountRole) HasForAccount() bool`

HasForAccount returns a boolean if a field has been set.

### GetRoles

`func (o *AccountRole) GetRoles() Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccountRole) GetRolesOk() (*Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccountRole) SetRoles(v Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AccountRole) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


