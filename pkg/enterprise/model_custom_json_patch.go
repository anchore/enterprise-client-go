/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.11.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the CustomJsonPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomJsonPatch{}

// CustomJsonPatch Anchore-specific adaptation of RFC6902 to be describeable in swagger/open-api 2.0. Each item is given an ID and the ids are ordered in the array.
type CustomJsonPatch struct {
	// Ordered list of the operations in the type-specific lists. This imparts the total ordering of patches to apply such that they can be moved into a single array.
	Operations []string `json:"operations,omitempty"`
	Add []JsonPatchAdd `json:"add,omitempty"`
	Remove []JsonPatchRemove `json:"remove,omitempty"`
	Replace []JsonPatchReplace `json:"replace,omitempty"`
	Move []JsonPatchMove `json:"move,omitempty"`
	Copy []JsonPatchCopy `json:"copy,omitempty"`
	Test []JsonPatchTest `json:"test,omitempty"`
}

// NewCustomJsonPatch instantiates a new CustomJsonPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomJsonPatch() *CustomJsonPatch {
	this := CustomJsonPatch{}
	return &this
}

// NewCustomJsonPatchWithDefaults instantiates a new CustomJsonPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomJsonPatchWithDefaults() *CustomJsonPatch {
	this := CustomJsonPatch{}
	return &this
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetOperations() []string {
	if o == nil || IsNil(o.Operations) {
		var ret []string
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetOperationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []string and assigns it to the Operations field.
func (o *CustomJsonPatch) SetOperations(v []string) {
	o.Operations = v
}

// GetAdd returns the Add field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetAdd() []JsonPatchAdd {
	if o == nil || IsNil(o.Add) {
		var ret []JsonPatchAdd
		return ret
	}
	return o.Add
}

// GetAddOk returns a tuple with the Add field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetAddOk() ([]JsonPatchAdd, bool) {
	if o == nil || IsNil(o.Add) {
		return nil, false
	}
	return o.Add, true
}

// HasAdd returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasAdd() bool {
	if o != nil && !IsNil(o.Add) {
		return true
	}

	return false
}

// SetAdd gets a reference to the given []JsonPatchAdd and assigns it to the Add field.
func (o *CustomJsonPatch) SetAdd(v []JsonPatchAdd) {
	o.Add = v
}

// GetRemove returns the Remove field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetRemove() []JsonPatchRemove {
	if o == nil || IsNil(o.Remove) {
		var ret []JsonPatchRemove
		return ret
	}
	return o.Remove
}

// GetRemoveOk returns a tuple with the Remove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetRemoveOk() ([]JsonPatchRemove, bool) {
	if o == nil || IsNil(o.Remove) {
		return nil, false
	}
	return o.Remove, true
}

// HasRemove returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasRemove() bool {
	if o != nil && !IsNil(o.Remove) {
		return true
	}

	return false
}

// SetRemove gets a reference to the given []JsonPatchRemove and assigns it to the Remove field.
func (o *CustomJsonPatch) SetRemove(v []JsonPatchRemove) {
	o.Remove = v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetReplace() []JsonPatchReplace {
	if o == nil || IsNil(o.Replace) {
		var ret []JsonPatchReplace
		return ret
	}
	return o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetReplaceOk() ([]JsonPatchReplace, bool) {
	if o == nil || IsNil(o.Replace) {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasReplace() bool {
	if o != nil && !IsNil(o.Replace) {
		return true
	}

	return false
}

// SetReplace gets a reference to the given []JsonPatchReplace and assigns it to the Replace field.
func (o *CustomJsonPatch) SetReplace(v []JsonPatchReplace) {
	o.Replace = v
}

// GetMove returns the Move field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetMove() []JsonPatchMove {
	if o == nil || IsNil(o.Move) {
		var ret []JsonPatchMove
		return ret
	}
	return o.Move
}

// GetMoveOk returns a tuple with the Move field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetMoveOk() ([]JsonPatchMove, bool) {
	if o == nil || IsNil(o.Move) {
		return nil, false
	}
	return o.Move, true
}

// HasMove returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasMove() bool {
	if o != nil && !IsNil(o.Move) {
		return true
	}

	return false
}

// SetMove gets a reference to the given []JsonPatchMove and assigns it to the Move field.
func (o *CustomJsonPatch) SetMove(v []JsonPatchMove) {
	o.Move = v
}

// GetCopy returns the Copy field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetCopy() []JsonPatchCopy {
	if o == nil || IsNil(o.Copy) {
		var ret []JsonPatchCopy
		return ret
	}
	return o.Copy
}

// GetCopyOk returns a tuple with the Copy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetCopyOk() ([]JsonPatchCopy, bool) {
	if o == nil || IsNil(o.Copy) {
		return nil, false
	}
	return o.Copy, true
}

// HasCopy returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasCopy() bool {
	if o != nil && !IsNil(o.Copy) {
		return true
	}

	return false
}

// SetCopy gets a reference to the given []JsonPatchCopy and assigns it to the Copy field.
func (o *CustomJsonPatch) SetCopy(v []JsonPatchCopy) {
	o.Copy = v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetTest() []JsonPatchTest {
	if o == nil || IsNil(o.Test) {
		var ret []JsonPatchTest
		return ret
	}
	return o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetTestOk() ([]JsonPatchTest, bool) {
	if o == nil || IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasTest() bool {
	if o != nil && !IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given []JsonPatchTest and assigns it to the Test field.
func (o *CustomJsonPatch) SetTest(v []JsonPatchTest) {
	o.Test = v
}

func (o CustomJsonPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomJsonPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	if !IsNil(o.Add) {
		toSerialize["add"] = o.Add
	}
	if !IsNil(o.Remove) {
		toSerialize["remove"] = o.Remove
	}
	if !IsNil(o.Replace) {
		toSerialize["replace"] = o.Replace
	}
	if !IsNil(o.Move) {
		toSerialize["move"] = o.Move
	}
	if !IsNil(o.Copy) {
		toSerialize["copy"] = o.Copy
	}
	if !IsNil(o.Test) {
		toSerialize["test"] = o.Test
	}
	return toSerialize, nil
}

type NullableCustomJsonPatch struct {
	value *CustomJsonPatch
	isSet bool
}

func (v NullableCustomJsonPatch) Get() *CustomJsonPatch {
	return v.value
}

func (v *NullableCustomJsonPatch) Set(val *CustomJsonPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomJsonPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomJsonPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomJsonPatch(val *CustomJsonPatch) *NullableCustomJsonPatch {
	return &NullableCustomJsonPatch{value: val, isSet: true}
}

func (v NullableCustomJsonPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomJsonPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


