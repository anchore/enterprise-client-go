# FeedMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of the feed | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date the metadata record was created in engine (first seen on source) | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Date the metadata was last updated | [optional] 
**Groups** | Pointer to [**[]FeedGroupMetadata**](FeedGroupMetadata.md) |  | [optional] 
**LastFullSync** | Pointer to **time.Time** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


