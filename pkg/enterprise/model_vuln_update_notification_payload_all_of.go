/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 1.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// VulnUpdateNotificationPayloadAllOf struct for VulnUpdateNotificationPayloadAllOf
type VulnUpdateNotificationPayloadAllOf struct {
	DiffVulnerabilityResult *VulnDiffResult `json:"diff_vulnerability_result,omitempty"`
	ImageDigest *string `json:"image_digest,omitempty"`
	// List of Corresponding Image Annotations
	Annotations interface{} `json:"annotations,omitempty"`
}

// NewVulnUpdateNotificationPayloadAllOf instantiates a new VulnUpdateNotificationPayloadAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnUpdateNotificationPayloadAllOf() *VulnUpdateNotificationPayloadAllOf {
	this := VulnUpdateNotificationPayloadAllOf{}
	return &this
}

// NewVulnUpdateNotificationPayloadAllOfWithDefaults instantiates a new VulnUpdateNotificationPayloadAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnUpdateNotificationPayloadAllOfWithDefaults() *VulnUpdateNotificationPayloadAllOf {
	this := VulnUpdateNotificationPayloadAllOf{}
	return &this
}

// GetDiffVulnerabilityResult returns the DiffVulnerabilityResult field value if set, zero value otherwise.
func (o *VulnUpdateNotificationPayloadAllOf) GetDiffVulnerabilityResult() VulnDiffResult {
	if o == nil || o.DiffVulnerabilityResult == nil {
		var ret VulnDiffResult
		return ret
	}
	return *o.DiffVulnerabilityResult
}

// GetDiffVulnerabilityResultOk returns a tuple with the DiffVulnerabilityResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationPayloadAllOf) GetDiffVulnerabilityResultOk() (*VulnDiffResult, bool) {
	if o == nil || o.DiffVulnerabilityResult == nil {
		return nil, false
	}
	return o.DiffVulnerabilityResult, true
}

// HasDiffVulnerabilityResult returns a boolean if a field has been set.
func (o *VulnUpdateNotificationPayloadAllOf) HasDiffVulnerabilityResult() bool {
	if o != nil && o.DiffVulnerabilityResult != nil {
		return true
	}

	return false
}

// SetDiffVulnerabilityResult gets a reference to the given VulnDiffResult and assigns it to the DiffVulnerabilityResult field.
func (o *VulnUpdateNotificationPayloadAllOf) SetDiffVulnerabilityResult(v VulnDiffResult) {
	o.DiffVulnerabilityResult = &v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *VulnUpdateNotificationPayloadAllOf) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationPayloadAllOf) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *VulnUpdateNotificationPayloadAllOf) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *VulnUpdateNotificationPayloadAllOf) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VulnUpdateNotificationPayloadAllOf) GetAnnotations() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VulnUpdateNotificationPayloadAllOf) GetAnnotationsOk() (*interface{}, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *VulnUpdateNotificationPayloadAllOf) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *VulnUpdateNotificationPayloadAllOf) SetAnnotations(v interface{}) {
	o.Annotations = v
}

func (o VulnUpdateNotificationPayloadAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DiffVulnerabilityResult != nil {
		toSerialize["diff_vulnerability_result"] = o.DiffVulnerabilityResult
	}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableVulnUpdateNotificationPayloadAllOf struct {
	value *VulnUpdateNotificationPayloadAllOf
	isSet bool
}

func (v NullableVulnUpdateNotificationPayloadAllOf) Get() *VulnUpdateNotificationPayloadAllOf {
	return v.value
}

func (v *NullableVulnUpdateNotificationPayloadAllOf) Set(val *VulnUpdateNotificationPayloadAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnUpdateNotificationPayloadAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnUpdateNotificationPayloadAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnUpdateNotificationPayloadAllOf(val *VulnUpdateNotificationPayloadAllOf) *NullableVulnUpdateNotificationPayloadAllOf {
	return &NullableVulnUpdateNotificationPayloadAllOf{value: val, isSet: true}
}

func (v NullableVulnUpdateNotificationPayloadAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnUpdateNotificationPayloadAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


