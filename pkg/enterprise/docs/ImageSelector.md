# ImageSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registry** | Pointer to **string** | The registry section of a pull string. e.g. with \&quot;docker.io/anchore/anchore-engine:latest\&quot;, this is \&quot;docker.io\&quot; | [optional] 
**Repository** | Pointer to **string** | The repository section of a pull string. e.g. with \&quot;docker.io/anchore/anchore-engine:latest\&quot;, this is \&quot;anchore/anchore-engine\&quot; | [optional] 
**Tag** | Pointer to **string** | The tag-only section of a pull string. e.g. with \&quot;docker.io/anchore/anchore-engine:latest\&quot;, this is \&quot;latest\&quot; | [optional] 

## Methods

### NewImageSelector

`func NewImageSelector() *ImageSelector`

NewImageSelector instantiates a new ImageSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageSelectorWithDefaults

`func NewImageSelectorWithDefaults() *ImageSelector`

NewImageSelectorWithDefaults instantiates a new ImageSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistry

`func (o *ImageSelector) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *ImageSelector) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *ImageSelector) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *ImageSelector) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRepository

`func (o *ImageSelector) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ImageSelector) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ImageSelector) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ImageSelector) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTag

`func (o *ImageSelector) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ImageSelector) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ImageSelector) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ImageSelector) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


