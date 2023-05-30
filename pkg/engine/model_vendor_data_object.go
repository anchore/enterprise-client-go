/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.7.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// VendorDataObject struct for VendorDataObject
type VendorDataObject struct {
	// Vendor Vulnerability ID
	Id *string `json:"id,omitempty"`
	CvssV2 *CVSSV2Scores `json:"cvss_v2,omitempty"`
	CvssV3 *CVSSV3Scores `json:"cvss_v3,omitempty"`
}

// NewVendorDataObject instantiates a new VendorDataObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorDataObject() *VendorDataObject {
	this := VendorDataObject{}
	return &this
}

// NewVendorDataObjectWithDefaults instantiates a new VendorDataObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorDataObjectWithDefaults() *VendorDataObject {
	this := VendorDataObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VendorDataObject) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorDataObject) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VendorDataObject) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VendorDataObject) SetId(v string) {
	o.Id = &v
}

// GetCvssV2 returns the CvssV2 field value if set, zero value otherwise.
func (o *VendorDataObject) GetCvssV2() CVSSV2Scores {
	if o == nil || o.CvssV2 == nil {
		var ret CVSSV2Scores
		return ret
	}
	return *o.CvssV2
}

// GetCvssV2Ok returns a tuple with the CvssV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorDataObject) GetCvssV2Ok() (*CVSSV2Scores, bool) {
	if o == nil || o.CvssV2 == nil {
		return nil, false
	}
	return o.CvssV2, true
}

// HasCvssV2 returns a boolean if a field has been set.
func (o *VendorDataObject) HasCvssV2() bool {
	if o != nil && o.CvssV2 != nil {
		return true
	}

	return false
}

// SetCvssV2 gets a reference to the given CVSSV2Scores and assigns it to the CvssV2 field.
func (o *VendorDataObject) SetCvssV2(v CVSSV2Scores) {
	o.CvssV2 = &v
}

// GetCvssV3 returns the CvssV3 field value if set, zero value otherwise.
func (o *VendorDataObject) GetCvssV3() CVSSV3Scores {
	if o == nil || o.CvssV3 == nil {
		var ret CVSSV3Scores
		return ret
	}
	return *o.CvssV3
}

// GetCvssV3Ok returns a tuple with the CvssV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorDataObject) GetCvssV3Ok() (*CVSSV3Scores, bool) {
	if o == nil || o.CvssV3 == nil {
		return nil, false
	}
	return o.CvssV3, true
}

// HasCvssV3 returns a boolean if a field has been set.
func (o *VendorDataObject) HasCvssV3() bool {
	if o != nil && o.CvssV3 != nil {
		return true
	}

	return false
}

// SetCvssV3 gets a reference to the given CVSSV3Scores and assigns it to the CvssV3 field.
func (o *VendorDataObject) SetCvssV3(v CVSSV3Scores) {
	o.CvssV3 = &v
}

func (o VendorDataObject) MarshalJSON() ([]byte, error) {
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

type NullableVendorDataObject struct {
	value *VendorDataObject
	isSet bool
}

func (v NullableVendorDataObject) Get() *VendorDataObject {
	return v.value
}

func (v *NullableVendorDataObject) Set(val *VendorDataObject) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorDataObject) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorDataObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorDataObject(val *VendorDataObject) *NullableVendorDataObject {
	return &NullableVendorDataObject{value: val, isSet: true}
}

func (v NullableVendorDataObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorDataObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


