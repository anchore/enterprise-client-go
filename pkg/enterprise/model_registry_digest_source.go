/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the RegistryDigestSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistryDigestSource{}

// RegistryDigestSource An image reference using a digest in a registry, includes some extra tag and timestamp info in addition to the pull string to allow proper tag history reconstruction.
type RegistryDigestSource struct {
	// A digest-based pull string (e.g. docker.io/nginx@sha256:123abc)
	PullString string `json:"pull_string"`
	// A valid docker tag reference (e.g. docker.io/nginx:latest) that will be associated with the image but not used to pull the image.
	Tag string `json:"tag"`
	// Optional override of the image creation time to support proper tag history construction in cases of out-of-order analysis compared to registry history for the tag
	CreationTimestampOverride *time.Time `json:"creation_timestamp_override,omitempty"`
	// Base64 encoded content of the dockerfile used to build the image, if available.
	Dockerfile *string `json:"dockerfile,omitempty" validate:"regexp=^[a-zA-Z0-9+\\/=]+$"`
}

type _RegistryDigestSource RegistryDigestSource

// NewRegistryDigestSource instantiates a new RegistryDigestSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryDigestSource(pullString string, tag string) *RegistryDigestSource {
	this := RegistryDigestSource{}
	this.PullString = pullString
	this.Tag = tag
	return &this
}

// NewRegistryDigestSourceWithDefaults instantiates a new RegistryDigestSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryDigestSourceWithDefaults() *RegistryDigestSource {
	this := RegistryDigestSource{}
	return &this
}

// GetPullString returns the PullString field value
func (o *RegistryDigestSource) GetPullString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PullString
}

// GetPullStringOk returns a tuple with the PullString field value
// and a boolean to check if the value has been set.
func (o *RegistryDigestSource) GetPullStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PullString, true
}

// SetPullString sets field value
func (o *RegistryDigestSource) SetPullString(v string) {
	o.PullString = v
}

// GetTag returns the Tag field value
func (o *RegistryDigestSource) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *RegistryDigestSource) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *RegistryDigestSource) SetTag(v string) {
	o.Tag = v
}

// GetCreationTimestampOverride returns the CreationTimestampOverride field value if set, zero value otherwise.
func (o *RegistryDigestSource) GetCreationTimestampOverride() time.Time {
	if o == nil || IsNil(o.CreationTimestampOverride) {
		var ret time.Time
		return ret
	}
	return *o.CreationTimestampOverride
}

// GetCreationTimestampOverrideOk returns a tuple with the CreationTimestampOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryDigestSource) GetCreationTimestampOverrideOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTimestampOverride) {
		return nil, false
	}
	return o.CreationTimestampOverride, true
}

// HasCreationTimestampOverride returns a boolean if a field has been set.
func (o *RegistryDigestSource) HasCreationTimestampOverride() bool {
	if o != nil && !IsNil(o.CreationTimestampOverride) {
		return true
	}

	return false
}

// SetCreationTimestampOverride gets a reference to the given time.Time and assigns it to the CreationTimestampOverride field.
func (o *RegistryDigestSource) SetCreationTimestampOverride(v time.Time) {
	o.CreationTimestampOverride = &v
}

// GetDockerfile returns the Dockerfile field value if set, zero value otherwise.
func (o *RegistryDigestSource) GetDockerfile() string {
	if o == nil || IsNil(o.Dockerfile) {
		var ret string
		return ret
	}
	return *o.Dockerfile
}

// GetDockerfileOk returns a tuple with the Dockerfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryDigestSource) GetDockerfileOk() (*string, bool) {
	if o == nil || IsNil(o.Dockerfile) {
		return nil, false
	}
	return o.Dockerfile, true
}

// HasDockerfile returns a boolean if a field has been set.
func (o *RegistryDigestSource) HasDockerfile() bool {
	if o != nil && !IsNil(o.Dockerfile) {
		return true
	}

	return false
}

// SetDockerfile gets a reference to the given string and assigns it to the Dockerfile field.
func (o *RegistryDigestSource) SetDockerfile(v string) {
	o.Dockerfile = &v
}

func (o RegistryDigestSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistryDigestSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pull_string"] = o.PullString
	toSerialize["tag"] = o.Tag
	if !IsNil(o.CreationTimestampOverride) {
		toSerialize["creation_timestamp_override"] = o.CreationTimestampOverride
	}
	if !IsNil(o.Dockerfile) {
		toSerialize["dockerfile"] = o.Dockerfile
	}
	return toSerialize, nil
}

func (o *RegistryDigestSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pull_string",
		"tag",
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

	varRegistryDigestSource := _RegistryDigestSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistryDigestSource)

	if err != nil {
		return err
	}

	*o = RegistryDigestSource(varRegistryDigestSource)

	return err
}

type NullableRegistryDigestSource struct {
	value *RegistryDigestSource
	isSet bool
}

func (v NullableRegistryDigestSource) Get() *RegistryDigestSource {
	return v.value
}

func (v *NullableRegistryDigestSource) Set(val *RegistryDigestSource) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryDigestSource) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryDigestSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryDigestSource(val *RegistryDigestSource) *NullableRegistryDigestSource {
	return &NullableRegistryDigestSource{value: val, isSet: true}
}

func (v NullableRegistryDigestSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryDigestSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


