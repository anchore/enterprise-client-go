# SourceManifestMetadataRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**CiWorkflowName** | Pointer to **NullableString** |  | [optional] 
**CiWorkflowExecutionTime** | Pointer to **NullableTime** |  | [optional] 
**BranchName** | Pointer to **NullableString** |  | [optional] 
**ChangeAuthor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSourceManifestMetadataRecords

`func NewSourceManifestMetadataRecords() *SourceManifestMetadataRecords`

NewSourceManifestMetadataRecords instantiates a new SourceManifestMetadataRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceManifestMetadataRecordsWithDefaults

`func NewSourceManifestMetadataRecordsWithDefaults() *SourceManifestMetadataRecords`

NewSourceManifestMetadataRecordsWithDefaults instantiates a new SourceManifestMetadataRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SourceManifestMetadataRecords) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SourceManifestMetadataRecords) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SourceManifestMetadataRecords) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SourceManifestMetadataRecords) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCiWorkflowName

`func (o *SourceManifestMetadataRecords) GetCiWorkflowName() string`

GetCiWorkflowName returns the CiWorkflowName field if non-nil, zero value otherwise.

### GetCiWorkflowNameOk

`func (o *SourceManifestMetadataRecords) GetCiWorkflowNameOk() (*string, bool)`

GetCiWorkflowNameOk returns a tuple with the CiWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiWorkflowName

`func (o *SourceManifestMetadataRecords) SetCiWorkflowName(v string)`

SetCiWorkflowName sets CiWorkflowName field to given value.

### HasCiWorkflowName

`func (o *SourceManifestMetadataRecords) HasCiWorkflowName() bool`

HasCiWorkflowName returns a boolean if a field has been set.

### SetCiWorkflowNameNil

`func (o *SourceManifestMetadataRecords) SetCiWorkflowNameNil(b bool)`

 SetCiWorkflowNameNil sets the value for CiWorkflowName to be an explicit nil

### UnsetCiWorkflowName
`func (o *SourceManifestMetadataRecords) UnsetCiWorkflowName()`

UnsetCiWorkflowName ensures that no value is present for CiWorkflowName, not even an explicit nil
### GetCiWorkflowExecutionTime

`func (o *SourceManifestMetadataRecords) GetCiWorkflowExecutionTime() time.Time`

GetCiWorkflowExecutionTime returns the CiWorkflowExecutionTime field if non-nil, zero value otherwise.

### GetCiWorkflowExecutionTimeOk

`func (o *SourceManifestMetadataRecords) GetCiWorkflowExecutionTimeOk() (*time.Time, bool)`

GetCiWorkflowExecutionTimeOk returns a tuple with the CiWorkflowExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiWorkflowExecutionTime

`func (o *SourceManifestMetadataRecords) SetCiWorkflowExecutionTime(v time.Time)`

SetCiWorkflowExecutionTime sets CiWorkflowExecutionTime field to given value.

### HasCiWorkflowExecutionTime

`func (o *SourceManifestMetadataRecords) HasCiWorkflowExecutionTime() bool`

HasCiWorkflowExecutionTime returns a boolean if a field has been set.

### SetCiWorkflowExecutionTimeNil

`func (o *SourceManifestMetadataRecords) SetCiWorkflowExecutionTimeNil(b bool)`

 SetCiWorkflowExecutionTimeNil sets the value for CiWorkflowExecutionTime to be an explicit nil

### UnsetCiWorkflowExecutionTime
`func (o *SourceManifestMetadataRecords) UnsetCiWorkflowExecutionTime()`

UnsetCiWorkflowExecutionTime ensures that no value is present for CiWorkflowExecutionTime, not even an explicit nil
### GetBranchName

`func (o *SourceManifestMetadataRecords) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *SourceManifestMetadataRecords) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *SourceManifestMetadataRecords) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *SourceManifestMetadataRecords) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### SetBranchNameNil

`func (o *SourceManifestMetadataRecords) SetBranchNameNil(b bool)`

 SetBranchNameNil sets the value for BranchName to be an explicit nil

### UnsetBranchName
`func (o *SourceManifestMetadataRecords) UnsetBranchName()`

UnsetBranchName ensures that no value is present for BranchName, not even an explicit nil
### GetChangeAuthor

`func (o *SourceManifestMetadataRecords) GetChangeAuthor() string`

GetChangeAuthor returns the ChangeAuthor field if non-nil, zero value otherwise.

### GetChangeAuthorOk

`func (o *SourceManifestMetadataRecords) GetChangeAuthorOk() (*string, bool)`

GetChangeAuthorOk returns a tuple with the ChangeAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAuthor

`func (o *SourceManifestMetadataRecords) SetChangeAuthor(v string)`

SetChangeAuthor sets ChangeAuthor field to given value.

### HasChangeAuthor

`func (o *SourceManifestMetadataRecords) HasChangeAuthor() bool`

HasChangeAuthor returns a boolean if a field has been set.

### SetChangeAuthorNil

`func (o *SourceManifestMetadataRecords) SetChangeAuthorNil(b bool)`

 SetChangeAuthorNil sets the value for ChangeAuthor to be an explicit nil

### UnsetChangeAuthor
`func (o *SourceManifestMetadataRecords) UnsetChangeAuthor()`

UnsetChangeAuthor ensures that no value is present for ChangeAuthor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


