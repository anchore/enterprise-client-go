# CustomJsonPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | Pointer to **[]string** | Ordered list of the operations in the type-specific lists. This imparts the total ordering of patches to apply such that they can be moved into a single array. This is a workaround for &#39;oneOf&#39; support in OpenAPI 2.0 | [optional] 
**Add** | Pointer to [**[]JsonPatchAdd**](JsonPatchAdd.md) |  | [optional] 
**Remove** | Pointer to [**[]JsonPatchRemove**](JsonPatchRemove.md) |  | [optional] 
**Replace** | Pointer to [**[]JsonPatchReplace**](JsonPatchReplace.md) |  | [optional] 
**Move** | Pointer to [**[]JsonPatchMove**](JsonPatchMove.md) |  | [optional] 
**Copy** | Pointer to [**[]JsonPatchCopy**](JsonPatchCopy.md) |  | [optional] 
**Test** | Pointer to [**[]JsonPatchTest**](JsonPatchTest.md) |  | [optional] 

## Methods

### NewCustomJsonPatch

`func NewCustomJsonPatch() *CustomJsonPatch`

NewCustomJsonPatch instantiates a new CustomJsonPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomJsonPatchWithDefaults

`func NewCustomJsonPatchWithDefaults() *CustomJsonPatch`

NewCustomJsonPatchWithDefaults instantiates a new CustomJsonPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *CustomJsonPatch) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *CustomJsonPatch) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *CustomJsonPatch) SetOperations(v []string)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *CustomJsonPatch) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetAdd

`func (o *CustomJsonPatch) GetAdd() []JsonPatchAdd`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *CustomJsonPatch) GetAddOk() (*[]JsonPatchAdd, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *CustomJsonPatch) SetAdd(v []JsonPatchAdd)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *CustomJsonPatch) HasAdd() bool`

HasAdd returns a boolean if a field has been set.

### GetRemove

`func (o *CustomJsonPatch) GetRemove() []JsonPatchRemove`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *CustomJsonPatch) GetRemoveOk() (*[]JsonPatchRemove, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *CustomJsonPatch) SetRemove(v []JsonPatchRemove)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *CustomJsonPatch) HasRemove() bool`

HasRemove returns a boolean if a field has been set.

### GetReplace

`func (o *CustomJsonPatch) GetReplace() []JsonPatchReplace`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *CustomJsonPatch) GetReplaceOk() (*[]JsonPatchReplace, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *CustomJsonPatch) SetReplace(v []JsonPatchReplace)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *CustomJsonPatch) HasReplace() bool`

HasReplace returns a boolean if a field has been set.

### GetMove

`func (o *CustomJsonPatch) GetMove() []JsonPatchMove`

GetMove returns the Move field if non-nil, zero value otherwise.

### GetMoveOk

`func (o *CustomJsonPatch) GetMoveOk() (*[]JsonPatchMove, bool)`

GetMoveOk returns a tuple with the Move field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMove

`func (o *CustomJsonPatch) SetMove(v []JsonPatchMove)`

SetMove sets Move field to given value.

### HasMove

`func (o *CustomJsonPatch) HasMove() bool`

HasMove returns a boolean if a field has been set.

### GetCopy

`func (o *CustomJsonPatch) GetCopy() []JsonPatchCopy`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *CustomJsonPatch) GetCopyOk() (*[]JsonPatchCopy, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *CustomJsonPatch) SetCopy(v []JsonPatchCopy)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *CustomJsonPatch) HasCopy() bool`

HasCopy returns a boolean if a field has been set.

### GetTest

`func (o *CustomJsonPatch) GetTest() []JsonPatchTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CustomJsonPatch) GetTestOk() (*[]JsonPatchTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CustomJsonPatch) SetTest(v []JsonPatchTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *CustomJsonPatch) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


