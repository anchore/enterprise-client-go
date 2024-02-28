/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.3.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// FilesContent struct for FilesContent
type FilesContent struct {
	Filename *string `json:"filename,omitempty"`
	Gid *int32 `json:"gid,omitempty"`
	Linkdest NullableString `json:"linkdest,omitempty"`
	Mode *string `json:"mode,omitempty"`
	Sha256 NullableString `json:"sha256,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Type *string `json:"type,omitempty"`
	Uid *int32 `json:"uid,omitempty"`
}

// NewFilesContent instantiates a new FilesContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesContent() *FilesContent {
	this := FilesContent{}
	return &this
}

// NewFilesContentWithDefaults instantiates a new FilesContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesContentWithDefaults() *FilesContent {
	this := FilesContent{}
	return &this
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *FilesContent) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesContent) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *FilesContent) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *FilesContent) SetFilename(v string) {
	o.Filename = &v
}

// GetGid returns the Gid field value if set, zero value otherwise.
func (o *FilesContent) GetGid() int32 {
	if o == nil || o.Gid == nil {
		var ret int32
		return ret
	}
	return *o.Gid
}

// GetGidOk returns a tuple with the Gid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesContent) GetGidOk() (*int32, bool) {
	if o == nil || o.Gid == nil {
		return nil, false
	}
	return o.Gid, true
}

// HasGid returns a boolean if a field has been set.
func (o *FilesContent) HasGid() bool {
	if o != nil && o.Gid != nil {
		return true
	}

	return false
}

// SetGid gets a reference to the given int32 and assigns it to the Gid field.
func (o *FilesContent) SetGid(v int32) {
	o.Gid = &v
}

// GetLinkdest returns the Linkdest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesContent) GetLinkdest() string {
	if o == nil || o.Linkdest.Get() == nil {
		var ret string
		return ret
	}
	return *o.Linkdest.Get()
}

// GetLinkdestOk returns a tuple with the Linkdest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesContent) GetLinkdestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Linkdest.Get(), o.Linkdest.IsSet()
}

// HasLinkdest returns a boolean if a field has been set.
func (o *FilesContent) HasLinkdest() bool {
	if o != nil && o.Linkdest.IsSet() {
		return true
	}

	return false
}

// SetLinkdest gets a reference to the given NullableString and assigns it to the Linkdest field.
func (o *FilesContent) SetLinkdest(v string) {
	o.Linkdest.Set(&v)
}
// SetLinkdestNil sets the value for Linkdest to be an explicit nil
func (o *FilesContent) SetLinkdestNil() {
	o.Linkdest.Set(nil)
}

// UnsetLinkdest ensures that no value is present for Linkdest, not even an explicit nil
func (o *FilesContent) UnsetLinkdest() {
	o.Linkdest.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *FilesContent) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesContent) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *FilesContent) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *FilesContent) SetMode(v string) {
	o.Mode = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesContent) GetSha256() string {
	if o == nil || o.Sha256.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sha256.Get()
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesContent) GetSha256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha256.Get(), o.Sha256.IsSet()
}

// HasSha256 returns a boolean if a field has been set.
func (o *FilesContent) HasSha256() bool {
	if o != nil && o.Sha256.IsSet() {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given NullableString and assigns it to the Sha256 field.
func (o *FilesContent) SetSha256(v string) {
	o.Sha256.Set(&v)
}
// SetSha256Nil sets the value for Sha256 to be an explicit nil
func (o *FilesContent) SetSha256Nil() {
	o.Sha256.Set(nil)
}

// UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
func (o *FilesContent) UnsetSha256() {
	o.Sha256.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FilesContent) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesContent) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FilesContent) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *FilesContent) SetSize(v int32) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FilesContent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesContent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FilesContent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FilesContent) SetType(v string) {
	o.Type = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *FilesContent) GetUid() int32 {
	if o == nil || o.Uid == nil {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesContent) GetUidOk() (*int32, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *FilesContent) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *FilesContent) SetUid(v int32) {
	o.Uid = &v
}

func (o FilesContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.Gid != nil {
		toSerialize["gid"] = o.Gid
	}
	if o.Linkdest.IsSet() {
		toSerialize["linkdest"] = o.Linkdest.Get()
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Sha256.IsSet() {
		toSerialize["sha256"] = o.Sha256.Get()
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	return json.Marshal(toSerialize)
}

type NullableFilesContent struct {
	value *FilesContent
	isSet bool
}

func (v NullableFilesContent) Get() *FilesContent {
	return v.value
}

func (v *NullableFilesContent) Set(val *FilesContent) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesContent) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesContent(val *FilesContent) *NullableFilesContent {
	return &NullableFilesContent{value: val, isSet: true}
}

func (v NullableFilesContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


