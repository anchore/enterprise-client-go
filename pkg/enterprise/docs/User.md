# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username to authenticate with | 
**Type** | Pointer to **string** | The user&#39;s type | [optional] 
**Source** | Pointer to **NullableString** | When the user &#39;type&#39; is &#39;saml&#39;, this will be the EntityId of the IDP that they are authenticating from. Otherwise, this will be set to null. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp of when the user record was created | [optional] 
**LastUpdated** | Pointer to **time.Time** | The timestamp of the last update to this record | [optional] 
**IdpName** | Pointer to **NullableString** | When the user &#39;type&#39; is &#39;saml&#39;, this will be the configured name of the IDP that they are authenticating from.  Otherwise, this will be set to null. | [optional] 
**PasswordLastUpdated** | Pointer to **NullableTime** | When the user &#39;type&#39; is &#39;native&#39;, this will be the timestamp of the last time this user&#39;s credentials were updated. | [optional] 
**UnifiedRoles** | Pointer to [**[]UnifiedRoles**](UnifiedRoles.md) | The unified list of RBAC roles this user currently has in this account. | [optional] 

## Methods

### NewUser

`func NewUser(username string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetType

`func (o *User) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *User) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *User) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *User) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSource

`func (o *User) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *User) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *User) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *User) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *User) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *User) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *User) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *User) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *User) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *User) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetIdpName

`func (o *User) GetIdpName() string`

GetIdpName returns the IdpName field if non-nil, zero value otherwise.

### GetIdpNameOk

`func (o *User) GetIdpNameOk() (*string, bool)`

GetIdpNameOk returns a tuple with the IdpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpName

`func (o *User) SetIdpName(v string)`

SetIdpName sets IdpName field to given value.

### HasIdpName

`func (o *User) HasIdpName() bool`

HasIdpName returns a boolean if a field has been set.

### SetIdpNameNil

`func (o *User) SetIdpNameNil(b bool)`

 SetIdpNameNil sets the value for IdpName to be an explicit nil

### UnsetIdpName
`func (o *User) UnsetIdpName()`

UnsetIdpName ensures that no value is present for IdpName, not even an explicit nil
### GetPasswordLastUpdated

`func (o *User) GetPasswordLastUpdated() time.Time`

GetPasswordLastUpdated returns the PasswordLastUpdated field if non-nil, zero value otherwise.

### GetPasswordLastUpdatedOk

`func (o *User) GetPasswordLastUpdatedOk() (*time.Time, bool)`

GetPasswordLastUpdatedOk returns a tuple with the PasswordLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastUpdated

`func (o *User) SetPasswordLastUpdated(v time.Time)`

SetPasswordLastUpdated sets PasswordLastUpdated field to given value.

### HasPasswordLastUpdated

`func (o *User) HasPasswordLastUpdated() bool`

HasPasswordLastUpdated returns a boolean if a field has been set.

### SetPasswordLastUpdatedNil

`func (o *User) SetPasswordLastUpdatedNil(b bool)`

 SetPasswordLastUpdatedNil sets the value for PasswordLastUpdated to be an explicit nil

### UnsetPasswordLastUpdated
`func (o *User) UnsetPasswordLastUpdated()`

UnsetPasswordLastUpdated ensures that no value is present for PasswordLastUpdated, not even an explicit nil
### GetUnifiedRoles

`func (o *User) GetUnifiedRoles() []UnifiedRoles`

GetUnifiedRoles returns the UnifiedRoles field if non-nil, zero value otherwise.

### GetUnifiedRolesOk

`func (o *User) GetUnifiedRolesOk() (*[]UnifiedRoles, bool)`

GetUnifiedRolesOk returns a tuple with the UnifiedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedRoles

`func (o *User) SetUnifiedRoles(v []UnifiedRoles)`

SetUnifiedRoles sets UnifiedRoles field to given value.

### HasUnifiedRoles

`func (o *User) HasUnifiedRoles() bool`

HasUnifiedRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


