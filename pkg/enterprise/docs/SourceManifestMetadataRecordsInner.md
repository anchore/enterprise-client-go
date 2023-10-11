# SourceManifestMetadataRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**CiWorkflowName** | Pointer to **NullableString** |  | [optional] 
**CiWorkflowExecutionTime** | Pointer to **NullableTime** |  | [optional] 
**BranchName** | Pointer to **NullableString** |  | [optional] 
**ChangeAuthor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSourceManifestMetadataRecordsInner

`func NewSourceManifestMetadataRecordsInner() *SourceManifestMetadataRecordsInner`

NewSourceManifestMetadataRecordsInner instantiates a new SourceManifestMetadataRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceManifestMetadataRecordsInnerWithDefaults

`func NewSourceManifestMetadataRecordsInnerWithDefaults() *SourceManifestMetadataRecordsInner`

NewSourceManifestMetadataRecordsInnerWithDefaults instantiates a new SourceManifestMetadataRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SourceManifestMetadataRecordsInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SourceManifestMetadataRecordsInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SourceManifestMetadataRecordsInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SourceManifestMetadataRecordsInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCiWorkflowName

`func (o *SourceManifestMetadataRecordsInner) GetCiWorkflowName() string`

GetCiWorkflowName returns the CiWorkflowName field if non-nil, zero value otherwise.

### GetCiWorkflowNameOk

`func (o *SourceManifestMetadataRecordsInner) GetCiWorkflowNameOk() (*string, bool)`

GetCiWorkflowNameOk returns a tuple with the CiWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiWorkflowName

`func (o *SourceManifestMetadataRecordsInner) SetCiWorkflowName(v string)`

SetCiWorkflowName sets CiWorkflowName field to given value.

### HasCiWorkflowName

`func (o *SourceManifestMetadataRecordsInner) HasCiWorkflowName() bool`

HasCiWorkflowName returns a boolean if a field has been set.

### SetCiWorkflowNameNil

`func (o *SourceManifestMetadataRecordsInner) SetCiWorkflowNameNil(b bool)`

 SetCiWorkflowNameNil sets the value for CiWorkflowName to be an explicit nil

### UnsetCiWorkflowName
`func (o *SourceManifestMetadataRecordsInner) UnsetCiWorkflowName()`

UnsetCiWorkflowName ensures that no value is present for CiWorkflowName, not even an explicit nil
### GetCiWorkflowExecutionTime

`func (o *SourceManifestMetadataRecordsInner) GetCiWorkflowExecutionTime() time.Time`

GetCiWorkflowExecutionTime returns the CiWorkflowExecutionTime field if non-nil, zero value otherwise.

### GetCiWorkflowExecutionTimeOk

`func (o *SourceManifestMetadataRecordsInner) GetCiWorkflowExecutionTimeOk() (*time.Time, bool)`

GetCiWorkflowExecutionTimeOk returns a tuple with the CiWorkflowExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiWorkflowExecutionTime

`func (o *SourceManifestMetadataRecordsInner) SetCiWorkflowExecutionTime(v time.Time)`

SetCiWorkflowExecutionTime sets CiWorkflowExecutionTime field to given value.

### HasCiWorkflowExecutionTime

`func (o *SourceManifestMetadataRecordsInner) HasCiWorkflowExecutionTime() bool`

HasCiWorkflowExecutionTime returns a boolean if a field has been set.

### SetCiWorkflowExecutionTimeNil

`func (o *SourceManifestMetadataRecordsInner) SetCiWorkflowExecutionTimeNil(b bool)`

 SetCiWorkflowExecutionTimeNil sets the value for CiWorkflowExecutionTime to be an explicit nil

### UnsetCiWorkflowExecutionTime
`func (o *SourceManifestMetadataRecordsInner) UnsetCiWorkflowExecutionTime()`

UnsetCiWorkflowExecutionTime ensures that no value is present for CiWorkflowExecutionTime, not even an explicit nil
### GetBranchName

`func (o *SourceManifestMetadataRecordsInner) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *SourceManifestMetadataRecordsInner) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *SourceManifestMetadataRecordsInner) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *SourceManifestMetadataRecordsInner) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### SetBranchNameNil

`func (o *SourceManifestMetadataRecordsInner) SetBranchNameNil(b bool)`

 SetBranchNameNil sets the value for BranchName to be an explicit nil

### UnsetBranchName
`func (o *SourceManifestMetadataRecordsInner) UnsetBranchName()`

UnsetBranchName ensures that no value is present for BranchName, not even an explicit nil
### GetChangeAuthor

`func (o *SourceManifestMetadataRecordsInner) GetChangeAuthor() string`

GetChangeAuthor returns the ChangeAuthor field if non-nil, zero value otherwise.

### GetChangeAuthorOk

`func (o *SourceManifestMetadataRecordsInner) GetChangeAuthorOk() (*string, bool)`

GetChangeAuthorOk returns a tuple with the ChangeAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAuthor

`func (o *SourceManifestMetadataRecordsInner) SetChangeAuthor(v string)`

SetChangeAuthor sets ChangeAuthor field to given value.

### HasChangeAuthor

`func (o *SourceManifestMetadataRecordsInner) HasChangeAuthor() bool`

HasChangeAuthor returns a boolean if a field has been set.

### SetChangeAuthorNil

`func (o *SourceManifestMetadataRecordsInner) SetChangeAuthorNil(b bool)`

 SetChangeAuthorNil sets the value for ChangeAuthor to be an explicit nil

### UnsetChangeAuthor
`func (o *SourceManifestMetadataRecordsInner) UnsetChangeAuthor()`

UnsetChangeAuthor ensures that no value is present for ChangeAuthor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


