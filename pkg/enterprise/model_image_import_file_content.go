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
	"bytes"
	"fmt"
)

// checks if the ImageImportFileContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageImportFileContent{}

// ImageImportFileContent struct for ImageImportFileContent
type ImageImportFileContent struct {
	Location ImportPackageLocation `json:"location"`
	Contents string `json:"contents"`
}

type _ImageImportFileContent ImageImportFileContent

// NewImageImportFileContent instantiates a new ImageImportFileContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageImportFileContent(location ImportPackageLocation, contents string) *ImageImportFileContent {
	this := ImageImportFileContent{}
	this.Location = location
	this.Contents = contents
	return &this
}

// NewImageImportFileContentWithDefaults instantiates a new ImageImportFileContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageImportFileContentWithDefaults() *ImageImportFileContent {
	this := ImageImportFileContent{}
	return &this
}

// GetLocation returns the Location field value
func (o *ImageImportFileContent) GetLocation() ImportPackageLocation {
	if o == nil {
		var ret ImportPackageLocation
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *ImageImportFileContent) GetLocationOk() (*ImportPackageLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *ImageImportFileContent) SetLocation(v ImportPackageLocation) {
	o.Location = v
}

// GetContents returns the Contents field value
func (o *ImageImportFileContent) GetContents() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value
// and a boolean to check if the value has been set.
func (o *ImageImportFileContent) GetContentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contents, true
}

// SetContents sets field value
func (o *ImageImportFileContent) SetContents(v string) {
	o.Contents = v
}

func (o ImageImportFileContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageImportFileContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	toSerialize["contents"] = o.Contents
	return toSerialize, nil
}

func (o *ImageImportFileContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"location",
		"contents",
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

	varImageImportFileContent := _ImageImportFileContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImageImportFileContent)

	if err != nil {
		return err
	}

	*o = ImageImportFileContent(varImageImportFileContent)

	return err
}

type NullableImageImportFileContent struct {
	value *ImageImportFileContent
	isSet bool
}

func (v NullableImageImportFileContent) Get() *ImageImportFileContent {
	return v.value
}

func (v *NullableImageImportFileContent) Set(val *ImageImportFileContent) {
	v.value = val
	v.isSet = true
}

func (v NullableImageImportFileContent) IsSet() bool {
	return v.isSet
}

func (v *NullableImageImportFileContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageImportFileContent(val *ImageImportFileContent) *NullableImageImportFileContent {
	return &NullableImageImportFileContent{value: val, isSet: true}
}

func (v NullableImageImportFileContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageImportFileContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


