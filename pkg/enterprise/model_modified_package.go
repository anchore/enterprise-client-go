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
	if o == nil || o.Source == nil {
		var ret Package
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifiedPackage) GetSourceOk() (*Package, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ModifiedPackage) HasSource() bool {
	if o != nil && o.Source != nil {
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
	if o == nil || o.Target == nil {
		var ret Package
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifiedPackage) GetTargetOk() (*Package, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ModifiedPackage) HasTarget() bool {
	if o != nil && o.Target != nil {
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
	if o == nil || o.Patch == nil {
		var ret CustomJsonPatch
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifiedPackage) GetPatchOk() (*CustomJsonPatch, bool) {
	if o == nil || o.Patch == nil {
		return nil, false
	}
	return o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *ModifiedPackage) HasPatch() bool {
	if o != nil && o.Patch != nil {
		return true
	}

	return false
}

// SetPatch gets a reference to the given CustomJsonPatch and assigns it to the Patch field.
func (o *ModifiedPackage) SetPatch(v CustomJsonPatch) {
	o.Patch = &v
}

func (o ModifiedPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Patch != nil {
		toSerialize["patch"] = o.Patch
	}
	return json.Marshal(toSerialize)
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


