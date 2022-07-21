/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.20
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// AnalysisUpdateNotificationDataAllOf struct for AnalysisUpdateNotificationDataAllOf
type AnalysisUpdateNotificationDataAllOf struct {
	NotificationPayload *AnalysisUpdateNotificationPayload `json:"notification_payload,omitempty"`
}

// NewAnalysisUpdateNotificationDataAllOf instantiates a new AnalysisUpdateNotificationDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisUpdateNotificationDataAllOf() *AnalysisUpdateNotificationDataAllOf {
	this := AnalysisUpdateNotificationDataAllOf{}
	return &this
}

// NewAnalysisUpdateNotificationDataAllOfWithDefaults instantiates a new AnalysisUpdateNotificationDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisUpdateNotificationDataAllOfWithDefaults() *AnalysisUpdateNotificationDataAllOf {
	this := AnalysisUpdateNotificationDataAllOf{}
	return &this
}

// GetNotificationPayload returns the NotificationPayload field value if set, zero value otherwise.
func (o *AnalysisUpdateNotificationDataAllOf) GetNotificationPayload() AnalysisUpdateNotificationPayload {
	if o == nil || o.NotificationPayload == nil {
		var ret AnalysisUpdateNotificationPayload
		return ret
	}
	return *o.NotificationPayload
}

// GetNotificationPayloadOk returns a tuple with the NotificationPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateNotificationDataAllOf) GetNotificationPayloadOk() (*AnalysisUpdateNotificationPayload, bool) {
	if o == nil || o.NotificationPayload == nil {
		return nil, false
	}
	return o.NotificationPayload, true
}

// HasNotificationPayload returns a boolean if a field has been set.
func (o *AnalysisUpdateNotificationDataAllOf) HasNotificationPayload() bool {
	if o != nil && o.NotificationPayload != nil {
		return true
	}

	return false
}

// SetNotificationPayload gets a reference to the given AnalysisUpdateNotificationPayload and assigns it to the NotificationPayload field.
func (o *AnalysisUpdateNotificationDataAllOf) SetNotificationPayload(v AnalysisUpdateNotificationPayload) {
	o.NotificationPayload = &v
}

func (o AnalysisUpdateNotificationDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NotificationPayload != nil {
		toSerialize["notification_payload"] = o.NotificationPayload
	}
	return json.Marshal(toSerialize)
}

type NullableAnalysisUpdateNotificationDataAllOf struct {
	value *AnalysisUpdateNotificationDataAllOf
	isSet bool
}

func (v NullableAnalysisUpdateNotificationDataAllOf) Get() *AnalysisUpdateNotificationDataAllOf {
	return v.value
}

func (v *NullableAnalysisUpdateNotificationDataAllOf) Set(val *AnalysisUpdateNotificationDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisUpdateNotificationDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisUpdateNotificationDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisUpdateNotificationDataAllOf(val *AnalysisUpdateNotificationDataAllOf) *NullableAnalysisUpdateNotificationDataAllOf {
	return &NullableAnalysisUpdateNotificationDataAllOf{value: val, isSet: true}
}

func (v NullableAnalysisUpdateNotificationDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisUpdateNotificationDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

