# ImportContentDigests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Packages** | **string** | Digest to use for the packages content | 
**ImageConfig** | **string** | Digest for reference content for image config | 
**Manifest** | **string** | Digest to reference content for the image manifest | 
**ParentManifest** | Pointer to **string** | Digest for reference content for parent manifest | [optional] 
**Dockerfile** | Pointer to **string** | Digest for reference content for dockerfile | [optional] 
**SecretSearches** | Pointer to **string** | Digest for reference content for secret search results | [optional] 
**ContentSearches** | Pointer to **string** | Digest for reference content for content search results | [optional] 
**FileContents** | Pointer to **string** | Digest for reference content for file retrieve content | [optional] 

## Methods

### NewImportContentDigests

`func NewImportContentDigests(packages string, imageConfig string, manifest string, ) *ImportContentDigests`

NewImportContentDigests instantiates a new ImportContentDigests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportContentDigestsWithDefaults

`func NewImportContentDigestsWithDefaults() *ImportContentDigests`

NewImportContentDigestsWithDefaults instantiates a new ImportContentDigests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackages

`func (o *ImportContentDigests) GetPackages() string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *ImportContentDigests) GetPackagesOk() (*string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *ImportContentDigests) SetPackages(v string)`

SetPackages sets Packages field to given value.


### GetImageConfig

`func (o *ImportContentDigests) GetImageConfig() string`

GetImageConfig returns the ImageConfig field if non-nil, zero value otherwise.

### GetImageConfigOk

`func (o *ImportContentDigests) GetImageConfigOk() (*string, bool)`

GetImageConfigOk returns a tuple with the ImageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfig

`func (o *ImportContentDigests) SetImageConfig(v string)`

SetImageConfig sets ImageConfig field to given value.


### GetManifest

`func (o *ImportContentDigests) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *ImportContentDigests) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *ImportContentDigests) SetManifest(v string)`

SetManifest sets Manifest field to given value.


### GetParentManifest

`func (o *ImportContentDigests) GetParentManifest() string`

GetParentManifest returns the ParentManifest field if non-nil, zero value otherwise.

### GetParentManifestOk

`func (o *ImportContentDigests) GetParentManifestOk() (*string, bool)`

GetParentManifestOk returns a tuple with the ParentManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentManifest

`func (o *ImportContentDigests) SetParentManifest(v string)`

SetParentManifest sets ParentManifest field to given value.

### HasParentManifest

`func (o *ImportContentDigests) HasParentManifest() bool`

HasParentManifest returns a boolean if a field has been set.

### GetDockerfile

`func (o *ImportContentDigests) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *ImportContentDigests) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *ImportContentDigests) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *ImportContentDigests) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetSecretSearches

`func (o *ImportContentDigests) GetSecretSearches() string`

GetSecretSearches returns the SecretSearches field if non-nil, zero value otherwise.

### GetSecretSearchesOk

`func (o *ImportContentDigests) GetSecretSearchesOk() (*string, bool)`

GetSecretSearchesOk returns a tuple with the SecretSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSearches

`func (o *ImportContentDigests) SetSecretSearches(v string)`

SetSecretSearches sets SecretSearches field to given value.

### HasSecretSearches

`func (o *ImportContentDigests) HasSecretSearches() bool`

HasSecretSearches returns a boolean if a field has been set.

### GetContentSearches

`func (o *ImportContentDigests) GetContentSearches() string`

GetContentSearches returns the ContentSearches field if non-nil, zero value otherwise.

### GetContentSearchesOk

`func (o *ImportContentDigests) GetContentSearchesOk() (*string, bool)`

GetContentSearchesOk returns a tuple with the ContentSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSearches

`func (o *ImportContentDigests) SetContentSearches(v string)`

SetContentSearches sets ContentSearches field to given value.

### HasContentSearches

`func (o *ImportContentDigests) HasContentSearches() bool`

HasContentSearches returns a boolean if a field has been set.

### GetFileContents

`func (o *ImportContentDigests) GetFileContents() string`

GetFileContents returns the FileContents field if non-nil, zero value otherwise.

### GetFileContentsOk

`func (o *ImportContentDigests) GetFileContentsOk() (*string, bool)`

GetFileContentsOk returns a tuple with the FileContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContents

`func (o *ImportContentDigests) SetFileContents(v string)`

SetFileContents sets FileContents field to given value.

### HasFileContents

`func (o *ImportContentDigests) HasFileContents() bool`

HasFileContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


