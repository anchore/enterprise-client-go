# AnalysisArchiveSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** | The image digest identify the analysis. Archived analyses are based on digest, tag records are restored as analysis is restored. | 

## Methods

### NewAnalysisArchiveSource

`func NewAnalysisArchiveSource(digest string, ) *AnalysisArchiveSource`

NewAnalysisArchiveSource instantiates a new AnalysisArchiveSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisArchiveSourceWithDefaults

`func NewAnalysisArchiveSourceWithDefaults() *AnalysisArchiveSource`

NewAnalysisArchiveSourceWithDefaults instantiates a new AnalysisArchiveSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *AnalysisArchiveSource) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *AnalysisArchiveSource) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *AnalysisArchiveSource) SetDigest(v string)`

SetDigest sets Digest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


