# UnifiedRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleName** | Pointer to **string** | The name of the RBAC Role | [optional] 
**DomainName** | Pointer to **string** | The domain (or account) name that provides the scope of the role | [optional] 
**Granter** | Pointer to **string** | The name of the user group that granted the role.  Will be null if the role was granted directly to the user. | [optional] 
**GrantType** | Pointer to **string** | The type of grant that was made | [optional] 
**GrantedAt** | Pointer to **time.Time** | The timestamp of when the role was granted | [optional] 

## Methods

### NewUnifiedRoles

`func NewUnifiedRoles() *UnifiedRoles`

NewUnifiedRoles instantiates a new UnifiedRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnifiedRolesWithDefaults

`func NewUnifiedRolesWithDefaults() *UnifiedRoles`

NewUnifiedRolesWithDefaults instantiates a new UnifiedRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleName

`func (o *UnifiedRoles) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *UnifiedRoles) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *UnifiedRoles) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *UnifiedRoles) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetDomainName

`func (o *UnifiedRoles) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *UnifiedRoles) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *UnifiedRoles) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *UnifiedRoles) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetGranter

`func (o *UnifiedRoles) GetGranter() string`

GetGranter returns the Granter field if non-nil, zero value otherwise.

### GetGranterOk

`func (o *UnifiedRoles) GetGranterOk() (*string, bool)`

GetGranterOk returns a tuple with the Granter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranter

`func (o *UnifiedRoles) SetGranter(v string)`

SetGranter sets Granter field to given value.

### HasGranter

`func (o *UnifiedRoles) HasGranter() bool`

HasGranter returns a boolean if a field has been set.

### GetGrantType

`func (o *UnifiedRoles) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *UnifiedRoles) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *UnifiedRoles) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *UnifiedRoles) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetGrantedAt

`func (o *UnifiedRoles) GetGrantedAt() time.Time`

GetGrantedAt returns the GrantedAt field if non-nil, zero value otherwise.

### GetGrantedAtOk

`func (o *UnifiedRoles) GetGrantedAtOk() (*time.Time, bool)`

GetGrantedAtOk returns a tuple with the GrantedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedAt

`func (o *UnifiedRoles) SetGrantedAt(v time.Time)`

SetGrantedAt sets GrantedAt field to given value.

### HasGrantedAt

`func (o *UnifiedRoles) HasGrantedAt() bool`

HasGrantedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


