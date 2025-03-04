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

// checks if the NotificationServiceVersionApi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationServiceVersionApi{}

// NotificationServiceVersionApi Api Version string
type NotificationServiceVersionApi struct {
	// Semantic version of the api
	Version *string `json:"version,omitempty"`
}

// NewNotificationServiceVersionApi instantiates a new NotificationServiceVersionApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationServiceVersionApi() *NotificationServiceVersionApi {
	this := NotificationServiceVersionApi{}
	return &this
}

// NewNotificationServiceVersionApiWithDefaults instantiates a new NotificationServiceVersionApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationServiceVersionApiWithDefaults() *NotificationServiceVersionApi {
	this := NotificationServiceVersionApi{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NotificationServiceVersionApi) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationServiceVersionApi) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NotificationServiceVersionApi) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NotificationServiceVersionApi) SetVersion(v string) {
	o.Version = &v
}

func (o NotificationServiceVersionApi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationServiceVersionApi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableNotificationServiceVersionApi struct {
	value *NotificationServiceVersionApi
	isSet bool
}

func (v NullableNotificationServiceVersionApi) Get() *NotificationServiceVersionApi {
	return v.value
}

func (v *NullableNotificationServiceVersionApi) Set(val *NotificationServiceVersionApi) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationServiceVersionApi) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationServiceVersionApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationServiceVersionApi(val *NotificationServiceVersionApi) *NullableNotificationServiceVersionApi {
	return &NullableNotificationServiceVersionApi{value: val, isSet: true}
}

func (v NullableNotificationServiceVersionApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationServiceVersionApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


