/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.9.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LoggingLevel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoggingLevel{}

// LoggingLevel struct for LoggingLevel
type LoggingLevel struct {
	ServiceName string `json:"service_name"`
	LoggingLevel string `json:"logging_level"`
}

type _LoggingLevel LoggingLevel

// NewLoggingLevel instantiates a new LoggingLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingLevel(serviceName string, loggingLevel string) *LoggingLevel {
	this := LoggingLevel{}
	this.ServiceName = serviceName
	this.LoggingLevel = loggingLevel
	return &this
}

// NewLoggingLevelWithDefaults instantiates a new LoggingLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingLevelWithDefaults() *LoggingLevel {
	this := LoggingLevel{}
	return &this
}

// GetServiceName returns the ServiceName field value
func (o *LoggingLevel) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *LoggingLevel) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *LoggingLevel) SetServiceName(v string) {
	o.ServiceName = v
}

// GetLoggingLevel returns the LoggingLevel field value
func (o *LoggingLevel) GetLoggingLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoggingLevel
}

// GetLoggingLevelOk returns a tuple with the LoggingLevel field value
// and a boolean to check if the value has been set.
func (o *LoggingLevel) GetLoggingLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoggingLevel, true
}

// SetLoggingLevel sets field value
func (o *LoggingLevel) SetLoggingLevel(v string) {
	o.LoggingLevel = v
}

func (o LoggingLevel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoggingLevel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service_name"] = o.ServiceName
	toSerialize["logging_level"] = o.LoggingLevel
	return toSerialize, nil
}

func (o *LoggingLevel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"service_name",
		"logging_level",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLoggingLevel := _LoggingLevel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLoggingLevel)

	if err != nil {
		return err
	}

	*o = LoggingLevel(varLoggingLevel)

	return err
}

type NullableLoggingLevel struct {
	value *LoggingLevel
	isSet bool
}

func (v NullableLoggingLevel) Get() *LoggingLevel {
	return v.value
}

func (v *NullableLoggingLevel) Set(val *LoggingLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingLevel(val *LoggingLevel) *NullableLoggingLevel {
	return &NullableLoggingLevel{value: val, isSet: true}
}

func (v NullableLoggingLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoggingLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


