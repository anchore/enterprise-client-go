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
)

// RegistryConfigurationRequest A registry record describing the endpoint and credentials for a registry
type RegistryConfigurationRequest struct {
	// Username portion of credential to use for this registry
	RegistryUser *string `json:"registry_user,omitempty"`
	// Password portion of credential to use for this registry
	RegistryPass *string `json:"registry_pass,omitempty"`
	// Type of registry
	RegistryType *string `json:"registry_type,omitempty"`
	// hostname:port string for accessing the registry, as would be used in a docker pull operation. May include some or all of a repository and wildcards (e.g. docker.io/library/_* or gcr.io/myproject/myrepository)
	Registry *string `json:"registry,omitempty"`
	// human readable name associated with registry record
	RegistryName *string `json:"registry_name,omitempty"`
	// Use TLS/SSL verification for the registry URL
	RegistryVerify *bool `json:"registry_verify,omitempty"`
}

// NewRegistryConfigurationRequest instantiates a new RegistryConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryConfigurationRequest() *RegistryConfigurationRequest {
	this := RegistryConfigurationRequest{}
	return &this
}

// NewRegistryConfigurationRequestWithDefaults instantiates a new RegistryConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryConfigurationRequestWithDefaults() *RegistryConfigurationRequest {
	this := RegistryConfigurationRequest{}
	return &this
}

// GetRegistryUser returns the RegistryUser field value if set, zero value otherwise.
func (o *RegistryConfigurationRequest) GetRegistryUser() string {
	if o == nil || o.RegistryUser == nil {
		var ret string
		return ret
	}
	return *o.RegistryUser
}

// GetRegistryUserOk returns a tuple with the RegistryUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfigurationRequest) GetRegistryUserOk() (*string, bool) {
	if o == nil || o.RegistryUser == nil {
		return nil, false
	}
	return o.RegistryUser, true
}

// HasRegistryUser returns a boolean if a field has been set.
func (o *RegistryConfigurationRequest) HasRegistryUser() bool {
	if o != nil && o.RegistryUser != nil {
		return true
	}

	return false
}

// SetRegistryUser gets a reference to the given string and assigns it to the RegistryUser field.
func (o *RegistryConfigurationRequest) SetRegistryUser(v string) {
	o.RegistryUser = &v
}

// GetRegistryPass returns the RegistryPass field value if set, zero value otherwise.
func (o *RegistryConfigurationRequest) GetRegistryPass() string {
	if o == nil || o.RegistryPass == nil {
		var ret string
		return ret
	}
	return *o.RegistryPass
}

// GetRegistryPassOk returns a tuple with the RegistryPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfigurationRequest) GetRegistryPassOk() (*string, bool) {
	if o == nil || o.RegistryPass == nil {
		return nil, false
	}
	return o.RegistryPass, true
}

// HasRegistryPass returns a boolean if a field has been set.
func (o *RegistryConfigurationRequest) HasRegistryPass() bool {
	if o != nil && o.RegistryPass != nil {
		return true
	}

	return false
}

// SetRegistryPass gets a reference to the given string and assigns it to the RegistryPass field.
func (o *RegistryConfigurationRequest) SetRegistryPass(v string) {
	o.RegistryPass = &v
}

// GetRegistryType returns the RegistryType field value if set, zero value otherwise.
func (o *RegistryConfigurationRequest) GetRegistryType() string {
	if o == nil || o.RegistryType == nil {
		var ret string
		return ret
	}
	return *o.RegistryType
}

// GetRegistryTypeOk returns a tuple with the RegistryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfigurationRequest) GetRegistryTypeOk() (*string, bool) {
	if o == nil || o.RegistryType == nil {
		return nil, false
	}
	return o.RegistryType, true
}

// HasRegistryType returns a boolean if a field has been set.
func (o *RegistryConfigurationRequest) HasRegistryType() bool {
	if o != nil && o.RegistryType != nil {
		return true
	}

	return false
}

// SetRegistryType gets a reference to the given string and assigns it to the RegistryType field.
func (o *RegistryConfigurationRequest) SetRegistryType(v string) {
	o.RegistryType = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *RegistryConfigurationRequest) GetRegistry() string {
	if o == nil || o.Registry == nil {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfigurationRequest) GetRegistryOk() (*string, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *RegistryConfigurationRequest) HasRegistry() bool {
	if o != nil && o.Registry != nil {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *RegistryConfigurationRequest) SetRegistry(v string) {
	o.Registry = &v
}

// GetRegistryName returns the RegistryName field value if set, zero value otherwise.
func (o *RegistryConfigurationRequest) GetRegistryName() string {
	if o == nil || o.RegistryName == nil {
		var ret string
		return ret
	}
	return *o.RegistryName
}

// GetRegistryNameOk returns a tuple with the RegistryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfigurationRequest) GetRegistryNameOk() (*string, bool) {
	if o == nil || o.RegistryName == nil {
		return nil, false
	}
	return o.RegistryName, true
}

// HasRegistryName returns a boolean if a field has been set.
func (o *RegistryConfigurationRequest) HasRegistryName() bool {
	if o != nil && o.RegistryName != nil {
		return true
	}

	return false
}

// SetRegistryName gets a reference to the given string and assigns it to the RegistryName field.
func (o *RegistryConfigurationRequest) SetRegistryName(v string) {
	o.RegistryName = &v
}

// GetRegistryVerify returns the RegistryVerify field value if set, zero value otherwise.
func (o *RegistryConfigurationRequest) GetRegistryVerify() bool {
	if o == nil || o.RegistryVerify == nil {
		var ret bool
		return ret
	}
	return *o.RegistryVerify
}

// GetRegistryVerifyOk returns a tuple with the RegistryVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryConfigurationRequest) GetRegistryVerifyOk() (*bool, bool) {
	if o == nil || o.RegistryVerify == nil {
		return nil, false
	}
	return o.RegistryVerify, true
}

// HasRegistryVerify returns a boolean if a field has been set.
func (o *RegistryConfigurationRequest) HasRegistryVerify() bool {
	if o != nil && o.RegistryVerify != nil {
		return true
	}

	return false
}

// SetRegistryVerify gets a reference to the given bool and assigns it to the RegistryVerify field.
func (o *RegistryConfigurationRequest) SetRegistryVerify(v bool) {
	o.RegistryVerify = &v
}

func (o RegistryConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RegistryUser != nil {
		toSerialize["registry_user"] = o.RegistryUser
	}
	if o.RegistryPass != nil {
		toSerialize["registry_pass"] = o.RegistryPass
	}
	if o.RegistryType != nil {
		toSerialize["registry_type"] = o.RegistryType
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

type NullableRegistryConfigurationRequest struct {
	value *RegistryConfigurationRequest
	isSet bool
}

func (v NullableRegistryConfigurationRequest) Get() *RegistryConfigurationRequest {
	return v.value
}

func (v *NullableRegistryConfigurationRequest) Set(val *RegistryConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryConfigurationRequest(val *RegistryConfigurationRequest) *NullableRegistryConfigurationRequest {
	return &NullableRegistryConfigurationRequest{value: val, isSet: true}
}

func (v NullableRegistryConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


