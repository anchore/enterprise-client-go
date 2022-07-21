/*
Anchore Enterprise RBAC API

Enterprise API for managing roles, permissions, and user mappings

API version: 0.0.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rbac

import (
	"encoding/json"
	"time"
)

// SamlConfiguration A named configuration for interaction with an Identity Provider that supports SAML 2.0
type SamlConfiguration struct {
	// The name to use for referencing this IDP configuration. This will configured as part of the url string the Idp must have the client POST the saml assertion to.
	Name string `json:"name"`
	// If this IDP configuration should be enabled for user logins
	Enabled bool `json:"enabled"`
	// The entity ID for this SP. Can be the same for all IDP configurations in this installation or unique to each. This is typically a URL, but you can use any value as long as you also configure the IDP to expect this value.
	SpEntityId string `json:"sp_entity_id"`
	// The URL the IDP can use to access the Assertion Consumer Service to provide the token for sso. This is the way to reach the rbac manager services /v1/saml/sso/{IDP_name} route externally
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
	// The existing anchore account to assign all users to from this IDP if no account attribute is mapped or present.
	DefaultAccount *string `json:"default_account,omitempty"`
	// The default role to apply to new users from this IDP if no attribute is mapped or found in the SAML assertions.
	DefaultRole *string `json:"default_role,omitempty"`
	// Require assertions in to be signed from the IDP
	RequireSignedAssertions *bool `json:"require_signed_assertions,omitempty"`
	// Require the authn response to be signed by the IDP
	RequireSignedResponse *bool `json:"require_signed_response,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewSamlConfiguration instantiates a new SamlConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlConfiguration(name string, enabled bool, spEntityId string, acsUrl string) *SamlConfiguration {
	this := SamlConfiguration{}
	this.Name = name
	this.Enabled = enabled
	this.SpEntityId = spEntityId
	this.AcsUrl = acsUrl
	var requireSignedAssertions bool = true
	this.RequireSignedAssertions = &requireSignedAssertions
	var requireSignedResponse bool = true
	this.RequireSignedResponse = &requireSignedResponse
	return &this
}

// NewSamlConfigurationWithDefaults instantiates a new SamlConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlConfigurationWithDefaults() *SamlConfiguration {
	this := SamlConfiguration{}
	var requireSignedAssertions bool = true
	this.RequireSignedAssertions = &requireSignedAssertions
	var requireSignedResponse bool = true
	this.RequireSignedResponse = &requireSignedResponse
	return &this
}

// GetName returns the Name field value
func (o *SamlConfiguration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SamlConfiguration) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *SamlConfiguration) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SamlConfiguration) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSpEntityId returns the SpEntityId field value
func (o *SamlConfiguration) GetSpEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpEntityId
}

// GetSpEntityIdOk returns a tuple with the SpEntityId field value
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetSpEntityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SpEntityId, true
}

// SetSpEntityId sets field value
func (o *SamlConfiguration) SetSpEntityId(v string) {
	o.SpEntityId = v
}

// GetAcsUrl returns the AcsUrl field value
func (o *SamlConfiguration) GetAcsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetAcsUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AcsUrl, true
}

// SetAcsUrl sets field value
func (o *SamlConfiguration) SetAcsUrl(v string) {
	o.AcsUrl = v
}

// GetAcsHttpsPort returns the AcsHttpsPort field value if set, zero value otherwise.
func (o *SamlConfiguration) GetAcsHttpsPort() int32 {
	if o == nil || o.AcsHttpsPort == nil {
		var ret int32
		return ret
	}
	return *o.AcsHttpsPort
}

// GetAcsHttpsPortOk returns a tuple with the AcsHttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetAcsHttpsPortOk() (*int32, bool) {
	if o == nil || o.AcsHttpsPort == nil {
		return nil, false
	}
	return o.AcsHttpsPort, true
}

// HasAcsHttpsPort returns a boolean if a field has been set.
func (o *SamlConfiguration) HasAcsHttpsPort() bool {
	if o != nil && o.AcsHttpsPort != nil {
		return true
	}

	return false
}

// SetAcsHttpsPort gets a reference to the given int32 and assigns it to the AcsHttpsPort field.
func (o *SamlConfiguration) SetAcsHttpsPort(v int32) {
	o.AcsHttpsPort = &v
}

// GetIdpMetadataUrl returns the IdpMetadataUrl field value if set, zero value otherwise.
func (o *SamlConfiguration) GetIdpMetadataUrl() string {
	if o == nil || o.IdpMetadataUrl == nil {
		var ret string
		return ret
	}
	return *o.IdpMetadataUrl
}

// GetIdpMetadataUrlOk returns a tuple with the IdpMetadataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetIdpMetadataUrlOk() (*string, bool) {
	if o == nil || o.IdpMetadataUrl == nil {
		return nil, false
	}
	return o.IdpMetadataUrl, true
}

// HasIdpMetadataUrl returns a boolean if a field has been set.
func (o *SamlConfiguration) HasIdpMetadataUrl() bool {
	if o != nil && o.IdpMetadataUrl != nil {
		return true
	}

	return false
}

// SetIdpMetadataUrl gets a reference to the given string and assigns it to the IdpMetadataUrl field.
func (o *SamlConfiguration) SetIdpMetadataUrl(v string) {
	o.IdpMetadataUrl = &v
}

// GetIdpMetadataXml returns the IdpMetadataXml field value if set, zero value otherwise.
func (o *SamlConfiguration) GetIdpMetadataXml() string {
	if o == nil || o.IdpMetadataXml == nil {
		var ret string
		return ret
	}
	return *o.IdpMetadataXml
}

// GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetIdpMetadataXmlOk() (*string, bool) {
	if o == nil || o.IdpMetadataXml == nil {
		return nil, false
	}
	return o.IdpMetadataXml, true
}

// HasIdpMetadataXml returns a boolean if a field has been set.
func (o *SamlConfiguration) HasIdpMetadataXml() bool {
	if o != nil && o.IdpMetadataXml != nil {
		return true
	}

	return false
}

// SetIdpMetadataXml gets a reference to the given string and assigns it to the IdpMetadataXml field.
func (o *SamlConfiguration) SetIdpMetadataXml(v string) {
	o.IdpMetadataXml = &v
}

// GetIdpUsernameAttribute returns the IdpUsernameAttribute field value if set, zero value otherwise.
func (o *SamlConfiguration) GetIdpUsernameAttribute() string {
	if o == nil || o.IdpUsernameAttribute == nil {
		var ret string
		return ret
	}
	return *o.IdpUsernameAttribute
}

// GetIdpUsernameAttributeOk returns a tuple with the IdpUsernameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetIdpUsernameAttributeOk() (*string, bool) {
	if o == nil || o.IdpUsernameAttribute == nil {
		return nil, false
	}
	return o.IdpUsernameAttribute, true
}

// HasIdpUsernameAttribute returns a boolean if a field has been set.
func (o *SamlConfiguration) HasIdpUsernameAttribute() bool {
	if o != nil && o.IdpUsernameAttribute != nil {
		return true
	}

	return false
}

// SetIdpUsernameAttribute gets a reference to the given string and assigns it to the IdpUsernameAttribute field.
func (o *SamlConfiguration) SetIdpUsernameAttribute(v string) {
	o.IdpUsernameAttribute = &v
}

// GetIdpAccountAttribute returns the IdpAccountAttribute field value if set, zero value otherwise.
func (o *SamlConfiguration) GetIdpAccountAttribute() string {
	if o == nil || o.IdpAccountAttribute == nil {
		var ret string
		return ret
	}
	return *o.IdpAccountAttribute
}

// GetIdpAccountAttributeOk returns a tuple with the IdpAccountAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetIdpAccountAttributeOk() (*string, bool) {
	if o == nil || o.IdpAccountAttribute == nil {
		return nil, false
	}
	return o.IdpAccountAttribute, true
}

// HasIdpAccountAttribute returns a boolean if a field has been set.
func (o *SamlConfiguration) HasIdpAccountAttribute() bool {
	if o != nil && o.IdpAccountAttribute != nil {
		return true
	}

	return false
}

// SetIdpAccountAttribute gets a reference to the given string and assigns it to the IdpAccountAttribute field.
func (o *SamlConfiguration) SetIdpAccountAttribute(v string) {
	o.IdpAccountAttribute = &v
}

// GetIdpRoleAttribute returns the IdpRoleAttribute field value if set, zero value otherwise.
func (o *SamlConfiguration) GetIdpRoleAttribute() string {
	if o == nil || o.IdpRoleAttribute == nil {
		var ret string
		return ret
	}
	return *o.IdpRoleAttribute
}

// GetIdpRoleAttributeOk returns a tuple with the IdpRoleAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetIdpRoleAttributeOk() (*string, bool) {
	if o == nil || o.IdpRoleAttribute == nil {
		return nil, false
	}
	return o.IdpRoleAttribute, true
}

// HasIdpRoleAttribute returns a boolean if a field has been set.
func (o *SamlConfiguration) HasIdpRoleAttribute() bool {
	if o != nil && o.IdpRoleAttribute != nil {
		return true
	}

	return false
}

// SetIdpRoleAttribute gets a reference to the given string and assigns it to the IdpRoleAttribute field.
func (o *SamlConfiguration) SetIdpRoleAttribute(v string) {
	o.IdpRoleAttribute = &v
}

// GetDefaultAccount returns the DefaultAccount field value if set, zero value otherwise.
func (o *SamlConfiguration) GetDefaultAccount() string {
	if o == nil || o.DefaultAccount == nil {
		var ret string
		return ret
	}
	return *o.DefaultAccount
}

// GetDefaultAccountOk returns a tuple with the DefaultAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetDefaultAccountOk() (*string, bool) {
	if o == nil || o.DefaultAccount == nil {
		return nil, false
	}
	return o.DefaultAccount, true
}

// HasDefaultAccount returns a boolean if a field has been set.
func (o *SamlConfiguration) HasDefaultAccount() bool {
	if o != nil && o.DefaultAccount != nil {
		return true
	}

	return false
}

// SetDefaultAccount gets a reference to the given string and assigns it to the DefaultAccount field.
func (o *SamlConfiguration) SetDefaultAccount(v string) {
	o.DefaultAccount = &v
}

// GetDefaultRole returns the DefaultRole field value if set, zero value otherwise.
func (o *SamlConfiguration) GetDefaultRole() string {
	if o == nil || o.DefaultRole == nil {
		var ret string
		return ret
	}
	return *o.DefaultRole
}

// GetDefaultRoleOk returns a tuple with the DefaultRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetDefaultRoleOk() (*string, bool) {
	if o == nil || o.DefaultRole == nil {
		return nil, false
	}
	return o.DefaultRole, true
}

// HasDefaultRole returns a boolean if a field has been set.
func (o *SamlConfiguration) HasDefaultRole() bool {
	if o != nil && o.DefaultRole != nil {
		return true
	}

	return false
}

// SetDefaultRole gets a reference to the given string and assigns it to the DefaultRole field.
func (o *SamlConfiguration) SetDefaultRole(v string) {
	o.DefaultRole = &v
}

// GetRequireSignedAssertions returns the RequireSignedAssertions field value if set, zero value otherwise.
func (o *SamlConfiguration) GetRequireSignedAssertions() bool {
	if o == nil || o.RequireSignedAssertions == nil {
		var ret bool
		return ret
	}
	return *o.RequireSignedAssertions
}

// GetRequireSignedAssertionsOk returns a tuple with the RequireSignedAssertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetRequireSignedAssertionsOk() (*bool, bool) {
	if o == nil || o.RequireSignedAssertions == nil {
		return nil, false
	}
	return o.RequireSignedAssertions, true
}

// HasRequireSignedAssertions returns a boolean if a field has been set.
func (o *SamlConfiguration) HasRequireSignedAssertions() bool {
	if o != nil && o.RequireSignedAssertions != nil {
		return true
	}

	return false
}

// SetRequireSignedAssertions gets a reference to the given bool and assigns it to the RequireSignedAssertions field.
func (o *SamlConfiguration) SetRequireSignedAssertions(v bool) {
	o.RequireSignedAssertions = &v
}

// GetRequireSignedResponse returns the RequireSignedResponse field value if set, zero value otherwise.
func (o *SamlConfiguration) GetRequireSignedResponse() bool {
	if o == nil || o.RequireSignedResponse == nil {
		var ret bool
		return ret
	}
	return *o.RequireSignedResponse
}

// GetRequireSignedResponseOk returns a tuple with the RequireSignedResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetRequireSignedResponseOk() (*bool, bool) {
	if o == nil || o.RequireSignedResponse == nil {
		return nil, false
	}
	return o.RequireSignedResponse, true
}

// HasRequireSignedResponse returns a boolean if a field has been set.
func (o *SamlConfiguration) HasRequireSignedResponse() bool {
	if o != nil && o.RequireSignedResponse != nil {
		return true
	}

	return false
}

// SetRequireSignedResponse gets a reference to the given bool and assigns it to the RequireSignedResponse field.
func (o *SamlConfiguration) SetRequireSignedResponse(v bool) {
	o.RequireSignedResponse = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SamlConfiguration) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SamlConfiguration) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SamlConfiguration) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *SamlConfiguration) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfiguration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *SamlConfiguration) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *SamlConfiguration) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o SamlConfiguration) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableSamlConfiguration struct {
	value *SamlConfiguration
	isSet bool
}

func (v NullableSamlConfiguration) Get() *SamlConfiguration {
	return v.value
}

func (v *NullableSamlConfiguration) Set(val *SamlConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlConfiguration(val *SamlConfiguration) *NullableSamlConfiguration {
	return &NullableSamlConfiguration{value: val, isSet: true}
}

func (v NullableSamlConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

