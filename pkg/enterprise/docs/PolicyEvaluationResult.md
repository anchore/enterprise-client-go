# PolicyEvaluationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**NullablePolicyEvaluationResultDetails**](PolicyEvaluationResultDetails.md) |  | [optional] 
**ComparisonImageDigest** | Pointer to **NullableString** | Image digest of the base image used during policy evaluation | [optional] 
**EvaluationTime** | **time.Time** | The date and time this policy evaluation was performed at | 
**EvaluationProblems** | [**[]PolicyEvaluationProblem**](PolicyEvaluationProblem.md) | list of error objects indicating errors encountered during evaluation execution | 
**Status** | **string** | The overall status of the policy evaluation | 
**FinalAction** | **string** | The overall outcome of the evaluation. | 
**FinalActionReason** | **string** | The reason for the final result | 
**ImageAllowlisted** | **bool** | Whether the evaluated image matched an allowlist rule | 
**MatchedAllowlistedImagesRule** | Pointer to [**PolicyEvaluationResultMatchedAllowlistedImagesRule**](PolicyEvaluationResultMatchedAllowlistedImagesRule.md) |  | [optional] 
**ImageDenylisted** | **bool** | Whether the evaluated image matched a denylist rule | 
**MatchedDenylistedImagesRule** | Pointer to [**PolicyEvaluationResultMatchedAllowlistedImagesRule**](PolicyEvaluationResultMatchedAllowlistedImagesRule.md) |  | [optional] 
**ImageMappedToRule** | **bool** | Whether the evaluated image matched a policy rule | 
**MatchedMappingRule** | Pointer to [**PolicyEvaluationResultMatchedMappingRule**](PolicyEvaluationResultMatchedMappingRule.md) |  | [optional] 
**NumberOfFindings** | **int32** | Number of policy findings in the response | 

## Methods

### NewPolicyEvaluationResult

`func NewPolicyEvaluationResult(evaluationTime time.Time, evaluationProblems []PolicyEvaluationProblem, status string, finalAction string, finalActionReason string, imageAllowlisted bool, imageDenylisted bool, imageMappedToRule bool, numberOfFindings int32, ) *PolicyEvaluationResult`

NewPolicyEvaluationResult instantiates a new PolicyEvaluationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationResultWithDefaults

`func NewPolicyEvaluationResultWithDefaults() *PolicyEvaluationResult`

NewPolicyEvaluationResultWithDefaults instantiates a new PolicyEvaluationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *PolicyEvaluationResult) GetDetails() PolicyEvaluationResultDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PolicyEvaluationResult) GetDetailsOk() (*PolicyEvaluationResultDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PolicyEvaluationResult) SetDetails(v PolicyEvaluationResultDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PolicyEvaluationResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *PolicyEvaluationResult) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *PolicyEvaluationResult) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetComparisonImageDigest

`func (o *PolicyEvaluationResult) GetComparisonImageDigest() string`

GetComparisonImageDigest returns the ComparisonImageDigest field if non-nil, zero value otherwise.

### GetComparisonImageDigestOk

`func (o *PolicyEvaluationResult) GetComparisonImageDigestOk() (*string, bool)`

GetComparisonImageDigestOk returns a tuple with the ComparisonImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonImageDigest

`func (o *PolicyEvaluationResult) SetComparisonImageDigest(v string)`

SetComparisonImageDigest sets ComparisonImageDigest field to given value.

### HasComparisonImageDigest

`func (o *PolicyEvaluationResult) HasComparisonImageDigest() bool`

HasComparisonImageDigest returns a boolean if a field has been set.

### SetComparisonImageDigestNil

`func (o *PolicyEvaluationResult) SetComparisonImageDigestNil(b bool)`

 SetComparisonImageDigestNil sets the value for ComparisonImageDigest to be an explicit nil

### UnsetComparisonImageDigest
`func (o *PolicyEvaluationResult) UnsetComparisonImageDigest()`

UnsetComparisonImageDigest ensures that no value is present for ComparisonImageDigest, not even an explicit nil
### GetEvaluationTime

`func (o *PolicyEvaluationResult) GetEvaluationTime() time.Time`

GetEvaluationTime returns the EvaluationTime field if non-nil, zero value otherwise.

### GetEvaluationTimeOk

`func (o *PolicyEvaluationResult) GetEvaluationTimeOk() (*time.Time, bool)`

GetEvaluationTimeOk returns a tuple with the EvaluationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTime

`func (o *PolicyEvaluationResult) SetEvaluationTime(v time.Time)`

SetEvaluationTime sets EvaluationTime field to given value.


### GetEvaluationProblems

`func (o *PolicyEvaluationResult) GetEvaluationProblems() []PolicyEvaluationProblem`

GetEvaluationProblems returns the EvaluationProblems field if non-nil, zero value otherwise.

### GetEvaluationProblemsOk

`func (o *PolicyEvaluationResult) GetEvaluationProblemsOk() (*[]PolicyEvaluationProblem, bool)`

GetEvaluationProblemsOk returns a tuple with the EvaluationProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationProblems

`func (o *PolicyEvaluationResult) SetEvaluationProblems(v []PolicyEvaluationProblem)`

SetEvaluationProblems sets EvaluationProblems field to given value.


### GetStatus

`func (o *PolicyEvaluationResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PolicyEvaluationResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PolicyEvaluationResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFinalAction

`func (o *PolicyEvaluationResult) GetFinalAction() string`

GetFinalAction returns the FinalAction field if non-nil, zero value otherwise.

### GetFinalActionOk

`func (o *PolicyEvaluationResult) GetFinalActionOk() (*string, bool)`

GetFinalActionOk returns a tuple with the FinalAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalAction

`func (o *PolicyEvaluationResult) SetFinalAction(v string)`

SetFinalAction sets FinalAction field to given value.


### GetFinalActionReason

`func (o *PolicyEvaluationResult) GetFinalActionReason() string`

GetFinalActionReason returns the FinalActionReason field if non-nil, zero value otherwise.

### GetFinalActionReasonOk

`func (o *PolicyEvaluationResult) GetFinalActionReasonOk() (*string, bool)`

GetFinalActionReasonOk returns a tuple with the FinalActionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalActionReason

`func (o *PolicyEvaluationResult) SetFinalActionReason(v string)`

SetFinalActionReason sets FinalActionReason field to given value.


### GetImageAllowlisted

`func (o *PolicyEvaluationResult) GetImageAllowlisted() bool`

GetImageAllowlisted returns the ImageAllowlisted field if non-nil, zero value otherwise.

### GetImageAllowlistedOk

`func (o *PolicyEvaluationResult) GetImageAllowlistedOk() (*bool, bool)`

GetImageAllowlistedOk returns a tuple with the ImageAllowlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAllowlisted

`func (o *PolicyEvaluationResult) SetImageAllowlisted(v bool)`

SetImageAllowlisted sets ImageAllowlisted field to given value.


### GetMatchedAllowlistedImagesRule

`func (o *PolicyEvaluationResult) GetMatchedAllowlistedImagesRule() PolicyEvaluationResultMatchedAllowlistedImagesRule`

GetMatchedAllowlistedImagesRule returns the MatchedAllowlistedImagesRule field if non-nil, zero value otherwise.

### GetMatchedAllowlistedImagesRuleOk

`func (o *PolicyEvaluationResult) GetMatchedAllowlistedImagesRuleOk() (*PolicyEvaluationResultMatchedAllowlistedImagesRule, bool)`

GetMatchedAllowlistedImagesRuleOk returns a tuple with the MatchedAllowlistedImagesRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedAllowlistedImagesRule

`func (o *PolicyEvaluationResult) SetMatchedAllowlistedImagesRule(v PolicyEvaluationResultMatchedAllowlistedImagesRule)`

SetMatchedAllowlistedImagesRule sets MatchedAllowlistedImagesRule field to given value.

### HasMatchedAllowlistedImagesRule

`func (o *PolicyEvaluationResult) HasMatchedAllowlistedImagesRule() bool`

HasMatchedAllowlistedImagesRule returns a boolean if a field has been set.

### GetImageDenylisted

`func (o *PolicyEvaluationResult) GetImageDenylisted() bool`

GetImageDenylisted returns the ImageDenylisted field if non-nil, zero value otherwise.

### GetImageDenylistedOk

`func (o *PolicyEvaluationResult) GetImageDenylistedOk() (*bool, bool)`

GetImageDenylistedOk returns a tuple with the ImageDenylisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDenylisted

`func (o *PolicyEvaluationResult) SetImageDenylisted(v bool)`

SetImageDenylisted sets ImageDenylisted field to given value.


### GetMatchedDenylistedImagesRule

`func (o *PolicyEvaluationResult) GetMatchedDenylistedImagesRule() PolicyEvaluationResultMatchedAllowlistedImagesRule`

GetMatchedDenylistedImagesRule returns the MatchedDenylistedImagesRule field if non-nil, zero value otherwise.

### GetMatchedDenylistedImagesRuleOk

`func (o *PolicyEvaluationResult) GetMatchedDenylistedImagesRuleOk() (*PolicyEvaluationResultMatchedAllowlistedImagesRule, bool)`

GetMatchedDenylistedImagesRuleOk returns a tuple with the MatchedDenylistedImagesRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedDenylistedImagesRule

`func (o *PolicyEvaluationResult) SetMatchedDenylistedImagesRule(v PolicyEvaluationResultMatchedAllowlistedImagesRule)`

SetMatchedDenylistedImagesRule sets MatchedDenylistedImagesRule field to given value.

### HasMatchedDenylistedImagesRule

`func (o *PolicyEvaluationResult) HasMatchedDenylistedImagesRule() bool`

HasMatchedDenylistedImagesRule returns a boolean if a field has been set.

### GetImageMappedToRule

`func (o *PolicyEvaluationResult) GetImageMappedToRule() bool`

GetImageMappedToRule returns the ImageMappedToRule field if non-nil, zero value otherwise.

### GetImageMappedToRuleOk

`func (o *PolicyEvaluationResult) GetImageMappedToRuleOk() (*bool, bool)`

GetImageMappedToRuleOk returns a tuple with the ImageMappedToRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMappedToRule

`func (o *PolicyEvaluationResult) SetImageMappedToRule(v bool)`

SetImageMappedToRule sets ImageMappedToRule field to given value.


### GetMatchedMappingRule

`func (o *PolicyEvaluationResult) GetMatchedMappingRule() PolicyEvaluationResultMatchedMappingRule`

GetMatchedMappingRule returns the MatchedMappingRule field if non-nil, zero value otherwise.

### GetMatchedMappingRuleOk

`func (o *PolicyEvaluationResult) GetMatchedMappingRuleOk() (*PolicyEvaluationResultMatchedMappingRule, bool)`

GetMatchedMappingRuleOk returns a tuple with the MatchedMappingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedMappingRule

`func (o *PolicyEvaluationResult) SetMatchedMappingRule(v PolicyEvaluationResultMatchedMappingRule)`

SetMatchedMappingRule sets MatchedMappingRule field to given value.

### HasMatchedMappingRule

`func (o *PolicyEvaluationResult) HasMatchedMappingRule() bool`

HasMatchedMappingRule returns a boolean if a field has been set.

### GetNumberOfFindings

`func (o *PolicyEvaluationResult) GetNumberOfFindings() int32`

GetNumberOfFindings returns the NumberOfFindings field if non-nil, zero value otherwise.

### GetNumberOfFindingsOk

`func (o *PolicyEvaluationResult) GetNumberOfFindingsOk() (*int32, bool)`

GetNumberOfFindingsOk returns a tuple with the NumberOfFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFindings

`func (o *PolicyEvaluationResult) SetNumberOfFindings(v int32)`

SetNumberOfFindings sets NumberOfFindings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


