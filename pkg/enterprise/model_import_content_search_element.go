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

// ImportContentSearchElement struct for ImportContentSearchElement
type ImportContentSearchElement struct {
	Classification string `json:"classification"`
	LineNumber int32 `json:"line_number"`
	LineOffset int32 `json:"line_offset"`
	SeekPosition int32 `json:"seek_position"`
	Length int32 `json:"length"`
}

// NewImportContentSearchElement instantiates a new ImportContentSearchElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportContentSearchElement(classification string, lineNumber int32, lineOffset int32, seekPosition int32, length int32) *ImportContentSearchElement {
	this := ImportContentSearchElement{}
	this.Classification = classification
	this.LineNumber = lineNumber
	this.LineOffset = lineOffset
	this.SeekPosition = seekPosition
	this.Length = length
	return &this
}

// NewImportContentSearchElementWithDefaults instantiates a new ImportContentSearchElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportContentSearchElementWithDefaults() *ImportContentSearchElement {
	this := ImportContentSearchElement{}
	return &this
}

// GetClassification returns the Classification field value
func (o *ImportContentSearchElement) GetClassification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value
// and a boolean to check if the value has been set.
func (o *ImportContentSearchElement) GetClassificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Classification, true
}

// SetClassification sets field value
func (o *ImportContentSearchElement) SetClassification(v string) {
	o.Classification = v
}

// GetLineNumber returns the LineNumber field value
func (o *ImportContentSearchElement) GetLineNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value
// and a boolean to check if the value has been set.
func (o *ImportContentSearchElement) GetLineNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LineNumber, true
}

// SetLineNumber sets field value
func (o *ImportContentSearchElement) SetLineNumber(v int32) {
	o.LineNumber = v
}

// GetLineOffset returns the LineOffset field value
func (o *ImportContentSearchElement) GetLineOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LineOffset
}

// GetLineOffsetOk returns a tuple with the LineOffset field value
// and a boolean to check if the value has been set.
func (o *ImportContentSearchElement) GetLineOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LineOffset, true
}

// SetLineOffset sets field value
func (o *ImportContentSearchElement) SetLineOffset(v int32) {
	o.LineOffset = v
}

// GetSeekPosition returns the SeekPosition field value
func (o *ImportContentSearchElement) GetSeekPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SeekPosition
}

// GetSeekPositionOk returns a tuple with the SeekPosition field value
// and a boolean to check if the value has been set.
func (o *ImportContentSearchElement) GetSeekPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeekPosition, true
}

// SetSeekPosition sets field value
func (o *ImportContentSearchElement) SetSeekPosition(v int32) {
	o.SeekPosition = v
}

// GetLength returns the Length field value
func (o *ImportContentSearchElement) GetLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *ImportContentSearchElement) GetLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *ImportContentSearchElement) SetLength(v int32) {
	o.Length = v
}

func (o ImportContentSearchElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["classification"] = o.Classification
	}
	if true {
		toSerialize["line_number"] = o.LineNumber
	}
	if true {
		toSerialize["line_offset"] = o.LineOffset
	}
	if true {
		toSerialize["seek_position"] = o.SeekPosition
	}
	if true {
		toSerialize["length"] = o.Length
	}
	return json.Marshal(toSerialize)
}

type NullableImportContentSearchElement struct {
	value *ImportContentSearchElement
	isSet bool
}

func (v NullableImportContentSearchElement) Get() *ImportContentSearchElement {
	return v.value
}

func (v *NullableImportContentSearchElement) Set(val *ImportContentSearchElement) {
	v.value = val
	v.isSet = true
}

func (v NullableImportContentSearchElement) IsSet() bool {
	return v.isSet
}

func (v *NullableImportContentSearchElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportContentSearchElement(val *ImportContentSearchElement) *NullableImportContentSearchElement {
	return &NullableImportContentSearchElement{value: val, isSet: true}
}

func (v NullableImportContentSearchElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportContentSearchElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


