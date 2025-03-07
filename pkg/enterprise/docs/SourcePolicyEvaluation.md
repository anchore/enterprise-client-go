# SourcePolicyEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** | The name of the account containing the source to evaluate | 
**EvaluationId** | **string** | The ID of this policy evaluation | 
**SourceId** | **string** | The ID of the source repository that was evaluated | 
**Host** | **string** | Host name for the repository location (e.g. github.com) | 
**RepositoryName** | **string** | The name of the repository on the host (e.g. &#39;anchore/anchore-engine&#39;) | 
**Revision** | **string** | The commit ID for a git repository | 
**Policy** | [**Policy**](Policy.md) |  | 
**SourceMappedToRule** | **bool** | Whether the evaluated source repository matched a policy rule | 
**MatchedMappingRule** | Pointer to **interface{}** | The policy mapping rule that the source repository being evaluated matched against. | [optional] 
**Findings** | [**[]SourcePolicyEvaluationFinding**](SourcePolicyEvaluationFinding.md) | The detailed policy findings | 
**NumberOfFindings** | **int64** | Number of policy findings in the response | 
**EvaluationTime** | **time.Time** | The date and time this policy evaluation was performed at | 
**FinalAction** | **string** | The overall outcome of the evaluation. | 
**FinalActionReason** | **string** | The reason for the final result | 
**EvaluationProblems** | [**[]PolicyEvaluationProblem**](PolicyEvaluationProblem.md) | list of error objects indicating errors encountered during evaluation execution | 
**Status** | **string** | The overall status of the policy evaluation | 

## Methods

### NewSourcePolicyEvaluation

`func NewSourcePolicyEvaluation(accountName string, evaluationId string, sourceId string, host string, repositoryName string, revision string, policy Policy, sourceMappedToRule bool, findings []SourcePolicyEvaluationFinding, numberOfFindings int64, evaluationTime time.Time, finalAction string, finalActionReason string, evaluationProblems []PolicyEvaluationProblem, status string, ) *SourcePolicyEvaluation`

NewSourcePolicyEvaluation instantiates a new SourcePolicyEvaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcePolicyEvaluationWithDefaults

`func NewSourcePolicyEvaluationWithDefaults() *SourcePolicyEvaluation`

NewSourcePolicyEvaluationWithDefaults instantiates a new SourcePolicyEvaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *SourcePolicyEvaluation) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SourcePolicyEvaluation) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SourcePolicyEvaluation) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetEvaluationId

`func (o *SourcePolicyEvaluation) GetEvaluationId() string`

GetEvaluationId returns the EvaluationId field if non-nil, zero value otherwise.

### GetEvaluationIdOk

`func (o *SourcePolicyEvaluation) GetEvaluationIdOk() (*string, bool)`

GetEvaluationIdOk returns a tuple with the EvaluationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationId

`func (o *SourcePolicyEvaluation) SetEvaluationId(v string)`

SetEvaluationId sets EvaluationId field to given value.


### GetSourceId

`func (o *SourcePolicyEvaluation) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SourcePolicyEvaluation) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SourcePolicyEvaluation) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetHost

`func (o *SourcePolicyEvaluation) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SourcePolicyEvaluation) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SourcePolicyEvaluation) SetHost(v string)`

SetHost sets Host field to given value.


### GetRepositoryName

`func (o *SourcePolicyEvaluation) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *SourcePolicyEvaluation) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *SourcePolicyEvaluation) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.


### GetRevision

`func (o *SourcePolicyEvaluation) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SourcePolicyEvaluation) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SourcePolicyEvaluation) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetPolicy

`func (o *SourcePolicyEvaluation) GetPolicy() Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SourcePolicyEvaluation) GetPolicyOk() (*Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SourcePolicyEvaluation) SetPolicy(v Policy)`

SetPolicy sets Policy field to given value.


### GetSourceMappedToRule

`func (o *SourcePolicyEvaluation) GetSourceMappedToRule() bool`

GetSourceMappedToRule returns the SourceMappedToRule field if non-nil, zero value otherwise.

### GetSourceMappedToRuleOk

`func (o *SourcePolicyEvaluation) GetSourceMappedToRuleOk() (*bool, bool)`

GetSourceMappedToRuleOk returns a tuple with the SourceMappedToRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMappedToRule

`func (o *SourcePolicyEvaluation) SetSourceMappedToRule(v bool)`

SetSourceMappedToRule sets SourceMappedToRule field to given value.


### GetMatchedMappingRule

`func (o *SourcePolicyEvaluation) GetMatchedMappingRule() interface{}`

GetMatchedMappingRule returns the MatchedMappingRule field if non-nil, zero value otherwise.

### GetMatchedMappingRuleOk

`func (o *SourcePolicyEvaluation) GetMatchedMappingRuleOk() (*interface{}, bool)`

GetMatchedMappingRuleOk returns a tuple with the MatchedMappingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedMappingRule

`func (o *SourcePolicyEvaluation) SetMatchedMappingRule(v interface{})`

SetMatchedMappingRule sets MatchedMappingRule field to given value.

### HasMatchedMappingRule

`func (o *SourcePolicyEvaluation) HasMatchedMappingRule() bool`

HasMatchedMappingRule returns a boolean if a field has been set.

### GetFindings

`func (o *SourcePolicyEvaluation) GetFindings() []SourcePolicyEvaluationFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *SourcePolicyEvaluation) GetFindingsOk() (*[]SourcePolicyEvaluationFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *SourcePolicyEvaluation) SetFindings(v []SourcePolicyEvaluationFinding)`

SetFindings sets Findings field to given value.


### GetNumberOfFindings

`func (o *SourcePolicyEvaluation) GetNumberOfFindings() int64`

GetNumberOfFindings returns the NumberOfFindings field if non-nil, zero value otherwise.

### GetNumberOfFindingsOk

`func (o *SourcePolicyEvaluation) GetNumberOfFindingsOk() (*int64, bool)`

GetNumberOfFindingsOk returns a tuple with the NumberOfFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFindings

`func (o *SourcePolicyEvaluation) SetNumberOfFindings(v int64)`

SetNumberOfFindings sets NumberOfFindings field to given value.


### GetEvaluationTime

`func (o *SourcePolicyEvaluation) GetEvaluationTime() time.Time`

GetEvaluationTime returns the EvaluationTime field if non-nil, zero value otherwise.

### GetEvaluationTimeOk

`func (o *SourcePolicyEvaluation) GetEvaluationTimeOk() (*time.Time, bool)`

GetEvaluationTimeOk returns a tuple with the EvaluationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTime

`func (o *SourcePolicyEvaluation) SetEvaluationTime(v time.Time)`

SetEvaluationTime sets EvaluationTime field to given value.


### GetFinalAction

`func (o *SourcePolicyEvaluation) GetFinalAction() string`

GetFinalAction returns the FinalAction field if non-nil, zero value otherwise.

### GetFinalActionOk

`func (o *SourcePolicyEvaluation) GetFinalActionOk() (*string, bool)`

GetFinalActionOk returns a tuple with the FinalAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalAction

`func (o *SourcePolicyEvaluation) SetFinalAction(v string)`

SetFinalAction sets FinalAction field to given value.


### GetFinalActionReason

`func (o *SourcePolicyEvaluation) GetFinalActionReason() string`

GetFinalActionReason returns the FinalActionReason field if non-nil, zero value otherwise.

### GetFinalActionReasonOk

`func (o *SourcePolicyEvaluation) GetFinalActionReasonOk() (*string, bool)`

GetFinalActionReasonOk returns a tuple with the FinalActionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalActionReason

`func (o *SourcePolicyEvaluation) SetFinalActionReason(v string)`

SetFinalActionReason sets FinalActionReason field to given value.


### GetEvaluationProblems

`func (o *SourcePolicyEvaluation) GetEvaluationProblems() []PolicyEvaluationProblem`

GetEvaluationProblems returns the EvaluationProblems field if non-nil, zero value otherwise.

### GetEvaluationProblemsOk

`func (o *SourcePolicyEvaluation) GetEvaluationProblemsOk() (*[]PolicyEvaluationProblem, bool)`

GetEvaluationProblemsOk returns a tuple with the EvaluationProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationProblems

`func (o *SourcePolicyEvaluation) SetEvaluationProblems(v []PolicyEvaluationProblem)`

SetEvaluationProblems sets EvaluationProblems field to given value.


### GetStatus

`func (o *SourcePolicyEvaluation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourcePolicyEvaluation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourcePolicyEvaluation) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


