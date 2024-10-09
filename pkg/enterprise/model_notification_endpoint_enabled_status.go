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
	"time"
)

// NotificationEndpointEnabledStatus struct for NotificationEndpointEnabledStatus
type NotificationEndpointEnabledStatus struct {
	// Is the endpoint enabled for use in the system. Affects all usage, including system-level if set to false.
	Enabled *bool `json:"enabled,omitempty"`
	// The timestamp of the last change to the status
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewNotificationEndpointEnabledStatus instantiates a new NotificationEndpointEnabledStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationEndpointEnabledStatus() *NotificationEndpointEnabledStatus {
	this := NotificationEndpointEnabledStatus{}
	return &this
}

// NewNotificationEndpointEnabledStatusWithDefaults instantiates a new NotificationEndpointEnabledStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationEndpointEnabledStatusWithDefaults() *NotificationEndpointEnabledStatus {
	this := NotificationEndpointEnabledStatus{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NotificationEndpointEnabledStatus) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEndpointEnabledStatus) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NotificationEndpointEnabledStatus) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NotificationEndpointEnabledStatus) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NotificationEndpointEnabledStatus) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEndpointEnabledStatus) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NotificationEndpointEnabledStatus) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NotificationEndpointEnabledStatus) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o NotificationEndpointEnabledStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationEndpointEnabledStatus struct {
	value *NotificationEndpointEnabledStatus
	isSet bool
}

func (v NullableNotificationEndpointEnabledStatus) Get() *NotificationEndpointEnabledStatus {
	return v.value
}

func (v *NullableNotificationEndpointEnabledStatus) Set(val *NotificationEndpointEnabledStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationEndpointEnabledStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationEndpointEnabledStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationEndpointEnabledStatus(val *NotificationEndpointEnabledStatus) *NullableNotificationEndpointEnabledStatus {
	return &NullableNotificationEndpointEnabledStatus{value: val, isSet: true}
}

func (v NullableNotificationEndpointEnabledStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationEndpointEnabledStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


