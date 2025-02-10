# AnchoreImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageContent** | Pointer to [**ImageContent**](ImageContent.md) |  | [optional] 
**ImageDetail** | Pointer to [**[]ImageDetail**](ImageDetail.md) | Details specific to an image reference and type such as tag and image source | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ImageDigest** | Pointer to **string** |  | [optional] 
**ParentDigest** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **interface{}** |  | [optional] 
**ImageStatus** | Pointer to **string** | State of the image | [optional] 
**AnalysisStatus** | Pointer to **string** | A state value for the current status of the analysis progress of the image | [optional] 
**RecordVersion** | Pointer to **string** | The version of the record, used for internal schema updates and data migrations. | [optional] 
**AnalysisStatusDetail** | Pointer to [**[]AnalysisStatusDetail**](AnalysisStatusDetail.md) |  | [optional] 

## Methods

### NewAnchoreImage

`func NewAnchoreImage() *AnchoreImage`

NewAnchoreImage instantiates a new AnchoreImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnchoreImageWithDefaults

`func NewAnchoreImageWithDefaults() *AnchoreImage`

NewAnchoreImageWithDefaults instantiates a new AnchoreImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageContent

`func (o *AnchoreImage) GetImageContent() ImageContent`

GetImageContent returns the ImageContent field if non-nil, zero value otherwise.

### GetImageContentOk

`func (o *AnchoreImage) GetImageContentOk() (*ImageContent, bool)`

GetImageContentOk returns a tuple with the ImageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageContent

`func (o *AnchoreImage) SetImageContent(v ImageContent)`

SetImageContent sets ImageContent field to given value.

### HasImageContent

`func (o *AnchoreImage) HasImageContent() bool`

HasImageContent returns a boolean if a field has been set.

### GetImageDetail

`func (o *AnchoreImage) GetImageDetail() []ImageDetail`

GetImageDetail returns the ImageDetail field if non-nil, zero value otherwise.

### GetImageDetailOk

`func (o *AnchoreImage) GetImageDetailOk() (*[]ImageDetail, bool)`

GetImageDetailOk returns a tuple with the ImageDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDetail

`func (o *AnchoreImage) SetImageDetail(v []ImageDetail)`

SetImageDetail sets ImageDetail field to given value.

### HasImageDetail

`func (o *AnchoreImage) HasImageDetail() bool`

HasImageDetail returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AnchoreImage) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AnchoreImage) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AnchoreImage) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AnchoreImage) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AnchoreImage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnchoreImage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnchoreImage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AnchoreImage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetImageDigest

`func (o *AnchoreImage) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *AnchoreImage) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *AnchoreImage) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *AnchoreImage) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetParentDigest

`func (o *AnchoreImage) GetParentDigest() string`

GetParentDigest returns the ParentDigest field if non-nil, zero value otherwise.

### GetParentDigestOk

`func (o *AnchoreImage) GetParentDigestOk() (*string, bool)`

GetParentDigestOk returns a tuple with the ParentDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDigest

`func (o *AnchoreImage) SetParentDigest(v string)`

SetParentDigest sets ParentDigest field to given value.

### HasParentDigest

`func (o *AnchoreImage) HasParentDigest() bool`

HasParentDigest returns a boolean if a field has been set.

### GetAccountName

`func (o *AnchoreImage) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AnchoreImage) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AnchoreImage) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AnchoreImage) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAnnotations

`func (o *AnchoreImage) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AnchoreImage) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AnchoreImage) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AnchoreImage) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetImageStatus

`func (o *AnchoreImage) GetImageStatus() string`

GetImageStatus returns the ImageStatus field if non-nil, zero value otherwise.

### GetImageStatusOk

`func (o *AnchoreImage) GetImageStatusOk() (*string, bool)`

GetImageStatusOk returns a tuple with the ImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStatus

`func (o *AnchoreImage) SetImageStatus(v string)`

SetImageStatus sets ImageStatus field to given value.

### HasImageStatus

`func (o *AnchoreImage) HasImageStatus() bool`

HasImageStatus returns a boolean if a field has been set.

### GetAnalysisStatus

`func (o *AnchoreImage) GetAnalysisStatus() string`

GetAnalysisStatus returns the AnalysisStatus field if non-nil, zero value otherwise.

### GetAnalysisStatusOk

`func (o *AnchoreImage) GetAnalysisStatusOk() (*string, bool)`

GetAnalysisStatusOk returns a tuple with the AnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatus

`func (o *AnchoreImage) SetAnalysisStatus(v string)`

SetAnalysisStatus sets AnalysisStatus field to given value.

### HasAnalysisStatus

`func (o *AnchoreImage) HasAnalysisStatus() bool`

HasAnalysisStatus returns a boolean if a field has been set.

### GetRecordVersion

`func (o *AnchoreImage) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *AnchoreImage) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *AnchoreImage) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *AnchoreImage) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetAnalysisStatusDetail

`func (o *AnchoreImage) GetAnalysisStatusDetail() []AnalysisStatusDetail`

GetAnalysisStatusDetail returns the AnalysisStatusDetail field if non-nil, zero value otherwise.

### GetAnalysisStatusDetailOk

`func (o *AnchoreImage) GetAnalysisStatusDetailOk() (*[]AnalysisStatusDetail, bool)`

GetAnalysisStatusDetailOk returns a tuple with the AnalysisStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatusDetail

`func (o *AnchoreImage) SetAnalysisStatusDetail(v []AnalysisStatusDetail)`

SetAnalysisStatusDetail sets AnalysisStatusDetail field to given value.

### HasAnalysisStatusDetail

`func (o *AnchoreImage) HasAnalysisStatusDetail() bool`

HasAnalysisStatusDetail returns a boolean if a field has been set.

### SetAnalysisStatusDetailNil

`func (o *AnchoreImage) SetAnalysisStatusDetailNil(b bool)`

 SetAnalysisStatusDetailNil sets the value for AnalysisStatusDetail to be an explicit nil

### UnsetAnalysisStatusDetail
`func (o *AnchoreImage) UnsetAnalysisStatusDetail()`

UnsetAnalysisStatusDetail ensures that no value is present for AnalysisStatusDetail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


