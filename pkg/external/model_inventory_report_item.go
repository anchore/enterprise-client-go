/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.2.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external

import (
	"encoding/json"
)

// InventoryReportItem Defines a particular context for an inventory
type InventoryReportItem struct {
	Namespace *string `json:"namespace,omitempty"`
	Images []InventoryReportImage `json:"images,omitempty"`
}

// NewInventoryReportItem instantiates a new InventoryReportItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReportItem() *InventoryReportItem {
	this := InventoryReportItem{}
	return &this
}

// NewInventoryReportItemWithDefaults instantiates a new InventoryReportItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReportItemWithDefaults() *InventoryReportItem {
	this := InventoryReportItem{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *InventoryReportItem) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReportItem) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *InventoryReportItem) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *InventoryReportItem) SetNamespace(v string) {
	o.Namespace = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *InventoryReportItem) GetImages() []InventoryReportImage {
	if o == nil || o.Images == nil {
		var ret []InventoryReportImage
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReportItem) GetImagesOk() ([]InventoryReportImage, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *InventoryReportItem) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []InventoryReportImage and assigns it to the Images field.
func (o *InventoryReportItem) SetImages(v []InventoryReportImage) {
	o.Images = v
}

func (o InventoryReportItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryReportItem struct {
	value *InventoryReportItem
	isSet bool
}

func (v NullableInventoryReportItem) Get() *InventoryReportItem {
	return v.value
}

func (v *NullableInventoryReportItem) Set(val *InventoryReportItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReportItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReportItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReportItem(val *InventoryReportItem) *NullableInventoryReportItem {
	return &NullableInventoryReportItem{value: val, isSet: true}
}

func (v NullableInventoryReportItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReportItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


