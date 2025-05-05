# ContentSearchMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyzedAt** | **time.Time** | The time the result was analyzed | 
**AnchorectlVersion** | **string** | The version of AnchoreCTL used to generate the result | 
**Platform** | **string** | The platform of AnchoreCTL used to generate the result | 

## Methods

### NewContentSearchMetadata

`func NewContentSearchMetadata(analyzedAt time.Time, anchorectlVersion string, platform string, ) *ContentSearchMetadata`

NewContentSearchMetadata instantiates a new ContentSearchMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSearchMetadataWithDefaults

`func NewContentSearchMetadataWithDefaults() *ContentSearchMetadata`

NewContentSearchMetadataWithDefaults instantiates a new ContentSearchMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzedAt

`func (o *ContentSearchMetadata) GetAnalyzedAt() time.Time`

GetAnalyzedAt returns the AnalyzedAt field if non-nil, zero value otherwise.

### GetAnalyzedAtOk

`func (o *ContentSearchMetadata) GetAnalyzedAtOk() (*time.Time, bool)`

GetAnalyzedAtOk returns a tuple with the AnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzedAt

`func (o *ContentSearchMetadata) SetAnalyzedAt(v time.Time)`

SetAnalyzedAt sets AnalyzedAt field to given value.


### GetAnchorectlVersion

`func (o *ContentSearchMetadata) GetAnchorectlVersion() string`

GetAnchorectlVersion returns the AnchorectlVersion field if non-nil, zero value otherwise.

### GetAnchorectlVersionOk

`func (o *ContentSearchMetadata) GetAnchorectlVersionOk() (*string, bool)`

GetAnchorectlVersionOk returns a tuple with the AnchorectlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorectlVersion

`func (o *ContentSearchMetadata) SetAnchorectlVersion(v string)`

SetAnchorectlVersion sets AnchorectlVersion field to given value.


### GetPlatform

`func (o *ContentSearchMetadata) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ContentSearchMetadata) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ContentSearchMetadata) SetPlatform(v string)`

SetPlatform sets Platform field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


