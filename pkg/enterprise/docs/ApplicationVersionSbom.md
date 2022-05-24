# ApplicationVersionSbom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**Application**](Application.md) |  | [optional] 
**ApplicationVersion** | Pointer to [**ApplicationVersion**](ApplicationVersion.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the application version sbom was created | [optional] 
**SourceSboms** | Pointer to **[]interface{}** |  | [optional] 
**ImageSboms** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewApplicationVersionSbom

`func NewApplicationVersionSbom() *ApplicationVersionSbom`

NewApplicationVersionSbom instantiates a new ApplicationVersionSbom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationVersionSbomWithDefaults

`func NewApplicationVersionSbomWithDefaults() *ApplicationVersionSbom`

NewApplicationVersionSbomWithDefaults instantiates a new ApplicationVersionSbom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ApplicationVersionSbom) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApplicationVersionSbom) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApplicationVersionSbom) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApplicationVersionSbom) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetApplicationVersion

`func (o *ApplicationVersionSbom) GetApplicationVersion() ApplicationVersion`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *ApplicationVersionSbom) GetApplicationVersionOk() (*ApplicationVersion, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *ApplicationVersionSbom) SetApplicationVersion(v ApplicationVersion)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *ApplicationVersionSbom) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationVersionSbom) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationVersionSbom) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationVersionSbom) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationVersionSbom) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSourceSboms

`func (o *ApplicationVersionSbom) GetSourceSboms() []interface{}`

GetSourceSboms returns the SourceSboms field if non-nil, zero value otherwise.

### GetSourceSbomsOk

`func (o *ApplicationVersionSbom) GetSourceSbomsOk() (*[]interface{}, bool)`

GetSourceSbomsOk returns a tuple with the SourceSboms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSboms

`func (o *ApplicationVersionSbom) SetSourceSboms(v []interface{})`

SetSourceSboms sets SourceSboms field to given value.

### HasSourceSboms

`func (o *ApplicationVersionSbom) HasSourceSboms() bool`

HasSourceSboms returns a boolean if a field has been set.

### GetImageSboms

`func (o *ApplicationVersionSbom) GetImageSboms() []interface{}`

GetImageSboms returns the ImageSboms field if non-nil, zero value otherwise.

### GetImageSbomsOk

`func (o *ApplicationVersionSbom) GetImageSbomsOk() (*[]interface{}, bool)`

GetImageSbomsOk returns a tuple with the ImageSboms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSboms

`func (o *ApplicationVersionSbom) SetImageSboms(v []interface{})`

SetImageSboms sets ImageSboms field to given value.

### HasImageSboms

`func (o *ApplicationVersionSbom) HasImageSboms() bool`

HasImageSboms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


