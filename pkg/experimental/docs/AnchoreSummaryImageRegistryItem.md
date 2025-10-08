# AnchoreSummaryImageRegistryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repo** | Pointer to **string** |  | [optional] 
**Analyzed** | Pointer to **int32** |  | [optional] 
**NotAnalyzed** | Pointer to **int32** |  | [optional] 
**Analyzing** | Pointer to **int32** |  | [optional] 
**AnalysisFailed** | Pointer to **int32** |  | [optional] 
**LatestAnalyzedAt** | Pointer to **time.Time** |  | [optional] 
**Alert** | Pointer to **bool** |  | [optional] 
**AlertId** | Pointer to **string** |  | [optional] 
**WatchRepo** | Pointer to **bool** |  | [optional] 
**WatchRepoId** | Pointer to **string** |  | [optional] 

## Methods

### NewAnchoreSummaryImageRegistryItem

`func NewAnchoreSummaryImageRegistryItem() *AnchoreSummaryImageRegistryItem`

NewAnchoreSummaryImageRegistryItem instantiates a new AnchoreSummaryImageRegistryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnchoreSummaryImageRegistryItemWithDefaults

`func NewAnchoreSummaryImageRegistryItemWithDefaults() *AnchoreSummaryImageRegistryItem`

NewAnchoreSummaryImageRegistryItemWithDefaults instantiates a new AnchoreSummaryImageRegistryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepo

`func (o *AnchoreSummaryImageRegistryItem) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AnchoreSummaryImageRegistryItem) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AnchoreSummaryImageRegistryItem) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AnchoreSummaryImageRegistryItem) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetAnalyzed

`func (o *AnchoreSummaryImageRegistryItem) GetAnalyzed() int32`

GetAnalyzed returns the Analyzed field if non-nil, zero value otherwise.

### GetAnalyzedOk

`func (o *AnchoreSummaryImageRegistryItem) GetAnalyzedOk() (*int32, bool)`

GetAnalyzedOk returns a tuple with the Analyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzed

`func (o *AnchoreSummaryImageRegistryItem) SetAnalyzed(v int32)`

SetAnalyzed sets Analyzed field to given value.

### HasAnalyzed

`func (o *AnchoreSummaryImageRegistryItem) HasAnalyzed() bool`

HasAnalyzed returns a boolean if a field has been set.

### GetNotAnalyzed

`func (o *AnchoreSummaryImageRegistryItem) GetNotAnalyzed() int32`

GetNotAnalyzed returns the NotAnalyzed field if non-nil, zero value otherwise.

### GetNotAnalyzedOk

`func (o *AnchoreSummaryImageRegistryItem) GetNotAnalyzedOk() (*int32, bool)`

GetNotAnalyzedOk returns a tuple with the NotAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAnalyzed

`func (o *AnchoreSummaryImageRegistryItem) SetNotAnalyzed(v int32)`

SetNotAnalyzed sets NotAnalyzed field to given value.

### HasNotAnalyzed

`func (o *AnchoreSummaryImageRegistryItem) HasNotAnalyzed() bool`

HasNotAnalyzed returns a boolean if a field has been set.

### GetAnalyzing

`func (o *AnchoreSummaryImageRegistryItem) GetAnalyzing() int32`

GetAnalyzing returns the Analyzing field if non-nil, zero value otherwise.

### GetAnalyzingOk

`func (o *AnchoreSummaryImageRegistryItem) GetAnalyzingOk() (*int32, bool)`

GetAnalyzingOk returns a tuple with the Analyzing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzing

`func (o *AnchoreSummaryImageRegistryItem) SetAnalyzing(v int32)`

SetAnalyzing sets Analyzing field to given value.

### HasAnalyzing

`func (o *AnchoreSummaryImageRegistryItem) HasAnalyzing() bool`

HasAnalyzing returns a boolean if a field has been set.

### GetAnalysisFailed

`func (o *AnchoreSummaryImageRegistryItem) GetAnalysisFailed() int32`

GetAnalysisFailed returns the AnalysisFailed field if non-nil, zero value otherwise.

### GetAnalysisFailedOk

`func (o *AnchoreSummaryImageRegistryItem) GetAnalysisFailedOk() (*int32, bool)`

GetAnalysisFailedOk returns a tuple with the AnalysisFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisFailed

`func (o *AnchoreSummaryImageRegistryItem) SetAnalysisFailed(v int32)`

SetAnalysisFailed sets AnalysisFailed field to given value.

### HasAnalysisFailed

`func (o *AnchoreSummaryImageRegistryItem) HasAnalysisFailed() bool`

HasAnalysisFailed returns a boolean if a field has been set.

### GetLatestAnalyzedAt

`func (o *AnchoreSummaryImageRegistryItem) GetLatestAnalyzedAt() time.Time`

GetLatestAnalyzedAt returns the LatestAnalyzedAt field if non-nil, zero value otherwise.

### GetLatestAnalyzedAtOk

`func (o *AnchoreSummaryImageRegistryItem) GetLatestAnalyzedAtOk() (*time.Time, bool)`

GetLatestAnalyzedAtOk returns a tuple with the LatestAnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAnalyzedAt

`func (o *AnchoreSummaryImageRegistryItem) SetLatestAnalyzedAt(v time.Time)`

SetLatestAnalyzedAt sets LatestAnalyzedAt field to given value.

### HasLatestAnalyzedAt

`func (o *AnchoreSummaryImageRegistryItem) HasLatestAnalyzedAt() bool`

HasLatestAnalyzedAt returns a boolean if a field has been set.

### GetAlert

`func (o *AnchoreSummaryImageRegistryItem) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *AnchoreSummaryImageRegistryItem) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *AnchoreSummaryImageRegistryItem) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *AnchoreSummaryImageRegistryItem) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetAlertId

`func (o *AnchoreSummaryImageRegistryItem) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *AnchoreSummaryImageRegistryItem) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *AnchoreSummaryImageRegistryItem) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *AnchoreSummaryImageRegistryItem) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetWatchRepo

`func (o *AnchoreSummaryImageRegistryItem) GetWatchRepo() bool`

GetWatchRepo returns the WatchRepo field if non-nil, zero value otherwise.

### GetWatchRepoOk

`func (o *AnchoreSummaryImageRegistryItem) GetWatchRepoOk() (*bool, bool)`

GetWatchRepoOk returns a tuple with the WatchRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchRepo

`func (o *AnchoreSummaryImageRegistryItem) SetWatchRepo(v bool)`

SetWatchRepo sets WatchRepo field to given value.

### HasWatchRepo

`func (o *AnchoreSummaryImageRegistryItem) HasWatchRepo() bool`

HasWatchRepo returns a boolean if a field has been set.

### GetWatchRepoId

`func (o *AnchoreSummaryImageRegistryItem) GetWatchRepoId() string`

GetWatchRepoId returns the WatchRepoId field if non-nil, zero value otherwise.

### GetWatchRepoIdOk

`func (o *AnchoreSummaryImageRegistryItem) GetWatchRepoIdOk() (*string, bool)`

GetWatchRepoIdOk returns a tuple with the WatchRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchRepoId

`func (o *AnchoreSummaryImageRegistryItem) SetWatchRepoId(v string)`

SetWatchRepoId sets WatchRepoId field to given value.

### HasWatchRepoId

`func (o *AnchoreSummaryImageRegistryItem) HasWatchRepoId() bool`

HasWatchRepoId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


