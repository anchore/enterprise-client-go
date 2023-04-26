# ImageImportFileCoordinate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The path on the filesystem of the file within the given layer | [optional] 
**LayerID** | Pointer to **string** | The image layer in which the file was found | [optional] 

## Methods

### NewImageImportFileCoordinate

`func NewImageImportFileCoordinate() *ImageImportFileCoordinate`

NewImageImportFileCoordinate instantiates a new ImageImportFileCoordinate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageImportFileCoordinateWithDefaults

`func NewImageImportFileCoordinateWithDefaults() *ImageImportFileCoordinate`

NewImageImportFileCoordinateWithDefaults instantiates a new ImageImportFileCoordinate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ImageImportFileCoordinate) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ImageImportFileCoordinate) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ImageImportFileCoordinate) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ImageImportFileCoordinate) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetLayerID

`func (o *ImageImportFileCoordinate) GetLayerID() string`

GetLayerID returns the LayerID field if non-nil, zero value otherwise.

### GetLayerIDOk

`func (o *ImageImportFileCoordinate) GetLayerIDOk() (*string, bool)`

GetLayerIDOk returns a tuple with the LayerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerID

`func (o *ImageImportFileCoordinate) SetLayerID(v string)`

SetLayerID sets LayerID field to given value.

### HasLayerID

`func (o *ImageImportFileCoordinate) HasLayerID() bool`

HasLayerID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


