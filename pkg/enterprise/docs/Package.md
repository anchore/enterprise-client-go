# Package

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**Sourcepkg** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Licenses** | Pointer to **[]string** |  | [optional] 
**MetadataType** | Pointer to **string** | The type of the metadata entry | [optional] 
**Metadata** | Pointer to **interface{}** | Package type specific metadata | [optional] 
**SpecificationVersion** | Pointer to **string** | Spec version for java packages | [optional] 
**ImplementationVersion** | Pointer to **string** | Implementation version for java packages | [optional] 
**MavenVersion** | Pointer to **string** | Maven version for java packages | [optional] 
**Cpes** | Pointer to **[]string** | List of CPE strings for this package | [optional] 

## Methods

### NewPackage

`func NewPackage() *Package`

NewPackage instantiates a new Package object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageWithDefaults

`func NewPackageWithDefaults() *Package`

NewPackageWithDefaults instantiates a new Package object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Package) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Package) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Package) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Package) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Package) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Package) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Package) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Package) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRelease

`func (o *Package) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *Package) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *Package) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *Package) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetSourcepkg

`func (o *Package) GetSourcepkg() string`

GetSourcepkg returns the Sourcepkg field if non-nil, zero value otherwise.

### GetSourcepkgOk

`func (o *Package) GetSourcepkgOk() (*string, bool)`

GetSourcepkgOk returns a tuple with the Sourcepkg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcepkg

`func (o *Package) SetSourcepkg(v string)`

SetSourcepkg sets Sourcepkg field to given value.

### HasSourcepkg

`func (o *Package) HasSourcepkg() bool`

HasSourcepkg returns a boolean if a field has been set.

### GetLocation

`func (o *Package) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Package) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Package) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Package) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetOrigin

`func (o *Package) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Package) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Package) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Package) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetSize

`func (o *Package) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Package) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Package) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Package) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLicenses

`func (o *Package) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *Package) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *Package) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *Package) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMetadataType

`func (o *Package) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *Package) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *Package) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *Package) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.

### GetMetadata

`func (o *Package) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Package) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Package) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Package) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpecificationVersion

`func (o *Package) GetSpecificationVersion() string`

GetSpecificationVersion returns the SpecificationVersion field if non-nil, zero value otherwise.

### GetSpecificationVersionOk

`func (o *Package) GetSpecificationVersionOk() (*string, bool)`

GetSpecificationVersionOk returns a tuple with the SpecificationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificationVersion

`func (o *Package) SetSpecificationVersion(v string)`

SetSpecificationVersion sets SpecificationVersion field to given value.

### HasSpecificationVersion

`func (o *Package) HasSpecificationVersion() bool`

HasSpecificationVersion returns a boolean if a field has been set.

### GetImplementationVersion

`func (o *Package) GetImplementationVersion() string`

GetImplementationVersion returns the ImplementationVersion field if non-nil, zero value otherwise.

### GetImplementationVersionOk

`func (o *Package) GetImplementationVersionOk() (*string, bool)`

GetImplementationVersionOk returns a tuple with the ImplementationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationVersion

`func (o *Package) SetImplementationVersion(v string)`

SetImplementationVersion sets ImplementationVersion field to given value.

### HasImplementationVersion

`func (o *Package) HasImplementationVersion() bool`

HasImplementationVersion returns a boolean if a field has been set.

### GetMavenVersion

`func (o *Package) GetMavenVersion() string`

GetMavenVersion returns the MavenVersion field if non-nil, zero value otherwise.

### GetMavenVersionOk

`func (o *Package) GetMavenVersionOk() (*string, bool)`

GetMavenVersionOk returns a tuple with the MavenVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMavenVersion

`func (o *Package) SetMavenVersion(v string)`

SetMavenVersion sets MavenVersion field to given value.

### HasMavenVersion

`func (o *Package) HasMavenVersion() bool`

HasMavenVersion returns a boolean if a field has been set.

### GetCpes

`func (o *Package) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *Package) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *Package) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *Package) HasCpes() bool`

HasCpes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


