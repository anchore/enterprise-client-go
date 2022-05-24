# TriggerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the trigger as it would appear in a policy document | [optional] 
**Description** | Pointer to **string** | Trigger description for what it tests and when it will fire during evaluation | [optional] 
**State** | Pointer to **string** | State of the trigger | [optional] 
**SupercededBy** | Pointer to **NullableString** | The name of another trigger that supercedes this on functionally if this is deprecated | [optional] 
**Parameters** | Pointer to [**[]TriggerParamSpec**](TriggerParamSpec.md) | The list of parameters that are valid for this trigger | [optional] 

## Methods

### NewTriggerSpec

`func NewTriggerSpec() *TriggerSpec`

NewTriggerSpec instantiates a new TriggerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerSpecWithDefaults

`func NewTriggerSpecWithDefaults() *TriggerSpec`

NewTriggerSpecWithDefaults instantiates a new TriggerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TriggerSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TriggerSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TriggerSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TriggerSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TriggerSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TriggerSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *TriggerSpec) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TriggerSpec) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TriggerSpec) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TriggerSpec) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSupercededBy

`func (o *TriggerSpec) GetSupercededBy() string`

GetSupercededBy returns the SupercededBy field if non-nil, zero value otherwise.

### GetSupercededByOk

`func (o *TriggerSpec) GetSupercededByOk() (*string, bool)`

GetSupercededByOk returns a tuple with the SupercededBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupercededBy

`func (o *TriggerSpec) SetSupercededBy(v string)`

SetSupercededBy sets SupercededBy field to given value.

### HasSupercededBy

`func (o *TriggerSpec) HasSupercededBy() bool`

HasSupercededBy returns a boolean if a field has been set.

### SetSupercededByNil

`func (o *TriggerSpec) SetSupercededByNil(b bool)`

 SetSupercededByNil sets the value for SupercededBy to be an explicit nil

### UnsetSupercededBy
`func (o *TriggerSpec) UnsetSupercededBy()`

UnsetSupercededBy ensures that no value is present for SupercededBy, not even an explicit nil
### GetParameters

`func (o *TriggerSpec) GetParameters() []TriggerParamSpec`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TriggerSpec) GetParametersOk() (*[]TriggerParamSpec, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TriggerSpec) SetParameters(v []TriggerParamSpec)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TriggerSpec) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


