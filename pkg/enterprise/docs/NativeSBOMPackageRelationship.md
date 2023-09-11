# NativeSBOMPackageRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | **string** |  | 
**Child** | **string** |  | 
**Type** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewNativeSBOMPackageRelationship

`func NewNativeSBOMPackageRelationship(parent string, child string, type_ string, ) *NativeSBOMPackageRelationship`

NewNativeSBOMPackageRelationship instantiates a new NativeSBOMPackageRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNativeSBOMPackageRelationshipWithDefaults

`func NewNativeSBOMPackageRelationshipWithDefaults() *NativeSBOMPackageRelationship`

NewNativeSBOMPackageRelationshipWithDefaults instantiates a new NativeSBOMPackageRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *NativeSBOMPackageRelationship) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *NativeSBOMPackageRelationship) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *NativeSBOMPackageRelationship) SetParent(v string)`

SetParent sets Parent field to given value.


### GetChild

`func (o *NativeSBOMPackageRelationship) GetChild() string`

GetChild returns the Child field if non-nil, zero value otherwise.

### GetChildOk

`func (o *NativeSBOMPackageRelationship) GetChildOk() (*string, bool)`

GetChildOk returns a tuple with the Child field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChild

`func (o *NativeSBOMPackageRelationship) SetChild(v string)`

SetChild sets Child field to given value.


### GetType

`func (o *NativeSBOMPackageRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NativeSBOMPackageRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NativeSBOMPackageRelationship) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *NativeSBOMPackageRelationship) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NativeSBOMPackageRelationship) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NativeSBOMPackageRelationship) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NativeSBOMPackageRelationship) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


