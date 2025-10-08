# AnalysisStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumAnalyzed** | Pointer to **int32** | The number of images in the \&quot;analyzed\&quot; state | [optional] 
**NumAnalyzing** | Pointer to **int32** | The number of images in the \&quot;analyzing\&quot; state | [optional] 
**NumNotAnalyzed** | Pointer to **int32** | The number of images in the \&quot;not_analyzed\&quot; state | [optional] 
**NumFailed** | Pointer to **int32** | The number of images in the \&quot;failed\&quot; state | [optional] 

## Methods

### NewAnalysisStatistics

`func NewAnalysisStatistics() *AnalysisStatistics`

NewAnalysisStatistics instantiates a new AnalysisStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisStatisticsWithDefaults

`func NewAnalysisStatisticsWithDefaults() *AnalysisStatistics`

NewAnalysisStatisticsWithDefaults instantiates a new AnalysisStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumAnalyzed

`func (o *AnalysisStatistics) GetNumAnalyzed() int32`

GetNumAnalyzed returns the NumAnalyzed field if non-nil, zero value otherwise.

### GetNumAnalyzedOk

`func (o *AnalysisStatistics) GetNumAnalyzedOk() (*int32, bool)`

GetNumAnalyzedOk returns a tuple with the NumAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAnalyzed

`func (o *AnalysisStatistics) SetNumAnalyzed(v int32)`

SetNumAnalyzed sets NumAnalyzed field to given value.

### HasNumAnalyzed

`func (o *AnalysisStatistics) HasNumAnalyzed() bool`

HasNumAnalyzed returns a boolean if a field has been set.

### GetNumAnalyzing

`func (o *AnalysisStatistics) GetNumAnalyzing() int32`

GetNumAnalyzing returns the NumAnalyzing field if non-nil, zero value otherwise.

### GetNumAnalyzingOk

`func (o *AnalysisStatistics) GetNumAnalyzingOk() (*int32, bool)`

GetNumAnalyzingOk returns a tuple with the NumAnalyzing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAnalyzing

`func (o *AnalysisStatistics) SetNumAnalyzing(v int32)`

SetNumAnalyzing sets NumAnalyzing field to given value.

### HasNumAnalyzing

`func (o *AnalysisStatistics) HasNumAnalyzing() bool`

HasNumAnalyzing returns a boolean if a field has been set.

### GetNumNotAnalyzed

`func (o *AnalysisStatistics) GetNumNotAnalyzed() int32`

GetNumNotAnalyzed returns the NumNotAnalyzed field if non-nil, zero value otherwise.

### GetNumNotAnalyzedOk

`func (o *AnalysisStatistics) GetNumNotAnalyzedOk() (*int32, bool)`

GetNumNotAnalyzedOk returns a tuple with the NumNotAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNotAnalyzed

`func (o *AnalysisStatistics) SetNumNotAnalyzed(v int32)`

SetNumNotAnalyzed sets NumNotAnalyzed field to given value.

### HasNumNotAnalyzed

`func (o *AnalysisStatistics) HasNumNotAnalyzed() bool`

HasNumNotAnalyzed returns a boolean if a field has been set.

### GetNumFailed

`func (o *AnalysisStatistics) GetNumFailed() int32`

GetNumFailed returns the NumFailed field if non-nil, zero value otherwise.

### GetNumFailedOk

`func (o *AnalysisStatistics) GetNumFailedOk() (*int32, bool)`

GetNumFailedOk returns a tuple with the NumFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailed

`func (o *AnalysisStatistics) SetNumFailed(v int32)`

SetNumFailed sets NumFailed field to given value.

### HasNumFailed

`func (o *AnalysisStatistics) HasNumFailed() bool`

HasNumFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


