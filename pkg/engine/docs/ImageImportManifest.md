# ImageImportManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**ImportContentDigests**](ImportContentDigests.md) |  | 
**Tags** | **[]string** |  | 
**Digest** | **string** |  | 
**ParentDigest** | Pointer to **string** | The digest of the image&#39;s manifest-list parent if it was accessed from a multi-arch tag where the tag pointed to a manifest-list. This allows preservation of that relationship in the data | [optional] 
**LocalImageId** | Pointer to **string** | An \&quot;image_id\&quot; as used by Docker if available | [optional] 
**OperationUuid** | **string** |  | 

## Methods

### NewImageImportManifest

`func NewImageImportManifest(contents ImportContentDigests, tags []string, digest string, operationUuid string, ) *ImageImportManifest`

NewImageImportManifest instantiates a new ImageImportManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageImportManifestWithDefaults

`func NewImageImportManifestWithDefaults() *ImageImportManifest`

NewImageImportManifestWithDefaults instantiates a new ImageImportManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *ImageImportManifest) GetContents() ImportContentDigests`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ImageImportManifest) GetContentsOk() (*ImportContentDigests, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ImageImportManifest) SetContents(v ImportContentDigests)`

SetContents sets Contents field to given value.


### GetTags

`func (o *ImageImportManifest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImageImportManifest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImageImportManifest) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetDigest

`func (o *ImageImportManifest) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ImageImportManifest) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ImageImportManifest) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetParentDigest

`func (o *ImageImportManifest) GetParentDigest() string`

GetParentDigest returns the ParentDigest field if non-nil, zero value otherwise.

### GetParentDigestOk

`func (o *ImageImportManifest) GetParentDigestOk() (*string, bool)`

GetParentDigestOk returns a tuple with the ParentDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDigest

`func (o *ImageImportManifest) SetParentDigest(v string)`

SetParentDigest sets ParentDigest field to given value.

### HasParentDigest

`func (o *ImageImportManifest) HasParentDigest() bool`

HasParentDigest returns a boolean if a field has been set.

### GetLocalImageId

`func (o *ImageImportManifest) GetLocalImageId() string`

GetLocalImageId returns the LocalImageId field if non-nil, zero value otherwise.

### GetLocalImageIdOk

`func (o *ImageImportManifest) GetLocalImageIdOk() (*string, bool)`

GetLocalImageIdOk returns a tuple with the LocalImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalImageId

`func (o *ImageImportManifest) SetLocalImageId(v string)`

SetLocalImageId sets LocalImageId field to given value.

### HasLocalImageId

`func (o *ImageImportManifest) HasLocalImageId() bool`

HasLocalImageId returns a boolean if a field has been set.

### GetOperationUuid

`func (o *ImageImportManifest) GetOperationUuid() string`

GetOperationUuid returns the OperationUuid field if non-nil, zero value otherwise.

### GetOperationUuidOk

`func (o *ImageImportManifest) GetOperationUuidOk() (*string, bool)`

GetOperationUuidOk returns a tuple with the OperationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationUuid

`func (o *ImageImportManifest) SetOperationUuid(v string)`

SetOperationUuid sets OperationUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


