# ContentSearchFindingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**ContentSearchFindingsInnerLocation**](ContentSearchFindingsInnerLocation.md) |  | 
**ContentSearches** | [**[]ContentSearchFindingsInnerContentSearchesInner**](ContentSearchFindingsInnerContentSearchesInner.md) |  | 

## Methods

### NewContentSearchFindingsInner

`func NewContentSearchFindingsInner(location ContentSearchFindingsInnerLocation, contentSearches []ContentSearchFindingsInnerContentSearchesInner, ) *ContentSearchFindingsInner`

NewContentSearchFindingsInner instantiates a new ContentSearchFindingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSearchFindingsInnerWithDefaults

`func NewContentSearchFindingsInnerWithDefaults() *ContentSearchFindingsInner`

NewContentSearchFindingsInnerWithDefaults instantiates a new ContentSearchFindingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ContentSearchFindingsInner) GetLocation() ContentSearchFindingsInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ContentSearchFindingsInner) GetLocationOk() (*ContentSearchFindingsInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ContentSearchFindingsInner) SetLocation(v ContentSearchFindingsInnerLocation)`

SetLocation sets Location field to given value.


### GetContentSearches

`func (o *ContentSearchFindingsInner) GetContentSearches() []ContentSearchFindingsInnerContentSearchesInner`

GetContentSearches returns the ContentSearches field if non-nil, zero value otherwise.

### GetContentSearchesOk

`func (o *ContentSearchFindingsInner) GetContentSearchesOk() (*[]ContentSearchFindingsInnerContentSearchesInner, bool)`

GetContentSearchesOk returns a tuple with the ContentSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSearches

`func (o *ContentSearchFindingsInner) SetContentSearches(v []ContentSearchFindingsInnerContentSearchesInner)`

SetContentSearches sets ContentSearches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


