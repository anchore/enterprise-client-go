# ContentFilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDigest** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Content** | Pointer to [**[]ContentFilesResponseContentInner**](ContentFilesResponseContentInner.md) |  | [optional] 

## Methods

### NewContentFilesResponse

`func NewContentFilesResponse() *ContentFilesResponse`

NewContentFilesResponse instantiates a new ContentFilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentFilesResponseWithDefaults

`func NewContentFilesResponseWithDefaults() *ContentFilesResponse`

NewContentFilesResponseWithDefaults instantiates a new ContentFilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDigest

`func (o *ContentFilesResponse) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ContentFilesResponse) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ContentFilesResponse) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ContentFilesResponse) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetContentType

`func (o *ContentFilesResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ContentFilesResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ContentFilesResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ContentFilesResponse) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContent

`func (o *ContentFilesResponse) GetContent() []ContentFilesResponseContentInner`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContentFilesResponse) GetContentOk() (*[]ContentFilesResponseContentInner, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContentFilesResponse) SetContent(v []ContentFilesResponseContentInner)`

SetContent sets Content field to given value.

### HasContent

`func (o *ContentFilesResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


