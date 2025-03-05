# NativeSBOMPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Version** | **string** |  | 
**Type** | **string** |  | 
**FoundBy** | Pointer to **string** |  | [optional] 
**Locations** | [**[]NativeSBOMPackageLocation**](NativeSBOMPackageLocation.md) |  | 
**Licenses** | [**[]ImportPackageLicensesInner**](ImportPackageLicensesInner.md) |  | 
**Language** | **string** |  | 
**Cpes** | [**[]NativeSBOMPackageCpesInner**](NativeSBOMPackageCpesInner.md) |  | 
**Purl** | Pointer to **string** |  | [optional] 
**MetadataType** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewNativeSBOMPackage

`func NewNativeSBOMPackage(name string, version string, type_ string, locations []NativeSBOMPackageLocation, licenses []ImportPackageLicensesInner, language string, cpes []NativeSBOMPackageCpesInner, ) *NativeSBOMPackage`

NewNativeSBOMPackage instantiates a new NativeSBOMPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNativeSBOMPackageWithDefaults

`func NewNativeSBOMPackageWithDefaults() *NativeSBOMPackage`

NewNativeSBOMPackageWithDefaults instantiates a new NativeSBOMPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NativeSBOMPackage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NativeSBOMPackage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NativeSBOMPackage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NativeSBOMPackage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NativeSBOMPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NativeSBOMPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NativeSBOMPackage) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *NativeSBOMPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NativeSBOMPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NativeSBOMPackage) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetType

`func (o *NativeSBOMPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NativeSBOMPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NativeSBOMPackage) SetType(v string)`

SetType sets Type field to given value.


### GetFoundBy

`func (o *NativeSBOMPackage) GetFoundBy() string`

GetFoundBy returns the FoundBy field if non-nil, zero value otherwise.

### GetFoundByOk

`func (o *NativeSBOMPackage) GetFoundByOk() (*string, bool)`

GetFoundByOk returns a tuple with the FoundBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundBy

`func (o *NativeSBOMPackage) SetFoundBy(v string)`

SetFoundBy sets FoundBy field to given value.

### HasFoundBy

`func (o *NativeSBOMPackage) HasFoundBy() bool`

HasFoundBy returns a boolean if a field has been set.

### GetLocations

`func (o *NativeSBOMPackage) GetLocations() []NativeSBOMPackageLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *NativeSBOMPackage) GetLocationsOk() (*[]NativeSBOMPackageLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *NativeSBOMPackage) SetLocations(v []NativeSBOMPackageLocation)`

SetLocations sets Locations field to given value.


### GetLicenses

`func (o *NativeSBOMPackage) GetLicenses() []ImportPackageLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *NativeSBOMPackage) GetLicensesOk() (*[]ImportPackageLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *NativeSBOMPackage) SetLicenses(v []ImportPackageLicensesInner)`

SetLicenses sets Licenses field to given value.


### GetLanguage

`func (o *NativeSBOMPackage) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *NativeSBOMPackage) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *NativeSBOMPackage) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetCpes

`func (o *NativeSBOMPackage) GetCpes() []NativeSBOMPackageCpesInner`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *NativeSBOMPackage) GetCpesOk() (*[]NativeSBOMPackageCpesInner, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *NativeSBOMPackage) SetCpes(v []NativeSBOMPackageCpesInner)`

SetCpes sets Cpes field to given value.


### GetPurl

`func (o *NativeSBOMPackage) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *NativeSBOMPackage) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *NativeSBOMPackage) SetPurl(v string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *NativeSBOMPackage) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

### GetMetadataType

`func (o *NativeSBOMPackage) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *NativeSBOMPackage) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *NativeSBOMPackage) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *NativeSBOMPackage) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.

### SetMetadataTypeNil

`func (o *NativeSBOMPackage) SetMetadataTypeNil(b bool)`

 SetMetadataTypeNil sets the value for MetadataType to be an explicit nil

### UnsetMetadataType
`func (o *NativeSBOMPackage) UnsetMetadataType()`

UnsetMetadataType ensures that no value is present for MetadataType, not even an explicit nil
### GetMetadata

`func (o *NativeSBOMPackage) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NativeSBOMPackage) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NativeSBOMPackage) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NativeSBOMPackage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *NativeSBOMPackage) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *NativeSBOMPackage) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


