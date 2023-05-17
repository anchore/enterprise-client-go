/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 1.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// RbacManagerSamlConfiguration A named configuration for interaction with an Identity Provider that supports SAML 2.0
type RbacManagerSamlConfiguration struct {
	// The name to use for referencing this IDP configuration. This will configured as part of the url string the Idp must have the client POST the saml assertion to.
	Name string `json:"name"`
	// If this IDP configuration should be enabled for user logins
	Enabled bool `json:"enabled"`
	// The entity ID for this SP. Can be the same for all IDP configurations in this installation or unique to each. This is typically a URL, but you can use any value as long as you also configure the IDP to expect this value.
	SpEntityId string `json:"sp_entity_id"`
	// The URL the IDP can use to access the Assertion Consumer Service to provide the token for sso. This is the way to reach the rbac manager services /service/sso/auth/{IDP_name} route externally
	AcsUrl string `json:"acs_url"`
	// The port number to use for https if not 443. If omitted or -1, 443 is assumed and used as a default
	AcsHttpsPort *int32 `json:"acs_https_port,omitempty"`
	// The url where the SP (anchore) can retrieve the metadata about the Identity Provider. Only one of this or metadata_xml should be set. This is typically provided by the IDP.
	IdpMetadataUrl *string `json:"idp_metadata_url,omitempty"`
	// The direct metadata xml payload, if a url is not available. Only one of this or metadata_url should be set.
	IdpMetadataXml *string `json:"idp_metadata_xml,omitempty"`
	// The SAML attribute to use from the response assertions to determine the anchore username. If unset, the subject is used.
	IdpUsernameAttribute *string `json:"idp_username_attribute,omitempty"`
	// The SAML attribute to use from the response assertions to determine the anchore account to use. If unset, the default is used.
	IdpAccountAttribute *string `json:"idp_account_attribute,omitempty"`
	// The SAML attribute to use from the response assertions to determine the anchore role(s) to assign a new user in the specified account. If unset, the default is used.
	IdpRoleAttribute *string `json:"idp_role_attribute,omitempty"`
	// The anchore account to assign all users to from this IDP if no account attribute is mapped or present.
	DefaultAccount *string `json:"default_account,omitempty"`
	// The default role to apply to new users from this IDP if no attribute is mapped or found in the SAML assertions.
	DefaultRole *string `json:"default_role,omitempty"`
	// Require assertions in to be signed from the IDP
	RequireSignedAssertions *bool `json:"require_signed_assertions,omitempty"`
	// Require the authn response to be signed by the IDP
	RequireSignedResponse *bool `json:"require_signed_response,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Indicates if Anchore will require an authenticating SSO user to already exist.  This field is ignored on POST/PUT Operations.
	RequireExistingUsers *bool `json:"require_existing_users,omitempty"`
}

// NewRbacManagerSamlConfiguration instantiates a new RbacManagerSamlConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerSamlConfiguration(name string, enabled bool, spEntityId string, acsUrl string) *RbacManagerSamlConfiguration {
	this := RbacManagerSamlConfiguration{}
	this.Name = name
	this.Enabled = enabled
	this.SpEntityId = spEntityId
	this.AcsUrl = acsUrl
	var requireSignedAssertions bool = true
	this.RequireSignedAssertions = &requireSignedAssertions
	var requireSignedResponse bool = true
	this.RequireSignedResponse = &requireSignedResponse
	var requireExistingUsers bool = false
	this.RequireExistingUsers = &requireExistingUsers
	return &this
}

// NewRbacManagerSamlConfigurationWithDefaults instantiates a new RbacManagerSamlConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerSamlConfigurationWithDefaults() *RbacManagerSamlConfiguration {
	this := RbacManagerSamlConfiguration{}
	var requireSignedAssertions bool = true
	this.RequireSignedAssertions = &requireSignedAssertions
	var requireSignedResponse bool = true
	this.RequireSignedResponse = &requireSignedResponse
	var requireExistingUsers bool = false
	this.RequireExistingUsers = &requireExistingUsers
	return &this
}

// GetName returns the Name field value
func (o *RbacManagerSamlConfiguration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RbacManagerSamlConfiguration) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *RbacManagerSamlConfiguration) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *RbacManagerSamlConfiguration) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSpEntityId returns the SpEntityId field value
func (o *RbacManagerSamlConfiguration) GetSpEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpEntityId
}

// GetSpEntityIdOk returns a tuple with the SpEntityId field value
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetSpEntityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SpEntityId, true
}

// SetSpEntityId sets field value
func (o *RbacManagerSamlConfiguration) SetSpEntityId(v string) {
	o.SpEntityId = v
}

// GetAcsUrl returns the AcsUrl field value
func (o *RbacManagerSamlConfiguration) GetAcsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetAcsUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AcsUrl, true
}

// SetAcsUrl sets field value
func (o *RbacManagerSamlConfiguration) SetAcsUrl(v string) {
	o.AcsUrl = v
}

// GetAcsHttpsPort returns the AcsHttpsPort field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetAcsHttpsPort() int32 {
	if o == nil || o.AcsHttpsPort == nil {
		var ret int32
		return ret
	}
	return *o.AcsHttpsPort
}

// GetAcsHttpsPortOk returns a tuple with the AcsHttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetAcsHttpsPortOk() (*int32, bool) {
	if o == nil || o.AcsHttpsPort == nil {
		return nil, false
	}
	return o.AcsHttpsPort, true
}

// HasAcsHttpsPort returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasAcsHttpsPort() bool {
	if o != nil && o.AcsHttpsPort != nil {
		return true
	}

	return false
}

// SetAcsHttpsPort gets a reference to the given int32 and assigns it to the AcsHttpsPort field.
func (o *RbacManagerSamlConfiguration) SetAcsHttpsPort(v int32) {
	o.AcsHttpsPort = &v
}

// GetIdpMetadataUrl returns the IdpMetadataUrl field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetIdpMetadataUrl() string {
	if o == nil || o.IdpMetadataUrl == nil {
		var ret string
		return ret
	}
	return *o.IdpMetadataUrl
}

// GetIdpMetadataUrlOk returns a tuple with the IdpMetadataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetIdpMetadataUrlOk() (*string, bool) {
	if o == nil || o.IdpMetadataUrl == nil {
		return nil, false
	}
	return o.IdpMetadataUrl, true
}

// HasIdpMetadataUrl returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasIdpMetadataUrl() bool {
	if o != nil && o.IdpMetadataUrl != nil {
		return true
	}

	return false
}

// SetIdpMetadataUrl gets a reference to the given string and assigns it to the IdpMetadataUrl field.
func (o *RbacManagerSamlConfiguration) SetIdpMetadataUrl(v string) {
	o.IdpMetadataUrl = &v
}

// GetIdpMetadataXml returns the IdpMetadataXml field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetIdpMetadataXml() string {
	if o == nil || o.IdpMetadataXml == nil {
		var ret string
		return ret
	}
	return *o.IdpMetadataXml
}

// GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetIdpMetadataXmlOk() (*string, bool) {
	if o == nil || o.IdpMetadataXml == nil {
		return nil, false
	}
	return o.IdpMetadataXml, true
}

// HasIdpMetadataXml returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasIdpMetadataXml() bool {
	if o != nil && o.IdpMetadataXml != nil {
		return true
	}

	return false
}

// SetIdpMetadataXml gets a reference to the given string and assigns it to the IdpMetadataXml field.
func (o *RbacManagerSamlConfiguration) SetIdpMetadataXml(v string) {
	o.IdpMetadataXml = &v
}

// GetIdpUsernameAttribute returns the IdpUsernameAttribute field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetIdpUsernameAttribute() string {
	if o == nil || o.IdpUsernameAttribute == nil {
		var ret string
		return ret
	}
	return *o.IdpUsernameAttribute
}

// GetIdpUsernameAttributeOk returns a tuple with the IdpUsernameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetIdpUsernameAttributeOk() (*string, bool) {
	if o == nil || o.IdpUsernameAttribute == nil {
		return nil, false
	}
	return o.IdpUsernameAttribute, true
}

// HasIdpUsernameAttribute returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasIdpUsernameAttribute() bool {
	if o != nil && o.IdpUsernameAttribute != nil {
		return true
	}

	return false
}

// SetIdpUsernameAttribute gets a reference to the given string and assigns it to the IdpUsernameAttribute field.
func (o *RbacManagerSamlConfiguration) SetIdpUsernameAttribute(v string) {
	o.IdpUsernameAttribute = &v
}

// GetIdpAccountAttribute returns the IdpAccountAttribute field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetIdpAccountAttribute() string {
	if o == nil || o.IdpAccountAttribute == nil {
		var ret string
		return ret
	}
	return *o.IdpAccountAttribute
}

// GetIdpAccountAttributeOk returns a tuple with the IdpAccountAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetIdpAccountAttributeOk() (*string, bool) {
	if o == nil || o.IdpAccountAttribute == nil {
		return nil, false
	}
	return o.IdpAccountAttribute, true
}

// HasIdpAccountAttribute returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasIdpAccountAttribute() bool {
	if o != nil && o.IdpAccountAttribute != nil {
		return true
	}

	return false
}

// SetIdpAccountAttribute gets a reference to the given string and assigns it to the IdpAccountAttribute field.
func (o *RbacManagerSamlConfiguration) SetIdpAccountAttribute(v string) {
	o.IdpAccountAttribute = &v
}

// GetIdpRoleAttribute returns the IdpRoleAttribute field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetIdpRoleAttribute() string {
	if o == nil || o.IdpRoleAttribute == nil {
		var ret string
		return ret
	}
	return *o.IdpRoleAttribute
}

// GetIdpRoleAttributeOk returns a tuple with the IdpRoleAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetIdpRoleAttributeOk() (*string, bool) {
	if o == nil || o.IdpRoleAttribute == nil {
		return nil, false
	}
	return o.IdpRoleAttribute, true
}

// HasIdpRoleAttribute returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasIdpRoleAttribute() bool {
	if o != nil && o.IdpRoleAttribute != nil {
		return true
	}

	return false
}

// SetIdpRoleAttribute gets a reference to the given string and assigns it to the IdpRoleAttribute field.
func (o *RbacManagerSamlConfiguration) SetIdpRoleAttribute(v string) {
	o.IdpRoleAttribute = &v
}

// GetDefaultAccount returns the DefaultAccount field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetDefaultAccount() string {
	if o == nil || o.DefaultAccount == nil {
		var ret string
		return ret
	}
	return *o.DefaultAccount
}

// GetDefaultAccountOk returns a tuple with the DefaultAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetDefaultAccountOk() (*string, bool) {
	if o == nil || o.DefaultAccount == nil {
		return nil, false
	}
	return o.DefaultAccount, true
}

// HasDefaultAccount returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasDefaultAccount() bool {
	if o != nil && o.DefaultAccount != nil {
		return true
	}

	return false
}

// SetDefaultAccount gets a reference to the given string and assigns it to the DefaultAccount field.
func (o *RbacManagerSamlConfiguration) SetDefaultAccount(v string) {
	o.DefaultAccount = &v
}

// GetDefaultRole returns the DefaultRole field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetDefaultRole() string {
	if o == nil || o.DefaultRole == nil {
		var ret string
		return ret
	}
	return *o.DefaultRole
}

// GetDefaultRoleOk returns a tuple with the DefaultRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetDefaultRoleOk() (*string, bool) {
	if o == nil || o.DefaultRole == nil {
		return nil, false
	}
	return o.DefaultRole, true
}

// HasDefaultRole returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasDefaultRole() bool {
	if o != nil && o.DefaultRole != nil {
		return true
	}

	return false
}

// SetDefaultRole gets a reference to the given string and assigns it to the DefaultRole field.
func (o *RbacManagerSamlConfiguration) SetDefaultRole(v string) {
	o.DefaultRole = &v
}

// GetRequireSignedAssertions returns the RequireSignedAssertions field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetRequireSignedAssertions() bool {
	if o == nil || o.RequireSignedAssertions == nil {
		var ret bool
		return ret
	}
	return *o.RequireSignedAssertions
}

// GetRequireSignedAssertionsOk returns a tuple with the RequireSignedAssertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetRequireSignedAssertionsOk() (*bool, bool) {
	if o == nil || o.RequireSignedAssertions == nil {
		return nil, false
	}
	return o.RequireSignedAssertions, true
}

// HasRequireSignedAssertions returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasRequireSignedAssertions() bool {
	if o != nil && o.RequireSignedAssertions != nil {
		return true
	}

	return false
}

// SetRequireSignedAssertions gets a reference to the given bool and assigns it to the RequireSignedAssertions field.
func (o *RbacManagerSamlConfiguration) SetRequireSignedAssertions(v bool) {
	o.RequireSignedAssertions = &v
}

// GetRequireSignedResponse returns the RequireSignedResponse field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetRequireSignedResponse() bool {
	if o == nil || o.RequireSignedResponse == nil {
		var ret bool
		return ret
	}
	return *o.RequireSignedResponse
}

// GetRequireSignedResponseOk returns a tuple with the RequireSignedResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetRequireSignedResponseOk() (*bool, bool) {
	if o == nil || o.RequireSignedResponse == nil {
		return nil, false
	}
	return o.RequireSignedResponse, true
}

// HasRequireSignedResponse returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasRequireSignedResponse() bool {
	if o != nil && o.RequireSignedResponse != nil {
		return true
	}

	return false
}

// SetRequireSignedResponse gets a reference to the given bool and assigns it to the RequireSignedResponse field.
func (o *RbacManagerSamlConfiguration) SetRequireSignedResponse(v bool) {
	o.RequireSignedResponse = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RbacManagerSamlConfiguration) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *RbacManagerSamlConfiguration) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetRequireExistingUsers returns the RequireExistingUsers field value if set, zero value otherwise.
func (o *RbacManagerSamlConfiguration) GetRequireExistingUsers() bool {
	if o == nil || o.RequireExistingUsers == nil {
		var ret bool
		return ret
	}
	return *o.RequireExistingUsers
}

// GetRequireExistingUsersOk returns a tuple with the RequireExistingUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfiguration) GetRequireExistingUsersOk() (*bool, bool) {
	if o == nil || o.RequireExistingUsers == nil {
		return nil, false
	}
	return o.RequireExistingUsers, true
}

// HasRequireExistingUsers returns a boolean if a field has been set.
func (o *RbacManagerSamlConfiguration) HasRequireExistingUsers() bool {
	if o != nil && o.RequireExistingUsers != nil {
		return true
	}

	return false
}

// SetRequireExistingUsers gets a reference to the given bool and assigns it to the RequireExistingUsers field.
func (o *RbacManagerSamlConfiguration) SetRequireExistingUsers(v bool) {
	o.RequireExistingUsers = &v
}

func (o RbacManagerSamlConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["sp_entity_id"] = o.SpEntityId
	}
	if true {
		toSerialize["acs_url"] = o.AcsUrl
	}
	if o.AcsHttpsPort != nil {
		toSerialize["acs_https_port"] = o.AcsHttpsPort
	}
	if o.IdpMetadataUrl != nil {
		toSerialize["idp_metadata_url"] = o.IdpMetadataUrl
	}
	if o.IdpMetadataXml != nil {
		toSerialize["idp_metadata_xml"] = o.IdpMetadataXml
	}
	if o.IdpUsernameAttribute != nil {
		toSerialize["idp_username_attribute"] = o.IdpUsernameAttribute
	}
	if o.IdpAccountAttribute != nil {
		toSerialize["idp_account_attribute"] = o.IdpAccountAttribute
	}
	if o.IdpRoleAttribute != nil {
		toSerialize["idp_role_attribute"] = o.IdpRoleAttribute
	}
	if o.DefaultAccount != nil {
		toSerialize["default_account"] = o.DefaultAccount
	}
	if o.DefaultRole != nil {
		toSerialize["default_role"] = o.DefaultRole
	}
	if o.RequireSignedAssertions != nil {
		toSerialize["require_signed_assertions"] = o.RequireSignedAssertions
	}
	if o.RequireSignedResponse != nil {
		toSerialize["require_signed_response"] = o.RequireSignedResponse
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.RequireExistingUsers != nil {
		toSerialize["require_existing_users"] = o.RequireExistingUsers
	}
	return json.Marshal(toSerialize)
}

type NullableRbacManagerSamlConfiguration struct {
	value *RbacManagerSamlConfiguration
	isSet bool
}

func (v NullableRbacManagerSamlConfiguration) Get() *RbacManagerSamlConfiguration {
	return v.value
}

func (v *NullableRbacManagerSamlConfiguration) Set(val *RbacManagerSamlConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerSamlConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerSamlConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerSamlConfiguration(val *RbacManagerSamlConfiguration) *NullableRbacManagerSamlConfiguration {
	return &NullableRbacManagerSamlConfiguration{value: val, isSet: true}
}

func (v NullableRbacManagerSamlConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerSamlConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


