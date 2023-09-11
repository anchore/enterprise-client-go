# PolicyEvalNotificationPayloadAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrEval** | Pointer to **interface{}** | The Current Policy Evaluation result | [optional] 
**LastEval** | Pointer to **interface{}** | The Previous Policy Evaluation result | [optional] 
**Annotations** | Pointer to **interface{}** | List of Corresponding Image Annotations | [optional] 

## Methods

### NewPolicyEvalNotificationPayloadAllOf

`func NewPolicyEvalNotificationPayloadAllOf() *PolicyEvalNotificationPayloadAllOf`

NewPolicyEvalNotificationPayloadAllOf instantiates a new PolicyEvalNotificationPayloadAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvalNotificationPayloadAllOfWithDefaults

`func NewPolicyEvalNotificationPayloadAllOfWithDefaults() *PolicyEvalNotificationPayloadAllOf`

NewPolicyEvalNotificationPayloadAllOfWithDefaults instantiates a new PolicyEvalNotificationPayloadAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrEval

`func (o *PolicyEvalNotificationPayloadAllOf) GetCurrEval() interface{}`

GetCurrEval returns the CurrEval field if non-nil, zero value otherwise.

### GetCurrEvalOk

`func (o *PolicyEvalNotificationPayloadAllOf) GetCurrEvalOk() (*interface{}, bool)`

GetCurrEvalOk returns a tuple with the CurrEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrEval

`func (o *PolicyEvalNotificationPayloadAllOf) SetCurrEval(v interface{})`

SetCurrEval sets CurrEval field to given value.

### HasCurrEval

`func (o *PolicyEvalNotificationPayloadAllOf) HasCurrEval() bool`

HasCurrEval returns a boolean if a field has been set.

### GetLastEval

`func (o *PolicyEvalNotificationPayloadAllOf) GetLastEval() interface{}`

GetLastEval returns the LastEval field if non-nil, zero value otherwise.

### GetLastEvalOk

`func (o *PolicyEvalNotificationPayloadAllOf) GetLastEvalOk() (*interface{}, bool)`

GetLastEvalOk returns a tuple with the LastEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEval

`func (o *PolicyEvalNotificationPayloadAllOf) SetLastEval(v interface{})`

SetLastEval sets LastEval field to given value.

### HasLastEval

`func (o *PolicyEvalNotificationPayloadAllOf) HasLastEval() bool`

HasLastEval returns a boolean if a field has been set.

### GetAnnotations

`func (o *PolicyEvalNotificationPayloadAllOf) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *PolicyEvalNotificationPayloadAllOf) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *PolicyEvalNotificationPayloadAllOf) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *PolicyEvalNotificationPayloadAllOf) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *PolicyEvalNotificationPayloadAllOf) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *PolicyEvalNotificationPayloadAllOf) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


