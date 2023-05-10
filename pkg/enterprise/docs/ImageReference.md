# ImageReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDigest** | Pointer to **string** | The image digest | [optional] 
**ImageId** | Pointer to **string** | The image id if available | [optional] 
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

### GetImageDigest

`func (o *ImageReference) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ImageReference) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ImageReference) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ImageReference) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetImageId

`func (o *ImageReference) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ImageReference) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ImageReference) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ImageReference) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

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


