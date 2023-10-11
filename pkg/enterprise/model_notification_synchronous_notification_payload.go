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

// checks if the NotificationSynchronousNotificationPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSynchronousNotificationPayload{}

// NotificationSynchronousNotificationPayload defines a notification payload that can be sent synchronously
type NotificationSynchronousNotificationPayload struct {
	Type *string `json:"type,omitempty"`
}

// NewNotificationSynchronousNotificationPayload instantiates a new NotificationSynchronousNotificationPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSynchronousNotificationPayload() *NotificationSynchronousNotificationPayload {
	this := NotificationSynchronousNotificationPayload{}
	return &this
}

// NewNotificationSynchronousNotificationPayloadWithDefaults instantiates a new NotificationSynchronousNotificationPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSynchronousNotificationPayloadWithDefaults() *NotificationSynchronousNotificationPayload {
	this := NotificationSynchronousNotificationPayload{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotificationSynchronousNotificationPayload) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSynchronousNotificationPayload) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotificationSynchronousNotificationPayload) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NotificationSynchronousNotificationPayload) SetType(v string) {
	o.Type = &v
}

func (o NotificationSynchronousNotificationPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSynchronousNotificationPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableNotificationSynchronousNotificationPayload struct {
	value *NotificationSynchronousNotificationPayload
	isSet bool
}

func (v NullableNotificationSynchronousNotificationPayload) Get() *NotificationSynchronousNotificationPayload {
	return v.value
}

func (v *NullableNotificationSynchronousNotificationPayload) Set(val *NotificationSynchronousNotificationPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSynchronousNotificationPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSynchronousNotificationPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSynchronousNotificationPayload(val *NotificationSynchronousNotificationPayload) *NullableNotificationSynchronousNotificationPayload {
	return &NullableNotificationSynchronousNotificationPayload{value: val, isSet: true}
}

func (v NullableNotificationSynchronousNotificationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSynchronousNotificationPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


