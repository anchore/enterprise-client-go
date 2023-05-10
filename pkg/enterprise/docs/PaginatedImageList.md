# PaginatedImageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **string** | The page number returned (should match the requested page query string param) | [optional] 
**NextPage** | Pointer to **string** | True if additional pages exist (page + 1) or False if this is the last page | [optional] 
**ReturnedCount** | Pointer to **int32** | The number of items sent in this response | [optional] 
**Images** | Pointer to [**[]ImageWithPackages**](ImageWithPackages.md) |  | [optional] 

## Methods

### NewPaginatedImageList

`func NewPaginatedImageList() *PaginatedImageList`

NewPaginatedImageList instantiates a new PaginatedImageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedImageListWithDefaults

`func NewPaginatedImageListWithDefaults() *PaginatedImageList`

NewPaginatedImageListWithDefaults instantiates a new PaginatedImageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PaginatedImageList) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginatedImageList) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginatedImageList) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaginatedImageList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetNextPage

`func (o *PaginatedImageList) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *PaginatedImageList) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *PaginatedImageList) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *PaginatedImageList) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetReturnedCount

`func (o *PaginatedImageList) GetReturnedCount() int32`

GetReturnedCount returns the ReturnedCount field if non-nil, zero value otherwise.

### GetReturnedCountOk

`func (o *PaginatedImageList) GetReturnedCountOk() (*int32, bool)`

GetReturnedCountOk returns a tuple with the ReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedCount

`func (o *PaginatedImageList) SetReturnedCount(v int32)`

SetReturnedCount sets ReturnedCount field to given value.

### HasReturnedCount

`func (o *PaginatedImageList) HasReturnedCount() bool`

HasReturnedCount returns a boolean if a field has been set.

### GetImages

`func (o *PaginatedImageList) GetImages() []ImageWithPackages`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *PaginatedImageList) GetImagesOk() (*[]ImageWithPackages, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *PaginatedImageList) SetImages(v []ImageWithPackages)`

SetImages sets Images field to given value.

### HasImages

`func (o *PaginatedImageList) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


