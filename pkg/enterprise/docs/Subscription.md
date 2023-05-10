# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionKey** | Pointer to **string** | The key value that the subscription references. E.g. a tag value or a repo name. | [optional] 
**SubscriptionType** | Pointer to **string** | The type of the subscription | [optional] 
**SubscriptionValue** | Pointer to **NullableString** | The value of the subscription target | [optional] 
**AccountName** | Pointer to **string** | The account_name of the subscribed user | [optional] 
**Active** | Pointer to **bool** | Is the subscription currently active | [optional] 
**SubscriptionId** | Pointer to **string** | the unique id for this subscription record | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionKey

`func (o *Subscription) GetSubscriptionKey() string`

GetSubscriptionKey returns the SubscriptionKey field if non-nil, zero value otherwise.

### GetSubscriptionKeyOk

`func (o *Subscription) GetSubscriptionKeyOk() (*string, bool)`

GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionKey

`func (o *Subscription) SetSubscriptionKey(v string)`

SetSubscriptionKey sets SubscriptionKey field to given value.

### HasSubscriptionKey

`func (o *Subscription) HasSubscriptionKey() bool`

HasSubscriptionKey returns a boolean if a field has been set.

### GetSubscriptionType

`func (o *Subscription) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *Subscription) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *Subscription) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *Subscription) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetSubscriptionValue

`func (o *Subscription) GetSubscriptionValue() string`

GetSubscriptionValue returns the SubscriptionValue field if non-nil, zero value otherwise.

### GetSubscriptionValueOk

`func (o *Subscription) GetSubscriptionValueOk() (*string, bool)`

GetSubscriptionValueOk returns a tuple with the SubscriptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionValue

`func (o *Subscription) SetSubscriptionValue(v string)`

SetSubscriptionValue sets SubscriptionValue field to given value.

### HasSubscriptionValue

`func (o *Subscription) HasSubscriptionValue() bool`

HasSubscriptionValue returns a boolean if a field has been set.

### SetSubscriptionValueNil

`func (o *Subscription) SetSubscriptionValueNil(b bool)`

 SetSubscriptionValueNil sets the value for SubscriptionValue to be an explicit nil

### UnsetSubscriptionValue
`func (o *Subscription) UnsetSubscriptionValue()`

UnsetSubscriptionValue ensures that no value is present for SubscriptionValue, not even an explicit nil
### GetAccountName

`func (o *Subscription) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Subscription) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Subscription) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Subscription) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetActive

`func (o *Subscription) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Subscription) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Subscription) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Subscription) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *Subscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Subscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Subscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *Subscription) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


