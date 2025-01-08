/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the NotificationServiceVersionEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationServiceVersionEngine{}

// NotificationServiceVersionEngine struct for NotificationServiceVersionEngine
type NotificationServiceVersionEngine struct {
	// Version of the installed engine library
	Version *string `json:"version,omitempty"`
	// Version of the installed engine db schema
	Db *string `json:"db,omitempty"`
}

// NewNotificationServiceVersionEngine instantiates a new NotificationServiceVersionEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationServiceVersionEngine() *NotificationServiceVersionEngine {
	this := NotificationServiceVersionEngine{}
	return &this
}

// NewNotificationServiceVersionEngineWithDefaults instantiates a new NotificationServiceVersionEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationServiceVersionEngineWithDefaults() *NotificationServiceVersionEngine {
	this := NotificationServiceVersionEngine{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NotificationServiceVersionEngine) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationServiceVersionEngine) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NotificationServiceVersionEngine) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NotificationServiceVersionEngine) SetVersion(v string) {
	o.Version = &v
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *NotificationServiceVersionEngine) GetDb() string {
	if o == nil || IsNil(o.Db) {
		var ret string
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationServiceVersionEngine) GetDbOk() (*string, bool) {
	if o == nil || IsNil(o.Db) {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *NotificationServiceVersionEngine) HasDb() bool {
	if o != nil && !IsNil(o.Db) {
		return true
	}

	return false
}

// SetDb gets a reference to the given string and assigns it to the Db field.
func (o *NotificationServiceVersionEngine) SetDb(v string) {
	o.Db = &v
}

func (o NotificationServiceVersionEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationServiceVersionEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Db) {
		toSerialize["db"] = o.Db
	}
	return toSerialize, nil
}

type NullableNotificationServiceVersionEngine struct {
	value *NotificationServiceVersionEngine
	isSet bool
}

func (v NullableNotificationServiceVersionEngine) Get() *NotificationServiceVersionEngine {
	return v.value
}

func (v *NullableNotificationServiceVersionEngine) Set(val *NotificationServiceVersionEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationServiceVersionEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationServiceVersionEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationServiceVersionEngine(val *NotificationServiceVersionEngine) *NullableNotificationServiceVersionEngine {
	return &NullableNotificationServiceVersionEngine{value: val, isSet: true}
}

func (v NullableNotificationServiceVersionEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationServiceVersionEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


