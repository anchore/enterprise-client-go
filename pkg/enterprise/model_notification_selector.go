/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.4.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NotificationSelector A selector for notifications that determines which notifications are passed to a specific endpoint configuration
type NotificationSelector struct {
	Uuid *string `json:"uuid,omitempty"`
	// UUID of the endpoint configuration bound to this selector
	ConfigurationUuid *string `json:"configuration_uuid,omitempty"`
	// The scope to filter events. 'global' scope encompasses all the events in the system, only the admin account can request this selector scope. 'account' covers events scoped to a specific account.
	Scope string `json:"scope"`
	Event NotificationEventSelector `json:"event"`
}

// NewNotificationSelector instantiates a new NotificationSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSelector(scope string, event NotificationEventSelector) *NotificationSelector {
	this := NotificationSelector{}
	this.Scope = scope
	this.Event = event
	return &this
}

// NewNotificationSelectorWithDefaults instantiates a new NotificationSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSelectorWithDefaults() *NotificationSelector {
	this := NotificationSelector{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NotificationSelector) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSelector) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NotificationSelector) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NotificationSelector) SetUuid(v string) {
	o.Uuid = &v
}

// GetConfigurationUuid returns the ConfigurationUuid field value if set, zero value otherwise.
func (o *NotificationSelector) GetConfigurationUuid() string {
	if o == nil || o.ConfigurationUuid == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationUuid
}

// GetConfigurationUuidOk returns a tuple with the ConfigurationUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSelector) GetConfigurationUuidOk() (*string, bool) {
	if o == nil || o.ConfigurationUuid == nil {
		return nil, false
	}
	return o.ConfigurationUuid, true
}

// HasConfigurationUuid returns a boolean if a field has been set.
func (o *NotificationSelector) HasConfigurationUuid() bool {
	if o != nil && o.ConfigurationUuid != nil {
		return true
	}

	return false
}

// SetConfigurationUuid gets a reference to the given string and assigns it to the ConfigurationUuid field.
func (o *NotificationSelector) SetConfigurationUuid(v string) {
	o.ConfigurationUuid = &v
}

// GetScope returns the Scope field value
func (o *NotificationSelector) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *NotificationSelector) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *NotificationSelector) SetScope(v string) {
	o.Scope = v
}

// GetEvent returns the Event field value
func (o *NotificationSelector) GetEvent() NotificationEventSelector {
	if o == nil {
		var ret NotificationEventSelector
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *NotificationSelector) GetEventOk() (*NotificationEventSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *NotificationSelector) SetEvent(v NotificationEventSelector) {
	o.Event = v
}

func (o NotificationSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.ConfigurationUuid != nil {
		toSerialize["configuration_uuid"] = o.ConfigurationUuid
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["event"] = o.Event
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationSelector struct {
	value *NotificationSelector
	isSet bool
}

func (v NullableNotificationSelector) Get() *NotificationSelector {
	return v.value
}

func (v *NullableNotificationSelector) Set(val *NotificationSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSelector(val *NotificationSelector) *NullableNotificationSelector {
	return &NullableNotificationSelector{value: val, isSet: true}
}

func (v NullableNotificationSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


