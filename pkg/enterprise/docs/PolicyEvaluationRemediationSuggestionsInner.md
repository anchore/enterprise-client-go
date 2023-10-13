# PolicyEvaluationRemediationSuggestionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The suggestion for resolving a finding | 
**Preferred** | **bool** | Indicates whether this suggestion is recommended | 

## Methods

### NewPolicyEvaluationRemediationSuggestionsInner

`func NewPolicyEvaluationRemediationSuggestionsInner(message string, preferred bool, ) *PolicyEvaluationRemediationSuggestionsInner`

NewPolicyEvaluationRemediationSuggestionsInner instantiates a new PolicyEvaluationRemediationSuggestionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationRemediationSuggestionsInnerWithDefaults

`func NewPolicyEvaluationRemediationSuggestionsInnerWithDefaults() *PolicyEvaluationRemediationSuggestionsInner`

NewPolicyEvaluationRemediationSuggestionsInnerWithDefaults instantiates a new PolicyEvaluationRemediationSuggestionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PolicyEvaluationRemediationSuggestionsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PolicyEvaluationRemediationSuggestionsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PolicyEvaluationRemediationSuggestionsInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPreferred

`func (o *PolicyEvaluationRemediationSuggestionsInner) GetPreferred() bool`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *PolicyEvaluationRemediationSuggestionsInner) GetPreferredOk() (*bool, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *PolicyEvaluationRemediationSuggestionsInner) SetPreferred(v bool)`

SetPreferred sets Preferred field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


