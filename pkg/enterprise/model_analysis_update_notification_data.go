/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.9.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the AnalysisUpdateNotificationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalysisUpdateNotificationData{}

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
	if o == nil || IsNil(o.NotificationUser) {
		var ret string
		return ret
	}
	return *o.NotificationUser
}

// GetNotificationUserOk returns a tuple with the NotificationUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationData) GetNotificationUserOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationUser) {
		return nil, false
	}
	return o.NotificationUser, true
}

// HasNotificationUser returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationData) HasNotificationUser() bool {
	if o != nil && !IsNil(o.NotificationUser) {
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
	if o == nil || IsNil(o.NotificationUserEmail) {
		var ret string
		return ret
	}
	return *o.NotificationUserEmail
}

// GetNotificationUserEmailOk returns a tuple with the NotificationUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationData) GetNotificationUserEmailOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationUserEmail) {
		return nil, false
	}
	return o.NotificationUserEmail, true
}

// HasNotificationUserEmail returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationData) HasNotificationUserEmail() bool {
	if o != nil && !IsNil(o.NotificationUserEmail) {
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
	if o == nil || IsNil(o.NotificationType) {
		var ret string
		return ret
	}
	return *o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationData) GetNotificationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationType) {
		return nil, false
	}
	return o.NotificationType, true
}

// HasNotificationType returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationData) HasNotificationType() bool {
	if o != nil && !IsNil(o.NotificationType) {
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
	if o == nil || IsNil(o.NotificationPayload) {
		var ret AnalysisUpdateNotificationPayload
		return ret
	}
	return *o.NotificationPayload
}

// GetNotificationPayloadOk returns a tuple with the NotificationPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationData) GetNotificationPayloadOk() (*AnalysisUpdateNotificationPayload, bool) {
	if o == nil || IsNil(o.NotificationPayload) {
		return nil, false
	}
	return o.NotificationPayload, true
}

// HasNotificationPayload returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationData) HasNotificationPayload() bool {
	if o != nil && !IsNil(o.NotificationPayload) {
		return true
	}

	return false
}

// SetNotificationPayload gets a reference to the given AnalysisUpdateNotificationPayload and assigns it to the NotificationPayload field.
func (o *AnalysisUpdateNotificationData) SetNotificationPayload(v AnalysisUpdateNotificationPayload) {
	o.NotificationPayload = &v
}

func (o AnalysisUpdateNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalysisUpdateNotificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationUser) {
		toSerialize["notification_user"] = o.NotificationUser
	}
	if !IsNil(o.NotificationUserEmail) {
		toSerialize["notification_user_email"] = o.NotificationUserEmail
	}
	if !IsNil(o.NotificationType) {
		toSerialize["notification_type"] = o.NotificationType
	}
	if !IsNil(o.NotificationPayload) {
		toSerialize["notification_payload"] = o.NotificationPayload
	}
	return toSerialize, nil
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


