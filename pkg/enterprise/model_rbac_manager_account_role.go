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
)

// checks if the RbacManagerAccountRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RbacManagerAccountRole{}

// RbacManagerAccountRole An account identifier and roles a user has within that account
type RbacManagerAccountRole struct {
	// Deprecated. Please use `domain_name' instead.  The account scope that applies to the set of roles
	// Deprecated
	ForAccount *string `json:"for_account,omitempty"`
	// The domain scope that applies to the set of roles
	DomainName *string `json:"domain_name,omitempty"`
	Roles *RbacManagerRole `json:"roles,omitempty"`
	Account *Account `json:"account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RbacManagerAccountRole RbacManagerAccountRole

// NewRbacManagerAccountRole instantiates a new RbacManagerAccountRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerAccountRole() *RbacManagerAccountRole {
	this := RbacManagerAccountRole{}
	return &this
}

// NewRbacManagerAccountRoleWithDefaults instantiates a new RbacManagerAccountRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerAccountRoleWithDefaults() *RbacManagerAccountRole {
	this := RbacManagerAccountRole{}
	return &this
}

// GetForAccount returns the ForAccount field value if set, zero value otherwise.
// Deprecated
func (o *RbacManagerAccountRole) GetForAccount() string {
	if o == nil || IsNil(o.ForAccount) {
		var ret string
		return ret
	}
	return *o.ForAccount
}

// GetForAccountOk returns a tuple with the ForAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *RbacManagerAccountRole) GetForAccountOk() (*string, bool) {
	if o == nil || IsNil(o.ForAccount) {
		return nil, false
	}
	return o.ForAccount, true
}

// HasForAccount returns a boolean if a field has been set.
func (o *RbacManagerAccountRole) HasForAccount() bool {
	if o != nil && !IsNil(o.ForAccount) {
		return true
	}

	return false
}

// SetForAccount gets a reference to the given string and assigns it to the ForAccount field.
// Deprecated
func (o *RbacManagerAccountRole) SetForAccount(v string) {
	o.ForAccount = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *RbacManagerAccountRole) GetDomainName() string {
	if o == nil || IsNil(o.DomainName) {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerAccountRole) GetDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *RbacManagerAccountRole) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *RbacManagerAccountRole) SetDomainName(v string) {
	o.DomainName = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *RbacManagerAccountRole) GetRoles() RbacManagerRole {
	if o == nil || IsNil(o.Roles) {
		var ret RbacManagerRole
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerAccountRole) GetRolesOk() (*RbacManagerRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *RbacManagerAccountRole) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given RbacManagerRole and assigns it to the Roles field.
func (o *RbacManagerAccountRole) SetRoles(v RbacManagerRole) {
	o.Roles = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *RbacManagerAccountRole) GetAccount() Account {
	if o == nil || IsNil(o.Account) {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerAccountRole) GetAccountOk() (*Account, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *RbacManagerAccountRole) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *RbacManagerAccountRole) SetAccount(v Account) {
	o.Account = &v
}

func (o RbacManagerAccountRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RbacManagerAccountRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForAccount) {
		toSerialize["for_account"] = o.ForAccount
	}
	if !IsNil(o.DomainName) {
		toSerialize["domain_name"] = o.DomainName
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RbacManagerAccountRole) UnmarshalJSON(data []byte) (err error) {
	varRbacManagerAccountRole := _RbacManagerAccountRole{}

	err = json.Unmarshal(data, &varRbacManagerAccountRole)

	if err != nil {
		return err
	}

	*o = RbacManagerAccountRole(varRbacManagerAccountRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "for_account")
		delete(additionalProperties, "domain_name")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRbacManagerAccountRole struct {
	value *RbacManagerAccountRole
	isSet bool
}

func (v NullableRbacManagerAccountRole) Get() *RbacManagerAccountRole {
	return v.value
}

func (v *NullableRbacManagerAccountRole) Set(val *RbacManagerAccountRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerAccountRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerAccountRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerAccountRole(val *RbacManagerAccountRole) *NullableRbacManagerAccountRole {
	return &NullableRbacManagerAccountRole{value: val, isSet: true}
}

func (v NullableRbacManagerAccountRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerAccountRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


