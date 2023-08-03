# RegistryTagSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PullString** | **string** | A docker pull string (e.g. docker.io/nginx:latest, or docker.io/nginx@sha256:abd) to retrieve the image | 
**Dockerfile** | Pointer to **string** | Base64 encoded content of the dockerfile used to build the image, if available. | [optional] 

## Methods

### NewRegistryTagSource

`func NewRegistryTagSource(pullString string, ) *RegistryTagSource`

NewRegistryTagSource instantiates a new RegistryTagSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryTagSourceWithDefaults

`func NewRegistryTagSourceWithDefaults() *RegistryTagSource`

NewRegistryTagSourceWithDefaults instantiates a new RegistryTagSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPullString

`func (o *RegistryTagSource) GetPullString() string`

GetPullString returns the PullString field if non-nil, zero value otherwise.

### GetPullStringOk

`func (o *RegistryTagSource) GetPullStringOk() (*string, bool)`

GetPullStringOk returns a tuple with the PullString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullString

`func (o *RegistryTagSource) SetPullString(v string)`

SetPullString sets PullString field to given value.


### GetDockerfile

`func (o *RegistryTagSource) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *RegistryTagSource) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *RegistryTagSource) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *RegistryTagSource) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


