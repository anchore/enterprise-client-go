/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the GateSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GateSpec{}

// GateSpec A description of the set of gates available in this engine and the triggers and parameters supported
type GateSpec struct {
	// Gate name, as it would appear in a policy document
	Name *string `json:"name,omitempty"`
	// Description of the gate
	Description *string `json:"description,omitempty"`
	SupportedArtifactType *string `json:"supported_artifact_type,omitempty"`
	// State of the gate and transitively all triggers it contains if not 'active'
	State *string `json:"state,omitempty"`
	// The name of another trigger that supersedes this on functionally if this is deprecated
	SupersededBy NullableString `json:"superseded_by,omitempty"`
	// List of the triggers that can fire for this Gate
	Triggers []TriggerSpec `json:"triggers,omitempty"`
}

// NewGateSpec instantiates a new GateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGateSpec() *GateSpec {
	this := GateSpec{}
	return &this
}

// NewGateSpecWithDefaults instantiates a new GateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGateSpecWithDefaults() *GateSpec {
	this := GateSpec{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GateSpec) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateSpec) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GateSpec) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GateSpec) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GateSpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GateSpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetSupportedArtifactType returns the SupportedArtifactType field value if set, zero value otherwise.
func (o *GateSpec) GetSupportedArtifactType() string {
	if o == nil || IsNil(o.SupportedArtifactType) {
		var ret string
		return ret
	}
	return *o.SupportedArtifactType
}

// GetSupportedArtifactTypeOk returns a tuple with the SupportedArtifactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateSpec) GetSupportedArtifactTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedArtifactType) {
		return nil, false
	}
	return o.SupportedArtifactType, true
}

// HasSupportedArtifactType returns a boolean if a field has been set.
func (o *GateSpec) HasSupportedArtifactType() bool {
	if o != nil && !IsNil(o.SupportedArtifactType) {
		return true
	}

	return false
}

// SetSupportedArtifactType gets a reference to the given string and assigns it to the SupportedArtifactType field.
func (o *GateSpec) SetSupportedArtifactType(v string) {
	o.SupportedArtifactType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GateSpec) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateSpec) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GateSpec) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GateSpec) SetState(v string) {
	o.State = &v
}

// GetSupersededBy returns the SupersededBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GateSpec) GetSupersededBy() string {
	if o == nil || IsNil(o.SupersededBy.Get()) {
		var ret string
		return ret
	}
	return *o.SupersededBy.Get()
}

// GetSupersededByOk returns a tuple with the SupersededBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GateSpec) GetSupersededByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupersededBy.Get(), o.SupersededBy.IsSet()
}

// HasSupersededBy returns a boolean if a field has been set.
func (o *GateSpec) HasSupersededBy() bool {
	if o != nil && o.SupersededBy.IsSet() {
		return true
	}

	return false
}

// SetSupersededBy gets a reference to the given NullableString and assigns it to the SupersededBy field.
func (o *GateSpec) SetSupersededBy(v string) {
	o.SupersededBy.Set(&v)
}
// SetSupersededByNil sets the value for SupersededBy to be an explicit nil
func (o *GateSpec) SetSupersededByNil() {
	o.SupersededBy.Set(nil)
}

// UnsetSupersededBy ensures that no value is present for SupersededBy, not even an explicit nil
func (o *GateSpec) UnsetSupersededBy() {
	o.SupersededBy.Unset()
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *GateSpec) GetTriggers() []TriggerSpec {
	if o == nil || IsNil(o.Triggers) {
		var ret []TriggerSpec
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateSpec) GetTriggersOk() ([]TriggerSpec, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *GateSpec) HasTriggers() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []TriggerSpec and assigns it to the Triggers field.
func (o *GateSpec) SetTriggers(v []TriggerSpec) {
	o.Triggers = v
}

func (o GateSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GateSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SupportedArtifactType) {
		toSerialize["supported_artifact_type"] = o.SupportedArtifactType
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.SupersededBy.IsSet() {
		toSerialize["superseded_by"] = o.SupersededBy.Get()
	}
	if !IsNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	return toSerialize, nil
}

type NullableGateSpec struct {
	value *GateSpec
	isSet bool
}

func (v NullableGateSpec) Get() *GateSpec {
	return v.value
}

func (v *NullableGateSpec) Set(val *GateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGateSpec(val *GateSpec) *NullableGateSpec {
	return &NullableGateSpec{value: val, isSet: true}
}

func (v NullableGateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


