# PolicyEvalNotificationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** |  | [optional] 
**SubscriptionKey** | Pointer to **string** |  | [optional] 
**SubscriptionType** | Pointer to **string** |  | [optional] 
**NotificationId** | Pointer to **string** |  | [optional] 
**CurrEval** | Pointer to **interface{}** | The Current Policy Evaluation result | [optional] 
**LastEval** | Pointer to **interface{}** | The Previous Policy Evaluation result | [optional] 
**Annotations** | Pointer to **interface{}** | List of Corresponding Image Annotations | [optional] 

## Methods

### NewPolicyEvalNotificationPayload

`func NewPolicyEvalNotificationPayload() *PolicyEvalNotificationPayload`

NewPolicyEvalNotificationPayload instantiates a new PolicyEvalNotificationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvalNotificationPayloadWithDefaults

`func NewPolicyEvalNotificationPayloadWithDefaults() *PolicyEvalNotificationPayload`

NewPolicyEvalNotificationPayloadWithDefaults instantiates a new PolicyEvalNotificationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *PolicyEvalNotificationPayload) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PolicyEvalNotificationPayload) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PolicyEvalNotificationPayload) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PolicyEvalNotificationPayload) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetSubscriptionKey

`func (o *PolicyEvalNotificationPayload) GetSubscriptionKey() string`

GetSubscriptionKey returns the SubscriptionKey field if non-nil, zero value otherwise.

### GetSubscriptionKeyOk

`func (o *PolicyEvalNotificationPayload) GetSubscriptionKeyOk() (*string, bool)`

GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionKey

`func (o *PolicyEvalNotificationPayload) SetSubscriptionKey(v string)`

SetSubscriptionKey sets SubscriptionKey field to given value.

### HasSubscriptionKey

`func (o *PolicyEvalNotificationPayload) HasSubscriptionKey() bool`

HasSubscriptionKey returns a boolean if a field has been set.

### GetSubscriptionType

`func (o *PolicyEvalNotificationPayload) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *PolicyEvalNotificationPayload) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *PolicyEvalNotificationPayload) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *PolicyEvalNotificationPayload) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetNotificationId

`func (o *PolicyEvalNotificationPayload) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *PolicyEvalNotificationPayload) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *PolicyEvalNotificationPayload) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *PolicyEvalNotificationPayload) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetCurrEval

`func (o *PolicyEvalNotificationPayload) GetCurrEval() interface{}`

GetCurrEval returns the CurrEval field if non-nil, zero value otherwise.

### GetCurrEvalOk

`func (o *PolicyEvalNotificationPayload) GetCurrEvalOk() (*interface{}, bool)`

GetCurrEvalOk returns a tuple with the CurrEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrEval

`func (o *PolicyEvalNotificationPayload) SetCurrEval(v interface{})`

SetCurrEval sets CurrEval field to given value.

### HasCurrEval

`func (o *PolicyEvalNotificationPayload) HasCurrEval() bool`

HasCurrEval returns a boolean if a field has been set.

### GetLastEval

`func (o *PolicyEvalNotificationPayload) GetLastEval() interface{}`

GetLastEval returns the LastEval field if non-nil, zero value otherwise.

### GetLastEvalOk

`func (o *PolicyEvalNotificationPayload) GetLastEvalOk() (*interface{}, bool)`

GetLastEvalOk returns a tuple with the LastEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEval

`func (o *PolicyEvalNotificationPayload) SetLastEval(v interface{})`

SetLastEval sets LastEval field to given value.

### HasLastEval

`func (o *PolicyEvalNotificationPayload) HasLastEval() bool`

HasLastEval returns a boolean if a field has been set.

### GetAnnotations

`func (o *PolicyEvalNotificationPayload) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *PolicyEvalNotificationPayload) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *PolicyEvalNotificationPayload) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *PolicyEvalNotificationPayload) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *PolicyEvalNotificationPayload) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *PolicyEvalNotificationPayload) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


