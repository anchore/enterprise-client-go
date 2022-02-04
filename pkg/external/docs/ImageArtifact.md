# ImageArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account id | [optional] 
**ImageDigest** | Pointer to **string** | The digest of the image | [optional] 
**Distro** | Pointer to **string** | The distro of the image | [optional] 
**DistroVersion** | Pointer to **string** | The distro version of the image | [optional] 
**AnalysisStatus** | Pointer to **string** | the analysis status of image | [optional] 
**ImageStatus** | Pointer to **string** | The status of the image | [optional] 
**AnalyzedAt** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the image was analyzed | [optional] 
**CreatedAt** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the image was created | [optional] 
**LastUpdated** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the image was last updated | [optional] 

## Methods

### NewImageArtifact

`func NewImageArtifact() *ImageArtifact`

NewImageArtifact instantiates a new ImageArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageArtifactWithDefaults

`func NewImageArtifactWithDefaults() *ImageArtifact`

NewImageArtifactWithDefaults instantiates a new ImageArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ImageArtifact) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ImageArtifact) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ImageArtifact) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ImageArtifact) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetImageDigest

`func (o *ImageArtifact) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ImageArtifact) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ImageArtifact) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ImageArtifact) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetDistro

`func (o *ImageArtifact) GetDistro() string`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *ImageArtifact) GetDistroOk() (*string, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *ImageArtifact) SetDistro(v string)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *ImageArtifact) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### GetDistroVersion

`func (o *ImageArtifact) GetDistroVersion() string`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *ImageArtifact) GetDistroVersionOk() (*string, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *ImageArtifact) SetDistroVersion(v string)`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *ImageArtifact) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### GetAnalysisStatus

`func (o *ImageArtifact) GetAnalysisStatus() string`

GetAnalysisStatus returns the AnalysisStatus field if non-nil, zero value otherwise.

### GetAnalysisStatusOk

`func (o *ImageArtifact) GetAnalysisStatusOk() (*string, bool)`

GetAnalysisStatusOk returns a tuple with the AnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatus

`func (o *ImageArtifact) SetAnalysisStatus(v string)`

SetAnalysisStatus sets AnalysisStatus field to given value.

### HasAnalysisStatus

`func (o *ImageArtifact) HasAnalysisStatus() bool`

HasAnalysisStatus returns a boolean if a field has been set.

### GetImageStatus

`func (o *ImageArtifact) GetImageStatus() string`

GetImageStatus returns the ImageStatus field if non-nil, zero value otherwise.

### GetImageStatusOk

`func (o *ImageArtifact) GetImageStatusOk() (*string, bool)`

GetImageStatusOk returns a tuple with the ImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStatus

`func (o *ImageArtifact) SetImageStatus(v string)`

SetImageStatus sets ImageStatus field to given value.

### HasImageStatus

`func (o *ImageArtifact) HasImageStatus() bool`

HasImageStatus returns a boolean if a field has been set.

### GetAnalyzedAt

`func (o *ImageArtifact) GetAnalyzedAt() time.Time`

GetAnalyzedAt returns the AnalyzedAt field if non-nil, zero value otherwise.

### GetAnalyzedAtOk

`func (o *ImageArtifact) GetAnalyzedAtOk() (*time.Time, bool)`

GetAnalyzedAtOk returns a tuple with the AnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzedAt

`func (o *ImageArtifact) SetAnalyzedAt(v time.Time)`

SetAnalyzedAt sets AnalyzedAt field to given value.

### HasAnalyzedAt

`func (o *ImageArtifact) HasAnalyzedAt() bool`

HasAnalyzedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ImageArtifact) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImageArtifact) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImageArtifact) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ImageArtifact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ImageArtifact) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ImageArtifact) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ImageArtifact) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ImageArtifact) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


