/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// ArtifactLifecyclePolicy A policy which exists in the system to manage the lifecycle of artifacts
type ArtifactLifecyclePolicy struct {
	// A system defined unique identifier.
	Uuid *string `json:"uuid,omitempty"`
	// The action that should be taken when the rule parameters are met.
	Action string `json:"action"`
	// A user provided name for the policy. This name must be unique for an Artifact Lifecycle Policy.
	Name string `json:"name"`
	// A user provided description for the policy.
	Description *string `json:"description,omitempty"`
	PolicyConditions ArtifactLifecyclePolicyConditions `json:"policy_conditions"`
	// Indicates if the policy should be active or not.  Defaulted to false.
	Enabled *bool `json:"enabled,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// NewArtifactLifecyclePolicy instantiates a new ArtifactLifecyclePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactLifecyclePolicy(action string, name string, policyConditions ArtifactLifecyclePolicyConditions) *ArtifactLifecyclePolicy {
	this := ArtifactLifecyclePolicy{}
	this.Action = action
	this.Name = name
	this.PolicyConditions = policyConditions
	return &this
}

// NewArtifactLifecyclePolicyWithDefaults instantiates a new ArtifactLifecyclePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactLifecyclePolicyWithDefaults() *ArtifactLifecyclePolicy {
	this := ArtifactLifecyclePolicy{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ArtifactLifecyclePolicy) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicy) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ArtifactLifecyclePolicy) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ArtifactLifecyclePolicy) SetUuid(v string) {
	o.Uuid = &v
}

// GetAction returns the Action field value
func (o *ArtifactLifecyclePolicy) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicy) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ArtifactLifecyclePolicy) SetAction(v string) {
	o.Action = v
}

// GetName returns the Name field value
func (o *ArtifactLifecyclePolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArtifactLifecyclePolicy) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArtifactLifecyclePolicy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArtifactLifecyclePolicy) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArtifactLifecyclePolicy) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyConditions returns the PolicyConditions field value
func (o *ArtifactLifecyclePolicy) GetPolicyConditions() ArtifactLifecyclePolicyConditions {
	if o == nil {
		var ret ArtifactLifecyclePolicyConditions
		return ret
	}

	return o.PolicyConditions
}

// GetPolicyConditionsOk returns a tuple with the PolicyConditions field value
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicy) GetPolicyConditionsOk() (*ArtifactLifecyclePolicyConditions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyConditions, true
}

// SetPolicyConditions sets field value
func (o *ArtifactLifecyclePolicy) SetPolicyConditions(v ArtifactLifecyclePolicyConditions) {
	o.PolicyConditions = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ArtifactLifecyclePolicy) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ArtifactLifecyclePolicy) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ArtifactLifecyclePolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ArtifactLifecyclePolicy) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ArtifactLifecyclePolicy) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ArtifactLifecyclePolicy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ArtifactLifecyclePolicy) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ArtifactLifecyclePolicy) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ArtifactLifecyclePolicy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ArtifactLifecyclePolicy) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactLifecyclePolicy) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ArtifactLifecyclePolicy) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *ArtifactLifecyclePolicy) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

func (o ArtifactLifecyclePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["policy_conditions"] = o.PolicyConditions
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableArtifactLifecyclePolicy struct {
	value *ArtifactLifecyclePolicy
	isSet bool
}

func (v NullableArtifactLifecyclePolicy) Get() *ArtifactLifecyclePolicy {
	return v.value
}

func (v *NullableArtifactLifecyclePolicy) Set(val *ArtifactLifecyclePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactLifecyclePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactLifecyclePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactLifecyclePolicy(val *ArtifactLifecyclePolicy) *NullableArtifactLifecyclePolicy {
	return &NullableArtifactLifecyclePolicy{value: val, isSet: true}
}

func (v NullableArtifactLifecyclePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactLifecyclePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


