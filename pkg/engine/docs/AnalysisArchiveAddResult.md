# AnalysisArchiveAddResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** | The image digest requested to be added | [optional] 
**Status** | Pointer to **string** | The status of the archive add operation. Typically either &#39;archived&#39; or &#39;error&#39; | [optional] 
**Detail** | Pointer to **string** | Details on the status, e.g. the error message | [optional] 

## Methods

### NewAnalysisArchiveAddResult

`func NewAnalysisArchiveAddResult() *AnalysisArchiveAddResult`

NewAnalysisArchiveAddResult instantiates a new AnalysisArchiveAddResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisArchiveAddResultWithDefaults

`func NewAnalysisArchiveAddResultWithDefaults() *AnalysisArchiveAddResult`

NewAnalysisArchiveAddResultWithDefaults instantiates a new AnalysisArchiveAddResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *AnalysisArchiveAddResult) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *AnalysisArchiveAddResult) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *AnalysisArchiveAddResult) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *AnalysisArchiveAddResult) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetStatus

`func (o *AnalysisArchiveAddResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnalysisArchiveAddResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnalysisArchiveAddResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AnalysisArchiveAddResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *AnalysisArchiveAddResult) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AnalysisArchiveAddResult) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AnalysisArchiveAddResult) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AnalysisArchiveAddResult) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


