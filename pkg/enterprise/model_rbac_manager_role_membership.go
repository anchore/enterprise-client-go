/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// RbacManagerRoleMembership Membership for a role in an account
type RbacManagerRoleMembership struct {
	// The name of the role the user has permissions for
	Role *string `json:"role,omitempty"`
	// The account for which the user has the role permission
	ForAccount *string `json:"for_account,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewRbacManagerRoleMembership instantiates a new RbacManagerRoleMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerRoleMembership() *RbacManagerRoleMembership {
	this := RbacManagerRoleMembership{}
	return &this
}

// NewRbacManagerRoleMembershipWithDefaults instantiates a new RbacManagerRoleMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerRoleMembershipWithDefaults() *RbacManagerRoleMembership {
	this := RbacManagerRoleMembership{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RbacManagerRoleMembership) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleMembership) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RbacManagerRoleMembership) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *RbacManagerRoleMembership) SetRole(v string) {
	o.Role = &v
}

// GetForAccount returns the ForAccount field value if set, zero value otherwise.
func (o *RbacManagerRoleMembership) GetForAccount() string {
	if o == nil || o.ForAccount == nil {
		var ret string
		return ret
	}
	return *o.ForAccount
}

// GetForAccountOk returns a tuple with the ForAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleMembership) GetForAccountOk() (*string, bool) {
	if o == nil || o.ForAccount == nil {
		return nil, false
	}
	return o.ForAccount, true
}

// HasForAccount returns a boolean if a field has been set.
func (o *RbacManagerRoleMembership) HasForAccount() bool {
	if o != nil && o.ForAccount != nil {
		return true
	}

	return false
}

// SetForAccount gets a reference to the given string and assigns it to the ForAccount field.
func (o *RbacManagerRoleMembership) SetForAccount(v string) {
	o.ForAccount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RbacManagerRoleMembership) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleMembership) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RbacManagerRoleMembership) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RbacManagerRoleMembership) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o RbacManagerRoleMembership) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.ForAccount != nil {
		toSerialize["for_account"] = o.ForAccount
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableRbacManagerRoleMembership struct {
	value *RbacManagerRoleMembership
	isSet bool
}

func (v NullableRbacManagerRoleMembership) Get() *RbacManagerRoleMembership {
	return v.value
}

func (v *NullableRbacManagerRoleMembership) Set(val *RbacManagerRoleMembership) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerRoleMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerRoleMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerRoleMembership(val *RbacManagerRoleMembership) *NullableRbacManagerRoleMembership {
	return &NullableRbacManagerRoleMembership{value: val, isSet: true}
}

func (v NullableRbacManagerRoleMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerRoleMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


