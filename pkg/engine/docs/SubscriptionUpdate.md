# SubscriptionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionValue** | Pointer to **NullableString** | The new subscription value, e.g. the new tag to be subscribed to | [optional] 
**Active** | Pointer to **bool** | Toggle the subscription processing on or off | [optional] 

## Methods

### NewSubscriptionUpdate

`func NewSubscriptionUpdate() *SubscriptionUpdate`

NewSubscriptionUpdate instantiates a new SubscriptionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionUpdateWithDefaults

`func NewSubscriptionUpdateWithDefaults() *SubscriptionUpdate`

NewSubscriptionUpdateWithDefaults instantiates a new SubscriptionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionValue

`func (o *SubscriptionUpdate) GetSubscriptionValue() string`

GetSubscriptionValue returns the SubscriptionValue field if non-nil, zero value otherwise.

### GetSubscriptionValueOk

`func (o *SubscriptionUpdate) GetSubscriptionValueOk() (*string, bool)`

GetSubscriptionValueOk returns a tuple with the SubscriptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionValue

`func (o *SubscriptionUpdate) SetSubscriptionValue(v string)`

SetSubscriptionValue sets SubscriptionValue field to given value.

### HasSubscriptionValue

`func (o *SubscriptionUpdate) HasSubscriptionValue() bool`

HasSubscriptionValue returns a boolean if a field has been set.

### SetSubscriptionValueNil

`func (o *SubscriptionUpdate) SetSubscriptionValueNil(b bool)`

 SetSubscriptionValueNil sets the value for SubscriptionValue to be an explicit nil

### UnsetSubscriptionValue
`func (o *SubscriptionUpdate) UnsetSubscriptionValue()`

UnsetSubscriptionValue ensures that no value is present for SubscriptionValue, not even an explicit nil
### GetActive

`func (o *SubscriptionUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SubscriptionUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SubscriptionUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SubscriptionUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


