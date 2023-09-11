# ImageSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to [**NullableRegistryTagSource**](RegistryTagSource.md) |  | [optional] 
**Digest** | Pointer to [**NullableRegistryDigestSource**](RegistryDigestSource.md) |  | [optional] 
**Archive** | Pointer to [**NullableAnalysisArchiveSource**](AnalysisArchiveSource.md) |  | [optional] 
**Import** | Pointer to [**NullableImageImportManifest**](ImageImportManifest.md) |  | [optional] 

## Methods

### NewImageSource

`func NewImageSource() *ImageSource`

NewImageSource instantiates a new ImageSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageSourceWithDefaults

`func NewImageSourceWithDefaults() *ImageSource`

NewImageSourceWithDefaults instantiates a new ImageSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *ImageSource) GetTag() RegistryTagSource`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ImageSource) GetTagOk() (*RegistryTagSource, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ImageSource) SetTag(v RegistryTagSource)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ImageSource) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ImageSource) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ImageSource) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetDigest

`func (o *ImageSource) GetDigest() RegistryDigestSource`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ImageSource) GetDigestOk() (*RegistryDigestSource, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ImageSource) SetDigest(v RegistryDigestSource)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ImageSource) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### SetDigestNil

`func (o *ImageSource) SetDigestNil(b bool)`

 SetDigestNil sets the value for Digest to be an explicit nil

### UnsetDigest
`func (o *ImageSource) UnsetDigest()`

UnsetDigest ensures that no value is present for Digest, not even an explicit nil
### GetArchive

`func (o *ImageSource) GetArchive() AnalysisArchiveSource`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *ImageSource) GetArchiveOk() (*AnalysisArchiveSource, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *ImageSource) SetArchive(v AnalysisArchiveSource)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *ImageSource) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### SetArchiveNil

`func (o *ImageSource) SetArchiveNil(b bool)`

 SetArchiveNil sets the value for Archive to be an explicit nil

### UnsetArchive
`func (o *ImageSource) UnsetArchive()`

UnsetArchive ensures that no value is present for Archive, not even an explicit nil
### GetImport

`func (o *ImageSource) GetImport() ImageImportManifest`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *ImageSource) GetImportOk() (*ImageImportManifest, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *ImageSource) SetImport(v ImageImportManifest)`

SetImport sets Import field to given value.

### HasImport

`func (o *ImageSource) HasImport() bool`

HasImport returns a boolean if a field has been set.

### SetImportNil

`func (o *ImageSource) SetImportNil(b bool)`

 SetImportNil sets the value for Import to be an explicit nil

### UnsetImport
`func (o *ImageSource) UnsetImport()`

UnsetImport ensures that no value is present for Import, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


