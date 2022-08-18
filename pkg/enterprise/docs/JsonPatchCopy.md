# JsonPatchCopy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Operation ID, referenced for ordering in the | [optional] 
**Op** | **string** |  | 
**Path** | **string** | A JSONPointer per RFC6901 | 
**From** | **string** | A JSONPointer per RFC6901 | 

## Methods

### NewJsonPatchCopy

`func NewJsonPatchCopy(op string, path string, from string, ) *JsonPatchCopy`

NewJsonPatchCopy instantiates a new JsonPatchCopy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonPatchCopyWithDefaults

`func NewJsonPatchCopyWithDefaults() *JsonPatchCopy`

NewJsonPatchCopyWithDefaults instantiates a new JsonPatchCopy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JsonPatchCopy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonPatchCopy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonPatchCopy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JsonPatchCopy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOp

`func (o *JsonPatchCopy) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JsonPatchCopy) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JsonPatchCopy) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *JsonPatchCopy) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JsonPatchCopy) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JsonPatchCopy) SetPath(v string)`

SetPath sets Path field to given value.


### GetFrom

`func (o *JsonPatchCopy) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *JsonPatchCopy) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *JsonPatchCopy) SetFrom(v string)`

SetFrom sets From field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


