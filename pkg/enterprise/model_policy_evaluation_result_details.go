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

// PolicyEvaluationResultDetails Contains additional details about the policy evaluation
type PolicyEvaluationResultDetails struct {
	Policy Policy `json:"policy"`
	// The detailed policy findings
	Findings []PolicyEvaluationFinding `json:"findings"`
	// The outcome of the policy evaluation, before allowlist/denylist evaluation.
	PolicyAction string `json:"policy_action"`
	// List of remediations for the findings
	Remediations []PolicyEvaluationRemediation `json:"remediations,omitempty"`
}

// NewPolicyEvaluationResultDetails instantiates a new PolicyEvaluationResultDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvaluationResultDetails(policy Policy, findings []PolicyEvaluationFinding, policyAction string) *PolicyEvaluationResultDetails {
	this := PolicyEvaluationResultDetails{}
	this.Policy = policy
	this.Findings = findings
	this.PolicyAction = policyAction
	return &this
}

// NewPolicyEvaluationResultDetailsWithDefaults instantiates a new PolicyEvaluationResultDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvaluationResultDetailsWithDefaults() *PolicyEvaluationResultDetails {
	this := PolicyEvaluationResultDetails{}
	return &this
}

// GetPolicy returns the Policy field value
func (o *PolicyEvaluationResultDetails) GetPolicy() Policy {
	if o == nil {
		var ret Policy
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResultDetails) GetPolicyOk() (*Policy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *PolicyEvaluationResultDetails) SetPolicy(v Policy) {
	o.Policy = v
}

// GetFindings returns the Findings field value
func (o *PolicyEvaluationResultDetails) GetFindings() []PolicyEvaluationFinding {
	if o == nil {
		var ret []PolicyEvaluationFinding
		return ret
	}

	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResultDetails) GetFindingsOk() ([]PolicyEvaluationFinding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Findings, true
}

// SetFindings sets field value
func (o *PolicyEvaluationResultDetails) SetFindings(v []PolicyEvaluationFinding) {
	o.Findings = v
}

// GetPolicyAction returns the PolicyAction field value
func (o *PolicyEvaluationResultDetails) GetPolicyAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyAction
}

// GetPolicyActionOk returns a tuple with the PolicyAction field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResultDetails) GetPolicyActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyAction, true
}

// SetPolicyAction sets field value
func (o *PolicyEvaluationResultDetails) SetPolicyAction(v string) {
	o.PolicyAction = v
}

// GetRemediations returns the Remediations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyEvaluationResultDetails) GetRemediations() []PolicyEvaluationRemediation {
	if o == nil {
		var ret []PolicyEvaluationRemediation
		return ret
	}
	return o.Remediations
}

// GetRemediationsOk returns a tuple with the Remediations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyEvaluationResultDetails) GetRemediationsOk() ([]PolicyEvaluationRemediation, bool) {
	if o == nil || o.Remediations == nil {
		return nil, false
	}
	return o.Remediations, true
}

// HasRemediations returns a boolean if a field has been set.
func (o *PolicyEvaluationResultDetails) HasRemediations() bool {
	if o != nil && o.Remediations != nil {
		return true
	}

	return false
}

// SetRemediations gets a reference to the given []PolicyEvaluationRemediation and assigns it to the Remediations field.
func (o *PolicyEvaluationResultDetails) SetRemediations(v []PolicyEvaluationRemediation) {
	o.Remediations = v
}

func (o PolicyEvaluationResultDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["policy"] = o.Policy
	}
	if true {
		toSerialize["findings"] = o.Findings
	}
	if true {
		toSerialize["policy_action"] = o.PolicyAction
	}
	if o.Remediations != nil {
		toSerialize["remediations"] = o.Remediations
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyEvaluationResultDetails struct {
	value *PolicyEvaluationResultDetails
	isSet bool
}

func (v NullablePolicyEvaluationResultDetails) Get() *PolicyEvaluationResultDetails {
	return v.value
}

func (v *NullablePolicyEvaluationResultDetails) Set(val *PolicyEvaluationResultDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluationResultDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluationResultDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluationResultDetails(val *PolicyEvaluationResultDetails) *NullablePolicyEvaluationResultDetails {
	return &NullablePolicyEvaluationResultDetails{value: val, isSet: true}
}

func (v NullablePolicyEvaluationResultDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluationResultDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


