# IntegrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**IntegrationLifecycleState**](IntegrationLifecycleState.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp when status was updated | [optional] 

## Methods

### NewIntegrationStatus

`func NewIntegrationStatus() *IntegrationStatus`

NewIntegrationStatus instantiates a new IntegrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationStatusWithDefaults

`func NewIntegrationStatusWithDefaults() *IntegrationStatus`

NewIntegrationStatusWithDefaults instantiates a new IntegrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *IntegrationStatus) GetState() IntegrationLifecycleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IntegrationStatus) GetStateOk() (*IntegrationLifecycleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IntegrationStatus) SetState(v IntegrationLifecycleState)`

SetState sets State field to given value.

### HasState

`func (o *IntegrationStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *IntegrationStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IntegrationStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IntegrationStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IntegrationStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDetails

`func (o *IntegrationStatus) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *IntegrationStatus) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *IntegrationStatus) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *IntegrationStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IntegrationStatus) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IntegrationStatus) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IntegrationStatus) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IntegrationStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


