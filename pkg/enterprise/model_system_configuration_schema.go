/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the SystemConfigurationSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemConfigurationSchema{}

// SystemConfigurationSchema struct for SystemConfigurationSchema
type SystemConfigurationSchema struct {
	AllowInfNan NullableBool `json:"allow_inf_nan,omitempty"`
	DecimalPlaces NullableInt32 `json:"decimal_places,omitempty"`
	DataType *string `json:"data_type,omitempty"`
	DefaultValue NullableSystemConfigurationValue `json:"default_value,omitempty"`
	Enum []string `json:"enum,omitempty"`
	Ge NullableFloat32 `json:"ge,omitempty"`
	Gt NullableFloat32 `json:"gt,omitempty"`
	IsArray NullableBool `json:"is_array,omitempty"`
	Le NullableFloat32 `json:"le,omitempty"`
	Lt NullableFloat32 `json:"lt,omitempty"`
	MaxDigits NullableInt32 `json:"max_digits,omitempty"`
	MaxItems NullableInt32 `json:"max_items,omitempty"`
	MaxLength NullableInt32 `json:"max_length,omitempty"`
	MinItems NullableInt32 `json:"min_items,omitempty"`
	MinLength NullableInt32 `json:"min_length,omitempty"`
	MultipleOf NullableFloat32 `json:"multiple_of,omitempty"`
	Regex NullableString `json:"regex,omitempty"`
	UniqueItems NullableBool `json:"unique_items,omitempty"`
}

// NewSystemConfigurationSchema instantiates a new SystemConfigurationSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemConfigurationSchema() *SystemConfigurationSchema {
	this := SystemConfigurationSchema{}
	return &this
}

// NewSystemConfigurationSchemaWithDefaults instantiates a new SystemConfigurationSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemConfigurationSchemaWithDefaults() *SystemConfigurationSchema {
	this := SystemConfigurationSchema{}
	return &this
}

// GetAllowInfNan returns the AllowInfNan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetAllowInfNan() bool {
	if o == nil || IsNil(o.AllowInfNan.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowInfNan.Get()
}

// GetAllowInfNanOk returns a tuple with the AllowInfNan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetAllowInfNanOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowInfNan.Get(), o.AllowInfNan.IsSet()
}

// HasAllowInfNan returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasAllowInfNan() bool {
	if o != nil && o.AllowInfNan.IsSet() {
		return true
	}

	return false
}

// SetAllowInfNan gets a reference to the given NullableBool and assigns it to the AllowInfNan field.
func (o *SystemConfigurationSchema) SetAllowInfNan(v bool) {
	o.AllowInfNan.Set(&v)
}
// SetAllowInfNanNil sets the value for AllowInfNan to be an explicit nil
func (o *SystemConfigurationSchema) SetAllowInfNanNil() {
	o.AllowInfNan.Set(nil)
}

// UnsetAllowInfNan ensures that no value is present for AllowInfNan, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetAllowInfNan() {
	o.AllowInfNan.Unset()
}

// GetDecimalPlaces returns the DecimalPlaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetDecimalPlaces() int32 {
	if o == nil || IsNil(o.DecimalPlaces.Get()) {
		var ret int32
		return ret
	}
	return *o.DecimalPlaces.Get()
}

// GetDecimalPlacesOk returns a tuple with the DecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetDecimalPlacesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecimalPlaces.Get(), o.DecimalPlaces.IsSet()
}

// HasDecimalPlaces returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasDecimalPlaces() bool {
	if o != nil && o.DecimalPlaces.IsSet() {
		return true
	}

	return false
}

// SetDecimalPlaces gets a reference to the given NullableInt32 and assigns it to the DecimalPlaces field.
func (o *SystemConfigurationSchema) SetDecimalPlaces(v int32) {
	o.DecimalPlaces.Set(&v)
}
// SetDecimalPlacesNil sets the value for DecimalPlaces to be an explicit nil
func (o *SystemConfigurationSchema) SetDecimalPlacesNil() {
	o.DecimalPlaces.Set(nil)
}

// UnsetDecimalPlaces ensures that no value is present for DecimalPlaces, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetDecimalPlaces() {
	o.DecimalPlaces.Unset()
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *SystemConfigurationSchema) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemConfigurationSchema) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *SystemConfigurationSchema) SetDataType(v string) {
	o.DataType = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetDefaultValue() SystemConfigurationValue {
	if o == nil || IsNil(o.DefaultValue.Get()) {
		var ret SystemConfigurationValue
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetDefaultValueOk() (*SystemConfigurationValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasDefaultValue() bool {
	if o != nil && o.DefaultValue.IsSet() {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given NullableSystemConfigurationValue and assigns it to the DefaultValue field.
func (o *SystemConfigurationSchema) SetDefaultValue(v SystemConfigurationValue) {
	o.DefaultValue.Set(&v)
}
// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil
func (o *SystemConfigurationSchema) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

// GetEnum returns the Enum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetEnum() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetEnumOk() ([]string, bool) {
	if o == nil || IsNil(o.Enum) {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasEnum() bool {
	if o != nil && !IsNil(o.Enum) {
		return true
	}

	return false
}

// SetEnum gets a reference to the given []string and assigns it to the Enum field.
func (o *SystemConfigurationSchema) SetEnum(v []string) {
	o.Enum = v
}

// GetGe returns the Ge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetGe() float32 {
	if o == nil || IsNil(o.Ge.Get()) {
		var ret float32
		return ret
	}
	return *o.Ge.Get()
}

// GetGeOk returns a tuple with the Ge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetGeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ge.Get(), o.Ge.IsSet()
}

// HasGe returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasGe() bool {
	if o != nil && o.Ge.IsSet() {
		return true
	}

	return false
}

// SetGe gets a reference to the given NullableFloat32 and assigns it to the Ge field.
func (o *SystemConfigurationSchema) SetGe(v float32) {
	o.Ge.Set(&v)
}
// SetGeNil sets the value for Ge to be an explicit nil
func (o *SystemConfigurationSchema) SetGeNil() {
	o.Ge.Set(nil)
}

// UnsetGe ensures that no value is present for Ge, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetGe() {
	o.Ge.Unset()
}

// GetGt returns the Gt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetGt() float32 {
	if o == nil || IsNil(o.Gt.Get()) {
		var ret float32
		return ret
	}
	return *o.Gt.Get()
}

// GetGtOk returns a tuple with the Gt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetGtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gt.Get(), o.Gt.IsSet()
}

// HasGt returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasGt() bool {
	if o != nil && o.Gt.IsSet() {
		return true
	}

	return false
}

// SetGt gets a reference to the given NullableFloat32 and assigns it to the Gt field.
func (o *SystemConfigurationSchema) SetGt(v float32) {
	o.Gt.Set(&v)
}
// SetGtNil sets the value for Gt to be an explicit nil
func (o *SystemConfigurationSchema) SetGtNil() {
	o.Gt.Set(nil)
}

// UnsetGt ensures that no value is present for Gt, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetGt() {
	o.Gt.Unset()
}

// GetIsArray returns the IsArray field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetIsArray() bool {
	if o == nil || IsNil(o.IsArray.Get()) {
		var ret bool
		return ret
	}
	return *o.IsArray.Get()
}

// GetIsArrayOk returns a tuple with the IsArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetIsArrayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsArray.Get(), o.IsArray.IsSet()
}

// HasIsArray returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasIsArray() bool {
	if o != nil && o.IsArray.IsSet() {
		return true
	}

	return false
}

// SetIsArray gets a reference to the given NullableBool and assigns it to the IsArray field.
func (o *SystemConfigurationSchema) SetIsArray(v bool) {
	o.IsArray.Set(&v)
}
// SetIsArrayNil sets the value for IsArray to be an explicit nil
func (o *SystemConfigurationSchema) SetIsArrayNil() {
	o.IsArray.Set(nil)
}

// UnsetIsArray ensures that no value is present for IsArray, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetIsArray() {
	o.IsArray.Unset()
}

// GetLe returns the Le field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetLe() float32 {
	if o == nil || IsNil(o.Le.Get()) {
		var ret float32
		return ret
	}
	return *o.Le.Get()
}

// GetLeOk returns a tuple with the Le field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetLeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Le.Get(), o.Le.IsSet()
}

// HasLe returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasLe() bool {
	if o != nil && o.Le.IsSet() {
		return true
	}

	return false
}

// SetLe gets a reference to the given NullableFloat32 and assigns it to the Le field.
func (o *SystemConfigurationSchema) SetLe(v float32) {
	o.Le.Set(&v)
}
// SetLeNil sets the value for Le to be an explicit nil
func (o *SystemConfigurationSchema) SetLeNil() {
	o.Le.Set(nil)
}

// UnsetLe ensures that no value is present for Le, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetLe() {
	o.Le.Unset()
}

// GetLt returns the Lt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetLt() float32 {
	if o == nil || IsNil(o.Lt.Get()) {
		var ret float32
		return ret
	}
	return *o.Lt.Get()
}

// GetLtOk returns a tuple with the Lt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetLtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lt.Get(), o.Lt.IsSet()
}

// HasLt returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasLt() bool {
	if o != nil && o.Lt.IsSet() {
		return true
	}

	return false
}

// SetLt gets a reference to the given NullableFloat32 and assigns it to the Lt field.
func (o *SystemConfigurationSchema) SetLt(v float32) {
	o.Lt.Set(&v)
}
// SetLtNil sets the value for Lt to be an explicit nil
func (o *SystemConfigurationSchema) SetLtNil() {
	o.Lt.Set(nil)
}

// UnsetLt ensures that no value is present for Lt, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetLt() {
	o.Lt.Unset()
}

// GetMaxDigits returns the MaxDigits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetMaxDigits() int32 {
	if o == nil || IsNil(o.MaxDigits.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxDigits.Get()
}

// GetMaxDigitsOk returns a tuple with the MaxDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetMaxDigitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxDigits.Get(), o.MaxDigits.IsSet()
}

// HasMaxDigits returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasMaxDigits() bool {
	if o != nil && o.MaxDigits.IsSet() {
		return true
	}

	return false
}

// SetMaxDigits gets a reference to the given NullableInt32 and assigns it to the MaxDigits field.
func (o *SystemConfigurationSchema) SetMaxDigits(v int32) {
	o.MaxDigits.Set(&v)
}
// SetMaxDigitsNil sets the value for MaxDigits to be an explicit nil
func (o *SystemConfigurationSchema) SetMaxDigitsNil() {
	o.MaxDigits.Set(nil)
}

// UnsetMaxDigits ensures that no value is present for MaxDigits, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetMaxDigits() {
	o.MaxDigits.Unset()
}

// GetMaxItems returns the MaxItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetMaxItems() int32 {
	if o == nil || IsNil(o.MaxItems.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxItems.Get()
}

// GetMaxItemsOk returns a tuple with the MaxItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetMaxItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxItems.Get(), o.MaxItems.IsSet()
}

// HasMaxItems returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasMaxItems() bool {
	if o != nil && o.MaxItems.IsSet() {
		return true
	}

	return false
}

// SetMaxItems gets a reference to the given NullableInt32 and assigns it to the MaxItems field.
func (o *SystemConfigurationSchema) SetMaxItems(v int32) {
	o.MaxItems.Set(&v)
}
// SetMaxItemsNil sets the value for MaxItems to be an explicit nil
func (o *SystemConfigurationSchema) SetMaxItemsNil() {
	o.MaxItems.Set(nil)
}

// UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetMaxItems() {
	o.MaxItems.Unset()
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetMaxLength() int32 {
	if o == nil || IsNil(o.MaxLength.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxLength.Get()
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetMaxLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxLength.Get(), o.MaxLength.IsSet()
}

// HasMaxLength returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasMaxLength() bool {
	if o != nil && o.MaxLength.IsSet() {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given NullableInt32 and assigns it to the MaxLength field.
func (o *SystemConfigurationSchema) SetMaxLength(v int32) {
	o.MaxLength.Set(&v)
}
// SetMaxLengthNil sets the value for MaxLength to be an explicit nil
func (o *SystemConfigurationSchema) SetMaxLengthNil() {
	o.MaxLength.Set(nil)
}

// UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetMaxLength() {
	o.MaxLength.Unset()
}

// GetMinItems returns the MinItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetMinItems() int32 {
	if o == nil || IsNil(o.MinItems.Get()) {
		var ret int32
		return ret
	}
	return *o.MinItems.Get()
}

// GetMinItemsOk returns a tuple with the MinItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetMinItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinItems.Get(), o.MinItems.IsSet()
}

// HasMinItems returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasMinItems() bool {
	if o != nil && o.MinItems.IsSet() {
		return true
	}

	return false
}

// SetMinItems gets a reference to the given NullableInt32 and assigns it to the MinItems field.
func (o *SystemConfigurationSchema) SetMinItems(v int32) {
	o.MinItems.Set(&v)
}
// SetMinItemsNil sets the value for MinItems to be an explicit nil
func (o *SystemConfigurationSchema) SetMinItemsNil() {
	o.MinItems.Set(nil)
}

// UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetMinItems() {
	o.MinItems.Unset()
}

// GetMinLength returns the MinLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength.Get()) {
		var ret int32
		return ret
	}
	return *o.MinLength.Get()
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetMinLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinLength.Get(), o.MinLength.IsSet()
}

// HasMinLength returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasMinLength() bool {
	if o != nil && o.MinLength.IsSet() {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given NullableInt32 and assigns it to the MinLength field.
func (o *SystemConfigurationSchema) SetMinLength(v int32) {
	o.MinLength.Set(&v)
}
// SetMinLengthNil sets the value for MinLength to be an explicit nil
func (o *SystemConfigurationSchema) SetMinLengthNil() {
	o.MinLength.Set(nil)
}

// UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetMinLength() {
	o.MinLength.Unset()
}

// GetMultipleOf returns the MultipleOf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetMultipleOf() float32 {
	if o == nil || IsNil(o.MultipleOf.Get()) {
		var ret float32
		return ret
	}
	return *o.MultipleOf.Get()
}

// GetMultipleOfOk returns a tuple with the MultipleOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetMultipleOfOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MultipleOf.Get(), o.MultipleOf.IsSet()
}

// HasMultipleOf returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasMultipleOf() bool {
	if o != nil && o.MultipleOf.IsSet() {
		return true
	}

	return false
}

// SetMultipleOf gets a reference to the given NullableFloat32 and assigns it to the MultipleOf field.
func (o *SystemConfigurationSchema) SetMultipleOf(v float32) {
	o.MultipleOf.Set(&v)
}
// SetMultipleOfNil sets the value for MultipleOf to be an explicit nil
func (o *SystemConfigurationSchema) SetMultipleOfNil() {
	o.MultipleOf.Set(nil)
}

// UnsetMultipleOf ensures that no value is present for MultipleOf, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetMultipleOf() {
	o.MultipleOf.Unset()
}

// GetRegex returns the Regex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetRegex() string {
	if o == nil || IsNil(o.Regex.Get()) {
		var ret string
		return ret
	}
	return *o.Regex.Get()
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regex.Get(), o.Regex.IsSet()
}

// HasRegex returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasRegex() bool {
	if o != nil && o.Regex.IsSet() {
		return true
	}

	return false
}

// SetRegex gets a reference to the given NullableString and assigns it to the Regex field.
func (o *SystemConfigurationSchema) SetRegex(v string) {
	o.Regex.Set(&v)
}
// SetRegexNil sets the value for Regex to be an explicit nil
func (o *SystemConfigurationSchema) SetRegexNil() {
	o.Regex.Set(nil)
}

// UnsetRegex ensures that no value is present for Regex, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetRegex() {
	o.Regex.Unset()
}

// GetUniqueItems returns the UniqueItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfigurationSchema) GetUniqueItems() bool {
	if o == nil || IsNil(o.UniqueItems.Get()) {
		var ret bool
		return ret
	}
	return *o.UniqueItems.Get()
}

// GetUniqueItemsOk returns a tuple with the UniqueItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationSchema) GetUniqueItemsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UniqueItems.Get(), o.UniqueItems.IsSet()
}

// HasUniqueItems returns a boolean if a field has been set.
func (o *SystemConfigurationSchema) HasUniqueItems() bool {
	if o != nil && o.UniqueItems.IsSet() {
		return true
	}

	return false
}

// SetUniqueItems gets a reference to the given NullableBool and assigns it to the UniqueItems field.
func (o *SystemConfigurationSchema) SetUniqueItems(v bool) {
	o.UniqueItems.Set(&v)
}
// SetUniqueItemsNil sets the value for UniqueItems to be an explicit nil
func (o *SystemConfigurationSchema) SetUniqueItemsNil() {
	o.UniqueItems.Set(nil)
}

// UnsetUniqueItems ensures that no value is present for UniqueItems, not even an explicit nil
func (o *SystemConfigurationSchema) UnsetUniqueItems() {
	o.UniqueItems.Unset()
}

func (o SystemConfigurationSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemConfigurationSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowInfNan.IsSet() {
		toSerialize["allow_inf_nan"] = o.AllowInfNan.Get()
	}
	if o.DecimalPlaces.IsSet() {
		toSerialize["decimal_places"] = o.DecimalPlaces.Get()
	}
	if !IsNil(o.DataType) {
		toSerialize["data_type"] = o.DataType
	}
	if o.DefaultValue.IsSet() {
		toSerialize["default_value"] = o.DefaultValue.Get()
	}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	if o.Ge.IsSet() {
		toSerialize["ge"] = o.Ge.Get()
	}
	if o.Gt.IsSet() {
		toSerialize["gt"] = o.Gt.Get()
	}
	if o.IsArray.IsSet() {
		toSerialize["is_array"] = o.IsArray.Get()
	}
	if o.Le.IsSet() {
		toSerialize["le"] = o.Le.Get()
	}
	if o.Lt.IsSet() {
		toSerialize["lt"] = o.Lt.Get()
	}
	if o.MaxDigits.IsSet() {
		toSerialize["max_digits"] = o.MaxDigits.Get()
	}
	if o.MaxItems.IsSet() {
		toSerialize["max_items"] = o.MaxItems.Get()
	}
	if o.MaxLength.IsSet() {
		toSerialize["max_length"] = o.MaxLength.Get()
	}
	if o.MinItems.IsSet() {
		toSerialize["min_items"] = o.MinItems.Get()
	}
	if o.MinLength.IsSet() {
		toSerialize["min_length"] = o.MinLength.Get()
	}
	if o.MultipleOf.IsSet() {
		toSerialize["multiple_of"] = o.MultipleOf.Get()
	}
	if o.Regex.IsSet() {
		toSerialize["regex"] = o.Regex.Get()
	}
	if o.UniqueItems.IsSet() {
		toSerialize["unique_items"] = o.UniqueItems.Get()
	}
	return toSerialize, nil
}

type NullableSystemConfigurationSchema struct {
	value *SystemConfigurationSchema
	isSet bool
}

func (v NullableSystemConfigurationSchema) Get() *SystemConfigurationSchema {
	return v.value
}

func (v *NullableSystemConfigurationSchema) Set(val *SystemConfigurationSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemConfigurationSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemConfigurationSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemConfigurationSchema(val *SystemConfigurationSchema) *NullableSystemConfigurationSchema {
	return &NullableSystemConfigurationSchema{value: val, isSet: true}
}

func (v NullableSystemConfigurationSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemConfigurationSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


