# SystemResourceLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The resource being limited | [optional] 
**WarningValue** | Pointer to **NullableInt32** | The value at which a warning will be generated | [optional] 
**LimitValue** | Pointer to **NullableInt32** | The value at which a limit will be enforced | [optional] 
**CurrentValue** | Pointer to **int32** | The current value of the resource | [optional] 
**LimitExceeded** | Pointer to **bool** | Whether the limit has been exceeded and is being enforced | [optional] 

## Methods

### NewSystemResourceLimit

`func NewSystemResourceLimit() *SystemResourceLimit`

NewSystemResourceLimit instantiates a new SystemResourceLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemResourceLimitWithDefaults

`func NewSystemResourceLimitWithDefaults() *SystemResourceLimit`

NewSystemResourceLimitWithDefaults instantiates a new SystemResourceLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SystemResourceLimit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemResourceLimit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemResourceLimit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemResourceLimit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWarningValue

`func (o *SystemResourceLimit) GetWarningValue() int32`

GetWarningValue returns the WarningValue field if non-nil, zero value otherwise.

### GetWarningValueOk

`func (o *SystemResourceLimit) GetWarningValueOk() (*int32, bool)`

GetWarningValueOk returns a tuple with the WarningValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningValue

`func (o *SystemResourceLimit) SetWarningValue(v int32)`

SetWarningValue sets WarningValue field to given value.

### HasWarningValue

`func (o *SystemResourceLimit) HasWarningValue() bool`

HasWarningValue returns a boolean if a field has been set.

### SetWarningValueNil

`func (o *SystemResourceLimit) SetWarningValueNil(b bool)`

 SetWarningValueNil sets the value for WarningValue to be an explicit nil

### UnsetWarningValue
`func (o *SystemResourceLimit) UnsetWarningValue()`

UnsetWarningValue ensures that no value is present for WarningValue, not even an explicit nil
### GetLimitValue

`func (o *SystemResourceLimit) GetLimitValue() int32`

GetLimitValue returns the LimitValue field if non-nil, zero value otherwise.

### GetLimitValueOk

`func (o *SystemResourceLimit) GetLimitValueOk() (*int32, bool)`

GetLimitValueOk returns a tuple with the LimitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitValue

`func (o *SystemResourceLimit) SetLimitValue(v int32)`

SetLimitValue sets LimitValue field to given value.

### HasLimitValue

`func (o *SystemResourceLimit) HasLimitValue() bool`

HasLimitValue returns a boolean if a field has been set.

### SetLimitValueNil

`func (o *SystemResourceLimit) SetLimitValueNil(b bool)`

 SetLimitValueNil sets the value for LimitValue to be an explicit nil

### UnsetLimitValue
`func (o *SystemResourceLimit) UnsetLimitValue()`

UnsetLimitValue ensures that no value is present for LimitValue, not even an explicit nil
### GetCurrentValue

`func (o *SystemResourceLimit) GetCurrentValue() int32`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *SystemResourceLimit) GetCurrentValueOk() (*int32, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *SystemResourceLimit) SetCurrentValue(v int32)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *SystemResourceLimit) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetLimitExceeded

`func (o *SystemResourceLimit) GetLimitExceeded() bool`

GetLimitExceeded returns the LimitExceeded field if non-nil, zero value otherwise.

### GetLimitExceededOk

`func (o *SystemResourceLimit) GetLimitExceededOk() (*bool, bool)`

GetLimitExceededOk returns a tuple with the LimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitExceeded

`func (o *SystemResourceLimit) SetLimitExceeded(v bool)`

SetLimitExceeded sets LimitExceeded field to given value.

### HasLimitExceeded

`func (o *SystemResourceLimit) HasLimitExceeded() bool`

HasLimitExceeded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


