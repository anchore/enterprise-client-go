# PolicyEvaluationRemediationSuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The suggestion for resolving a finding | 
**Preferred** | **bool** | Indicates whether this suggestion is recommended | 

## Methods

### NewPolicyEvaluationRemediationSuggestion

`func NewPolicyEvaluationRemediationSuggestion(message string, preferred bool, ) *PolicyEvaluationRemediationSuggestion`

NewPolicyEvaluationRemediationSuggestion instantiates a new PolicyEvaluationRemediationSuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationRemediationSuggestionWithDefaults

`func NewPolicyEvaluationRemediationSuggestionWithDefaults() *PolicyEvaluationRemediationSuggestion`

NewPolicyEvaluationRemediationSuggestionWithDefaults instantiates a new PolicyEvaluationRemediationSuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PolicyEvaluationRemediationSuggestion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PolicyEvaluationRemediationSuggestion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PolicyEvaluationRemediationSuggestion) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPreferred

`func (o *PolicyEvaluationRemediationSuggestion) GetPreferred() bool`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *PolicyEvaluationRemediationSuggestion) GetPreferredOk() (*bool, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *PolicyEvaluationRemediationSuggestion) SetPreferred(v bool)`

SetPreferred sets Preferred field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


