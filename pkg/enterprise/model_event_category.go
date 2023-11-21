/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// EventCategory A collection of event subcategories
type EventCategory struct {
	Category *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	Subcategories []EventSubcategory `json:"subcategories,omitempty"`
}

// NewEventCategory instantiates a new EventCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCategory() *EventCategory {
	this := EventCategory{}
	return &this
}

// NewEventCategoryWithDefaults instantiates a new EventCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCategoryWithDefaults() *EventCategory {
	this := EventCategory{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *EventCategory) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCategory) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *EventCategory) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *EventCategory) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventCategory) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCategory) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventCategory) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventCategory) SetDescription(v string) {
	o.Description = &v
}

// GetSubcategories returns the Subcategories field value if set, zero value otherwise.
func (o *EventCategory) GetSubcategories() []EventSubcategory {
	if o == nil || o.Subcategories == nil {
		var ret []EventSubcategory
		return ret
	}
	return o.Subcategories
}

// GetSubcategoriesOk returns a tuple with the Subcategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCategory) GetSubcategoriesOk() ([]EventSubcategory, bool) {
	if o == nil || o.Subcategories == nil {
		return nil, false
	}
	return o.Subcategories, true
}

// HasSubcategories returns a boolean if a field has been set.
func (o *EventCategory) HasSubcategories() bool {
	if o != nil && o.Subcategories != nil {
		return true
	}

	return false
}

// SetSubcategories gets a reference to the given []EventSubcategory and assigns it to the Subcategories field.
func (o *EventCategory) SetSubcategories(v []EventSubcategory) {
	o.Subcategories = v
}

func (o EventCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Subcategories != nil {
		toSerialize["subcategories"] = o.Subcategories
	}
	return json.Marshal(toSerialize)
}

type NullableEventCategory struct {
	value *EventCategory
	isSet bool
}

func (v NullableEventCategory) Get() *EventCategory {
	return v.value
}

func (v *NullableEventCategory) Set(val *EventCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCategory(val *EventCategory) *NullableEventCategory {
	return &NullableEventCategory{value: val, isSet: true}
}

func (v NullableEventCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


