/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

<<<<<<< HEAD
API version: 2.0.0
=======
API version: 0.1.0
>>>>>>> 48fc108 (feat: updated the enterprise ref)
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// TagUpdateNotificationDataAllOf struct for TagUpdateNotificationDataAllOf
type TagUpdateNotificationDataAllOf struct {
	NotificationPayload *TagUpdateNotificationPayload `json:"notification_payload,omitempty"`
}

// NewTagUpdateNotificationDataAllOf instantiates a new TagUpdateNotificationDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagUpdateNotificationDataAllOf() *TagUpdateNotificationDataAllOf {
	this := TagUpdateNotificationDataAllOf{}
	return &this
}

// NewTagUpdateNotificationDataAllOfWithDefaults instantiates a new TagUpdateNotificationDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagUpdateNotificationDataAllOfWithDefaults() *TagUpdateNotificationDataAllOf {
	this := TagUpdateNotificationDataAllOf{}
	return &this
}

// GetNotificationPayload returns the NotificationPayload field value if set, zero value otherwise.
func (o *TagUpdateNotificationDataAllOf) GetNotificationPayload() TagUpdateNotificationPayload {
	if o == nil || o.NotificationPayload == nil {
		var ret TagUpdateNotificationPayload
		return ret
	}
	return *o.NotificationPayload
}

// GetNotificationPayloadOk returns a tuple with the NotificationPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationDataAllOf) GetNotificationPayloadOk() (*TagUpdateNotificationPayload, bool) {
	if o == nil || o.NotificationPayload == nil {
		return nil, false
	}
	return o.NotificationPayload, true
}

// HasNotificationPayload returns a boolean if a field has been set.
func (o *TagUpdateNotificationDataAllOf) HasNotificationPayload() bool {
	if o != nil && o.NotificationPayload != nil {
		return true
	}

	return false
}

// SetNotificationPayload gets a reference to the given TagUpdateNotificationPayload and assigns it to the NotificationPayload field.
func (o *TagUpdateNotificationDataAllOf) SetNotificationPayload(v TagUpdateNotificationPayload) {
	o.NotificationPayload = &v
}

func (o TagUpdateNotificationDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NotificationPayload != nil {
		toSerialize["notification_payload"] = o.NotificationPayload
	}
	return json.Marshal(toSerialize)
}

type NullableTagUpdateNotificationDataAllOf struct {
	value *TagUpdateNotificationDataAllOf
	isSet bool
}

func (v NullableTagUpdateNotificationDataAllOf) Get() *TagUpdateNotificationDataAllOf {
	return v.value
}

func (v *NullableTagUpdateNotificationDataAllOf) Set(val *TagUpdateNotificationDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTagUpdateNotificationDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTagUpdateNotificationDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagUpdateNotificationDataAllOf(val *TagUpdateNotificationDataAllOf) *NullableTagUpdateNotificationDataAllOf {
	return &NullableTagUpdateNotificationDataAllOf{value: val, isSet: true}
}

func (v NullableTagUpdateNotificationDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagUpdateNotificationDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


