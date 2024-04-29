/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NativeSBOMPackageCPE struct for NativeSBOMPackageCPE
type NativeSBOMPackageCPE struct {
	Cpe string `json:"cpe"`
	Source string `json:"source"`
	AdditionalProperties map[string]interface{}
}

type _NativeSBOMPackageCPE NativeSBOMPackageCPE

// NewNativeSBOMPackageCPE instantiates a new NativeSBOMPackageCPE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeSBOMPackageCPE(cpe string, source string) *NativeSBOMPackageCPE {
	this := NativeSBOMPackageCPE{}
	this.Cpe = cpe
	this.Source = source
	return &this
}

// NewNativeSBOMPackageCPEWithDefaults instantiates a new NativeSBOMPackageCPE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeSBOMPackageCPEWithDefaults() *NativeSBOMPackageCPE {
	this := NativeSBOMPackageCPE{}
	return &this
}

// GetCpe returns the Cpe field value
func (o *NativeSBOMPackageCPE) GetCpe() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cpe
}

// GetCpeOk returns a tuple with the Cpe field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackageCPE) GetCpeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpe, true
}

// SetCpe sets field value
func (o *NativeSBOMPackageCPE) SetCpe(v string) {
	o.Cpe = v
}

// GetSource returns the Source field value
func (o *NativeSBOMPackageCPE) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackageCPE) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *NativeSBOMPackageCPE) SetSource(v string) {
	o.Source = v
}

func (o NativeSBOMPackageCPE) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cpe"] = o.Cpe
	}
	if true {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NativeSBOMPackageCPE) UnmarshalJSON(bytes []byte) (err error) {
	varNativeSBOMPackageCPE := _NativeSBOMPackageCPE{}

	if err = json.Unmarshal(bytes, &varNativeSBOMPackageCPE); err == nil {
		*o = NativeSBOMPackageCPE(varNativeSBOMPackageCPE)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cpe")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNativeSBOMPackageCPE struct {
	value *NativeSBOMPackageCPE
	isSet bool
}

func (v NullableNativeSBOMPackageCPE) Get() *NativeSBOMPackageCPE {
	return v.value
}

func (v *NullableNativeSBOMPackageCPE) Set(val *NativeSBOMPackageCPE) {
	v.value = val
	v.isSet = true
}

func (v NullableNativeSBOMPackageCPE) IsSet() bool {
	return v.isSet
}

func (v *NullableNativeSBOMPackageCPE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNativeSBOMPackageCPE(val *NativeSBOMPackageCPE) *NullableNativeSBOMPackageCPE {
	return &NullableNativeSBOMPackageCPE{value: val, isSet: true}
}

func (v NullableNativeSBOMPackageCPE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNativeSBOMPackageCPE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


