# ContentSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaType** | **string** | The media type of this result | 
**Metadata** | [**ContentSearchMetadata**](ContentSearchMetadata.md) |  | 
**Findings** | [**[]ContentSearchFindingsInner**](ContentSearchFindingsInner.md) |  | 

## Methods

### NewContentSearch

`func NewContentSearch(mediaType string, metadata ContentSearchMetadata, findings []ContentSearchFindingsInner, ) *ContentSearch`

NewContentSearch instantiates a new ContentSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSearchWithDefaults

`func NewContentSearchWithDefaults() *ContentSearch`

NewContentSearchWithDefaults instantiates a new ContentSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaType

`func (o *ContentSearch) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *ContentSearch) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *ContentSearch) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetMetadata

`func (o *ContentSearch) GetMetadata() ContentSearchMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ContentSearch) GetMetadataOk() (*ContentSearchMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ContentSearch) SetMetadata(v ContentSearchMetadata)`

SetMetadata sets Metadata field to given value.


### GetFindings

`func (o *ContentSearch) GetFindings() []ContentSearchFindingsInner`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *ContentSearch) GetFindingsOk() (*[]ContentSearchFindingsInner, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *ContentSearch) SetFindings(v []ContentSearchFindingsInner)`

SetFindings sets Findings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


