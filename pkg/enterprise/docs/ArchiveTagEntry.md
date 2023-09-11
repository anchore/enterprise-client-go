# ArchiveTagEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PullString** | Pointer to **string** | The pullable string for the tag. E.g. \&quot;docker.io/library/node:latest\&quot; | [optional] 
**Registry** | Pointer to **string** | The registry hostname:port section of the pull string | [optional] 
**Repository** | Pointer to **string** | The repository section of the pull string | [optional] 
**Tag** | Pointer to **string** | The tag-only section of the pull string | [optional] 
**DetectedAt** | Pointer to **time.Time** | The timestamp at which the Anchore Engine detected this tag was mapped to the image digest. Does not necessarily indicate when the tag was actually pushed to the registry. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp at which Anchore Engine archived this image digest. | [optional] 
**LastUpdated** | Pointer to **time.Time** | The timestamp that the last change was made to this record. | [optional] 

## Methods

### NewArchiveTagEntry

`func NewArchiveTagEntry() *ArchiveTagEntry`

NewArchiveTagEntry instantiates a new ArchiveTagEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveTagEntryWithDefaults

`func NewArchiveTagEntryWithDefaults() *ArchiveTagEntry`

NewArchiveTagEntryWithDefaults instantiates a new ArchiveTagEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPullString

`func (o *ArchiveTagEntry) GetPullString() string`

GetPullString returns the PullString field if non-nil, zero value otherwise.

### GetPullStringOk

`func (o *ArchiveTagEntry) GetPullStringOk() (*string, bool)`

GetPullStringOk returns a tuple with the PullString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullString

`func (o *ArchiveTagEntry) SetPullString(v string)`

SetPullString sets PullString field to given value.

### HasPullString

`func (o *ArchiveTagEntry) HasPullString() bool`

HasPullString returns a boolean if a field has been set.

### GetRegistry

`func (o *ArchiveTagEntry) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *ArchiveTagEntry) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *ArchiveTagEntry) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *ArchiveTagEntry) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRepository

`func (o *ArchiveTagEntry) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ArchiveTagEntry) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ArchiveTagEntry) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ArchiveTagEntry) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTag

`func (o *ArchiveTagEntry) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ArchiveTagEntry) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ArchiveTagEntry) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ArchiveTagEntry) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetDetectedAt

`func (o *ArchiveTagEntry) GetDetectedAt() time.Time`

GetDetectedAt returns the DetectedAt field if non-nil, zero value otherwise.

### GetDetectedAtOk

`func (o *ArchiveTagEntry) GetDetectedAtOk() (*time.Time, bool)`

GetDetectedAtOk returns a tuple with the DetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedAt

`func (o *ArchiveTagEntry) SetDetectedAt(v time.Time)`

SetDetectedAt sets DetectedAt field to given value.

### HasDetectedAt

`func (o *ArchiveTagEntry) HasDetectedAt() bool`

HasDetectedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArchiveTagEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArchiveTagEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArchiveTagEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArchiveTagEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ArchiveTagEntry) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ArchiveTagEntry) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ArchiveTagEntry) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ArchiveTagEntry) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


