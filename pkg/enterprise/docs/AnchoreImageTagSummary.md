# AnchoreImageTagSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDigest** | Pointer to **string** |  | [optional] 
**ParentDigest** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**AnalysisStatus** | Pointer to **string** |  | [optional] 
**FullTag** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**AnalyzedAt** | Pointer to **int32** |  | [optional] 
**TagDetectedAt** | Pointer to **int32** |  | [optional] 
**ImageStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewAnchoreImageTagSummary

`func NewAnchoreImageTagSummary() *AnchoreImageTagSummary`

NewAnchoreImageTagSummary instantiates a new AnchoreImageTagSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnchoreImageTagSummaryWithDefaults

`func NewAnchoreImageTagSummaryWithDefaults() *AnchoreImageTagSummary`

NewAnchoreImageTagSummaryWithDefaults instantiates a new AnchoreImageTagSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDigest

`func (o *AnchoreImageTagSummary) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *AnchoreImageTagSummary) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *AnchoreImageTagSummary) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *AnchoreImageTagSummary) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetParentDigest

`func (o *AnchoreImageTagSummary) GetParentDigest() string`

GetParentDigest returns the ParentDigest field if non-nil, zero value otherwise.

### GetParentDigestOk

`func (o *AnchoreImageTagSummary) GetParentDigestOk() (*string, bool)`

GetParentDigestOk returns a tuple with the ParentDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDigest

`func (o *AnchoreImageTagSummary) SetParentDigest(v string)`

SetParentDigest sets ParentDigest field to given value.

### HasParentDigest

`func (o *AnchoreImageTagSummary) HasParentDigest() bool`

HasParentDigest returns a boolean if a field has been set.

### GetImageId

`func (o *AnchoreImageTagSummary) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AnchoreImageTagSummary) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AnchoreImageTagSummary) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *AnchoreImageTagSummary) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetAnalysisStatus

`func (o *AnchoreImageTagSummary) GetAnalysisStatus() string`

GetAnalysisStatus returns the AnalysisStatus field if non-nil, zero value otherwise.

### GetAnalysisStatusOk

`func (o *AnchoreImageTagSummary) GetAnalysisStatusOk() (*string, bool)`

GetAnalysisStatusOk returns a tuple with the AnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatus

`func (o *AnchoreImageTagSummary) SetAnalysisStatus(v string)`

SetAnalysisStatus sets AnalysisStatus field to given value.

### HasAnalysisStatus

`func (o *AnchoreImageTagSummary) HasAnalysisStatus() bool`

HasAnalysisStatus returns a boolean if a field has been set.

### GetFullTag

`func (o *AnchoreImageTagSummary) GetFullTag() string`

GetFullTag returns the FullTag field if non-nil, zero value otherwise.

### GetFullTagOk

`func (o *AnchoreImageTagSummary) GetFullTagOk() (*string, bool)`

GetFullTagOk returns a tuple with the FullTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTag

`func (o *AnchoreImageTagSummary) SetFullTag(v string)`

SetFullTag sets FullTag field to given value.

### HasFullTag

`func (o *AnchoreImageTagSummary) HasFullTag() bool`

HasFullTag returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AnchoreImageTagSummary) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnchoreImageTagSummary) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnchoreImageTagSummary) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AnchoreImageTagSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAnalyzedAt

`func (o *AnchoreImageTagSummary) GetAnalyzedAt() int32`

GetAnalyzedAt returns the AnalyzedAt field if non-nil, zero value otherwise.

### GetAnalyzedAtOk

`func (o *AnchoreImageTagSummary) GetAnalyzedAtOk() (*int32, bool)`

GetAnalyzedAtOk returns a tuple with the AnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzedAt

`func (o *AnchoreImageTagSummary) SetAnalyzedAt(v int32)`

SetAnalyzedAt sets AnalyzedAt field to given value.

### HasAnalyzedAt

`func (o *AnchoreImageTagSummary) HasAnalyzedAt() bool`

HasAnalyzedAt returns a boolean if a field has been set.

### GetTagDetectedAt

`func (o *AnchoreImageTagSummary) GetTagDetectedAt() int32`

GetTagDetectedAt returns the TagDetectedAt field if non-nil, zero value otherwise.

### GetTagDetectedAtOk

`func (o *AnchoreImageTagSummary) GetTagDetectedAtOk() (*int32, bool)`

GetTagDetectedAtOk returns a tuple with the TagDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagDetectedAt

`func (o *AnchoreImageTagSummary) SetTagDetectedAt(v int32)`

SetTagDetectedAt sets TagDetectedAt field to given value.

### HasTagDetectedAt

`func (o *AnchoreImageTagSummary) HasTagDetectedAt() bool`

HasTagDetectedAt returns a boolean if a field has been set.

### GetImageStatus

`func (o *AnchoreImageTagSummary) GetImageStatus() string`

GetImageStatus returns the ImageStatus field if non-nil, zero value otherwise.

### GetImageStatusOk

`func (o *AnchoreImageTagSummary) GetImageStatusOk() (*string, bool)`

GetImageStatusOk returns a tuple with the ImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStatus

`func (o *AnchoreImageTagSummary) SetImageStatus(v string)`

SetImageStatus sets ImageStatus field to given value.

### HasImageStatus

`func (o *AnchoreImageTagSummary) HasImageStatus() bool`

HasImageStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


