<<<<<<< HEAD
# Policy
=======
# RuleSet
>>>>>>> 48fc108 (feat: updated the enterprise ref)

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
<<<<<<< HEAD
**Id** | **string** | Id of the policy | 
**Name** | **string** | Human readable name for the policy | 
**Description** | Pointer to **string** | Description of the policy, human readable | [optional] 
**Version** | **string** | Version id for this policy format | 
**Allowlists** | Pointer to [**[]Allowlist**](Allowlist.md) | Allowlists which define which policy matches to disregard explicitly in the final policy decision | [optional] 
**RuleSets** | [**[]RuleSet**](RuleSet.md) | Collections of policy rules which define the go/stop/warn status of an image using rule matches on image properties | 
**SourceMappings** | Pointer to [**[]SourceMappingRule**](SourceMappingRule.md) | Mapping rules for defining which policy and allowlist(s) to apply to a source based on a match of the host and repo name. Evaluated in order. | [optional] 
**Mappings** | [**[]MappingRule**](MappingRule.md) | Mapping rules for defining which policy and allowlist(s) to apply to an image based on a match of the image tag or id. Evaluated in order. | 
**AllowlistedImages** | Pointer to [**[]ImageSelectionRule**](ImageSelectionRule.md) | List of mapping rules that define which images should always be passed (unless also on the denylist), regardless of policy result. | [optional] 
**DenylistedImages** | Pointer to [**[]ImageSelectionRule**](ImageSelectionRule.md) | List of mapping rules that define which images should always result in a STOP/FAIL policy result regardless of policy content or presence in allowlisted_images | [optional] 
**LastUpdated** | Pointer to **float32** | The time at which the policy was last updated, informational only | [optional] 

## Methods

### NewPolicy

`func NewPolicy(id string, name string, version string, ruleSets []RuleSet, mappings []MappingRule, ) *Policy`

NewPolicy instantiates a new Policy object
=======
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
>>>>>>> 48fc108 (feat: updated the enterprise ref)
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

<<<<<<< HEAD
### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
=======
### NewRuleSetWithDefaults

`func NewRuleSetWithDefaults() *RuleSet`

NewRuleSetWithDefaults instantiates a new RuleSet object
>>>>>>> 48fc108 (feat: updated the enterprise ref)
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

<<<<<<< HEAD
`func (o *Policy) GetId() string`
=======
`func (o *RuleSet) GetId() string`
>>>>>>> 48fc108 (feat: updated the enterprise ref)

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

<<<<<<< HEAD
`func (o *Policy) GetIdOk() (*string, bool)`
=======
`func (o *RuleSet) GetIdOk() (*string, bool)`
>>>>>>> 48fc108 (feat: updated the enterprise ref)

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

<<<<<<< HEAD
`func (o *Policy) SetId(v string)`
=======
`func (o *RuleSet) SetId(v string)`
>>>>>>> 48fc108 (feat: updated the enterprise ref)

SetId sets Id field to given value.


### GetName

<<<<<<< HEAD
`func (o *Policy) GetName() string`
=======
`func (o *RuleSet) GetName() string`
>>>>>>> 48fc108 (feat: updated the enterprise ref)

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

<<<<<<< HEAD
`func (o *Policy) GetNameOk() (*string, bool)`
=======
`func (o *RuleSet) GetNameOk() (*string, bool)`
>>>>>>> 48fc108 (feat: updated the enterprise ref)

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

<<<<<<< HEAD
`func (o *Policy) SetName(v string)`
=======
`func (o *RuleSet) SetName(v string)`
>>>>>>> 48fc108 (feat: updated the enterprise ref)

SetName sets Name field to given value.


### GetDescription

<<<<<<< HEAD
=======
`func (o *RuleSet) GetDescription() string`

<<<<<<<< HEAD:pkg/enterprise/docs/RuleSet.md
GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleSet) GetDescriptionOk() (*string, bool)`

========
### GetDescription

>>>>>>> 48fc108 (feat: updated the enterprise ref)
`func (o *Policy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Policy) GetDescriptionOk() (*string, bool)`

<<<<<<< HEAD
=======
>>>>>>>> 48fc108 (feat: updated the enterprise ref):pkg/enterprise/docs/Policy.md
>>>>>>> 48fc108 (feat: updated the enterprise ref)
GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

<<<<<<< HEAD
`func (o *Policy) SetDescription(v string)`
=======
<<<<<<<< HEAD:pkg/enterprise/docs/RuleSet.md
`func (o *RuleSet) SetDescription(v string)`
========
`func (o *Policy) SetDescription(v string)`
>>>>>>>> 48fc108 (feat: updated the enterprise ref):pkg/enterprise/docs/Policy.md
>>>>>>> 48fc108 (feat: updated the enterprise ref)

SetDescription sets Description field to given value.

### HasDescription

<<<<<<< HEAD
`func (o *Policy) HasDescription() bool`
=======
<<<<<<<< HEAD:pkg/enterprise/docs/RuleSet.md
`func (o *RuleSet) HasDescription() bool`
========
`func (o *Policy) HasDescription() bool`
>>>>>>>> 48fc108 (feat: updated the enterprise ref):pkg/enterprise/docs/Policy.md
>>>>>>> 48fc108 (feat: updated the enterprise ref)

HasDescription returns a boolean if a field has been set.

### GetVersion

<<<<<<< HEAD
`func (o *Policy) GetVersion() string`
=======
`func (o *RuleSet) GetVersion() string`
>>>>>>> 48fc108 (feat: updated the enterprise ref)

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

<<<<<<< HEAD
`func (o *Policy) GetVersionOk() (*string, bool)`
=======
`func (o *RuleSet) GetVersionOk() (*string, bool)`
>>>>>>> 48fc108 (feat: updated the enterprise ref)

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

<<<<<<< HEAD
`func (o *Policy) SetVersion(v string)`
=======
`func (o *RuleSet) SetVersion(v string)`
>>>>>>> 48fc108 (feat: updated the enterprise ref)

SetVersion sets Version field to given value.


<<<<<<< HEAD
### GetAllowlists

`func (o *Policy) GetAllowlists() []Allowlist`

GetAllowlists returns the Allowlists field if non-nil, zero value otherwise.

### GetAllowlistsOk

`func (o *Policy) GetAllowlistsOk() (*[]Allowlist, bool)`

GetAllowlistsOk returns a tuple with the Allowlists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlists

`func (o *Policy) SetAllowlists(v []Allowlist)`

SetAllowlists sets Allowlists field to given value.

### HasAllowlists

`func (o *Policy) HasAllowlists() bool`

HasAllowlists returns a boolean if a field has been set.

### GetRuleSets

`func (o *Policy) GetRuleSets() []RuleSet`

GetRuleSets returns the RuleSets field if non-nil, zero value otherwise.

### GetRuleSetsOk

`func (o *Policy) GetRuleSetsOk() (*[]RuleSet, bool)`

GetRuleSetsOk returns a tuple with the RuleSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSets

`func (o *Policy) SetRuleSets(v []RuleSet)`

SetRuleSets sets RuleSets field to given value.


### GetSourceMappings

`func (o *Policy) GetSourceMappings() []SourceMappingRule`

GetSourceMappings returns the SourceMappings field if non-nil, zero value otherwise.

### GetSourceMappingsOk

`func (o *Policy) GetSourceMappingsOk() (*[]SourceMappingRule, bool)`

GetSourceMappingsOk returns a tuple with the SourceMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMappings

`func (o *Policy) SetSourceMappings(v []SourceMappingRule)`

SetSourceMappings sets SourceMappings field to given value.

### HasSourceMappings

`func (o *Policy) HasSourceMappings() bool`

HasSourceMappings returns a boolean if a field has been set.

### GetMappings

`func (o *Policy) GetMappings() []MappingRule`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *Policy) GetMappingsOk() (*[]MappingRule, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *Policy) SetMappings(v []MappingRule)`

SetMappings sets Mappings field to given value.


### GetAllowlistedImages

`func (o *Policy) GetAllowlistedImages() []ImageSelectionRule`

GetAllowlistedImages returns the AllowlistedImages field if non-nil, zero value otherwise.

### GetAllowlistedImagesOk

`func (o *Policy) GetAllowlistedImagesOk() (*[]ImageSelectionRule, bool)`

GetAllowlistedImagesOk returns a tuple with the AllowlistedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistedImages

`func (o *Policy) SetAllowlistedImages(v []ImageSelectionRule)`

SetAllowlistedImages sets AllowlistedImages field to given value.

### HasAllowlistedImages

`func (o *Policy) HasAllowlistedImages() bool`

HasAllowlistedImages returns a boolean if a field has been set.

### GetDenylistedImages

`func (o *Policy) GetDenylistedImages() []ImageSelectionRule`

GetDenylistedImages returns the DenylistedImages field if non-nil, zero value otherwise.

### GetDenylistedImagesOk

`func (o *Policy) GetDenylistedImagesOk() (*[]ImageSelectionRule, bool)`

GetDenylistedImagesOk returns a tuple with the DenylistedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenylistedImages

`func (o *Policy) SetDenylistedImages(v []ImageSelectionRule)`

SetDenylistedImages sets DenylistedImages field to given value.

### HasDenylistedImages

`func (o *Policy) HasDenylistedImages() bool`

HasDenylistedImages returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Policy) GetLastUpdated() float32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Policy) GetLastUpdatedOk() (*float32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Policy) SetLastUpdated(v float32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Policy) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

=======
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


>>>>>>> 48fc108 (feat: updated the enterprise ref)

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


