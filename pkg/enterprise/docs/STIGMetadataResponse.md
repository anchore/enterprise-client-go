# STIGMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentUuid** | Pointer to **string** |  | [optional] 
**UploadedAt** | Pointer to **time.Time** |  | [optional] 
**StigProfile** | Pointer to **string** | The name of the STIG profile that produced this result | [optional] 
**DocumentType** | Pointer to **string** | The format of the STIG content | [optional] 

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

### GetContentUuid

`func (o *STIGMetadataResponse) GetContentUuid() string`

GetContentUuid returns the ContentUuid field if non-nil, zero value otherwise.

### GetContentUuidOk

`func (o *STIGMetadataResponse) GetContentUuidOk() (*string, bool)`

GetContentUuidOk returns a tuple with the ContentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUuid

`func (o *STIGMetadataResponse) SetContentUuid(v string)`

SetContentUuid sets ContentUuid field to given value.

### HasContentUuid

`func (o *STIGMetadataResponse) HasContentUuid() bool`

HasContentUuid returns a boolean if a field has been set.

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

### GetDocumentType

`func (o *STIGMetadataResponse) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *STIGMetadataResponse) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *STIGMetadataResponse) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *STIGMetadataResponse) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


