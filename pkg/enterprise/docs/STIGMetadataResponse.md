# STIGMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluationUuid** | Pointer to **string** |  | [optional] 
**UploadedAt** | Pointer to **time.Time** |  | [optional] 
**StigProfile** | Pointer to **string** | The name of the STIG profile that produced this result | [optional] 
**EvaluationFormat** | Pointer to **string** | The format of the STIG content | [optional] 

## Methods

### NewSTIGMetadataResponse

`func NewSTIGMetadataResponse() *STIGMetadataResponse`

NewSTIGMetadataResponse instantiates a new STIGMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSTIGMetadataResponseWithDefaults

`func NewSTIGMetadataResponseWithDefaults() *STIGMetadataResponse`

NewSTIGMetadataResponseWithDefaults instantiates a new STIGMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluationUuid

`func (o *STIGMetadataResponse) GetEvaluationUuid() string`

GetEvaluationUuid returns the EvaluationUuid field if non-nil, zero value otherwise.

### GetEvaluationUuidOk

`func (o *STIGMetadataResponse) GetEvaluationUuidOk() (*string, bool)`

GetEvaluationUuidOk returns a tuple with the EvaluationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationUuid

`func (o *STIGMetadataResponse) SetEvaluationUuid(v string)`

SetEvaluationUuid sets EvaluationUuid field to given value.

### HasEvaluationUuid

`func (o *STIGMetadataResponse) HasEvaluationUuid() bool`

HasEvaluationUuid returns a boolean if a field has been set.

### GetUploadedAt

`func (o *STIGMetadataResponse) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *STIGMetadataResponse) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *STIGMetadataResponse) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *STIGMetadataResponse) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetStigProfile

`func (o *STIGMetadataResponse) GetStigProfile() string`

GetStigProfile returns the StigProfile field if non-nil, zero value otherwise.

### GetStigProfileOk

`func (o *STIGMetadataResponse) GetStigProfileOk() (*string, bool)`

GetStigProfileOk returns a tuple with the StigProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStigProfile

`func (o *STIGMetadataResponse) SetStigProfile(v string)`

SetStigProfile sets StigProfile field to given value.

### HasStigProfile

`func (o *STIGMetadataResponse) HasStigProfile() bool`

HasStigProfile returns a boolean if a field has been set.

### GetEvaluationFormat

`func (o *STIGMetadataResponse) GetEvaluationFormat() string`

GetEvaluationFormat returns the EvaluationFormat field if non-nil, zero value otherwise.

### GetEvaluationFormatOk

`func (o *STIGMetadataResponse) GetEvaluationFormatOk() (*string, bool)`

GetEvaluationFormatOk returns a tuple with the EvaluationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationFormat

`func (o *STIGMetadataResponse) SetEvaluationFormat(v string)`

SetEvaluationFormat sets EvaluationFormat field to given value.

### HasEvaluationFormat

`func (o *STIGMetadataResponse) HasEvaluationFormat() bool`

HasEvaluationFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


