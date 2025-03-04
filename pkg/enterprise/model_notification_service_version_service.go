/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the NotificationServiceVersionService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationServiceVersionService{}

// NotificationServiceVersionService struct for NotificationServiceVersionService
type NotificationServiceVersionService struct {
	// Semantic Version string of the service implementation
	Version *string `json:"version,omitempty"`
}

// NewNotificationServiceVersionService instantiates a new NotificationServiceVersionService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationServiceVersionService() *NotificationServiceVersionService {
	this := NotificationServiceVersionService{}
	return &this
}

// NewNotificationServiceVersionServiceWithDefaults instantiates a new NotificationServiceVersionService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationServiceVersionServiceWithDefaults() *NotificationServiceVersionService {
	this := NotificationServiceVersionService{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NotificationServiceVersionService) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationServiceVersionService) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NotificationServiceVersionService) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NotificationServiceVersionService) SetVersion(v string) {
	o.Version = &v
}

func (o NotificationServiceVersionService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationServiceVersionService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableNotificationServiceVersionService struct {
	value *NotificationServiceVersionService
	isSet bool
}

func (v NullableNotificationServiceVersionService) Get() *NotificationServiceVersionService {
	return v.value
}

func (v *NullableNotificationServiceVersionService) Set(val *NotificationServiceVersionService) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationServiceVersionService) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationServiceVersionService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationServiceVersionService(val *NotificationServiceVersionService) *NullableNotificationServiceVersionService {
	return &NullableNotificationServiceVersionService{value: val, isSet: true}
}

func (v NullableNotificationServiceVersionService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationServiceVersionService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


