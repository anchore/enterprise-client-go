# SystemConfigurationPatchInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | [**SystemConfigurationValue**](SystemConfigurationValue.md) |  | 

## Methods

### NewSystemConfigurationPatchInner

`func NewSystemConfigurationPatchInner(key string, value SystemConfigurationValue, ) *SystemConfigurationPatchInner`

NewSystemConfigurationPatchInner instantiates a new SystemConfigurationPatchInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemConfigurationPatchInnerWithDefaults

`func NewSystemConfigurationPatchInnerWithDefaults() *SystemConfigurationPatchInner`

NewSystemConfigurationPatchInnerWithDefaults instantiates a new SystemConfigurationPatchInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SystemConfigurationPatchInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SystemConfigurationPatchInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SystemConfigurationPatchInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *SystemConfigurationPatchInner) GetValue() SystemConfigurationValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SystemConfigurationPatchInner) GetValueOk() (*SystemConfigurationValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SystemConfigurationPatchInner) SetValue(v SystemConfigurationValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

