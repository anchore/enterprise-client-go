# RegistryDigestSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PullString** | **string** | A digest-based pull string (e.g. docker.io/nginx@sha256:123abc) | 
**Tag** | **string** | A valid docker tag reference (e.g. docker.io/nginx:latest) that will be associated with the image but not used to pull the image. | 
**CreationTimestampOverride** | Pointer to **time.Time** | Optional override of the image creation time to support proper tag history construction in cases of out-of-order analysis compared to registry history for the tag | [optional] 
**Dockerfile** | Pointer to **string** | Base64 encoded content of the dockerfile used to build the image, if available. | [optional] 

## Methods

### NewRegistryDigestSource

`func NewRegistryDigestSource(pullString string, tag string, ) *RegistryDigestSource`

NewRegistryDigestSource instantiates a new RegistryDigestSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryDigestSourceWithDefaults

`func NewRegistryDigestSourceWithDefaults() *RegistryDigestSource`

NewRegistryDigestSourceWithDefaults instantiates a new RegistryDigestSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPullString

`func (o *RegistryDigestSource) GetPullString() string`

GetPullString returns the PullString field if non-nil, zero value otherwise.

### GetPullStringOk

`func (o *RegistryDigestSource) GetPullStringOk() (*string, bool)`

GetPullStringOk returns a tuple with the PullString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullString

`func (o *RegistryDigestSource) SetPullString(v string)`

SetPullString sets PullString field to given value.


### GetTag

`func (o *RegistryDigestSource) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RegistryDigestSource) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RegistryDigestSource) SetTag(v string)`

SetTag sets Tag field to given value.


### GetCreationTimestampOverride

`func (o *RegistryDigestSource) GetCreationTimestampOverride() time.Time`

GetCreationTimestampOverride returns the CreationTimestampOverride field if non-nil, zero value otherwise.

### GetCreationTimestampOverrideOk

`func (o *RegistryDigestSource) GetCreationTimestampOverrideOk() (*time.Time, bool)`

GetCreationTimestampOverrideOk returns a tuple with the CreationTimestampOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestampOverride

`func (o *RegistryDigestSource) SetCreationTimestampOverride(v time.Time)`

SetCreationTimestampOverride sets CreationTimestampOverride field to given value.

### HasCreationTimestampOverride

`func (o *RegistryDigestSource) HasCreationTimestampOverride() bool`

HasCreationTimestampOverride returns a boolean if a field has been set.

### GetDockerfile

`func (o *RegistryDigestSource) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *RegistryDigestSource) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *RegistryDigestSource) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *RegistryDigestSource) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


