# SourceImportMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CiWorkflowName** | Pointer to **string** |  | [optional] 
**CiWorkflowExecutionTime** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**RepositoryName** | Pointer to **string** |  | [optional] 
**BranchName** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**ChangeAuthor** | Pointer to **string** |  | [optional] 
**Contents** | Pointer to [**SourceImportMetadataContents**](SourceImportMetadataContents.md) |  | [optional] 

## Methods

### NewSourceImportMetadata

`func NewSourceImportMetadata() *SourceImportMetadata`

NewSourceImportMetadata instantiates a new SourceImportMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceImportMetadataWithDefaults

`func NewSourceImportMetadataWithDefaults() *SourceImportMetadata`

NewSourceImportMetadataWithDefaults instantiates a new SourceImportMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiWorkflowName

`func (o *SourceImportMetadata) GetCiWorkflowName() string`

GetCiWorkflowName returns the CiWorkflowName field if non-nil, zero value otherwise.

### GetCiWorkflowNameOk

`func (o *SourceImportMetadata) GetCiWorkflowNameOk() (*string, bool)`

GetCiWorkflowNameOk returns a tuple with the CiWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiWorkflowName

`func (o *SourceImportMetadata) SetCiWorkflowName(v string)`

SetCiWorkflowName sets CiWorkflowName field to given value.

### HasCiWorkflowName

`func (o *SourceImportMetadata) HasCiWorkflowName() bool`

HasCiWorkflowName returns a boolean if a field has been set.

### GetCiWorkflowExecutionTime

`func (o *SourceImportMetadata) GetCiWorkflowExecutionTime() time.Time`

GetCiWorkflowExecutionTime returns the CiWorkflowExecutionTime field if non-nil, zero value otherwise.

### GetCiWorkflowExecutionTimeOk

`func (o *SourceImportMetadata) GetCiWorkflowExecutionTimeOk() (*time.Time, bool)`

GetCiWorkflowExecutionTimeOk returns a tuple with the CiWorkflowExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiWorkflowExecutionTime

`func (o *SourceImportMetadata) SetCiWorkflowExecutionTime(v time.Time)`

SetCiWorkflowExecutionTime sets CiWorkflowExecutionTime field to given value.

### HasCiWorkflowExecutionTime

`func (o *SourceImportMetadata) HasCiWorkflowExecutionTime() bool`

HasCiWorkflowExecutionTime returns a boolean if a field has been set.

### GetHost

`func (o *SourceImportMetadata) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SourceImportMetadata) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SourceImportMetadata) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SourceImportMetadata) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetRepositoryName

`func (o *SourceImportMetadata) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *SourceImportMetadata) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *SourceImportMetadata) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.

### HasRepositoryName

`func (o *SourceImportMetadata) HasRepositoryName() bool`

HasRepositoryName returns a boolean if a field has been set.

### GetBranchName

`func (o *SourceImportMetadata) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *SourceImportMetadata) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *SourceImportMetadata) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *SourceImportMetadata) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetRevision

`func (o *SourceImportMetadata) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SourceImportMetadata) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SourceImportMetadata) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *SourceImportMetadata) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetChangeAuthor

`func (o *SourceImportMetadata) GetChangeAuthor() string`

GetChangeAuthor returns the ChangeAuthor field if non-nil, zero value otherwise.

### GetChangeAuthorOk

`func (o *SourceImportMetadata) GetChangeAuthorOk() (*string, bool)`

GetChangeAuthorOk returns a tuple with the ChangeAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAuthor

`func (o *SourceImportMetadata) SetChangeAuthor(v string)`

SetChangeAuthor sets ChangeAuthor field to given value.

### HasChangeAuthor

`func (o *SourceImportMetadata) HasChangeAuthor() bool`

HasChangeAuthor returns a boolean if a field has been set.

### GetContents

`func (o *SourceImportMetadata) GetContents() SourceImportMetadataContents`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *SourceImportMetadata) GetContentsOk() (*SourceImportMetadataContents, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *SourceImportMetadata) SetContents(v SourceImportMetadataContents)`

SetContents sets Contents field to given value.

### HasContents

`func (o *SourceImportMetadata) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


