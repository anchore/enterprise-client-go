# ImageAnalysisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageType** | Pointer to **string** | Optional. The type of image this is adding, defaults to \&quot;docker\&quot;. | [optional] 
**Annotations** | Pointer to **interface{}** | Annotations to be associated with the added image in key/value form | [optional] 
**Source** | Pointer to [**ImageSource**](ImageSource.md) |  | [optional] 

## Methods

### NewImageAnalysisRequest

`func NewImageAnalysisRequest() *ImageAnalysisRequest`

NewImageAnalysisRequest instantiates a new ImageAnalysisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageAnalysisRequestWithDefaults

`func NewImageAnalysisRequestWithDefaults() *ImageAnalysisRequest`

NewImageAnalysisRequestWithDefaults instantiates a new ImageAnalysisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageType

`func (o *ImageAnalysisRequest) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ImageAnalysisRequest) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ImageAnalysisRequest) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *ImageAnalysisRequest) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetAnnotations

`func (o *ImageAnalysisRequest) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ImageAnalysisRequest) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ImageAnalysisRequest) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ImageAnalysisRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetSource

`func (o *ImageAnalysisRequest) GetSource() ImageSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ImageAnalysisRequest) GetSourceOk() (*ImageSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ImageAnalysisRequest) SetSource(v ImageSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ImageAnalysisRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


