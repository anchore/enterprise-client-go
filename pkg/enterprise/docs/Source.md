# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | A system-assigned identifier unique for each source analysis | [optional] 
**AccountName** | Pointer to **string** | The anchore account id that owns this resource | [optional] 
**Host** | Pointer to **string** | Host name for the repository location (e.g. github.com) | [optional] 
**RepositoryName** | Pointer to **string** | The name of the repository on the host (e.g. &#39;anchore/anchore-engine&#39;) | [optional] 
**Revision** | Pointer to **string** | The commit ID for a git repository | [optional] 
**AnalysisStatus** | Pointer to **string** | The analysis state of the source | [optional] 
**SourceStatus** | Pointer to **string** | The state of the source | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSource

`func NewSource() *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Source) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Source) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Source) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Source) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccountName

`func (o *Source) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Source) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Source) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Source) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetHost

`func (o *Source) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Source) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Source) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Source) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetRepositoryName

`func (o *Source) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *Source) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *Source) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.

### HasRepositoryName

`func (o *Source) HasRepositoryName() bool`

HasRepositoryName returns a boolean if a field has been set.

### GetRevision

`func (o *Source) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Source) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Source) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *Source) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetAnalysisStatus

`func (o *Source) GetAnalysisStatus() string`

GetAnalysisStatus returns the AnalysisStatus field if non-nil, zero value otherwise.

### GetAnalysisStatusOk

`func (o *Source) GetAnalysisStatusOk() (*string, bool)`

GetAnalysisStatusOk returns a tuple with the AnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatus

`func (o *Source) SetAnalysisStatus(v string)`

SetAnalysisStatus sets AnalysisStatus field to given value.

### HasAnalysisStatus

`func (o *Source) HasAnalysisStatus() bool`

HasAnalysisStatus returns a boolean if a field has been set.

### GetSourceStatus

`func (o *Source) GetSourceStatus() string`

GetSourceStatus returns the SourceStatus field if non-nil, zero value otherwise.

### GetSourceStatusOk

`func (o *Source) GetSourceStatusOk() (*string, bool)`

GetSourceStatusOk returns a tuple with the SourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStatus

`func (o *Source) SetSourceStatus(v string)`

SetSourceStatus sets SourceStatus field to given value.

### HasSourceStatus

`func (o *Source) HasSourceStatus() bool`

HasSourceStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Source) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Source) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Source) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Source) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Source) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Source) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Source) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Source) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


