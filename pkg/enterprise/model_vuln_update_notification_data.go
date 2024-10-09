/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// VulnUpdateNotificationData struct for VulnUpdateNotificationData
type VulnUpdateNotificationData struct {
	NotificationUser *string `json:"notification_user,omitempty"`
	NotificationUserEmail *string `json:"notification_user_email,omitempty"`
	NotificationType *string `json:"notification_type,omitempty"`
	NotificationPayload *VulnUpdateNotificationPayload `json:"notification_payload,omitempty"`
}

// NewVulnUpdateNotificationData instantiates a new VulnUpdateNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnUpdateNotificationData() *VulnUpdateNotificationData {
	this := VulnUpdateNotificationData{}
	return &this
}

// NewVulnUpdateNotificationDataWithDefaults instantiates a new VulnUpdateNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnUpdateNotificationDataWithDefaults() *VulnUpdateNotificationData {
	this := VulnUpdateNotificationData{}
	return &this
}

// GetNotificationUser returns the NotificationUser field value if set, zero value otherwise.
func (o *VulnUpdateNotificationData) GetNotificationUser() string {
	if o == nil || o.NotificationUser == nil {
		var ret string
		return ret
	}
	return *o.NotificationUser
}

// GetNotificationUserOk returns a tuple with the NotificationUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationData) GetNotificationUserOk() (*string, bool) {
	if o == nil || o.NotificationUser == nil {
		return nil, false
	}
	return o.NotificationUser, true
}

// HasNotificationUser returns a boolean if a field has been set.
func (o *VulnUpdateNotificationData) HasNotificationUser() bool {
	if o != nil && o.NotificationUser != nil {
		return true
	}

	return false
}

// SetNotificationUser gets a reference to the given string and assigns it to the NotificationUser field.
func (o *VulnUpdateNotificationData) SetNotificationUser(v string) {
	o.NotificationUser = &v
}

// GetNotificationUserEmail returns the NotificationUserEmail field value if set, zero value otherwise.
func (o *VulnUpdateNotificationData) GetNotificationUserEmail() string {
	if o == nil || o.NotificationUserEmail == nil {
		var ret string
		return ret
	}
	return *o.NotificationUserEmail
}

// GetNotificationUserEmailOk returns a tuple with the NotificationUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationData) GetNotificationUserEmailOk() (*string, bool) {
	if o == nil || o.NotificationUserEmail == nil {
		return nil, false
	}
	return o.NotificationUserEmail, true
}

// HasNotificationUserEmail returns a boolean if a field has been set.
func (o *VulnUpdateNotificationData) HasNotificationUserEmail() bool {
	if o != nil && o.NotificationUserEmail != nil {
		return true
	}

	return false
}

// SetNotificationUserEmail gets a reference to the given string and assigns it to the NotificationUserEmail field.
func (o *VulnUpdateNotificationData) SetNotificationUserEmail(v string) {
	o.NotificationUserEmail = &v
}

// GetNotificationType returns the NotificationType field value if set, zero value otherwise.
func (o *VulnUpdateNotificationData) GetNotificationType() string {
	if o == nil || o.NotificationType == nil {
		var ret string
		return ret
	}
	return *o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationData) GetNotificationTypeOk() (*string, bool) {
	if o == nil || o.NotificationType == nil {
		return nil, false
	}
	return o.NotificationType, true
}

// HasNotificationType returns a boolean if a field has been set.
func (o *VulnUpdateNotificationData) HasNotificationType() bool {
	if o != nil && o.NotificationType != nil {
		return true
	}

	return false
}

// SetNotificationType gets a reference to the given string and assigns it to the NotificationType field.
func (o *VulnUpdateNotificationData) SetNotificationType(v string) {
	o.NotificationType = &v
}

// GetNotificationPayload returns the NotificationPayload field value if set, zero value otherwise.
func (o *VulnUpdateNotificationData) GetNotificationPayload() VulnUpdateNotificationPayload {
	if o == nil || o.NotificationPayload == nil {
		var ret VulnUpdateNotificationPayload
		return ret
	}
	return *o.NotificationPayload
}

// GetNotificationPayloadOk returns a tuple with the NotificationPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationData) GetNotificationPayloadOk() (*VulnUpdateNotificationPayload, bool) {
	if o == nil || o.NotificationPayload == nil {
		return nil, false
	}
	return o.NotificationPayload, true
}

// HasNotificationPayload returns a boolean if a field has been set.
func (o *VulnUpdateNotificationData) HasNotificationPayload() bool {
	if o != nil && o.NotificationPayload != nil {
		return true
	}

	return false
}

// SetNotificationPayload gets a reference to the given VulnUpdateNotificationPayload and assigns it to the NotificationPayload field.
func (o *VulnUpdateNotificationData) SetNotificationPayload(v VulnUpdateNotificationPayload) {
	o.NotificationPayload = &v
}

func (o VulnUpdateNotificationData) MarshalJSON() ([]byte, error) {
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

type NullableVulnUpdateNotificationData struct {
	value *VulnUpdateNotificationData
	isSet bool
}

func (v NullableVulnUpdateNotificationData) Get() *VulnUpdateNotificationData {
	return v.value
}

func (v *NullableVulnUpdateNotificationData) Set(val *VulnUpdateNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnUpdateNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnUpdateNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnUpdateNotificationData(val *VulnUpdateNotificationData) *NullableVulnUpdateNotificationData {
	return &NullableVulnUpdateNotificationData{value: val, isSet: true}
}

func (v NullableVulnUpdateNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnUpdateNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


