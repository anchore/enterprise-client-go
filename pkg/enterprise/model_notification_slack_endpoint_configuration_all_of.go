/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.2.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NotificationSlackEndpointConfigurationAllOf struct for NotificationSlackEndpointConfigurationAllOf
type NotificationSlackEndpointConfigurationAllOf struct {
	// url to POST to, including any query parameters, should begin with 'http://' or 'https://'
	Url string `json:"url"`
}

// NewNotificationSlackEndpointConfigurationAllOf instantiates a new NotificationSlackEndpointConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSlackEndpointConfigurationAllOf(url string) *NotificationSlackEndpointConfigurationAllOf {
	this := NotificationSlackEndpointConfigurationAllOf{}
	this.Url = url
	return &this
}

// NewNotificationSlackEndpointConfigurationAllOfWithDefaults instantiates a new NotificationSlackEndpointConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSlackEndpointConfigurationAllOfWithDefaults() *NotificationSlackEndpointConfigurationAllOf {
	this := NotificationSlackEndpointConfigurationAllOf{}
	return &this
}

// GetUrl returns the Url field value
func (o *NotificationSlackEndpointConfigurationAllOf) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NotificationSlackEndpointConfigurationAllOf) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NotificationSlackEndpointConfigurationAllOf) SetUrl(v string) {
	o.Url = v
}

func (o NotificationSlackEndpointConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationSlackEndpointConfigurationAllOf struct {
	value *NotificationSlackEndpointConfigurationAllOf
	isSet bool
}

func (v NullableNotificationSlackEndpointConfigurationAllOf) Get() *NotificationSlackEndpointConfigurationAllOf {
	return v.value
}

func (v *NullableNotificationSlackEndpointConfigurationAllOf) Set(val *NotificationSlackEndpointConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSlackEndpointConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSlackEndpointConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSlackEndpointConfigurationAllOf(val *NotificationSlackEndpointConfigurationAllOf) *NullableNotificationSlackEndpointConfigurationAllOf {
	return &NullableNotificationSlackEndpointConfigurationAllOf{value: val, isSet: true}
}

func (v NullableNotificationSlackEndpointConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSlackEndpointConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


