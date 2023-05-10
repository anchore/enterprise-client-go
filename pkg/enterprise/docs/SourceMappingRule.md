# SourceMappingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**AllowlistIds** | Pointer to **[]string** |  | [optional] 
**PolicyId** | Pointer to **string** | Optional single policy to evaluate, if set will override any value in policy_ids, for backwards compatibility. Generally, policy_ids should be used even with a array of length 1. | [optional] 
**PolicyIds** | Pointer to **[]string** | List of policy_ids to evaluate in order, to completion | [optional] 
**Host** | **string** |  | 
**Repository** | **string** |  | 
**Description** | Pointer to **string** | Description of the source to policy rule, human readable | [optional] 

## Methods

### NewSourceMappingRule

`func NewSourceMappingRule(name string, host string, repository string, ) *SourceMappingRule`

NewSourceMappingRule instantiates a new SourceMappingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceMappingRuleWithDefaults

`func NewSourceMappingRuleWithDefaults() *SourceMappingRule`

NewSourceMappingRuleWithDefaults instantiates a new SourceMappingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceMappingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceMappingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceMappingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SourceMappingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SourceMappingRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceMappingRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceMappingRule) SetName(v string)`

SetName sets Name field to given value.


### GetAllowlistIds

`func (o *SourceMappingRule) GetAllowlistIds() []string`

GetAllowlistIds returns the AllowlistIds field if non-nil, zero value otherwise.

### GetAllowlistIdsOk

`func (o *SourceMappingRule) GetAllowlistIdsOk() (*[]string, bool)`

GetAllowlistIdsOk returns a tuple with the AllowlistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistIds

`func (o *SourceMappingRule) SetAllowlistIds(v []string)`

SetAllowlistIds sets AllowlistIds field to given value.

### HasAllowlistIds

`func (o *SourceMappingRule) HasAllowlistIds() bool`

HasAllowlistIds returns a boolean if a field has been set.

### GetPolicyId

`func (o *SourceMappingRule) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *SourceMappingRule) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *SourceMappingRule) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *SourceMappingRule) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyIds

`func (o *SourceMappingRule) GetPolicyIds() []string`

GetPolicyIds returns the PolicyIds field if non-nil, zero value otherwise.

### GetPolicyIdsOk

`func (o *SourceMappingRule) GetPolicyIdsOk() (*[]string, bool)`

GetPolicyIdsOk returns a tuple with the PolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIds

`func (o *SourceMappingRule) SetPolicyIds(v []string)`

SetPolicyIds sets PolicyIds field to given value.

### HasPolicyIds

`func (o *SourceMappingRule) HasPolicyIds() bool`

HasPolicyIds returns a boolean if a field has been set.

### GetHost

`func (o *SourceMappingRule) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SourceMappingRule) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SourceMappingRule) SetHost(v string)`

SetHost sets Host field to given value.


### GetRepository

`func (o *SourceMappingRule) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *SourceMappingRule) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *SourceMappingRule) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetDescription

`func (o *SourceMappingRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SourceMappingRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SourceMappingRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SourceMappingRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


