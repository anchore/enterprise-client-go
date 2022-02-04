# CorrectionFieldMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | The package field name to match | 
**FieldValue** | **string** | The package field value for the corresponding field_name above to match. If field_name corresponds to a list value, this will search the list | 

## Methods

### NewCorrectionFieldMatch

`func NewCorrectionFieldMatch(fieldName string, fieldValue string, ) *CorrectionFieldMatch`

NewCorrectionFieldMatch instantiates a new CorrectionFieldMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrectionFieldMatchWithDefaults

`func NewCorrectionFieldMatchWithDefaults() *CorrectionFieldMatch`

NewCorrectionFieldMatchWithDefaults instantiates a new CorrectionFieldMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *CorrectionFieldMatch) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *CorrectionFieldMatch) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *CorrectionFieldMatch) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetFieldValue

`func (o *CorrectionFieldMatch) GetFieldValue() string`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *CorrectionFieldMatch) GetFieldValueOk() (*string, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *CorrectionFieldMatch) SetFieldValue(v string)`

SetFieldValue sets FieldValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


