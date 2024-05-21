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

// BaseNotificationData Every notification has a payload, which follows this basic structure
type BaseNotificationData struct {
	NotificationUser *string `json:"notification_user,omitempty"`
	NotificationUserEmail *string `json:"notification_user_email,omitempty"`
	NotificationType *string `json:"notification_type,omitempty"`
}

// NewBaseNotificationData instantiates a new BaseNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseNotificationData() *BaseNotificationData {
	this := BaseNotificationData{}
	return &this
}

// NewBaseNotificationDataWithDefaults instantiates a new BaseNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseNotificationDataWithDefaults() *BaseNotificationData {
	this := BaseNotificationData{}
	return &this
}

// GetNotificationUser returns the NotificationUser field value if set, zero value otherwise.
func (o *BaseNotificationData) GetNotificationUser() string {
	if o == nil || o.NotificationUser == nil {
		var ret string
		return ret
	}
	return *o.NotificationUser
}

// GetNotificationUserOk returns a tuple with the NotificationUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseNotificationData) GetNotificationUserOk() (*string, bool) {
	if o == nil || o.NotificationUser == nil {
		return nil, false
	}
	return o.NotificationUser, true
}

// HasNotificationUser returns a boolean if a field has been set.
func (o *BaseNotificationData) HasNotificationUser() bool {
	if o != nil && o.NotificationUser != nil {
		return true
	}

	return false
}

// SetNotificationUser gets a reference to the given string and assigns it to the NotificationUser field.
func (o *BaseNotificationData) SetNotificationUser(v string) {
	o.NotificationUser = &v
}

// GetNotificationUserEmail returns the NotificationUserEmail field value if set, zero value otherwise.
func (o *BaseNotificationData) GetNotificationUserEmail() string {
	if o == nil || o.NotificationUserEmail == nil {
		var ret string
		return ret
	}
	return *o.NotificationUserEmail
}

// GetNotificationUserEmailOk returns a tuple with the NotificationUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseNotificationData) GetNotificationUserEmailOk() (*string, bool) {
	if o == nil || o.NotificationUserEmail == nil {
		return nil, false
	}
	return o.NotificationUserEmail, true
}

// HasNotificationUserEmail returns a boolean if a field has been set.
func (o *BaseNotificationData) HasNotificationUserEmail() bool {
	if o != nil && o.NotificationUserEmail != nil {
		return true
	}

	return false
}

// SetNotificationUserEmail gets a reference to the given string and assigns it to the NotificationUserEmail field.
func (o *BaseNotificationData) SetNotificationUserEmail(v string) {
	o.NotificationUserEmail = &v
}

// GetNotificationType returns the NotificationType field value if set, zero value otherwise.
func (o *BaseNotificationData) GetNotificationType() string {
	if o == nil || o.NotificationType == nil {
		var ret string
		return ret
	}
	return *o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseNotificationData) GetNotificationTypeOk() (*string, bool) {
	if o == nil || o.NotificationType == nil {
		return nil, false
	}
	return o.NotificationType, true
}

// HasNotificationType returns a boolean if a field has been set.
func (o *BaseNotificationData) HasNotificationType() bool {
	if o != nil && o.NotificationType != nil {
		return true
	}

	return false
}

// SetNotificationType gets a reference to the given string and assigns it to the NotificationType field.
func (o *BaseNotificationData) SetNotificationType(v string) {
	o.NotificationType = &v
}

func (o BaseNotificationData) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableBaseNotificationData struct {
	value *BaseNotificationData
	isSet bool
}

func (v NullableBaseNotificationData) Get() *BaseNotificationData {
	return v.value
}

func (v *NullableBaseNotificationData) Set(val *BaseNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseNotificationData(val *BaseNotificationData) *NullableBaseNotificationData {
	return &NullableBaseNotificationData{value: val, isSet: true}
}

func (v NullableBaseNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


