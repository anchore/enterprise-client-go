# PolicyEvaluationProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | **string** | Severity of the policy evaluation problem. Problems with a severity of \&quot;error\&quot; prevent the policy from being evaluated, while severity \&quot;warn\&quot; indicates the policy was evaluated but the result may require additional attention. | 
**ProblemType** | **string** | the type of problem encountered, such as a misconfiguration or unavailable data | 
**Details** | **string** | Details about the problem itself and how to fix it | 

## Methods

### NewPolicyEvaluationProblem

`func NewPolicyEvaluationProblem(severity string, problemType string, details string, ) *PolicyEvaluationProblem`

NewPolicyEvaluationProblem instantiates a new PolicyEvaluationProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationProblemWithDefaults

`func NewPolicyEvaluationProblemWithDefaults() *PolicyEvaluationProblem`

NewPolicyEvaluationProblemWithDefaults instantiates a new PolicyEvaluationProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *PolicyEvaluationProblem) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *PolicyEvaluationProblem) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *PolicyEvaluationProblem) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetProblemType

`func (o *PolicyEvaluationProblem) GetProblemType() string`

GetProblemType returns the ProblemType field if non-nil, zero value otherwise.

### GetProblemTypeOk

`func (o *PolicyEvaluationProblem) GetProblemTypeOk() (*string, bool)`

GetProblemTypeOk returns a tuple with the ProblemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemType

`func (o *PolicyEvaluationProblem) SetProblemType(v string)`

SetProblemType sets ProblemType field to given value.


### GetDetails

`func (o *PolicyEvaluationProblem) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PolicyEvaluationProblem) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PolicyEvaluationProblem) SetDetails(v string)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


