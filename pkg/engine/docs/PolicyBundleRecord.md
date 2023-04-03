# PolicyBundleRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**PolicyId** | Pointer to **string** | The bundle&#39;s identifier | [optional] 
**Active** | Pointer to **bool** | True if the bundle is currently defined to be used automatically | [optional] 
**UserId** | Pointer to **string** | UserId of the user that owns the bundle | [optional] 
**PolicySource** | Pointer to **string** | Source location of where the policy bundle originated | [optional] 
**Policybundle** | Pointer to [**PolicyBundle**](PolicyBundle.md) |  | [optional] 
**Policybundlemeta** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** | Name of the policy bundle | [optional] 
**Description** | Pointer to **string** | Description of the bundle, human readable | [optional] 

## Methods

### NewPolicyBundleRecord

`func NewPolicyBundleRecord() *PolicyBundleRecord`

NewPolicyBundleRecord instantiates a new PolicyBundleRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyBundleRecordWithDefaults

`func NewPolicyBundleRecordWithDefaults() *PolicyBundleRecord`

NewPolicyBundleRecordWithDefaults instantiates a new PolicyBundleRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PolicyBundleRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyBundleRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyBundleRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PolicyBundleRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PolicyBundleRecord) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PolicyBundleRecord) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PolicyBundleRecord) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PolicyBundleRecord) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPolicyId

`func (o *PolicyBundleRecord) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyBundleRecord) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyBundleRecord) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyBundleRecord) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetActive

`func (o *PolicyBundleRecord) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PolicyBundleRecord) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PolicyBundleRecord) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PolicyBundleRecord) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUserId

`func (o *PolicyBundleRecord) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PolicyBundleRecord) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PolicyBundleRecord) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PolicyBundleRecord) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPolicySource

`func (o *PolicyBundleRecord) GetPolicySource() string`

GetPolicySource returns the PolicySource field if non-nil, zero value otherwise.

### GetPolicySourceOk

`func (o *PolicyBundleRecord) GetPolicySourceOk() (*string, bool)`

GetPolicySourceOk returns a tuple with the PolicySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySource

`func (o *PolicyBundleRecord) SetPolicySource(v string)`

SetPolicySource sets PolicySource field to given value.

### HasPolicySource

`func (o *PolicyBundleRecord) HasPolicySource() bool`

HasPolicySource returns a boolean if a field has been set.

### GetPolicybundle

`func (o *PolicyBundleRecord) GetPolicybundle() PolicyBundle`

GetPolicybundle returns the Policybundle field if non-nil, zero value otherwise.

### GetPolicybundleOk

`func (o *PolicyBundleRecord) GetPolicybundleOk() (*PolicyBundle, bool)`

GetPolicybundleOk returns a tuple with the Policybundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicybundle

`func (o *PolicyBundleRecord) SetPolicybundle(v PolicyBundle)`

SetPolicybundle sets Policybundle field to given value.

### HasPolicybundle

`func (o *PolicyBundleRecord) HasPolicybundle() bool`

HasPolicybundle returns a boolean if a field has been set.

### GetPolicybundlemeta

`func (o *PolicyBundleRecord) GetPolicybundlemeta() interface{}`

GetPolicybundlemeta returns the Policybundlemeta field if non-nil, zero value otherwise.

### GetPolicybundlemetaOk

`func (o *PolicyBundleRecord) GetPolicybundlemetaOk() (*interface{}, bool)`

GetPolicybundlemetaOk returns a tuple with the Policybundlemeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicybundlemeta

`func (o *PolicyBundleRecord) SetPolicybundlemeta(v interface{})`

SetPolicybundlemeta sets Policybundlemeta field to given value.

### HasPolicybundlemeta

`func (o *PolicyBundleRecord) HasPolicybundlemeta() bool`

HasPolicybundlemeta returns a boolean if a field has been set.

### GetName

`func (o *PolicyBundleRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyBundleRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyBundleRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyBundleRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyBundleRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyBundleRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyBundleRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyBundleRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


