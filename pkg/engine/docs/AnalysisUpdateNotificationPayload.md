# AnalysisUpdateNotificationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**SubscriptionKey** | Pointer to **string** |  | [optional] 
**SubscriptionType** | Pointer to **string** |  | [optional] 
**NotificationId** | Pointer to **string** |  | [optional] 
**CurrEval** | Pointer to [**AnalysisUpdateEval**](AnalysisUpdateEval.md) |  | [optional] 
**LastEval** | Pointer to [**AnalysisUpdateEval**](AnalysisUpdateEval.md) |  | [optional] 
**Annotations** | Pointer to **interface{}** | List of Corresponding Image Annotations | [optional] 

## Methods

### NewAnalysisUpdateNotificationPayload

`func NewAnalysisUpdateNotificationPayload() *AnalysisUpdateNotificationPayload`

NewAnalysisUpdateNotificationPayload instantiates a new AnalysisUpdateNotificationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisUpdateNotificationPayloadWithDefaults

`func NewAnalysisUpdateNotificationPayloadWithDefaults() *AnalysisUpdateNotificationPayload`

NewAnalysisUpdateNotificationPayloadWithDefaults instantiates a new AnalysisUpdateNotificationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AnalysisUpdateNotificationPayload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AnalysisUpdateNotificationPayload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AnalysisUpdateNotificationPayload) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AnalysisUpdateNotificationPayload) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSubscriptionKey

`func (o *AnalysisUpdateNotificationPayload) GetSubscriptionKey() string`

GetSubscriptionKey returns the SubscriptionKey field if non-nil, zero value otherwise.

### GetSubscriptionKeyOk

`func (o *AnalysisUpdateNotificationPayload) GetSubscriptionKeyOk() (*string, bool)`

GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionKey

`func (o *AnalysisUpdateNotificationPayload) SetSubscriptionKey(v string)`

SetSubscriptionKey sets SubscriptionKey field to given value.

### HasSubscriptionKey

`func (o *AnalysisUpdateNotificationPayload) HasSubscriptionKey() bool`

HasSubscriptionKey returns a boolean if a field has been set.

### GetSubscriptionType

`func (o *AnalysisUpdateNotificationPayload) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *AnalysisUpdateNotificationPayload) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *AnalysisUpdateNotificationPayload) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *AnalysisUpdateNotificationPayload) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetNotificationId

`func (o *AnalysisUpdateNotificationPayload) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *AnalysisUpdateNotificationPayload) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *AnalysisUpdateNotificationPayload) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *AnalysisUpdateNotificationPayload) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetCurrEval

`func (o *AnalysisUpdateNotificationPayload) GetCurrEval() AnalysisUpdateEval`

GetCurrEval returns the CurrEval field if non-nil, zero value otherwise.

### GetCurrEvalOk

`func (o *AnalysisUpdateNotificationPayload) GetCurrEvalOk() (*AnalysisUpdateEval, bool)`

GetCurrEvalOk returns a tuple with the CurrEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrEval

`func (o *AnalysisUpdateNotificationPayload) SetCurrEval(v AnalysisUpdateEval)`

SetCurrEval sets CurrEval field to given value.

### HasCurrEval

`func (o *AnalysisUpdateNotificationPayload) HasCurrEval() bool`

HasCurrEval returns a boolean if a field has been set.

### GetLastEval

`func (o *AnalysisUpdateNotificationPayload) GetLastEval() AnalysisUpdateEval`

GetLastEval returns the LastEval field if non-nil, zero value otherwise.

### GetLastEvalOk

`func (o *AnalysisUpdateNotificationPayload) GetLastEvalOk() (*AnalysisUpdateEval, bool)`

GetLastEvalOk returns a tuple with the LastEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEval

`func (o *AnalysisUpdateNotificationPayload) SetLastEval(v AnalysisUpdateEval)`

SetLastEval sets LastEval field to given value.

### HasLastEval

`func (o *AnalysisUpdateNotificationPayload) HasLastEval() bool`

HasLastEval returns a boolean if a field has been set.

### GetAnnotations

`func (o *AnalysisUpdateNotificationPayload) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AnalysisUpdateNotificationPayload) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AnalysisUpdateNotificationPayload) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AnalysisUpdateNotificationPayload) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *AnalysisUpdateNotificationPayload) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *AnalysisUpdateNotificationPayload) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


