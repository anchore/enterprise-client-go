# ImageImportContentSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**ImportPackageLocation**](ImportPackageLocation.md) |  | 
**Contentsearches** | [**[]ImportContentSearchElement**](ImportContentSearchElement.md) |  | 

## Methods

### NewImageImportContentSearch

`func NewImageImportContentSearch(location ImportPackageLocation, contentsearches []ImportContentSearchElement, ) *ImageImportContentSearch`

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


### GetContentsearches

`func (o *ImageImportContentSearch) GetContentsearches() []ImportContentSearchElement`

GetContentsearches returns the Contentsearches field if non-nil, zero value otherwise.

### GetContentsearchesOk

`func (o *ImageImportContentSearch) GetContentsearchesOk() (*[]ImportContentSearchElement, bool)`

GetContentsearchesOk returns a tuple with the Contentsearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsearches

`func (o *ImageImportContentSearch) SetContentsearches(v []ImportContentSearchElement)`

SetContentsearches sets Contentsearches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


