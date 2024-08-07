/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// RbacManagerStatusResponse System status response
type RbacManagerStatusResponse struct {
	Busy *bool `json:"busy,omitempty"`
	Up *bool `json:"up,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewRbacManagerStatusResponse instantiates a new RbacManagerStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerStatusResponse() *RbacManagerStatusResponse {
	this := RbacManagerStatusResponse{}
	return &this
}

// NewRbacManagerStatusResponseWithDefaults instantiates a new RbacManagerStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerStatusResponseWithDefaults() *RbacManagerStatusResponse {
	this := RbacManagerStatusResponse{}
	return &this
}

// GetBusy returns the Busy field value if set, zero value otherwise.
func (o *RbacManagerStatusResponse) GetBusy() bool {
	if o == nil || o.Busy == nil {
		var ret bool
		return ret
	}
	return *o.Busy
}

// GetBusyOk returns a tuple with the Busy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerStatusResponse) GetBusyOk() (*bool, bool) {
	if o == nil || o.Busy == nil {
		return nil, false
	}
	return o.Busy, true
}

// HasBusy returns a boolean if a field has been set.
func (o *RbacManagerStatusResponse) HasBusy() bool {
	if o != nil && o.Busy != nil {
		return true
	}

	return false
}

// SetBusy gets a reference to the given bool and assigns it to the Busy field.
func (o *RbacManagerStatusResponse) SetBusy(v bool) {
	o.Busy = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *RbacManagerStatusResponse) GetUp() bool {
	if o == nil || o.Up == nil {
		var ret bool
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerStatusResponse) GetUpOk() (*bool, bool) {
	if o == nil || o.Up == nil {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *RbacManagerStatusResponse) HasUp() bool {
	if o != nil && o.Up != nil {
		return true
	}

	return false
}

// SetUp gets a reference to the given bool and assigns it to the Up field.
func (o *RbacManagerStatusResponse) SetUp(v bool) {
	o.Up = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RbacManagerStatusResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerStatusResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RbacManagerStatusResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RbacManagerStatusResponse) SetMessage(v string) {
	o.Message = &v
}

func (o RbacManagerStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Busy != nil {
		toSerialize["busy"] = o.Busy
	}
	if o.Up != nil {
		toSerialize["up"] = o.Up
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableRbacManagerStatusResponse struct {
	value *RbacManagerStatusResponse
	isSet bool
}

func (v NullableRbacManagerStatusResponse) Get() *RbacManagerStatusResponse {
	return v.value
}

func (v *NullableRbacManagerStatusResponse) Set(val *RbacManagerStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerStatusResponse(val *RbacManagerStatusResponse) *NullableRbacManagerStatusResponse {
	return &NullableRbacManagerStatusResponse{value: val, isSet: true}
}

func (v NullableRbacManagerStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


