# SystemConfigurationPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | [**SystemConfigurationSchemaDefault**](SystemConfigurationSchemaDefault.md) |  | 

## Methods

### NewSystemConfigurationPut

`func NewSystemConfigurationPut(key string, value SystemConfigurationSchemaDefault, ) *SystemConfigurationPut`

NewSystemConfigurationPut instantiates a new SystemConfigurationPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemConfigurationPutWithDefaults

`func NewSystemConfigurationPutWithDefaults() *SystemConfigurationPut`

NewSystemConfigurationPutWithDefaults instantiates a new SystemConfigurationPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SystemConfigurationPut) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SystemConfigurationPut) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SystemConfigurationPut) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *SystemConfigurationPut) GetValue() SystemConfigurationSchemaDefault`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SystemConfigurationPut) GetValueOk() (*SystemConfigurationSchemaDefault, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SystemConfigurationPut) SetValue(v SystemConfigurationSchemaDefault)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


