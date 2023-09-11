# ImageImportContentSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**ImportPackageLocation**](ImportPackageLocation.md) |  | 
**ContentSearches** | [**[]ImportContentSearchElement**](ImportContentSearchElement.md) |  | 

## Methods

### NewImageImportContentSearch

`func NewImageImportContentSearch(location ImportPackageLocation, contentSearches []ImportContentSearchElement, ) *ImageImportContentSearch`

NewImageImportContentSearch instantiates a new ImageImportContentSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageImportContentSearchWithDefaults

`func NewImageImportContentSearchWithDefaults() *ImageImportContentSearch`

NewImageImportContentSearchWithDefaults instantiates a new ImageImportContentSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ImageImportContentSearch) GetLocation() ImportPackageLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ImageImportContentSearch) GetLocationOk() (*ImportPackageLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ImageImportContentSearch) SetLocation(v ImportPackageLocation)`

SetLocation sets Location field to given value.


### GetContentSearches

`func (o *ImageImportContentSearch) GetContentSearches() []ImportContentSearchElement`

GetContentSearches returns the ContentSearches field if non-nil, zero value otherwise.

### GetContentSearchesOk

`func (o *ImageImportContentSearch) GetContentSearchesOk() (*[]ImportContentSearchElement, bool)`

GetContentSearchesOk returns a tuple with the ContentSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSearches

`func (o *ImageImportContentSearch) SetContentSearches(v []ImportContentSearchElement)`

SetContentSearches sets ContentSearches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


