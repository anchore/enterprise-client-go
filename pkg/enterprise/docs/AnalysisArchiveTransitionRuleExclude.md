# AnalysisArchiveTransitionRuleExclude

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selector** | Pointer to [**ImageSelector**](ImageSelector.md) |  | [optional] 
**ExpirationDays** | Pointer to **int32** | How long the image selected will be excluded from the archive transition | [optional] 
**LastSeenInDays** | Pointer to **int32** | Exclude image from archive if last seen in inventory within defined number of days | [optional] 

## Methods

### NewAnalysisArchiveTransitionRuleExclude

`func NewAnalysisArchiveTransitionRuleExclude() *AnalysisArchiveTransitionRuleExclude`

NewAnalysisArchiveTransitionRuleExclude instantiates a new AnalysisArchiveTransitionRuleExclude object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisArchiveTransitionRuleExcludeWithDefaults

`func NewAnalysisArchiveTransitionRuleExcludeWithDefaults() *AnalysisArchiveTransitionRuleExclude`

NewAnalysisArchiveTransitionRuleExcludeWithDefaults instantiates a new AnalysisArchiveTransitionRuleExclude object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelector

`func (o *AnalysisArchiveTransitionRuleExclude) GetSelector() ImageSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *AnalysisArchiveTransitionRuleExclude) GetSelectorOk() (*ImageSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *AnalysisArchiveTransitionRuleExclude) SetSelector(v ImageSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *AnalysisArchiveTransitionRuleExclude) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetExpirationDays

`func (o *AnalysisArchiveTransitionRuleExclude) GetExpirationDays() int32`

GetExpirationDays returns the ExpirationDays field if non-nil, zero value otherwise.

### GetExpirationDaysOk

`func (o *AnalysisArchiveTransitionRuleExclude) GetExpirationDaysOk() (*int32, bool)`

GetExpirationDaysOk returns a tuple with the ExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDays

`func (o *AnalysisArchiveTransitionRuleExclude) SetExpirationDays(v int32)`

SetExpirationDays sets ExpirationDays field to given value.

### HasExpirationDays

`func (o *AnalysisArchiveTransitionRuleExclude) HasExpirationDays() bool`

HasExpirationDays returns a boolean if a field has been set.

### GetLastSeenInDays

`func (o *AnalysisArchiveTransitionRuleExclude) GetLastSeenInDays() int32`

GetLastSeenInDays returns the LastSeenInDays field if non-nil, zero value otherwise.

### GetLastSeenInDaysOk

`func (o *AnalysisArchiveTransitionRuleExclude) GetLastSeenInDaysOk() (*int32, bool)`

GetLastSeenInDaysOk returns a tuple with the LastSeenInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenInDays

`func (o *AnalysisArchiveTransitionRuleExclude) SetLastSeenInDays(v int32)`

SetLastSeenInDays sets LastSeenInDays field to given value.

### HasLastSeenInDays

`func (o *AnalysisArchiveTransitionRuleExclude) HasLastSeenInDays() bool`

HasLastSeenInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


