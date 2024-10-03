/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the RbacManagerSamlConfigurationGet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RbacManagerSamlConfigurationGet{}

// RbacManagerSamlConfigurationGet struct for RbacManagerSamlConfigurationGet
type RbacManagerSamlConfigurationGet struct {
	// The name to use for referencing this IDP configuration. This will configured as part of the url string the Idp must have the client POST the saml assertion to.
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9_-]+$"`
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
	// The default value is `memberOf`. This field is available to overwrite the SAML attribute if your IDP is using a different SSO group value during the response assertions.
	IdpGroupsAttribute *string `json:"idp_groups_attribute,omitempty"`
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
	// List of user groups associated with this IDP (Only for GET operations)
	UserGroups []RbacManagerSamlConfigurationGetAllOfUserGroups `json:"user_groups,omitempty"`
}

type _RbacManagerSamlConfigurationGet RbacManagerSamlConfigurationGet

// NewRbacManagerSamlConfigurationGet instantiates a new RbacManagerSamlConfigurationGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerSamlConfigurationGet(name string, enabled bool, spEntityId string, acsUrl string) *RbacManagerSamlConfigurationGet {
	this := RbacManagerSamlConfigurationGet{}
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

// NewRbacManagerSamlConfigurationGetWithDefaults instantiates a new RbacManagerSamlConfigurationGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerSamlConfigurationGetWithDefaults() *RbacManagerSamlConfigurationGet {
	this := RbacManagerSamlConfigurationGet{}
	var requireSignedAssertions bool = true
	this.RequireSignedAssertions = &requireSignedAssertions
	var requireSignedResponse bool = true
	this.RequireSignedResponse = &requireSignedResponse
	var requireExistingUsers bool = false
	this.RequireExistingUsers = &requireExistingUsers
	return &this
}

// GetName returns the Name field value
func (o *RbacManagerSamlConfigurationGet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RbacManagerSamlConfigurationGet) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *RbacManagerSamlConfigurationGet) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *RbacManagerSamlConfigurationGet) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSpEntityId returns the SpEntityId field value
func (o *RbacManagerSamlConfigurationGet) GetSpEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpEntityId
}

// GetSpEntityIdOk returns a tuple with the SpEntityId field value
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetSpEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpEntityId, true
}

// SetSpEntityId sets field value
func (o *RbacManagerSamlConfigurationGet) SetSpEntityId(v string) {
	o.SpEntityId = v
}

// GetAcsUrl returns the AcsUrl field value
func (o *RbacManagerSamlConfigurationGet) GetAcsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetAcsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcsUrl, true
}

// SetAcsUrl sets field value
func (o *RbacManagerSamlConfigurationGet) SetAcsUrl(v string) {
	o.AcsUrl = v
}

// GetAcsHttpsPort returns the AcsHttpsPort field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetAcsHttpsPort() int32 {
	if o == nil || IsNil(o.AcsHttpsPort) {
		var ret int32
		return ret
	}
	return *o.AcsHttpsPort
}

// GetAcsHttpsPortOk returns a tuple with the AcsHttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetAcsHttpsPortOk() (*int32, bool) {
	if o == nil || IsNil(o.AcsHttpsPort) {
		return nil, false
	}
	return o.AcsHttpsPort, true
}

// HasAcsHttpsPort returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasAcsHttpsPort() bool {
	if o != nil && !IsNil(o.AcsHttpsPort) {
		return true
	}

	return false
}

// SetAcsHttpsPort gets a reference to the given int32 and assigns it to the AcsHttpsPort field.
func (o *RbacManagerSamlConfigurationGet) SetAcsHttpsPort(v int32) {
	o.AcsHttpsPort = &v
}

// GetIdpMetadataUrl returns the IdpMetadataUrl field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetIdpMetadataUrl() string {
	if o == nil || IsNil(o.IdpMetadataUrl) {
		var ret string
		return ret
	}
	return *o.IdpMetadataUrl
}

// GetIdpMetadataUrlOk returns a tuple with the IdpMetadataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetIdpMetadataUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IdpMetadataUrl) {
		return nil, false
	}
	return o.IdpMetadataUrl, true
}

// HasIdpMetadataUrl returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasIdpMetadataUrl() bool {
	if o != nil && !IsNil(o.IdpMetadataUrl) {
		return true
	}

	return false
}

// SetIdpMetadataUrl gets a reference to the given string and assigns it to the IdpMetadataUrl field.
func (o *RbacManagerSamlConfigurationGet) SetIdpMetadataUrl(v string) {
	o.IdpMetadataUrl = &v
}

// GetIdpMetadataXml returns the IdpMetadataXml field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetIdpMetadataXml() string {
	if o == nil || IsNil(o.IdpMetadataXml) {
		var ret string
		return ret
	}
	return *o.IdpMetadataXml
}

// GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetIdpMetadataXmlOk() (*string, bool) {
	if o == nil || IsNil(o.IdpMetadataXml) {
		return nil, false
	}
	return o.IdpMetadataXml, true
}

// HasIdpMetadataXml returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasIdpMetadataXml() bool {
	if o != nil && !IsNil(o.IdpMetadataXml) {
		return true
	}

	return false
}

// SetIdpMetadataXml gets a reference to the given string and assigns it to the IdpMetadataXml field.
func (o *RbacManagerSamlConfigurationGet) SetIdpMetadataXml(v string) {
	o.IdpMetadataXml = &v
}

// GetIdpUsernameAttribute returns the IdpUsernameAttribute field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetIdpUsernameAttribute() string {
	if o == nil || IsNil(o.IdpUsernameAttribute) {
		var ret string
		return ret
	}
	return *o.IdpUsernameAttribute
}

// GetIdpUsernameAttributeOk returns a tuple with the IdpUsernameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetIdpUsernameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.IdpUsernameAttribute) {
		return nil, false
	}
	return o.IdpUsernameAttribute, true
}

// HasIdpUsernameAttribute returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasIdpUsernameAttribute() bool {
	if o != nil && !IsNil(o.IdpUsernameAttribute) {
		return true
	}

	return false
}

// SetIdpUsernameAttribute gets a reference to the given string and assigns it to the IdpUsernameAttribute field.
func (o *RbacManagerSamlConfigurationGet) SetIdpUsernameAttribute(v string) {
	o.IdpUsernameAttribute = &v
}

// GetIdpAccountAttribute returns the IdpAccountAttribute field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetIdpAccountAttribute() string {
	if o == nil || IsNil(o.IdpAccountAttribute) {
		var ret string
		return ret
	}
	return *o.IdpAccountAttribute
}

// GetIdpAccountAttributeOk returns a tuple with the IdpAccountAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetIdpAccountAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.IdpAccountAttribute) {
		return nil, false
	}
	return o.IdpAccountAttribute, true
}

// HasIdpAccountAttribute returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasIdpAccountAttribute() bool {
	if o != nil && !IsNil(o.IdpAccountAttribute) {
		return true
	}

	return false
}

// SetIdpAccountAttribute gets a reference to the given string and assigns it to the IdpAccountAttribute field.
func (o *RbacManagerSamlConfigurationGet) SetIdpAccountAttribute(v string) {
	o.IdpAccountAttribute = &v
}

// GetIdpRoleAttribute returns the IdpRoleAttribute field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetIdpRoleAttribute() string {
	if o == nil || IsNil(o.IdpRoleAttribute) {
		var ret string
		return ret
	}
	return *o.IdpRoleAttribute
}

// GetIdpRoleAttributeOk returns a tuple with the IdpRoleAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetIdpRoleAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.IdpRoleAttribute) {
		return nil, false
	}
	return o.IdpRoleAttribute, true
}

// HasIdpRoleAttribute returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasIdpRoleAttribute() bool {
	if o != nil && !IsNil(o.IdpRoleAttribute) {
		return true
	}

	return false
}

// SetIdpRoleAttribute gets a reference to the given string and assigns it to the IdpRoleAttribute field.
func (o *RbacManagerSamlConfigurationGet) SetIdpRoleAttribute(v string) {
	o.IdpRoleAttribute = &v
}

// GetIdpGroupsAttribute returns the IdpGroupsAttribute field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetIdpGroupsAttribute() string {
	if o == nil || IsNil(o.IdpGroupsAttribute) {
		var ret string
		return ret
	}
	return *o.IdpGroupsAttribute
}

// GetIdpGroupsAttributeOk returns a tuple with the IdpGroupsAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetIdpGroupsAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.IdpGroupsAttribute) {
		return nil, false
	}
	return o.IdpGroupsAttribute, true
}

// HasIdpGroupsAttribute returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasIdpGroupsAttribute() bool {
	if o != nil && !IsNil(o.IdpGroupsAttribute) {
		return true
	}

	return false
}

// SetIdpGroupsAttribute gets a reference to the given string and assigns it to the IdpGroupsAttribute field.
func (o *RbacManagerSamlConfigurationGet) SetIdpGroupsAttribute(v string) {
	o.IdpGroupsAttribute = &v
}

// GetDefaultAccount returns the DefaultAccount field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetDefaultAccount() string {
	if o == nil || IsNil(o.DefaultAccount) {
		var ret string
		return ret
	}
	return *o.DefaultAccount
}

// GetDefaultAccountOk returns a tuple with the DefaultAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetDefaultAccountOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultAccount) {
		return nil, false
	}
	return o.DefaultAccount, true
}

// HasDefaultAccount returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasDefaultAccount() bool {
	if o != nil && !IsNil(o.DefaultAccount) {
		return true
	}

	return false
}

// SetDefaultAccount gets a reference to the given string and assigns it to the DefaultAccount field.
func (o *RbacManagerSamlConfigurationGet) SetDefaultAccount(v string) {
	o.DefaultAccount = &v
}

// GetDefaultRole returns the DefaultRole field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetDefaultRole() string {
	if o == nil || IsNil(o.DefaultRole) {
		var ret string
		return ret
	}
	return *o.DefaultRole
}

// GetDefaultRoleOk returns a tuple with the DefaultRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetDefaultRoleOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultRole) {
		return nil, false
	}
	return o.DefaultRole, true
}

// HasDefaultRole returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasDefaultRole() bool {
	if o != nil && !IsNil(o.DefaultRole) {
		return true
	}

	return false
}

// SetDefaultRole gets a reference to the given string and assigns it to the DefaultRole field.
func (o *RbacManagerSamlConfigurationGet) SetDefaultRole(v string) {
	o.DefaultRole = &v
}

// GetRequireSignedAssertions returns the RequireSignedAssertions field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetRequireSignedAssertions() bool {
	if o == nil || IsNil(o.RequireSignedAssertions) {
		var ret bool
		return ret
	}
	return *o.RequireSignedAssertions
}

// GetRequireSignedAssertionsOk returns a tuple with the RequireSignedAssertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetRequireSignedAssertionsOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSignedAssertions) {
		return nil, false
	}
	return o.RequireSignedAssertions, true
}

// HasRequireSignedAssertions returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasRequireSignedAssertions() bool {
	if o != nil && !IsNil(o.RequireSignedAssertions) {
		return true
	}

	return false
}

// SetRequireSignedAssertions gets a reference to the given bool and assigns it to the RequireSignedAssertions field.
func (o *RbacManagerSamlConfigurationGet) SetRequireSignedAssertions(v bool) {
	o.RequireSignedAssertions = &v
}

// GetRequireSignedResponse returns the RequireSignedResponse field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetRequireSignedResponse() bool {
	if o == nil || IsNil(o.RequireSignedResponse) {
		var ret bool
		return ret
	}
	return *o.RequireSignedResponse
}

// GetRequireSignedResponseOk returns a tuple with the RequireSignedResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetRequireSignedResponseOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSignedResponse) {
		return nil, false
	}
	return o.RequireSignedResponse, true
}

// HasRequireSignedResponse returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasRequireSignedResponse() bool {
	if o != nil && !IsNil(o.RequireSignedResponse) {
		return true
	}

	return false
}

// SetRequireSignedResponse gets a reference to the given bool and assigns it to the RequireSignedResponse field.
func (o *RbacManagerSamlConfigurationGet) SetRequireSignedResponse(v bool) {
	o.RequireSignedResponse = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RbacManagerSamlConfigurationGet) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *RbacManagerSamlConfigurationGet) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetRequireExistingUsers returns the RequireExistingUsers field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetRequireExistingUsers() bool {
	if o == nil || IsNil(o.RequireExistingUsers) {
		var ret bool
		return ret
	}
	return *o.RequireExistingUsers
}

// GetRequireExistingUsersOk returns a tuple with the RequireExistingUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetRequireExistingUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireExistingUsers) {
		return nil, false
	}
	return o.RequireExistingUsers, true
}

// HasRequireExistingUsers returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasRequireExistingUsers() bool {
	if o != nil && !IsNil(o.RequireExistingUsers) {
		return true
	}

	return false
}

// SetRequireExistingUsers gets a reference to the given bool and assigns it to the RequireExistingUsers field.
func (o *RbacManagerSamlConfigurationGet) SetRequireExistingUsers(v bool) {
	o.RequireExistingUsers = &v
}

// GetUserGroups returns the UserGroups field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGet) GetUserGroups() []RbacManagerSamlConfigurationGetAllOfUserGroups {
	if o == nil || IsNil(o.UserGroups) {
		var ret []RbacManagerSamlConfigurationGetAllOfUserGroups
		return ret
	}
	return o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGet) GetUserGroupsOk() ([]RbacManagerSamlConfigurationGetAllOfUserGroups, bool) {
	if o == nil || IsNil(o.UserGroups) {
		return nil, false
	}
	return o.UserGroups, true
}

// HasUserGroups returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGet) HasUserGroups() bool {
	if o != nil && !IsNil(o.UserGroups) {
		return true
	}

	return false
}

// SetUserGroups gets a reference to the given []RbacManagerSamlConfigurationGetAllOfUserGroups and assigns it to the UserGroups field.
func (o *RbacManagerSamlConfigurationGet) SetUserGroups(v []RbacManagerSamlConfigurationGetAllOfUserGroups) {
	o.UserGroups = v
}

func (o RbacManagerSamlConfigurationGet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RbacManagerSamlConfigurationGet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["enabled"] = o.Enabled
	toSerialize["sp_entity_id"] = o.SpEntityId
	toSerialize["acs_url"] = o.AcsUrl
	if !IsNil(o.AcsHttpsPort) {
		toSerialize["acs_https_port"] = o.AcsHttpsPort
	}
	if !IsNil(o.IdpMetadataUrl) {
		toSerialize["idp_metadata_url"] = o.IdpMetadataUrl
	}
	if !IsNil(o.IdpMetadataXml) {
		toSerialize["idp_metadata_xml"] = o.IdpMetadataXml
	}
	if !IsNil(o.IdpUsernameAttribute) {
		toSerialize["idp_username_attribute"] = o.IdpUsernameAttribute
	}
	if !IsNil(o.IdpAccountAttribute) {
		toSerialize["idp_account_attribute"] = o.IdpAccountAttribute
	}
	if !IsNil(o.IdpRoleAttribute) {
		toSerialize["idp_role_attribute"] = o.IdpRoleAttribute
	}
	if !IsNil(o.IdpGroupsAttribute) {
		toSerialize["idp_groups_attribute"] = o.IdpGroupsAttribute
	}
	if !IsNil(o.DefaultAccount) {
		toSerialize["default_account"] = o.DefaultAccount
	}
	if !IsNil(o.DefaultRole) {
		toSerialize["default_role"] = o.DefaultRole
	}
	if !IsNil(o.RequireSignedAssertions) {
		toSerialize["require_signed_assertions"] = o.RequireSignedAssertions
	}
	if !IsNil(o.RequireSignedResponse) {
		toSerialize["require_signed_response"] = o.RequireSignedResponse
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if !IsNil(o.RequireExistingUsers) {
		toSerialize["require_existing_users"] = o.RequireExistingUsers
	}
	if !IsNil(o.UserGroups) {
		toSerialize["user_groups"] = o.UserGroups
	}
	return toSerialize, nil
}

func (o *RbacManagerSamlConfigurationGet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"enabled",
		"sp_entity_id",
		"acs_url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRbacManagerSamlConfigurationGet := _RbacManagerSamlConfigurationGet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRbacManagerSamlConfigurationGet)

	if err != nil {
		return err
	}

	*o = RbacManagerSamlConfigurationGet(varRbacManagerSamlConfigurationGet)

	return err
}

type NullableRbacManagerSamlConfigurationGet struct {
	value *RbacManagerSamlConfigurationGet
	isSet bool
}

func (v NullableRbacManagerSamlConfigurationGet) Get() *RbacManagerSamlConfigurationGet {
	return v.value
}

func (v *NullableRbacManagerSamlConfigurationGet) Set(val *RbacManagerSamlConfigurationGet) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerSamlConfigurationGet) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerSamlConfigurationGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerSamlConfigurationGet(val *RbacManagerSamlConfigurationGet) *NullableRbacManagerSamlConfigurationGet {
	return &NullableRbacManagerSamlConfigurationGet{value: val, isSet: true}
}

func (v NullableRbacManagerSamlConfigurationGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerSamlConfigurationGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


