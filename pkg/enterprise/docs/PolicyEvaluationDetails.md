# PolicyEvaluationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | [**Policy**](Policy.md) |  | 
**Findings** | [**[]PolicyEvaluationFinding**](PolicyEvaluationFinding.md) | The detailed policy findings | 
**Remediations** | Pointer to [**[]PolicyEvaluationRemediation**](PolicyEvaluationRemediation.md) | List of remediations for the findings | [optional] 

## Methods

### NewPolicyEvaluationDetails

`func NewPolicyEvaluationDetails(policy Policy, findings []PolicyEvaluationFinding, ) *PolicyEvaluationDetails`

NewPolicyEvaluationDetails instantiates a new PolicyEvaluationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationDetailsWithDefaults

`func NewPolicyEvaluationDetailsWithDefaults() *PolicyEvaluationDetails`

NewPolicyEvaluationDetailsWithDefaults instantiates a new PolicyEvaluationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *PolicyEvaluationDetails) GetPolicy() Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicyEvaluationDetails) GetPolicyOk() (*Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicyEvaluationDetails) SetPolicy(v Policy)`

SetPolicy sets Policy field to given value.


### GetFindings

`func (o *PolicyEvaluationDetails) GetFindings() []PolicyEvaluationFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *PolicyEvaluationDetails) GetFindingsOk() (*[]PolicyEvaluationFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *PolicyEvaluationDetails) SetFindings(v []PolicyEvaluationFinding)`

SetFindings sets Findings field to given value.


### GetRemediations

`func (o *PolicyEvaluationDetails) GetRemediations() []PolicyEvaluationRemediation`

GetRemediations returns the Remediations field if non-nil, zero value otherwise.

### GetRemediationsOk

`func (o *PolicyEvaluationDetails) GetRemediationsOk() (*[]PolicyEvaluationRemediation, bool)`

GetRemediationsOk returns a tuple with the Remediations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediations

`func (o *PolicyEvaluationDetails) SetRemediations(v []PolicyEvaluationRemediation)`

SetRemediations sets Remediations field to given value.

### HasRemediations

`func (o *PolicyEvaluationDetails) HasRemediations() bool`

HasRemediations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


