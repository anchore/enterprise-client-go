# GateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gate name, as it would appear in a policy document | [optional] 
**Description** | Pointer to **string** | Description of the gate | [optional] 
**SupportedArtifactType** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | State of the gate and transitively all triggers it contains if not &#39;active&#39; | [optional] 
**SupersededBy** | Pointer to **NullableString** | The name of another trigger that supersedes this on functionally if this is deprecated | [optional] 
**Triggers** | Pointer to [**[]TriggerSpec**](TriggerSpec.md) | List of the triggers that can fire for this Gate | [optional] 

## Methods

### NewGateSpec

`func NewGateSpec() *GateSpec`

NewGateSpec instantiates a new GateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGateSpecWithDefaults

`func NewGateSpecWithDefaults() *GateSpec`

NewGateSpecWithDefaults instantiates a new GateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GateSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GateSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSupportedArtifactType

`func (o *GateSpec) GetSupportedArtifactType() string`

GetSupportedArtifactType returns the SupportedArtifactType field if non-nil, zero value otherwise.

### GetSupportedArtifactTypeOk

`func (o *GateSpec) GetSupportedArtifactTypeOk() (*string, bool)`

GetSupportedArtifactTypeOk returns a tuple with the SupportedArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedArtifactType

`func (o *GateSpec) SetSupportedArtifactType(v string)`

SetSupportedArtifactType sets SupportedArtifactType field to given value.

### HasSupportedArtifactType

`func (o *GateSpec) HasSupportedArtifactType() bool`

HasSupportedArtifactType returns a boolean if a field has been set.

### GetState

`func (o *GateSpec) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GateSpec) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GateSpec) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GateSpec) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSupersededBy

`func (o *GateSpec) GetSupersededBy() string`

GetSupersededBy returns the SupersededBy field if non-nil, zero value otherwise.

### GetSupersededByOk

`func (o *GateSpec) GetSupersededByOk() (*string, bool)`

GetSupersededByOk returns a tuple with the SupersededBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersededBy

`func (o *GateSpec) SetSupersededBy(v string)`

SetSupersededBy sets SupersededBy field to given value.

### HasSupersededBy

`func (o *GateSpec) HasSupersededBy() bool`

HasSupersededBy returns a boolean if a field has been set.

### SetSupersededByNil

`func (o *GateSpec) SetSupersededByNil(b bool)`

 SetSupersededByNil sets the value for SupersededBy to be an explicit nil

### UnsetSupersededBy
`func (o *GateSpec) UnsetSupersededBy()`

UnsetSupersededBy ensures that no value is present for SupersededBy, not even an explicit nil
### GetTriggers

`func (o *GateSpec) GetTriggers() []TriggerSpec`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *GateSpec) GetTriggersOk() (*[]TriggerSpec, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *GateSpec) SetTriggers(v []TriggerSpec)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *GateSpec) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


