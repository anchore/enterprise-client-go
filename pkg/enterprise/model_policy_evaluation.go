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
)

// checks if the PolicyEvaluation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyEvaluation{}

// PolicyEvaluation Evaluation response object
type PolicyEvaluation struct {
	// The ID of the policy used to evaluate the image
	PolicyId *string `json:"policy_id,omitempty"`
	// Image digest of the image being evaluated
	ImageDigest *string `json:"image_digest,omitempty"`
	// Image tag used to evaluate policy for the given image
	EvaluatedTag *string `json:"evaluated_tag,omitempty"`
	// List of policy evaluations. Always has at least one result, may contain multiple when the evaluation history is requested.
	Evaluations []PolicyEvaluationResult `json:"evaluations,omitempty"`
}

// NewPolicyEvaluation instantiates a new PolicyEvaluation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvaluation() *PolicyEvaluation {
	this := PolicyEvaluation{}
	return &this
}

// NewPolicyEvaluationWithDefaults instantiates a new PolicyEvaluation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvaluationWithDefaults() *PolicyEvaluation {
	this := PolicyEvaluation{}
	return &this
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *PolicyEvaluation) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluation) GetPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *PolicyEvaluation) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *PolicyEvaluation) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *PolicyEvaluation) GetImageDigest() string {
	if o == nil || IsNil(o.ImageDigest) {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluation) GetImageDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ImageDigest) {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *PolicyEvaluation) HasImageDigest() bool {
	if o != nil && !IsNil(o.ImageDigest) {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *PolicyEvaluation) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetEvaluatedTag returns the EvaluatedTag field value if set, zero value otherwise.
func (o *PolicyEvaluation) GetEvaluatedTag() string {
	if o == nil || IsNil(o.EvaluatedTag) {
		var ret string
		return ret
	}
	return *o.EvaluatedTag
}

// GetEvaluatedTagOk returns a tuple with the EvaluatedTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluation) GetEvaluatedTagOk() (*string, bool) {
	if o == nil || IsNil(o.EvaluatedTag) {
		return nil, false
	}
	return o.EvaluatedTag, true
}

// HasEvaluatedTag returns a boolean if a field has been set.
func (o *PolicyEvaluation) HasEvaluatedTag() bool {
	if o != nil && !IsNil(o.EvaluatedTag) {
		return true
	}

	return false
}

// SetEvaluatedTag gets a reference to the given string and assigns it to the EvaluatedTag field.
func (o *PolicyEvaluation) SetEvaluatedTag(v string) {
	o.EvaluatedTag = &v
}

// GetEvaluations returns the Evaluations field value if set, zero value otherwise.
func (o *PolicyEvaluation) GetEvaluations() []PolicyEvaluationResult {
	if o == nil || IsNil(o.Evaluations) {
		var ret []PolicyEvaluationResult
		return ret
	}
	return o.Evaluations
}

// GetEvaluationsOk returns a tuple with the Evaluations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluation) GetEvaluationsOk() ([]PolicyEvaluationResult, bool) {
	if o == nil || IsNil(o.Evaluations) {
		return nil, false
	}
	return o.Evaluations, true
}

// HasEvaluations returns a boolean if a field has been set.
func (o *PolicyEvaluation) HasEvaluations() bool {
	if o != nil && !IsNil(o.Evaluations) {
		return true
	}

	return false
}

// SetEvaluations gets a reference to the given []PolicyEvaluationResult and assigns it to the Evaluations field.
func (o *PolicyEvaluation) SetEvaluations(v []PolicyEvaluationResult) {
	o.Evaluations = v
}

func (o PolicyEvaluation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyEvaluation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyId) {
		toSerialize["policy_id"] = o.PolicyId
	}
	if !IsNil(o.ImageDigest) {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if !IsNil(o.EvaluatedTag) {
		toSerialize["evaluated_tag"] = o.EvaluatedTag
	}
	if !IsNil(o.Evaluations) {
		toSerialize["evaluations"] = o.Evaluations
	}
	return toSerialize, nil
}

type NullablePolicyEvaluation struct {
	value *PolicyEvaluation
	isSet bool
}

func (v NullablePolicyEvaluation) Get() *PolicyEvaluation {
	return v.value
}

func (v *NullablePolicyEvaluation) Set(val *PolicyEvaluation) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluation) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluation(val *PolicyEvaluation) *NullablePolicyEvaluation {
	return &NullablePolicyEvaluation{value: val, isSet: true}
}

func (v NullablePolicyEvaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


