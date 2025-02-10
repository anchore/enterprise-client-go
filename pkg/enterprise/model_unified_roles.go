/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the UnifiedRoles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnifiedRoles{}

// UnifiedRoles struct for UnifiedRoles
type UnifiedRoles struct {
	// The name of the RBAC Role
	RoleName *string `json:"role_name,omitempty"`
	// The domain (or account) name that provides the scope of the role
	DomainName *string `json:"domain_name,omitempty"`
	// The name of the user group that granted the role.  Will be null if the role was granted directly to the user.
	Granter *string `json:"granter,omitempty"`
	// The type of grant that was made
	GrantType *string `json:"grant_type,omitempty"`
	// The timestamp of when the role was granted
	GrantedAt *time.Time `json:"granted_at,omitempty"`
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
	if o == nil || IsNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoles) GetRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleName) {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *UnifiedRoles) HasRoleName() bool {
	if o != nil && !IsNil(o.RoleName) {
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
	if o == nil || IsNil(o.DomainName) {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoles) GetDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *UnifiedRoles) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *UnifiedRoles) SetDomainName(v string) {
	o.DomainName = &v
}

// GetGranter returns the Granter field value if set, zero value otherwise.
func (o *UnifiedRoles) GetGranter() string {
	if o == nil || IsNil(o.Granter) {
		var ret string
		return ret
	}
	return *o.Granter
}

// GetGranterOk returns a tuple with the Granter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoles) GetGranterOk() (*string, bool) {
	if o == nil || IsNil(o.Granter) {
		return nil, false
	}
	return o.Granter, true
}

// HasGranter returns a boolean if a field has been set.
func (o *UnifiedRoles) HasGranter() bool {
	if o != nil && !IsNil(o.Granter) {
		return true
	}

	return false
}

// SetGranter gets a reference to the given string and assigns it to the Granter field.
func (o *UnifiedRoles) SetGranter(v string) {
	o.Granter = &v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *UnifiedRoles) GetGrantType() string {
	if o == nil || IsNil(o.GrantType) {
		var ret string
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoles) GetGrantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GrantType) {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *UnifiedRoles) HasGrantType() bool {
	if o != nil && !IsNil(o.GrantType) {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given string and assigns it to the GrantType field.
func (o *UnifiedRoles) SetGrantType(v string) {
	o.GrantType = &v
}

// GetGrantedAt returns the GrantedAt field value if set, zero value otherwise.
func (o *UnifiedRoles) GetGrantedAt() time.Time {
	if o == nil || IsNil(o.GrantedAt) {
		var ret time.Time
		return ret
	}
	return *o.GrantedAt
}

// GetGrantedAtOk returns a tuple with the GrantedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoles) GetGrantedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.GrantedAt) {
		return nil, false
	}
	return o.GrantedAt, true
}

// HasGrantedAt returns a boolean if a field has been set.
func (o *UnifiedRoles) HasGrantedAt() bool {
	if o != nil && !IsNil(o.GrantedAt) {
		return true
	}

	return false
}

// SetGrantedAt gets a reference to the given time.Time and assigns it to the GrantedAt field.
func (o *UnifiedRoles) SetGrantedAt(v time.Time) {
	o.GrantedAt = &v
}

func (o UnifiedRoles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnifiedRoles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoleName) {
		toSerialize["role_name"] = o.RoleName
	}
	if !IsNil(o.DomainName) {
		toSerialize["domain_name"] = o.DomainName
	}
	if !IsNil(o.Granter) {
		toSerialize["granter"] = o.Granter
	}
	if !IsNil(o.GrantType) {
		toSerialize["grant_type"] = o.GrantType
	}
	if !IsNil(o.GrantedAt) {
		toSerialize["granted_at"] = o.GrantedAt
	}
	return toSerialize, nil
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


