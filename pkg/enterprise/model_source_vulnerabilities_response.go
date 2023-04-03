/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.7.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// SourceVulnerabilitiesResponse Envelope containing list of vulnerabilities for a source repo
type SourceVulnerabilitiesResponse struct {
	SourceId *string `json:"source_id,omitempty"`
	VulnerabilityType *string `json:"vulnerability_type,omitempty"`
	// List of Vulnerability objects
	Vulnerabilities *[]Vulnerability `json:"vulnerabilities,omitempty"`
}

// NewSourceVulnerabilitiesResponse instantiates a new SourceVulnerabilitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceVulnerabilitiesResponse() *SourceVulnerabilitiesResponse {
	this := SourceVulnerabilitiesResponse{}
	return &this
}

// NewSourceVulnerabilitiesResponseWithDefaults instantiates a new SourceVulnerabilitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceVulnerabilitiesResponseWithDefaults() *SourceVulnerabilitiesResponse {
	this := SourceVulnerabilitiesResponse{}
	return &this
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *SourceVulnerabilitiesResponse) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceVulnerabilitiesResponse) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *SourceVulnerabilitiesResponse) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *SourceVulnerabilitiesResponse) SetSourceId(v string) {
	o.SourceId = &v
}

// GetVulnerabilityType returns the VulnerabilityType field value if set, zero value otherwise.
func (o *SourceVulnerabilitiesResponse) GetVulnerabilityType() string {
	if o == nil || o.VulnerabilityType == nil {
		var ret string
		return ret
	}
	return *o.VulnerabilityType
}

// GetVulnerabilityTypeOk returns a tuple with the VulnerabilityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceVulnerabilitiesResponse) GetVulnerabilityTypeOk() (*string, bool) {
	if o == nil || o.VulnerabilityType == nil {
		return nil, false
	}
	return o.VulnerabilityType, true
}

// HasVulnerabilityType returns a boolean if a field has been set.
func (o *SourceVulnerabilitiesResponse) HasVulnerabilityType() bool {
	if o != nil && o.VulnerabilityType != nil {
		return true
	}

	return false
}

// SetVulnerabilityType gets a reference to the given string and assigns it to the VulnerabilityType field.
func (o *SourceVulnerabilitiesResponse) SetVulnerabilityType(v string) {
	o.VulnerabilityType = &v
}

// GetVulnerabilities returns the Vulnerabilities field value if set, zero value otherwise.
func (o *SourceVulnerabilitiesResponse) GetVulnerabilities() []Vulnerability {
	if o == nil || o.Vulnerabilities == nil {
		var ret []Vulnerability
		return ret
	}
	return *o.Vulnerabilities
}

// GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceVulnerabilitiesResponse) GetVulnerabilitiesOk() (*[]Vulnerability, bool) {
	if o == nil || o.Vulnerabilities == nil {
		return nil, false
	}
	return o.Vulnerabilities, true
}

// HasVulnerabilities returns a boolean if a field has been set.
func (o *SourceVulnerabilitiesResponse) HasVulnerabilities() bool {
	if o != nil && o.Vulnerabilities != nil {
		return true
	}

	return false
}

// SetVulnerabilities gets a reference to the given []Vulnerability and assigns it to the Vulnerabilities field.
func (o *SourceVulnerabilitiesResponse) SetVulnerabilities(v []Vulnerability) {
	o.Vulnerabilities = &v
}

func (o SourceVulnerabilitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceId != nil {
		toSerialize["source_id"] = o.SourceId
	}
	if o.VulnerabilityType != nil {
		toSerialize["vulnerability_type"] = o.VulnerabilityType
	}
	if o.Vulnerabilities != nil {
		toSerialize["vulnerabilities"] = o.Vulnerabilities
	}
	return json.Marshal(toSerialize)
}

type NullableSourceVulnerabilitiesResponse struct {
	value *SourceVulnerabilitiesResponse
	isSet bool
}

func (v NullableSourceVulnerabilitiesResponse) Get() *SourceVulnerabilitiesResponse {
	return v.value
}

func (v *NullableSourceVulnerabilitiesResponse) Set(val *SourceVulnerabilitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceVulnerabilitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceVulnerabilitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceVulnerabilitiesResponse(val *SourceVulnerabilitiesResponse) *NullableSourceVulnerabilitiesResponse {
	return &NullableSourceVulnerabilitiesResponse{value: val, isSet: true}
}

func (v NullableSourceVulnerabilitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceVulnerabilitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


