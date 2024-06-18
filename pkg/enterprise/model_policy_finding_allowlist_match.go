/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// PolicyFindingAllowlistMatch Details about (possible) allowlist match
type PolicyFindingAllowlistMatch struct {
	// ID of the allowlist that matched this finding
	Id *string `json:"id,omitempty"`
	// Name of the allowlist that matched this finding
	Name *string `json:"name,omitempty"`
	// ID of the rule within the allowlist that matched this finding
	MatchedRuleId *string `json:"matched_rule_id,omitempty"`
}

// NewPolicyFindingAllowlistMatch instantiates a new PolicyFindingAllowlistMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyFindingAllowlistMatch() *PolicyFindingAllowlistMatch {
	this := PolicyFindingAllowlistMatch{}
	return &this
}

// NewPolicyFindingAllowlistMatchWithDefaults instantiates a new PolicyFindingAllowlistMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyFindingAllowlistMatchWithDefaults() *PolicyFindingAllowlistMatch {
	this := PolicyFindingAllowlistMatch{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PolicyFindingAllowlistMatch) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFindingAllowlistMatch) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PolicyFindingAllowlistMatch) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PolicyFindingAllowlistMatch) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PolicyFindingAllowlistMatch) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFindingAllowlistMatch) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PolicyFindingAllowlistMatch) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PolicyFindingAllowlistMatch) SetName(v string) {
	o.Name = &v
}

// GetMatchedRuleId returns the MatchedRuleId field value if set, zero value otherwise.
func (o *PolicyFindingAllowlistMatch) GetMatchedRuleId() string {
	if o == nil || o.MatchedRuleId == nil {
		var ret string
		return ret
	}
	return *o.MatchedRuleId
}

// GetMatchedRuleIdOk returns a tuple with the MatchedRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFindingAllowlistMatch) GetMatchedRuleIdOk() (*string, bool) {
	if o == nil || o.MatchedRuleId == nil {
		return nil, false
	}
	return o.MatchedRuleId, true
}

// HasMatchedRuleId returns a boolean if a field has been set.
func (o *PolicyFindingAllowlistMatch) HasMatchedRuleId() bool {
	if o != nil && o.MatchedRuleId != nil {
		return true
	}

	return false
}

// SetMatchedRuleId gets a reference to the given string and assigns it to the MatchedRuleId field.
func (o *PolicyFindingAllowlistMatch) SetMatchedRuleId(v string) {
	o.MatchedRuleId = &v
}

func (o PolicyFindingAllowlistMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.MatchedRuleId != nil {
		toSerialize["matched_rule_id"] = o.MatchedRuleId
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyFindingAllowlistMatch struct {
	value *PolicyFindingAllowlistMatch
	isSet bool
}

func (v NullablePolicyFindingAllowlistMatch) Get() *PolicyFindingAllowlistMatch {
	return v.value
}

func (v *NullablePolicyFindingAllowlistMatch) Set(val *PolicyFindingAllowlistMatch) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyFindingAllowlistMatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyFindingAllowlistMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyFindingAllowlistMatch(val *PolicyFindingAllowlistMatch) *NullablePolicyFindingAllowlistMatch {
	return &NullablePolicyFindingAllowlistMatch{value: val, isSet: true}
}

func (v NullablePolicyFindingAllowlistMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyFindingAllowlistMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


