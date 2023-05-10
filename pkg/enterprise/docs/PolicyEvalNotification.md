# PolicyEvalNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueId** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**DataId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**LastUpdated** | Pointer to **int32** |  | [optional] 
**RecordStateKey** | Pointer to **string** |  | [optional] [default to "active"]
**RecordStateVal** | Pointer to **NullableString** |  | [optional] 
**Tries** | Pointer to **int32** |  | [optional] 
**MaxTries** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**PolicyEvalNotificationData**](PolicyEvalNotificationData.md) |  | [optional] 

## Methods

### NewPolicyEvalNotification

`func NewPolicyEvalNotification() *PolicyEvalNotification`

NewPolicyEvalNotification instantiates a new PolicyEvalNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvalNotificationWithDefaults

`func NewPolicyEvalNotificationWithDefaults() *PolicyEvalNotification`

NewPolicyEvalNotificationWithDefaults instantiates a new PolicyEvalNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueId

`func (o *PolicyEvalNotification) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *PolicyEvalNotification) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *PolicyEvalNotification) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *PolicyEvalNotification) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetAccountName

`func (o *PolicyEvalNotification) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PolicyEvalNotification) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PolicyEvalNotification) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PolicyEvalNotification) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetDataId

`func (o *PolicyEvalNotification) GetDataId() string`

GetDataId returns the DataId field if non-nil, zero value otherwise.

### GetDataIdOk

`func (o *PolicyEvalNotification) GetDataIdOk() (*string, bool)`

GetDataIdOk returns a tuple with the DataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataId

`func (o *PolicyEvalNotification) SetDataId(v string)`

SetDataId sets DataId field to given value.

### HasDataId

`func (o *PolicyEvalNotification) HasDataId() bool`

HasDataId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PolicyEvalNotification) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyEvalNotification) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyEvalNotification) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PolicyEvalNotification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PolicyEvalNotification) GetLastUpdated() int32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PolicyEvalNotification) GetLastUpdatedOk() (*int32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PolicyEvalNotification) SetLastUpdated(v int32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PolicyEvalNotification) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRecordStateKey

`func (o *PolicyEvalNotification) GetRecordStateKey() string`

GetRecordStateKey returns the RecordStateKey field if non-nil, zero value otherwise.

### GetRecordStateKeyOk

`func (o *PolicyEvalNotification) GetRecordStateKeyOk() (*string, bool)`

GetRecordStateKeyOk returns a tuple with the RecordStateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStateKey

`func (o *PolicyEvalNotification) SetRecordStateKey(v string)`

SetRecordStateKey sets RecordStateKey field to given value.

### HasRecordStateKey

`func (o *PolicyEvalNotification) HasRecordStateKey() bool`

HasRecordStateKey returns a boolean if a field has been set.

### GetRecordStateVal

`func (o *PolicyEvalNotification) GetRecordStateVal() string`

GetRecordStateVal returns the RecordStateVal field if non-nil, zero value otherwise.

### GetRecordStateValOk

`func (o *PolicyEvalNotification) GetRecordStateValOk() (*string, bool)`

GetRecordStateValOk returns a tuple with the RecordStateVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStateVal

`func (o *PolicyEvalNotification) SetRecordStateVal(v string)`

SetRecordStateVal sets RecordStateVal field to given value.

### HasRecordStateVal

`func (o *PolicyEvalNotification) HasRecordStateVal() bool`

HasRecordStateVal returns a boolean if a field has been set.

### SetRecordStateValNil

`func (o *PolicyEvalNotification) SetRecordStateValNil(b bool)`

 SetRecordStateValNil sets the value for RecordStateVal to be an explicit nil

### UnsetRecordStateVal
`func (o *PolicyEvalNotification) UnsetRecordStateVal()`

UnsetRecordStateVal ensures that no value is present for RecordStateVal, not even an explicit nil
### GetTries

`func (o *PolicyEvalNotification) GetTries() int32`

GetTries returns the Tries field if non-nil, zero value otherwise.

### GetTriesOk

`func (o *PolicyEvalNotification) GetTriesOk() (*int32, bool)`

GetTriesOk returns a tuple with the Tries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTries

`func (o *PolicyEvalNotification) SetTries(v int32)`

SetTries sets Tries field to given value.

### HasTries

`func (o *PolicyEvalNotification) HasTries() bool`

HasTries returns a boolean if a field has been set.

### GetMaxTries

`func (o *PolicyEvalNotification) GetMaxTries() int32`

GetMaxTries returns the MaxTries field if non-nil, zero value otherwise.

### GetMaxTriesOk

`func (o *PolicyEvalNotification) GetMaxTriesOk() (*int32, bool)`

GetMaxTriesOk returns a tuple with the MaxTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTries

`func (o *PolicyEvalNotification) SetMaxTries(v int32)`

SetMaxTries sets MaxTries field to given value.

### HasMaxTries

`func (o *PolicyEvalNotification) HasMaxTries() bool`

HasMaxTries returns a boolean if a field has been set.

### GetData

`func (o *PolicyEvalNotification) GetData() PolicyEvalNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PolicyEvalNotification) GetDataOk() (*PolicyEvalNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PolicyEvalNotification) SetData(v PolicyEvalNotificationData)`

SetData sets Data field to given value.

### HasData

`func (o *PolicyEvalNotification) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


