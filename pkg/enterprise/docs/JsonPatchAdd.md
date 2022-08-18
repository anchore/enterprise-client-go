# JsonPatchAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Operation ID, referenced for ordering in the | [optional] 
**Op** | **string** |  | 
**Path** | **string** | A JSONPointer per RFC6901 | 
**Value** | **interface{}** | A valid json value, can be any valid json type | 

## Methods

### NewJsonPatchAdd

`func NewJsonPatchAdd(op string, path string, value interface{}, ) *JsonPatchAdd`

NewJsonPatchAdd instantiates a new JsonPatchAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonPatchAddWithDefaults

`func NewJsonPatchAddWithDefaults() *JsonPatchAdd`

NewJsonPatchAddWithDefaults instantiates a new JsonPatchAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JsonPatchAdd) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonPatchAdd) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonPatchAdd) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JsonPatchAdd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOp

`func (o *JsonPatchAdd) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JsonPatchAdd) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JsonPatchAdd) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *JsonPatchAdd) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JsonPatchAdd) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JsonPatchAdd) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *JsonPatchAdd) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JsonPatchAdd) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JsonPatchAdd) SetValue(v interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


