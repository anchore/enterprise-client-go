# ImportPackageRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | **string** |  | 
**Child** | **string** |  | 
**Type** | **string** |  | 
**Metadata** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewImportPackageRelationship

`func NewImportPackageRelationship(parent string, child string, type_ string, ) *ImportPackageRelationship`

NewImportPackageRelationship instantiates a new ImportPackageRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportPackageRelationshipWithDefaults

`func NewImportPackageRelationshipWithDefaults() *ImportPackageRelationship`

NewImportPackageRelationshipWithDefaults instantiates a new ImportPackageRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *ImportPackageRelationship) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ImportPackageRelationship) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ImportPackageRelationship) SetParent(v string)`

SetParent sets Parent field to given value.


### GetChild

`func (o *ImportPackageRelationship) GetChild() string`

GetChild returns the Child field if non-nil, zero value otherwise.

### GetChildOk

`func (o *ImportPackageRelationship) GetChildOk() (*string, bool)`

GetChildOk returns a tuple with the Child field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChild

`func (o *ImportPackageRelationship) SetChild(v string)`

SetChild sets Child field to given value.


### GetType

`func (o *ImportPackageRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportPackageRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportPackageRelationship) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *ImportPackageRelationship) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImportPackageRelationship) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImportPackageRelationship) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ImportPackageRelationship) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


