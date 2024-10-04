/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the PolicyEvalNotificationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyEvalNotificationData{}

// PolicyEvalNotificationData struct for PolicyEvalNotificationData
type PolicyEvalNotificationData struct {
	NotificationUser *string `json:"notification_user,omitempty"`
	NotificationUserEmail *string `json:"notification_user_email,omitempty"`
	NotificationType *string `json:"notification_type,omitempty"`
	NotificationPayload *PolicyEvalNotificationPayload `json:"notification_payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyEvalNotificationData PolicyEvalNotificationData

// NewPolicyEvalNotificationData instantiates a new PolicyEvalNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvalNotificationData() *PolicyEvalNotificationData {
	this := PolicyEvalNotificationData{}
	return &this
}

// NewPolicyEvalNotificationDataWithDefaults instantiates a new PolicyEvalNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvalNotificationDataWithDefaults() *PolicyEvalNotificationData {
	this := PolicyEvalNotificationData{}
	return &this
}

// GetNotificationUser returns the NotificationUser field value if set, zero value otherwise.
func (o *PolicyEvalNotificationData) GetNotificationUser() string {
	if o == nil || IsNil(o.NotificationUser) {
		var ret string
		return ret
	}
	return *o.NotificationUser
}

// GetNotificationUserOk returns a tuple with the NotificationUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvalNotificationData) GetNotificationUserOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationUser) {
		return nil, false
	}
	return o.NotificationUser, true
}

// HasNotificationUser returns a boolean if a field has been set.
func (o *PolicyEvalNotificationData) HasNotificationUser() bool {
	if o != nil && !IsNil(o.NotificationUser) {
		return true
	}

	return false
}

// SetNotificationUser gets a reference to the given string and assigns it to the NotificationUser field.
func (o *PolicyEvalNotificationData) SetNotificationUser(v string) {
	o.NotificationUser = &v
}

// GetNotificationUserEmail returns the NotificationUserEmail field value if set, zero value otherwise.
func (o *PolicyEvalNotificationData) GetNotificationUserEmail() string {
	if o == nil || IsNil(o.NotificationUserEmail) {
		var ret string
		return ret
	}
	return *o.NotificationUserEmail
}

// GetNotificationUserEmailOk returns a tuple with the NotificationUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvalNotificationData) GetNotificationUserEmailOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationUserEmail) {
		return nil, false
	}
	return o.NotificationUserEmail, true
}

// HasNotificationUserEmail returns a boolean if a field has been set.
func (o *PolicyEvalNotificationData) HasNotificationUserEmail() bool {
	if o != nil && !IsNil(o.NotificationUserEmail) {
		return true
	}

	return false
}

// SetNotificationUserEmail gets a reference to the given string and assigns it to the NotificationUserEmail field.
func (o *PolicyEvalNotificationData) SetNotificationUserEmail(v string) {
	o.NotificationUserEmail = &v
}

// GetNotificationType returns the NotificationType field value if set, zero value otherwise.
func (o *PolicyEvalNotificationData) GetNotificationType() string {
	if o == nil || IsNil(o.NotificationType) {
		var ret string
		return ret
	}
	return *o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvalNotificationData) GetNotificationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationType) {
		return nil, false
	}
	return o.NotificationType, true
}

// HasNotificationType returns a boolean if a field has been set.
func (o *PolicyEvalNotificationData) HasNotificationType() bool {
	if o != nil && !IsNil(o.NotificationType) {
		return true
	}

	return false
}

// SetNotificationType gets a reference to the given string and assigns it to the NotificationType field.
func (o *PolicyEvalNotificationData) SetNotificationType(v string) {
	o.NotificationType = &v
}

// GetNotificationPayload returns the NotificationPayload field value if set, zero value otherwise.
func (o *PolicyEvalNotificationData) GetNotificationPayload() PolicyEvalNotificationPayload {
	if o == nil || IsNil(o.NotificationPayload) {
		var ret PolicyEvalNotificationPayload
		return ret
	}
	return *o.NotificationPayload
}

// GetNotificationPayloadOk returns a tuple with the NotificationPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvalNotificationData) GetNotificationPayloadOk() (*PolicyEvalNotificationPayload, bool) {
	if o == nil || IsNil(o.NotificationPayload) {
		return nil, false
	}
	return o.NotificationPayload, true
}

// HasNotificationPayload returns a boolean if a field has been set.
func (o *PolicyEvalNotificationData) HasNotificationPayload() bool {
	if o != nil && !IsNil(o.NotificationPayload) {
		return true
	}

	return false
}

// SetNotificationPayload gets a reference to the given PolicyEvalNotificationPayload and assigns it to the NotificationPayload field.
func (o *PolicyEvalNotificationData) SetNotificationPayload(v PolicyEvalNotificationPayload) {
	o.NotificationPayload = &v
}

func (o PolicyEvalNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyEvalNotificationData) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyEvalNotificationData) UnmarshalJSON(data []byte) (err error) {
	varPolicyEvalNotificationData := _PolicyEvalNotificationData{}

	err = json.Unmarshal(data, &varPolicyEvalNotificationData)

	if err != nil {
		return err
	}

	*o = PolicyEvalNotificationData(varPolicyEvalNotificationData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "notification_user")
		delete(additionalProperties, "notification_user_email")
		delete(additionalProperties, "notification_type")
		delete(additionalProperties, "notification_payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyEvalNotificationData struct {
	value *PolicyEvalNotificationData
	isSet bool
}

func (v NullablePolicyEvalNotificationData) Get() *PolicyEvalNotificationData {
	return v.value
}

func (v *NullablePolicyEvalNotificationData) Set(val *PolicyEvalNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvalNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvalNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvalNotificationData(val *PolicyEvalNotificationData) *NullablePolicyEvalNotificationData {
	return &NullablePolicyEvalNotificationData{value: val, isSet: true}
}

func (v NullablePolicyEvalNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvalNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


