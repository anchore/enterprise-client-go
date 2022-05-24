# ArtifactAssociationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationId** | Pointer to **string** | The id of the association between the application version and the artifact | [optional] 
**CreatedAt** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the artifact was associated with the application version | [optional] 
**LastUpdated** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the artifact association was last updated | [optional] 

## Methods

### NewArtifactAssociationMetadata

`func NewArtifactAssociationMetadata() *ArtifactAssociationMetadata`

NewArtifactAssociationMetadata instantiates a new ArtifactAssociationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactAssociationMetadataWithDefaults

`func NewArtifactAssociationMetadataWithDefaults() *ArtifactAssociationMetadata`

NewArtifactAssociationMetadataWithDefaults instantiates a new ArtifactAssociationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationId

`func (o *ArtifactAssociationMetadata) GetAssociationId() string`

GetAssociationId returns the AssociationId field if non-nil, zero value otherwise.

### GetAssociationIdOk

`func (o *ArtifactAssociationMetadata) GetAssociationIdOk() (*string, bool)`

GetAssociationIdOk returns a tuple with the AssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationId

`func (o *ArtifactAssociationMetadata) SetAssociationId(v string)`

SetAssociationId sets AssociationId field to given value.

### HasAssociationId

`func (o *ArtifactAssociationMetadata) HasAssociationId() bool`

HasAssociationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArtifactAssociationMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactAssociationMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactAssociationMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactAssociationMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ArtifactAssociationMetadata) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ArtifactAssociationMetadata) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ArtifactAssociationMetadata) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ArtifactAssociationMetadata) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


