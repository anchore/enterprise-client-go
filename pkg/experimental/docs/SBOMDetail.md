# SBOMDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**RevisionUuid** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**FileFormat** | Pointer to **string** |  | [optional] 
**FileSchema** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **float32** |  | [optional] 
**FileHash** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]SBOMDetailGroupsInner**](SBOMDetailGroupsInner.md) |  | [optional] 
**RevisionHistory** | Pointer to [**[]SBOMDetailRevisionHistoryInner**](SBOMDetailRevisionHistoryInner.md) |  | [optional] 
**DocumentInsights** | Pointer to [**SBOMDetailDocumentInsights**](SBOMDetailDocumentInsights.md) |  | [optional] 

## Methods

### NewSBOMDetail

`func NewSBOMDetail() *SBOMDetail`

NewSBOMDetail instantiates a new SBOMDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSBOMDetailWithDefaults

`func NewSBOMDetailWithDefaults() *SBOMDetail`

NewSBOMDetailWithDefaults instantiates a new SBOMDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *SBOMDetail) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SBOMDetail) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SBOMDetail) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *SBOMDetail) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAnnotations

`func (o *SBOMDetail) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SBOMDetail) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SBOMDetail) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SBOMDetail) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetName

`func (o *SBOMDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SBOMDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SBOMDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SBOMDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SBOMDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SBOMDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SBOMDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SBOMDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *SBOMDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SBOMDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SBOMDetail) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SBOMDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUuid

`func (o *SBOMDetail) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SBOMDetail) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SBOMDetail) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SBOMDetail) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetRevisionUuid

`func (o *SBOMDetail) GetRevisionUuid() string`

GetRevisionUuid returns the RevisionUuid field if non-nil, zero value otherwise.

### GetRevisionUuidOk

`func (o *SBOMDetail) GetRevisionUuidOk() (*string, bool)`

GetRevisionUuidOk returns a tuple with the RevisionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionUuid

`func (o *SBOMDetail) SetRevisionUuid(v string)`

SetRevisionUuid sets RevisionUuid field to given value.

### HasRevisionUuid

`func (o *SBOMDetail) HasRevisionUuid() bool`

HasRevisionUuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SBOMDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SBOMDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SBOMDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SBOMDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFileFormat

`func (o *SBOMDetail) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *SBOMDetail) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *SBOMDetail) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *SBOMDetail) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### GetFileSchema

`func (o *SBOMDetail) GetFileSchema() string`

GetFileSchema returns the FileSchema field if non-nil, zero value otherwise.

### GetFileSchemaOk

`func (o *SBOMDetail) GetFileSchemaOk() (*string, bool)`

GetFileSchemaOk returns a tuple with the FileSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSchema

`func (o *SBOMDetail) SetFileSchema(v string)`

SetFileSchema sets FileSchema field to given value.

### HasFileSchema

`func (o *SBOMDetail) HasFileSchema() bool`

HasFileSchema returns a boolean if a field has been set.

### GetFileName

`func (o *SBOMDetail) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SBOMDetail) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SBOMDetail) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SBOMDetail) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *SBOMDetail) GetFileSize() float32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *SBOMDetail) GetFileSizeOk() (*float32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *SBOMDetail) SetFileSize(v float32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *SBOMDetail) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileHash

`func (o *SBOMDetail) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *SBOMDetail) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *SBOMDetail) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.

### HasFileHash

`func (o *SBOMDetail) HasFileHash() bool`

HasFileHash returns a boolean if a field has been set.

### GetGroups

`func (o *SBOMDetail) GetGroups() []SBOMDetailGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SBOMDetail) GetGroupsOk() (*[]SBOMDetailGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SBOMDetail) SetGroups(v []SBOMDetailGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SBOMDetail) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetRevisionHistory

`func (o *SBOMDetail) GetRevisionHistory() []SBOMDetailRevisionHistoryInner`

GetRevisionHistory returns the RevisionHistory field if non-nil, zero value otherwise.

### GetRevisionHistoryOk

`func (o *SBOMDetail) GetRevisionHistoryOk() (*[]SBOMDetailRevisionHistoryInner, bool)`

GetRevisionHistoryOk returns a tuple with the RevisionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionHistory

`func (o *SBOMDetail) SetRevisionHistory(v []SBOMDetailRevisionHistoryInner)`

SetRevisionHistory sets RevisionHistory field to given value.

### HasRevisionHistory

`func (o *SBOMDetail) HasRevisionHistory() bool`

HasRevisionHistory returns a boolean if a field has been set.

### GetDocumentInsights

`func (o *SBOMDetail) GetDocumentInsights() SBOMDetailDocumentInsights`

GetDocumentInsights returns the DocumentInsights field if non-nil, zero value otherwise.

### GetDocumentInsightsOk

`func (o *SBOMDetail) GetDocumentInsightsOk() (*SBOMDetailDocumentInsights, bool)`

GetDocumentInsightsOk returns a tuple with the DocumentInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentInsights

`func (o *SBOMDetail) SetDocumentInsights(v SBOMDetailDocumentInsights)`

SetDocumentInsights sets DocumentInsights field to given value.

### HasDocumentInsights

`func (o *SBOMDetail) HasDocumentInsights() bool`

HasDocumentInsights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


