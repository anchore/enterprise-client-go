/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 0.1.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// RegistryConfiguration A registry entry describing the endpoint and credentials for a registry to pull images from
type RegistryConfiguration struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Username portion of credential to use for this registry
	RegistryUser *string `json:"registry_user,omitempty"`
	// Type of registry
	RegistryType *string `json:"registry_type,omitempty"`
	// Engine user that owns this registry entry
	AccountName *string `json:"account_name,omitempty"`
	// hostname:port string for accessing the registry, as would be used in a docker pull operation
	Registry *string `json:"registry,omitempty"`
	// human readable name associated with registry record
	RegistryName *string `json:"registry_name,omitempty"`
	// Use TLS/SSL verification for the registry URL
	RegistryVerify *bool `json:"registry_verify,omitempty"`
}

// NewRegistryConfiguration instantiates a new RegistryConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryConfiguration() *RegistryConfiguration {
	this := RegistryConfiguration{}
	return &this
}

// NewRegistryConfigurationWithDefaults instantiates a new RegistryConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryConfigurationWithDefaults() *RegistryConfiguration {
	this := RegistryConfiguration{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RegistryConfiguration) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfiguration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RegistryConfiguration) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RegistryConfiguration) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *RegistryConfiguration) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfiguration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *RegistryConfiguration) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *RegistryConfiguration) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetRegistryUser returns the RegistryUser field value if set, zero value otherwise.
func (o *RegistryConfiguration) GetRegistryUser() string {
	if o == nil || o.RegistryUser == nil {
		var ret string
		return ret
	}
	return *o.RegistryUser
}

// GetRegistryUserOk returns a tuple with the RegistryUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfiguration) GetRegistryUserOk() (*string, bool) {
	if o == nil || o.RegistryUser == nil {
		return nil, false
	}
	return o.RegistryUser, true
}

// HasRegistryUser returns a boolean if a field has been set.
func (o *RegistryConfiguration) HasRegistryUser() bool {
	if o != nil && o.RegistryUser != nil {
		return true
	}

	return false
}

// SetRegistryUser gets a reference to the given string and assigns it to the RegistryUser field.
func (o *RegistryConfiguration) SetRegistryUser(v string) {
	o.RegistryUser = &v
}

// GetRegistryType returns the RegistryType field value if set, zero value otherwise.
func (o *RegistryConfiguration) GetRegistryType() string {
	if o == nil || o.RegistryType == nil {
		var ret string
		return ret
	}
	return *o.RegistryType
}

// GetRegistryTypeOk returns a tuple with the RegistryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfiguration) GetRegistryTypeOk() (*string, bool) {
	if o == nil || o.RegistryType == nil {
		return nil, false
	}
	return o.RegistryType, true
}

// HasRegistryType returns a boolean if a field has been set.
func (o *RegistryConfiguration) HasRegistryType() bool {
	if o != nil && o.RegistryType != nil {
		return true
	}

	return false
}

// SetRegistryType gets a reference to the given string and assigns it to the RegistryType field.
func (o *RegistryConfiguration) SetRegistryType(v string) {
	o.RegistryType = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *RegistryConfiguration) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfiguration) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *RegistryConfiguration) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *RegistryConfiguration) SetAccountName(v string) {
	o.AccountName = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *RegistryConfiguration) GetRegistry() string {
	if o == nil || o.Registry == nil {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfiguration) GetRegistryOk() (*string, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *RegistryConfiguration) HasRegistry() bool {
	if o != nil && o.Registry != nil {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *RegistryConfiguration) SetRegistry(v string) {
	o.Registry = &v
}

// GetRegistryName returns the RegistryName field value if set, zero value otherwise.
func (o *RegistryConfiguration) GetRegistryName() string {
	if o == nil || o.RegistryName == nil {
		var ret string
		return ret
	}
	return *o.RegistryName
}

// GetRegistryNameOk returns a tuple with the RegistryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfiguration) GetRegistryNameOk() (*string, bool) {
	if o == nil || o.RegistryName == nil {
		return nil, false
	}
	return o.RegistryName, true
}

// HasRegistryName returns a boolean if a field has been set.
func (o *RegistryConfiguration) HasRegistryName() bool {
	if o != nil && o.RegistryName != nil {
		return true
	}

	return false
}

// SetRegistryName gets a reference to the given string and assigns it to the RegistryName field.
func (o *RegistryConfiguration) SetRegistryName(v string) {
	o.RegistryName = &v
}

// GetRegistryVerify returns the RegistryVerify field value if set, zero value otherwise.
func (o *RegistryConfiguration) GetRegistryVerify() bool {
	if o == nil || o.RegistryVerify == nil {
		var ret bool
		return ret
	}
	return *o.RegistryVerify
}

// GetRegistryVerifyOk returns a tuple with the RegistryVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfiguration) GetRegistryVerifyOk() (*bool, bool) {
	if o == nil || o.RegistryVerify == nil {
		return nil, false
	}
	return o.RegistryVerify, true
}

// HasRegistryVerify returns a boolean if a field has been set.
func (o *RegistryConfiguration) HasRegistryVerify() bool {
	if o != nil && o.RegistryVerify != nil {
		return true
	}

	return false
}

// SetRegistryVerify gets a reference to the given bool and assigns it to the RegistryVerify field.
func (o *RegistryConfiguration) SetRegistryVerify(v bool) {
	o.RegistryVerify = &v
}

func (o RegistryConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.RegistryUser != nil {
		toSerialize["registry_user"] = o.RegistryUser
	}
	if o.RegistryType != nil {
		toSerialize["registry_type"] = o.RegistryType
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.Registry != nil {
		toSerialize["registry"] = o.Registry
	}
	if o.RegistryName != nil {
		toSerialize["registry_name"] = o.RegistryName
	}
	if o.RegistryVerify != nil {
		toSerialize["registry_verify"] = o.RegistryVerify
	}
	return json.Marshal(toSerialize)
}

type NullableRegistryConfiguration struct {
	value *RegistryConfiguration
	isSet bool
}

func (v NullableRegistryConfiguration) Get() *RegistryConfiguration {
	return v.value
}

func (v *NullableRegistryConfiguration) Set(val *RegistryConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryConfiguration(val *RegistryConfiguration) *NullableRegistryConfiguration {
	return &NullableRegistryConfiguration{value: val, isSet: true}
}

func (v NullableRegistryConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


