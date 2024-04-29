/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// UnifiedRoles struct for UnifiedRoles
type UnifiedRoles struct {
	// The name of the RBAC Role
	RoleName *string `json:"role_name,omitempty"`
	// The domain (or account) name that provides the scope of the role
	DomainName *string `json:"domain_name,omitempty"`
}

// NewUnifiedRoles instantiates a new UnifiedRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnifiedRoles() *UnifiedRoles {
	this := UnifiedRoles{}
	return &this
}

// NewUnifiedRolesWithDefaults instantiates a new UnifiedRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnifiedRolesWithDefaults() *UnifiedRoles {
	this := UnifiedRoles{}
	return &this
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *UnifiedRoles) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoles) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *UnifiedRoles) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *UnifiedRoles) SetRoleName(v string) {
	o.RoleName = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *UnifiedRoles) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoles) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *UnifiedRoles) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *UnifiedRoles) SetDomainName(v string) {
	o.DomainName = &v
}

func (o UnifiedRoles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	if o.DomainName != nil {
		toSerialize["domain_name"] = o.DomainName
	}
	return json.Marshal(toSerialize)
}

type NullableUnifiedRoles struct {
	value *UnifiedRoles
	isSet bool
}

func (v NullableUnifiedRoles) Get() *UnifiedRoles {
	return v.value
}

func (v *NullableUnifiedRoles) Set(val *UnifiedRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableUnifiedRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableUnifiedRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnifiedRoles(val *UnifiedRoles) *NullableUnifiedRoles {
	return &NullableUnifiedRoles{value: val, isSet: true}
}

func (v NullableUnifiedRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnifiedRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


