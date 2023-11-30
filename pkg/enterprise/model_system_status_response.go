/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.1.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// SystemStatusResponse System status response
type SystemStatusResponse struct {
	// A list of service objects
	ServiceStates []Service `json:"service_states,omitempty"`
}

// NewSystemStatusResponse instantiates a new SystemStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemStatusResponse() *SystemStatusResponse {
	this := SystemStatusResponse{}
	return &this
}

// NewSystemStatusResponseWithDefaults instantiates a new SystemStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemStatusResponseWithDefaults() *SystemStatusResponse {
	this := SystemStatusResponse{}
	return &this
}

// GetServiceStates returns the ServiceStates field value if set, zero value otherwise.
func (o *SystemStatusResponse) GetServiceStates() []Service {
	if o == nil || o.ServiceStates == nil {
		var ret []Service
		return ret
	}
	return o.ServiceStates
}

// GetServiceStatesOk returns a tuple with the ServiceStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemStatusResponse) GetServiceStatesOk() ([]Service, bool) {
	if o == nil || o.ServiceStates == nil {
		return nil, false
	}
	return o.ServiceStates, true
}

// HasServiceStates returns a boolean if a field has been set.
func (o *SystemStatusResponse) HasServiceStates() bool {
	if o != nil && o.ServiceStates != nil {
		return true
	}

	return false
}

// SetServiceStates gets a reference to the given []Service and assigns it to the ServiceStates field.
func (o *SystemStatusResponse) SetServiceStates(v []Service) {
	o.ServiceStates = v
}

func (o SystemStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceStates != nil {
		toSerialize["service_states"] = o.ServiceStates
	}
	return json.Marshal(toSerialize)
}

type NullableSystemStatusResponse struct {
	value *SystemStatusResponse
	isSet bool
}

func (v NullableSystemStatusResponse) Get() *SystemStatusResponse {
	return v.value
}

func (v *NullableSystemStatusResponse) Set(val *SystemStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemStatusResponse(val *SystemStatusResponse) *NullableSystemStatusResponse {
	return &NullableSystemStatusResponse{value: val, isSet: true}
}

func (v NullableSystemStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


