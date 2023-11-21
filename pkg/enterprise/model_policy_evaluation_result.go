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

// PolicyEvaluationResult struct for PolicyEvaluationResult
type PolicyEvaluationResult struct {
	Details NullablePolicyEvaluationResultDetails `json:"details,omitempty"`
	// Image digest of the base image used during policy evaluation
	ComparisonImageDigest NullableString `json:"comparison_image_digest,omitempty"`
	// The date and time this policy evaluation was performed at
	EvaluationTime time.Time `json:"evaluation_time"`
	// list of error objects indicating errors encountered during evaluation execution
	EvaluationProblems []PolicyEvaluationProblem `json:"evaluation_problems"`
	// The overall status of the policy evaluation
	Status string `json:"status"`
	// The overall outcome of the evaluation.
	FinalAction string `json:"final_action"`
	// The reason for the final result
	FinalActionReason string `json:"final_action_reason"`
	// Whether the evaluated image matched an allowlist rule
	ImageAllowlisted bool `json:"image_allowlisted"`
	MatchedAllowlistedImagesRule *PolicyEvaluationResultMatchedAllowlistedImagesRule `json:"matched_allowlisted_images_rule,omitempty"`
	// Whether the evaluated image matched a denylist rule
	ImageDenylisted bool `json:"image_denylisted"`
	MatchedDenylistedImagesRule *PolicyEvaluationResultMatchedAllowlistedImagesRule `json:"matched_denylisted_images_rule,omitempty"`
	// Whether the evaluated image matched a policy rule
	ImageMappedToRule bool `json:"image_mapped_to_rule"`
	MatchedMappingRule *PolicyEvaluationResultMatchedMappingRule `json:"matched_mapping_rule,omitempty"`
	// Number of policy findings in the response
	NumberOfFindings int32 `json:"number_of_findings"`
}

// NewPolicyEvaluationResult instantiates a new PolicyEvaluationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvaluationResult(evaluationTime time.Time, evaluationProblems []PolicyEvaluationProblem, status string, finalAction string, finalActionReason string, imageAllowlisted bool, imageDenylisted bool, imageMappedToRule bool, numberOfFindings int32) *PolicyEvaluationResult {
	this := PolicyEvaluationResult{}
	this.EvaluationTime = evaluationTime
	this.EvaluationProblems = evaluationProblems
	this.Status = status
	this.FinalAction = finalAction
	this.FinalActionReason = finalActionReason
	this.ImageAllowlisted = imageAllowlisted
	this.ImageDenylisted = imageDenylisted
	this.ImageMappedToRule = imageMappedToRule
	this.NumberOfFindings = numberOfFindings
	return &this
}

// NewPolicyEvaluationResultWithDefaults instantiates a new PolicyEvaluationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvaluationResultWithDefaults() *PolicyEvaluationResult {
	this := PolicyEvaluationResult{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyEvaluationResult) GetDetails() PolicyEvaluationResultDetails {
	if o == nil || o.Details.Get() == nil {
		var ret PolicyEvaluationResultDetails
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyEvaluationResult) GetDetailsOk() (*PolicyEvaluationResultDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullablePolicyEvaluationResultDetails and assigns it to the Details field.
func (o *PolicyEvaluationResult) SetDetails(v PolicyEvaluationResultDetails) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *PolicyEvaluationResult) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *PolicyEvaluationResult) UnsetDetails() {
	o.Details.Unset()
}

// GetComparisonImageDigest returns the ComparisonImageDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyEvaluationResult) GetComparisonImageDigest() string {
	if o == nil || o.ComparisonImageDigest.Get() == nil {
		var ret string
		return ret
	}
	return *o.ComparisonImageDigest.Get()
}

// GetComparisonImageDigestOk returns a tuple with the ComparisonImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyEvaluationResult) GetComparisonImageDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComparisonImageDigest.Get(), o.ComparisonImageDigest.IsSet()
}

// HasComparisonImageDigest returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasComparisonImageDigest() bool {
	if o != nil && o.ComparisonImageDigest.IsSet() {
		return true
	}

	return false
}

// SetComparisonImageDigest gets a reference to the given NullableString and assigns it to the ComparisonImageDigest field.
func (o *PolicyEvaluationResult) SetComparisonImageDigest(v string) {
	o.ComparisonImageDigest.Set(&v)
}
// SetComparisonImageDigestNil sets the value for ComparisonImageDigest to be an explicit nil
func (o *PolicyEvaluationResult) SetComparisonImageDigestNil() {
	o.ComparisonImageDigest.Set(nil)
}

// UnsetComparisonImageDigest ensures that no value is present for ComparisonImageDigest, not even an explicit nil
func (o *PolicyEvaluationResult) UnsetComparisonImageDigest() {
	o.ComparisonImageDigest.Unset()
}

// GetEvaluationTime returns the EvaluationTime field value
func (o *PolicyEvaluationResult) GetEvaluationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EvaluationTime
}

// GetEvaluationTimeOk returns a tuple with the EvaluationTime field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetEvaluationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationTime, true
}

// SetEvaluationTime sets field value
func (o *PolicyEvaluationResult) SetEvaluationTime(v time.Time) {
	o.EvaluationTime = v
}

// GetEvaluationProblems returns the EvaluationProblems field value
func (o *PolicyEvaluationResult) GetEvaluationProblems() []PolicyEvaluationProblem {
	if o == nil {
		var ret []PolicyEvaluationProblem
		return ret
	}

	return o.EvaluationProblems
}

// GetEvaluationProblemsOk returns a tuple with the EvaluationProblems field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetEvaluationProblemsOk() ([]PolicyEvaluationProblem, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvaluationProblems, true
}

// SetEvaluationProblems sets field value
func (o *PolicyEvaluationResult) SetEvaluationProblems(v []PolicyEvaluationProblem) {
	o.EvaluationProblems = v
}

// GetStatus returns the Status field value
func (o *PolicyEvaluationResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PolicyEvaluationResult) SetStatus(v string) {
	o.Status = v
}

// GetFinalAction returns the FinalAction field value
func (o *PolicyEvaluationResult) GetFinalAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinalAction
}

// GetFinalActionOk returns a tuple with the FinalAction field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetFinalActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalAction, true
}

// SetFinalAction sets field value
func (o *PolicyEvaluationResult) SetFinalAction(v string) {
	o.FinalAction = v
}

// GetFinalActionReason returns the FinalActionReason field value
func (o *PolicyEvaluationResult) GetFinalActionReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinalActionReason
}

// GetFinalActionReasonOk returns a tuple with the FinalActionReason field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetFinalActionReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalActionReason, true
}

// SetFinalActionReason sets field value
func (o *PolicyEvaluationResult) SetFinalActionReason(v string) {
	o.FinalActionReason = v
}

// GetImageAllowlisted returns the ImageAllowlisted field value
func (o *PolicyEvaluationResult) GetImageAllowlisted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ImageAllowlisted
}

// GetImageAllowlistedOk returns a tuple with the ImageAllowlisted field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetImageAllowlistedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageAllowlisted, true
}

// SetImageAllowlisted sets field value
func (o *PolicyEvaluationResult) SetImageAllowlisted(v bool) {
	o.ImageAllowlisted = v
}

// GetMatchedAllowlistedImagesRule returns the MatchedAllowlistedImagesRule field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetMatchedAllowlistedImagesRule() PolicyEvaluationResultMatchedAllowlistedImagesRule {
	if o == nil || o.MatchedAllowlistedImagesRule == nil {
		var ret PolicyEvaluationResultMatchedAllowlistedImagesRule
		return ret
	}
	return *o.MatchedAllowlistedImagesRule
}

// GetMatchedAllowlistedImagesRuleOk returns a tuple with the MatchedAllowlistedImagesRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetMatchedAllowlistedImagesRuleOk() (*PolicyEvaluationResultMatchedAllowlistedImagesRule, bool) {
	if o == nil || o.MatchedAllowlistedImagesRule == nil {
		return nil, false
	}
	return o.MatchedAllowlistedImagesRule, true
}

// HasMatchedAllowlistedImagesRule returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasMatchedAllowlistedImagesRule() bool {
	if o != nil && o.MatchedAllowlistedImagesRule != nil {
		return true
	}

	return false
}

// SetMatchedAllowlistedImagesRule gets a reference to the given PolicyEvaluationResultMatchedAllowlistedImagesRule and assigns it to the MatchedAllowlistedImagesRule field.
func (o *PolicyEvaluationResult) SetMatchedAllowlistedImagesRule(v PolicyEvaluationResultMatchedAllowlistedImagesRule) {
	o.MatchedAllowlistedImagesRule = &v
}

// GetImageDenylisted returns the ImageDenylisted field value
func (o *PolicyEvaluationResult) GetImageDenylisted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ImageDenylisted
}

// GetImageDenylistedOk returns a tuple with the ImageDenylisted field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetImageDenylistedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageDenylisted, true
}

// SetImageDenylisted sets field value
func (o *PolicyEvaluationResult) SetImageDenylisted(v bool) {
	o.ImageDenylisted = v
}

// GetMatchedDenylistedImagesRule returns the MatchedDenylistedImagesRule field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetMatchedDenylistedImagesRule() PolicyEvaluationResultMatchedAllowlistedImagesRule {
	if o == nil || o.MatchedDenylistedImagesRule == nil {
		var ret PolicyEvaluationResultMatchedAllowlistedImagesRule
		return ret
	}
	return *o.MatchedDenylistedImagesRule
}

// GetMatchedDenylistedImagesRuleOk returns a tuple with the MatchedDenylistedImagesRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetMatchedDenylistedImagesRuleOk() (*PolicyEvaluationResultMatchedAllowlistedImagesRule, bool) {
	if o == nil || o.MatchedDenylistedImagesRule == nil {
		return nil, false
	}
	return o.MatchedDenylistedImagesRule, true
}

// HasMatchedDenylistedImagesRule returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasMatchedDenylistedImagesRule() bool {
	if o != nil && o.MatchedDenylistedImagesRule != nil {
		return true
	}

	return false
}

// SetMatchedDenylistedImagesRule gets a reference to the given PolicyEvaluationResultMatchedAllowlistedImagesRule and assigns it to the MatchedDenylistedImagesRule field.
func (o *PolicyEvaluationResult) SetMatchedDenylistedImagesRule(v PolicyEvaluationResultMatchedAllowlistedImagesRule) {
	o.MatchedDenylistedImagesRule = &v
}

// GetImageMappedToRule returns the ImageMappedToRule field value
func (o *PolicyEvaluationResult) GetImageMappedToRule() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ImageMappedToRule
}

// GetImageMappedToRuleOk returns a tuple with the ImageMappedToRule field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetImageMappedToRuleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageMappedToRule, true
}

// SetImageMappedToRule sets field value
func (o *PolicyEvaluationResult) SetImageMappedToRule(v bool) {
	o.ImageMappedToRule = v
}

// GetMatchedMappingRule returns the MatchedMappingRule field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetMatchedMappingRule() PolicyEvaluationResultMatchedMappingRule {
	if o == nil || o.MatchedMappingRule == nil {
		var ret PolicyEvaluationResultMatchedMappingRule
		return ret
	}
	return *o.MatchedMappingRule
}

// GetMatchedMappingRuleOk returns a tuple with the MatchedMappingRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetMatchedMappingRuleOk() (*PolicyEvaluationResultMatchedMappingRule, bool) {
	if o == nil || o.MatchedMappingRule == nil {
		return nil, false
	}
	return o.MatchedMappingRule, true
}

// HasMatchedMappingRule returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasMatchedMappingRule() bool {
	if o != nil && o.MatchedMappingRule != nil {
		return true
	}

	return false
}

// SetMatchedMappingRule gets a reference to the given PolicyEvaluationResultMatchedMappingRule and assigns it to the MatchedMappingRule field.
func (o *PolicyEvaluationResult) SetMatchedMappingRule(v PolicyEvaluationResultMatchedMappingRule) {
	o.MatchedMappingRule = &v
}

// GetNumberOfFindings returns the NumberOfFindings field value
func (o *PolicyEvaluationResult) GetNumberOfFindings() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfFindings
}

// GetNumberOfFindingsOk returns a tuple with the NumberOfFindings field value
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetNumberOfFindingsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfFindings, true
}

// SetNumberOfFindings sets field value
func (o *PolicyEvaluationResult) SetNumberOfFindings(v int32) {
	o.NumberOfFindings = v
}

func (o PolicyEvaluationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.ComparisonImageDigest.IsSet() {
		toSerialize["comparison_image_digest"] = o.ComparisonImageDigest.Get()
	}
	if true {
		toSerialize["evaluation_time"] = o.EvaluationTime
	}
	if true {
		toSerialize["evaluation_problems"] = o.EvaluationProblems
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["final_action"] = o.FinalAction
	}
	if true {
		toSerialize["final_action_reason"] = o.FinalActionReason
	}
	if true {
		toSerialize["image_allowlisted"] = o.ImageAllowlisted
	}
	if o.MatchedAllowlistedImagesRule != nil {
		toSerialize["matched_allowlisted_images_rule"] = o.MatchedAllowlistedImagesRule
	}
	if true {
		toSerialize["image_denylisted"] = o.ImageDenylisted
	}
	if o.MatchedDenylistedImagesRule != nil {
		toSerialize["matched_denylisted_images_rule"] = o.MatchedDenylistedImagesRule
	}
	if true {
		toSerialize["image_mapped_to_rule"] = o.ImageMappedToRule
	}
	if o.MatchedMappingRule != nil {
		toSerialize["matched_mapping_rule"] = o.MatchedMappingRule
	}
	if true {
		toSerialize["number_of_findings"] = o.NumberOfFindings
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyEvaluationResult struct {
	value *PolicyEvaluationResult
	isSet bool
}

func (v NullablePolicyEvaluationResult) Get() *PolicyEvaluationResult {
	return v.value
}

func (v *NullablePolicyEvaluationResult) Set(val *PolicyEvaluationResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluationResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluationResult(val *PolicyEvaluationResult) *NullablePolicyEvaluationResult {
	return &NullablePolicyEvaluationResult{value: val, isSet: true}
}

func (v NullablePolicyEvaluationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

