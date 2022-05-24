# ComplianceViolationAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Identifier for the alert | [optional] 
**Type** | Pointer to **string** | Type of alert generated | [optional] 
**State** | Pointer to **string** | Current state of the alert | [optional] 
**Resource** | Pointer to [**ComplianceResource**](ComplianceResource.md) |  | [optional] 
**ClosedBy** | Pointer to **string** | Account that closed the alert | [optional] 
**ClosedReason** | Pointer to **string** | Reason for closing the alert | [optional] 
**CreatedAt** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the alert was generated | [optional] 
**LastUpdated** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the alert was last modified | [optional] 
**ComplianceStatusReason** | Pointer to **string** | Reason for compliance check status. Compliance check could fail due to policy evaluation or blacklisting or errors evaluating compliance | [optional] 
**ViolationsCount** | Pointer to **int32** | Number of STOP action results in the compliance check report | [optional] 

## Methods

### NewComplianceViolationAlert

`func NewComplianceViolationAlert() *ComplianceViolationAlert`

NewComplianceViolationAlert instantiates a new ComplianceViolationAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceViolationAlertWithDefaults

`func NewComplianceViolationAlertWithDefaults() *ComplianceViolationAlert`

NewComplianceViolationAlertWithDefaults instantiates a new ComplianceViolationAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ComplianceViolationAlert) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComplianceViolationAlert) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComplianceViolationAlert) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComplianceViolationAlert) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *ComplianceViolationAlert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComplianceViolationAlert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComplianceViolationAlert) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ComplianceViolationAlert) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *ComplianceViolationAlert) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ComplianceViolationAlert) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ComplianceViolationAlert) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ComplianceViolationAlert) HasState() bool`

HasState returns a boolean if a field has been set.

### GetResource

`func (o *ComplianceViolationAlert) GetResource() ComplianceResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ComplianceViolationAlert) GetResourceOk() (*ComplianceResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ComplianceViolationAlert) SetResource(v ComplianceResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ComplianceViolationAlert) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetClosedBy

`func (o *ComplianceViolationAlert) GetClosedBy() string`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *ComplianceViolationAlert) GetClosedByOk() (*string, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *ComplianceViolationAlert) SetClosedBy(v string)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *ComplianceViolationAlert) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetClosedReason

`func (o *ComplianceViolationAlert) GetClosedReason() string`

GetClosedReason returns the ClosedReason field if non-nil, zero value otherwise.

### GetClosedReasonOk

`func (o *ComplianceViolationAlert) GetClosedReasonOk() (*string, bool)`

GetClosedReasonOk returns a tuple with the ClosedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedReason

`func (o *ComplianceViolationAlert) SetClosedReason(v string)`

SetClosedReason sets ClosedReason field to given value.

### HasClosedReason

`func (o *ComplianceViolationAlert) HasClosedReason() bool`

HasClosedReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ComplianceViolationAlert) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComplianceViolationAlert) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComplianceViolationAlert) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ComplianceViolationAlert) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ComplianceViolationAlert) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ComplianceViolationAlert) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ComplianceViolationAlert) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ComplianceViolationAlert) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetComplianceStatusReason

`func (o *ComplianceViolationAlert) GetComplianceStatusReason() string`

GetComplianceStatusReason returns the ComplianceStatusReason field if non-nil, zero value otherwise.

### GetComplianceStatusReasonOk

`func (o *ComplianceViolationAlert) GetComplianceStatusReasonOk() (*string, bool)`

GetComplianceStatusReasonOk returns a tuple with the ComplianceStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceStatusReason

`func (o *ComplianceViolationAlert) SetComplianceStatusReason(v string)`

SetComplianceStatusReason sets ComplianceStatusReason field to given value.

### HasComplianceStatusReason

`func (o *ComplianceViolationAlert) HasComplianceStatusReason() bool`

HasComplianceStatusReason returns a boolean if a field has been set.

### GetViolationsCount

`func (o *ComplianceViolationAlert) GetViolationsCount() int32`

GetViolationsCount returns the ViolationsCount field if non-nil, zero value otherwise.

### GetViolationsCountOk

`func (o *ComplianceViolationAlert) GetViolationsCountOk() (*int32, bool)`

GetViolationsCountOk returns a tuple with the ViolationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationsCount

`func (o *ComplianceViolationAlert) SetViolationsCount(v int32)`

SetViolationsCount sets ViolationsCount field to given value.

### HasViolationsCount

`func (o *ComplianceViolationAlert) HasViolationsCount() bool`

HasViolationsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


