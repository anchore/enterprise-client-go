# PolicyEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | Pointer to **string** | The ID of the policy used to evaluate the image | [optional] 
**ImageDigest** | Pointer to **string** | Image digest of the image being evaluated | [optional] 
**EvaluatedTag** | Pointer to **string** | Image tag used to evaluate policy for the given image | [optional] 
<<<<<<< HEAD
**Evaluations** | Pointer to [**[]PolicyEvaluationEvaluationsInner**](PolicyEvaluationEvaluationsInner.md) | List of policy evaluations. Always has at least one result, may contain multiple when the evaluation history is requested. | [optional] 
=======
**Evaluations** | Pointer to [**[]PolicyEvaluationEvaluations**](PolicyEvaluationEvaluations.md) | List of policy evaluations. Always has at least one result, may contain multiple when the evaluation history is requested. | [optional] 
>>>>>>> main

## Methods

### NewPolicyEvaluation

`func NewPolicyEvaluation() *PolicyEvaluation`

NewPolicyEvaluation instantiates a new PolicyEvaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationWithDefaults

`func NewPolicyEvaluationWithDefaults() *PolicyEvaluation`

NewPolicyEvaluationWithDefaults instantiates a new PolicyEvaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *PolicyEvaluation) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyEvaluation) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyEvaluation) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyEvaluation) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetImageDigest

`func (o *PolicyEvaluation) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *PolicyEvaluation) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *PolicyEvaluation) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *PolicyEvaluation) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetEvaluatedTag

`func (o *PolicyEvaluation) GetEvaluatedTag() string`

GetEvaluatedTag returns the EvaluatedTag field if non-nil, zero value otherwise.

### GetEvaluatedTagOk

`func (o *PolicyEvaluation) GetEvaluatedTagOk() (*string, bool)`

GetEvaluatedTagOk returns a tuple with the EvaluatedTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedTag

`func (o *PolicyEvaluation) SetEvaluatedTag(v string)`

SetEvaluatedTag sets EvaluatedTag field to given value.

### HasEvaluatedTag

`func (o *PolicyEvaluation) HasEvaluatedTag() bool`

HasEvaluatedTag returns a boolean if a field has been set.

### GetEvaluations

<<<<<<< HEAD
`func (o *PolicyEvaluation) GetEvaluations() []PolicyEvaluationEvaluationsInner`
=======
`func (o *PolicyEvaluation) GetEvaluations() []PolicyEvaluationEvaluations`
>>>>>>> main

GetEvaluations returns the Evaluations field if non-nil, zero value otherwise.

### GetEvaluationsOk

<<<<<<< HEAD
`func (o *PolicyEvaluation) GetEvaluationsOk() (*[]PolicyEvaluationEvaluationsInner, bool)`
=======
`func (o *PolicyEvaluation) GetEvaluationsOk() (*[]PolicyEvaluationEvaluations, bool)`
>>>>>>> main

GetEvaluationsOk returns a tuple with the Evaluations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluations

<<<<<<< HEAD
`func (o *PolicyEvaluation) SetEvaluations(v []PolicyEvaluationEvaluationsInner)`
=======
`func (o *PolicyEvaluation) SetEvaluations(v []PolicyEvaluationEvaluations)`
>>>>>>> main

SetEvaluations sets Evaluations field to given value.

### HasEvaluations

`func (o *PolicyEvaluation) HasEvaluations() bool`

HasEvaluations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


