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
)

// checks if the NvdDataObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NvdDataObject{}

// NvdDataObject struct for NvdDataObject
type NvdDataObject struct {
	// NVD Vulnerability ID
	Id *string `json:"id,omitempty"`
	// The full NVD description text for the vulnerability
	Description *string `json:"description,omitempty"`
	// Indicates the classification of the CVSS
	Type *string `json:"type,omitempty"`
	// Identifies the organization or entity that generated or provided the CVSS score
	Source *string `json:"source,omitempty"`
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NvdDataObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NvdDataObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NvdDataObject) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NvdDataObject) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NvdDataObject) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NvdDataObject) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NvdDataObject) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NvdDataObject) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NvdDataObject) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NvdDataObject) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NvdDataObject) SetType(v string) {
	o.Type = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *NvdDataObject) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NvdDataObject) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *NvdDataObject) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *NvdDataObject) SetSource(v string) {
	o.Source = &v
}

// GetCvssV2 returns the CvssV2 field value if set, zero value otherwise.
func (o *NvdDataObject) GetCvssV2() CVSSV2Scores {
	if o == nil || IsNil(o.CvssV2) {
		var ret CVSSV2Scores
		return ret
	}
	return *o.CvssV2
}

// GetCvssV2Ok returns a tuple with the CvssV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NvdDataObject) GetCvssV2Ok() (*CVSSV2Scores, bool) {
	if o == nil || IsNil(o.CvssV2) {
		return nil, false
	}
	return o.CvssV2, true
}

// HasCvssV2 returns a boolean if a field has been set.
func (o *NvdDataObject) HasCvssV2() bool {
	if o != nil && !IsNil(o.CvssV2) {
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
	if o == nil || IsNil(o.CvssV3) {
		var ret CVSSV3Scores
		return ret
	}
	return *o.CvssV3
}

// GetCvssV3Ok returns a tuple with the CvssV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NvdDataObject) GetCvssV3Ok() (*CVSSV3Scores, bool) {
	if o == nil || IsNil(o.CvssV3) {
		return nil, false
	}
	return o.CvssV3, true
}

// HasCvssV3 returns a boolean if a field has been set.
func (o *NvdDataObject) HasCvssV3() bool {
	if o != nil && !IsNil(o.CvssV3) {
		return true
	}

	return false
}

// SetCvssV3 gets a reference to the given CVSSV3Scores and assigns it to the CvssV3 field.
func (o *NvdDataObject) SetCvssV3(v CVSSV3Scores) {
	o.CvssV3 = &v
}

func (o NvdDataObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NvdDataObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.CvssV2) {
		toSerialize["cvss_v2"] = o.CvssV2
	}
	if !IsNil(o.CvssV3) {
		toSerialize["cvss_v3"] = o.CvssV3
	}
	return toSerialize, nil
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


