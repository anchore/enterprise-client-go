# ImageContentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | The CPU architecture of the image | [optional] 
**Distro** | Pointer to **string** | The distro of the image | [optional] 
**DistroVersion** | Pointer to **string** | The distro version of the image | [optional] 
**DockerfileMode** | Pointer to **string** | The mode of the dockerfile | [optional] 
**ImageSize** | Pointer to **int32** | The size of the image in bytes | [optional] 
**LayerCount** | Pointer to **int32** | The number of layers in the image | [optional] 

## Methods

### NewImageContentMetadata

`func NewImageContentMetadata() *ImageContentMetadata`

NewImageContentMetadata instantiates a new ImageContentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageContentMetadataWithDefaults

`func NewImageContentMetadataWithDefaults() *ImageContentMetadata`

NewImageContentMetadataWithDefaults instantiates a new ImageContentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *ImageContentMetadata) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ImageContentMetadata) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ImageContentMetadata) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ImageContentMetadata) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetDistro

`func (o *ImageContentMetadata) GetDistro() string`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *ImageContentMetadata) GetDistroOk() (*string, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *ImageContentMetadata) SetDistro(v string)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *ImageContentMetadata) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### GetDistroVersion

`func (o *ImageContentMetadata) GetDistroVersion() string`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *ImageContentMetadata) GetDistroVersionOk() (*string, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *ImageContentMetadata) SetDistroVersion(v string)`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *ImageContentMetadata) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### GetDockerfileMode

`func (o *ImageContentMetadata) GetDockerfileMode() string`

GetDockerfileMode returns the DockerfileMode field if non-nil, zero value otherwise.

### GetDockerfileModeOk

`func (o *ImageContentMetadata) GetDockerfileModeOk() (*string, bool)`

GetDockerfileModeOk returns a tuple with the DockerfileMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfileMode

`func (o *ImageContentMetadata) SetDockerfileMode(v string)`

SetDockerfileMode sets DockerfileMode field to given value.

### HasDockerfileMode

`func (o *ImageContentMetadata) HasDockerfileMode() bool`

HasDockerfileMode returns a boolean if a field has been set.

### GetImageSize

`func (o *ImageContentMetadata) GetImageSize() int32`

GetImageSize returns the ImageSize field if non-nil, zero value otherwise.

### GetImageSizeOk

`func (o *ImageContentMetadata) GetImageSizeOk() (*int32, bool)`

GetImageSizeOk returns a tuple with the ImageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSize

`func (o *ImageContentMetadata) SetImageSize(v int32)`

SetImageSize sets ImageSize field to given value.

### HasImageSize

`func (o *ImageContentMetadata) HasImageSize() bool`

HasImageSize returns a boolean if a field has been set.

### GetLayerCount

`func (o *ImageContentMetadata) GetLayerCount() int32`

GetLayerCount returns the LayerCount field if non-nil, zero value otherwise.

### GetLayerCountOk

`func (o *ImageContentMetadata) GetLayerCountOk() (*int32, bool)`

GetLayerCountOk returns a tuple with the LayerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerCount

`func (o *ImageContentMetadata) SetLayerCount(v int32)`

SetLayerCount sets LayerCount field to given value.

### HasLayerCount

`func (o *ImageContentMetadata) HasLayerCount() bool`

HasLayerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


