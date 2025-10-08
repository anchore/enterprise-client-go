# SystemStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**MetricBasis** | Pointer to **string** |  | [optional] 
**ValueType** | **string** |  | 
**Value** | **float32** |  | 
**Detail** | Pointer to [**SystemStatisticsDetail**](SystemStatisticsDetail.md) |  | [optional] 

## Methods

### NewSystemStatistics

`func NewSystemStatistics(name string, description string, valueType string, value float32, ) *SystemStatistics`

NewSystemStatistics instantiates a new SystemStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemStatisticsWithDefaults

`func NewSystemStatisticsWithDefaults() *SystemStatistics`

NewSystemStatisticsWithDefaults instantiates a new SystemStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SystemStatistics) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemStatistics) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemStatistics) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SystemStatistics) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemStatistics) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemStatistics) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMetricBasis

`func (o *SystemStatistics) GetMetricBasis() string`

GetMetricBasis returns the MetricBasis field if non-nil, zero value otherwise.

### GetMetricBasisOk

`func (o *SystemStatistics) GetMetricBasisOk() (*string, bool)`

GetMetricBasisOk returns a tuple with the MetricBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricBasis

`func (o *SystemStatistics) SetMetricBasis(v string)`

SetMetricBasis sets MetricBasis field to given value.

### HasMetricBasis

`func (o *SystemStatistics) HasMetricBasis() bool`

HasMetricBasis returns a boolean if a field has been set.

### GetValueType

`func (o *SystemStatistics) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *SystemStatistics) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *SystemStatistics) SetValueType(v string)`

SetValueType sets ValueType field to given value.


### GetValue

`func (o *SystemStatistics) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SystemStatistics) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SystemStatistics) SetValue(v float32)`

SetValue sets Value field to given value.


### GetDetail

`func (o *SystemStatistics) GetDetail() SystemStatisticsDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SystemStatistics) GetDetailOk() (*SystemStatisticsDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SystemStatistics) SetDetail(v SystemStatisticsDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SystemStatistics) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


