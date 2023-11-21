/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ECSServices Services defined in ECS
type ECSServices struct {
	Services []ECSService `json:"services,omitempty"`
}

// NewECSServices instantiates a new ECSServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSServices() *ECSServices {
	this := ECSServices{}
	return &this
}

// NewECSServicesWithDefaults instantiates a new ECSServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSServicesWithDefaults() *ECSServices {
	this := ECSServices{}
	return &this
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ECSServices) GetServices() []ECSService {
	if o == nil || o.Services == nil {
		var ret []ECSService
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSServices) GetServicesOk() ([]ECSService, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ECSServices) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []ECSService and assigns it to the Services field.
func (o *ECSServices) SetServices(v []ECSService) {
	o.Services = v
}

func (o ECSServices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableECSServices struct {
	value *ECSServices
	isSet bool
}

func (v NullableECSServices) Get() *ECSServices {
	return v.value
}

func (v *NullableECSServices) Set(val *ECSServices) {
	v.value = val
	v.isSet = true
}

func (v NullableECSServices) IsSet() bool {
	return v.isSet
}

func (v *NullableECSServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSServices(val *ECSServices) *NullableECSServices {
	return &NullableECSServices{value: val, isSet: true}
}

func (v NullableECSServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


