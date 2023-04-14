# TagUpdateNotificationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** |  | [optional] 
**SubscriptionKey** | Pointer to **string** |  | [optional] 
**SubscriptionType** | Pointer to **string** |  | [optional] 
**NotificationId** | Pointer to **string** |  | [optional] 
**CurrEval** | Pointer to **[]interface{}** | A list containing the current image digest | [optional] 
**LastEval** | Pointer to **[]interface{}** | A list containing the previous image digests | [optional] 
**Annotations** | Pointer to **interface{}** | List of Corresponding Image Annotations | [optional] 

## Methods

### NewTagUpdateNotificationPayload

`func NewTagUpdateNotificationPayload() *TagUpdateNotificationPayload`

NewTagUpdateNotificationPayload instantiates a new TagUpdateNotificationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagUpdateNotificationPayloadWithDefaults

`func NewTagUpdateNotificationPayloadWithDefaults() *TagUpdateNotificationPayload`

NewTagUpdateNotificationPayloadWithDefaults instantiates a new TagUpdateNotificationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *TagUpdateNotificationPayload) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *TagUpdateNotificationPayload) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *TagUpdateNotificationPayload) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *TagUpdateNotificationPayload) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetSubscriptionKey

`func (o *TagUpdateNotificationPayload) GetSubscriptionKey() string`

GetSubscriptionKey returns the SubscriptionKey field if non-nil, zero value otherwise.

### GetSubscriptionKeyOk

`func (o *TagUpdateNotificationPayload) GetSubscriptionKeyOk() (*string, bool)`

GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionKey

`func (o *TagUpdateNotificationPayload) SetSubscriptionKey(v string)`

SetSubscriptionKey sets SubscriptionKey field to given value.

### HasSubscriptionKey

`func (o *TagUpdateNotificationPayload) HasSubscriptionKey() bool`

HasSubscriptionKey returns a boolean if a field has been set.

### GetSubscriptionType

`func (o *TagUpdateNotificationPayload) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *TagUpdateNotificationPayload) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *TagUpdateNotificationPayload) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *TagUpdateNotificationPayload) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetNotificationId

`func (o *TagUpdateNotificationPayload) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *TagUpdateNotificationPayload) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *TagUpdateNotificationPayload) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *TagUpdateNotificationPayload) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetCurrEval

`func (o *TagUpdateNotificationPayload) GetCurrEval() []interface{}`

GetCurrEval returns the CurrEval field if non-nil, zero value otherwise.

### GetCurrEvalOk

`func (o *TagUpdateNotificationPayload) GetCurrEvalOk() (*[]interface{}, bool)`

GetCurrEvalOk returns a tuple with the CurrEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrEval

`func (o *TagUpdateNotificationPayload) SetCurrEval(v []interface{})`

SetCurrEval sets CurrEval field to given value.

### HasCurrEval

`func (o *TagUpdateNotificationPayload) HasCurrEval() bool`

HasCurrEval returns a boolean if a field has been set.

### GetLastEval

`func (o *TagUpdateNotificationPayload) GetLastEval() []interface{}`

GetLastEval returns the LastEval field if non-nil, zero value otherwise.

### GetLastEvalOk

`func (o *TagUpdateNotificationPayload) GetLastEvalOk() (*[]interface{}, bool)`

GetLastEvalOk returns a tuple with the LastEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEval

`func (o *TagUpdateNotificationPayload) SetLastEval(v []interface{})`

SetLastEval sets LastEval field to given value.

### HasLastEval

`func (o *TagUpdateNotificationPayload) HasLastEval() bool`

HasLastEval returns a boolean if a field has been set.

### GetAnnotations

`func (o *TagUpdateNotificationPayload) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *TagUpdateNotificationPayload) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *TagUpdateNotificationPayload) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *TagUpdateNotificationPayload) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *TagUpdateNotificationPayload) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *TagUpdateNotificationPayload) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


