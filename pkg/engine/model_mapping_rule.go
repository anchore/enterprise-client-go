/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.20
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// MappingRule struct for MappingRule
type MappingRule struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	WhitelistIds *[]string `json:"whitelist_ids,omitempty"`
	// Optional single policy to evalute, if set will override any value in policy_ids, for backwards compatibility. Generally, policy_ids should be used even with a array of length 1.
	PolicyId *string `json:"policy_id,omitempty"`
	// List of policyIds to evaluate in order, to completion
	PolicyIds *[]string `json:"policy_ids,omitempty"`
	Registry string `json:"registry"`
	Repository string `json:"repository"`
	Image ImageRef `json:"image"`
}

// NewMappingRule instantiates a new MappingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMappingRule(name string, registry string, repository string, image ImageRef) *MappingRule {
	this := MappingRule{}
	this.Name = name
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

// GetId returns the Id field value if set, zero value otherwise.
func (o *MappingRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MappingRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MappingRule) SetId(v string) {
	o.Id = &v
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
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MappingRule) SetName(v string) {
	o.Name = v
}

// GetWhitelistIds returns the WhitelistIds field value if set, zero value otherwise.
func (o *MappingRule) GetWhitelistIds() []string {
	if o == nil || o.WhitelistIds == nil {
		var ret []string
		return ret
	}
	return *o.WhitelistIds
}

// GetWhitelistIdsOk returns a tuple with the WhitelistIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingRule) GetWhitelistIdsOk() (*[]string, bool) {
	if o == nil || o.WhitelistIds == nil {
		return nil, false
	}
	return o.WhitelistIds, true
}

// HasWhitelistIds returns a boolean if a field has been set.
func (o *MappingRule) HasWhitelistIds() bool {
	if o != nil && o.WhitelistIds != nil {
		return true
	}

	return false
}

// SetWhitelistIds gets a reference to the given []string and assigns it to the WhitelistIds field.
func (o *MappingRule) SetWhitelistIds(v []string) {
	o.WhitelistIds = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *MappingRule) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingRule) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *MappingRule) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *MappingRule) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetPolicyIds returns the PolicyIds field value if set, zero value otherwise.
func (o *MappingRule) GetPolicyIds() []string {
	if o == nil || o.PolicyIds == nil {
		var ret []string
		return ret
	}
	return *o.PolicyIds
}

// GetPolicyIdsOk returns a tuple with the PolicyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingRule) GetPolicyIdsOk() (*[]string, bool) {
	if o == nil || o.PolicyIds == nil {
		return nil, false
	}
	return o.PolicyIds, true
}

// HasPolicyIds returns a boolean if a field has been set.
func (o *MappingRule) HasPolicyIds() bool {
	if o != nil && o.PolicyIds != nil {
		return true
	}

	return false
}

// SetPolicyIds gets a reference to the given []string and assigns it to the PolicyIds field.
func (o *MappingRule) SetPolicyIds(v []string) {
	o.PolicyIds = &v
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
	if o == nil  {
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
	if o == nil  {
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
	if o == nil  {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *MappingRule) SetImage(v ImageRef) {
	o.Image = v
}

func (o MappingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.WhitelistIds != nil {
		toSerialize["whitelist_ids"] = o.WhitelistIds
	}
	if o.PolicyId != nil {
		toSerialize["policy_id"] = o.PolicyId
	}
	if o.PolicyIds != nil {
		toSerialize["policy_ids"] = o.PolicyIds
	}
	if true {
		toSerialize["registry"] = o.Registry
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["image"] = o.Image
	}
	return json.Marshal(toSerialize)
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


