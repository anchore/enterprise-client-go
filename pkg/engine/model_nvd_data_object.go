/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.21
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// NvdDataObject struct for NvdDataObject
type NvdDataObject struct {
	// NVD Vulnerability ID
	Id *string `json:"id,omitempty"`
	CvssV2 *CVSSV2Scores `json:"cvss_v2,omitempty"`
	CvssV3 *CVSSV3Scores `json:"cvss_v3,omitempty"`
}

// NewNvdDataObject instantiates a new NvdDataObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNvdDataObject() *NvdDataObject {
	this := NvdDataObject{}
	return &this
}

// NewNvdDataObjectWithDefaults instantiates a new NvdDataObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNvdDataObjectWithDefaults() *NvdDataObject {
	this := NvdDataObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NvdDataObject) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NvdDataObject) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NvdDataObject) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NvdDataObject) SetId(v string) {
	o.Id = &v
}

// GetCvssV2 returns the CvssV2 field value if set, zero value otherwise.
func (o *NvdDataObject) GetCvssV2() CVSSV2Scores {
	if o == nil || o.CvssV2 == nil {
		var ret CVSSV2Scores
		return ret
	}
	return *o.CvssV2
}

// GetCvssV2Ok returns a tuple with the CvssV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NvdDataObject) GetCvssV2Ok() (*CVSSV2Scores, bool) {
	if o == nil || o.CvssV2 == nil {
		return nil, false
	}
	return o.CvssV2, true
}

// HasCvssV2 returns a boolean if a field has been set.
func (o *NvdDataObject) HasCvssV2() bool {
	if o != nil && o.CvssV2 != nil {
		return true
	}

	return false
}

// SetCvssV2 gets a reference to the given CVSSV2Scores and assigns it to the CvssV2 field.
func (o *NvdDataObject) SetCvssV2(v CVSSV2Scores) {
	o.CvssV2 = &v
}

// GetCvssV3 returns the CvssV3 field value if set, zero value otherwise.
func (o *NvdDataObject) GetCvssV3() CVSSV3Scores {
	if o == nil || o.CvssV3 == nil {
		var ret CVSSV3Scores
		return ret
	}
	return *o.CvssV3
}

// GetCvssV3Ok returns a tuple with the CvssV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NvdDataObject) GetCvssV3Ok() (*CVSSV3Scores, bool) {
	if o == nil || o.CvssV3 == nil {
		return nil, false
	}
	return o.CvssV3, true
}

// HasCvssV3 returns a boolean if a field has been set.
func (o *NvdDataObject) HasCvssV3() bool {
	if o != nil && o.CvssV3 != nil {
		return true
	}

	return false
}

// SetCvssV3 gets a reference to the given CVSSV3Scores and assigns it to the CvssV3 field.
func (o *NvdDataObject) SetCvssV3(v CVSSV3Scores) {
	o.CvssV3 = &v
}

func (o NvdDataObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CvssV2 != nil {
		toSerialize["cvss_v2"] = o.CvssV2
	}
	if o.CvssV3 != nil {
		toSerialize["cvss_v3"] = o.CvssV3
	}
	return json.Marshal(toSerialize)
}

type NullableNvdDataObject struct {
	value *NvdDataObject
	isSet bool
}

func (v NullableNvdDataObject) Get() *NvdDataObject {
	return v.value
}

func (v *NullableNvdDataObject) Set(val *NvdDataObject) {
	v.value = val
	v.isSet = true
}

func (v NullableNvdDataObject) IsSet() bool {
	return v.isSet
}

func (v *NullableNvdDataObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNvdDataObject(val *NvdDataObject) *NullableNvdDataObject {
	return &NullableNvdDataObject{value: val, isSet: true}
}

func (v NullableNvdDataObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNvdDataObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


