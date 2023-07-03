# PolicyEvaluationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**EvaluationId** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**VcsHost** | Pointer to **string** |  | [optional] 
**RepositoryName** | Pointer to **string** |  | [optional] 
**FinalAction** | Pointer to **string** |  | [optional] 
**EvaluationUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Result** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPolicyEvaluationResult

`func NewPolicyEvaluationResult() *PolicyEvaluationResult`

NewPolicyEvaluationResult instantiates a new PolicyEvaluationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationResultWithDefaults

`func NewPolicyEvaluationResultWithDefaults() *PolicyEvaluationResult`

NewPolicyEvaluationResultWithDefaults instantiates a new PolicyEvaluationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *PolicyEvaluationResult) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PolicyEvaluationResult) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PolicyEvaluationResult) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PolicyEvaluationResult) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetPolicyId

`func (o *PolicyEvaluationResult) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyEvaluationResult) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyEvaluationResult) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyEvaluationResult) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetEvaluationId

`func (o *PolicyEvaluationResult) GetEvaluationId() string`

GetEvaluationId returns the EvaluationId field if non-nil, zero value otherwise.

### GetEvaluationIdOk

`func (o *PolicyEvaluationResult) GetEvaluationIdOk() (*string, bool)`

GetEvaluationIdOk returns a tuple with the EvaluationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationId

`func (o *PolicyEvaluationResult) SetEvaluationId(v string)`

SetEvaluationId sets EvaluationId field to given value.

### HasEvaluationId

`func (o *PolicyEvaluationResult) HasEvaluationId() bool`

HasEvaluationId returns a boolean if a field has been set.

### GetSourceId

`func (o *PolicyEvaluationResult) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *PolicyEvaluationResult) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *PolicyEvaluationResult) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *PolicyEvaluationResult) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetVcsHost

`func (o *PolicyEvaluationResult) GetVcsHost() string`

GetVcsHost returns the VcsHost field if non-nil, zero value otherwise.

### GetVcsHostOk

`func (o *PolicyEvaluationResult) GetVcsHostOk() (*string, bool)`

GetVcsHostOk returns a tuple with the VcsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsHost

`func (o *PolicyEvaluationResult) SetVcsHost(v string)`

SetVcsHost sets VcsHost field to given value.

### HasVcsHost

`func (o *PolicyEvaluationResult) HasVcsHost() bool`

HasVcsHost returns a boolean if a field has been set.

### GetRepositoryName

`func (o *PolicyEvaluationResult) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *PolicyEvaluationResult) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *PolicyEvaluationResult) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.

### HasRepositoryName

`func (o *PolicyEvaluationResult) HasRepositoryName() bool`

HasRepositoryName returns a boolean if a field has been set.

### GetFinalAction

`func (o *PolicyEvaluationResult) GetFinalAction() string`

GetFinalAction returns the FinalAction field if non-nil, zero value otherwise.

### GetFinalActionOk

`func (o *PolicyEvaluationResult) GetFinalActionOk() (*string, bool)`

GetFinalActionOk returns a tuple with the FinalAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalAction

`func (o *PolicyEvaluationResult) SetFinalAction(v string)`

SetFinalAction sets FinalAction field to given value.

### HasFinalAction

`func (o *PolicyEvaluationResult) HasFinalAction() bool`

HasFinalAction returns a boolean if a field has been set.

### GetEvaluationUrl

`func (o *PolicyEvaluationResult) GetEvaluationUrl() string`

GetEvaluationUrl returns the EvaluationUrl field if non-nil, zero value otherwise.

### GetEvaluationUrlOk

`func (o *PolicyEvaluationResult) GetEvaluationUrlOk() (*string, bool)`

GetEvaluationUrlOk returns a tuple with the EvaluationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationUrl

`func (o *PolicyEvaluationResult) SetEvaluationUrl(v string)`

SetEvaluationUrl sets EvaluationUrl field to given value.

### HasEvaluationUrl

`func (o *PolicyEvaluationResult) HasEvaluationUrl() bool`

HasEvaluationUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PolicyEvaluationResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyEvaluationResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyEvaluationResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PolicyEvaluationResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PolicyEvaluationResult) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PolicyEvaluationResult) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PolicyEvaluationResult) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PolicyEvaluationResult) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetResult

`func (o *PolicyEvaluationResult) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PolicyEvaluationResult) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PolicyEvaluationResult) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *PolicyEvaluationResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


