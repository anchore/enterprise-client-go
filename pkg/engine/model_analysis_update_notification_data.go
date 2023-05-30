/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.7.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// AnalysisUpdateNotificationData struct for AnalysisUpdateNotificationData
type AnalysisUpdateNotificationData struct {
	NotificationUser *string `json:"notification_user,omitempty"`
	NotificationUserEmail *string `json:"notification_user_email,omitempty"`
	NotificationType *string `json:"notification_type,omitempty"`
	NotificationPayload *AnalysisUpdateNotificationPayload `json:"notification_payload,omitempty"`
}

// NewAnalysisUpdateNotificationData instantiates a new AnalysisUpdateNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisUpdateNotificationData() *AnalysisUpdateNotificationData {
	this := AnalysisUpdateNotificationData{}
	return &this
}

// NewAnalysisUpdateNotificationDataWithDefaults instantiates a new AnalysisUpdateNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisUpdateNotificationDataWithDefaults() *AnalysisUpdateNotificationData {
	this := AnalysisUpdateNotificationData{}
	return &this
}

// GetNotificationUser returns the NotificationUser field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationData) GetNotificationUser() string {
	if o == nil || o.NotificationUser == nil {
		var ret string
		return ret
	}
	return *o.NotificationUser
}

// GetNotificationUserOk returns a tuple with the NotificationUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationData) GetNotificationUserOk() (*string, bool) {
	if o == nil || o.NotificationUser == nil {
		return nil, false
	}
	return o.NotificationUser, true
}

// HasNotificationUser returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationData) HasNotificationUser() bool {
	if o != nil && o.NotificationUser != nil {
		return true
	}

	return false
}

// SetNotificationUser gets a reference to the given string and assigns it to the NotificationUser field.
func (o *AnalysisUpdateNotificationData) SetNotificationUser(v string) {
	o.NotificationUser = &v
}

// GetNotificationUserEmail returns the NotificationUserEmail field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationData) GetNotificationUserEmail() string {
	if o == nil || o.NotificationUserEmail == nil {
		var ret string
		return ret
	}
	return *o.NotificationUserEmail
}

// GetNotificationUserEmailOk returns a tuple with the NotificationUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationData) GetNotificationUserEmailOk() (*string, bool) {
	if o == nil || o.NotificationUserEmail == nil {
		return nil, false
	}
	return o.NotificationUserEmail, true
}

// HasNotificationUserEmail returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationData) HasNotificationUserEmail() bool {
	if o != nil && o.NotificationUserEmail != nil {
		return true
	}

	return false
}

// SetNotificationUserEmail gets a reference to the given string and assigns it to the NotificationUserEmail field.
func (o *AnalysisUpdateNotificationData) SetNotificationUserEmail(v string) {
	o.NotificationUserEmail = &v
}

// GetNotificationType returns the NotificationType field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationData) GetNotificationType() string {
	if o == nil || o.NotificationType == nil {
		var ret string
		return ret
	}
	return *o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationData) GetNotificationTypeOk() (*string, bool) {
	if o == nil || o.NotificationType == nil {
		return nil, false
	}
	return o.NotificationType, true
}

// HasNotificationType returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationData) HasNotificationType() bool {
	if o != nil && o.NotificationType != nil {
		return true
	}

	return false
}

// SetNotificationType gets a reference to the given string and assigns it to the NotificationType field.
func (o *AnalysisUpdateNotificationData) SetNotificationType(v string) {
	o.NotificationType = &v
}

// GetNotificationPayload returns the NotificationPayload field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationData) GetNotificationPayload() AnalysisUpdateNotificationPayload {
	if o == nil || o.NotificationPayload == nil {
		var ret AnalysisUpdateNotificationPayload
		return ret
	}
	return *o.NotificationPayload
}

// GetNotificationPayloadOk returns a tuple with the NotificationPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationData) GetNotificationPayloadOk() (*AnalysisUpdateNotificationPayload, bool) {
	if o == nil || o.NotificationPayload == nil {
		return nil, false
	}
	return o.NotificationPayload, true
}

// HasNotificationPayload returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationData) HasNotificationPayload() bool {
	if o != nil && o.NotificationPayload != nil {
		return true
	}

	return false
}

// SetNotificationPayload gets a reference to the given AnalysisUpdateNotificationPayload and assigns it to the NotificationPayload field.
func (o *AnalysisUpdateNotificationData) SetNotificationPayload(v AnalysisUpdateNotificationPayload) {
	o.NotificationPayload = &v
}

func (o AnalysisUpdateNotificationData) MarshalJSON() ([]byte, error) {
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

type NullableAnalysisUpdateNotificationData struct {
	value *AnalysisUpdateNotificationData
	isSet bool
}

func (v NullableAnalysisUpdateNotificationData) Get() *AnalysisUpdateNotificationData {
	return v.value
}

func (v *NullableAnalysisUpdateNotificationData) Set(val *AnalysisUpdateNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisUpdateNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisUpdateNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisUpdateNotificationData(val *AnalysisUpdateNotificationData) *NullableAnalysisUpdateNotificationData {
	return &NullableAnalysisUpdateNotificationData{value: val, isSet: true}
}

func (v NullableAnalysisUpdateNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisUpdateNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


