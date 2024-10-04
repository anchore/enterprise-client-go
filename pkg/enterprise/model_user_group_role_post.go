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
	"bytes"
	"fmt"
)

// checks if the UserGroupRolePost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupRolePost{}

// UserGroupRolePost struct for UserGroupRolePost
type UserGroupRolePost struct {
	// Deprecated. Please use domain_name instead. The account
	// Deprecated
	ForAccount *string `json:"for_account,omitempty"`
	// The domain scope for this role. This may be an account name when the domain is an account.
	DomainName *string `json:"domain_name,omitempty"`
	Roles []UserGroupRolePostRolesInner `json:"roles"`
}

type _UserGroupRolePost UserGroupRolePost

// NewUserGroupRolePost instantiates a new UserGroupRolePost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupRolePost(roles []UserGroupRolePostRolesInner) *UserGroupRolePost {
	this := UserGroupRolePost{}
	this.Roles = roles
	return &this
}

// NewUserGroupRolePostWithDefaults instantiates a new UserGroupRolePost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupRolePostWithDefaults() *UserGroupRolePost {
	this := UserGroupRolePost{}
	return &this
}

// GetForAccount returns the ForAccount field value if set, zero value otherwise.
// Deprecated
func (o *UserGroupRolePost) GetForAccount() string {
	if o == nil || IsNil(o.ForAccount) {
		var ret string
		return ret
	}
	return *o.ForAccount
}

// GetForAccountOk returns a tuple with the ForAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UserGroupRolePost) GetForAccountOk() (*string, bool) {
	if o == nil || IsNil(o.ForAccount) {
		return nil, false
	}
	return o.ForAccount, true
}

// HasForAccount returns a boolean if a field has been set.
func (o *UserGroupRolePost) HasForAccount() bool {
	if o != nil && !IsNil(o.ForAccount) {
		return true
	}

	return false
}

// SetForAccount gets a reference to the given string and assigns it to the ForAccount field.
// Deprecated
func (o *UserGroupRolePost) SetForAccount(v string) {
	o.ForAccount = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *UserGroupRolePost) GetDomainName() string {
	if o == nil || IsNil(o.DomainName) {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupRolePost) GetDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *UserGroupRolePost) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *UserGroupRolePost) SetDomainName(v string) {
	o.DomainName = &v
}

// GetRoles returns the Roles field value
func (o *UserGroupRolePost) GetRoles() []UserGroupRolePostRolesInner {
	if o == nil {
		var ret []UserGroupRolePostRolesInner
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UserGroupRolePost) GetRolesOk() ([]UserGroupRolePostRolesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *UserGroupRolePost) SetRoles(v []UserGroupRolePostRolesInner) {
	o.Roles = v
}

func (o UserGroupRolePost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupRolePost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForAccount) {
		toSerialize["for_account"] = o.ForAccount
	}
	if !IsNil(o.DomainName) {
		toSerialize["domain_name"] = o.DomainName
	}
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

func (o *UserGroupRolePost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roles",
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

	varUserGroupRolePost := _UserGroupRolePost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGroupRolePost)

	if err != nil {
		return err
	}

	*o = UserGroupRolePost(varUserGroupRolePost)

	return err
}

type NullableUserGroupRolePost struct {
	value *UserGroupRolePost
	isSet bool
}

func (v NullableUserGroupRolePost) Get() *UserGroupRolePost {
	return v.value
}

func (v *NullableUserGroupRolePost) Set(val *UserGroupRolePost) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupRolePost) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupRolePost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupRolePost(val *UserGroupRolePost) *NullableUserGroupRolePost {
	return &NullableUserGroupRolePost{value: val, isSet: true}
}

func (v NullableUserGroupRolePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupRolePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


