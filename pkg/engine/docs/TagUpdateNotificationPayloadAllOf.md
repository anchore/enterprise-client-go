# TagUpdateNotificationPayloadAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrEval** | Pointer to **[]interface{}** | A list containing the current image digest | [optional] 
**LastEval** | Pointer to **[]interface{}** | A list containing the previous image digests | [optional] 
**Annotations** | Pointer to **interface{}** | List of Corresponding Image Annotations | [optional] 

## Methods

### NewTagUpdateNotificationPayloadAllOf

`func NewTagUpdateNotificationPayloadAllOf() *TagUpdateNotificationPayloadAllOf`

NewTagUpdateNotificationPayloadAllOf instantiates a new TagUpdateNotificationPayloadAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagUpdateNotificationPayloadAllOfWithDefaults

`func NewTagUpdateNotificationPayloadAllOfWithDefaults() *TagUpdateNotificationPayloadAllOf`

NewTagUpdateNotificationPayloadAllOfWithDefaults instantiates a new TagUpdateNotificationPayloadAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrEval

`func (o *TagUpdateNotificationPayloadAllOf) GetCurrEval() []interface{}`

GetCurrEval returns the CurrEval field if non-nil, zero value otherwise.

### GetCurrEvalOk

`func (o *TagUpdateNotificationPayloadAllOf) GetCurrEvalOk() (*[]interface{}, bool)`

GetCurrEvalOk returns a tuple with the CurrEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrEval

`func (o *TagUpdateNotificationPayloadAllOf) SetCurrEval(v []interface{})`

SetCurrEval sets CurrEval field to given value.

### HasCurrEval

`func (o *TagUpdateNotificationPayloadAllOf) HasCurrEval() bool`

HasCurrEval returns a boolean if a field has been set.

### GetLastEval

`func (o *TagUpdateNotificationPayloadAllOf) GetLastEval() []interface{}`

GetLastEval returns the LastEval field if non-nil, zero value otherwise.

### GetLastEvalOk

`func (o *TagUpdateNotificationPayloadAllOf) GetLastEvalOk() (*[]interface{}, bool)`

GetLastEvalOk returns a tuple with the LastEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEval

`func (o *TagUpdateNotificationPayloadAllOf) SetLastEval(v []interface{})`

SetLastEval sets LastEval field to given value.

### HasLastEval

`func (o *TagUpdateNotificationPayloadAllOf) HasLastEval() bool`

HasLastEval returns a boolean if a field has been set.

### GetAnnotations

`func (o *TagUpdateNotificationPayloadAllOf) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *TagUpdateNotificationPayloadAllOf) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *TagUpdateNotificationPayloadAllOf) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *TagUpdateNotificationPayloadAllOf) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *TagUpdateNotificationPayloadAllOf) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *TagUpdateNotificationPayloadAllOf) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


