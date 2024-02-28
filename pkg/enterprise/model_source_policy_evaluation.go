/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.3.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// SourcePolicyEvaluation A policy bundle evaluation result for a specific image, tag, policy tuple
type SourcePolicyEvaluation struct {
	// The name of the account containing the source to evaluate
	AccountName string `json:"account_name"`
	// The ID of this policy evaluation
	EvaluationId string `json:"evaluation_id"`
	// The ID of the source repository that was evaluated
	SourceId string `json:"source_id"`
	// Host name for the repository location (e.g. github.com)
	Host string `json:"host"`
	// The name of the repository on the host (e.g. 'anchore/anchore-engine')
	RepositoryName string `json:"repository_name"`
	// The commit ID for a git repository
	Revision string `json:"revision"`
	Policy Policy `json:"policy"`
	// Whether the evaluated source repository matched a policy rule
	SourceMappedToRule bool `json:"source_mapped_to_rule"`
	// The policy mapping rule that the source repository being evaluated matched against.
	MatchedMappingRule interface{} `json:"matched_mapping_rule,omitempty"`
	// The detailed policy findings
	Findings []SourcePolicyEvaluationFinding `json:"findings"`
	// Number of policy findings in the response
	NumberOfFindings int32 `json:"number_of_findings"`
	// The date and time this policy evaluation was performed at
	EvaluationTime time.Time `json:"evaluation_time"`
	// The overall outcome of the evaluation.
	FinalAction string `json:"final_action"`
	// The reason for the final result
	FinalActionReason string `json:"final_action_reason"`
	// list of error objects indicating errors encountered during evaluation execution
	EvaluationProblems []PolicyEvaluationProblem `json:"evaluation_problems"`
	// The overall status of the policy evaluation
	Status string `json:"status"`
}

// NewSourcePolicyEvaluation instantiates a new SourcePolicyEvaluation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourcePolicyEvaluation(accountName string, evaluationId string, sourceId string, host string, repositoryName string, revision string, policy Policy, sourceMappedToRule bool, findings []SourcePolicyEvaluationFinding, numberOfFindings int32, evaluationTime time.Time, finalAction string, finalActionReason string, evaluationProblems []PolicyEvaluationProblem, status string) *SourcePolicyEvaluation {
	this := SourcePolicyEvaluation{}
	this.AccountName = accountName
	this.EvaluationId = evaluationId
	this.SourceId = sourceId
	this.Host = host
	this.RepositoryName = repositoryName
	this.Revision = revision
	this.Policy = policy
	this.SourceMappedToRule = sourceMappedToRule
	this.Findings = findings
	this.NumberOfFindings = numberOfFindings
	this.EvaluationTime = evaluationTime
	this.FinalAction = finalAction
	this.FinalActionReason = finalActionReason
	this.EvaluationProblems = evaluationProblems
	this.Status = status
	return &this
}

// NewSourcePolicyEvaluationWithDefaults instantiates a new SourcePolicyEvaluation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourcePolicyEvaluationWithDefaults() *SourcePolicyEvaluation {
	this := SourcePolicyEvaluation{}
	return &this
}

// GetAccountName returns the AccountName field value
func (o *SourcePolicyEvaluation) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *SourcePolicyEvaluation) SetAccountName(v string) {
	o.AccountName = v
}

// GetEvaluationId returns the EvaluationId field value
func (o *SourcePolicyEvaluation) GetEvaluationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvaluationId
}

// GetEvaluationIdOk returns a tuple with the EvaluationId field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetEvaluationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationId, true
}

// SetEvaluationId sets field value
func (o *SourcePolicyEvaluation) SetEvaluationId(v string) {
	o.EvaluationId = v
}

// GetSourceId returns the SourceId field value
func (o *SourcePolicyEvaluation) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *SourcePolicyEvaluation) SetSourceId(v string) {
	o.SourceId = v
}

// GetHost returns the Host field value
func (o *SourcePolicyEvaluation) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *SourcePolicyEvaluation) SetHost(v string) {
	o.Host = v
}

// GetRepositoryName returns the RepositoryName field value
func (o *SourcePolicyEvaluation) GetRepositoryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoryName
}

// GetRepositoryNameOk returns a tuple with the RepositoryName field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetRepositoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryName, true
}

// SetRepositoryName sets field value
func (o *SourcePolicyEvaluation) SetRepositoryName(v string) {
	o.RepositoryName = v
}

// GetRevision returns the Revision field value
func (o *SourcePolicyEvaluation) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *SourcePolicyEvaluation) SetRevision(v string) {
	o.Revision = v
}

// GetPolicy returns the Policy field value
func (o *SourcePolicyEvaluation) GetPolicy() Policy {
	if o == nil {
		var ret Policy
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetPolicyOk() (*Policy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *SourcePolicyEvaluation) SetPolicy(v Policy) {
	o.Policy = v
}

// GetSourceMappedToRule returns the SourceMappedToRule field value
func (o *SourcePolicyEvaluation) GetSourceMappedToRule() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SourceMappedToRule
}

// GetSourceMappedToRuleOk returns a tuple with the SourceMappedToRule field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetSourceMappedToRuleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceMappedToRule, true
}

// SetSourceMappedToRule sets field value
func (o *SourcePolicyEvaluation) SetSourceMappedToRule(v bool) {
	o.SourceMappedToRule = v
}

// GetMatchedMappingRule returns the MatchedMappingRule field value if set, zero value otherwise.
func (o *SourcePolicyEvaluation) GetMatchedMappingRule() interface{} {
	if o == nil || o.MatchedMappingRule == nil {
		var ret interface{}
		return ret
	}
	return o.MatchedMappingRule
}

// GetMatchedMappingRuleOk returns a tuple with the MatchedMappingRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetMatchedMappingRuleOk() (interface{}, bool) {
	if o == nil || o.MatchedMappingRule == nil {
		return nil, false
	}
	return o.MatchedMappingRule, true
}

// HasMatchedMappingRule returns a boolean if a field has been set.
func (o *SourcePolicyEvaluation) HasMatchedMappingRule() bool {
	if o != nil && o.MatchedMappingRule != nil {
		return true
	}

	return false
}

// SetMatchedMappingRule gets a reference to the given interface{} and assigns it to the MatchedMappingRule field.
func (o *SourcePolicyEvaluation) SetMatchedMappingRule(v interface{}) {
	o.MatchedMappingRule = v
}

// GetFindings returns the Findings field value
func (o *SourcePolicyEvaluation) GetFindings() []SourcePolicyEvaluationFinding {
	if o == nil {
		var ret []SourcePolicyEvaluationFinding
		return ret
	}

	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetFindingsOk() ([]SourcePolicyEvaluationFinding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Findings, true
}

// SetFindings sets field value
func (o *SourcePolicyEvaluation) SetFindings(v []SourcePolicyEvaluationFinding) {
	o.Findings = v
}

// GetNumberOfFindings returns the NumberOfFindings field value
func (o *SourcePolicyEvaluation) GetNumberOfFindings() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfFindings
}

// GetNumberOfFindingsOk returns a tuple with the NumberOfFindings field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetNumberOfFindingsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfFindings, true
}

// SetNumberOfFindings sets field value
func (o *SourcePolicyEvaluation) SetNumberOfFindings(v int32) {
	o.NumberOfFindings = v
}

// GetEvaluationTime returns the EvaluationTime field value
func (o *SourcePolicyEvaluation) GetEvaluationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EvaluationTime
}

// GetEvaluationTimeOk returns a tuple with the EvaluationTime field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetEvaluationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationTime, true
}

// SetEvaluationTime sets field value
func (o *SourcePolicyEvaluation) SetEvaluationTime(v time.Time) {
	o.EvaluationTime = v
}

// GetFinalAction returns the FinalAction field value
func (o *SourcePolicyEvaluation) GetFinalAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinalAction
}

// GetFinalActionOk returns a tuple with the FinalAction field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetFinalActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalAction, true
}

// SetFinalAction sets field value
func (o *SourcePolicyEvaluation) SetFinalAction(v string) {
	o.FinalAction = v
}

// GetFinalActionReason returns the FinalActionReason field value
func (o *SourcePolicyEvaluation) GetFinalActionReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinalActionReason
}

// GetFinalActionReasonOk returns a tuple with the FinalActionReason field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetFinalActionReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalActionReason, true
}

// SetFinalActionReason sets field value
func (o *SourcePolicyEvaluation) SetFinalActionReason(v string) {
	o.FinalActionReason = v
}

// GetEvaluationProblems returns the EvaluationProblems field value
func (o *SourcePolicyEvaluation) GetEvaluationProblems() []PolicyEvaluationProblem {
	if o == nil {
		var ret []PolicyEvaluationProblem
		return ret
	}

	return o.EvaluationProblems
}

// GetEvaluationProblemsOk returns a tuple with the EvaluationProblems field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetEvaluationProblemsOk() ([]PolicyEvaluationProblem, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvaluationProblems, true
}

// SetEvaluationProblems sets field value
func (o *SourcePolicyEvaluation) SetEvaluationProblems(v []PolicyEvaluationProblem) {
	o.EvaluationProblems = v
}

// GetStatus returns the Status field value
func (o *SourcePolicyEvaluation) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SourcePolicyEvaluation) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SourcePolicyEvaluation) SetStatus(v string) {
	o.Status = v
}

func (o SourcePolicyEvaluation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_name"] = o.AccountName
	}
	if true {
		toSerialize["evaluation_id"] = o.EvaluationId
	}
	if true {
		toSerialize["source_id"] = o.SourceId
	}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["repository_name"] = o.RepositoryName
	}
	if true {
		toSerialize["revision"] = o.Revision
	}
	if true {
		toSerialize["policy"] = o.Policy
	}
	if true {
		toSerialize["source_mapped_to_rule"] = o.SourceMappedToRule
	}
	if o.MatchedMappingRule != nil {
		toSerialize["matched_mapping_rule"] = o.MatchedMappingRule
	}
	if true {
		toSerialize["findings"] = o.Findings
	}
	if true {
		toSerialize["number_of_findings"] = o.NumberOfFindings
	}
	if true {
		toSerialize["evaluation_time"] = o.EvaluationTime
	}
	if true {
		toSerialize["final_action"] = o.FinalAction
	}
	if true {
		toSerialize["final_action_reason"] = o.FinalActionReason
	}
	if true {
		toSerialize["evaluation_problems"] = o.EvaluationProblems
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSourcePolicyEvaluation struct {
	value *SourcePolicyEvaluation
	isSet bool
}

func (v NullableSourcePolicyEvaluation) Get() *SourcePolicyEvaluation {
	return v.value
}

func (v *NullableSourcePolicyEvaluation) Set(val *SourcePolicyEvaluation) {
	v.value = val
	v.isSet = true
}

func (v NullableSourcePolicyEvaluation) IsSet() bool {
	return v.isSet
}

func (v *NullableSourcePolicyEvaluation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourcePolicyEvaluation(val *SourcePolicyEvaluation) *NullableSourcePolicyEvaluation {
	return &NullableSourcePolicyEvaluation{value: val, isSet: true}
}

func (v NullableSourcePolicyEvaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourcePolicyEvaluation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


