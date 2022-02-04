# SourceManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**RepositoryName** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**AnalysisStatus** | Pointer to **string** |  | [optional] 
**SourceStatus** | Pointer to **string** |  | [optional] 
**Artifacts** | Pointer to [**[]SourceManifestArtifacts**](SourceManifestArtifacts.md) | Digest of content to use in the final import | [optional] 
**MetadataRecords** | Pointer to [**[]SourceManifestMetadataRecords**](SourceManifestMetadataRecords.md) | Array of metadata available | [optional] 

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

### GetArtifacts

`func (o *SourceManifest) GetArtifacts() []SourceManifestArtifacts`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *SourceManifest) GetArtifactsOk() (*[]SourceManifestArtifacts, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *SourceManifest) SetArtifacts(v []SourceManifestArtifacts)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *SourceManifest) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetMetadataRecords

`func (o *SourceManifest) GetMetadataRecords() []SourceManifestMetadataRecords`

GetMetadataRecords returns the MetadataRecords field if non-nil, zero value otherwise.

### GetMetadataRecordsOk

`func (o *SourceManifest) GetMetadataRecordsOk() (*[]SourceManifestMetadataRecords, bool)`

GetMetadataRecordsOk returns a tuple with the MetadataRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataRecords

`func (o *SourceManifest) SetMetadataRecords(v []SourceManifestMetadataRecords)`

SetMetadataRecords sets MetadataRecords field to given value.

### HasMetadataRecords

`func (o *SourceManifest) HasMetadataRecords() bool`

HasMetadataRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


