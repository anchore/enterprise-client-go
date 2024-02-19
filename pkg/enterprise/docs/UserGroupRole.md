# UserGroupRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** | The account for this role | 
**Roles** | **[]string** |  | 

## Methods

### NewUserGroupRole

`func NewUserGroupRole(accountName string, roles []string, ) *UserGroupRole`

NewUserGroupRole instantiates a new UserGroupRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupRoleWithDefaults

`func NewUserGroupRoleWithDefaults() *UserGroupRole`

NewUserGroupRoleWithDefaults instantiates a new UserGroupRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *UserGroupRole) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *UserGroupRole) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *UserGroupRole) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetRoles

`func (o *UserGroupRole) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserGroupRole) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserGroupRole) SetRoles(v []string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


