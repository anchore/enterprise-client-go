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

// ImportSchema struct for ImportSchema
type ImportSchema struct {
	Version string `json:"version"`
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _ImportSchema ImportSchema

// NewImportSchema instantiates a new ImportSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportSchema(version string, url string) *ImportSchema {
	this := ImportSchema{}
	this.Version = version
	this.Url = url
	return &this
}

// NewImportSchemaWithDefaults instantiates a new ImportSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportSchemaWithDefaults() *ImportSchema {
	this := ImportSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *ImportSchema) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ImportSchema) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ImportSchema) SetVersion(v string) {
	o.Version = v
}

// GetUrl returns the Url field value
func (o *ImportSchema) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ImportSchema) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ImportSchema) SetUrl(v string) {
	o.Url = v
}

func (o ImportSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ImportSchema) UnmarshalJSON(bytes []byte) (err error) {
	varImportSchema := _ImportSchema{}

	if err = json.Unmarshal(bytes, &varImportSchema); err == nil {
		*o = ImportSchema(varImportSchema)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportSchema struct {
	value *ImportSchema
	isSet bool
}

func (v NullableImportSchema) Get() *ImportSchema {
	return v.value
}

func (v *NullableImportSchema) Set(val *ImportSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableImportSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableImportSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportSchema(val *ImportSchema) *NullableImportSchema {
	return &NullableImportSchema{value: val, isSet: true}
}

func (v NullableImportSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


