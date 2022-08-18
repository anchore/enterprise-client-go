# JsonPatchRemove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Operation ID, referenced for ordering in the | [optional] 
**Op** | **string** |  | 
**Path** | **string** | A JSONPointer per RFC6901 | 

## Methods

### NewJsonPatchRemove

`func NewJsonPatchRemove(op string, path string, ) *JsonPatchRemove`

NewJsonPatchRemove instantiates a new JsonPatchRemove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonPatchRemoveWithDefaults

`func NewJsonPatchRemoveWithDefaults() *JsonPatchRemove`

NewJsonPatchRemoveWithDefaults instantiates a new JsonPatchRemove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JsonPatchRemove) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonPatchRemove) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonPatchRemove) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JsonPatchRemove) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOp

`func (o *JsonPatchRemove) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JsonPatchRemove) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JsonPatchRemove) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *JsonPatchRemove) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JsonPatchRemove) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JsonPatchRemove) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


