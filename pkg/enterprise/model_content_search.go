/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ContentSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentSearch{}

// ContentSearch struct for ContentSearch
type ContentSearch struct {
	// The media type of this result
	MediaType string `json:"media_type"`
	Metadata ContentSearchMetadata `json:"metadata"`
	Findings []ContentSearchFindingsInner `json:"findings"`
}

type _ContentSearch ContentSearch

// NewContentSearch instantiates a new ContentSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentSearch(mediaType string, metadata ContentSearchMetadata, findings []ContentSearchFindingsInner) *ContentSearch {
	this := ContentSearch{}
	this.MediaType = mediaType
	this.Metadata = metadata
	this.Findings = findings
	return &this
}

// NewContentSearchWithDefaults instantiates a new ContentSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentSearchWithDefaults() *ContentSearch {
	this := ContentSearch{}
	return &this
}

// GetMediaType returns the MediaType field value
func (o *ContentSearch) GetMediaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value
// and a boolean to check if the value has been set.
func (o *ContentSearch) GetMediaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaType, true
}

// SetMediaType sets field value
func (o *ContentSearch) SetMediaType(v string) {
	o.MediaType = v
}

// GetMetadata returns the Metadata field value
func (o *ContentSearch) GetMetadata() ContentSearchMetadata {
	if o == nil {
		var ret ContentSearchMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ContentSearch) GetMetadataOk() (*ContentSearchMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ContentSearch) SetMetadata(v ContentSearchMetadata) {
	o.Metadata = v
}

// GetFindings returns the Findings field value
func (o *ContentSearch) GetFindings() []ContentSearchFindingsInner {
	if o == nil {
		var ret []ContentSearchFindingsInner
		return ret
	}

	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *ContentSearch) GetFindingsOk() ([]ContentSearchFindingsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Findings, true
}

// SetFindings sets field value
func (o *ContentSearch) SetFindings(v []ContentSearchFindingsInner) {
	o.Findings = v
}

func (o ContentSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["media_type"] = o.MediaType
	toSerialize["metadata"] = o.Metadata
	toSerialize["findings"] = o.Findings
	return toSerialize, nil
}

func (o *ContentSearch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"media_type",
		"metadata",
		"findings",
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

	varContentSearch := _ContentSearch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContentSearch)

	if err != nil {
		return err
	}

	*o = ContentSearch(varContentSearch)

	return err
}

type NullableContentSearch struct {
	value *ContentSearch
	isSet bool
}

func (v NullableContentSearch) Get() *ContentSearch {
	return v.value
}

func (v *NullableContentSearch) Set(val *ContentSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableContentSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableContentSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentSearch(val *ContentSearch) *NullableContentSearch {
	return &NullableContentSearch{value: val, isSet: true}
}

func (v NullableContentSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


