# IntegrationReportedStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**IntegrationHealthState**](IntegrationHealthState.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewIntegrationReportedStatus

`func NewIntegrationReportedStatus() *IntegrationReportedStatus`

NewIntegrationReportedStatus instantiates a new IntegrationReportedStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationReportedStatusWithDefaults

`func NewIntegrationReportedStatusWithDefaults() *IntegrationReportedStatus`

NewIntegrationReportedStatusWithDefaults instantiates a new IntegrationReportedStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *IntegrationReportedStatus) GetState() IntegrationHealthState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IntegrationReportedStatus) GetStateOk() (*IntegrationHealthState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IntegrationReportedStatus) SetState(v IntegrationHealthState)`

SetState sets State field to given value.

### HasState

`func (o *IntegrationReportedStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *IntegrationReportedStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IntegrationReportedStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IntegrationReportedStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IntegrationReportedStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDetails

`func (o *IntegrationReportedStatus) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *IntegrationReportedStatus) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *IntegrationReportedStatus) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *IntegrationReportedStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


