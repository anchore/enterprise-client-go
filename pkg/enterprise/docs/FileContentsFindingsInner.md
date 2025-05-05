# FileContentsFindingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**FileContentsFindingsInnerLocation**](FileContentsFindingsInnerLocation.md) |  | 
**Contents** | **string** | The contents of the file | 

## Methods

### NewFileContentsFindingsInner

`func NewFileContentsFindingsInner(location FileContentsFindingsInnerLocation, contents string, ) *FileContentsFindingsInner`

NewFileContentsFindingsInner instantiates a new FileContentsFindingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileContentsFindingsInnerWithDefaults

`func NewFileContentsFindingsInnerWithDefaults() *FileContentsFindingsInner`

NewFileContentsFindingsInnerWithDefaults instantiates a new FileContentsFindingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *FileContentsFindingsInner) GetLocation() FileContentsFindingsInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FileContentsFindingsInner) GetLocationOk() (*FileContentsFindingsInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FileContentsFindingsInner) SetLocation(v FileContentsFindingsInnerLocation)`

SetLocation sets Location field to given value.


### GetContents

`func (o *FileContentsFindingsInner) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *FileContentsFindingsInner) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *FileContentsFindingsInner) SetContents(v string)`

SetContents sets Contents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


