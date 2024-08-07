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

// ImportFile struct for ImportFile
type ImportFile struct {
	// Unique identifier within the sbom for the file for other elements in the sbom to reference
	Id string `json:"id"`
	Location ImageImportFileCoordinate `json:"location"`
	// File metadata such as mode, size, etc. This is populated by anchorectl analysis but is not available in older syft-generated SBOMs
	Metadata interface{} `json:"metadata,omitempty"`
	Digests []ImportFileDigest `json:"digests,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportFile ImportFile

// NewImportFile instantiates a new ImportFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportFile(id string, location ImageImportFileCoordinate) *ImportFile {
	this := ImportFile{}
	this.Id = id
	this.Location = location
	return &this
}

// NewImportFileWithDefaults instantiates a new ImportFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportFileWithDefaults() *ImportFile {
	this := ImportFile{}
	return &this
}

// GetId returns the Id field value
func (o *ImportFile) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImportFile) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImportFile) SetId(v string) {
	o.Id = v
}

// GetLocation returns the Location field value
func (o *ImportFile) GetLocation() ImageImportFileCoordinate {
	if o == nil {
		var ret ImageImportFileCoordinate
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *ImportFile) GetLocationOk() (*ImageImportFileCoordinate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *ImportFile) SetLocation(v ImageImportFileCoordinate) {
	o.Location = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ImportFile) GetMetadata() interface{} {
	if o == nil || o.Metadata == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetMetadataOk() (interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ImportFile) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *ImportFile) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetDigests returns the Digests field value if set, zero value otherwise.
func (o *ImportFile) GetDigests() []ImportFileDigest {
	if o == nil || o.Digests == nil {
		var ret []ImportFileDigest
		return ret
	}
	return o.Digests
}

// GetDigestsOk returns a tuple with the Digests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetDigestsOk() ([]ImportFileDigest, bool) {
	if o == nil || o.Digests == nil {
		return nil, false
	}
	return o.Digests, true
}

// HasDigests returns a boolean if a field has been set.
func (o *ImportFile) HasDigests() bool {
	if o != nil && o.Digests != nil {
		return true
	}

	return false
}

// SetDigests gets a reference to the given []ImportFileDigest and assigns it to the Digests field.
func (o *ImportFile) SetDigests(v []ImportFileDigest) {
	o.Digests = v
}

func (o ImportFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Digests != nil {
		toSerialize["digests"] = o.Digests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ImportFile) UnmarshalJSON(bytes []byte) (err error) {
	varImportFile := _ImportFile{}

	if err = json.Unmarshal(bytes, &varImportFile); err == nil {
		*o = ImportFile(varImportFile)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "location")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "digests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportFile struct {
	value *ImportFile
	isSet bool
}

func (v NullableImportFile) Get() *ImportFile {
	return v.value
}

func (v *NullableImportFile) Set(val *ImportFile) {
	v.value = val
	v.isSet = true
}

func (v NullableImportFile) IsSet() bool {
	return v.isSet
}

func (v *NullableImportFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportFile(val *ImportFile) *NullableImportFile {
	return &NullableImportFile{value: val, isSet: true}
}

func (v NullableImportFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


