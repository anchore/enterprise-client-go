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

// checks if the SourceMappingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceMappingRule{}

// SourceMappingRule struct for SourceMappingRule
type SourceMappingRule struct {
	Id string `json:"id"`
	Name string `json:"name"`
	AllowlistIds []string `json:"allowlist_ids"`
	// List of rule_set_ids to evaluate in order, to completion
	RuleSetIds []string `json:"rule_set_ids"`
	Host string `json:"host"`
	Repository string `json:"repository"`
	// Description of the source to policy rule, human readable
	Description *string `json:"description,omitempty"`
}

type _SourceMappingRule SourceMappingRule

// NewSourceMappingRule instantiates a new SourceMappingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceMappingRule(id string, name string, allowlistIds []string, ruleSetIds []string, host string, repository string) *SourceMappingRule {
	this := SourceMappingRule{}
	this.Id = id
	this.Name = name
	this.AllowlistIds = allowlistIds
	this.RuleSetIds = ruleSetIds
	this.Host = host
	this.Repository = repository
	return &this
}

// NewSourceMappingRuleWithDefaults instantiates a new SourceMappingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceMappingRuleWithDefaults() *SourceMappingRule {
	this := SourceMappingRule{}
	return &this
}

// GetId returns the Id field value
func (o *SourceMappingRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceMappingRule) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SourceMappingRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceMappingRule) SetName(v string) {
	o.Name = v
}

// GetAllowlistIds returns the AllowlistIds field value
func (o *SourceMappingRule) GetAllowlistIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowlistIds
}

// GetAllowlistIdsOk returns a tuple with the AllowlistIds field value
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetAllowlistIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowlistIds, true
}

// SetAllowlistIds sets field value
func (o *SourceMappingRule) SetAllowlistIds(v []string) {
	o.AllowlistIds = v
}

// GetRuleSetIds returns the RuleSetIds field value
func (o *SourceMappingRule) GetRuleSetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RuleSetIds
}

// GetRuleSetIdsOk returns a tuple with the RuleSetIds field value
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetRuleSetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleSetIds, true
}

// SetRuleSetIds sets field value
func (o *SourceMappingRule) SetRuleSetIds(v []string) {
	o.RuleSetIds = v
}

// GetHost returns the Host field value
func (o *SourceMappingRule) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *SourceMappingRule) SetHost(v string) {
	o.Host = v
}

// GetRepository returns the Repository field value
func (o *SourceMappingRule) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *SourceMappingRule) SetRepository(v string) {
	o.Repository = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SourceMappingRule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceMappingRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SourceMappingRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SourceMappingRule) SetDescription(v string) {
	o.Description = &v
}

func (o SourceMappingRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceMappingRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["allowlist_ids"] = o.AllowlistIds
	toSerialize["rule_set_ids"] = o.RuleSetIds
	toSerialize["host"] = o.Host
	toSerialize["repository"] = o.Repository
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *SourceMappingRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"allowlist_ids",
		"rule_set_ids",
		"host",
		"repository",
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

	varSourceMappingRule := _SourceMappingRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSourceMappingRule)

	if err != nil {
		return err
	}

	*o = SourceMappingRule(varSourceMappingRule)

	return err
}

type NullableSourceMappingRule struct {
	value *SourceMappingRule
	isSet bool
}

func (v NullableSourceMappingRule) Get() *SourceMappingRule {
	return v.value
}

func (v *NullableSourceMappingRule) Set(val *SourceMappingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceMappingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceMappingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceMappingRule(val *SourceMappingRule) *NullableSourceMappingRule {
	return &NullableSourceMappingRule{value: val, isSet: true}
}

func (v NullableSourceMappingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceMappingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


