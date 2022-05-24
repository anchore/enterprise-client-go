# AnalysisArchiveRulesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of rules for this account | [optional] 
**LastUpdated** | Pointer to **time.Time** | The newest last_updated timestamp from the set of rules | [optional] 

## Methods

### NewAnalysisArchiveRulesSummary

`func NewAnalysisArchiveRulesSummary() *AnalysisArchiveRulesSummary`

NewAnalysisArchiveRulesSummary instantiates a new AnalysisArchiveRulesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisArchiveRulesSummaryWithDefaults

`func NewAnalysisArchiveRulesSummaryWithDefaults() *AnalysisArchiveRulesSummary`

NewAnalysisArchiveRulesSummaryWithDefaults instantiates a new AnalysisArchiveRulesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AnalysisArchiveRulesSummary) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AnalysisArchiveRulesSummary) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AnalysisArchiveRulesSummary) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AnalysisArchiveRulesSummary) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AnalysisArchiveRulesSummary) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AnalysisArchiveRulesSummary) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AnalysisArchiveRulesSummary) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AnalysisArchiveRulesSummary) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


