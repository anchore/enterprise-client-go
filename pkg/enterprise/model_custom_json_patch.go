/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.2.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// CustomJsonPatch Anchore-specific adaptation of RFC6902 to be describeable in swagger/open-api 2.0. Each item is given an ID and the ids are ordered in the array.
type CustomJsonPatch struct {
	// Ordered list of the operations in the type-specific lists. This imparts the total ordering of patches to apply such that they can be moved into a single array. This is a workaround for 'oneOf' support in OpenAPI 2.0
	Operations *[]string `json:"operations,omitempty"`
	Add *[]JsonPatchAdd `json:"add,omitempty"`
	Remove *[]JsonPatchRemove `json:"remove,omitempty"`
	Replace *[]JsonPatchReplace `json:"replace,omitempty"`
	Move *[]JsonPatchMove `json:"move,omitempty"`
	Copy *[]JsonPatchCopy `json:"copy,omitempty"`
	Test *[]JsonPatchTest `json:"test,omitempty"`
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
	if o == nil || o.Operations == nil {
		var ret []string
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetOperationsOk() (*[]string, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []string and assigns it to the Operations field.
func (o *CustomJsonPatch) SetOperations(v []string) {
	o.Operations = &v
}

// GetAdd returns the Add field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetAdd() []JsonPatchAdd {
	if o == nil || o.Add == nil {
		var ret []JsonPatchAdd
		return ret
	}
	return *o.Add
}

// GetAddOk returns a tuple with the Add field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetAddOk() (*[]JsonPatchAdd, bool) {
	if o == nil || o.Add == nil {
		return nil, false
	}
	return o.Add, true
}

// HasAdd returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasAdd() bool {
	if o != nil && o.Add != nil {
		return true
	}

	return false
}

// SetAdd gets a reference to the given []JsonPatchAdd and assigns it to the Add field.
func (o *CustomJsonPatch) SetAdd(v []JsonPatchAdd) {
	o.Add = &v
}

// GetRemove returns the Remove field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetRemove() []JsonPatchRemove {
	if o == nil || o.Remove == nil {
		var ret []JsonPatchRemove
		return ret
	}
	return *o.Remove
}

// GetRemoveOk returns a tuple with the Remove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetRemoveOk() (*[]JsonPatchRemove, bool) {
	if o == nil || o.Remove == nil {
		return nil, false
	}
	return o.Remove, true
}

// HasRemove returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasRemove() bool {
	if o != nil && o.Remove != nil {
		return true
	}

	return false
}

// SetRemove gets a reference to the given []JsonPatchRemove and assigns it to the Remove field.
func (o *CustomJsonPatch) SetRemove(v []JsonPatchRemove) {
	o.Remove = &v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetReplace() []JsonPatchReplace {
	if o == nil || o.Replace == nil {
		var ret []JsonPatchReplace
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetReplaceOk() (*[]JsonPatchReplace, bool) {
	if o == nil || o.Replace == nil {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasReplace() bool {
	if o != nil && o.Replace != nil {
		return true
	}

	return false
}

// SetReplace gets a reference to the given []JsonPatchReplace and assigns it to the Replace field.
func (o *CustomJsonPatch) SetReplace(v []JsonPatchReplace) {
	o.Replace = &v
}

// GetMove returns the Move field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetMove() []JsonPatchMove {
	if o == nil || o.Move == nil {
		var ret []JsonPatchMove
		return ret
	}
	return *o.Move
}

// GetMoveOk returns a tuple with the Move field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetMoveOk() (*[]JsonPatchMove, bool) {
	if o == nil || o.Move == nil {
		return nil, false
	}
	return o.Move, true
}

// HasMove returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasMove() bool {
	if o != nil && o.Move != nil {
		return true
	}

	return false
}

// SetMove gets a reference to the given []JsonPatchMove and assigns it to the Move field.
func (o *CustomJsonPatch) SetMove(v []JsonPatchMove) {
	o.Move = &v
}

// GetCopy returns the Copy field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetCopy() []JsonPatchCopy {
	if o == nil || o.Copy == nil {
		var ret []JsonPatchCopy
		return ret
	}
	return *o.Copy
}

// GetCopyOk returns a tuple with the Copy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetCopyOk() (*[]JsonPatchCopy, bool) {
	if o == nil || o.Copy == nil {
		return nil, false
	}
	return o.Copy, true
}

// HasCopy returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasCopy() bool {
	if o != nil && o.Copy != nil {
		return true
	}

	return false
}

// SetCopy gets a reference to the given []JsonPatchCopy and assigns it to the Copy field.
func (o *CustomJsonPatch) SetCopy(v []JsonPatchCopy) {
	o.Copy = &v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *CustomJsonPatch) GetTest() []JsonPatchTest {
	if o == nil || o.Test == nil {
		var ret []JsonPatchTest
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomJsonPatch) GetTestOk() (*[]JsonPatchTest, bool) {
	if o == nil || o.Test == nil {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *CustomJsonPatch) HasTest() bool {
	if o != nil && o.Test != nil {
		return true
	}

	return false
}

// SetTest gets a reference to the given []JsonPatchTest and assigns it to the Test field.
func (o *CustomJsonPatch) SetTest(v []JsonPatchTest) {
	o.Test = &v
}

func (o CustomJsonPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Operations != nil {
		toSerialize["operations"] = o.Operations
	}
	if o.Add != nil {
		toSerialize["add"] = o.Add
	}
	if o.Remove != nil {
		toSerialize["remove"] = o.Remove
	}
	if o.Replace != nil {
		toSerialize["replace"] = o.Replace
	}
	if o.Move != nil {
		toSerialize["move"] = o.Move
	}
	if o.Copy != nil {
		toSerialize["copy"] = o.Copy
	}
	if o.Test != nil {
		toSerialize["test"] = o.Test
	}
	return json.Marshal(toSerialize)
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


