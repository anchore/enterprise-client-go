/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

<<<<<<< HEAD
API version: 2.0.0
=======
API version: 0.1.0
>>>>>>> 48fc108 (feat: updated the enterprise ref)
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// EventSubcategory A collection of events related to each other
type EventSubcategory struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Events *[]EventDescription `json:"events,omitempty"`
}

// NewEventSubcategory instantiates a new EventSubcategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSubcategory() *EventSubcategory {
	this := EventSubcategory{}
	return &this
}

// NewEventSubcategoryWithDefaults instantiates a new EventSubcategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSubcategoryWithDefaults() *EventSubcategory {
	this := EventSubcategory{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventSubcategory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubcategory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventSubcategory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventSubcategory) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventSubcategory) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubcategory) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventSubcategory) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventSubcategory) SetDescription(v string) {
	o.Description = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *EventSubcategory) GetEvents() []EventDescription {
	if o == nil || o.Events == nil {
		var ret []EventDescription
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubcategory) GetEventsOk() (*[]EventDescription, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *EventSubcategory) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []EventDescription and assigns it to the Events field.
func (o *EventSubcategory) SetEvents(v []EventDescription) {
	o.Events = &v
}

func (o EventSubcategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableEventSubcategory struct {
	value *EventSubcategory
	isSet bool
}

func (v NullableEventSubcategory) Get() *EventSubcategory {
	return v.value
}

func (v *NullableEventSubcategory) Set(val *EventSubcategory) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSubcategory) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSubcategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSubcategory(val *EventSubcategory) *NullableEventSubcategory {
	return &NullableEventSubcategory{value: val, isSet: true}
}

func (v NullableEventSubcategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSubcategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


