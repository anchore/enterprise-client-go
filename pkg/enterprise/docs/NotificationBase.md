# NotificationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueId** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**DataId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**LastUpdated** | Pointer to **int64** |  | [optional] 
**RecordStateKey** | Pointer to **string** |  | [optional] [default to "active"]
**RecordStateVal** | Pointer to **NullableString** |  | [optional] 
**Tries** | Pointer to **int64** |  | [optional] 
**MaxTries** | Pointer to **int64** |  | [optional] 

## Methods

### NewNotificationBase

`func NewNotificationBase() *NotificationBase`

NewNotificationBase instantiates a new NotificationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationBaseWithDefaults

`func NewNotificationBaseWithDefaults() *NotificationBase`

NewNotificationBaseWithDefaults instantiates a new NotificationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueId

`func (o *NotificationBase) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *NotificationBase) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *NotificationBase) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *NotificationBase) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetAccountName

`func (o *NotificationBase) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *NotificationBase) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *NotificationBase) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *NotificationBase) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetDataId

`func (o *NotificationBase) GetDataId() string`

GetDataId returns the DataId field if non-nil, zero value otherwise.

### GetDataIdOk

`func (o *NotificationBase) GetDataIdOk() (*string, bool)`

GetDataIdOk returns a tuple with the DataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataId

`func (o *NotificationBase) SetDataId(v string)`

SetDataId sets DataId field to given value.

### HasDataId

`func (o *NotificationBase) HasDataId() bool`

HasDataId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationBase) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationBase) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationBase) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationBase) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NotificationBase) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NotificationBase) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NotificationBase) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NotificationBase) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRecordStateKey

`func (o *NotificationBase) GetRecordStateKey() string`

GetRecordStateKey returns the RecordStateKey field if non-nil, zero value otherwise.

### GetRecordStateKeyOk

`func (o *NotificationBase) GetRecordStateKeyOk() (*string, bool)`

GetRecordStateKeyOk returns a tuple with the RecordStateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStateKey

`func (o *NotificationBase) SetRecordStateKey(v string)`

SetRecordStateKey sets RecordStateKey field to given value.

### HasRecordStateKey

`func (o *NotificationBase) HasRecordStateKey() bool`

HasRecordStateKey returns a boolean if a field has been set.

### GetRecordStateVal

`func (o *NotificationBase) GetRecordStateVal() string`

GetRecordStateVal returns the RecordStateVal field if non-nil, zero value otherwise.

### GetRecordStateValOk

`func (o *NotificationBase) GetRecordStateValOk() (*string, bool)`

GetRecordStateValOk returns a tuple with the RecordStateVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStateVal

`func (o *NotificationBase) SetRecordStateVal(v string)`

SetRecordStateVal sets RecordStateVal field to given value.

### HasRecordStateVal

`func (o *NotificationBase) HasRecordStateVal() bool`

HasRecordStateVal returns a boolean if a field has been set.

### SetRecordStateValNil

`func (o *NotificationBase) SetRecordStateValNil(b bool)`

 SetRecordStateValNil sets the value for RecordStateVal to be an explicit nil

### UnsetRecordStateVal
`func (o *NotificationBase) UnsetRecordStateVal()`

UnsetRecordStateVal ensures that no value is present for RecordStateVal, not even an explicit nil
### GetTries

`func (o *NotificationBase) GetTries() int64`

GetTries returns the Tries field if non-nil, zero value otherwise.

### GetTriesOk

`func (o *NotificationBase) GetTriesOk() (*int64, bool)`

GetTriesOk returns a tuple with the Tries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTries

`func (o *NotificationBase) SetTries(v int64)`

SetTries sets Tries field to given value.

### HasTries

`func (o *NotificationBase) HasTries() bool`

HasTries returns a boolean if a field has been set.

### GetMaxTries

`func (o *NotificationBase) GetMaxTries() int64`

GetMaxTries returns the MaxTries field if non-nil, zero value otherwise.

### GetMaxTriesOk

`func (o *NotificationBase) GetMaxTriesOk() (*int64, bool)`

GetMaxTriesOk returns a tuple with the MaxTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTries

`func (o *NotificationBase) SetMaxTries(v int64)`

SetMaxTries sets MaxTries field to given value.

### HasMaxTries

`func (o *NotificationBase) HasMaxTries() bool`

HasMaxTries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


