# ModifiedPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**Package**](Package.md) |  | [optional] 
**Target** | Pointer to [**Package**](Package.md) |  | [optional] 
**Patch** | Pointer to [**CustomJsonPatch**](CustomJsonPatch.md) |  | [optional] 

## Methods

### NewModifiedPackage

`func NewModifiedPackage() *ModifiedPackage`

NewModifiedPackage instantiates a new ModifiedPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifiedPackageWithDefaults

`func NewModifiedPackageWithDefaults() *ModifiedPackage`

NewModifiedPackageWithDefaults instantiates a new ModifiedPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ModifiedPackage) GetSource() Package`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ModifiedPackage) GetSourceOk() (*Package, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ModifiedPackage) SetSource(v Package)`

SetSource sets Source field to given value.

### HasSource

`func (o *ModifiedPackage) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *ModifiedPackage) GetTarget() Package`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ModifiedPackage) GetTargetOk() (*Package, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ModifiedPackage) SetTarget(v Package)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ModifiedPackage) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetPatch

`func (o *ModifiedPackage) GetPatch() CustomJsonPatch`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *ModifiedPackage) GetPatchOk() (*CustomJsonPatch, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *ModifiedPackage) SetPatch(v CustomJsonPatch)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *ModifiedPackage) HasPatch() bool`

HasPatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


