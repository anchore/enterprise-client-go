# RbacManagerSamlConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name to use for referencing this IDP configuration. This will configured as part of the url string the Idp must have the client POST the saml assertion to. | 
**Enabled** | **bool** | If this IDP configuration should be enabled for user logins | 
**SpEntityId** | **string** | The entity ID for this SP. Can be the same for all IDP configurations in this installation or unique to each. This is typically a URL, but you can use any value as long as you also configure the IDP to expect this value. | 
**AcsUrl** | **string** | The URL the IDP can use to access the Assertion Consumer Service to provide the token for sso. This is the way to reach the rbac manager services /service/sso/auth/{IDP_name} route externally | 
**AcsHttpsPort** | Pointer to **int32** | The port number to use for https if not 443. If omitted or -1, 443 is assumed and used as a default | [optional] 
**IdpMetadataUrl** | Pointer to **string** | The url where the SP (anchore) can retrieve the metadata about the Identity Provider. Only one of this or metadata_xml should be set. This is typically provided by the IDP. | [optional] 
**IdpMetadataXml** | Pointer to **string** | The direct metadata xml payload, if a url is not available. Only one of this or metadata_url should be set. | [optional] 
**IdpUsernameAttribute** | Pointer to **string** | The SAML attribute to use from the response assertions to determine the anchore username. If unset, the subject is used. | [optional] 
**IdpAccountAttribute** | Pointer to **string** | The SAML attribute to use from the response assertions to determine the anchore account to use. If unset, the default is used. | [optional] 
**IdpRoleAttribute** | Pointer to **string** | The SAML attribute to use from the response assertions to determine the anchore role(s) to assign a new user in the specified account. If unset, the default is used. | [optional] 
**IdpGroupsAttribute** | Pointer to **string** | The SAML attribute to use from the response assertions to determine the SSO groups the user is a member of. If unset, the default is used (memberOf). | [optional] 
**DefaultAccount** | Pointer to **string** | The anchore account to assign all users to from this IDP if no account attribute is mapped or present. | [optional] 
**DefaultRole** | Pointer to **string** | The default role to apply to new users from this IDP if no attribute is mapped or found in the SAML assertions. | [optional] 
**RequireSignedAssertions** | Pointer to **bool** | Require assertions in to be signed from the IDP | [optional] [default to true]
**RequireSignedResponse** | Pointer to **bool** | Require the authn response to be signed by the IDP | [optional] [default to true]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**RequireExistingUsers** | Pointer to **bool** | Indicates if Anchore will require an authenticating SSO user to already exist.  This field is ignored on POST/PUT Operations. | [optional] [default to false]

## Methods

### NewRbacManagerSamlConfiguration

`func NewRbacManagerSamlConfiguration(name string, enabled bool, spEntityId string, acsUrl string, ) *RbacManagerSamlConfiguration`

NewRbacManagerSamlConfiguration instantiates a new RbacManagerSamlConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacManagerSamlConfigurationWithDefaults

`func NewRbacManagerSamlConfigurationWithDefaults() *RbacManagerSamlConfiguration`

NewRbacManagerSamlConfigurationWithDefaults instantiates a new RbacManagerSamlConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RbacManagerSamlConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RbacManagerSamlConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RbacManagerSamlConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *RbacManagerSamlConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RbacManagerSamlConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RbacManagerSamlConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSpEntityId

`func (o *RbacManagerSamlConfiguration) GetSpEntityId() string`

GetSpEntityId returns the SpEntityId field if non-nil, zero value otherwise.

### GetSpEntityIdOk

`func (o *RbacManagerSamlConfiguration) GetSpEntityIdOk() (*string, bool)`

GetSpEntityIdOk returns a tuple with the SpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpEntityId

`func (o *RbacManagerSamlConfiguration) SetSpEntityId(v string)`

SetSpEntityId sets SpEntityId field to given value.


### GetAcsUrl

`func (o *RbacManagerSamlConfiguration) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *RbacManagerSamlConfiguration) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *RbacManagerSamlConfiguration) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.


### GetAcsHttpsPort

`func (o *RbacManagerSamlConfiguration) GetAcsHttpsPort() int32`

GetAcsHttpsPort returns the AcsHttpsPort field if non-nil, zero value otherwise.

### GetAcsHttpsPortOk

`func (o *RbacManagerSamlConfiguration) GetAcsHttpsPortOk() (*int32, bool)`

GetAcsHttpsPortOk returns a tuple with the AcsHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsHttpsPort

`func (o *RbacManagerSamlConfiguration) SetAcsHttpsPort(v int32)`

SetAcsHttpsPort sets AcsHttpsPort field to given value.

### HasAcsHttpsPort

`func (o *RbacManagerSamlConfiguration) HasAcsHttpsPort() bool`

HasAcsHttpsPort returns a boolean if a field has been set.

### GetIdpMetadataUrl

`func (o *RbacManagerSamlConfiguration) GetIdpMetadataUrl() string`

GetIdpMetadataUrl returns the IdpMetadataUrl field if non-nil, zero value otherwise.

### GetIdpMetadataUrlOk

`func (o *RbacManagerSamlConfiguration) GetIdpMetadataUrlOk() (*string, bool)`

GetIdpMetadataUrlOk returns a tuple with the IdpMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataUrl

`func (o *RbacManagerSamlConfiguration) SetIdpMetadataUrl(v string)`

SetIdpMetadataUrl sets IdpMetadataUrl field to given value.

### HasIdpMetadataUrl

`func (o *RbacManagerSamlConfiguration) HasIdpMetadataUrl() bool`

HasIdpMetadataUrl returns a boolean if a field has been set.

### GetIdpMetadataXml

`func (o *RbacManagerSamlConfiguration) GetIdpMetadataXml() string`

GetIdpMetadataXml returns the IdpMetadataXml field if non-nil, zero value otherwise.

### GetIdpMetadataXmlOk

`func (o *RbacManagerSamlConfiguration) GetIdpMetadataXmlOk() (*string, bool)`

GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataXml

`func (o *RbacManagerSamlConfiguration) SetIdpMetadataXml(v string)`

SetIdpMetadataXml sets IdpMetadataXml field to given value.

### HasIdpMetadataXml

`func (o *RbacManagerSamlConfiguration) HasIdpMetadataXml() bool`

HasIdpMetadataXml returns a boolean if a field has been set.

### GetIdpUsernameAttribute

`func (o *RbacManagerSamlConfiguration) GetIdpUsernameAttribute() string`

GetIdpUsernameAttribute returns the IdpUsernameAttribute field if non-nil, zero value otherwise.

### GetIdpUsernameAttributeOk

`func (o *RbacManagerSamlConfiguration) GetIdpUsernameAttributeOk() (*string, bool)`

GetIdpUsernameAttributeOk returns a tuple with the IdpUsernameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpUsernameAttribute

`func (o *RbacManagerSamlConfiguration) SetIdpUsernameAttribute(v string)`

SetIdpUsernameAttribute sets IdpUsernameAttribute field to given value.

### HasIdpUsernameAttribute

`func (o *RbacManagerSamlConfiguration) HasIdpUsernameAttribute() bool`

HasIdpUsernameAttribute returns a boolean if a field has been set.

### GetIdpAccountAttribute

`func (o *RbacManagerSamlConfiguration) GetIdpAccountAttribute() string`

GetIdpAccountAttribute returns the IdpAccountAttribute field if non-nil, zero value otherwise.

### GetIdpAccountAttributeOk

`func (o *RbacManagerSamlConfiguration) GetIdpAccountAttributeOk() (*string, bool)`

GetIdpAccountAttributeOk returns a tuple with the IdpAccountAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpAccountAttribute

`func (o *RbacManagerSamlConfiguration) SetIdpAccountAttribute(v string)`

SetIdpAccountAttribute sets IdpAccountAttribute field to given value.

### HasIdpAccountAttribute

`func (o *RbacManagerSamlConfiguration) HasIdpAccountAttribute() bool`

HasIdpAccountAttribute returns a boolean if a field has been set.

### GetIdpRoleAttribute

`func (o *RbacManagerSamlConfiguration) GetIdpRoleAttribute() string`

GetIdpRoleAttribute returns the IdpRoleAttribute field if non-nil, zero value otherwise.

### GetIdpRoleAttributeOk

`func (o *RbacManagerSamlConfiguration) GetIdpRoleAttributeOk() (*string, bool)`

GetIdpRoleAttributeOk returns a tuple with the IdpRoleAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpRoleAttribute

`func (o *RbacManagerSamlConfiguration) SetIdpRoleAttribute(v string)`

SetIdpRoleAttribute sets IdpRoleAttribute field to given value.

### HasIdpRoleAttribute

`func (o *RbacManagerSamlConfiguration) HasIdpRoleAttribute() bool`

HasIdpRoleAttribute returns a boolean if a field has been set.

### GetIdpGroupsAttribute

`func (o *RbacManagerSamlConfiguration) GetIdpGroupsAttribute() string`

GetIdpGroupsAttribute returns the IdpGroupsAttribute field if non-nil, zero value otherwise.

### GetIdpGroupsAttributeOk

`func (o *RbacManagerSamlConfiguration) GetIdpGroupsAttributeOk() (*string, bool)`

GetIdpGroupsAttributeOk returns a tuple with the IdpGroupsAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpGroupsAttribute

`func (o *RbacManagerSamlConfiguration) SetIdpGroupsAttribute(v string)`

SetIdpGroupsAttribute sets IdpGroupsAttribute field to given value.

### HasIdpGroupsAttribute

`func (o *RbacManagerSamlConfiguration) HasIdpGroupsAttribute() bool`

HasIdpGroupsAttribute returns a boolean if a field has been set.

### GetDefaultAccount

`func (o *RbacManagerSamlConfiguration) GetDefaultAccount() string`

GetDefaultAccount returns the DefaultAccount field if non-nil, zero value otherwise.

### GetDefaultAccountOk

`func (o *RbacManagerSamlConfiguration) GetDefaultAccountOk() (*string, bool)`

GetDefaultAccountOk returns a tuple with the DefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccount

`func (o *RbacManagerSamlConfiguration) SetDefaultAccount(v string)`

SetDefaultAccount sets DefaultAccount field to given value.

### HasDefaultAccount

`func (o *RbacManagerSamlConfiguration) HasDefaultAccount() bool`

HasDefaultAccount returns a boolean if a field has been set.

### GetDefaultRole

`func (o *RbacManagerSamlConfiguration) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *RbacManagerSamlConfiguration) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *RbacManagerSamlConfiguration) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *RbacManagerSamlConfiguration) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetRequireSignedAssertions

`func (o *RbacManagerSamlConfiguration) GetRequireSignedAssertions() bool`

GetRequireSignedAssertions returns the RequireSignedAssertions field if non-nil, zero value otherwise.

### GetRequireSignedAssertionsOk

`func (o *RbacManagerSamlConfiguration) GetRequireSignedAssertionsOk() (*bool, bool)`

GetRequireSignedAssertionsOk returns a tuple with the RequireSignedAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedAssertions

`func (o *RbacManagerSamlConfiguration) SetRequireSignedAssertions(v bool)`

SetRequireSignedAssertions sets RequireSignedAssertions field to given value.

### HasRequireSignedAssertions

`func (o *RbacManagerSamlConfiguration) HasRequireSignedAssertions() bool`

HasRequireSignedAssertions returns a boolean if a field has been set.

### GetRequireSignedResponse

`func (o *RbacManagerSamlConfiguration) GetRequireSignedResponse() bool`

GetRequireSignedResponse returns the RequireSignedResponse field if non-nil, zero value otherwise.

### GetRequireSignedResponseOk

`func (o *RbacManagerSamlConfiguration) GetRequireSignedResponseOk() (*bool, bool)`

GetRequireSignedResponseOk returns a tuple with the RequireSignedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedResponse

`func (o *RbacManagerSamlConfiguration) SetRequireSignedResponse(v bool)`

SetRequireSignedResponse sets RequireSignedResponse field to given value.

### HasRequireSignedResponse

`func (o *RbacManagerSamlConfiguration) HasRequireSignedResponse() bool`

HasRequireSignedResponse returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RbacManagerSamlConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RbacManagerSamlConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RbacManagerSamlConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RbacManagerSamlConfiguration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RbacManagerSamlConfiguration) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RbacManagerSamlConfiguration) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RbacManagerSamlConfiguration) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *RbacManagerSamlConfiguration) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRequireExistingUsers

`func (o *RbacManagerSamlConfiguration) GetRequireExistingUsers() bool`

GetRequireExistingUsers returns the RequireExistingUsers field if non-nil, zero value otherwise.

### GetRequireExistingUsersOk

`func (o *RbacManagerSamlConfiguration) GetRequireExistingUsersOk() (*bool, bool)`

GetRequireExistingUsersOk returns a tuple with the RequireExistingUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExistingUsers

`func (o *RbacManagerSamlConfiguration) SetRequireExistingUsers(v bool)`

SetRequireExistingUsers sets RequireExistingUsers field to given value.

### HasRequireExistingUsers

`func (o *RbacManagerSamlConfiguration) HasRequireExistingUsers() bool`

HasRequireExistingUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


