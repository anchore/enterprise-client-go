# ImageAncestor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDigest** | Pointer to **string** | The digest of the image | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Layers** | Pointer to **[]string** | The full set of layers for this image | [optional] 
**UserMarkedBase** | Pointer to **bool** | True if a specific ancestor has been marked by the user as the chosen base image | [optional] 
**ChosenBaseImage** | Pointer to **bool** | True for the specific ancestor that has been identified as the base image by the system. This image will be used internally for comparisons. | [optional] 

## Methods

### NewImageAncestor

`func NewImageAncestor() *ImageAncestor`

NewImageAncestor instantiates a new ImageAncestor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageAncestorWithDefaults

`func NewImageAncestorWithDefaults() *ImageAncestor`

NewImageAncestorWithDefaults instantiates a new ImageAncestor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDigest

`func (o *ImageAncestor) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ImageAncestor) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ImageAncestor) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ImageAncestor) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetTags

`func (o *ImageAncestor) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImageAncestor) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImageAncestor) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImageAncestor) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLayers

`func (o *ImageAncestor) GetLayers() []string`

GetLayers returns the Layers field if non-nil, zero value otherwise.

### GetLayersOk

`func (o *ImageAncestor) GetLayersOk() (*[]string, bool)`

GetLayersOk returns a tuple with the Layers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayers

`func (o *ImageAncestor) SetLayers(v []string)`

SetLayers sets Layers field to given value.

### HasLayers

`func (o *ImageAncestor) HasLayers() bool`

HasLayers returns a boolean if a field has been set.

### GetUserMarkedBase

`func (o *ImageAncestor) GetUserMarkedBase() bool`

GetUserMarkedBase returns the UserMarkedBase field if non-nil, zero value otherwise.

### GetUserMarkedBaseOk

`func (o *ImageAncestor) GetUserMarkedBaseOk() (*bool, bool)`

GetUserMarkedBaseOk returns a tuple with the UserMarkedBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMarkedBase

`func (o *ImageAncestor) SetUserMarkedBase(v bool)`

SetUserMarkedBase sets UserMarkedBase field to given value.

### HasUserMarkedBase

`func (o *ImageAncestor) HasUserMarkedBase() bool`

HasUserMarkedBase returns a boolean if a field has been set.

### GetChosenBaseImage

`func (o *ImageAncestor) GetChosenBaseImage() bool`

GetChosenBaseImage returns the ChosenBaseImage field if non-nil, zero value otherwise.

### GetChosenBaseImageOk

`func (o *ImageAncestor) GetChosenBaseImageOk() (*bool, bool)`

GetChosenBaseImageOk returns a tuple with the ChosenBaseImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChosenBaseImage

`func (o *ImageAncestor) SetChosenBaseImage(v bool)`

SetChosenBaseImage sets ChosenBaseImage field to given value.

### HasChosenBaseImage

`func (o *ImageAncestor) HasChosenBaseImage() bool`

HasChosenBaseImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


