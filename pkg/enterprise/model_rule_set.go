/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.11.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RuleSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleSet{}

// RuleSet struct for RuleSet
type RuleSet struct {
	Id string `json:"id"`
	Name string `json:"name"`
	// Description of the Policy, human readable
	Description *string `json:"description,omitempty"`
	Version string `json:"version"`
	ArtifactType *string `json:"artifact_type,omitempty"`
	Rules []PolicyRule `json:"rules"`
}

type _RuleSet RuleSet

// NewRuleSet instantiates a new RuleSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleSet(id string, name string, version string, rules []PolicyRule) *RuleSet {
	this := RuleSet{}
	this.Id = id
	this.Name = name
	this.Version = version
	this.Rules = rules
	return &this
}

// NewRuleSetWithDefaults instantiates a new RuleSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleSetWithDefaults() *RuleSet {
	this := RuleSet{}
	return &this
}

// GetId returns the Id field value
func (o *RuleSet) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RuleSet) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RuleSet) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RuleSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RuleSet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RuleSet) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RuleSet) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSet) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RuleSet) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RuleSet) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value
func (o *RuleSet) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RuleSet) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RuleSet) SetVersion(v string) {
	o.Version = v
}

// GetArtifactType returns the ArtifactType field value if set, zero value otherwise.
func (o *RuleSet) GetArtifactType() string {
	if o == nil || IsNil(o.ArtifactType) {
		var ret string
		return ret
	}
	return *o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSet) GetArtifactTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ArtifactType) {
		return nil, false
	}
	return o.ArtifactType, true
}

// HasArtifactType returns a boolean if a field has been set.
func (o *RuleSet) HasArtifactType() bool {
	if o != nil && !IsNil(o.ArtifactType) {
		return true
	}

	return false
}

// SetArtifactType gets a reference to the given string and assigns it to the ArtifactType field.
func (o *RuleSet) SetArtifactType(v string) {
	o.ArtifactType = &v
}

// GetRules returns the Rules field value
func (o *RuleSet) GetRules() []PolicyRule {
	if o == nil {
		var ret []PolicyRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *RuleSet) GetRulesOk() ([]PolicyRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *RuleSet) SetRules(v []PolicyRule) {
	o.Rules = v
}

func (o RuleSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["version"] = o.Version
	if !IsNil(o.ArtifactType) {
		toSerialize["artifact_type"] = o.ArtifactType
	}
	toSerialize["rules"] = o.Rules
	return toSerialize, nil
}

func (o *RuleSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"version",
		"rules",
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

	varRuleSet := _RuleSet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRuleSet)

	if err != nil {
		return err
	}

	*o = RuleSet(varRuleSet)

	return err
}

type NullableRuleSet struct {
	value *RuleSet
	isSet bool
}

func (v NullableRuleSet) Get() *RuleSet {
	return v.value
}

func (v *NullableRuleSet) Set(val *RuleSet) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleSet) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleSet(val *RuleSet) *NullableRuleSet {
	return &NullableRuleSet{value: val, isSet: true}
}

func (v NullableRuleSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


