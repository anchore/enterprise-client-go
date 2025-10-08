# AnalysisStatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneratedAt** | Pointer to **time.Time** | The time the statistics were generated | [optional] 
**Global** | Pointer to [**AnalysisStatisticsByType**](AnalysisStatisticsByType.md) |  | [optional] 
**Details** | Pointer to [**[]AnalysisStatisticsAccount**](AnalysisStatisticsAccount.md) |  | [optional] 

## Methods

### NewAnalysisStatisticsResponse

`func NewAnalysisStatisticsResponse() *AnalysisStatisticsResponse`

NewAnalysisStatisticsResponse instantiates a new AnalysisStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisStatisticsResponseWithDefaults

`func NewAnalysisStatisticsResponseWithDefaults() *AnalysisStatisticsResponse`

NewAnalysisStatisticsResponseWithDefaults instantiates a new AnalysisStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratedAt

`func (o *AnalysisStatisticsResponse) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *AnalysisStatisticsResponse) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *AnalysisStatisticsResponse) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *AnalysisStatisticsResponse) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetGlobal

`func (o *AnalysisStatisticsResponse) GetGlobal() AnalysisStatisticsByType`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *AnalysisStatisticsResponse) GetGlobalOk() (*AnalysisStatisticsByType, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *AnalysisStatisticsResponse) SetGlobal(v AnalysisStatisticsByType)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *AnalysisStatisticsResponse) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetDetails

`func (o *AnalysisStatisticsResponse) GetDetails() []AnalysisStatisticsAccount`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AnalysisStatisticsResponse) GetDetailsOk() (*[]AnalysisStatisticsAccount, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AnalysisStatisticsResponse) SetDetails(v []AnalysisStatisticsAccount)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AnalysisStatisticsResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


