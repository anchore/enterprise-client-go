# SystemStatisticsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TwelveMonths** | Pointer to [**SystemStatisticsDetailTwelveMonths**](SystemStatisticsDetailTwelveMonths.md) |  | [optional] 
**ThirtyDays** | Pointer to [**SystemStatisticsDetailTwelveMonths**](SystemStatisticsDetailTwelveMonths.md) |  | [optional] 
**CalendarYear** | Pointer to **float32** |  | [optional] 
**CalendarMonth** | Pointer to **float32** |  | [optional] 
**TrackedSince** | **string** |  | 
**MonthlyTwelveMonths** | [**[]SystemStatisticsDetailMonthlyTwelveMonthsInner**](SystemStatisticsDetailMonthlyTwelveMonthsInner.md) |  | 
**DailyThirtyDays** | [**[]SystemStatisticsDetailMonthlyTwelveMonthsInner**](SystemStatisticsDetailMonthlyTwelveMonthsInner.md) |  | 

## Methods

### NewSystemStatisticsDetail

`func NewSystemStatisticsDetail(trackedSince string, monthlyTwelveMonths []SystemStatisticsDetailMonthlyTwelveMonthsInner, dailyThirtyDays []SystemStatisticsDetailMonthlyTwelveMonthsInner, ) *SystemStatisticsDetail`

NewSystemStatisticsDetail instantiates a new SystemStatisticsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemStatisticsDetailWithDefaults

`func NewSystemStatisticsDetailWithDefaults() *SystemStatisticsDetail`

NewSystemStatisticsDetailWithDefaults instantiates a new SystemStatisticsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTwelveMonths

`func (o *SystemStatisticsDetail) GetTwelveMonths() SystemStatisticsDetailTwelveMonths`

GetTwelveMonths returns the TwelveMonths field if non-nil, zero value otherwise.

### GetTwelveMonthsOk

`func (o *SystemStatisticsDetail) GetTwelveMonthsOk() (*SystemStatisticsDetailTwelveMonths, bool)`

GetTwelveMonthsOk returns a tuple with the TwelveMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwelveMonths

`func (o *SystemStatisticsDetail) SetTwelveMonths(v SystemStatisticsDetailTwelveMonths)`

SetTwelveMonths sets TwelveMonths field to given value.

### HasTwelveMonths

`func (o *SystemStatisticsDetail) HasTwelveMonths() bool`

HasTwelveMonths returns a boolean if a field has been set.

### GetThirtyDays

`func (o *SystemStatisticsDetail) GetThirtyDays() SystemStatisticsDetailTwelveMonths`

GetThirtyDays returns the ThirtyDays field if non-nil, zero value otherwise.

### GetThirtyDaysOk

`func (o *SystemStatisticsDetail) GetThirtyDaysOk() (*SystemStatisticsDetailTwelveMonths, bool)`

GetThirtyDaysOk returns a tuple with the ThirtyDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirtyDays

`func (o *SystemStatisticsDetail) SetThirtyDays(v SystemStatisticsDetailTwelveMonths)`

SetThirtyDays sets ThirtyDays field to given value.

### HasThirtyDays

`func (o *SystemStatisticsDetail) HasThirtyDays() bool`

HasThirtyDays returns a boolean if a field has been set.

### GetCalendarYear

`func (o *SystemStatisticsDetail) GetCalendarYear() float32`

GetCalendarYear returns the CalendarYear field if non-nil, zero value otherwise.

### GetCalendarYearOk

`func (o *SystemStatisticsDetail) GetCalendarYearOk() (*float32, bool)`

GetCalendarYearOk returns a tuple with the CalendarYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarYear

`func (o *SystemStatisticsDetail) SetCalendarYear(v float32)`

SetCalendarYear sets CalendarYear field to given value.

### HasCalendarYear

`func (o *SystemStatisticsDetail) HasCalendarYear() bool`

HasCalendarYear returns a boolean if a field has been set.

### GetCalendarMonth

`func (o *SystemStatisticsDetail) GetCalendarMonth() float32`

GetCalendarMonth returns the CalendarMonth field if non-nil, zero value otherwise.

### GetCalendarMonthOk

`func (o *SystemStatisticsDetail) GetCalendarMonthOk() (*float32, bool)`

GetCalendarMonthOk returns a tuple with the CalendarMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarMonth

`func (o *SystemStatisticsDetail) SetCalendarMonth(v float32)`

SetCalendarMonth sets CalendarMonth field to given value.

### HasCalendarMonth

`func (o *SystemStatisticsDetail) HasCalendarMonth() bool`

HasCalendarMonth returns a boolean if a field has been set.

### GetTrackedSince

`func (o *SystemStatisticsDetail) GetTrackedSince() string`

GetTrackedSince returns the TrackedSince field if non-nil, zero value otherwise.

### GetTrackedSinceOk

`func (o *SystemStatisticsDetail) GetTrackedSinceOk() (*string, bool)`

GetTrackedSinceOk returns a tuple with the TrackedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedSince

`func (o *SystemStatisticsDetail) SetTrackedSince(v string)`

SetTrackedSince sets TrackedSince field to given value.


### GetMonthlyTwelveMonths

`func (o *SystemStatisticsDetail) GetMonthlyTwelveMonths() []SystemStatisticsDetailMonthlyTwelveMonthsInner`

GetMonthlyTwelveMonths returns the MonthlyTwelveMonths field if non-nil, zero value otherwise.

### GetMonthlyTwelveMonthsOk

`func (o *SystemStatisticsDetail) GetMonthlyTwelveMonthsOk() (*[]SystemStatisticsDetailMonthlyTwelveMonthsInner, bool)`

GetMonthlyTwelveMonthsOk returns a tuple with the MonthlyTwelveMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyTwelveMonths

`func (o *SystemStatisticsDetail) SetMonthlyTwelveMonths(v []SystemStatisticsDetailMonthlyTwelveMonthsInner)`

SetMonthlyTwelveMonths sets MonthlyTwelveMonths field to given value.


### GetDailyThirtyDays

`func (o *SystemStatisticsDetail) GetDailyThirtyDays() []SystemStatisticsDetailMonthlyTwelveMonthsInner`

GetDailyThirtyDays returns the DailyThirtyDays field if non-nil, zero value otherwise.

### GetDailyThirtyDaysOk

`func (o *SystemStatisticsDetail) GetDailyThirtyDaysOk() (*[]SystemStatisticsDetailMonthlyTwelveMonthsInner, bool)`

GetDailyThirtyDaysOk returns a tuple with the DailyThirtyDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyThirtyDays

`func (o *SystemStatisticsDetail) SetDailyThirtyDays(v []SystemStatisticsDetailMonthlyTwelveMonthsInner)`

SetDailyThirtyDays sets DailyThirtyDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


