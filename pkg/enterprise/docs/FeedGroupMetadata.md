# FeedGroupMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this feed group. | [optional] 
**DataServiceBuiltAt** | Pointer to **time.Time** | The time when the Anchore Data Service started to build this feed group. | [optional] 
**EnterpriseReceivedAt** | Pointer to **time.Time** | The time when the Enterprise received this feed group either via the data-syncer service or air gapped workflow. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Deprecated - The time when Enterprise uploaded this feed group. | [optional] 
**LastSync** | Pointer to **time.Time** | Deprecated - The last successful update for this feed group. | [optional] 
**RecordCount** | Pointer to **int32** | The number of records in this feed group. | [optional] 
**Enabled** | Pointer to **bool** | Deprecated - If group is enabled. | [optional] 

## Methods

### NewFeedGroupMetadata

`func NewFeedGroupMetadata() *FeedGroupMetadata`

NewFeedGroupMetadata instantiates a new FeedGroupMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedGroupMetadataWithDefaults

`func NewFeedGroupMetadataWithDefaults() *FeedGroupMetadata`

NewFeedGroupMetadataWithDefaults instantiates a new FeedGroupMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeedGroupMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeedGroupMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeedGroupMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeedGroupMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataServiceBuiltAt

`func (o *FeedGroupMetadata) GetDataServiceBuiltAt() time.Time`

GetDataServiceBuiltAt returns the DataServiceBuiltAt field if non-nil, zero value otherwise.

### GetDataServiceBuiltAtOk

`func (o *FeedGroupMetadata) GetDataServiceBuiltAtOk() (*time.Time, bool)`

GetDataServiceBuiltAtOk returns a tuple with the DataServiceBuiltAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceBuiltAt

`func (o *FeedGroupMetadata) SetDataServiceBuiltAt(v time.Time)`

SetDataServiceBuiltAt sets DataServiceBuiltAt field to given value.

### HasDataServiceBuiltAt

`func (o *FeedGroupMetadata) HasDataServiceBuiltAt() bool`

HasDataServiceBuiltAt returns a boolean if a field has been set.

### GetEnterpriseReceivedAt

`func (o *FeedGroupMetadata) GetEnterpriseReceivedAt() time.Time`

GetEnterpriseReceivedAt returns the EnterpriseReceivedAt field if non-nil, zero value otherwise.

### GetEnterpriseReceivedAtOk

`func (o *FeedGroupMetadata) GetEnterpriseReceivedAtOk() (*time.Time, bool)`

GetEnterpriseReceivedAtOk returns a tuple with the EnterpriseReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseReceivedAt

`func (o *FeedGroupMetadata) SetEnterpriseReceivedAt(v time.Time)`

SetEnterpriseReceivedAt sets EnterpriseReceivedAt field to given value.

### HasEnterpriseReceivedAt

`func (o *FeedGroupMetadata) HasEnterpriseReceivedAt() bool`

HasEnterpriseReceivedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FeedGroupMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeedGroupMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeedGroupMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FeedGroupMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastSync

`func (o *FeedGroupMetadata) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *FeedGroupMetadata) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *FeedGroupMetadata) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *FeedGroupMetadata) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetRecordCount

`func (o *FeedGroupMetadata) GetRecordCount() int32`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *FeedGroupMetadata) GetRecordCountOk() (*int32, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *FeedGroupMetadata) SetRecordCount(v int32)`

SetRecordCount sets RecordCount field to given value.

### HasRecordCount

`func (o *FeedGroupMetadata) HasRecordCount() bool`

HasRecordCount returns a boolean if a field has been set.

### GetEnabled

`func (o *FeedGroupMetadata) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FeedGroupMetadata) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FeedGroupMetadata) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FeedGroupMetadata) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


