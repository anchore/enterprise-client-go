/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the RbacManagerRoleMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RbacManagerRoleMember{}

// RbacManagerRoleMember A mapping between a username and a role within a domain
type RbacManagerRoleMember struct {
	Username string
	// Deprecated. Please use domain_name instead.
	// Deprecated
	ForAccount *string
	// The domain scope that applies to the set of roles. This will be the account name if the domain scope is an account.
	DomainName *string
	CreatedAt *time.Time
	AdditionalProperties map[string]interface{}
}

type _RbacManagerRoleMember RbacManagerRoleMember

// NewRbacManagerRoleMember instantiates a new RbacManagerRoleMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerRoleMember(username string) *RbacManagerRoleMember {
	this := RbacManagerRoleMember{}
	return &this
}

// NewRbacManagerRoleMemberWithDefaults instantiates a new RbacManagerRoleMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerRoleMemberWithDefaults() *RbacManagerRoleMember {
	this := RbacManagerRoleMember{}
	return &this
}

// GetUsername returns the Username field value
func (o *RbacManagerRoleMember) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleMember) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *RbacManagerRoleMember) SetUsername(v string) {
	o.Username = v
}

// GetForAccount returns the ForAccount field value if set, zero value otherwise.
// Deprecated
func (o *RbacManagerRoleMember) GetForAccount() string {
	if o == nil || IsNil(o.ForAccount) {
		var ret string
		return ret
	}
	return *o.ForAccount
}

// GetForAccountOk returns a tuple with the ForAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *RbacManagerRoleMember) GetForAccountOk() (*string, bool) {
	if o == nil || IsNil(o.ForAccount) {
		return nil, false
	}
	return o.ForAccount, true
}

// HasForAccount returns a boolean if a field has been set.
func (o *RbacManagerRoleMember) HasForAccount() bool {
	if o != nil && !IsNil(o.ForAccount) {
		return true
	}

	return false
}

// SetForAccount gets a reference to the given string and assigns it to the ForAccount field.
// Deprecated
func (o *RbacManagerRoleMember) SetForAccount(v string) {
	o.ForAccount = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *RbacManagerRoleMember) GetDomainName() string {
	if o == nil || IsNil(o.DomainName) {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleMember) GetDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *RbacManagerRoleMember) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *RbacManagerRoleMember) SetDomainName(v string) {
	o.DomainName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RbacManagerRoleMember) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleMember) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RbacManagerRoleMember) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RbacManagerRoleMember) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o RbacManagerRoleMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RbacManagerRoleMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	if !IsNil(o.ForAccount) {
		toSerialize["for_account"] = o.ForAccount
	}
	if !IsNil(o.DomainName) {
		toSerialize["domain_name"] = o.DomainName
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RbacManagerRoleMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
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

	varRbacManagerRoleMember := _RbacManagerRoleMember{}

	err = json.Unmarshal(data, &varRbacManagerRoleMember)

	if err != nil {
		return err
	}

	*o = RbacManagerRoleMember(varRbacManagerRoleMember)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "for_account")
		delete(additionalProperties, "domain_name")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRbacManagerRoleMember struct {
	value *RbacManagerRoleMember
	isSet bool
}

func (v NullableRbacManagerRoleMember) Get() *RbacManagerRoleMember {
	return v.value
}

func (v *NullableRbacManagerRoleMember) Set(val *RbacManagerRoleMember) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerRoleMember) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerRoleMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerRoleMember(val *RbacManagerRoleMember) *NullableRbacManagerRoleMember {
	return &NullableRbacManagerRoleMember{value: val, isSet: true}
}

func (v NullableRbacManagerRoleMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerRoleMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


