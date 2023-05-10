# PolicyBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the bundle | 
**Name** | Pointer to **string** | Human readable name for the bundle | [optional] 
**Description** | Pointer to **string** | Description of the bundle, human readable | [optional] 
**Version** | **string** | Version id for this bundle format | 
**Allowlists** | Pointer to [**[]Allowlist**](Allowlist.md) | Allowlists which define which policy matches to disregard explicitly in the final policy decision | [optional] 
**Policies** | [**[]Policy**](Policy.md) | Policies which define the go/stop/warn status of an image using rule matches on image properties | 
**SourceMappings** | Pointer to [**[]SourceMappingRule**](SourceMappingRule.md) | Mapping rules for defining which policy and allowlist(s) to apply to a source based on a match of the host and repo name. Evaluated in order. | [optional] 
**Mappings** | [**[]MappingRule**](MappingRule.md) | Mapping rules for defining which policy and allowlist(s) to apply to an image based on a match of the image tag or id. Evaluated in order. | 
**AllowlistedImages** | Pointer to [**[]ImageSelectionRule**](ImageSelectionRule.md) | List of mapping rules that define which images should always be passed (unless also on the denylist), regardless of policy result. | [optional] 
**DenylistedImages** | Pointer to [**[]ImageSelectionRule**](ImageSelectionRule.md) | List of mapping rules that define which images should always result in a STOP/FAIL policy result regardless of policy content or presence in allowlisted_images | [optional] 
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


### GetAllowlists

`func (o *PolicyBundle) GetAllowlists() []Allowlist`

GetAllowlists returns the Allowlists field if non-nil, zero value otherwise.

### GetAllowlistsOk

`func (o *PolicyBundle) GetAllowlistsOk() (*[]Allowlist, bool)`

GetAllowlistsOk returns a tuple with the Allowlists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlists

`func (o *PolicyBundle) SetAllowlists(v []Allowlist)`

SetAllowlists sets Allowlists field to given value.

### HasAllowlists

`func (o *PolicyBundle) HasAllowlists() bool`

HasAllowlists returns a boolean if a field has been set.

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


### GetAllowlistedImages

`func (o *PolicyBundle) GetAllowlistedImages() []ImageSelectionRule`

GetAllowlistedImages returns the AllowlistedImages field if non-nil, zero value otherwise.

### GetAllowlistedImagesOk

`func (o *PolicyBundle) GetAllowlistedImagesOk() (*[]ImageSelectionRule, bool)`

GetAllowlistedImagesOk returns a tuple with the AllowlistedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistedImages

`func (o *PolicyBundle) SetAllowlistedImages(v []ImageSelectionRule)`

SetAllowlistedImages sets AllowlistedImages field to given value.

### HasAllowlistedImages

`func (o *PolicyBundle) HasAllowlistedImages() bool`

HasAllowlistedImages returns a boolean if a field has been set.

### GetDenylistedImages

`func (o *PolicyBundle) GetDenylistedImages() []ImageSelectionRule`

GetDenylistedImages returns the DenylistedImages field if non-nil, zero value otherwise.

### GetDenylistedImagesOk

`func (o *PolicyBundle) GetDenylistedImagesOk() (*[]ImageSelectionRule, bool)`

GetDenylistedImagesOk returns a tuple with the DenylistedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenylistedImages

`func (o *PolicyBundle) SetDenylistedImages(v []ImageSelectionRule)`

SetDenylistedImages sets DenylistedImages field to given value.

### HasDenylistedImages

`func (o *PolicyBundle) HasDenylistedImages() bool`

HasDenylistedImages returns a boolean if a field has been set.

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


