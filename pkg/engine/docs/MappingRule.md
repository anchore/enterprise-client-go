# MappingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**WhitelistIds** | Pointer to **[]string** |  | [optional] 
**PolicyId** | Pointer to **string** | Optional single policy to evalute, if set will override any value in policy_ids, for backwards compatibility. Generally, policy_ids should be used even with a array of length 1. | [optional] 
**PolicyIds** | Pointer to **[]string** | List of policyIds to evaluate in order, to completion | [optional] 
**Registry** | **string** |  | 
**Repository** | **string** |  | 
**Image** | [**ImageRef**](ImageRef.md) |  | 

## Methods

### NewMappingRule

`func NewMappingRule(name string, registry string, repository string, image ImageRef, ) *MappingRule`

NewMappingRule instantiates a new MappingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingRuleWithDefaults

`func NewMappingRuleWithDefaults() *MappingRule`

NewMappingRuleWithDefaults instantiates a new MappingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MappingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MappingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MappingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MappingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MappingRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MappingRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MappingRule) SetName(v string)`

SetName sets Name field to given value.


### GetWhitelistIds

`func (o *MappingRule) GetWhitelistIds() []string`

GetWhitelistIds returns the WhitelistIds field if non-nil, zero value otherwise.

### GetWhitelistIdsOk

`func (o *MappingRule) GetWhitelistIdsOk() (*[]string, bool)`

GetWhitelistIdsOk returns a tuple with the WhitelistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistIds

`func (o *MappingRule) SetWhitelistIds(v []string)`

SetWhitelistIds sets WhitelistIds field to given value.

### HasWhitelistIds

`func (o *MappingRule) HasWhitelistIds() bool`

HasWhitelistIds returns a boolean if a field has been set.

### GetPolicyId

`func (o *MappingRule) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *MappingRule) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *MappingRule) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *MappingRule) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyIds

`func (o *MappingRule) GetPolicyIds() []string`

GetPolicyIds returns the PolicyIds field if non-nil, zero value otherwise.

### GetPolicyIdsOk

`func (o *MappingRule) GetPolicyIdsOk() (*[]string, bool)`

GetPolicyIdsOk returns a tuple with the PolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIds

`func (o *MappingRule) SetPolicyIds(v []string)`

SetPolicyIds sets PolicyIds field to given value.

### HasPolicyIds

`func (o *MappingRule) HasPolicyIds() bool`

HasPolicyIds returns a boolean if a field has been set.

### GetRegistry

`func (o *MappingRule) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *MappingRule) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *MappingRule) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetRepository

`func (o *MappingRule) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *MappingRule) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *MappingRule) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetImage

`func (o *MappingRule) GetImage() ImageRef`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *MappingRule) GetImageOk() (*ImageRef, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *MappingRule) SetImage(v ImageRef)`

SetImage sets Image field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


