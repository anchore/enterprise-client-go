/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.4.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// UserGroupRoleRolesInner struct for UserGroupRoleRolesInner
type UserGroupRoleRolesInner struct {
	// The role name
	Role *string `json:"role,omitempty"`
	// The timestamp of when the role membership was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The unique identifier for the role membership
	MembershipId *string `json:"membership_id,omitempty"`
}

// NewUserGroupRoleRolesInner instantiates a new UserGroupRoleRolesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupRoleRolesInner() *UserGroupRoleRolesInner {
	this := UserGroupRoleRolesInner{}
	return &this
}

// NewUserGroupRoleRolesInnerWithDefaults instantiates a new UserGroupRoleRolesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupRoleRolesInnerWithDefaults() *UserGroupRoleRolesInner {
	this := UserGroupRoleRolesInner{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UserGroupRoleRolesInner) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupRoleRolesInner) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UserGroupRoleRolesInner) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UserGroupRoleRolesInner) SetRole(v string) {
	o.Role = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserGroupRoleRolesInner) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupRoleRolesInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserGroupRoleRolesInner) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UserGroupRoleRolesInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetMembershipId returns the MembershipId field value if set, zero value otherwise.
func (o *UserGroupRoleRolesInner) GetMembershipId() string {
	if o == nil || o.MembershipId == nil {
		var ret string
		return ret
	}
	return *o.MembershipId
}

// GetMembershipIdOk returns a tuple with the MembershipId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupRoleRolesInner) GetMembershipIdOk() (*string, bool) {
	if o == nil || o.MembershipId == nil {
		return nil, false
	}
	return o.MembershipId, true
}

// HasMembershipId returns a boolean if a field has been set.
func (o *UserGroupRoleRolesInner) HasMembershipId() bool {
	if o != nil && o.MembershipId != nil {
		return true
	}

	return false
}

// SetMembershipId gets a reference to the given string and assigns it to the MembershipId field.
func (o *UserGroupRoleRolesInner) SetMembershipId(v string) {
	o.MembershipId = &v
}

func (o UserGroupRoleRolesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.MembershipId != nil {
		toSerialize["membership_id"] = o.MembershipId
	}
	return json.Marshal(toSerialize)
}

type NullableUserGroupRoleRolesInner struct {
	value *UserGroupRoleRolesInner
	isSet bool
}

func (v NullableUserGroupRoleRolesInner) Get() *UserGroupRoleRolesInner {
	return v.value
}

func (v *NullableUserGroupRoleRolesInner) Set(val *UserGroupRoleRolesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupRoleRolesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupRoleRolesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupRoleRolesInner(val *UserGroupRoleRolesInner) *NullableUserGroupRoleRolesInner {
	return &NullableUserGroupRoleRolesInner{value: val, isSet: true}
}

func (v NullableUserGroupRoleRolesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupRoleRolesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


