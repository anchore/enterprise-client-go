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
)

// PolicyEvaluationProblem Details for an error or warning indicating a problem during policy evaluation
type PolicyEvaluationProblem struct {
	// Severity of the policy evaluation problem. Problems with a severity of \"error\" prevent the policy from being evaluated, while severity \"warn\" indicates the policy was evaluated but the result may require additional attention.
	Severity string `json:"severity"`
	// the type of problem encountered, such as a misconfiguration or unavailable data
	ProblemType string `json:"problem_type"`
	// Details about the problem itself and how to fix it
	Details string `json:"details"`
}

// NewPolicyEvaluationProblem instantiates a new PolicyEvaluationProblem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvaluationProblem(severity string, problemType string, details string) *PolicyEvaluationProblem {
	this := PolicyEvaluationProblem{}
	this.Severity = severity
	this.ProblemType = problemType
	this.Details = details
	return &this
}

// NewPolicyEvaluationProblemWithDefaults instantiates a new PolicyEvaluationProblem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvaluationProblemWithDefaults() *PolicyEvaluationProblem {
	this := PolicyEvaluationProblem{}
	return &this
}

// GetSeverity returns the Severity field value
func (o *PolicyEvaluationProblem) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationProblem) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *PolicyEvaluationProblem) SetSeverity(v string) {
	o.Severity = v
}

// GetProblemType returns the ProblemType field value
func (o *PolicyEvaluationProblem) GetProblemType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProblemType
}

// GetProblemTypeOk returns a tuple with the ProblemType field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationProblem) GetProblemTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProblemType, true
}

// SetProblemType sets field value
func (o *PolicyEvaluationProblem) SetProblemType(v string) {
	o.ProblemType = v
}

// GetDetails returns the Details field value
func (o *PolicyEvaluationProblem) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationProblem) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *PolicyEvaluationProblem) SetDetails(v string) {
	o.Details = v
}

func (o PolicyEvaluationProblem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["problem_type"] = o.ProblemType
	}
	if true {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyEvaluationProblem struct {
	value *PolicyEvaluationProblem
	isSet bool
}

func (v NullablePolicyEvaluationProblem) Get() *PolicyEvaluationProblem {
	return v.value
}

func (v *NullablePolicyEvaluationProblem) Set(val *PolicyEvaluationProblem) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluationProblem) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluationProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluationProblem(val *PolicyEvaluationProblem) *NullablePolicyEvaluationProblem {
	return &NullablePolicyEvaluationProblem{value: val, isSet: true}
}

func (v NullablePolicyEvaluationProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluationProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


