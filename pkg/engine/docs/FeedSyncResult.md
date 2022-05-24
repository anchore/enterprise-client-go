# FeedSyncResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feed** | Pointer to **string** | The name of the feed synced | [optional] 
**Status** | Pointer to **string** | The result of the sync operations, either co | [optional] 
**TotalTimeSeconds** | Pointer to **float32** | The duratin, in seconds, of the sync of the feed, the sum of all the group syncs | [optional] 
**Groups** | Pointer to [**[]GroupSyncResult**](GroupSyncResult.md) | Array of group sync results | [optional] 

## Methods

### NewFeedSyncResult

`func NewFeedSyncResult() *FeedSyncResult`

NewFeedSyncResult instantiates a new FeedSyncResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedSyncResultWithDefaults

`func NewFeedSyncResultWithDefaults() *FeedSyncResult`

NewFeedSyncResultWithDefaults instantiates a new FeedSyncResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeed

`func (o *FeedSyncResult) GetFeed() string`

GetFeed returns the Feed field if non-nil, zero value otherwise.

### GetFeedOk

`func (o *FeedSyncResult) GetFeedOk() (*string, bool)`

GetFeedOk returns a tuple with the Feed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeed

`func (o *FeedSyncResult) SetFeed(v string)`

SetFeed sets Feed field to given value.

### HasFeed

`func (o *FeedSyncResult) HasFeed() bool`

HasFeed returns a boolean if a field has been set.

### GetStatus

`func (o *FeedSyncResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FeedSyncResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FeedSyncResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FeedSyncResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalTimeSeconds

`func (o *FeedSyncResult) GetTotalTimeSeconds() float32`

GetTotalTimeSeconds returns the TotalTimeSeconds field if non-nil, zero value otherwise.

### GetTotalTimeSecondsOk

`func (o *FeedSyncResult) GetTotalTimeSecondsOk() (*float32, bool)`

GetTotalTimeSecondsOk returns a tuple with the TotalTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeSeconds

`func (o *FeedSyncResult) SetTotalTimeSeconds(v float32)`

SetTotalTimeSeconds sets TotalTimeSeconds field to given value.

### HasTotalTimeSeconds

`func (o *FeedSyncResult) HasTotalTimeSeconds() bool`

HasTotalTimeSeconds returns a boolean if a field has been set.

### GetGroups

`func (o *FeedSyncResult) GetGroups() []GroupSyncResult`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *FeedSyncResult) GetGroupsOk() (*[]GroupSyncResult, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *FeedSyncResult) SetGroups(v []GroupSyncResult)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *FeedSyncResult) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


