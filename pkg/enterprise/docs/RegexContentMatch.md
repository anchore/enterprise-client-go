# RegexContentMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name associated with the regular expression | [optional] 
**Regex** | Pointer to **string** | The regular expression used for the match | [optional] 
**Lines** | Pointer to **[]int64** | A list of line numbers in the file that matched the regex | [optional] 

## Methods

### NewRegexContentMatch

`func NewRegexContentMatch() *RegexContentMatch`

NewRegexContentMatch instantiates a new RegexContentMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegexContentMatchWithDefaults

`func NewRegexContentMatchWithDefaults() *RegexContentMatch`

NewRegexContentMatchWithDefaults instantiates a new RegexContentMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegexContentMatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegexContentMatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegexContentMatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegexContentMatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegex

`func (o *RegexContentMatch) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *RegexContentMatch) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *RegexContentMatch) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *RegexContentMatch) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetLines

`func (o *RegexContentMatch) GetLines() []int64`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *RegexContentMatch) GetLinesOk() (*[]int64, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *RegexContentMatch) SetLines(v []int64)`

SetLines sets Lines field to given value.

### HasLines

`func (o *RegexContentMatch) HasLines() bool`

HasLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


