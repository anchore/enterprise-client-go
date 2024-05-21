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
)

// NotificationOperationalStatus Operational status for a specific notification endpoint configuration
type NotificationOperationalStatus struct {
	Status *string `json:"status,omitempty"`
}

// NewNotificationOperationalStatus instantiates a new NotificationOperationalStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationOperationalStatus() *NotificationOperationalStatus {
	this := NotificationOperationalStatus{}
	return &this
}

// NewNotificationOperationalStatusWithDefaults instantiates a new NotificationOperationalStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationOperationalStatusWithDefaults() *NotificationOperationalStatus {
	this := NotificationOperationalStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotificationOperationalStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationOperationalStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotificationOperationalStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NotificationOperationalStatus) SetStatus(v string) {
	o.Status = &v
}

func (o NotificationOperationalStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationOperationalStatus struct {
	value *NotificationOperationalStatus
	isSet bool
}

func (v NullableNotificationOperationalStatus) Get() *NotificationOperationalStatus {
	return v.value
}

func (v *NullableNotificationOperationalStatus) Set(val *NotificationOperationalStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationOperationalStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationOperationalStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationOperationalStatus(val *NotificationOperationalStatus) *NullableNotificationOperationalStatus {
	return &NullableNotificationOperationalStatus{value: val, isSet: true}
}

func (v NullableNotificationOperationalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationOperationalStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


