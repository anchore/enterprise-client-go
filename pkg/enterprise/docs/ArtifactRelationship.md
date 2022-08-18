# ArtifactRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Source** | Pointer to [**ArtifactReference**](ArtifactReference.md) |  | [optional] 
**Target** | Pointer to [**ArtifactReference**](ArtifactReference.md) |  | [optional] 
**RelationshipType** | Pointer to [**RelationshipType**](RelationshipType.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**UserMetadata** | Pointer to **interface{}** | User-provided metadata about the relationship | [optional] 

## Methods

### NewArtifactRelationship

`func NewArtifactRelationship() *ArtifactRelationship`

NewArtifactRelationship instantiates a new ArtifactRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactRelationshipWithDefaults

`func NewArtifactRelationshipWithDefaults() *ArtifactRelationship`

NewArtifactRelationshipWithDefaults instantiates a new ArtifactRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ArtifactRelationship) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactRelationship) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactRelationship) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactRelationship) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSource

`func (o *ArtifactRelationship) GetSource() ArtifactReference`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ArtifactRelationship) GetSourceOk() (*ArtifactReference, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ArtifactRelationship) SetSource(v ArtifactReference)`

SetSource sets Source field to given value.

### HasSource

`func (o *ArtifactRelationship) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *ArtifactRelationship) GetTarget() ArtifactReference`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ArtifactRelationship) GetTargetOk() (*ArtifactReference, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ArtifactRelationship) SetTarget(v ArtifactReference)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ArtifactRelationship) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetRelationshipType

`func (o *ArtifactRelationship) GetRelationshipType() RelationshipType`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *ArtifactRelationship) GetRelationshipTypeOk() (*RelationshipType, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *ArtifactRelationship) SetRelationshipType(v RelationshipType)`

SetRelationshipType sets RelationshipType field to given value.

### HasRelationshipType

`func (o *ArtifactRelationship) HasRelationshipType() bool`

HasRelationshipType returns a boolean if a field has been set.

### GetComment

`func (o *ArtifactRelationship) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ArtifactRelationship) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ArtifactRelationship) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ArtifactRelationship) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetUserMetadata

`func (o *ArtifactRelationship) GetUserMetadata() interface{}`

GetUserMetadata returns the UserMetadata field if non-nil, zero value otherwise.

### GetUserMetadataOk

`func (o *ArtifactRelationship) GetUserMetadataOk() (*interface{}, bool)`

GetUserMetadataOk returns a tuple with the UserMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetadata

`func (o *ArtifactRelationship) SetUserMetadata(v interface{})`

SetUserMetadata sets UserMetadata field to given value.

### HasUserMetadata

`func (o *ArtifactRelationship) HasUserMetadata() bool`

HasUserMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


