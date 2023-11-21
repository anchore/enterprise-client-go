# PolicyEvaluationEvaluationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**NullablePolicyEvaluationEvaluationsInnerDetails**](PolicyEvaluationEvaluationsInnerDetails.md) |  | [optional] 
**ComparisonImageDigest** | Pointer to **NullableString** | Image digest of the base image used during policy evaluation | [optional] 
**EvaluationTime** | **time.Time** | The date and time this policy evaluation was performed at | 
**EvaluationProblems** | [**[]PolicyEvaluationProblem**](PolicyEvaluationProblem.md) | list of error objects indicating errors encountered during evaluation execution | 
**Status** | **string** | The overall status of the policy evaluation | 
**FinalAction** | **string** | The overall outcome of the evaluation. | 
**FinalActionReason** | **string** | The reason for the final result | 
**ImageAllowlisted** | **bool** | Whether the evaluated image matched an allowlist rule | 
**MatchedAllowlistedImagesRule** | Pointer to [**PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule**](PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule.md) |  | [optional] 
**ImageDenylisted** | **bool** | Whether the evaluated image matched a denylist rule | 
**MatchedDenylistedImagesRule** | Pointer to [**PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule**](PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule.md) |  | [optional] 
**ImageMappedToRule** | **bool** | Whether the evaluated image matched a policy rule | 
**MatchedMappingRule** | Pointer to [**PolicyEvaluationEvaluationsInnerMatchedMappingRule**](PolicyEvaluationEvaluationsInnerMatchedMappingRule.md) |  | [optional] 
**NumberOfFindings** | **int32** | Number of policy findings in the response | 

## Methods

### NewPolicyEvaluationEvaluationsInner

`func NewPolicyEvaluationEvaluationsInner(evaluationTime time.Time, evaluationProblems []PolicyEvaluationProblem, status string, finalAction string, finalActionReason string, imageAllowlisted bool, imageDenylisted bool, imageMappedToRule bool, numberOfFindings int32, ) *PolicyEvaluationEvaluationsInner`

NewPolicyEvaluationEvaluationsInner instantiates a new PolicyEvaluationEvaluationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationEvaluationsInnerWithDefaults

`func NewPolicyEvaluationEvaluationsInnerWithDefaults() *PolicyEvaluationEvaluationsInner`

NewPolicyEvaluationEvaluationsInnerWithDefaults instantiates a new PolicyEvaluationEvaluationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *PolicyEvaluationEvaluationsInner) GetDetails() PolicyEvaluationEvaluationsInnerDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PolicyEvaluationEvaluationsInner) GetDetailsOk() (*PolicyEvaluationEvaluationsInnerDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PolicyEvaluationEvaluationsInner) SetDetails(v PolicyEvaluationEvaluationsInnerDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PolicyEvaluationEvaluationsInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *PolicyEvaluationEvaluationsInner) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *PolicyEvaluationEvaluationsInner) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetComparisonImageDigest

`func (o *PolicyEvaluationEvaluationsInner) GetComparisonImageDigest() string`

GetComparisonImageDigest returns the ComparisonImageDigest field if non-nil, zero value otherwise.

### GetComparisonImageDigestOk

`func (o *PolicyEvaluationEvaluationsInner) GetComparisonImageDigestOk() (*string, bool)`

GetComparisonImageDigestOk returns a tuple with the ComparisonImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonImageDigest

`func (o *PolicyEvaluationEvaluationsInner) SetComparisonImageDigest(v string)`

SetComparisonImageDigest sets ComparisonImageDigest field to given value.

### HasComparisonImageDigest

`func (o *PolicyEvaluationEvaluationsInner) HasComparisonImageDigest() bool`

HasComparisonImageDigest returns a boolean if a field has been set.

### SetComparisonImageDigestNil

`func (o *PolicyEvaluationEvaluationsInner) SetComparisonImageDigestNil(b bool)`

 SetComparisonImageDigestNil sets the value for ComparisonImageDigest to be an explicit nil

### UnsetComparisonImageDigest
`func (o *PolicyEvaluationEvaluationsInner) UnsetComparisonImageDigest()`

UnsetComparisonImageDigest ensures that no value is present for ComparisonImageDigest, not even an explicit nil
### GetEvaluationTime

`func (o *PolicyEvaluationEvaluationsInner) GetEvaluationTime() time.Time`

GetEvaluationTime returns the EvaluationTime field if non-nil, zero value otherwise.

### GetEvaluationTimeOk

`func (o *PolicyEvaluationEvaluationsInner) GetEvaluationTimeOk() (*time.Time, bool)`

GetEvaluationTimeOk returns a tuple with the EvaluationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTime

`func (o *PolicyEvaluationEvaluationsInner) SetEvaluationTime(v time.Time)`

SetEvaluationTime sets EvaluationTime field to given value.


### GetEvaluationProblems

`func (o *PolicyEvaluationEvaluationsInner) GetEvaluationProblems() []PolicyEvaluationProblem`

GetEvaluationProblems returns the EvaluationProblems field if non-nil, zero value otherwise.

### GetEvaluationProblemsOk

`func (o *PolicyEvaluationEvaluationsInner) GetEvaluationProblemsOk() (*[]PolicyEvaluationProblem, bool)`

GetEvaluationProblemsOk returns a tuple with the EvaluationProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationProblems

`func (o *PolicyEvaluationEvaluationsInner) SetEvaluationProblems(v []PolicyEvaluationProblem)`

SetEvaluationProblems sets EvaluationProblems field to given value.


### GetStatus

`func (o *PolicyEvaluationEvaluationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PolicyEvaluationEvaluationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PolicyEvaluationEvaluationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFinalAction

`func (o *PolicyEvaluationEvaluationsInner) GetFinalAction() string`

GetFinalAction returns the FinalAction field if non-nil, zero value otherwise.

### GetFinalActionOk

`func (o *PolicyEvaluationEvaluationsInner) GetFinalActionOk() (*string, bool)`

GetFinalActionOk returns a tuple with the FinalAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalAction

`func (o *PolicyEvaluationEvaluationsInner) SetFinalAction(v string)`

SetFinalAction sets FinalAction field to given value.


### GetFinalActionReason

`func (o *PolicyEvaluationEvaluationsInner) GetFinalActionReason() string`

GetFinalActionReason returns the FinalActionReason field if non-nil, zero value otherwise.

### GetFinalActionReasonOk

`func (o *PolicyEvaluationEvaluationsInner) GetFinalActionReasonOk() (*string, bool)`

GetFinalActionReasonOk returns a tuple with the FinalActionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalActionReason

`func (o *PolicyEvaluationEvaluationsInner) SetFinalActionReason(v string)`

SetFinalActionReason sets FinalActionReason field to given value.


### GetImageAllowlisted

`func (o *PolicyEvaluationEvaluationsInner) GetImageAllowlisted() bool`

GetImageAllowlisted returns the ImageAllowlisted field if non-nil, zero value otherwise.

### GetImageAllowlistedOk

`func (o *PolicyEvaluationEvaluationsInner) GetImageAllowlistedOk() (*bool, bool)`

GetImageAllowlistedOk returns a tuple with the ImageAllowlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAllowlisted

`func (o *PolicyEvaluationEvaluationsInner) SetImageAllowlisted(v bool)`

SetImageAllowlisted sets ImageAllowlisted field to given value.


### GetMatchedAllowlistedImagesRule

`func (o *PolicyEvaluationEvaluationsInner) GetMatchedAllowlistedImagesRule() PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule`

GetMatchedAllowlistedImagesRule returns the MatchedAllowlistedImagesRule field if non-nil, zero value otherwise.

### GetMatchedAllowlistedImagesRuleOk

`func (o *PolicyEvaluationEvaluationsInner) GetMatchedAllowlistedImagesRuleOk() (*PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule, bool)`

GetMatchedAllowlistedImagesRuleOk returns a tuple with the MatchedAllowlistedImagesRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedAllowlistedImagesRule

`func (o *PolicyEvaluationEvaluationsInner) SetMatchedAllowlistedImagesRule(v PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule)`

SetMatchedAllowlistedImagesRule sets MatchedAllowlistedImagesRule field to given value.

### HasMatchedAllowlistedImagesRule

`func (o *PolicyEvaluationEvaluationsInner) HasMatchedAllowlistedImagesRule() bool`

HasMatchedAllowlistedImagesRule returns a boolean if a field has been set.

### GetImageDenylisted

`func (o *PolicyEvaluationEvaluationsInner) GetImageDenylisted() bool`

GetImageDenylisted returns the ImageDenylisted field if non-nil, zero value otherwise.

### GetImageDenylistedOk

`func (o *PolicyEvaluationEvaluationsInner) GetImageDenylistedOk() (*bool, bool)`

GetImageDenylistedOk returns a tuple with the ImageDenylisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDenylisted

`func (o *PolicyEvaluationEvaluationsInner) SetImageDenylisted(v bool)`

SetImageDenylisted sets ImageDenylisted field to given value.


### GetMatchedDenylistedImagesRule

`func (o *PolicyEvaluationEvaluationsInner) GetMatchedDenylistedImagesRule() PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule`

GetMatchedDenylistedImagesRule returns the MatchedDenylistedImagesRule field if non-nil, zero value otherwise.

### GetMatchedDenylistedImagesRuleOk

`func (o *PolicyEvaluationEvaluationsInner) GetMatchedDenylistedImagesRuleOk() (*PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule, bool)`

GetMatchedDenylistedImagesRuleOk returns a tuple with the MatchedDenylistedImagesRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedDenylistedImagesRule

`func (o *PolicyEvaluationEvaluationsInner) SetMatchedDenylistedImagesRule(v PolicyEvaluationEvaluationsInnerMatchedAllowlistedImagesRule)`

SetMatchedDenylistedImagesRule sets MatchedDenylistedImagesRule field to given value.

### HasMatchedDenylistedImagesRule

`func (o *PolicyEvaluationEvaluationsInner) HasMatchedDenylistedImagesRule() bool`

HasMatchedDenylistedImagesRule returns a boolean if a field has been set.

### GetImageMappedToRule

`func (o *PolicyEvaluationEvaluationsInner) GetImageMappedToRule() bool`

GetImageMappedToRule returns the ImageMappedToRule field if non-nil, zero value otherwise.

### GetImageMappedToRuleOk

`func (o *PolicyEvaluationEvaluationsInner) GetImageMappedToRuleOk() (*bool, bool)`

GetImageMappedToRuleOk returns a tuple with the ImageMappedToRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMappedToRule

`func (o *PolicyEvaluationEvaluationsInner) SetImageMappedToRule(v bool)`

SetImageMappedToRule sets ImageMappedToRule field to given value.


### GetMatchedMappingRule

`func (o *PolicyEvaluationEvaluationsInner) GetMatchedMappingRule() PolicyEvaluationEvaluationsInnerMatchedMappingRule`

GetMatchedMappingRule returns the MatchedMappingRule field if non-nil, zero value otherwise.

### GetMatchedMappingRuleOk

`func (o *PolicyEvaluationEvaluationsInner) GetMatchedMappingRuleOk() (*PolicyEvaluationEvaluationsInnerMatchedMappingRule, bool)`

GetMatchedMappingRuleOk returns a tuple with the MatchedMappingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedMappingRule

`func (o *PolicyEvaluationEvaluationsInner) SetMatchedMappingRule(v PolicyEvaluationEvaluationsInnerMatchedMappingRule)`

SetMatchedMappingRule sets MatchedMappingRule field to given value.

### HasMatchedMappingRule

`func (o *PolicyEvaluationEvaluationsInner) HasMatchedMappingRule() bool`

HasMatchedMappingRule returns a boolean if a field has been set.

### GetNumberOfFindings

`func (o *PolicyEvaluationEvaluationsInner) GetNumberOfFindings() int32`

GetNumberOfFindings returns the NumberOfFindings field if non-nil, zero value otherwise.

### GetNumberOfFindingsOk

`func (o *PolicyEvaluationEvaluationsInner) GetNumberOfFindingsOk() (*int32, bool)`

GetNumberOfFindingsOk returns a tuple with the NumberOfFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFindings

`func (o *PolicyEvaluationEvaluationsInner) SetNumberOfFindings(v int32)`

SetNumberOfFindings sets NumberOfFindings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


