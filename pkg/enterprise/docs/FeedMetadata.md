# FeedMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the feed. | [optional] 
**Version** | Pointer to **string** | The version of the feed. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Deprecated - The time when Enterprise uploaded this feed. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Deprecated - The last time the policy-engine service pinged the feed service to see if there was a new grypedb available. | [optional] 
**DataServiceBuiltAt** | Pointer to **time.Time** | The time when the Anchore Data Service started to build this feed. | [optional] 
**EnterpriseReceivedAt** | Pointer to **time.Time** | The time when the Enterprise received this feed either via the data-syncer service or air gapped workflow. | [optional] 
**Groups** | Pointer to [**[]FeedGroupMetadata**](FeedGroupMetadata.md) |  | [optional] 
**LastFullSync** | Pointer to **time.Time** | Deprecated - The last time that policy-engine service downloaded a new grypedb. | [optional] 
**Enabled** | Pointer to **bool** | Deprecated - If feed is enabled | [optional] 
**DatasetName** | Pointer to **string** | The name of the dataset that provides this feed | [optional] 
**DatasetChecksum** | Pointer to **string** | The checksum of the dataset that provides this feed | [optional] 

## Methods

### NewFeedMetadata

`func NewFeedMetadata() *FeedMetadata`

NewFeedMetadata instantiates a new FeedMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedMetadataWithDefaults

`func NewFeedMetadataWithDefaults() *FeedMetadata`

NewFeedMetadataWithDefaults instantiates a new FeedMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeedMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeedMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeedMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeedMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *FeedMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeedMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeedMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FeedMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FeedMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeedMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeedMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FeedMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FeedMetadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FeedMetadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FeedMetadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FeedMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDataServiceBuiltAt

`func (o *FeedMetadata) GetDataServiceBuiltAt() time.Time`

GetDataServiceBuiltAt returns the DataServiceBuiltAt field if non-nil, zero value otherwise.

### GetDataServiceBuiltAtOk

`func (o *FeedMetadata) GetDataServiceBuiltAtOk() (*time.Time, bool)`

GetDataServiceBuiltAtOk returns a tuple with the DataServiceBuiltAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceBuiltAt

`func (o *FeedMetadata) SetDataServiceBuiltAt(v time.Time)`

SetDataServiceBuiltAt sets DataServiceBuiltAt field to given value.

### HasDataServiceBuiltAt

`func (o *FeedMetadata) HasDataServiceBuiltAt() bool`

HasDataServiceBuiltAt returns a boolean if a field has been set.

### GetEnterpriseReceivedAt

`func (o *FeedMetadata) GetEnterpriseReceivedAt() time.Time`

GetEnterpriseReceivedAt returns the EnterpriseReceivedAt field if non-nil, zero value otherwise.

### GetEnterpriseReceivedAtOk

`func (o *FeedMetadata) GetEnterpriseReceivedAtOk() (*time.Time, bool)`

GetEnterpriseReceivedAtOk returns a tuple with the EnterpriseReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseReceivedAt

`func (o *FeedMetadata) SetEnterpriseReceivedAt(v time.Time)`

SetEnterpriseReceivedAt sets EnterpriseReceivedAt field to given value.

### HasEnterpriseReceivedAt

`func (o *FeedMetadata) HasEnterpriseReceivedAt() bool`

HasEnterpriseReceivedAt returns a boolean if a field has been set.

### GetGroups

`func (o *FeedMetadata) GetGroups() []FeedGroupMetadata`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *FeedMetadata) GetGroupsOk() (*[]FeedGroupMetadata, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *FeedMetadata) SetGroups(v []FeedGroupMetadata)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *FeedMetadata) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetLastFullSync

`func (o *FeedMetadata) GetLastFullSync() time.Time`

GetLastFullSync returns the LastFullSync field if non-nil, zero value otherwise.

### GetLastFullSyncOk

`func (o *FeedMetadata) GetLastFullSyncOk() (*time.Time, bool)`

GetLastFullSyncOk returns a tuple with the LastFullSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFullSync

`func (o *FeedMetadata) SetLastFullSync(v time.Time)`

SetLastFullSync sets LastFullSync field to given value.

### HasLastFullSync

`func (o *FeedMetadata) HasLastFullSync() bool`

HasLastFullSync returns a boolean if a field has been set.

### GetEnabled

`func (o *FeedMetadata) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FeedMetadata) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FeedMetadata) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FeedMetadata) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDatasetName

`func (o *FeedMetadata) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *FeedMetadata) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *FeedMetadata) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *FeedMetadata) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.

### GetDatasetChecksum

`func (o *FeedMetadata) GetDatasetChecksum() string`

GetDatasetChecksum returns the DatasetChecksum field if non-nil, zero value otherwise.

### GetDatasetChecksumOk

`func (o *FeedMetadata) GetDatasetChecksumOk() (*string, bool)`

GetDatasetChecksumOk returns a tuple with the DatasetChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetChecksum

`func (o *FeedMetadata) SetDatasetChecksum(v string)`

SetDatasetChecksum sets DatasetChecksum field to given value.

### HasDatasetChecksum

`func (o *FeedMetadata) HasDatasetChecksum() bool`

HasDatasetChecksum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


