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
	"bytes"
	"fmt"
)

// checks if the Policy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Policy{}

// Policy A policy containing a rule-set, allowlists, and rules for mapping them to specific images
type Policy struct {
	// Id of the policy
	Id string `json:"id"`
	// Human readable name for the policy
	Name string `json:"name"`
	// Description of the policy, human readable
	Description *string `json:"description,omitempty"`
	// Version id for this policy format
	Version string `json:"version"`
	// Allowlists which define which policy matches to disregard explicitly in the final policy decision
	Allowlists []Allowlist `json:"allowlists,omitempty"`
	// Collections of policy rules which define the go/stop/warn status of an image using rule matches on image properties
	RuleSets []RuleSet `json:"rule_sets"`
	// Mapping rules for defining which policy and allowlist(s) to apply to a source based on a match of the host and repo name. Evaluated in order.
	SourceMappings []SourceMappingRule `json:"source_mappings,omitempty"`
	// Mapping rules for defining which policy and allowlist(s) to apply to an image based on a match of the image tag or id. Evaluated in order.
	Mappings []MappingRule `json:"mappings"`
	// List of mapping rules that define which images should always be passed (unless also on the denylist), regardless of policy result.
	AllowlistedImages []ImageSelectionRule `json:"allowlisted_images,omitempty"`
	// List of mapping rules that define which images should always result in a STOP/FAIL policy result regardless of policy content or presence in allowlisted_images
	DenylistedImages []ImageSelectionRule `json:"denylisted_images,omitempty"`
	// The time at which the policy was last updated, informational only
	LastUpdated *float32 `json:"last_updated,omitempty"`
}

type _Policy Policy

// NewPolicy instantiates a new Policy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicy(id string, name string, version string, ruleSets []RuleSet, mappings []MappingRule) *Policy {
	this := Policy{}
	this.Id = id
	this.Name = name
	this.Version = version
	this.RuleSets = ruleSets
	this.Mappings = mappings
	return &this
}

// NewPolicyWithDefaults instantiates a new Policy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyWithDefaults() *Policy {
	this := Policy{}
	return &this
}

// GetId returns the Id field value
func (o *Policy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Policy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Policy) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Policy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Policy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Policy) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Policy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Policy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Policy) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value
func (o *Policy) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Policy) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Policy) SetVersion(v string) {
	o.Version = v
}

// GetAllowlists returns the Allowlists field value if set, zero value otherwise.
func (o *Policy) GetAllowlists() []Allowlist {
	if o == nil || IsNil(o.Allowlists) {
		var ret []Allowlist
		return ret
	}
	return o.Allowlists
}

// GetAllowlistsOk returns a tuple with the Allowlists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetAllowlistsOk() ([]Allowlist, bool) {
	if o == nil || IsNil(o.Allowlists) {
		return nil, false
	}
	return o.Allowlists, true
}

// HasAllowlists returns a boolean if a field has been set.
func (o *Policy) HasAllowlists() bool {
	if o != nil && !IsNil(o.Allowlists) {
		return true
	}

	return false
}

// SetAllowlists gets a reference to the given []Allowlist and assigns it to the Allowlists field.
func (o *Policy) SetAllowlists(v []Allowlist) {
	o.Allowlists = v
}

// GetRuleSets returns the RuleSets field value
func (o *Policy) GetRuleSets() []RuleSet {
	if o == nil {
		var ret []RuleSet
		return ret
	}

	return o.RuleSets
}

// GetRuleSetsOk returns a tuple with the RuleSets field value
// and a boolean to check if the value has been set.
func (o *Policy) GetRuleSetsOk() ([]RuleSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleSets, true
}

// SetRuleSets sets field value
func (o *Policy) SetRuleSets(v []RuleSet) {
	o.RuleSets = v
}

// GetSourceMappings returns the SourceMappings field value if set, zero value otherwise.
func (o *Policy) GetSourceMappings() []SourceMappingRule {
	if o == nil || IsNil(o.SourceMappings) {
		var ret []SourceMappingRule
		return ret
	}
	return o.SourceMappings
}

// GetSourceMappingsOk returns a tuple with the SourceMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetSourceMappingsOk() ([]SourceMappingRule, bool) {
	if o == nil || IsNil(o.SourceMappings) {
		return nil, false
	}
	return o.SourceMappings, true
}

// HasSourceMappings returns a boolean if a field has been set.
func (o *Policy) HasSourceMappings() bool {
	if o != nil && !IsNil(o.SourceMappings) {
		return true
	}

	return false
}

// SetSourceMappings gets a reference to the given []SourceMappingRule and assigns it to the SourceMappings field.
func (o *Policy) SetSourceMappings(v []SourceMappingRule) {
	o.SourceMappings = v
}

// GetMappings returns the Mappings field value
func (o *Policy) GetMappings() []MappingRule {
	if o == nil {
		var ret []MappingRule
		return ret
	}

	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value
// and a boolean to check if the value has been set.
func (o *Policy) GetMappingsOk() ([]MappingRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mappings, true
}

// SetMappings sets field value
func (o *Policy) SetMappings(v []MappingRule) {
	o.Mappings = v
}

// GetAllowlistedImages returns the AllowlistedImages field value if set, zero value otherwise.
func (o *Policy) GetAllowlistedImages() []ImageSelectionRule {
	if o == nil || IsNil(o.AllowlistedImages) {
		var ret []ImageSelectionRule
		return ret
	}
	return o.AllowlistedImages
}

// GetAllowlistedImagesOk returns a tuple with the AllowlistedImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetAllowlistedImagesOk() ([]ImageSelectionRule, bool) {
	if o == nil || IsNil(o.AllowlistedImages) {
		return nil, false
	}
	return o.AllowlistedImages, true
}

// HasAllowlistedImages returns a boolean if a field has been set.
func (o *Policy) HasAllowlistedImages() bool {
	if o != nil && !IsNil(o.AllowlistedImages) {
		return true
	}

	return false
}

// SetAllowlistedImages gets a reference to the given []ImageSelectionRule and assigns it to the AllowlistedImages field.
func (o *Policy) SetAllowlistedImages(v []ImageSelectionRule) {
	o.AllowlistedImages = v
}

// GetDenylistedImages returns the DenylistedImages field value if set, zero value otherwise.
func (o *Policy) GetDenylistedImages() []ImageSelectionRule {
	if o == nil || IsNil(o.DenylistedImages) {
		var ret []ImageSelectionRule
		return ret
	}
	return o.DenylistedImages
}

// GetDenylistedImagesOk returns a tuple with the DenylistedImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetDenylistedImagesOk() ([]ImageSelectionRule, bool) {
	if o == nil || IsNil(o.DenylistedImages) {
		return nil, false
	}
	return o.DenylistedImages, true
}

// HasDenylistedImages returns a boolean if a field has been set.
func (o *Policy) HasDenylistedImages() bool {
	if o != nil && !IsNil(o.DenylistedImages) {
		return true
	}

	return false
}

// SetDenylistedImages gets a reference to the given []ImageSelectionRule and assigns it to the DenylistedImages field.
func (o *Policy) SetDenylistedImages(v []ImageSelectionRule) {
	o.DenylistedImages = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Policy) GetLastUpdated() float32 {
	if o == nil || IsNil(o.LastUpdated) {
		var ret float32
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetLastUpdatedOk() (*float32, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Policy) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given float32 and assigns it to the LastUpdated field.
func (o *Policy) SetLastUpdated(v float32) {
	o.LastUpdated = &v
}

func (o Policy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Policy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["version"] = o.Version
	if !IsNil(o.Allowlists) {
		toSerialize["allowlists"] = o.Allowlists
	}
	toSerialize["rule_sets"] = o.RuleSets
	if !IsNil(o.SourceMappings) {
		toSerialize["source_mappings"] = o.SourceMappings
	}
	toSerialize["mappings"] = o.Mappings
	if !IsNil(o.AllowlistedImages) {
		toSerialize["allowlisted_images"] = o.AllowlistedImages
	}
	if !IsNil(o.DenylistedImages) {
		toSerialize["denylisted_images"] = o.DenylistedImages
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return toSerialize, nil
}

func (o *Policy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"version",
		"rule_sets",
		"mappings",
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

	varPolicy := _Policy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolicy)

	if err != nil {
		return err
	}

	*o = Policy(varPolicy)

	return err
}

type NullablePolicy struct {
	value *Policy
	isSet bool
}

func (v NullablePolicy) Get() *Policy {
	return v.value
}

func (v *NullablePolicy) Set(val *Policy) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicy(val *Policy) *NullablePolicy {
	return &NullablePolicy{value: val, isSet: true}
}

func (v NullablePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


