# FileContentsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyzedAt** | Pointer to **time.Time** | The time the result was analyzed | [optional] 
**AnchorectlVersion** | **string** | The version of AnchoreCTL used to generate the result | 
**Platform** | Pointer to **string** | The platform of AnchoreCTL used to generate the result | [optional] 
**AnchorectlPlatform** | **string** | The platform of AnchoreCTL used to generate the result | 

## Methods

### NewFileContentsMetadata

`func NewFileContentsMetadata(anchorectlVersion string, anchorectlPlatform string, ) *FileContentsMetadata`

NewFileContentsMetadata instantiates a new FileContentsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileContentsMetadataWithDefaults

`func NewFileContentsMetadataWithDefaults() *FileContentsMetadata`

NewFileContentsMetadataWithDefaults instantiates a new FileContentsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzedAt

`func (o *FileContentsMetadata) GetAnalyzedAt() time.Time`

GetAnalyzedAt returns the AnalyzedAt field if non-nil, zero value otherwise.

### GetAnalyzedAtOk

`func (o *FileContentsMetadata) GetAnalyzedAtOk() (*time.Time, bool)`

GetAnalyzedAtOk returns a tuple with the AnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzedAt

`func (o *FileContentsMetadata) SetAnalyzedAt(v time.Time)`

SetAnalyzedAt sets AnalyzedAt field to given value.

### HasAnalyzedAt

`func (o *FileContentsMetadata) HasAnalyzedAt() bool`

HasAnalyzedAt returns a boolean if a field has been set.

### GetAnchorectlVersion

`func (o *FileContentsMetadata) GetAnchorectlVersion() string`

GetAnchorectlVersion returns the AnchorectlVersion field if non-nil, zero value otherwise.

### GetAnchorectlVersionOk

`func (o *FileContentsMetadata) GetAnchorectlVersionOk() (*string, bool)`

GetAnchorectlVersionOk returns a tuple with the AnchorectlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorectlVersion

`func (o *FileContentsMetadata) SetAnchorectlVersion(v string)`

SetAnchorectlVersion sets AnchorectlVersion field to given value.


### GetPlatform

`func (o *FileContentsMetadata) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *FileContentsMetadata) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *FileContentsMetadata) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *FileContentsMetadata) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAnchorectlPlatform

`func (o *FileContentsMetadata) GetAnchorectlPlatform() string`

GetAnchorectlPlatform returns the AnchorectlPlatform field if non-nil, zero value otherwise.

### GetAnchorectlPlatformOk

`func (o *FileContentsMetadata) GetAnchorectlPlatformOk() (*string, bool)`

GetAnchorectlPlatformOk returns a tuple with the AnchorectlPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorectlPlatform

`func (o *FileContentsMetadata) SetAnchorectlPlatform(v string)`

SetAnchorectlPlatform sets AnchorectlPlatform field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


