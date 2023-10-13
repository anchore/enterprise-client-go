# PolicyRecordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewPolicyRecordPolicy

`func NewPolicyRecordPolicy(id string, name string, version string, ruleSets []RuleSet, mappings []MappingRule, ) *PolicyRecordPolicy`

NewPolicyRecordPolicy instantiates a new PolicyRecordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRecordPolicyWithDefaults

`func NewPolicyRecordPolicyWithDefaults() *PolicyRecordPolicy`

NewPolicyRecordPolicyWithDefaults instantiates a new PolicyRecordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyRecordPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyRecordPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyRecordPolicy) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PolicyRecordPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRecordPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRecordPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PolicyRecordPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyRecordPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyRecordPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyRecordPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *PolicyRecordPolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PolicyRecordPolicy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PolicyRecordPolicy) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetAllowlists

`func (o *PolicyRecordPolicy) GetAllowlists() []Allowlist`

GetAllowlists returns the Allowlists field if non-nil, zero value otherwise.

### GetAllowlistsOk

`func (o *PolicyRecordPolicy) GetAllowlistsOk() (*[]Allowlist, bool)`

GetAllowlistsOk returns a tuple with the Allowlists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlists

`func (o *PolicyRecordPolicy) SetAllowlists(v []Allowlist)`

SetAllowlists sets Allowlists field to given value.

### HasAllowlists

`func (o *PolicyRecordPolicy) HasAllowlists() bool`

HasAllowlists returns a boolean if a field has been set.

### GetRuleSets

`func (o *PolicyRecordPolicy) GetRuleSets() []RuleSet`

GetRuleSets returns the RuleSets field if non-nil, zero value otherwise.

### GetRuleSetsOk

`func (o *PolicyRecordPolicy) GetRuleSetsOk() (*[]RuleSet, bool)`

GetRuleSetsOk returns a tuple with the RuleSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSets

`func (o *PolicyRecordPolicy) SetRuleSets(v []RuleSet)`

SetRuleSets sets RuleSets field to given value.


### GetSourceMappings

`func (o *PolicyRecordPolicy) GetSourceMappings() []SourceMappingRule`

GetSourceMappings returns the SourceMappings field if non-nil, zero value otherwise.

### GetSourceMappingsOk

`func (o *PolicyRecordPolicy) GetSourceMappingsOk() (*[]SourceMappingRule, bool)`

GetSourceMappingsOk returns a tuple with the SourceMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMappings

`func (o *PolicyRecordPolicy) SetSourceMappings(v []SourceMappingRule)`

SetSourceMappings sets SourceMappings field to given value.

### HasSourceMappings

`func (o *PolicyRecordPolicy) HasSourceMappings() bool`

HasSourceMappings returns a boolean if a field has been set.

### GetMappings

`func (o *PolicyRecordPolicy) GetMappings() []MappingRule`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *PolicyRecordPolicy) GetMappingsOk() (*[]MappingRule, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *PolicyRecordPolicy) SetMappings(v []MappingRule)`

SetMappings sets Mappings field to given value.


### GetAllowlistedImages

`func (o *PolicyRecordPolicy) GetAllowlistedImages() []ImageSelectionRule`

GetAllowlistedImages returns the AllowlistedImages field if non-nil, zero value otherwise.

### GetAllowlistedImagesOk

`func (o *PolicyRecordPolicy) GetAllowlistedImagesOk() (*[]ImageSelectionRule, bool)`

GetAllowlistedImagesOk returns a tuple with the AllowlistedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistedImages

`func (o *PolicyRecordPolicy) SetAllowlistedImages(v []ImageSelectionRule)`

SetAllowlistedImages sets AllowlistedImages field to given value.

### HasAllowlistedImages

`func (o *PolicyRecordPolicy) HasAllowlistedImages() bool`

HasAllowlistedImages returns a boolean if a field has been set.

### GetDenylistedImages

`func (o *PolicyRecordPolicy) GetDenylistedImages() []ImageSelectionRule`

GetDenylistedImages returns the DenylistedImages field if non-nil, zero value otherwise.

### GetDenylistedImagesOk

`func (o *PolicyRecordPolicy) GetDenylistedImagesOk() (*[]ImageSelectionRule, bool)`

GetDenylistedImagesOk returns a tuple with the DenylistedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenylistedImages

`func (o *PolicyRecordPolicy) SetDenylistedImages(v []ImageSelectionRule)`

SetDenylistedImages sets DenylistedImages field to given value.

### HasDenylistedImages

`func (o *PolicyRecordPolicy) HasDenylistedImages() bool`

HasDenylistedImages returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PolicyRecordPolicy) GetLastUpdated() float32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PolicyRecordPolicy) GetLastUpdatedOk() (*float32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PolicyRecordPolicy) SetLastUpdated(v float32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PolicyRecordPolicy) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


