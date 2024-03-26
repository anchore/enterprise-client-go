# DeploymentHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **string** |  | [optional] 
**UpgradeId** | Pointer to **string** |  | [optional] 
**ToSystemVersion** | Pointer to **string** |  | [optional] 
**FromSystemVersion** | Pointer to **string** |  | [optional] 
**ToDatabaseVersion** | Pointer to **string** |  | [optional] 
**FromDatabaseVersion** | Pointer to **string** |  | [optional] 
**Outcome** | Pointer to **string** |  | [optional] 
**DbUpgradeDuration** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDeploymentHistory

`func NewDeploymentHistory() *DeploymentHistory`

NewDeploymentHistory instantiates a new DeploymentHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryWithDefaults

`func NewDeploymentHistoryWithDefaults() *DeploymentHistory`

NewDeploymentHistoryWithDefaults instantiates a new DeploymentHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *DeploymentHistory) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DeploymentHistory) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DeploymentHistory) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *DeploymentHistory) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetUpgradeId

`func (o *DeploymentHistory) GetUpgradeId() string`

GetUpgradeId returns the UpgradeId field if non-nil, zero value otherwise.

### GetUpgradeIdOk

`func (o *DeploymentHistory) GetUpgradeIdOk() (*string, bool)`

GetUpgradeIdOk returns a tuple with the UpgradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeId

`func (o *DeploymentHistory) SetUpgradeId(v string)`

SetUpgradeId sets UpgradeId field to given value.

### HasUpgradeId

`func (o *DeploymentHistory) HasUpgradeId() bool`

HasUpgradeId returns a boolean if a field has been set.

### GetToSystemVersion

`func (o *DeploymentHistory) GetToSystemVersion() string`

GetToSystemVersion returns the ToSystemVersion field if non-nil, zero value otherwise.

### GetToSystemVersionOk

`func (o *DeploymentHistory) GetToSystemVersionOk() (*string, bool)`

GetToSystemVersionOk returns a tuple with the ToSystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSystemVersion

`func (o *DeploymentHistory) SetToSystemVersion(v string)`

SetToSystemVersion sets ToSystemVersion field to given value.

### HasToSystemVersion

`func (o *DeploymentHistory) HasToSystemVersion() bool`

HasToSystemVersion returns a boolean if a field has been set.

### GetFromSystemVersion

`func (o *DeploymentHistory) GetFromSystemVersion() string`

GetFromSystemVersion returns the FromSystemVersion field if non-nil, zero value otherwise.

### GetFromSystemVersionOk

`func (o *DeploymentHistory) GetFromSystemVersionOk() (*string, bool)`

GetFromSystemVersionOk returns a tuple with the FromSystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSystemVersion

`func (o *DeploymentHistory) SetFromSystemVersion(v string)`

SetFromSystemVersion sets FromSystemVersion field to given value.

### HasFromSystemVersion

`func (o *DeploymentHistory) HasFromSystemVersion() bool`

HasFromSystemVersion returns a boolean if a field has been set.

### GetToDatabaseVersion

`func (o *DeploymentHistory) GetToDatabaseVersion() string`

GetToDatabaseVersion returns the ToDatabaseVersion field if non-nil, zero value otherwise.

### GetToDatabaseVersionOk

`func (o *DeploymentHistory) GetToDatabaseVersionOk() (*string, bool)`

GetToDatabaseVersionOk returns a tuple with the ToDatabaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDatabaseVersion

`func (o *DeploymentHistory) SetToDatabaseVersion(v string)`

SetToDatabaseVersion sets ToDatabaseVersion field to given value.

### HasToDatabaseVersion

`func (o *DeploymentHistory) HasToDatabaseVersion() bool`

HasToDatabaseVersion returns a boolean if a field has been set.

### GetFromDatabaseVersion

`func (o *DeploymentHistory) GetFromDatabaseVersion() string`

GetFromDatabaseVersion returns the FromDatabaseVersion field if non-nil, zero value otherwise.

### GetFromDatabaseVersionOk

`func (o *DeploymentHistory) GetFromDatabaseVersionOk() (*string, bool)`

GetFromDatabaseVersionOk returns a tuple with the FromDatabaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDatabaseVersion

`func (o *DeploymentHistory) SetFromDatabaseVersion(v string)`

SetFromDatabaseVersion sets FromDatabaseVersion field to given value.

### HasFromDatabaseVersion

`func (o *DeploymentHistory) HasFromDatabaseVersion() bool`

HasFromDatabaseVersion returns a boolean if a field has been set.

### GetOutcome

`func (o *DeploymentHistory) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *DeploymentHistory) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *DeploymentHistory) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *DeploymentHistory) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetDbUpgradeDuration

`func (o *DeploymentHistory) GetDbUpgradeDuration() float32`

GetDbUpgradeDuration returns the DbUpgradeDuration field if non-nil, zero value otherwise.

### GetDbUpgradeDurationOk

`func (o *DeploymentHistory) GetDbUpgradeDurationOk() (*float32, bool)`

GetDbUpgradeDurationOk returns a tuple with the DbUpgradeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUpgradeDuration

`func (o *DeploymentHistory) SetDbUpgradeDuration(v float32)`

SetDbUpgradeDuration sets DbUpgradeDuration field to given value.

### HasDbUpgradeDuration

`func (o *DeploymentHistory) HasDbUpgradeDuration() bool`

HasDbUpgradeDuration returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeploymentHistory) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentHistory) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentHistory) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeploymentHistory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


