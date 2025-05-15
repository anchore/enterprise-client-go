# ContentJavaPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDigest** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Content** | Pointer to [**[]JavaPackageContent**](JavaPackageContent.md) |  | [optional] 

## Methods

### NewContentJavaPackageResponse

`func NewContentJavaPackageResponse() *ContentJavaPackageResponse`

NewContentJavaPackageResponse instantiates a new ContentJavaPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentJavaPackageResponseWithDefaults

`func NewContentJavaPackageResponseWithDefaults() *ContentJavaPackageResponse`

NewContentJavaPackageResponseWithDefaults instantiates a new ContentJavaPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDigest

`func (o *ContentJavaPackageResponse) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ContentJavaPackageResponse) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ContentJavaPackageResponse) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ContentJavaPackageResponse) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetContentType

`func (o *ContentJavaPackageResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ContentJavaPackageResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ContentJavaPackageResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ContentJavaPackageResponse) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContent

`func (o *ContentJavaPackageResponse) GetContent() []JavaPackageContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContentJavaPackageResponse) GetContentOk() (*[]JavaPackageContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContentJavaPackageResponse) SetContent(v []JavaPackageContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *ContentJavaPackageResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


