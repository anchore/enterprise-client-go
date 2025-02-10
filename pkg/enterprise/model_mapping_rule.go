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
	"bytes"
	"fmt"
)

// checks if the MappingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MappingRule{}

// MappingRule struct for MappingRule
type MappingRule struct {
	Id string `json:"id"`
	Name string `json:"name"`
	AllowlistIds []string `json:"allowlist_ids"`
	// List of rule_set_ids to evaluate in order, to completion
	RuleSetIds []string `json:"rule_set_ids"`
	Registry string `json:"registry"`
	Repository string `json:"repository"`
	Image ImageRef `json:"image"`
	// Description of the image to policy mapping rule, human readable
	Description *string `json:"description,omitempty"`
}

type _MappingRule MappingRule

// NewMappingRule instantiates a new MappingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMappingRule(id string, name string, allowlistIds []string, ruleSetIds []string, registry string, repository string, image ImageRef) *MappingRule {
	this := MappingRule{}
	this.Id = id
	this.Name = name
	this.AllowlistIds = allowlistIds
	this.RuleSetIds = ruleSetIds
	this.Registry = registry
	this.Repository = repository
	this.Image = image
	return &this
}

// NewMappingRuleWithDefaults instantiates a new MappingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappingRuleWithDefaults() *MappingRule {
	this := MappingRule{}
	return &this
}

// GetId returns the Id field value
func (o *MappingRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MappingRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MappingRule) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MappingRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MappingRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MappingRule) SetName(v string) {
	o.Name = v
}

// GetAllowlistIds returns the AllowlistIds field value
func (o *MappingRule) GetAllowlistIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowlistIds
}

// GetAllowlistIdsOk returns a tuple with the AllowlistIds field value
// and a boolean to check if the value has been set.
func (o *MappingRule) GetAllowlistIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowlistIds, true
}

// SetAllowlistIds sets field value
func (o *MappingRule) SetAllowlistIds(v []string) {
	o.AllowlistIds = v
}

// GetRuleSetIds returns the RuleSetIds field value
func (o *MappingRule) GetRuleSetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RuleSetIds
}

// GetRuleSetIdsOk returns a tuple with the RuleSetIds field value
// and a boolean to check if the value has been set.
func (o *MappingRule) GetRuleSetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleSetIds, true
}

// SetRuleSetIds sets field value
func (o *MappingRule) SetRuleSetIds(v []string) {
	o.RuleSetIds = v
}

// GetRegistry returns the Registry field value
func (o *MappingRule) GetRegistry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value
// and a boolean to check if the value has been set.
func (o *MappingRule) GetRegistryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Registry, true
}

// SetRegistry sets field value
func (o *MappingRule) SetRegistry(v string) {
	o.Registry = v
}

// GetRepository returns the Repository field value
func (o *MappingRule) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *MappingRule) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *MappingRule) SetRepository(v string) {
	o.Repository = v
}

// GetImage returns the Image field value
func (o *MappingRule) GetImage() ImageRef {
	if o == nil {
		var ret ImageRef
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *MappingRule) GetImageOk() (*ImageRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *MappingRule) SetImage(v ImageRef) {
	o.Image = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MappingRule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MappingRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MappingRule) SetDescription(v string) {
	o.Description = &v
}

func (o MappingRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MappingRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["allowlist_ids"] = o.AllowlistIds
	toSerialize["rule_set_ids"] = o.RuleSetIds
	toSerialize["registry"] = o.Registry
	toSerialize["repository"] = o.Repository
	toSerialize["image"] = o.Image
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *MappingRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"allowlist_ids",
		"rule_set_ids",
		"registry",
		"repository",
		"image",
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

	varMappingRule := _MappingRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMappingRule)

	if err != nil {
		return err
	}

	*o = MappingRule(varMappingRule)

	return err
}

type NullableMappingRule struct {
	value *MappingRule
	isSet bool
}

func (v NullableMappingRule) Get() *MappingRule {
	return v.value
}

func (v *NullableMappingRule) Set(val *MappingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableMappingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableMappingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMappingRule(val *MappingRule) *NullableMappingRule {
	return &NullableMappingRule{value: val, isSet: true}
}

func (v NullableMappingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMappingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


