# AnchoreSummaryImagesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registry** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**TagCount** | Pointer to **int32** |  | [optional] 
**LatestAnalyzedAt** | Pointer to **time.Time** |  | [optional] 
**Alert** | Pointer to **bool** |  | [optional] 
**AlertId** | Pointer to **string** |  | [optional] 
**WatchRepo** | Pointer to **bool** |  | [optional] 
**WatchRepoId** | Pointer to **string** |  | [optional] 

## Methods

### NewAnchoreSummaryImagesItem

`func NewAnchoreSummaryImagesItem() *AnchoreSummaryImagesItem`

NewAnchoreSummaryImagesItem instantiates a new AnchoreSummaryImagesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnchoreSummaryImagesItemWithDefaults

`func NewAnchoreSummaryImagesItemWithDefaults() *AnchoreSummaryImagesItem`

NewAnchoreSummaryImagesItemWithDefaults instantiates a new AnchoreSummaryImagesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistry

`func (o *AnchoreSummaryImagesItem) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *AnchoreSummaryImagesItem) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *AnchoreSummaryImagesItem) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *AnchoreSummaryImagesItem) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRepo

`func (o *AnchoreSummaryImagesItem) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AnchoreSummaryImagesItem) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AnchoreSummaryImagesItem) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AnchoreSummaryImagesItem) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetTagCount

`func (o *AnchoreSummaryImagesItem) GetTagCount() int32`

GetTagCount returns the TagCount field if non-nil, zero value otherwise.

### GetTagCountOk

`func (o *AnchoreSummaryImagesItem) GetTagCountOk() (*int32, bool)`

GetTagCountOk returns a tuple with the TagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCount

`func (o *AnchoreSummaryImagesItem) SetTagCount(v int32)`

SetTagCount sets TagCount field to given value.

### HasTagCount

`func (o *AnchoreSummaryImagesItem) HasTagCount() bool`

HasTagCount returns a boolean if a field has been set.

### GetLatestAnalyzedAt

`func (o *AnchoreSummaryImagesItem) GetLatestAnalyzedAt() time.Time`

GetLatestAnalyzedAt returns the LatestAnalyzedAt field if non-nil, zero value otherwise.

### GetLatestAnalyzedAtOk

`func (o *AnchoreSummaryImagesItem) GetLatestAnalyzedAtOk() (*time.Time, bool)`

GetLatestAnalyzedAtOk returns a tuple with the LatestAnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAnalyzedAt

`func (o *AnchoreSummaryImagesItem) SetLatestAnalyzedAt(v time.Time)`

SetLatestAnalyzedAt sets LatestAnalyzedAt field to given value.

### HasLatestAnalyzedAt

`func (o *AnchoreSummaryImagesItem) HasLatestAnalyzedAt() bool`

HasLatestAnalyzedAt returns a boolean if a field has been set.

### GetAlert

`func (o *AnchoreSummaryImagesItem) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *AnchoreSummaryImagesItem) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *AnchoreSummaryImagesItem) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *AnchoreSummaryImagesItem) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetAlertId

`func (o *AnchoreSummaryImagesItem) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *AnchoreSummaryImagesItem) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *AnchoreSummaryImagesItem) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *AnchoreSummaryImagesItem) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetWatchRepo

`func (o *AnchoreSummaryImagesItem) GetWatchRepo() bool`

GetWatchRepo returns the WatchRepo field if non-nil, zero value otherwise.

### GetWatchRepoOk

`func (o *AnchoreSummaryImagesItem) GetWatchRepoOk() (*bool, bool)`

GetWatchRepoOk returns a tuple with the WatchRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchRepo

`func (o *AnchoreSummaryImagesItem) SetWatchRepo(v bool)`

SetWatchRepo sets WatchRepo field to given value.

### HasWatchRepo

`func (o *AnchoreSummaryImagesItem) HasWatchRepo() bool`

HasWatchRepo returns a boolean if a field has been set.

### GetWatchRepoId

`func (o *AnchoreSummaryImagesItem) GetWatchRepoId() string`

GetWatchRepoId returns the WatchRepoId field if non-nil, zero value otherwise.

### GetWatchRepoIdOk

`func (o *AnchoreSummaryImagesItem) GetWatchRepoIdOk() (*string, bool)`

GetWatchRepoIdOk returns a tuple with the WatchRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchRepoId

`func (o *AnchoreSummaryImagesItem) SetWatchRepoId(v string)`

SetWatchRepoId sets WatchRepoId field to given value.

### HasWatchRepoId

`func (o *AnchoreSummaryImagesItem) HasWatchRepoId() bool`

HasWatchRepoId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


