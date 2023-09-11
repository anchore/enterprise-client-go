/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// PolicyRecord A policy plus some metadata
type PolicyRecord struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// The policy's identifier
	PolicyId string `json:"policy_id"`
	// True if the policy is currently defined to be used automatically
	Active bool `json:"active"`
	// UserId of the user that owns the policy
	AccountName string `json:"account_name"`
	// Source location of where the policy originated
	PolicySource string `json:"policy_source"`
	Policy NullablePolicy `json:"policy,omitempty"`
	// Name of the policy
	Name string `json:"name"`
	// Description of the policy, human readable
	Description *string `json:"description,omitempty"`
}

// NewPolicyRecord instantiates a new PolicyRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRecord(policyId string, active bool, accountName string, policySource string, name string) *PolicyRecord {
	this := PolicyRecord{}
	this.PolicyId = policyId
	this.Active = active
	this.AccountName = accountName
	this.PolicySource = policySource
	this.Name = name
	return &this
}

// NewPolicyRecordWithDefaults instantiates a new PolicyRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRecordWithDefaults() *PolicyRecord {
	this := PolicyRecord{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PolicyRecord) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRecord) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PolicyRecord) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PolicyRecord) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PolicyRecord) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRecord) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PolicyRecord) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *PolicyRecord) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetPolicyId returns the PolicyId field value
func (o *PolicyRecord) GetPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
func (o *PolicyRecord) GetPolicyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PolicyId, true
}

// SetPolicyId sets field value
func (o *PolicyRecord) SetPolicyId(v string) {
	o.PolicyId = v
}

// GetActive returns the Active field value
func (o *PolicyRecord) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *PolicyRecord) GetActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *PolicyRecord) SetActive(v bool) {
	o.Active = v
}

// GetAccountName returns the AccountName field value
func (o *PolicyRecord) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *PolicyRecord) GetAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *PolicyRecord) SetAccountName(v string) {
	o.AccountName = v
}

// GetPolicySource returns the PolicySource field value
func (o *PolicyRecord) GetPolicySource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicySource
}

// GetPolicySourceOk returns a tuple with the PolicySource field value
// and a boolean to check if the value has been set.
func (o *PolicyRecord) GetPolicySourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PolicySource, true
}

// SetPolicySource sets field value
func (o *PolicyRecord) SetPolicySource(v string) {
	o.PolicySource = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRecord) GetPolicy() Policy {
	if o == nil || o.Policy.Get() == nil {
		var ret Policy
		return ret
	}
	return *o.Policy.Get()
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRecord) GetPolicyOk() (*Policy, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Policy.Get(), o.Policy.IsSet()
}

// HasPolicy returns a boolean if a field has been set.
func (o *PolicyRecord) HasPolicy() bool {
	if o != nil && o.Policy.IsSet() {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given NullablePolicy and assigns it to the Policy field.
func (o *PolicyRecord) SetPolicy(v Policy) {
	o.Policy.Set(&v)
}
// SetPolicyNil sets the value for Policy to be an explicit nil
func (o *PolicyRecord) SetPolicyNil() {
	o.Policy.Set(nil)
}

// UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
func (o *PolicyRecord) UnsetPolicy() {
	o.Policy.Unset()
}

// GetName returns the Name field value
func (o *PolicyRecord) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PolicyRecord) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PolicyRecord) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PolicyRecord) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRecord) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyRecord) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PolicyRecord) SetDescription(v string) {
	o.Description = &v
}

func (o PolicyRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if true {
		toSerialize["policy_id"] = o.PolicyId
	}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["account_name"] = o.AccountName
	}
	if true {
		toSerialize["policy_source"] = o.PolicySource
	}
	if o.Policy.IsSet() {
		toSerialize["policy"] = o.Policy.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyRecord struct {
	value *PolicyRecord
	isSet bool
}

func (v NullablePolicyRecord) Get() *PolicyRecord {
	return v.value
}

func (v *NullablePolicyRecord) Set(val *PolicyRecord) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRecord) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRecord(val *PolicyRecord) *NullablePolicyRecord {
	return &NullablePolicyRecord{value: val, isSet: true}
}

func (v NullablePolicyRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


