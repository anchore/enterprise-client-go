# AnalysisUpdateNotificationPayloadAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrEval** | Pointer to [**AnalysisUpdateEval**](AnalysisUpdateEval.md) |  | [optional] 
**LastEval** | Pointer to [**AnalysisUpdateEval**](AnalysisUpdateEval.md) |  | [optional] 
**Annotations** | Pointer to **interface{}** | List of Corresponding Image Annotations | [optional] 

## Methods

### NewAnalysisUpdateNotificationPayloadAllOf

`func NewAnalysisUpdateNotificationPayloadAllOf() *AnalysisUpdateNotificationPayloadAllOf`

NewAnalysisUpdateNotificationPayloadAllOf instantiates a new AnalysisUpdateNotificationPayloadAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisUpdateNotificationPayloadAllOfWithDefaults

`func NewAnalysisUpdateNotificationPayloadAllOfWithDefaults() *AnalysisUpdateNotificationPayloadAllOf`

NewAnalysisUpdateNotificationPayloadAllOfWithDefaults instantiates a new AnalysisUpdateNotificationPayloadAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrEval

`func (o *AnalysisUpdateNotificationPayloadAllOf) GetCurrEval() AnalysisUpdateEval`

GetCurrEval returns the CurrEval field if non-nil, zero value otherwise.

### GetCurrEvalOk

`func (o *AnalysisUpdateNotificationPayloadAllOf) GetCurrEvalOk() (*AnalysisUpdateEval, bool)`

GetCurrEvalOk returns a tuple with the CurrEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrEval

`func (o *AnalysisUpdateNotificationPayloadAllOf) SetCurrEval(v AnalysisUpdateEval)`

SetCurrEval sets CurrEval field to given value.

### HasCurrEval

`func (o *AnalysisUpdateNotificationPayloadAllOf) HasCurrEval() bool`

HasCurrEval returns a boolean if a field has been set.

### GetLastEval

`func (o *AnalysisUpdateNotificationPayloadAllOf) GetLastEval() AnalysisUpdateEval`

GetLastEval returns the LastEval field if non-nil, zero value otherwise.

### GetLastEvalOk

`func (o *AnalysisUpdateNotificationPayloadAllOf) GetLastEvalOk() (*AnalysisUpdateEval, bool)`

GetLastEvalOk returns a tuple with the LastEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEval

`func (o *AnalysisUpdateNotificationPayloadAllOf) SetLastEval(v AnalysisUpdateEval)`

SetLastEval sets LastEval field to given value.

### HasLastEval

`func (o *AnalysisUpdateNotificationPayloadAllOf) HasLastEval() bool`

HasLastEval returns a boolean if a field has been set.

### GetAnnotations

`func (o *AnalysisUpdateNotificationPayloadAllOf) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AnalysisUpdateNotificationPayloadAllOf) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AnalysisUpdateNotificationPayloadAllOf) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AnalysisUpdateNotificationPayloadAllOf) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *AnalysisUpdateNotificationPayloadAllOf) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *AnalysisUpdateNotificationPayloadAllOf) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


