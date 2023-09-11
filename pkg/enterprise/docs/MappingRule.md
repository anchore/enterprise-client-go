# MappingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**AllowlistIds** | **[]string** |  | 
**RuleSetIds** | **[]string** | List of rule_set_ids to evaluate in order, to completion | 
**Registry** | **string** |  | 
**Repository** | **string** |  | 
**Image** | [**ImageRef**](ImageRef.md) |  | 
**Description** | Pointer to **string** | Description of the image to policy mapping rule, human readable | [optional] 

## Methods

### NewMappingRule

`func NewMappingRule(id string, name string, allowlistIds []string, ruleSetIds []string, registry string, repository string, image ImageRef, ) *MappingRule`

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


### GetAllowlistIds

`func (o *MappingRule) GetAllowlistIds() []string`

GetAllowlistIds returns the AllowlistIds field if non-nil, zero value otherwise.

### GetAllowlistIdsOk

`func (o *MappingRule) GetAllowlistIdsOk() (*[]string, bool)`

GetAllowlistIdsOk returns a tuple with the AllowlistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistIds

`func (o *MappingRule) SetAllowlistIds(v []string)`

SetAllowlistIds sets AllowlistIds field to given value.


### GetRuleSetIds

`func (o *MappingRule) GetRuleSetIds() []string`

GetRuleSetIds returns the RuleSetIds field if non-nil, zero value otherwise.

### GetRuleSetIdsOk

`func (o *MappingRule) GetRuleSetIdsOk() (*[]string, bool)`

GetRuleSetIdsOk returns a tuple with the RuleSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSetIds

`func (o *MappingRule) SetRuleSetIds(v []string)`

SetRuleSetIds sets RuleSetIds field to given value.


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


### GetDescription

`func (o *MappingRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MappingRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MappingRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MappingRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


