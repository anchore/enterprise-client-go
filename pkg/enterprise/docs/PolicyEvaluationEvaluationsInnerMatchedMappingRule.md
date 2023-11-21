# PolicyEvaluationEvaluationsInnerMatchedMappingRule

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

### NewPolicyEvaluationEvaluationsInnerMatchedMappingRule

`func NewPolicyEvaluationEvaluationsInnerMatchedMappingRule(id string, name string, allowlistIds []string, ruleSetIds []string, registry string, repository string, image ImageRef, ) *PolicyEvaluationEvaluationsInnerMatchedMappingRule`

NewPolicyEvaluationEvaluationsInnerMatchedMappingRule instantiates a new PolicyEvaluationEvaluationsInnerMatchedMappingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationEvaluationsInnerMatchedMappingRuleWithDefaults

`func NewPolicyEvaluationEvaluationsInnerMatchedMappingRuleWithDefaults() *PolicyEvaluationEvaluationsInnerMatchedMappingRule`

NewPolicyEvaluationEvaluationsInnerMatchedMappingRuleWithDefaults instantiates a new PolicyEvaluationEvaluationsInnerMatchedMappingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) SetName(v string)`

SetName sets Name field to given value.


### GetAllowlistIds

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetAllowlistIds() []string`

GetAllowlistIds returns the AllowlistIds field if non-nil, zero value otherwise.

### GetAllowlistIdsOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetAllowlistIdsOk() (*[]string, bool)`

GetAllowlistIdsOk returns a tuple with the AllowlistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistIds

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) SetAllowlistIds(v []string)`

SetAllowlistIds sets AllowlistIds field to given value.


### GetRuleSetIds

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetRuleSetIds() []string`

GetRuleSetIds returns the RuleSetIds field if non-nil, zero value otherwise.

### GetRuleSetIdsOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetRuleSetIdsOk() (*[]string, bool)`

GetRuleSetIdsOk returns a tuple with the RuleSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSetIds

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) SetRuleSetIds(v []string)`

SetRuleSetIds sets RuleSetIds field to given value.


### GetRegistry

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetRepository

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetImage

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetImage() ImageRef`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetImageOk() (*ImageRef, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) SetImage(v ImageRef)`

SetImage sets Image field to given value.


### GetDescription

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyEvaluationEvaluationsInnerMatchedMappingRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


