# TagEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fulltag** | Pointer to **string** | The pullable string for the tag. E.g. \&quot;docker.io/library/node:latest\&quot; | [optional] 
**Registry** | Pointer to **string** | The registry hostname:port section of the pull string | [optional] 
**Repo** | Pointer to **string** | The repository section of the pull string | [optional] 
**Tag** | Pointer to **string** | The tag-only section of the pull string | [optional] 
**TagDetectedAt** | Pointer to **time.Time** | The timestamp at which the Anchore Engine detected this tag was mapped to the image digest. Does not necessarily indicate when the tag was actually pushed to the registry. | [optional] 

## Methods

### NewTagEntry

`func NewTagEntry() *TagEntry`

NewTagEntry instantiates a new TagEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagEntryWithDefaults

`func NewTagEntryWithDefaults() *TagEntry`

NewTagEntryWithDefaults instantiates a new TagEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulltag

`func (o *TagEntry) GetFulltag() string`

GetFulltag returns the Fulltag field if non-nil, zero value otherwise.

### GetFulltagOk

`func (o *TagEntry) GetFulltagOk() (*string, bool)`

GetFulltagOk returns a tuple with the Fulltag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulltag

`func (o *TagEntry) SetFulltag(v string)`

SetFulltag sets Fulltag field to given value.

### HasFulltag

`func (o *TagEntry) HasFulltag() bool`

HasFulltag returns a boolean if a field has been set.

### GetRegistry

`func (o *TagEntry) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *TagEntry) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *TagEntry) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *TagEntry) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRepo

`func (o *TagEntry) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *TagEntry) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *TagEntry) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *TagEntry) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetTag

`func (o *TagEntry) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TagEntry) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TagEntry) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *TagEntry) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTagDetectedAt

`func (o *TagEntry) GetTagDetectedAt() time.Time`

GetTagDetectedAt returns the TagDetectedAt field if non-nil, zero value otherwise.

### GetTagDetectedAtOk

`func (o *TagEntry) GetTagDetectedAtOk() (*time.Time, bool)`

GetTagDetectedAtOk returns a tuple with the TagDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagDetectedAt

`func (o *TagEntry) SetTagDetectedAt(v time.Time)`

SetTagDetectedAt sets TagDetectedAt field to given value.

### HasTagDetectedAt

`func (o *TagEntry) HasTagDetectedAt() bool`

HasTagDetectedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


