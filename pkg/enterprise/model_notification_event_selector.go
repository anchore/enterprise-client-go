/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the NotificationEventSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationEventSelector{}

// NotificationEventSelector A selector of event properties
type NotificationEventSelector struct {
	// The level of the event to filter. '*' matches events of all levels. 'info' and 'error' match related events respectively
	Level string `json:"level"`
	// The type of resource to filter. '*' matches all resource types. Some examples of resource type are 'image_digest' or 'service'
	ResourceType string `json:"resource_type"`
	// The type of event to filter, using wildcards against type field of the event. Event types have a structured format <category>.<subcategory>.<event>. Thus, '*' matches all types of events. 'system.*' matches all system events, 'user.*' matches events that are relevant to individual consumption, and omitting an asterisk will do an exact match. See the GET /event_types route definition in the engine's external API for the list of event types. 
	Type string `json:"type" validate:"regexp=^[a-z0-9-_.*]+$"`
}

type _NotificationEventSelector NotificationEventSelector

// NewNotificationEventSelector instantiates a new NotificationEventSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationEventSelector(level string, resourceType string, type_ string) *NotificationEventSelector {
	this := NotificationEventSelector{}
	this.Level = level
	this.ResourceType = resourceType
	this.Type = type_
	return &this
}

// NewNotificationEventSelectorWithDefaults instantiates a new NotificationEventSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationEventSelectorWithDefaults() *NotificationEventSelector {
	this := NotificationEventSelector{}
	return &this
}

// GetLevel returns the Level field value
func (o *NotificationEventSelector) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *NotificationEventSelector) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *NotificationEventSelector) SetLevel(v string) {
	o.Level = v
}

// GetResourceType returns the ResourceType field value
func (o *NotificationEventSelector) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *NotificationEventSelector) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *NotificationEventSelector) SetResourceType(v string) {
	o.ResourceType = v
}

// GetType returns the Type field value
func (o *NotificationEventSelector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationEventSelector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationEventSelector) SetType(v string) {
	o.Type = v
}

func (o NotificationEventSelector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationEventSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *NotificationEventSelector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"level",
		"resource_type",
		"type",
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

	varNotificationEventSelector := _NotificationEventSelector{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationEventSelector)

	if err != nil {
		return err
	}

	*o = NotificationEventSelector(varNotificationEventSelector)

	return err
}

type NullableNotificationEventSelector struct {
	value *NotificationEventSelector
	isSet bool
}

func (v NullableNotificationEventSelector) Get() *NotificationEventSelector {
	return v.value
}

func (v *NullableNotificationEventSelector) Set(val *NotificationEventSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationEventSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationEventSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationEventSelector(val *NotificationEventSelector) *NullableNotificationEventSelector {
	return &NullableNotificationEventSelector{value: val, isSet: true}
}

func (v NullableNotificationEventSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationEventSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


