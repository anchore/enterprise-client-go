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

// NotificationStatusResponse System status response
type NotificationStatusResponse struct {
	Busy *bool `json:"busy,omitempty"`
	Up *bool `json:"up,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewNotificationStatusResponse instantiates a new NotificationStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationStatusResponse() *NotificationStatusResponse {
	this := NotificationStatusResponse{}
	return &this
}

// NewNotificationStatusResponseWithDefaults instantiates a new NotificationStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationStatusResponseWithDefaults() *NotificationStatusResponse {
	this := NotificationStatusResponse{}
	return &this
}

// GetBusy returns the Busy field value if set, zero value otherwise.
func (o *NotificationStatusResponse) GetBusy() bool {
	if o == nil || o.Busy == nil {
		var ret bool
		return ret
	}
	return *o.Busy
}

// GetBusyOk returns a tuple with the Busy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationStatusResponse) GetBusyOk() (*bool, bool) {
	if o == nil || o.Busy == nil {
		return nil, false
	}
	return o.Busy, true
}

// HasBusy returns a boolean if a field has been set.
func (o *NotificationStatusResponse) HasBusy() bool {
	if o != nil && o.Busy != nil {
		return true
	}

	return false
}

// SetBusy gets a reference to the given bool and assigns it to the Busy field.
func (o *NotificationStatusResponse) SetBusy(v bool) {
	o.Busy = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *NotificationStatusResponse) GetUp() bool {
	if o == nil || o.Up == nil {
		var ret bool
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationStatusResponse) GetUpOk() (*bool, bool) {
	if o == nil || o.Up == nil {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *NotificationStatusResponse) HasUp() bool {
	if o != nil && o.Up != nil {
		return true
	}

	return false
}

// SetUp gets a reference to the given bool and assigns it to the Up field.
func (o *NotificationStatusResponse) SetUp(v bool) {
	o.Up = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NotificationStatusResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationStatusResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NotificationStatusResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NotificationStatusResponse) SetMessage(v string) {
	o.Message = &v
}

func (o NotificationStatusResponse) MarshalJSON() ([]byte, error) {
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

type NullableNotificationStatusResponse struct {
	value *NotificationStatusResponse
	isSet bool
}

func (v NullableNotificationStatusResponse) Get() *NotificationStatusResponse {
	return v.value
}

func (v *NullableNotificationStatusResponse) Set(val *NotificationStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationStatusResponse(val *NotificationStatusResponse) *NullableNotificationStatusResponse {
	return &NullableNotificationStatusResponse{value: val, isSet: true}
}

func (v NullableNotificationStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


