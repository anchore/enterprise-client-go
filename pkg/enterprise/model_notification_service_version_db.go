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

// NotificationServiceVersionDb struct for NotificationServiceVersionDb
type NotificationServiceVersionDb struct {
	// Semantic version of the db schema
	SchemaVersion *string `json:"schema_version,omitempty"`
}

// NewNotificationServiceVersionDb instantiates a new NotificationServiceVersionDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationServiceVersionDb() *NotificationServiceVersionDb {
	this := NotificationServiceVersionDb{}
	return &this
}

// NewNotificationServiceVersionDbWithDefaults instantiates a new NotificationServiceVersionDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationServiceVersionDbWithDefaults() *NotificationServiceVersionDb {
	this := NotificationServiceVersionDb{}
	return &this
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *NotificationServiceVersionDb) GetSchemaVersion() string {
	if o == nil || o.SchemaVersion == nil {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationServiceVersionDb) GetSchemaVersionOk() (*string, bool) {
	if o == nil || o.SchemaVersion == nil {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *NotificationServiceVersionDb) HasSchemaVersion() bool {
	if o != nil && o.SchemaVersion != nil {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *NotificationServiceVersionDb) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

func (o NotificationServiceVersionDb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SchemaVersion != nil {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationServiceVersionDb struct {
	value *NotificationServiceVersionDb
	isSet bool
}

func (v NullableNotificationServiceVersionDb) Get() *NotificationServiceVersionDb {
	return v.value
}

func (v *NullableNotificationServiceVersionDb) Set(val *NotificationServiceVersionDb) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationServiceVersionDb) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationServiceVersionDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationServiceVersionDb(val *NotificationServiceVersionDb) *NullableNotificationServiceVersionDb {
	return &NullableNotificationServiceVersionDb{value: val, isSet: true}
}

func (v NullableNotificationServiceVersionDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationServiceVersionDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


