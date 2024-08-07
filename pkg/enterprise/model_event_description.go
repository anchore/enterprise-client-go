/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// EventDescription A description of an event type
type EventDescription struct {
	// The event type. The last component of the fully-qualified event_type (category.subcategory.event)
	Name *string `json:"name,omitempty"`
	// The fully qualified event type as would be seen in the event payload
	Type *string `json:"type,omitempty"`
	// The message associated with the event type
	Message *string `json:"message,omitempty"`
	// The type of resource this event is generated from
	ResourceType *string `json:"resource_type,omitempty"`
}

// NewEventDescription instantiates a new EventDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventDescription() *EventDescription {
	this := EventDescription{}
	return &this
}

// NewEventDescriptionWithDefaults instantiates a new EventDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventDescriptionWithDefaults() *EventDescription {
	this := EventDescription{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventDescription) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDescription) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventDescription) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventDescription) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventDescription) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDescription) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventDescription) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventDescription) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EventDescription) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDescription) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EventDescription) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EventDescription) SetMessage(v string) {
	o.Message = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *EventDescription) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDescription) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *EventDescription) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *EventDescription) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o EventDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	return json.Marshal(toSerialize)
}

type NullableEventDescription struct {
	value *EventDescription
	isSet bool
}

func (v NullableEventDescription) Get() *EventDescription {
	return v.value
}

func (v *NullableEventDescription) Set(val *EventDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableEventDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableEventDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventDescription(val *EventDescription) *NullableEventDescription {
	return &NullableEventDescription{value: val, isSet: true}
}

func (v NullableEventDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


