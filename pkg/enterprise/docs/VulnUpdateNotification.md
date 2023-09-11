# VulnUpdateNotification

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
**Data** | Pointer to [**VulnUpdateNotificationData**](VulnUpdateNotificationData.md) |  | [optional] 

## Methods

### NewVulnUpdateNotification

`func NewVulnUpdateNotification() *VulnUpdateNotification`

NewVulnUpdateNotification instantiates a new VulnUpdateNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnUpdateNotificationWithDefaults

`func NewVulnUpdateNotificationWithDefaults() *VulnUpdateNotification`

NewVulnUpdateNotificationWithDefaults instantiates a new VulnUpdateNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueId

`func (o *VulnUpdateNotification) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *VulnUpdateNotification) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *VulnUpdateNotification) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *VulnUpdateNotification) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetAccountName

`func (o *VulnUpdateNotification) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *VulnUpdateNotification) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *VulnUpdateNotification) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *VulnUpdateNotification) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetDataId

`func (o *VulnUpdateNotification) GetDataId() string`

GetDataId returns the DataId field if non-nil, zero value otherwise.

### GetDataIdOk

`func (o *VulnUpdateNotification) GetDataIdOk() (*string, bool)`

GetDataIdOk returns a tuple with the DataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataId

`func (o *VulnUpdateNotification) SetDataId(v string)`

SetDataId sets DataId field to given value.

### HasDataId

`func (o *VulnUpdateNotification) HasDataId() bool`

HasDataId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VulnUpdateNotification) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VulnUpdateNotification) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VulnUpdateNotification) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VulnUpdateNotification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *VulnUpdateNotification) GetLastUpdated() int32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VulnUpdateNotification) GetLastUpdatedOk() (*int32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VulnUpdateNotification) SetLastUpdated(v int32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *VulnUpdateNotification) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRecordStateKey

`func (o *VulnUpdateNotification) GetRecordStateKey() string`

GetRecordStateKey returns the RecordStateKey field if non-nil, zero value otherwise.

### GetRecordStateKeyOk

`func (o *VulnUpdateNotification) GetRecordStateKeyOk() (*string, bool)`

GetRecordStateKeyOk returns a tuple with the RecordStateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStateKey

`func (o *VulnUpdateNotification) SetRecordStateKey(v string)`

SetRecordStateKey sets RecordStateKey field to given value.

### HasRecordStateKey

`func (o *VulnUpdateNotification) HasRecordStateKey() bool`

HasRecordStateKey returns a boolean if a field has been set.

### GetRecordStateVal

`func (o *VulnUpdateNotification) GetRecordStateVal() string`

GetRecordStateVal returns the RecordStateVal field if non-nil, zero value otherwise.

### GetRecordStateValOk

`func (o *VulnUpdateNotification) GetRecordStateValOk() (*string, bool)`

GetRecordStateValOk returns a tuple with the RecordStateVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStateVal

`func (o *VulnUpdateNotification) SetRecordStateVal(v string)`

SetRecordStateVal sets RecordStateVal field to given value.

### HasRecordStateVal

`func (o *VulnUpdateNotification) HasRecordStateVal() bool`

HasRecordStateVal returns a boolean if a field has been set.

### SetRecordStateValNil

`func (o *VulnUpdateNotification) SetRecordStateValNil(b bool)`

 SetRecordStateValNil sets the value for RecordStateVal to be an explicit nil

### UnsetRecordStateVal
`func (o *VulnUpdateNotification) UnsetRecordStateVal()`

UnsetRecordStateVal ensures that no value is present for RecordStateVal, not even an explicit nil
### GetTries

`func (o *VulnUpdateNotification) GetTries() int32`

GetTries returns the Tries field if non-nil, zero value otherwise.

### GetTriesOk

`func (o *VulnUpdateNotification) GetTriesOk() (*int32, bool)`

GetTriesOk returns a tuple with the Tries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTries

`func (o *VulnUpdateNotification) SetTries(v int32)`

SetTries sets Tries field to given value.

### HasTries

`func (o *VulnUpdateNotification) HasTries() bool`

HasTries returns a boolean if a field has been set.

### GetMaxTries

`func (o *VulnUpdateNotification) GetMaxTries() int32`

GetMaxTries returns the MaxTries field if non-nil, zero value otherwise.

### GetMaxTriesOk

`func (o *VulnUpdateNotification) GetMaxTriesOk() (*int32, bool)`

GetMaxTriesOk returns a tuple with the MaxTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTries

`func (o *VulnUpdateNotification) SetMaxTries(v int32)`

SetMaxTries sets MaxTries field to given value.

### HasMaxTries

`func (o *VulnUpdateNotification) HasMaxTries() bool`

HasMaxTries returns a boolean if a field has been set.

### GetData

`func (o *VulnUpdateNotification) GetData() VulnUpdateNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VulnUpdateNotification) GetDataOk() (*VulnUpdateNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VulnUpdateNotification) SetData(v VulnUpdateNotificationData)`

SetData sets Data field to given value.

### HasData

`func (o *VulnUpdateNotification) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


