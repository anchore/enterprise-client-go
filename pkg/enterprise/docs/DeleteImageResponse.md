# DeleteImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDigest** | **string** |  | 
**Status** | **string** | Current status of the image deletion | 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteImageResponse

`func NewDeleteImageResponse(imageDigest string, status string, ) *DeleteImageResponse`

NewDeleteImageResponse instantiates a new DeleteImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteImageResponseWithDefaults

`func NewDeleteImageResponseWithDefaults() *DeleteImageResponse`

NewDeleteImageResponseWithDefaults instantiates a new DeleteImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDigest

`func (o *DeleteImageResponse) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *DeleteImageResponse) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *DeleteImageResponse) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.


### GetStatus

`func (o *DeleteImageResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeleteImageResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeleteImageResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *DeleteImageResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *DeleteImageResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *DeleteImageResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *DeleteImageResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


