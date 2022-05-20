# ArchivedAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDigest** | Pointer to **string** | The image digest (digest of the manifest describing the image, per docker spec) | [optional] 
**ParentDigest** | Pointer to **string** | The digest of a parent manifest (for manifest-list images) | [optional] 
**Annotations** | Pointer to **interface{}** | User provided annotations as key-value pairs | [optional] 
**Status** | Pointer to **string** | The archival status | [optional] 
**ImageDetail** | Pointer to [**[]TagEntry**](TagEntry.md) | List of tags associated with the image digest | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**AnalyzedAt** | Pointer to **time.Time** |  | [optional] 
**ArchiveSizeBytes** | Pointer to **int32** | The size, in bytes, of the analysis archive file | [optional] 

## Methods

### NewArchivedAnalysis

`func NewArchivedAnalysis() *ArchivedAnalysis`

NewArchivedAnalysis instantiates a new ArchivedAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivedAnalysisWithDefaults

`func NewArchivedAnalysisWithDefaults() *ArchivedAnalysis`

NewArchivedAnalysisWithDefaults instantiates a new ArchivedAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDigest

`func (o *ArchivedAnalysis) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ArchivedAnalysis) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ArchivedAnalysis) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ArchivedAnalysis) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetParentDigest

`func (o *ArchivedAnalysis) GetParentDigest() string`

GetParentDigest returns the ParentDigest field if non-nil, zero value otherwise.

### GetParentDigestOk

`func (o *ArchivedAnalysis) GetParentDigestOk() (*string, bool)`

GetParentDigestOk returns a tuple with the ParentDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDigest

`func (o *ArchivedAnalysis) SetParentDigest(v string)`

SetParentDigest sets ParentDigest field to given value.

### HasParentDigest

`func (o *ArchivedAnalysis) HasParentDigest() bool`

HasParentDigest returns a boolean if a field has been set.

### GetAnnotations

`func (o *ArchivedAnalysis) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ArchivedAnalysis) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ArchivedAnalysis) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ArchivedAnalysis) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetStatus

`func (o *ArchivedAnalysis) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArchivedAnalysis) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArchivedAnalysis) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArchivedAnalysis) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImageDetail

`func (o *ArchivedAnalysis) GetImageDetail() []TagEntry`

GetImageDetail returns the ImageDetail field if non-nil, zero value otherwise.

### GetImageDetailOk

`func (o *ArchivedAnalysis) GetImageDetailOk() (*[]TagEntry, bool)`

GetImageDetailOk returns a tuple with the ImageDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDetail

`func (o *ArchivedAnalysis) SetImageDetail(v []TagEntry)`

SetImageDetail sets ImageDetail field to given value.

### HasImageDetail

`func (o *ArchivedAnalysis) HasImageDetail() bool`

HasImageDetail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArchivedAnalysis) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArchivedAnalysis) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArchivedAnalysis) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArchivedAnalysis) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ArchivedAnalysis) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ArchivedAnalysis) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ArchivedAnalysis) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ArchivedAnalysis) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAnalyzedAt

`func (o *ArchivedAnalysis) GetAnalyzedAt() time.Time`

GetAnalyzedAt returns the AnalyzedAt field if non-nil, zero value otherwise.

### GetAnalyzedAtOk

`func (o *ArchivedAnalysis) GetAnalyzedAtOk() (*time.Time, bool)`

GetAnalyzedAtOk returns a tuple with the AnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzedAt

`func (o *ArchivedAnalysis) SetAnalyzedAt(v time.Time)`

SetAnalyzedAt sets AnalyzedAt field to given value.

### HasAnalyzedAt

`func (o *ArchivedAnalysis) HasAnalyzedAt() bool`

HasAnalyzedAt returns a boolean if a field has been set.

### GetArchiveSizeBytes

`func (o *ArchivedAnalysis) GetArchiveSizeBytes() int32`

GetArchiveSizeBytes returns the ArchiveSizeBytes field if non-nil, zero value otherwise.

### GetArchiveSizeBytesOk

`func (o *ArchivedAnalysis) GetArchiveSizeBytesOk() (*int32, bool)`

GetArchiveSizeBytesOk returns a tuple with the ArchiveSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveSizeBytes

`func (o *ArchivedAnalysis) SetArchiveSizeBytes(v int32)`

SetArchiveSizeBytes sets ArchiveSizeBytes field to given value.

### HasArchiveSizeBytes

`func (o *ArchivedAnalysis) HasArchiveSizeBytes() bool`

HasArchiveSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


