/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// User A username for authenticating with one or more types of credentials. User type defines the expected credentials allowed for the user. Native users have passwords, other users have no credential internally. Internal users are service/system users for inter-service communication.
type User struct {
	// The username to authenticate with
	Username string `json:"username"`
	// The user's type
	Type *string `json:"type,omitempty"`
	// When the user 'type' is 'saml', this will be the EntityId of the IDP that they are authenticating from. Otherwise, this will be set to null.
	Source NullableString `json:"source,omitempty"`
	// The timestamp of when the user record was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The timestamp of the last update to this record
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// When the user 'type' is 'saml', this will be the configured name of the IDP that they are authenticating from.  Otherwise, this will be set to null.
	IdpName NullableString `json:"idp_name,omitempty"`
	// When the user 'type' is 'native', this will be the timestamp of the last time this user's credentials were updated.
	PasswordLastUpdated NullableTime `json:"password_last_updated,omitempty"`
	// The unified list of RBAC roles this user currently has in this account.
	UnifiedRoles []UnifiedRoles `json:"unified_roles,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(username string) *User {
	this := User{}
	this.Username = username
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetUsername returns the Username field value
func (o *User) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *User) SetUsername(v string) {
	o.Username = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *User) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *User) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *User) SetType(v string) {
	o.Type = &v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetSource() string {
	if o == nil || o.Source.Get() == nil {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *User) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableString and assigns it to the Source field.
func (o *User) SetSource(v string) {
	o.Source.Set(&v)
}
// SetSourceNil sets the value for Source to be an explicit nil
func (o *User) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *User) UnsetSource() {
	o.Source.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *User) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *User) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *User) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetIdpName returns the IdpName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetIdpName() string {
	if o == nil || o.IdpName.Get() == nil {
		var ret string
		return ret
	}
	return *o.IdpName.Get()
}

// GetIdpNameOk returns a tuple with the IdpName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIdpNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdpName.Get(), o.IdpName.IsSet()
}

// HasIdpName returns a boolean if a field has been set.
func (o *User) HasIdpName() bool {
	if o != nil && o.IdpName.IsSet() {
		return true
	}

	return false
}

// SetIdpName gets a reference to the given NullableString and assigns it to the IdpName field.
func (o *User) SetIdpName(v string) {
	o.IdpName.Set(&v)
}
// SetIdpNameNil sets the value for IdpName to be an explicit nil
func (o *User) SetIdpNameNil() {
	o.IdpName.Set(nil)
}

// UnsetIdpName ensures that no value is present for IdpName, not even an explicit nil
func (o *User) UnsetIdpName() {
	o.IdpName.Unset()
}

// GetPasswordLastUpdated returns the PasswordLastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetPasswordLastUpdated() time.Time {
	if o == nil || o.PasswordLastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.PasswordLastUpdated.Get()
}

// GetPasswordLastUpdatedOk returns a tuple with the PasswordLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPasswordLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordLastUpdated.Get(), o.PasswordLastUpdated.IsSet()
}

// HasPasswordLastUpdated returns a boolean if a field has been set.
func (o *User) HasPasswordLastUpdated() bool {
	if o != nil && o.PasswordLastUpdated.IsSet() {
		return true
	}

	return false
}

// SetPasswordLastUpdated gets a reference to the given NullableTime and assigns it to the PasswordLastUpdated field.
func (o *User) SetPasswordLastUpdated(v time.Time) {
	o.PasswordLastUpdated.Set(&v)
}
// SetPasswordLastUpdatedNil sets the value for PasswordLastUpdated to be an explicit nil
func (o *User) SetPasswordLastUpdatedNil() {
	o.PasswordLastUpdated.Set(nil)
}

// UnsetPasswordLastUpdated ensures that no value is present for PasswordLastUpdated, not even an explicit nil
func (o *User) UnsetPasswordLastUpdated() {
	o.PasswordLastUpdated.Unset()
}

// GetUnifiedRoles returns the UnifiedRoles field value if set, zero value otherwise.
func (o *User) GetUnifiedRoles() []UnifiedRoles {
	if o == nil || o.UnifiedRoles == nil {
		var ret []UnifiedRoles
		return ret
	}
	return o.UnifiedRoles
}

// GetUnifiedRolesOk returns a tuple with the UnifiedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUnifiedRolesOk() ([]UnifiedRoles, bool) {
	if o == nil || o.UnifiedRoles == nil {
		return nil, false
	}
	return o.UnifiedRoles, true
}

// HasUnifiedRoles returns a boolean if a field has been set.
func (o *User) HasUnifiedRoles() bool {
	if o != nil && o.UnifiedRoles != nil {
		return true
	}

	return false
}

// SetUnifiedRoles gets a reference to the given []UnifiedRoles and assigns it to the UnifiedRoles field.
func (o *User) SetUnifiedRoles(v []UnifiedRoles) {
	o.UnifiedRoles = v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.IdpName.IsSet() {
		toSerialize["idp_name"] = o.IdpName.Get()
	}
	if o.PasswordLastUpdated.IsSet() {
		toSerialize["password_last_updated"] = o.PasswordLastUpdated.Get()
	}
	if o.UnifiedRoles != nil {
		toSerialize["unified_roles"] = o.UnifiedRoles
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


