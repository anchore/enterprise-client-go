# ImageImportContentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewImageImportContentResponse

`func NewImageImportContentResponse() *ImageImportContentResponse`

NewImageImportContentResponse instantiates a new ImageImportContentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageImportContentResponseWithDefaults

`func NewImageImportContentResponseWithDefaults() *ImageImportContentResponse`

NewImageImportContentResponseWithDefaults instantiates a new ImageImportContentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *ImageImportContentResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ImageImportContentResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ImageImportContentResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ImageImportContentResponse) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ImageImportContentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImageImportContentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImageImportContentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ImageImportContentResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


