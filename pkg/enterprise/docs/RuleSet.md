# RuleSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
<<<<<<<< HEAD:pkg/enterprise/docs/RuleSet.md
**Name** | **string** |  | 
========
**Name** | Pointer to **string** |  | [optional] 
>>>>>>>> 48fc108 (feat: updated the enterprise ref):pkg/enterprise/docs/Policy.md
**Description** | Pointer to **string** | Description of the Policy, human readable | [optional] 
**Version** | **string** |  | 
**ArtifactType** | Pointer to **string** |  | [optional] 
**Rules** | [**[]PolicyRule**](PolicyRule.md) |  | 

## Methods

### NewRuleSet

`func NewRuleSet(id string, name string, version string, rules []PolicyRule, ) *RuleSet`

NewRuleSet instantiates a new RuleSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSetWithDefaults

`func NewRuleSetWithDefaults() *RuleSet`

NewRuleSetWithDefaults instantiates a new RuleSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleSet) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RuleSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleSet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RuleSet) GetDescription() string`

<<<<<<<< HEAD:pkg/enterprise/docs/RuleSet.md
GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleSet) GetDescriptionOk() (*string, bool)`

========
### GetDescription

`func (o *Policy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Policy) GetDescriptionOk() (*string, bool)`

>>>>>>>> 48fc108 (feat: updated the enterprise ref):pkg/enterprise/docs/Policy.md
GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

<<<<<<<< HEAD:pkg/enterprise/docs/RuleSet.md
`func (o *RuleSet) SetDescription(v string)`
========
`func (o *Policy) SetDescription(v string)`
>>>>>>>> 48fc108 (feat: updated the enterprise ref):pkg/enterprise/docs/Policy.md

SetDescription sets Description field to given value.

### HasDescription

<<<<<<<< HEAD:pkg/enterprise/docs/RuleSet.md
`func (o *RuleSet) HasDescription() bool`
========
`func (o *Policy) HasDescription() bool`
>>>>>>>> 48fc108 (feat: updated the enterprise ref):pkg/enterprise/docs/Policy.md

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *RuleSet) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RuleSet) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RuleSet) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetArtifactType

`func (o *RuleSet) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *RuleSet) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *RuleSet) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *RuleSet) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetRules

`func (o *RuleSet) GetRules() []PolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *RuleSet) GetRulesOk() (*[]PolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *RuleSet) SetRules(v []PolicyRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


