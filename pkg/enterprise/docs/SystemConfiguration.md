# SystemConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Key** | **string** |  | 
**Value** | [**SystemConfigurationValue**](SystemConfigurationValue.md) |  | 
**RequiresSystemRestart** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ConfigSchema** | Pointer to [**SystemConfigurationSchema**](SystemConfigurationSchema.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**IsEditable** | Pointer to **bool** |  | [optional] 
**EditableReason** | Pointer to **string** |  | [optional] 
**IsSecret** | Pointer to **bool** |  | [optional] 
**IsDeprecated** | Pointer to **bool** |  | [optional] 
**DeprecatedReason** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewSystemConfiguration

`func NewSystemConfiguration(key string, value SystemConfigurationValue, ) *SystemConfiguration`

NewSystemConfiguration instantiates a new SystemConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemConfigurationWithDefaults

`func NewSystemConfigurationWithDefaults() *SystemConfiguration`

NewSystemConfigurationWithDefaults instantiates a new SystemConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SystemConfiguration) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SystemConfiguration) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SystemConfiguration) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SystemConfiguration) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetKey

`func (o *SystemConfiguration) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SystemConfiguration) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SystemConfiguration) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *SystemConfiguration) GetValue() SystemConfigurationValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SystemConfiguration) GetValueOk() (*SystemConfigurationValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SystemConfiguration) SetValue(v SystemConfigurationValue)`

SetValue sets Value field to given value.


### GetRequiresSystemRestart

`func (o *SystemConfiguration) GetRequiresSystemRestart() bool`

GetRequiresSystemRestart returns the RequiresSystemRestart field if non-nil, zero value otherwise.

### GetRequiresSystemRestartOk

`func (o *SystemConfiguration) GetRequiresSystemRestartOk() (*bool, bool)`

GetRequiresSystemRestartOk returns a tuple with the RequiresSystemRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresSystemRestart

`func (o *SystemConfiguration) SetRequiresSystemRestart(v bool)`

SetRequiresSystemRestart sets RequiresSystemRestart field to given value.

### HasRequiresSystemRestart

`func (o *SystemConfiguration) HasRequiresSystemRestart() bool`

HasRequiresSystemRestart returns a boolean if a field has been set.

### GetCategory

`func (o *SystemConfiguration) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SystemConfiguration) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SystemConfiguration) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SystemConfiguration) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfigSchema

`func (o *SystemConfiguration) GetConfigSchema() SystemConfigurationSchema`

GetConfigSchema returns the ConfigSchema field if non-nil, zero value otherwise.

### GetConfigSchemaOk

`func (o *SystemConfiguration) GetConfigSchemaOk() (*SystemConfigurationSchema, bool)`

GetConfigSchemaOk returns a tuple with the ConfigSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSchema

`func (o *SystemConfiguration) SetConfigSchema(v SystemConfigurationSchema)`

SetConfigSchema sets ConfigSchema field to given value.

### HasConfigSchema

`func (o *SystemConfiguration) HasConfigSchema() bool`

HasConfigSchema returns a boolean if a field has been set.

### GetTitle

`func (o *SystemConfiguration) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SystemConfiguration) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SystemConfiguration) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SystemConfiguration) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *SystemConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValue

`func (o *SystemConfiguration) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SystemConfiguration) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SystemConfiguration) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SystemConfiguration) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetSource

`func (o *SystemConfiguration) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SystemConfiguration) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SystemConfiguration) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SystemConfiguration) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetIsEditable

`func (o *SystemConfiguration) GetIsEditable() bool`

GetIsEditable returns the IsEditable field if non-nil, zero value otherwise.

### GetIsEditableOk

`func (o *SystemConfiguration) GetIsEditableOk() (*bool, bool)`

GetIsEditableOk returns a tuple with the IsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEditable

`func (o *SystemConfiguration) SetIsEditable(v bool)`

SetIsEditable sets IsEditable field to given value.

### HasIsEditable

`func (o *SystemConfiguration) HasIsEditable() bool`

HasIsEditable returns a boolean if a field has been set.

### GetEditableReason

`func (o *SystemConfiguration) GetEditableReason() string`

GetEditableReason returns the EditableReason field if non-nil, zero value otherwise.

### GetEditableReasonOk

`func (o *SystemConfiguration) GetEditableReasonOk() (*string, bool)`

GetEditableReasonOk returns a tuple with the EditableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableReason

`func (o *SystemConfiguration) SetEditableReason(v string)`

SetEditableReason sets EditableReason field to given value.

### HasEditableReason

`func (o *SystemConfiguration) HasEditableReason() bool`

HasEditableReason returns a boolean if a field has been set.

### GetIsSecret

`func (o *SystemConfiguration) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *SystemConfiguration) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *SystemConfiguration) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *SystemConfiguration) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.

### GetIsDeprecated

`func (o *SystemConfiguration) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *SystemConfiguration) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *SystemConfiguration) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.

### HasIsDeprecated

`func (o *SystemConfiguration) HasIsDeprecated() bool`

HasIsDeprecated returns a boolean if a field has been set.

### GetDeprecatedReason

`func (o *SystemConfiguration) GetDeprecatedReason() string`

GetDeprecatedReason returns the DeprecatedReason field if non-nil, zero value otherwise.

### GetDeprecatedReasonOk

`func (o *SystemConfiguration) GetDeprecatedReasonOk() (*string, bool)`

GetDeprecatedReasonOk returns a tuple with the DeprecatedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedReason

`func (o *SystemConfiguration) SetDeprecatedReason(v string)`

SetDeprecatedReason sets DeprecatedReason field to given value.

### HasDeprecatedReason

`func (o *SystemConfiguration) HasDeprecatedReason() bool`

HasDeprecatedReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SystemConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SystemConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SystemConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SystemConfiguration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SystemConfiguration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SystemConfiguration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SystemConfiguration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SystemConfiguration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


