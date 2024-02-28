/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.3.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// TagUpdateNotificationData struct for TagUpdateNotificationData
type TagUpdateNotificationData struct {
	NotificationUser *string `json:"notification_user,omitempty"`
	NotificationUserEmail *string `json:"notification_user_email,omitempty"`
	NotificationType *string `json:"notification_type,omitempty"`
	NotificationPayload *TagUpdateNotificationPayload `json:"notification_payload,omitempty"`
}

// NewTagUpdateNotificationData instantiates a new TagUpdateNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagUpdateNotificationData() *TagUpdateNotificationData {
	this := TagUpdateNotificationData{}
	return &this
}

// NewTagUpdateNotificationDataWithDefaults instantiates a new TagUpdateNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagUpdateNotificationDataWithDefaults() *TagUpdateNotificationData {
	this := TagUpdateNotificationData{}
	return &this
}

// GetNotificationUser returns the NotificationUser field value if set, zero value otherwise.
func (o *TagUpdateNotificationData) GetNotificationUser() string {
	if o == nil || o.NotificationUser == nil {
		var ret string
		return ret
	}
	return *o.NotificationUser
}

// GetNotificationUserOk returns a tuple with the NotificationUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationData) GetNotificationUserOk() (*string, bool) {
	if o == nil || o.NotificationUser == nil {
		return nil, false
	}
	return o.NotificationUser, true
}

// HasNotificationUser returns a boolean if a field has been set.
func (o *TagUpdateNotificationData) HasNotificationUser() bool {
	if o != nil && o.NotificationUser != nil {
		return true
	}

	return false
}

// SetNotificationUser gets a reference to the given string and assigns it to the NotificationUser field.
func (o *TagUpdateNotificationData) SetNotificationUser(v string) {
	o.NotificationUser = &v
}

// GetNotificationUserEmail returns the NotificationUserEmail field value if set, zero value otherwise.
func (o *TagUpdateNotificationData) GetNotificationUserEmail() string {
	if o == nil || o.NotificationUserEmail == nil {
		var ret string
		return ret
	}
	return *o.NotificationUserEmail
}

// GetNotificationUserEmailOk returns a tuple with the NotificationUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationData) GetNotificationUserEmailOk() (*string, bool) {
	if o == nil || o.NotificationUserEmail == nil {
		return nil, false
	}
	return o.NotificationUserEmail, true
}

// HasNotificationUserEmail returns a boolean if a field has been set.
func (o *TagUpdateNotificationData) HasNotificationUserEmail() bool {
	if o != nil && o.NotificationUserEmail != nil {
		return true
	}

	return false
}

// SetNotificationUserEmail gets a reference to the given string and assigns it to the NotificationUserEmail field.
func (o *TagUpdateNotificationData) SetNotificationUserEmail(v string) {
	o.NotificationUserEmail = &v
}

// GetNotificationType returns the NotificationType field value if set, zero value otherwise.
func (o *TagUpdateNotificationData) GetNotificationType() string {
	if o == nil || o.NotificationType == nil {
		var ret string
		return ret
	}
	return *o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationData) GetNotificationTypeOk() (*string, bool) {
	if o == nil || o.NotificationType == nil {
		return nil, false
	}
	return o.NotificationType, true
}

// HasNotificationType returns a boolean if a field has been set.
func (o *TagUpdateNotificationData) HasNotificationType() bool {
	if o != nil && o.NotificationType != nil {
		return true
	}

	return false
}

// SetNotificationType gets a reference to the given string and assigns it to the NotificationType field.
func (o *TagUpdateNotificationData) SetNotificationType(v string) {
	o.NotificationType = &v
}

// GetNotificationPayload returns the NotificationPayload field value if set, zero value otherwise.
func (o *TagUpdateNotificationData) GetNotificationPayload() TagUpdateNotificationPayload {
	if o == nil || o.NotificationPayload == nil {
		var ret TagUpdateNotificationPayload
		return ret
	}
	return *o.NotificationPayload
}

// GetNotificationPayloadOk returns a tuple with the NotificationPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdateNotificationData) GetNotificationPayloadOk() (*TagUpdateNotificationPayload, bool) {
	if o == nil || o.NotificationPayload == nil {
		return nil, false
	}
	return o.NotificationPayload, true
}

// HasNotificationPayload returns a boolean if a field has been set.
func (o *TagUpdateNotificationData) HasNotificationPayload() bool {
	if o != nil && o.NotificationPayload != nil {
		return true
	}

	return false
}

// SetNotificationPayload gets a reference to the given TagUpdateNotificationPayload and assigns it to the NotificationPayload field.
func (o *TagUpdateNotificationData) SetNotificationPayload(v TagUpdateNotificationPayload) {
	o.NotificationPayload = &v
}

func (o TagUpdateNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NotificationUser != nil {
		toSerialize["notification_user"] = o.NotificationUser
	}
	if o.NotificationUserEmail != nil {
		toSerialize["notification_user_email"] = o.NotificationUserEmail
	}
	if o.NotificationType != nil {
		toSerialize["notification_type"] = o.NotificationType
	}
	if o.NotificationPayload != nil {
		toSerialize["notification_payload"] = o.NotificationPayload
	}
	return json.Marshal(toSerialize)
}

type NullableTagUpdateNotificationData struct {
	value *TagUpdateNotificationData
	isSet bool
}

func (v NullableTagUpdateNotificationData) Get() *TagUpdateNotificationData {
	return v.value
}

func (v *NullableTagUpdateNotificationData) Set(val *TagUpdateNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableTagUpdateNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableTagUpdateNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagUpdateNotificationData(val *TagUpdateNotificationData) *NullableTagUpdateNotificationData {
	return &NullableTagUpdateNotificationData{value: val, isSet: true}
}

func (v NullableTagUpdateNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagUpdateNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


