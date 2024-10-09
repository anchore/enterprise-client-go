# UserGroupRolePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForAccount** | Pointer to **string** | Deprecated. Please use domain_name instead. The account | [optional] 
**DomainName** | Pointer to **string** | The domain scope for this role. This may be an account name when the domain is an account. | [optional] 
**Roles** | [**[]UserGroupRolePostRolesInner**](UserGroupRolePostRolesInner.md) |  | 

## Methods

### NewUserGroupRolePost

`func NewUserGroupRolePost(roles []UserGroupRolePostRolesInner, ) *UserGroupRolePost`

NewUserGroupRolePost instantiates a new UserGroupRolePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupRolePostWithDefaults

`func NewUserGroupRolePostWithDefaults() *UserGroupRolePost`

NewUserGroupRolePostWithDefaults instantiates a new UserGroupRolePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForAccount

`func (o *UserGroupRolePost) GetForAccount() string`

GetForAccount returns the ForAccount field if non-nil, zero value otherwise.

### GetForAccountOk

`func (o *UserGroupRolePost) GetForAccountOk() (*string, bool)`

GetForAccountOk returns a tuple with the ForAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAccount

`func (o *UserGroupRolePost) SetForAccount(v string)`

SetForAccount sets ForAccount field to given value.

### HasForAccount

`func (o *UserGroupRolePost) HasForAccount() bool`

HasForAccount returns a boolean if a field has been set.

### GetDomainName

`func (o *UserGroupRolePost) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *UserGroupRolePost) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *UserGroupRolePost) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *UserGroupRolePost) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetRoles

`func (o *UserGroupRolePost) GetRoles() []UserGroupRolePostRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserGroupRolePost) GetRolesOk() (*[]UserGroupRolePostRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserGroupRolePost) SetRoles(v []UserGroupRolePostRolesInner)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


