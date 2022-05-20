# SourceImportMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CiWorkflowName** | Pointer to **NullableString** |  | [optional] 
**CiWorkflowExecutionTime** | Pointer to **NullableTime** |  | [optional] 
**Host** | **string** |  | 
**RepositoryName** | **string** |  | 
**BranchName** | Pointer to **NullableString** |  | [optional] 
**Revision** | **string** |  | 
**ChangeAuthor** | Pointer to **NullableString** |  | [optional] 
**Contents** | [**SourceImportMetadataContents**](SourceImportMetadataContents.md) |  | 

## Methods

### NewSourceImportMetadata

`func NewSourceImportMetadata(host string, repositoryName string, revision string, contents SourceImportMetadataContents, ) *SourceImportMetadata`

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

### SetCiWorkflowNameNil

`func (o *SourceImportMetadata) SetCiWorkflowNameNil(b bool)`

 SetCiWorkflowNameNil sets the value for CiWorkflowName to be an explicit nil

### UnsetCiWorkflowName
`func (o *SourceImportMetadata) UnsetCiWorkflowName()`

UnsetCiWorkflowName ensures that no value is present for CiWorkflowName, not even an explicit nil
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

### SetCiWorkflowExecutionTimeNil

`func (o *SourceImportMetadata) SetCiWorkflowExecutionTimeNil(b bool)`

 SetCiWorkflowExecutionTimeNil sets the value for CiWorkflowExecutionTime to be an explicit nil

### UnsetCiWorkflowExecutionTime
`func (o *SourceImportMetadata) UnsetCiWorkflowExecutionTime()`

UnsetCiWorkflowExecutionTime ensures that no value is present for CiWorkflowExecutionTime, not even an explicit nil
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

### SetBranchNameNil

`func (o *SourceImportMetadata) SetBranchNameNil(b bool)`

 SetBranchNameNil sets the value for BranchName to be an explicit nil

### UnsetBranchName
`func (o *SourceImportMetadata) UnsetBranchName()`

UnsetBranchName ensures that no value is present for BranchName, not even an explicit nil
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

### SetChangeAuthorNil

`func (o *SourceImportMetadata) SetChangeAuthorNil(b bool)`

 SetChangeAuthorNil sets the value for ChangeAuthor to be an explicit nil

### UnsetChangeAuthor
`func (o *SourceImportMetadata) UnsetChangeAuthor()`

UnsetChangeAuthor ensures that no value is present for ChangeAuthor, not even an explicit nil
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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


