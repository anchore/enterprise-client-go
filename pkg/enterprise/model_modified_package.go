/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.9.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the ModifiedPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifiedPackage{}

// ModifiedPackage A combined modification record showing the source and target packages and the json patch to apply to the relationship-target object to result in the relationship-source package. Source and target are the packages from the respective sides of the relationship.
type ModifiedPackage struct {
	Source *Package `json:"source,omitempty"`
	Target *Package `json:"target,omitempty"`
	Patch *CustomJsonPatch `json:"patch,omitempty"`
}

// NewModifiedPackage instantiates a new ModifiedPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifiedPackage() *ModifiedPackage {
	this := ModifiedPackage{}
	return &this
}

// NewModifiedPackageWithDefaults instantiates a new ModifiedPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifiedPackageWithDefaults() *ModifiedPackage {
	this := ModifiedPackage{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ModifiedPackage) GetSource() Package {
	if o == nil || IsNil(o.Source) {
		var ret Package
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifiedPackage) GetSourceOk() (*Package, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ModifiedPackage) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Package and assigns it to the Source field.
func (o *ModifiedPackage) SetSource(v Package) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ModifiedPackage) GetTarget() Package {
	if o == nil || IsNil(o.Target) {
		var ret Package
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifiedPackage) GetTargetOk() (*Package, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ModifiedPackage) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given Package and assigns it to the Target field.
func (o *ModifiedPackage) SetTarget(v Package) {
	o.Target = &v
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *ModifiedPackage) GetPatch() CustomJsonPatch {
	if o == nil || IsNil(o.Patch) {
		var ret CustomJsonPatch
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifiedPackage) GetPatchOk() (*CustomJsonPatch, bool) {
	if o == nil || IsNil(o.Patch) {
		return nil, false
	}
	return o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *ModifiedPackage) HasPatch() bool {
	if o != nil && !IsNil(o.Patch) {
		return true
	}

	return false
}

// SetPatch gets a reference to the given CustomJsonPatch and assigns it to the Patch field.
func (o *ModifiedPackage) SetPatch(v CustomJsonPatch) {
	o.Patch = &v
}

func (o ModifiedPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifiedPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Patch) {
		toSerialize["patch"] = o.Patch
	}
	return toSerialize, nil
}

type NullableModifiedPackage struct {
	value *ModifiedPackage
	isSet bool
}

func (v NullableModifiedPackage) Get() *ModifiedPackage {
	return v.value
}

func (v *NullableModifiedPackage) Set(val *ModifiedPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableModifiedPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableModifiedPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifiedPackage(val *ModifiedPackage) *NullableModifiedPackage {
	return &NullableModifiedPackage{value: val, isSet: true}
}

func (v NullableModifiedPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifiedPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


