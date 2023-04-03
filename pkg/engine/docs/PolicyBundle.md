# PolicyBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the bundle | 
**Name** | Pointer to **string** | Human readable name for the bundle | [optional] 
**Comment** | Pointer to **string** | Description of the bundle, human readable (Deprecated, use the description field instead) | [optional] 
**Description** | Pointer to **string** | Description of the bundle, human readable | [optional] 
**Version** | **string** | Version id for this bundle format | 
**Whitelists** | Pointer to [**[]Whitelist**](Whitelist.md) | Whitelists which define which policy matches to disregard explicitly in the final policy decision | [optional] 
**Policies** | [**[]Policy**](Policy.md) | Policies which define the go/stop/warn status of an image using rule matches on image properties | 
**SourceMappings** | Pointer to [**[]SourceMappingRule**](SourceMappingRule.md) | Mapping rules for defining which policy and whitelist(s) to apply to a source based on a match of the host and repo name. Evaluated in order. | [optional] 
**Mappings** | [**[]MappingRule**](MappingRule.md) | Mapping rules for defining which policy and whitelist(s) to apply to an image based on a match of the image tag or id. Evaluated in order. | 
**WhitelistedImages** | Pointer to [**[]ImageSelectionRule**](ImageSelectionRule.md) | List of mapping rules that define which images should always be passed (unless also on the blacklist), regardless of policy result. | [optional] 
**BlacklistedImages** | Pointer to [**[]ImageSelectionRule**](ImageSelectionRule.md) | List of mapping rules that define which images should always result in a STOP/FAIL policy result regardless of policy content or presence in whitelisted_images | [optional] 
**LastUpdated** | Pointer to **float32** | The time at which the policy was last updated, informational only | [optional] 

## Methods

### NewPolicyBundle

`func NewPolicyBundle(id string, version string, policies []Policy, mappings []MappingRule, ) *PolicyBundle`

NewPolicyBundle instantiates a new PolicyBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyBundleWithDefaults

`func NewPolicyBundleWithDefaults() *PolicyBundle`

NewPolicyBundleWithDefaults instantiates a new PolicyBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyBundle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyBundle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyBundle) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PolicyBundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyBundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyBundle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyBundle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *PolicyBundle) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PolicyBundle) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PolicyBundle) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PolicyBundle) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyBundle) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyBundle) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyBundle) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyBundle) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *PolicyBundle) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PolicyBundle) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PolicyBundle) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetWhitelists

`func (o *PolicyBundle) GetWhitelists() []Whitelist`

GetWhitelists returns the Whitelists field if non-nil, zero value otherwise.

### GetWhitelistsOk

`func (o *PolicyBundle) GetWhitelistsOk() (*[]Whitelist, bool)`

GetWhitelistsOk returns a tuple with the Whitelists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelists

`func (o *PolicyBundle) SetWhitelists(v []Whitelist)`

SetWhitelists sets Whitelists field to given value.

### HasWhitelists

`func (o *PolicyBundle) HasWhitelists() bool`

HasWhitelists returns a boolean if a field has been set.

### GetPolicies

`func (o *PolicyBundle) GetPolicies() []Policy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *PolicyBundle) GetPoliciesOk() (*[]Policy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *PolicyBundle) SetPolicies(v []Policy)`

SetPolicies sets Policies field to given value.


### GetSourceMappings

`func (o *PolicyBundle) GetSourceMappings() []SourceMappingRule`

GetSourceMappings returns the SourceMappings field if non-nil, zero value otherwise.

### GetSourceMappingsOk

`func (o *PolicyBundle) GetSourceMappingsOk() (*[]SourceMappingRule, bool)`

GetSourceMappingsOk returns a tuple with the SourceMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMappings

`func (o *PolicyBundle) SetSourceMappings(v []SourceMappingRule)`

SetSourceMappings sets SourceMappings field to given value.

### HasSourceMappings

`func (o *PolicyBundle) HasSourceMappings() bool`

HasSourceMappings returns a boolean if a field has been set.

### GetMappings

`func (o *PolicyBundle) GetMappings() []MappingRule`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *PolicyBundle) GetMappingsOk() (*[]MappingRule, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *PolicyBundle) SetMappings(v []MappingRule)`

SetMappings sets Mappings field to given value.


### GetWhitelistedImages

`func (o *PolicyBundle) GetWhitelistedImages() []ImageSelectionRule`

GetWhitelistedImages returns the WhitelistedImages field if non-nil, zero value otherwise.

### GetWhitelistedImagesOk

`func (o *PolicyBundle) GetWhitelistedImagesOk() (*[]ImageSelectionRule, bool)`

GetWhitelistedImagesOk returns a tuple with the WhitelistedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedImages

`func (o *PolicyBundle) SetWhitelistedImages(v []ImageSelectionRule)`

SetWhitelistedImages sets WhitelistedImages field to given value.

### HasWhitelistedImages

`func (o *PolicyBundle) HasWhitelistedImages() bool`

HasWhitelistedImages returns a boolean if a field has been set.

### GetBlacklistedImages

`func (o *PolicyBundle) GetBlacklistedImages() []ImageSelectionRule`

GetBlacklistedImages returns the BlacklistedImages field if non-nil, zero value otherwise.

### GetBlacklistedImagesOk

`func (o *PolicyBundle) GetBlacklistedImagesOk() (*[]ImageSelectionRule, bool)`

GetBlacklistedImagesOk returns a tuple with the BlacklistedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedImages

`func (o *PolicyBundle) SetBlacklistedImages(v []ImageSelectionRule)`

SetBlacklistedImages sets BlacklistedImages field to given value.

### HasBlacklistedImages

`func (o *PolicyBundle) HasBlacklistedImages() bool`

HasBlacklistedImages returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PolicyBundle) GetLastUpdated() float32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PolicyBundle) GetLastUpdatedOk() (*float32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PolicyBundle) SetLastUpdated(v float32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PolicyBundle) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


