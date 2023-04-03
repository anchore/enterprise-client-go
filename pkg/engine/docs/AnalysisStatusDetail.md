# AnalysisStatusDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromStatus** | **string** |  | 
**ToStatus** | **string** |  | 
**Timestamp** | **string** |  | 
**Source** | [**ServiceReference**](ServiceReference.md) |  | 

## Methods

### NewAnalysisStatusDetail

`func NewAnalysisStatusDetail(fromStatus string, toStatus string, timestamp string, source ServiceReference, ) *AnalysisStatusDetail`

NewAnalysisStatusDetail instantiates a new AnalysisStatusDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisStatusDetailWithDefaults

`func NewAnalysisStatusDetailWithDefaults() *AnalysisStatusDetail`

NewAnalysisStatusDetailWithDefaults instantiates a new AnalysisStatusDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromStatus

`func (o *AnalysisStatusDetail) GetFromStatus() string`

GetFromStatus returns the FromStatus field if non-nil, zero value otherwise.

### GetFromStatusOk

`func (o *AnalysisStatusDetail) GetFromStatusOk() (*string, bool)`

GetFromStatusOk returns a tuple with the FromStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromStatus

`func (o *AnalysisStatusDetail) SetFromStatus(v string)`

SetFromStatus sets FromStatus field to given value.


### GetToStatus

`func (o *AnalysisStatusDetail) GetToStatus() string`

GetToStatus returns the ToStatus field if non-nil, zero value otherwise.

### GetToStatusOk

`func (o *AnalysisStatusDetail) GetToStatusOk() (*string, bool)`

GetToStatusOk returns a tuple with the ToStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToStatus

`func (o *AnalysisStatusDetail) SetToStatus(v string)`

SetToStatus sets ToStatus field to given value.


### GetTimestamp

`func (o *AnalysisStatusDetail) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AnalysisStatusDetail) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AnalysisStatusDetail) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetSource

`func (o *AnalysisStatusDetail) GetSource() ServiceReference`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AnalysisStatusDetail) GetSourceOk() (*ServiceReference, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AnalysisStatusDetail) SetSource(v ServiceReference)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


