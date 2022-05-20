# ImageReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** | The image digest | [optional] 
**Id** | Pointer to **string** | The image id if available | [optional] 
**AnalyzedAt** | Pointer to **string** | Timestamp, in rfc3339 format, indicating when the image state became &#39;analyzed&#39; in Anchore Engine. | [optional] 
**TagHistory** | Pointer to [**[]TagEntry**](TagEntry.md) |  | [optional] 

## Methods

### NewImageReference

`func NewImageReference() *ImageReference`

NewImageReference instantiates a new ImageReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageReferenceWithDefaults

`func NewImageReferenceWithDefaults() *ImageReference`

NewImageReferenceWithDefaults instantiates a new ImageReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *ImageReference) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ImageReference) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ImageReference) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ImageReference) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetId

`func (o *ImageReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageReference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageReference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAnalyzedAt

`func (o *ImageReference) GetAnalyzedAt() string`

GetAnalyzedAt returns the AnalyzedAt field if non-nil, zero value otherwise.

### GetAnalyzedAtOk

`func (o *ImageReference) GetAnalyzedAtOk() (*string, bool)`

GetAnalyzedAtOk returns a tuple with the AnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzedAt

`func (o *ImageReference) SetAnalyzedAt(v string)`

SetAnalyzedAt sets AnalyzedAt field to given value.

### HasAnalyzedAt

`func (o *ImageReference) HasAnalyzedAt() bool`

HasAnalyzedAt returns a boolean if a field has been set.

### GetTagHistory

`func (o *ImageReference) GetTagHistory() []TagEntry`

GetTagHistory returns the TagHistory field if non-nil, zero value otherwise.

### GetTagHistoryOk

`func (o *ImageReference) GetTagHistoryOk() (*[]TagEntry, bool)`

GetTagHistoryOk returns a tuple with the TagHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagHistory

`func (o *ImageReference) SetTagHistory(v []TagEntry)`

SetTagHistory sets TagHistory field to given value.

### HasTagHistory

`func (o *ImageReference) HasTagHistory() bool`

HasTagHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


