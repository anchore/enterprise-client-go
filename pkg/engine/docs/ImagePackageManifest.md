# ImagePackageManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | [**[]ImportPackage**](ImportPackage.md) |  | 
**Source** | [**ImportSource**](ImportSource.md) |  | 
**Distro** | [**ImportDistribution**](ImportDistribution.md) |  | 
**Descriptor** | Pointer to [**ImportDescriptor**](ImportDescriptor.md) |  | [optional] 
**Schema** | Pointer to [**ImportSchema**](ImportSchema.md) |  | [optional] 
**ArtifactRelationships** | Pointer to [**[]ImportPackageRelationship**](ImportPackageRelationship.md) |  | [optional] 

## Methods

### NewImagePackageManifest

`func NewImagePackageManifest(artifacts []ImportPackage, source ImportSource, distro ImportDistribution, ) *ImagePackageManifest`

NewImagePackageManifest instantiates a new ImagePackageManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePackageManifestWithDefaults

`func NewImagePackageManifestWithDefaults() *ImagePackageManifest`

NewImagePackageManifestWithDefaults instantiates a new ImagePackageManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *ImagePackageManifest) GetArtifacts() []ImportPackage`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ImagePackageManifest) GetArtifactsOk() (*[]ImportPackage, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ImagePackageManifest) SetArtifacts(v []ImportPackage)`

SetArtifacts sets Artifacts field to given value.


### GetSource

`func (o *ImagePackageManifest) GetSource() ImportSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ImagePackageManifest) GetSourceOk() (*ImportSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ImagePackageManifest) SetSource(v ImportSource)`

SetSource sets Source field to given value.


### GetDistro

`func (o *ImagePackageManifest) GetDistro() ImportDistribution`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *ImagePackageManifest) GetDistroOk() (*ImportDistribution, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *ImagePackageManifest) SetDistro(v ImportDistribution)`

SetDistro sets Distro field to given value.


### GetDescriptor

`func (o *ImagePackageManifest) GetDescriptor() ImportDescriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *ImagePackageManifest) GetDescriptorOk() (*ImportDescriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *ImagePackageManifest) SetDescriptor(v ImportDescriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *ImagePackageManifest) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetSchema

`func (o *ImagePackageManifest) GetSchema() ImportSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ImagePackageManifest) GetSchemaOk() (*ImportSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ImagePackageManifest) SetSchema(v ImportSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ImagePackageManifest) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetArtifactRelationships

`func (o *ImagePackageManifest) GetArtifactRelationships() []ImportPackageRelationship`

GetArtifactRelationships returns the ArtifactRelationships field if non-nil, zero value otherwise.

### GetArtifactRelationshipsOk

`func (o *ImagePackageManifest) GetArtifactRelationshipsOk() (*[]ImportPackageRelationship, bool)`

GetArtifactRelationshipsOk returns a tuple with the ArtifactRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactRelationships

`func (o *ImagePackageManifest) SetArtifactRelationships(v []ImportPackageRelationship)`

SetArtifactRelationships sets ArtifactRelationships field to given value.

### HasArtifactRelationships

`func (o *ImagePackageManifest) HasArtifactRelationships() bool`

HasArtifactRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


