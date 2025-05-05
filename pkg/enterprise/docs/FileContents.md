# FileContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaType** | **string** | The media type of this result | 
**Metadata** | [**ContentSearchMetadata**](ContentSearchMetadata.md) |  | 
**Findings** | [**[]FileContentsFindingsInner**](FileContentsFindingsInner.md) |  | 

## Methods

### NewFileContents

`func NewFileContents(mediaType string, metadata ContentSearchMetadata, findings []FileContentsFindingsInner, ) *FileContents`

NewFileContents instantiates a new FileContents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileContentsWithDefaults

`func NewFileContentsWithDefaults() *FileContents`

NewFileContentsWithDefaults instantiates a new FileContents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaType

`func (o *FileContents) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *FileContents) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *FileContents) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetMetadata

`func (o *FileContents) GetMetadata() ContentSearchMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FileContents) GetMetadataOk() (*ContentSearchMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FileContents) SetMetadata(v ContentSearchMetadata)`

SetMetadata sets Metadata field to given value.


### GetFindings

`func (o *FileContents) GetFindings() []FileContentsFindingsInner`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *FileContents) GetFindingsOk() (*[]FileContentsFindingsInner, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *FileContents) SetFindings(v []FileContentsFindingsInner)`

SetFindings sets Findings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


