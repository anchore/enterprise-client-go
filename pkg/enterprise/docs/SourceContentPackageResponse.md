# SourceContentPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Content** | Pointer to [**[]SourceContentPackageResponseContent**](SourceContentPackageResponseContent.md) |  | [optional] 

## Methods

### NewSourceContentPackageResponse

`func NewSourceContentPackageResponse() *SourceContentPackageResponse`

NewSourceContentPackageResponse instantiates a new SourceContentPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceContentPackageResponseWithDefaults

`func NewSourceContentPackageResponseWithDefaults() *SourceContentPackageResponse`

NewSourceContentPackageResponseWithDefaults instantiates a new SourceContentPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *SourceContentPackageResponse) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SourceContentPackageResponse) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SourceContentPackageResponse) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *SourceContentPackageResponse) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetContentType

`func (o *SourceContentPackageResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SourceContentPackageResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SourceContentPackageResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SourceContentPackageResponse) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContent

`func (o *SourceContentPackageResponse) GetContent() []SourceContentPackageResponseContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SourceContentPackageResponse) GetContentOk() (*[]SourceContentPackageResponseContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SourceContentPackageResponse) SetContent(v []SourceContentPackageResponseContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *SourceContentPackageResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


