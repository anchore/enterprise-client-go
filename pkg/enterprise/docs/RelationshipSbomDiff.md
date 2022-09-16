# RelationshipSbomDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceOnly** | Pointer to [**[]Package**](Package.md) | Packages added based on the type of relationship. A \&quot;contains\&quot; relationship means packages present in the source artifact (image) not present in the target (source repo) of the relationship. | [optional] 
**TargetOnly** | Pointer to [**[]Package**](Package.md) | Packages removed based on the type of relationship. A \&quot;contains\&quot; relationship means packages not present in the source artifact (image) present in the target (source repo) of the relationship. | [optional] 
**SourceModified** | Pointer to [**[]ModifiedPackage**](ModifiedPackage.md) |  | [optional] 
**BothUnmodified** | Pointer to [**[]Package**](Package.md) |  | [optional] 

## Methods

### NewRelationshipSbomDiff

`func NewRelationshipSbomDiff() *RelationshipSbomDiff`

NewRelationshipSbomDiff instantiates a new RelationshipSbomDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipSbomDiffWithDefaults

`func NewRelationshipSbomDiffWithDefaults() *RelationshipSbomDiff`

NewRelationshipSbomDiffWithDefaults instantiates a new RelationshipSbomDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceOnly

`func (o *RelationshipSbomDiff) GetSourceOnly() []Package`

GetSourceOnly returns the SourceOnly field if non-nil, zero value otherwise.

### GetSourceOnlyOk

`func (o *RelationshipSbomDiff) GetSourceOnlyOk() (*[]Package, bool)`

GetSourceOnlyOk returns a tuple with the SourceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOnly

`func (o *RelationshipSbomDiff) SetSourceOnly(v []Package)`

SetSourceOnly sets SourceOnly field to given value.

### HasSourceOnly

`func (o *RelationshipSbomDiff) HasSourceOnly() bool`

HasSourceOnly returns a boolean if a field has been set.

### GetTargetOnly

`func (o *RelationshipSbomDiff) GetTargetOnly() []Package`

GetTargetOnly returns the TargetOnly field if non-nil, zero value otherwise.

### GetTargetOnlyOk

`func (o *RelationshipSbomDiff) GetTargetOnlyOk() (*[]Package, bool)`

GetTargetOnlyOk returns a tuple with the TargetOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOnly

`func (o *RelationshipSbomDiff) SetTargetOnly(v []Package)`

SetTargetOnly sets TargetOnly field to given value.

### HasTargetOnly

`func (o *RelationshipSbomDiff) HasTargetOnly() bool`

HasTargetOnly returns a boolean if a field has been set.

### GetSourceModified

`func (o *RelationshipSbomDiff) GetSourceModified() []ModifiedPackage`

GetSourceModified returns the SourceModified field if non-nil, zero value otherwise.

### GetSourceModifiedOk

`func (o *RelationshipSbomDiff) GetSourceModifiedOk() (*[]ModifiedPackage, bool)`

GetSourceModifiedOk returns a tuple with the SourceModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceModified

`func (o *RelationshipSbomDiff) SetSourceModified(v []ModifiedPackage)`

SetSourceModified sets SourceModified field to given value.

### HasSourceModified

`func (o *RelationshipSbomDiff) HasSourceModified() bool`

HasSourceModified returns a boolean if a field has been set.

### GetBothUnmodified

`func (o *RelationshipSbomDiff) GetBothUnmodified() []Package`

GetBothUnmodified returns the BothUnmodified field if non-nil, zero value otherwise.

### GetBothUnmodifiedOk

`func (o *RelationshipSbomDiff) GetBothUnmodifiedOk() (*[]Package, bool)`

GetBothUnmodifiedOk returns a tuple with the BothUnmodified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBothUnmodified

`func (o *RelationshipSbomDiff) SetBothUnmodified(v []Package)`

SetBothUnmodified sets BothUnmodified field to given value.

### HasBothUnmodified

`func (o *RelationshipSbomDiff) HasBothUnmodified() bool`

HasBothUnmodified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


