# NativeSBOM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | [**[]NativeSBOMPackage**](NativeSBOMPackage.md) |  | 
**Source** | [**NativeSBOMSource**](NativeSBOMSource.md) |  | 
**Distro** | [**NativeSBOMDistribution**](NativeSBOMDistribution.md) |  | 
**Descriptor** | Pointer to [**NativeSBOMDescriptor**](NativeSBOMDescriptor.md) |  | [optional] 
**Schema** | Pointer to [**NativeSBOMSchema**](NativeSBOMSchema.md) |  | [optional] 
**ArtifactRelationships** | Pointer to [**[]NativeSBOMPackageRelationship**](NativeSBOMPackageRelationship.md) |  | [optional] 

## Methods

### NewNativeSBOM

`func NewNativeSBOM(artifacts []NativeSBOMPackage, source NativeSBOMSource, distro NativeSBOMDistribution, ) *NativeSBOM`

NewNativeSBOM instantiates a new NativeSBOM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNativeSBOMWithDefaults

`func NewNativeSBOMWithDefaults() *NativeSBOM`

NewNativeSBOMWithDefaults instantiates a new NativeSBOM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *NativeSBOM) GetArtifacts() []NativeSBOMPackage`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *NativeSBOM) GetArtifactsOk() (*[]NativeSBOMPackage, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *NativeSBOM) SetArtifacts(v []NativeSBOMPackage)`

SetArtifacts sets Artifacts field to given value.


### GetSource

`func (o *NativeSBOM) GetSource() NativeSBOMSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NativeSBOM) GetSourceOk() (*NativeSBOMSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NativeSBOM) SetSource(v NativeSBOMSource)`

SetSource sets Source field to given value.


### GetDistro

`func (o *NativeSBOM) GetDistro() NativeSBOMDistribution`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *NativeSBOM) GetDistroOk() (*NativeSBOMDistribution, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *NativeSBOM) SetDistro(v NativeSBOMDistribution)`

SetDistro sets Distro field to given value.


### GetDescriptor

`func (o *NativeSBOM) GetDescriptor() NativeSBOMDescriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *NativeSBOM) GetDescriptorOk() (*NativeSBOMDescriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *NativeSBOM) SetDescriptor(v NativeSBOMDescriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *NativeSBOM) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetSchema

`func (o *NativeSBOM) GetSchema() NativeSBOMSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *NativeSBOM) GetSchemaOk() (*NativeSBOMSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *NativeSBOM) SetSchema(v NativeSBOMSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *NativeSBOM) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetArtifactRelationships

`func (o *NativeSBOM) GetArtifactRelationships() []NativeSBOMPackageRelationship`

GetArtifactRelationships returns the ArtifactRelationships field if non-nil, zero value otherwise.

### GetArtifactRelationshipsOk

`func (o *NativeSBOM) GetArtifactRelationshipsOk() (*[]NativeSBOMPackageRelationship, bool)`

GetArtifactRelationshipsOk returns a tuple with the ArtifactRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactRelationships

`func (o *NativeSBOM) SetArtifactRelationships(v []NativeSBOMPackageRelationship)`

SetArtifactRelationships sets ArtifactRelationships field to given value.

### HasArtifactRelationships

`func (o *NativeSBOM) HasArtifactRelationships() bool`

HasArtifactRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


