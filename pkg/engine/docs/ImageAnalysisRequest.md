# ImageAnalysisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dockerfile** | Pointer to **string** | Base64 encoded content of the dockerfile for the image, if available. Deprecated in favor of the &#39;source&#39; field. | [optional] 
**Digest** | Pointer to **string** | A digest string for an image, maybe a pull string or just a digest. e.g. nginx@sha256:123 or sha256:abc123. If a pull string, it must have same registry/repo as the tag field. Deprecated in favor of the &#39;source&#39; field | [optional] 
**Tag** | Pointer to **string** | Full pullable tag reference for image. e.g. docker.io/nginx:latest. Deprecated in favor of the &#39;source&#39; field | [optional] 
**CreatedAt** | Pointer to **time.Time** | Optional override of the image creation time, only honored when both tag and digest are also supplied  e.g. 2018-10-17T18:14:00Z. Deprecated in favor of the &#39;source&#39; field | [optional] 
**ImageType** | Pointer to **string** | Optional. The type of image this is adding, defaults to \&quot;docker\&quot;. This can be ommitted until multiple image types are supported. | [optional] 
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

### GetDockerfile

`func (o *ImageAnalysisRequest) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *ImageAnalysisRequest) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *ImageAnalysisRequest) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *ImageAnalysisRequest) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetDigest

`func (o *ImageAnalysisRequest) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ImageAnalysisRequest) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ImageAnalysisRequest) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ImageAnalysisRequest) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetTag

`func (o *ImageAnalysisRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ImageAnalysisRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ImageAnalysisRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ImageAnalysisRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ImageAnalysisRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImageAnalysisRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImageAnalysisRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ImageAnalysisRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

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


