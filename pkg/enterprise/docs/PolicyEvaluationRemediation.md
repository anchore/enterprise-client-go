# PolicyEvaluationRemediation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
<<<<<<< HEAD
**Suggestions** | [**[]PolicyEvaluationRemediationSuggestionsInner**](PolicyEvaluationRemediationSuggestionsInner.md) | Anchore generated options for resolving a finding | 
=======
**Suggestions** | [**[]PolicyEvaluationRemediationSuggestions**](PolicyEvaluationRemediationSuggestions.md) | Anchore generated options for resolving a finding | 
>>>>>>> main
**TriggerIds** | **[]string** | List of trigger IDs that these remediation suggestions apply to | 

## Methods

### NewPolicyEvaluationRemediation

<<<<<<< HEAD
`func NewPolicyEvaluationRemediation(suggestions []PolicyEvaluationRemediationSuggestionsInner, triggerIds []string, ) *PolicyEvaluationRemediation`
=======
`func NewPolicyEvaluationRemediation(suggestions []PolicyEvaluationRemediationSuggestions, triggerIds []string, ) *PolicyEvaluationRemediation`
>>>>>>> main

NewPolicyEvaluationRemediation instantiates a new PolicyEvaluationRemediation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationRemediationWithDefaults

`func NewPolicyEvaluationRemediationWithDefaults() *PolicyEvaluationRemediation`

NewPolicyEvaluationRemediationWithDefaults instantiates a new PolicyEvaluationRemediation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestions

<<<<<<< HEAD
`func (o *PolicyEvaluationRemediation) GetSuggestions() []PolicyEvaluationRemediationSuggestionsInner`
=======
`func (o *PolicyEvaluationRemediation) GetSuggestions() []PolicyEvaluationRemediationSuggestions`
>>>>>>> main

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

<<<<<<< HEAD
`func (o *PolicyEvaluationRemediation) GetSuggestionsOk() (*[]PolicyEvaluationRemediationSuggestionsInner, bool)`
=======
`func (o *PolicyEvaluationRemediation) GetSuggestionsOk() (*[]PolicyEvaluationRemediationSuggestions, bool)`
>>>>>>> main

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

<<<<<<< HEAD
`func (o *PolicyEvaluationRemediation) SetSuggestions(v []PolicyEvaluationRemediationSuggestionsInner)`
=======
`func (o *PolicyEvaluationRemediation) SetSuggestions(v []PolicyEvaluationRemediationSuggestions)`
>>>>>>> main

SetSuggestions sets Suggestions field to given value.


### GetTriggerIds

`func (o *PolicyEvaluationRemediation) GetTriggerIds() []string`

GetTriggerIds returns the TriggerIds field if non-nil, zero value otherwise.

### GetTriggerIdsOk

`func (o *PolicyEvaluationRemediation) GetTriggerIdsOk() (*[]string, bool)`

GetTriggerIdsOk returns a tuple with the TriggerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerIds

`func (o *PolicyEvaluationRemediation) SetTriggerIds(v []string)`

SetTriggerIds sets TriggerIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


