# SystemConfigurationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInfNan** | Pointer to **NullableBool** |  | [optional] 
**DecimalPlaces** | Pointer to **NullableInt64** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to [**NullableSystemConfigurationValue**](SystemConfigurationValue.md) |  | [optional] 
**Enum** | Pointer to **[]string** |  | [optional] 
**Ge** | Pointer to **NullableFloat32** |  | [optional] 
**Gt** | Pointer to **NullableFloat32** |  | [optional] 
**IsArray** | Pointer to **NullableBool** |  | [optional] 
**Le** | Pointer to **NullableFloat32** |  | [optional] 
**Lt** | Pointer to **NullableFloat32** |  | [optional] 
**MaxDigits** | Pointer to **NullableInt64** |  | [optional] 
**MaxItems** | Pointer to **NullableInt64** |  | [optional] 
**MaxLength** | Pointer to **NullableInt64** |  | [optional] 
**MinItems** | Pointer to **NullableInt64** |  | [optional] 
**MinLength** | Pointer to **NullableInt64** |  | [optional] 
**MultipleOf** | Pointer to **NullableFloat32** |  | [optional] 
**Regex** | Pointer to **NullableString** |  | [optional] 
**UniqueItems** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSystemConfigurationSchema

`func NewSystemConfigurationSchema() *SystemConfigurationSchema`

NewSystemConfigurationSchema instantiates a new SystemConfigurationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemConfigurationSchemaWithDefaults

`func NewSystemConfigurationSchemaWithDefaults() *SystemConfigurationSchema`

NewSystemConfigurationSchemaWithDefaults instantiates a new SystemConfigurationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInfNan

`func (o *SystemConfigurationSchema) GetAllowInfNan() bool`

GetAllowInfNan returns the AllowInfNan field if non-nil, zero value otherwise.

### GetAllowInfNanOk

`func (o *SystemConfigurationSchema) GetAllowInfNanOk() (*bool, bool)`

GetAllowInfNanOk returns a tuple with the AllowInfNan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInfNan

`func (o *SystemConfigurationSchema) SetAllowInfNan(v bool)`

SetAllowInfNan sets AllowInfNan field to given value.

### HasAllowInfNan

`func (o *SystemConfigurationSchema) HasAllowInfNan() bool`

HasAllowInfNan returns a boolean if a field has been set.

### SetAllowInfNanNil

`func (o *SystemConfigurationSchema) SetAllowInfNanNil(b bool)`

 SetAllowInfNanNil sets the value for AllowInfNan to be an explicit nil

### UnsetAllowInfNan
`func (o *SystemConfigurationSchema) UnsetAllowInfNan()`

UnsetAllowInfNan ensures that no value is present for AllowInfNan, not even an explicit nil
### GetDecimalPlaces

`func (o *SystemConfigurationSchema) GetDecimalPlaces() int64`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *SystemConfigurationSchema) GetDecimalPlacesOk() (*int64, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *SystemConfigurationSchema) SetDecimalPlaces(v int64)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *SystemConfigurationSchema) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### SetDecimalPlacesNil

`func (o *SystemConfigurationSchema) SetDecimalPlacesNil(b bool)`

 SetDecimalPlacesNil sets the value for DecimalPlaces to be an explicit nil

### UnsetDecimalPlaces
`func (o *SystemConfigurationSchema) UnsetDecimalPlaces()`

UnsetDecimalPlaces ensures that no value is present for DecimalPlaces, not even an explicit nil
### GetDataType

`func (o *SystemConfigurationSchema) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *SystemConfigurationSchema) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *SystemConfigurationSchema) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *SystemConfigurationSchema) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDefaultValue

`func (o *SystemConfigurationSchema) GetDefaultValue() SystemConfigurationValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SystemConfigurationSchema) GetDefaultValueOk() (*SystemConfigurationValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SystemConfigurationSchema) SetDefaultValue(v SystemConfigurationValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SystemConfigurationSchema) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *SystemConfigurationSchema) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *SystemConfigurationSchema) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetEnum

`func (o *SystemConfigurationSchema) GetEnum() []string`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *SystemConfigurationSchema) GetEnumOk() (*[]string, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *SystemConfigurationSchema) SetEnum(v []string)`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *SystemConfigurationSchema) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### SetEnumNil

`func (o *SystemConfigurationSchema) SetEnumNil(b bool)`

 SetEnumNil sets the value for Enum to be an explicit nil

### UnsetEnum
`func (o *SystemConfigurationSchema) UnsetEnum()`

UnsetEnum ensures that no value is present for Enum, not even an explicit nil
### GetGe

`func (o *SystemConfigurationSchema) GetGe() float32`

GetGe returns the Ge field if non-nil, zero value otherwise.

### GetGeOk

`func (o *SystemConfigurationSchema) GetGeOk() (*float32, bool)`

GetGeOk returns a tuple with the Ge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGe

`func (o *SystemConfigurationSchema) SetGe(v float32)`

SetGe sets Ge field to given value.

### HasGe

`func (o *SystemConfigurationSchema) HasGe() bool`

HasGe returns a boolean if a field has been set.

### SetGeNil

`func (o *SystemConfigurationSchema) SetGeNil(b bool)`

 SetGeNil sets the value for Ge to be an explicit nil

### UnsetGe
`func (o *SystemConfigurationSchema) UnsetGe()`

UnsetGe ensures that no value is present for Ge, not even an explicit nil
### GetGt

`func (o *SystemConfigurationSchema) GetGt() float32`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *SystemConfigurationSchema) GetGtOk() (*float32, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *SystemConfigurationSchema) SetGt(v float32)`

SetGt sets Gt field to given value.

### HasGt

`func (o *SystemConfigurationSchema) HasGt() bool`

HasGt returns a boolean if a field has been set.

### SetGtNil

`func (o *SystemConfigurationSchema) SetGtNil(b bool)`

 SetGtNil sets the value for Gt to be an explicit nil

### UnsetGt
`func (o *SystemConfigurationSchema) UnsetGt()`

UnsetGt ensures that no value is present for Gt, not even an explicit nil
### GetIsArray

`func (o *SystemConfigurationSchema) GetIsArray() bool`

GetIsArray returns the IsArray field if non-nil, zero value otherwise.

### GetIsArrayOk

`func (o *SystemConfigurationSchema) GetIsArrayOk() (*bool, bool)`

GetIsArrayOk returns a tuple with the IsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArray

`func (o *SystemConfigurationSchema) SetIsArray(v bool)`

SetIsArray sets IsArray field to given value.

### HasIsArray

`func (o *SystemConfigurationSchema) HasIsArray() bool`

HasIsArray returns a boolean if a field has been set.

### SetIsArrayNil

`func (o *SystemConfigurationSchema) SetIsArrayNil(b bool)`

 SetIsArrayNil sets the value for IsArray to be an explicit nil

### UnsetIsArray
`func (o *SystemConfigurationSchema) UnsetIsArray()`

UnsetIsArray ensures that no value is present for IsArray, not even an explicit nil
### GetLe

`func (o *SystemConfigurationSchema) GetLe() float32`

GetLe returns the Le field if non-nil, zero value otherwise.

### GetLeOk

`func (o *SystemConfigurationSchema) GetLeOk() (*float32, bool)`

GetLeOk returns a tuple with the Le field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLe

`func (o *SystemConfigurationSchema) SetLe(v float32)`

SetLe sets Le field to given value.

### HasLe

`func (o *SystemConfigurationSchema) HasLe() bool`

HasLe returns a boolean if a field has been set.

### SetLeNil

`func (o *SystemConfigurationSchema) SetLeNil(b bool)`

 SetLeNil sets the value for Le to be an explicit nil

### UnsetLe
`func (o *SystemConfigurationSchema) UnsetLe()`

UnsetLe ensures that no value is present for Le, not even an explicit nil
### GetLt

`func (o *SystemConfigurationSchema) GetLt() float32`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *SystemConfigurationSchema) GetLtOk() (*float32, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *SystemConfigurationSchema) SetLt(v float32)`

SetLt sets Lt field to given value.

### HasLt

`func (o *SystemConfigurationSchema) HasLt() bool`

HasLt returns a boolean if a field has been set.

### SetLtNil

`func (o *SystemConfigurationSchema) SetLtNil(b bool)`

 SetLtNil sets the value for Lt to be an explicit nil

### UnsetLt
`func (o *SystemConfigurationSchema) UnsetLt()`

UnsetLt ensures that no value is present for Lt, not even an explicit nil
### GetMaxDigits

`func (o *SystemConfigurationSchema) GetMaxDigits() int64`

GetMaxDigits returns the MaxDigits field if non-nil, zero value otherwise.

### GetMaxDigitsOk

`func (o *SystemConfigurationSchema) GetMaxDigitsOk() (*int64, bool)`

GetMaxDigitsOk returns a tuple with the MaxDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDigits

`func (o *SystemConfigurationSchema) SetMaxDigits(v int64)`

SetMaxDigits sets MaxDigits field to given value.

### HasMaxDigits

`func (o *SystemConfigurationSchema) HasMaxDigits() bool`

HasMaxDigits returns a boolean if a field has been set.

### SetMaxDigitsNil

`func (o *SystemConfigurationSchema) SetMaxDigitsNil(b bool)`

 SetMaxDigitsNil sets the value for MaxDigits to be an explicit nil

### UnsetMaxDigits
`func (o *SystemConfigurationSchema) UnsetMaxDigits()`

UnsetMaxDigits ensures that no value is present for MaxDigits, not even an explicit nil
### GetMaxItems

`func (o *SystemConfigurationSchema) GetMaxItems() int64`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *SystemConfigurationSchema) GetMaxItemsOk() (*int64, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *SystemConfigurationSchema) SetMaxItems(v int64)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *SystemConfigurationSchema) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *SystemConfigurationSchema) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *SystemConfigurationSchema) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetMaxLength

`func (o *SystemConfigurationSchema) GetMaxLength() int64`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *SystemConfigurationSchema) GetMaxLengthOk() (*int64, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *SystemConfigurationSchema) SetMaxLength(v int64)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *SystemConfigurationSchema) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *SystemConfigurationSchema) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *SystemConfigurationSchema) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetMinItems

`func (o *SystemConfigurationSchema) GetMinItems() int64`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *SystemConfigurationSchema) GetMinItemsOk() (*int64, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *SystemConfigurationSchema) SetMinItems(v int64)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *SystemConfigurationSchema) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### SetMinItemsNil

`func (o *SystemConfigurationSchema) SetMinItemsNil(b bool)`

 SetMinItemsNil sets the value for MinItems to be an explicit nil

### UnsetMinItems
`func (o *SystemConfigurationSchema) UnsetMinItems()`

UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
### GetMinLength

`func (o *SystemConfigurationSchema) GetMinLength() int64`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *SystemConfigurationSchema) GetMinLengthOk() (*int64, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *SystemConfigurationSchema) SetMinLength(v int64)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *SystemConfigurationSchema) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *SystemConfigurationSchema) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *SystemConfigurationSchema) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMultipleOf

`func (o *SystemConfigurationSchema) GetMultipleOf() float32`

GetMultipleOf returns the MultipleOf field if non-nil, zero value otherwise.

### GetMultipleOfOk

`func (o *SystemConfigurationSchema) GetMultipleOfOk() (*float32, bool)`

GetMultipleOfOk returns a tuple with the MultipleOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleOf

`func (o *SystemConfigurationSchema) SetMultipleOf(v float32)`

SetMultipleOf sets MultipleOf field to given value.

### HasMultipleOf

`func (o *SystemConfigurationSchema) HasMultipleOf() bool`

HasMultipleOf returns a boolean if a field has been set.

### SetMultipleOfNil

`func (o *SystemConfigurationSchema) SetMultipleOfNil(b bool)`

 SetMultipleOfNil sets the value for MultipleOf to be an explicit nil

### UnsetMultipleOf
`func (o *SystemConfigurationSchema) UnsetMultipleOf()`

UnsetMultipleOf ensures that no value is present for MultipleOf, not even an explicit nil
### GetRegex

`func (o *SystemConfigurationSchema) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *SystemConfigurationSchema) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *SystemConfigurationSchema) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *SystemConfigurationSchema) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *SystemConfigurationSchema) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *SystemConfigurationSchema) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetUniqueItems

`func (o *SystemConfigurationSchema) GetUniqueItems() bool`

GetUniqueItems returns the UniqueItems field if non-nil, zero value otherwise.

### GetUniqueItemsOk

`func (o *SystemConfigurationSchema) GetUniqueItemsOk() (*bool, bool)`

GetUniqueItemsOk returns a tuple with the UniqueItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueItems

`func (o *SystemConfigurationSchema) SetUniqueItems(v bool)`

SetUniqueItems sets UniqueItems field to given value.

### HasUniqueItems

`func (o *SystemConfigurationSchema) HasUniqueItems() bool`

HasUniqueItems returns a boolean if a field has been set.

### SetUniqueItemsNil

`func (o *SystemConfigurationSchema) SetUniqueItemsNil(b bool)`

 SetUniqueItemsNil sets the value for UniqueItems to be an explicit nil

### UnsetUniqueItems
`func (o *SystemConfigurationSchema) UnsetUniqueItems()`

UnsetUniqueItems ensures that no value is present for UniqueItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


