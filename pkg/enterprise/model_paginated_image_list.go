/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// PaginatedImageList Pagination wrapped list of images that match some filter
type PaginatedImageList struct {
	// The page number returned (should match the requested page query string param)
	Page *string `json:"page,omitempty"`
	// True if additional pages exist (page + 1) or False if this is the last page
	NextPage *string `json:"next_page,omitempty"`
	// The number of items sent in this response
	ReturnedCount *int32 `json:"returned_count,omitempty"`
	Images []ImageWithPackages `json:"images,omitempty"`
}

// NewPaginatedImageList instantiates a new PaginatedImageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedImageList() *PaginatedImageList {
	this := PaginatedImageList{}
	return &this
}

// NewPaginatedImageListWithDefaults instantiates a new PaginatedImageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedImageListWithDefaults() *PaginatedImageList {
	this := PaginatedImageList{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginatedImageList) GetPage() string {
	if o == nil || o.Page == nil {
		var ret string
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedImageList) GetPageOk() (*string, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginatedImageList) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *PaginatedImageList) SetPage(v string) {
	o.Page = &v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *PaginatedImageList) GetNextPage() string {
	if o == nil || o.NextPage == nil {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedImageList) GetNextPageOk() (*string, bool) {
	if o == nil || o.NextPage == nil {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *PaginatedImageList) HasNextPage() bool {
	if o != nil && o.NextPage != nil {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *PaginatedImageList) SetNextPage(v string) {
	o.NextPage = &v
}

// GetReturnedCount returns the ReturnedCount field value if set, zero value otherwise.
func (o *PaginatedImageList) GetReturnedCount() int32 {
	if o == nil || o.ReturnedCount == nil {
		var ret int32
		return ret
	}
	return *o.ReturnedCount
}

// GetReturnedCountOk returns a tuple with the ReturnedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedImageList) GetReturnedCountOk() (*int32, bool) {
	if o == nil || o.ReturnedCount == nil {
		return nil, false
	}
	return o.ReturnedCount, true
}

// HasReturnedCount returns a boolean if a field has been set.
func (o *PaginatedImageList) HasReturnedCount() bool {
	if o != nil && o.ReturnedCount != nil {
		return true
	}

	return false
}

// SetReturnedCount gets a reference to the given int32 and assigns it to the ReturnedCount field.
func (o *PaginatedImageList) SetReturnedCount(v int32) {
	o.ReturnedCount = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *PaginatedImageList) GetImages() []ImageWithPackages {
	if o == nil || o.Images == nil {
		var ret []ImageWithPackages
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedImageList) GetImagesOk() ([]ImageWithPackages, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *PaginatedImageList) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []ImageWithPackages and assigns it to the Images field.
func (o *PaginatedImageList) SetImages(v []ImageWithPackages) {
	o.Images = v
}

func (o PaginatedImageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.NextPage != nil {
		toSerialize["next_page"] = o.NextPage
	}
	if o.ReturnedCount != nil {
		toSerialize["returned_count"] = o.ReturnedCount
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedImageList struct {
	value *PaginatedImageList
	isSet bool
}

func (v NullablePaginatedImageList) Get() *PaginatedImageList {
	return v.value
}

func (v *NullablePaginatedImageList) Set(val *PaginatedImageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedImageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedImageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedImageList(val *PaginatedImageList) *NullablePaginatedImageList {
	return &NullablePaginatedImageList{value: val, isSet: true}
}

func (v NullablePaginatedImageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedImageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


