# HealthReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Universally unique identifier for this health report | 
**ProtocolVersion** | **int64** | The version of the health report exchange protocol | 
**Timestamp** | **time.Time** | timestamp for this health report | 
**Uptime** | **float32** | Running time of integration instance in seconds | 
**HealthReportInterval** | **int64** | Interval (in seconds) between health reports | 
**HealthData** | [**HealthData**](HealthData.md) |  | 

## Methods

### NewHealthReport

`func NewHealthReport(uuid string, protocolVersion int64, timestamp time.Time, uptime float32, healthReportInterval int64, healthData HealthData, ) *HealthReport`

NewHealthReport instantiates a new HealthReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthReportWithDefaults

`func NewHealthReportWithDefaults() *HealthReport`

NewHealthReportWithDefaults instantiates a new HealthReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *HealthReport) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HealthReport) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HealthReport) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetProtocolVersion

`func (o *HealthReport) GetProtocolVersion() int64`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *HealthReport) GetProtocolVersionOk() (*int64, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *HealthReport) SetProtocolVersion(v int64)`

SetProtocolVersion sets ProtocolVersion field to given value.


### GetTimestamp

`func (o *HealthReport) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HealthReport) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HealthReport) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetUptime

`func (o *HealthReport) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *HealthReport) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *HealthReport) SetUptime(v float32)`

SetUptime sets Uptime field to given value.


### GetHealthReportInterval

`func (o *HealthReport) GetHealthReportInterval() int64`

GetHealthReportInterval returns the HealthReportInterval field if non-nil, zero value otherwise.

### GetHealthReportIntervalOk

`func (o *HealthReport) GetHealthReportIntervalOk() (*int64, bool)`

GetHealthReportIntervalOk returns a tuple with the HealthReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReportInterval

`func (o *HealthReport) SetHealthReportInterval(v int64)`

SetHealthReportInterval sets HealthReportInterval field to given value.


### GetHealthData

`func (o *HealthReport) GetHealthData() HealthData`

GetHealthData returns the HealthData field if non-nil, zero value otherwise.

### GetHealthDataOk

`func (o *HealthReport) GetHealthDataOk() (*HealthData, bool)`

GetHealthDataOk returns a tuple with the HealthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthData

`func (o *HealthReport) SetHealthData(v HealthData)`

SetHealthData sets HealthData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


