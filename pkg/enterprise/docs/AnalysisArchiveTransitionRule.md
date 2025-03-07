# AnalysisArchiveTransitionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selector** | Pointer to [**ImageSelector**](ImageSelector.md) |  | [optional] 
**RuleId** | Pointer to **string** | Unique identifier for archive rule | [optional] 
**TagVersionsNewer** | Pointer to **int64** | Number of images mapped to the tag that are newer | [optional] 
**AnalysisAgeDays** | Pointer to **int64** | Matches if the analysis is strictly older than this number of days | [optional] 
**Transition** | **string** | The type of transition to make. If \&quot;archive\&quot;, then archive an image from the working set and remove it from the working set. If \&quot;delete\&quot;, then match against archived images and delete from the archive if match. | 
**SystemGlobal** | Pointer to **bool** | True if the rule applies to all accounts in the system. This is only available to admin users to update/modify, but all users with permission to list rules can see them | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Exclude** | Pointer to [**AnalysisArchiveTransitionRuleExclude**](AnalysisArchiveTransitionRuleExclude.md) |  | [optional] 
**MaxImagesPerAccount** | Pointer to **int64** | This is the maximum number of image analyses an account can have. Can only be set on system_global rules | [optional] 

## Methods

### NewAnalysisArchiveTransitionRule

`func NewAnalysisArchiveTransitionRule(transition string, ) *AnalysisArchiveTransitionRule`

NewAnalysisArchiveTransitionRule instantiates a new AnalysisArchiveTransitionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisArchiveTransitionRuleWithDefaults

`func NewAnalysisArchiveTransitionRuleWithDefaults() *AnalysisArchiveTransitionRule`

NewAnalysisArchiveTransitionRuleWithDefaults instantiates a new AnalysisArchiveTransitionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelector

`func (o *AnalysisArchiveTransitionRule) GetSelector() ImageSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *AnalysisArchiveTransitionRule) GetSelectorOk() (*ImageSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *AnalysisArchiveTransitionRule) SetSelector(v ImageSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *AnalysisArchiveTransitionRule) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetRuleId

`func (o *AnalysisArchiveTransitionRule) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *AnalysisArchiveTransitionRule) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *AnalysisArchiveTransitionRule) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *AnalysisArchiveTransitionRule) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetTagVersionsNewer

`func (o *AnalysisArchiveTransitionRule) GetTagVersionsNewer() int64`

GetTagVersionsNewer returns the TagVersionsNewer field if non-nil, zero value otherwise.

### GetTagVersionsNewerOk

`func (o *AnalysisArchiveTransitionRule) GetTagVersionsNewerOk() (*int64, bool)`

GetTagVersionsNewerOk returns a tuple with the TagVersionsNewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagVersionsNewer

`func (o *AnalysisArchiveTransitionRule) SetTagVersionsNewer(v int64)`

SetTagVersionsNewer sets TagVersionsNewer field to given value.

### HasTagVersionsNewer

`func (o *AnalysisArchiveTransitionRule) HasTagVersionsNewer() bool`

HasTagVersionsNewer returns a boolean if a field has been set.

### GetAnalysisAgeDays

`func (o *AnalysisArchiveTransitionRule) GetAnalysisAgeDays() int64`

GetAnalysisAgeDays returns the AnalysisAgeDays field if non-nil, zero value otherwise.

### GetAnalysisAgeDaysOk

`func (o *AnalysisArchiveTransitionRule) GetAnalysisAgeDaysOk() (*int64, bool)`

GetAnalysisAgeDaysOk returns a tuple with the AnalysisAgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisAgeDays

`func (o *AnalysisArchiveTransitionRule) SetAnalysisAgeDays(v int64)`

SetAnalysisAgeDays sets AnalysisAgeDays field to given value.

### HasAnalysisAgeDays

`func (o *AnalysisArchiveTransitionRule) HasAnalysisAgeDays() bool`

HasAnalysisAgeDays returns a boolean if a field has been set.

### GetTransition

`func (o *AnalysisArchiveTransitionRule) GetTransition() string`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *AnalysisArchiveTransitionRule) GetTransitionOk() (*string, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *AnalysisArchiveTransitionRule) SetTransition(v string)`

SetTransition sets Transition field to given value.


### GetSystemGlobal

`func (o *AnalysisArchiveTransitionRule) GetSystemGlobal() bool`

GetSystemGlobal returns the SystemGlobal field if non-nil, zero value otherwise.

### GetSystemGlobalOk

`func (o *AnalysisArchiveTransitionRule) GetSystemGlobalOk() (*bool, bool)`

GetSystemGlobalOk returns a tuple with the SystemGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemGlobal

`func (o *AnalysisArchiveTransitionRule) SetSystemGlobal(v bool)`

SetSystemGlobal sets SystemGlobal field to given value.

### HasSystemGlobal

`func (o *AnalysisArchiveTransitionRule) HasSystemGlobal() bool`

HasSystemGlobal returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AnalysisArchiveTransitionRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnalysisArchiveTransitionRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnalysisArchiveTransitionRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AnalysisArchiveTransitionRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AnalysisArchiveTransitionRule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AnalysisArchiveTransitionRule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AnalysisArchiveTransitionRule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AnalysisArchiveTransitionRule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetExclude

`func (o *AnalysisArchiveTransitionRule) GetExclude() AnalysisArchiveTransitionRuleExclude`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *AnalysisArchiveTransitionRule) GetExcludeOk() (*AnalysisArchiveTransitionRuleExclude, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *AnalysisArchiveTransitionRule) SetExclude(v AnalysisArchiveTransitionRuleExclude)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *AnalysisArchiveTransitionRule) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetMaxImagesPerAccount

`func (o *AnalysisArchiveTransitionRule) GetMaxImagesPerAccount() int64`

GetMaxImagesPerAccount returns the MaxImagesPerAccount field if non-nil, zero value otherwise.

### GetMaxImagesPerAccountOk

`func (o *AnalysisArchiveTransitionRule) GetMaxImagesPerAccountOk() (*int64, bool)`

GetMaxImagesPerAccountOk returns a tuple with the MaxImagesPerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxImagesPerAccount

`func (o *AnalysisArchiveTransitionRule) SetMaxImagesPerAccount(v int64)`

SetMaxImagesPerAccount sets MaxImagesPerAccount field to given value.

### HasMaxImagesPerAccount

`func (o *AnalysisArchiveTransitionRule) HasMaxImagesPerAccount() bool`

HasMaxImagesPerAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


