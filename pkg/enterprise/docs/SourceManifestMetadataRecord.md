# SourceManifestMetadataRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**CiWorkflowName** | Pointer to **NullableString** |  | [optional] 
**CiWorkflowExecutionTime** | Pointer to **NullableTime** |  | [optional] 
**BranchName** | Pointer to **NullableString** |  | [optional] 
**ChangeAuthor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSourceManifestMetadataRecord

`func NewSourceManifestMetadataRecord() *SourceManifestMetadataRecord`

NewSourceManifestMetadataRecord instantiates a new SourceManifestMetadataRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceManifestMetadataRecordWithDefaults

`func NewSourceManifestMetadataRecordWithDefaults() *SourceManifestMetadataRecord`

NewSourceManifestMetadataRecordWithDefaults instantiates a new SourceManifestMetadataRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SourceManifestMetadataRecord) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SourceManifestMetadataRecord) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SourceManifestMetadataRecord) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SourceManifestMetadataRecord) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCiWorkflowName

`func (o *SourceManifestMetadataRecord) GetCiWorkflowName() string`

GetCiWorkflowName returns the CiWorkflowName field if non-nil, zero value otherwise.

### GetCiWorkflowNameOk

`func (o *SourceManifestMetadataRecord) GetCiWorkflowNameOk() (*string, bool)`

GetCiWorkflowNameOk returns a tuple with the CiWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiWorkflowName

`func (o *SourceManifestMetadataRecord) SetCiWorkflowName(v string)`

SetCiWorkflowName sets CiWorkflowName field to given value.

### HasCiWorkflowName

`func (o *SourceManifestMetadataRecord) HasCiWorkflowName() bool`

HasCiWorkflowName returns a boolean if a field has been set.

### SetCiWorkflowNameNil

`func (o *SourceManifestMetadataRecord) SetCiWorkflowNameNil(b bool)`

 SetCiWorkflowNameNil sets the value for CiWorkflowName to be an explicit nil

### UnsetCiWorkflowName
`func (o *SourceManifestMetadataRecord) UnsetCiWorkflowName()`

UnsetCiWorkflowName ensures that no value is present for CiWorkflowName, not even an explicit nil
### GetCiWorkflowExecutionTime

`func (o *SourceManifestMetadataRecord) GetCiWorkflowExecutionTime() time.Time`

GetCiWorkflowExecutionTime returns the CiWorkflowExecutionTime field if non-nil, zero value otherwise.

### GetCiWorkflowExecutionTimeOk

`func (o *SourceManifestMetadataRecord) GetCiWorkflowExecutionTimeOk() (*time.Time, bool)`

GetCiWorkflowExecutionTimeOk returns a tuple with the CiWorkflowExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiWorkflowExecutionTime

`func (o *SourceManifestMetadataRecord) SetCiWorkflowExecutionTime(v time.Time)`

SetCiWorkflowExecutionTime sets CiWorkflowExecutionTime field to given value.

### HasCiWorkflowExecutionTime

`func (o *SourceManifestMetadataRecord) HasCiWorkflowExecutionTime() bool`

HasCiWorkflowExecutionTime returns a boolean if a field has been set.

### SetCiWorkflowExecutionTimeNil

`func (o *SourceManifestMetadataRecord) SetCiWorkflowExecutionTimeNil(b bool)`

 SetCiWorkflowExecutionTimeNil sets the value for CiWorkflowExecutionTime to be an explicit nil

### UnsetCiWorkflowExecutionTime
`func (o *SourceManifestMetadataRecord) UnsetCiWorkflowExecutionTime()`

UnsetCiWorkflowExecutionTime ensures that no value is present for CiWorkflowExecutionTime, not even an explicit nil
### GetBranchName

`func (o *SourceManifestMetadataRecord) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *SourceManifestMetadataRecord) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *SourceManifestMetadataRecord) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *SourceManifestMetadataRecord) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### SetBranchNameNil

`func (o *SourceManifestMetadataRecord) SetBranchNameNil(b bool)`

 SetBranchNameNil sets the value for BranchName to be an explicit nil

### UnsetBranchName
`func (o *SourceManifestMetadataRecord) UnsetBranchName()`

UnsetBranchName ensures that no value is present for BranchName, not even an explicit nil
### GetChangeAuthor

`func (o *SourceManifestMetadataRecord) GetChangeAuthor() string`

GetChangeAuthor returns the ChangeAuthor field if non-nil, zero value otherwise.

### GetChangeAuthorOk

`func (o *SourceManifestMetadataRecord) GetChangeAuthorOk() (*string, bool)`

GetChangeAuthorOk returns a tuple with the ChangeAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAuthor

`func (o *SourceManifestMetadataRecord) SetChangeAuthor(v string)`

SetChangeAuthor sets ChangeAuthor field to given value.

### HasChangeAuthor

`func (o *SourceManifestMetadataRecord) HasChangeAuthor() bool`

HasChangeAuthor returns a boolean if a field has been set.

### SetChangeAuthorNil

`func (o *SourceManifestMetadataRecord) SetChangeAuthorNil(b bool)`

 SetChangeAuthorNil sets the value for ChangeAuthor to be an explicit nil

### UnsetChangeAuthor
`func (o *SourceManifestMetadataRecord) UnsetChangeAuthor()`

UnsetChangeAuthor ensures that no value is present for ChangeAuthor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


