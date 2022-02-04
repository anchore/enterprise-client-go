# ArtifactAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactType** | **string** | The type of the artifact | 
**ArtifactKeys** | **interface{}** | A json with key-pair values to query on | 

## Methods

### NewArtifactAssociationRequest

`func NewArtifactAssociationRequest(artifactType string, artifactKeys interface{}, ) *ArtifactAssociationRequest`

NewArtifactAssociationRequest instantiates a new ArtifactAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactAssociationRequestWithDefaults

`func NewArtifactAssociationRequestWithDefaults() *ArtifactAssociationRequest`

NewArtifactAssociationRequestWithDefaults instantiates a new ArtifactAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactType

`func (o *ArtifactAssociationRequest) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *ArtifactAssociationRequest) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *ArtifactAssociationRequest) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.


### GetArtifactKeys

`func (o *ArtifactAssociationRequest) GetArtifactKeys() interface{}`

GetArtifactKeys returns the ArtifactKeys field if non-nil, zero value otherwise.

### GetArtifactKeysOk

`func (o *ArtifactAssociationRequest) GetArtifactKeysOk() (*interface{}, bool)`

GetArtifactKeysOk returns a tuple with the ArtifactKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactKeys

`func (o *ArtifactAssociationRequest) SetArtifactKeys(v interface{})`

SetArtifactKeys sets ArtifactKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


