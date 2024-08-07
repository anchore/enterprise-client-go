/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// PolicyEvalNotificationDataAllOf struct for PolicyEvalNotificationDataAllOf
type PolicyEvalNotificationDataAllOf struct {
	NotificationPayload *PolicyEvalNotificationPayload `json:"notification_payload,omitempty"`
}

// NewPolicyEvalNotificationDataAllOf instantiates a new PolicyEvalNotificationDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvalNotificationDataAllOf() *PolicyEvalNotificationDataAllOf {
	this := PolicyEvalNotificationDataAllOf{}
	return &this
}

// NewPolicyEvalNotificationDataAllOfWithDefaults instantiates a new PolicyEvalNotificationDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvalNotificationDataAllOfWithDefaults() *PolicyEvalNotificationDataAllOf {
	this := PolicyEvalNotificationDataAllOf{}
	return &this
}

// GetNotificationPayload returns the NotificationPayload field value if set, zero value otherwise.
func (o *PolicyEvalNotificationDataAllOf) GetNotificationPayload() PolicyEvalNotificationPayload {
	if o == nil || o.NotificationPayload == nil {
		var ret PolicyEvalNotificationPayload
		return ret
	}
	return *o.NotificationPayload
}

// GetNotificationPayloadOk returns a tuple with the NotificationPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvalNotificationDataAllOf) GetNotificationPayloadOk() (*PolicyEvalNotificationPayload, bool) {
	if o == nil || o.NotificationPayload == nil {
		return nil, false
	}
	return o.NotificationPayload, true
}

// HasNotificationPayload returns a boolean if a field has been set.
func (o *PolicyEvalNotificationDataAllOf) HasNotificationPayload() bool {
	if o != nil && o.NotificationPayload != nil {
		return true
	}

	return false
}

// SetNotificationPayload gets a reference to the given PolicyEvalNotificationPayload and assigns it to the NotificationPayload field.
func (o *PolicyEvalNotificationDataAllOf) SetNotificationPayload(v PolicyEvalNotificationPayload) {
	o.NotificationPayload = &v
}

func (o PolicyEvalNotificationDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NotificationPayload != nil {
		toSerialize["notification_payload"] = o.NotificationPayload
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyEvalNotificationDataAllOf struct {
	value *PolicyEvalNotificationDataAllOf
	isSet bool
}

func (v NullablePolicyEvalNotificationDataAllOf) Get() *PolicyEvalNotificationDataAllOf {
	return v.value
}

func (v *NullablePolicyEvalNotificationDataAllOf) Set(val *PolicyEvalNotificationDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvalNotificationDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvalNotificationDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvalNotificationDataAllOf(val *PolicyEvalNotificationDataAllOf) *NullablePolicyEvalNotificationDataAllOf {
	return &NullablePolicyEvalNotificationDataAllOf{value: val, isSet: true}
}

func (v NullablePolicyEvalNotificationDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvalNotificationDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


