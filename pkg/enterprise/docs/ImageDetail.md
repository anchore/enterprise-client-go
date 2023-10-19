# ImageDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**FullTag** | Pointer to **string** | Full docker-pullable tag string referencing the image | [optional] 
**FullDigest** | Pointer to **string** | Full docker-pullable digest string including the registry url and repository necessary get the image | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**Registry** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Dockerfile** | Pointer to **NullableString** |  | [optional] 
**ImageDigest** | Pointer to **string** | The parent Anchore Image record to which this detail maps | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**TagDetectedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewImageDetail

`func NewImageDetail() *ImageDetail`

NewImageDetail instantiates a new ImageDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDetailWithDefaults

`func NewImageDetailWithDefaults() *ImageDetail`

NewImageDetailWithDefaults instantiates a new ImageDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ImageDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImageDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImageDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ImageDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ImageDetail) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ImageDetail) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ImageDetail) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ImageDetail) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetFullTag

`func (o *ImageDetail) GetFullTag() string`

GetFullTag returns the FullTag field if non-nil, zero value otherwise.

### GetFullTagOk

`func (o *ImageDetail) GetFullTagOk() (*string, bool)`

GetFullTagOk returns a tuple with the FullTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTag

`func (o *ImageDetail) SetFullTag(v string)`

SetFullTag sets FullTag field to given value.

### HasFullTag

`func (o *ImageDetail) HasFullTag() bool`

HasFullTag returns a boolean if a field has been set.

### GetFullDigest

`func (o *ImageDetail) GetFullDigest() string`

GetFullDigest returns the FullDigest field if non-nil, zero value otherwise.

### GetFullDigestOk

`func (o *ImageDetail) GetFullDigestOk() (*string, bool)`

GetFullDigestOk returns a tuple with the FullDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDigest

`func (o *ImageDetail) SetFullDigest(v string)`

SetFullDigest sets FullDigest field to given value.

### HasFullDigest

`func (o *ImageDetail) HasFullDigest() bool`

HasFullDigest returns a boolean if a field has been set.

### GetAccountName

`func (o *ImageDetail) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ImageDetail) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ImageDetail) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *ImageDetail) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetImageId

`func (o *ImageDetail) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ImageDetail) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ImageDetail) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ImageDetail) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetRegistry

`func (o *ImageDetail) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *ImageDetail) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *ImageDetail) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *ImageDetail) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRepo

`func (o *ImageDetail) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ImageDetail) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ImageDetail) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ImageDetail) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetDockerfile

`func (o *ImageDetail) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *ImageDetail) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *ImageDetail) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *ImageDetail) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### SetDockerfileNil

`func (o *ImageDetail) SetDockerfileNil(b bool)`

 SetDockerfileNil sets the value for Dockerfile to be an explicit nil

### UnsetDockerfile
`func (o *ImageDetail) UnsetDockerfile()`

UnsetDockerfile ensures that no value is present for Dockerfile, not even an explicit nil
### GetImageDigest

`func (o *ImageDetail) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ImageDetail) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ImageDetail) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ImageDetail) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetTag

`func (o *ImageDetail) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ImageDetail) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ImageDetail) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ImageDetail) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTagDetectedAt

`func (o *ImageDetail) GetTagDetectedAt() time.Time`

GetTagDetectedAt returns the TagDetectedAt field if non-nil, zero value otherwise.

### GetTagDetectedAtOk

`func (o *ImageDetail) GetTagDetectedAtOk() (*time.Time, bool)`

GetTagDetectedAtOk returns a tuple with the TagDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagDetectedAt

`func (o *ImageDetail) SetTagDetectedAt(v time.Time)`

SetTagDetectedAt sets TagDetectedAt field to given value.

### HasTagDetectedAt

`func (o *ImageDetail) HasTagDetectedAt() bool`

HasTagDetectedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


