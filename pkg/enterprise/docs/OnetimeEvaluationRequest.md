# OnetimeEvaluationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sbom** | Pointer to [**NativeSBOM**](NativeSBOM.md) |  | [optional] 
**SecretScan** | Pointer to [**ContentSearch**](ContentSearch.md) |  | [optional] 
**ContentSearch** | Pointer to [**ContentSearch**](ContentSearch.md) |  | [optional] 
**RetrievedFiles** | Pointer to [**FileContents**](FileContents.md) |  | [optional] 

## Methods

### NewOnetimeEvaluationRequest

`func NewOnetimeEvaluationRequest() *OnetimeEvaluationRequest`

NewOnetimeEvaluationRequest instantiates a new OnetimeEvaluationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnetimeEvaluationRequestWithDefaults

`func NewOnetimeEvaluationRequestWithDefaults() *OnetimeEvaluationRequest`

NewOnetimeEvaluationRequestWithDefaults instantiates a new OnetimeEvaluationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSbom

`func (o *OnetimeEvaluationRequest) GetSbom() NativeSBOM`

GetSbom returns the Sbom field if non-nil, zero value otherwise.

### GetSbomOk

`func (o *OnetimeEvaluationRequest) GetSbomOk() (*NativeSBOM, bool)`

GetSbomOk returns a tuple with the Sbom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbom

`func (o *OnetimeEvaluationRequest) SetSbom(v NativeSBOM)`

SetSbom sets Sbom field to given value.

### HasSbom

`func (o *OnetimeEvaluationRequest) HasSbom() bool`

HasSbom returns a boolean if a field has been set.

### GetSecretScan

`func (o *OnetimeEvaluationRequest) GetSecretScan() ContentSearch`

GetSecretScan returns the SecretScan field if non-nil, zero value otherwise.

### GetSecretScanOk

`func (o *OnetimeEvaluationRequest) GetSecretScanOk() (*ContentSearch, bool)`

GetSecretScanOk returns a tuple with the SecretScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretScan

`func (o *OnetimeEvaluationRequest) SetSecretScan(v ContentSearch)`

SetSecretScan sets SecretScan field to given value.

### HasSecretScan

`func (o *OnetimeEvaluationRequest) HasSecretScan() bool`

HasSecretScan returns a boolean if a field has been set.

### GetContentSearch

`func (o *OnetimeEvaluationRequest) GetContentSearch() ContentSearch`

GetContentSearch returns the ContentSearch field if non-nil, zero value otherwise.

### GetContentSearchOk

`func (o *OnetimeEvaluationRequest) GetContentSearchOk() (*ContentSearch, bool)`

GetContentSearchOk returns a tuple with the ContentSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSearch

`func (o *OnetimeEvaluationRequest) SetContentSearch(v ContentSearch)`

SetContentSearch sets ContentSearch field to given value.

### HasContentSearch

`func (o *OnetimeEvaluationRequest) HasContentSearch() bool`

HasContentSearch returns a boolean if a field has been set.

### GetRetrievedFiles

`func (o *OnetimeEvaluationRequest) GetRetrievedFiles() FileContents`

GetRetrievedFiles returns the RetrievedFiles field if non-nil, zero value otherwise.

### GetRetrievedFilesOk

`func (o *OnetimeEvaluationRequest) GetRetrievedFilesOk() (*FileContents, bool)`

GetRetrievedFilesOk returns a tuple with the RetrievedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedFiles

`func (o *OnetimeEvaluationRequest) SetRetrievedFiles(v FileContents)`

SetRetrievedFiles sets RetrievedFiles field to given value.

### HasRetrievedFiles

`func (o *OnetimeEvaluationRequest) HasRetrievedFiles() bool`

HasRetrievedFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


