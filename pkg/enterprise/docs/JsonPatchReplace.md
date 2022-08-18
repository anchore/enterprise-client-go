# JsonPatchReplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Operation ID, referenced for ordering in the | [optional] 
**Op** | **string** |  | 
**Path** | **string** | A JSONPointer per RFC6901 | 
**Value** | **interface{}** | A valid json value, can be any valid json type | 

## Methods

### NewJsonPatchReplace

`func NewJsonPatchReplace(op string, path string, value interface{}, ) *JsonPatchReplace`

NewJsonPatchReplace instantiates a new JsonPatchReplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonPatchReplaceWithDefaults

`func NewJsonPatchReplaceWithDefaults() *JsonPatchReplace`

NewJsonPatchReplaceWithDefaults instantiates a new JsonPatchReplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JsonPatchReplace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonPatchReplace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonPatchReplace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JsonPatchReplace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOp

`func (o *JsonPatchReplace) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JsonPatchReplace) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JsonPatchReplace) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *JsonPatchReplace) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JsonPatchReplace) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JsonPatchReplace) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *JsonPatchReplace) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JsonPatchReplace) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JsonPatchReplace) SetValue(v interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


