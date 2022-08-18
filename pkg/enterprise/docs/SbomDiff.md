# SbomDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | Pointer to [**[]Package**](Package.md) | Packages added based on the type of relationship. A \&quot;contains\&quot; relationship means packages present in the source artifact (image) not present in the target (source repo) of the relationship. | [optional] 
**Removed** | Pointer to [**[]Package**](Package.md) | Packages removed based on the type of relationship. A \&quot;contains\&quot; relationship means packages not present in the source artifact (image) present in the target (source repo) of the relationship. | [optional] 
**Modified** | Pointer to [**[]ModifiedPackage**](ModifiedPackage.md) |  | [optional] 
**Unmodified** | Pointer to [**[]Package**](Package.md) |  | [optional] 

## Methods

### NewSbomDiff

`func NewSbomDiff() *SbomDiff`

NewSbomDiff instantiates a new SbomDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSbomDiffWithDefaults

`func NewSbomDiffWithDefaults() *SbomDiff`

NewSbomDiffWithDefaults instantiates a new SbomDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *SbomDiff) GetAdded() []Package`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *SbomDiff) GetAddedOk() (*[]Package, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *SbomDiff) SetAdded(v []Package)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *SbomDiff) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetRemoved

`func (o *SbomDiff) GetRemoved() []Package`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *SbomDiff) GetRemovedOk() (*[]Package, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *SbomDiff) SetRemoved(v []Package)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *SbomDiff) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### GetModified

`func (o *SbomDiff) GetModified() []ModifiedPackage`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *SbomDiff) GetModifiedOk() (*[]ModifiedPackage, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *SbomDiff) SetModified(v []ModifiedPackage)`

SetModified sets Modified field to given value.

### HasModified

`func (o *SbomDiff) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetUnmodified

`func (o *SbomDiff) GetUnmodified() []Package`

GetUnmodified returns the Unmodified field if non-nil, zero value otherwise.

### GetUnmodifiedOk

`func (o *SbomDiff) GetUnmodifiedOk() (*[]Package, bool)`

GetUnmodifiedOk returns a tuple with the Unmodified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmodified

`func (o *SbomDiff) SetUnmodified(v []Package)`

SetUnmodified sets Unmodified field to given value.

### HasUnmodified

`func (o *SbomDiff) HasUnmodified() bool`

HasUnmodified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


