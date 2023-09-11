# GroupSyncResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | The name of the group | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedImageCount** | Pointer to **int32** | The number of images updated by the this group sync, across all accounts. This is typically only non-zero for vulnerability feeds which update images&#39; vulnerability results during the sync. | [optional] 
**UpdatedRecordCount** | Pointer to **int32** | The number of feed data records synced down as either updates or new records | [optional] 
**TotalTimeSeconds** | Pointer to **float32** | The duration of the group sync in seconds | [optional] 

## Methods

### NewGroupSyncResult

`func NewGroupSyncResult() *GroupSyncResult`

NewGroupSyncResult instantiates a new GroupSyncResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSyncResultWithDefaults

`func NewGroupSyncResultWithDefaults() *GroupSyncResult`

NewGroupSyncResultWithDefaults instantiates a new GroupSyncResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *GroupSyncResult) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GroupSyncResult) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GroupSyncResult) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GroupSyncResult) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetStatus

`func (o *GroupSyncResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GroupSyncResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GroupSyncResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GroupSyncResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedImageCount

`func (o *GroupSyncResult) GetUpdatedImageCount() int32`

GetUpdatedImageCount returns the UpdatedImageCount field if non-nil, zero value otherwise.

### GetUpdatedImageCountOk

`func (o *GroupSyncResult) GetUpdatedImageCountOk() (*int32, bool)`

GetUpdatedImageCountOk returns a tuple with the UpdatedImageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedImageCount

`func (o *GroupSyncResult) SetUpdatedImageCount(v int32)`

SetUpdatedImageCount sets UpdatedImageCount field to given value.

### HasUpdatedImageCount

`func (o *GroupSyncResult) HasUpdatedImageCount() bool`

HasUpdatedImageCount returns a boolean if a field has been set.

### GetUpdatedRecordCount

`func (o *GroupSyncResult) GetUpdatedRecordCount() int32`

GetUpdatedRecordCount returns the UpdatedRecordCount field if non-nil, zero value otherwise.

### GetUpdatedRecordCountOk

`func (o *GroupSyncResult) GetUpdatedRecordCountOk() (*int32, bool)`

GetUpdatedRecordCountOk returns a tuple with the UpdatedRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedRecordCount

`func (o *GroupSyncResult) SetUpdatedRecordCount(v int32)`

SetUpdatedRecordCount sets UpdatedRecordCount field to given value.

### HasUpdatedRecordCount

`func (o *GroupSyncResult) HasUpdatedRecordCount() bool`

HasUpdatedRecordCount returns a boolean if a field has been set.

### GetTotalTimeSeconds

`func (o *GroupSyncResult) GetTotalTimeSeconds() float32`

GetTotalTimeSeconds returns the TotalTimeSeconds field if non-nil, zero value otherwise.

### GetTotalTimeSecondsOk

`func (o *GroupSyncResult) GetTotalTimeSecondsOk() (*float32, bool)`

GetTotalTimeSecondsOk returns a tuple with the TotalTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeSeconds

`func (o *GroupSyncResult) SetTotalTimeSeconds(v float32)`

SetTotalTimeSeconds sets TotalTimeSeconds field to given value.

### HasTotalTimeSeconds

`func (o *GroupSyncResult) HasTotalTimeSeconds() bool`

HasTotalTimeSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


