# SecretSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Matches** | Pointer to [**[]RegexContentMatch**](RegexContentMatch.md) |  | [optional] 

## Methods

### NewSecretSearchResult

`func NewSecretSearchResult() *SecretSearchResult`

NewSecretSearchResult instantiates a new SecretSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretSearchResultWithDefaults

`func NewSecretSearchResultWithDefaults() *SecretSearchResult`

NewSecretSearchResultWithDefaults instantiates a new SecretSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *SecretSearchResult) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SecretSearchResult) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SecretSearchResult) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SecretSearchResult) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMatches

`func (o *SecretSearchResult) GetMatches() []RegexContentMatch`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *SecretSearchResult) GetMatchesOk() (*[]RegexContentMatch, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *SecretSearchResult) SetMatches(v []RegexContentMatch)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *SecretSearchResult) HasMatches() bool`

HasMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


