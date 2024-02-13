/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.2.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// SourceImportMetadataContents Digest of content to use in the final import
type SourceImportMetadataContents struct {
	// Digest to use for the sbom
	Sbom string `json:"sbom"`
}

// NewSourceImportMetadataContents instantiates a new SourceImportMetadataContents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceImportMetadataContents(sbom string) *SourceImportMetadataContents {
	this := SourceImportMetadataContents{}
	this.Sbom = sbom
	return &this
}

// NewSourceImportMetadataContentsWithDefaults instantiates a new SourceImportMetadataContents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceImportMetadataContentsWithDefaults() *SourceImportMetadataContents {
	this := SourceImportMetadataContents{}
	return &this
}

// GetSbom returns the Sbom field value
func (o *SourceImportMetadataContents) GetSbom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sbom
}

// GetSbomOk returns a tuple with the Sbom field value
// and a boolean to check if the value has been set.
func (o *SourceImportMetadataContents) GetSbomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sbom, true
}

// SetSbom sets field value
func (o *SourceImportMetadataContents) SetSbom(v string) {
	o.Sbom = v
}

func (o SourceImportMetadataContents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sbom"] = o.Sbom
	}
	return json.Marshal(toSerialize)
}

type NullableSourceImportMetadataContents struct {
	value *SourceImportMetadataContents
	isSet bool
}

func (v NullableSourceImportMetadataContents) Get() *SourceImportMetadataContents {
	return v.value
}

func (v *NullableSourceImportMetadataContents) Set(val *SourceImportMetadataContents) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceImportMetadataContents) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceImportMetadataContents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceImportMetadataContents(val *SourceImportMetadataContents) *NullableSourceImportMetadataContents {
	return &NullableSourceImportMetadataContents{value: val, isSet: true}
}

func (v NullableSourceImportMetadataContents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceImportMetadataContents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


