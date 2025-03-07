# AnalysisUpdateNotification

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
**Data** | Pointer to [**AnalysisUpdateNotificationData**](AnalysisUpdateNotificationData.md) |  | [optional] 

## Methods

### NewAnalysisUpdateNotification

`func NewAnalysisUpdateNotification() *AnalysisUpdateNotification`

NewAnalysisUpdateNotification instantiates a new AnalysisUpdateNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisUpdateNotificationWithDefaults

`func NewAnalysisUpdateNotificationWithDefaults() *AnalysisUpdateNotification`

NewAnalysisUpdateNotificationWithDefaults instantiates a new AnalysisUpdateNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueId

`func (o *AnalysisUpdateNotification) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *AnalysisUpdateNotification) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *AnalysisUpdateNotification) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *AnalysisUpdateNotification) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetAccountName

`func (o *AnalysisUpdateNotification) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AnalysisUpdateNotification) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AnalysisUpdateNotification) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AnalysisUpdateNotification) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetDataId

`func (o *AnalysisUpdateNotification) GetDataId() string`

GetDataId returns the DataId field if non-nil, zero value otherwise.

### GetDataIdOk

`func (o *AnalysisUpdateNotification) GetDataIdOk() (*string, bool)`

GetDataIdOk returns a tuple with the DataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataId

`func (o *AnalysisUpdateNotification) SetDataId(v string)`

SetDataId sets DataId field to given value.

### HasDataId

`func (o *AnalysisUpdateNotification) HasDataId() bool`

HasDataId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AnalysisUpdateNotification) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnalysisUpdateNotification) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnalysisUpdateNotification) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AnalysisUpdateNotification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AnalysisUpdateNotification) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AnalysisUpdateNotification) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AnalysisUpdateNotification) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AnalysisUpdateNotification) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRecordStateKey

`func (o *AnalysisUpdateNotification) GetRecordStateKey() string`

GetRecordStateKey returns the RecordStateKey field if non-nil, zero value otherwise.

### GetRecordStateKeyOk

`func (o *AnalysisUpdateNotification) GetRecordStateKeyOk() (*string, bool)`

GetRecordStateKeyOk returns a tuple with the RecordStateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStateKey

`func (o *AnalysisUpdateNotification) SetRecordStateKey(v string)`

SetRecordStateKey sets RecordStateKey field to given value.

### HasRecordStateKey

`func (o *AnalysisUpdateNotification) HasRecordStateKey() bool`

HasRecordStateKey returns a boolean if a field has been set.

### GetRecordStateVal

`func (o *AnalysisUpdateNotification) GetRecordStateVal() string`

GetRecordStateVal returns the RecordStateVal field if non-nil, zero value otherwise.

### GetRecordStateValOk

`func (o *AnalysisUpdateNotification) GetRecordStateValOk() (*string, bool)`

GetRecordStateValOk returns a tuple with the RecordStateVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStateVal

`func (o *AnalysisUpdateNotification) SetRecordStateVal(v string)`

SetRecordStateVal sets RecordStateVal field to given value.

### HasRecordStateVal

`func (o *AnalysisUpdateNotification) HasRecordStateVal() bool`

HasRecordStateVal returns a boolean if a field has been set.

### SetRecordStateValNil

`func (o *AnalysisUpdateNotification) SetRecordStateValNil(b bool)`

 SetRecordStateValNil sets the value for RecordStateVal to be an explicit nil

### UnsetRecordStateVal
`func (o *AnalysisUpdateNotification) UnsetRecordStateVal()`

UnsetRecordStateVal ensures that no value is present for RecordStateVal, not even an explicit nil
### GetTries

`func (o *AnalysisUpdateNotification) GetTries() int64`

GetTries returns the Tries field if non-nil, zero value otherwise.

### GetTriesOk

`func (o *AnalysisUpdateNotification) GetTriesOk() (*int64, bool)`

GetTriesOk returns a tuple with the Tries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTries

`func (o *AnalysisUpdateNotification) SetTries(v int64)`

SetTries sets Tries field to given value.

### HasTries

`func (o *AnalysisUpdateNotification) HasTries() bool`

HasTries returns a boolean if a field has been set.

### GetMaxTries

`func (o *AnalysisUpdateNotification) GetMaxTries() int64`

GetMaxTries returns the MaxTries field if non-nil, zero value otherwise.

### GetMaxTriesOk

`func (o *AnalysisUpdateNotification) GetMaxTriesOk() (*int64, bool)`

GetMaxTriesOk returns a tuple with the MaxTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTries

`func (o *AnalysisUpdateNotification) SetMaxTries(v int64)`

SetMaxTries sets MaxTries field to given value.

### HasMaxTries

`func (o *AnalysisUpdateNotification) HasMaxTries() bool`

HasMaxTries returns a boolean if a field has been set.

### GetData

`func (o *AnalysisUpdateNotification) GetData() AnalysisUpdateNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnalysisUpdateNotification) GetDataOk() (*AnalysisUpdateNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnalysisUpdateNotification) SetData(v AnalysisUpdateNotificationData)`

SetData sets Data field to given value.

### HasData

`func (o *AnalysisUpdateNotification) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


