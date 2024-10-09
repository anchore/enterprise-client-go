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

// AnchoreImageTagSummaryList a list of AnchoreImageTagSummary objects
type AnchoreImageTagSummaryList struct {
	Items []AnchoreImageTagSummary `json:"items,omitempty"`
	TotalRows *int32 `json:"total_rows,omitempty"`
}

// NewAnchoreImageTagSummaryList instantiates a new AnchoreImageTagSummaryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnchoreImageTagSummaryList() *AnchoreImageTagSummaryList {
	this := AnchoreImageTagSummaryList{}
	return &this
}

// NewAnchoreImageTagSummaryListWithDefaults instantiates a new AnchoreImageTagSummaryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnchoreImageTagSummaryListWithDefaults() *AnchoreImageTagSummaryList {
	this := AnchoreImageTagSummaryList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AnchoreImageTagSummaryList) GetItems() []AnchoreImageTagSummary {
	if o == nil || o.Items == nil {
		var ret []AnchoreImageTagSummary
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageTagSummaryList) GetItemsOk() ([]AnchoreImageTagSummary, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AnchoreImageTagSummaryList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AnchoreImageTagSummary and assigns it to the Items field.
func (o *AnchoreImageTagSummaryList) SetItems(v []AnchoreImageTagSummary) {
	o.Items = v
}

// GetTotalRows returns the TotalRows field value if set, zero value otherwise.
func (o *AnchoreImageTagSummaryList) GetTotalRows() int32 {
	if o == nil || o.TotalRows == nil {
		var ret int32
		return ret
	}
	return *o.TotalRows
}

// GetTotalRowsOk returns a tuple with the TotalRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageTagSummaryList) GetTotalRowsOk() (*int32, bool) {
	if o == nil || o.TotalRows == nil {
		return nil, false
	}
	return o.TotalRows, true
}

// HasTotalRows returns a boolean if a field has been set.
func (o *AnchoreImageTagSummaryList) HasTotalRows() bool {
	if o != nil && o.TotalRows != nil {
		return true
	}

	return false
}

// SetTotalRows gets a reference to the given int32 and assigns it to the TotalRows field.
func (o *AnchoreImageTagSummaryList) SetTotalRows(v int32) {
	o.TotalRows = &v
}

func (o AnchoreImageTagSummaryList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.TotalRows != nil {
		toSerialize["total_rows"] = o.TotalRows
	}
	return json.Marshal(toSerialize)
}

type NullableAnchoreImageTagSummaryList struct {
	value *AnchoreImageTagSummaryList
	isSet bool
}

func (v NullableAnchoreImageTagSummaryList) Get() *AnchoreImageTagSummaryList {
	return v.value
}

func (v *NullableAnchoreImageTagSummaryList) Set(val *AnchoreImageTagSummaryList) {
	v.value = val
	v.isSet = true
}

func (v NullableAnchoreImageTagSummaryList) IsSet() bool {
	return v.isSet
}

func (v *NullableAnchoreImageTagSummaryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnchoreImageTagSummaryList(val *AnchoreImageTagSummaryList) *NullableAnchoreImageTagSummaryList {
	return &NullableAnchoreImageTagSummaryList{value: val, isSet: true}
}

func (v NullableAnchoreImageTagSummaryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnchoreImageTagSummaryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


