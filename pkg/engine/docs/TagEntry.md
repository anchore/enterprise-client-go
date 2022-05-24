# TagEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pullstring** | Pointer to **string** | The pullable string for the tag. E.g. \&quot;docker.io/library/node:latest\&quot; | [optional] 
**Registry** | Pointer to **string** | The registry hostname:port section of the pull string | [optional] 
**Repository** | Pointer to **string** | The repository section of the pull string | [optional] 
**Tag** | Pointer to **string** | The tag-only section of the pull string | [optional] 
**DetectedAt** | Pointer to **time.Time** | The timestamp at which the Anchore Engine detected this tag was mapped to the image digest. Does not necessarily indicate when the tag was actually pushed to the registry. | [optional] 

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

### GetPullstring

`func (o *TagEntry) GetPullstring() string`

GetPullstring returns the Pullstring field if non-nil, zero value otherwise.

### GetPullstringOk

`func (o *TagEntry) GetPullstringOk() (*string, bool)`

GetPullstringOk returns a tuple with the Pullstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullstring

`func (o *TagEntry) SetPullstring(v string)`

SetPullstring sets Pullstring field to given value.

### HasPullstring

`func (o *TagEntry) HasPullstring() bool`

HasPullstring returns a boolean if a field has been set.

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

### GetRepository

`func (o *TagEntry) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *TagEntry) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *TagEntry) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *TagEntry) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

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

### GetDetectedAt

`func (o *TagEntry) GetDetectedAt() time.Time`

GetDetectedAt returns the DetectedAt field if non-nil, zero value otherwise.

### GetDetectedAtOk

`func (o *TagEntry) GetDetectedAtOk() (*time.Time, bool)`

GetDetectedAtOk returns a tuple with the DetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedAt

`func (o *TagEntry) SetDetectedAt(v time.Time)`

SetDetectedAt sets DetectedAt field to given value.

### HasDetectedAt

`func (o *TagEntry) HasDetectedAt() bool`

HasDetectedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


