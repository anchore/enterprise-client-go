# ScanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sbom** | Pointer to [**NativeSBOM**](NativeSBOM.md) |  | [optional] 
**SecretScan** | Pointer to [**ContentSearch**](ContentSearch.md) |  | [optional] 
**ContentSearch** | Pointer to [**ContentSearch**](ContentSearch.md) |  | [optional] 
**RetrievedFiles** | Pointer to [**FileContents**](FileContents.md) |  | [optional] 
**Dockerfile** | Pointer to **NullableString** | Base64 encoded content of the dockerfile used to build the image, if available. | [optional] 

## Methods

### NewScanRequest

`func NewScanRequest() *ScanRequest`

NewScanRequest instantiates a new ScanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanRequestWithDefaults

`func NewScanRequestWithDefaults() *ScanRequest`

NewScanRequestWithDefaults instantiates a new ScanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSbom

`func (o *ScanRequest) GetSbom() NativeSBOM`

GetSbom returns the Sbom field if non-nil, zero value otherwise.

### GetSbomOk

`func (o *ScanRequest) GetSbomOk() (*NativeSBOM, bool)`

GetSbomOk returns a tuple with the Sbom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbom

`func (o *ScanRequest) SetSbom(v NativeSBOM)`

SetSbom sets Sbom field to given value.

### HasSbom

`func (o *ScanRequest) HasSbom() bool`

HasSbom returns a boolean if a field has been set.

### GetSecretScan

`func (o *ScanRequest) GetSecretScan() ContentSearch`

GetSecretScan returns the SecretScan field if non-nil, zero value otherwise.

### GetSecretScanOk

`func (o *ScanRequest) GetSecretScanOk() (*ContentSearch, bool)`

GetSecretScanOk returns a tuple with the SecretScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretScan

`func (o *ScanRequest) SetSecretScan(v ContentSearch)`

SetSecretScan sets SecretScan field to given value.

### HasSecretScan

`func (o *ScanRequest) HasSecretScan() bool`

HasSecretScan returns a boolean if a field has been set.

### GetContentSearch

`func (o *ScanRequest) GetContentSearch() ContentSearch`

GetContentSearch returns the ContentSearch field if non-nil, zero value otherwise.

### GetContentSearchOk

`func (o *ScanRequest) GetContentSearchOk() (*ContentSearch, bool)`

GetContentSearchOk returns a tuple with the ContentSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSearch

`func (o *ScanRequest) SetContentSearch(v ContentSearch)`

SetContentSearch sets ContentSearch field to given value.

### HasContentSearch

`func (o *ScanRequest) HasContentSearch() bool`

HasContentSearch returns a boolean if a field has been set.

### GetRetrievedFiles

`func (o *ScanRequest) GetRetrievedFiles() FileContents`

GetRetrievedFiles returns the RetrievedFiles field if non-nil, zero value otherwise.

### GetRetrievedFilesOk

`func (o *ScanRequest) GetRetrievedFilesOk() (*FileContents, bool)`

GetRetrievedFilesOk returns a tuple with the RetrievedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedFiles

`func (o *ScanRequest) SetRetrievedFiles(v FileContents)`

SetRetrievedFiles sets RetrievedFiles field to given value.

### HasRetrievedFiles

`func (o *ScanRequest) HasRetrievedFiles() bool`

HasRetrievedFiles returns a boolean if a field has been set.

### GetDockerfile

`func (o *ScanRequest) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *ScanRequest) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *ScanRequest) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *ScanRequest) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### SetDockerfileNil

`func (o *ScanRequest) SetDockerfileNil(b bool)`

 SetDockerfileNil sets the value for Dockerfile to be an explicit nil

### UnsetDockerfile
`func (o *ScanRequest) UnsetDockerfile()`

UnsetDockerfile ensures that no value is present for Dockerfile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


