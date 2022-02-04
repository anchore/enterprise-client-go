# AlertSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Identifier for the alert | [optional] 
**Type** | Pointer to **string** | Type of the alert | [optional] 
**State** | Pointer to **string** | Current state of the alert | [optional] 
**ResourceLabels** | Pointer to [**[]ResourceLabel**](ResourceLabel.md) |  | [optional] 
**ClosedBy** | Pointer to **string** | Account that closed the alert | [optional] 
**ClosedReason** | Pointer to **string** | Reason for closing the alert | [optional] 
**CreatedAt** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the alert was generated | [optional] 
**LastUpdated** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the alert was last modified | [optional] 

## Methods

### NewAlertSummary

`func NewAlertSummary() *AlertSummary`

NewAlertSummary instantiates a new AlertSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertSummaryWithDefaults

`func NewAlertSummaryWithDefaults() *AlertSummary`

NewAlertSummaryWithDefaults instantiates a new AlertSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AlertSummary) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AlertSummary) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AlertSummary) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AlertSummary) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *AlertSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AlertSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *AlertSummary) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AlertSummary) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AlertSummary) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AlertSummary) HasState() bool`

HasState returns a boolean if a field has been set.

### GetResourceLabels

`func (o *AlertSummary) GetResourceLabels() []ResourceLabel`

GetResourceLabels returns the ResourceLabels field if non-nil, zero value otherwise.

### GetResourceLabelsOk

`func (o *AlertSummary) GetResourceLabelsOk() (*[]ResourceLabel, bool)`

GetResourceLabelsOk returns a tuple with the ResourceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLabels

`func (o *AlertSummary) SetResourceLabels(v []ResourceLabel)`

SetResourceLabels sets ResourceLabels field to given value.

### HasResourceLabels

`func (o *AlertSummary) HasResourceLabels() bool`

HasResourceLabels returns a boolean if a field has been set.

### GetClosedBy

`func (o *AlertSummary) GetClosedBy() string`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *AlertSummary) GetClosedByOk() (*string, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *AlertSummary) SetClosedBy(v string)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *AlertSummary) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetClosedReason

`func (o *AlertSummary) GetClosedReason() string`

GetClosedReason returns the ClosedReason field if non-nil, zero value otherwise.

### GetClosedReasonOk

`func (o *AlertSummary) GetClosedReasonOk() (*string, bool)`

GetClosedReasonOk returns a tuple with the ClosedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedReason

`func (o *AlertSummary) SetClosedReason(v string)`

SetClosedReason sets ClosedReason field to given value.

### HasClosedReason

`func (o *AlertSummary) HasClosedReason() bool`

HasClosedReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AlertSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AlertSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AlertSummary) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AlertSummary) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AlertSummary) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AlertSummary) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


