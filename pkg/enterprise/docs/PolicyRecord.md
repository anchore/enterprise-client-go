# PolicyRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**PolicyId** | Pointer to **string** | The policy&#39;s identifier | [optional] 
**Active** | Pointer to **bool** | True if the policy is currently defined to be used automatically | [optional] 
**AccountName** | Pointer to **string** | UserId of the user that owns the policy | [optional] 
**PolicySource** | Pointer to **string** | Source location of where the policy originated | [optional] 
**Policy** | Pointer to [**Policy**](Policy.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the policy | [optional] 
**Description** | Pointer to **string** | Description of the policy, human readable | [optional] 

## Methods

### NewPolicyRecord

`func NewPolicyRecord() *PolicyRecord`

NewPolicyRecord instantiates a new PolicyRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRecordWithDefaults

`func NewPolicyRecordWithDefaults() *PolicyRecord`

NewPolicyRecordWithDefaults instantiates a new PolicyRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PolicyRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PolicyRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PolicyRecord) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PolicyRecord) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PolicyRecord) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PolicyRecord) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPolicyId

`func (o *PolicyRecord) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyRecord) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyRecord) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyRecord) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetActive

`func (o *PolicyRecord) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PolicyRecord) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PolicyRecord) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PolicyRecord) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAccountName

`func (o *PolicyRecord) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PolicyRecord) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PolicyRecord) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PolicyRecord) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetPolicySource

`func (o *PolicyRecord) GetPolicySource() string`

GetPolicySource returns the PolicySource field if non-nil, zero value otherwise.

### GetPolicySourceOk

`func (o *PolicyRecord) GetPolicySourceOk() (*string, bool)`

GetPolicySourceOk returns a tuple with the PolicySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySource

`func (o *PolicyRecord) SetPolicySource(v string)`

SetPolicySource sets PolicySource field to given value.

### HasPolicySource

`func (o *PolicyRecord) HasPolicySource() bool`

HasPolicySource returns a boolean if a field has been set.

### GetPolicy

`func (o *PolicyRecord) GetPolicy() Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicyRecord) GetPolicyOk() (*Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicyRecord) SetPolicy(v Policy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PolicyRecord) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetName

`func (o *PolicyRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


