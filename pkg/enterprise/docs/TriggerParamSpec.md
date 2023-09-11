# TriggerParamSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Parameter name as it appears in policy document | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Example** | Pointer to **NullableString** | An example value for the parameter (encoded as a string if the parameter is an object or list type) | [optional] 
**State** | Pointer to **string** | State of the trigger parameter | [optional] 
**SupersededBy** | Pointer to **NullableString** | The name of another trigger that supersedes this on functionally if this is deprecated | [optional] 
**Required** | Pointer to **bool** | Is this a required parameter or optional | [optional] 
**Validator** | Pointer to **interface{}** | If present, a definition for validation of input. Typically a jsonschema object that can be used to validate an input against. | [optional] 

## Methods

### NewTriggerParamSpec

`func NewTriggerParamSpec() *TriggerParamSpec`

NewTriggerParamSpec instantiates a new TriggerParamSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerParamSpecWithDefaults

`func NewTriggerParamSpecWithDefaults() *TriggerParamSpec`

NewTriggerParamSpecWithDefaults instantiates a new TriggerParamSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TriggerParamSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerParamSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerParamSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TriggerParamSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TriggerParamSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TriggerParamSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TriggerParamSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TriggerParamSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *TriggerParamSpec) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *TriggerParamSpec) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *TriggerParamSpec) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *TriggerParamSpec) HasExample() bool`

HasExample returns a boolean if a field has been set.

### SetExampleNil

`func (o *TriggerParamSpec) SetExampleNil(b bool)`

 SetExampleNil sets the value for Example to be an explicit nil

### UnsetExample
`func (o *TriggerParamSpec) UnsetExample()`

UnsetExample ensures that no value is present for Example, not even an explicit nil
### GetState

`func (o *TriggerParamSpec) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TriggerParamSpec) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TriggerParamSpec) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TriggerParamSpec) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSupersededBy

`func (o *TriggerParamSpec) GetSupersededBy() string`

GetSupersededBy returns the SupersededBy field if non-nil, zero value otherwise.

### GetSupersededByOk

`func (o *TriggerParamSpec) GetSupersededByOk() (*string, bool)`

GetSupersededByOk returns a tuple with the SupersededBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersededBy

`func (o *TriggerParamSpec) SetSupersededBy(v string)`

SetSupersededBy sets SupersededBy field to given value.

### HasSupersededBy

`func (o *TriggerParamSpec) HasSupersededBy() bool`

HasSupersededBy returns a boolean if a field has been set.

### SetSupersededByNil

`func (o *TriggerParamSpec) SetSupersededByNil(b bool)`

 SetSupersededByNil sets the value for SupersededBy to be an explicit nil

### UnsetSupersededBy
`func (o *TriggerParamSpec) UnsetSupersededBy()`

UnsetSupersededBy ensures that no value is present for SupersededBy, not even an explicit nil
### GetRequired

`func (o *TriggerParamSpec) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *TriggerParamSpec) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *TriggerParamSpec) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *TriggerParamSpec) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetValidator

`func (o *TriggerParamSpec) GetValidator() interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *TriggerParamSpec) GetValidatorOk() (*interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *TriggerParamSpec) SetValidator(v interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *TriggerParamSpec) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


