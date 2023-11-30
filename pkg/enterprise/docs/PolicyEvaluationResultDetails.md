# PolicyEvaluationResultDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | [**Policy**](Policy.md) |  | 
**Findings** | [**[]PolicyEvaluationFinding**](PolicyEvaluationFinding.md) | The detailed policy findings | 
**PolicyAction** | **string** | The outcome of the policy evaluation, before allowlist/denylist evaluation. | 
**Remediations** | Pointer to [**[]PolicyEvaluationRemediation**](PolicyEvaluationRemediation.md) | List of remediations for the findings | [optional] 

## Methods

### NewPolicyEvaluationResultDetails

`func NewPolicyEvaluationResultDetails(policy Policy, findings []PolicyEvaluationFinding, policyAction string, ) *PolicyEvaluationResultDetails`

NewPolicyEvaluationResultDetails instantiates a new PolicyEvaluationResultDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationResultDetailsWithDefaults

`func NewPolicyEvaluationResultDetailsWithDefaults() *PolicyEvaluationResultDetails`

NewPolicyEvaluationResultDetailsWithDefaults instantiates a new PolicyEvaluationResultDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *PolicyEvaluationResultDetails) GetPolicy() Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicyEvaluationResultDetails) GetPolicyOk() (*Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicyEvaluationResultDetails) SetPolicy(v Policy)`

SetPolicy sets Policy field to given value.


### GetFindings

`func (o *PolicyEvaluationResultDetails) GetFindings() []PolicyEvaluationFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *PolicyEvaluationResultDetails) GetFindingsOk() (*[]PolicyEvaluationFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *PolicyEvaluationResultDetails) SetFindings(v []PolicyEvaluationFinding)`

SetFindings sets Findings field to given value.


### GetPolicyAction

`func (o *PolicyEvaluationResultDetails) GetPolicyAction() string`

GetPolicyAction returns the PolicyAction field if non-nil, zero value otherwise.

### GetPolicyActionOk

`func (o *PolicyEvaluationResultDetails) GetPolicyActionOk() (*string, bool)`

GetPolicyActionOk returns a tuple with the PolicyAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyAction

`func (o *PolicyEvaluationResultDetails) SetPolicyAction(v string)`

SetPolicyAction sets PolicyAction field to given value.


### GetRemediations

`func (o *PolicyEvaluationResultDetails) GetRemediations() []PolicyEvaluationRemediation`

GetRemediations returns the Remediations field if non-nil, zero value otherwise.

### GetRemediationsOk

`func (o *PolicyEvaluationResultDetails) GetRemediationsOk() (*[]PolicyEvaluationRemediation, bool)`

GetRemediationsOk returns a tuple with the Remediations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediations

`func (o *PolicyEvaluationResultDetails) SetRemediations(v []PolicyEvaluationRemediation)`

SetRemediations sets Remediations field to given value.

### HasRemediations

`func (o *PolicyEvaluationResultDetails) HasRemediations() bool`

HasRemediations returns a boolean if a field has been set.

### SetRemediationsNil

`func (o *PolicyEvaluationResultDetails) SetRemediationsNil(b bool)`

 SetRemediationsNil sets the value for Remediations to be an explicit nil

### UnsetRemediations
`func (o *PolicyEvaluationResultDetails) UnsetRemediations()`

UnsetRemediations ensures that no value is present for Remediations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


