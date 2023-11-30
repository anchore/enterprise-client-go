# SourceManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**VcsType** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**RepositoryName** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**AnalysisStatus** | Pointer to **string** |  | [optional] 
**SourceStatus** | Pointer to **string** |  | [optional] 
**MetadataRecords** | Pointer to [**[]SourceManifestMetadataRecord**](SourceManifestMetadataRecord.md) | Array of metadata available | [optional] 

## Methods

### NewSourceManifest

`func NewSourceManifest() *SourceManifest`

NewSourceManifest instantiates a new SourceManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceManifestWithDefaults

`func NewSourceManifestWithDefaults() *SourceManifest`

NewSourceManifestWithDefaults instantiates a new SourceManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SourceManifest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SourceManifest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SourceManifest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SourceManifest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccountName

`func (o *SourceManifest) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SourceManifest) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SourceManifest) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *SourceManifest) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetVcsType

`func (o *SourceManifest) GetVcsType() string`

GetVcsType returns the VcsType field if non-nil, zero value otherwise.

### GetVcsTypeOk

`func (o *SourceManifest) GetVcsTypeOk() (*string, bool)`

GetVcsTypeOk returns a tuple with the VcsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsType

`func (o *SourceManifest) SetVcsType(v string)`

SetVcsType sets VcsType field to given value.

### HasVcsType

`func (o *SourceManifest) HasVcsType() bool`

HasVcsType returns a boolean if a field has been set.

### GetHost

`func (o *SourceManifest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SourceManifest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SourceManifest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SourceManifest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetRepositoryName

`func (o *SourceManifest) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *SourceManifest) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *SourceManifest) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.

### HasRepositoryName

`func (o *SourceManifest) HasRepositoryName() bool`

HasRepositoryName returns a boolean if a field has been set.

### GetRevision

`func (o *SourceManifest) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SourceManifest) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SourceManifest) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *SourceManifest) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SourceManifest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SourceManifest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SourceManifest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SourceManifest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SourceManifest) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SourceManifest) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SourceManifest) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SourceManifest) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAnalysisStatus

`func (o *SourceManifest) GetAnalysisStatus() string`

GetAnalysisStatus returns the AnalysisStatus field if non-nil, zero value otherwise.

### GetAnalysisStatusOk

`func (o *SourceManifest) GetAnalysisStatusOk() (*string, bool)`

GetAnalysisStatusOk returns a tuple with the AnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatus

`func (o *SourceManifest) SetAnalysisStatus(v string)`

SetAnalysisStatus sets AnalysisStatus field to given value.

### HasAnalysisStatus

`func (o *SourceManifest) HasAnalysisStatus() bool`

HasAnalysisStatus returns a boolean if a field has been set.

### GetSourceStatus

`func (o *SourceManifest) GetSourceStatus() string`

GetSourceStatus returns the SourceStatus field if non-nil, zero value otherwise.

### GetSourceStatusOk

`func (o *SourceManifest) GetSourceStatusOk() (*string, bool)`

GetSourceStatusOk returns a tuple with the SourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStatus

`func (o *SourceManifest) SetSourceStatus(v string)`

SetSourceStatus sets SourceStatus field to given value.

### HasSourceStatus

`func (o *SourceManifest) HasSourceStatus() bool`

HasSourceStatus returns a boolean if a field has been set.

### GetMetadataRecords

`func (o *SourceManifest) GetMetadataRecords() []SourceManifestMetadataRecord`

GetMetadataRecords returns the MetadataRecords field if non-nil, zero value otherwise.

### GetMetadataRecordsOk

`func (o *SourceManifest) GetMetadataRecordsOk() (*[]SourceManifestMetadataRecord, bool)`

GetMetadataRecordsOk returns a tuple with the MetadataRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataRecords

`func (o *SourceManifest) SetMetadataRecords(v []SourceManifestMetadataRecord)`

SetMetadataRecords sets MetadataRecords field to given value.

### HasMetadataRecords

`func (o *SourceManifest) HasMetadataRecords() bool`

HasMetadataRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


