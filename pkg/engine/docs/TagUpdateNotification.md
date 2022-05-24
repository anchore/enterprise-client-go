# TagUpdateNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**DataId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**LastUpdated** | Pointer to **int32** |  | [optional] 
**RecordStateKey** | Pointer to **string** |  | [optional] [default to "active"]
**RecordStateVal** | Pointer to **NullableString** |  | [optional] 
**Tries** | Pointer to **int32** |  | [optional] 
**MaxTries** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**TagUpdateNotificationData**](TagUpdateNotificationData.md) |  | [optional] 

## Methods

### NewTagUpdateNotification

`func NewTagUpdateNotification() *TagUpdateNotification`

NewTagUpdateNotification instantiates a new TagUpdateNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagUpdateNotificationWithDefaults

`func NewTagUpdateNotificationWithDefaults() *TagUpdateNotification`

NewTagUpdateNotificationWithDefaults instantiates a new TagUpdateNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueId

`func (o *TagUpdateNotification) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *TagUpdateNotification) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *TagUpdateNotification) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *TagUpdateNotification) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetUserId

`func (o *TagUpdateNotification) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TagUpdateNotification) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TagUpdateNotification) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TagUpdateNotification) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDataId

`func (o *TagUpdateNotification) GetDataId() string`

GetDataId returns the DataId field if non-nil, zero value otherwise.

### GetDataIdOk

`func (o *TagUpdateNotification) GetDataIdOk() (*string, bool)`

GetDataIdOk returns a tuple with the DataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataId

`func (o *TagUpdateNotification) SetDataId(v string)`

SetDataId sets DataId field to given value.

### HasDataId

`func (o *TagUpdateNotification) HasDataId() bool`

HasDataId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TagUpdateNotification) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagUpdateNotification) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagUpdateNotification) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TagUpdateNotification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *TagUpdateNotification) GetLastUpdated() int32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TagUpdateNotification) GetLastUpdatedOk() (*int32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TagUpdateNotification) SetLastUpdated(v int32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *TagUpdateNotification) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRecordStateKey

`func (o *TagUpdateNotification) GetRecordStateKey() string`

GetRecordStateKey returns the RecordStateKey field if non-nil, zero value otherwise.

### GetRecordStateKeyOk

`func (o *TagUpdateNotification) GetRecordStateKeyOk() (*string, bool)`

GetRecordStateKeyOk returns a tuple with the RecordStateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStateKey

`func (o *TagUpdateNotification) SetRecordStateKey(v string)`

SetRecordStateKey sets RecordStateKey field to given value.

### HasRecordStateKey

`func (o *TagUpdateNotification) HasRecordStateKey() bool`

HasRecordStateKey returns a boolean if a field has been set.

### GetRecordStateVal

`func (o *TagUpdateNotification) GetRecordStateVal() string`

GetRecordStateVal returns the RecordStateVal field if non-nil, zero value otherwise.

### GetRecordStateValOk

`func (o *TagUpdateNotification) GetRecordStateValOk() (*string, bool)`

GetRecordStateValOk returns a tuple with the RecordStateVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStateVal

`func (o *TagUpdateNotification) SetRecordStateVal(v string)`

SetRecordStateVal sets RecordStateVal field to given value.

### HasRecordStateVal

`func (o *TagUpdateNotification) HasRecordStateVal() bool`

HasRecordStateVal returns a boolean if a field has been set.

### SetRecordStateValNil

`func (o *TagUpdateNotification) SetRecordStateValNil(b bool)`

 SetRecordStateValNil sets the value for RecordStateVal to be an explicit nil

### UnsetRecordStateVal
`func (o *TagUpdateNotification) UnsetRecordStateVal()`

UnsetRecordStateVal ensures that no value is present for RecordStateVal, not even an explicit nil
### GetTries

`func (o *TagUpdateNotification) GetTries() int32`

GetTries returns the Tries field if non-nil, zero value otherwise.

### GetTriesOk

`func (o *TagUpdateNotification) GetTriesOk() (*int32, bool)`

GetTriesOk returns a tuple with the Tries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTries

`func (o *TagUpdateNotification) SetTries(v int32)`

SetTries sets Tries field to given value.

### HasTries

`func (o *TagUpdateNotification) HasTries() bool`

HasTries returns a boolean if a field has been set.

### GetMaxTries

`func (o *TagUpdateNotification) GetMaxTries() int32`

GetMaxTries returns the MaxTries field if non-nil, zero value otherwise.

### GetMaxTriesOk

`func (o *TagUpdateNotification) GetMaxTriesOk() (*int32, bool)`

GetMaxTriesOk returns a tuple with the MaxTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTries

`func (o *TagUpdateNotification) SetMaxTries(v int32)`

SetMaxTries sets MaxTries field to given value.

### HasMaxTries

`func (o *TagUpdateNotification) HasMaxTries() bool`

HasMaxTries returns a boolean if a field has been set.

### GetData

`func (o *TagUpdateNotification) GetData() TagUpdateNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TagUpdateNotification) GetDataOk() (*TagUpdateNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TagUpdateNotification) SetData(v TagUpdateNotificationData)`

SetData sets Data field to given value.

### HasData

`func (o *TagUpdateNotification) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


