# ImportPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Version** | **string** |  | 
**Type** | **string** |  | 
**FoundBy** | Pointer to **string** |  | [optional] 
**Locations** | [**[]ImportPackageLocation**](ImportPackageLocation.md) |  | 
**Licenses** | **[]string** |  | 
**Language** | **string** |  | 
**Cpes** | **[]string** |  | 
**Purl** | Pointer to **string** |  | [optional] 
**MetadataType** | **string** |  | 
**Metadata** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewImportPackage

`func NewImportPackage(name string, version string, type_ string, locations []ImportPackageLocation, licenses []string, language string, cpes []string, metadataType string, ) *ImportPackage`

NewImportPackage instantiates a new ImportPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportPackageWithDefaults

`func NewImportPackageWithDefaults() *ImportPackage`

NewImportPackageWithDefaults instantiates a new ImportPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportPackage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportPackage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportPackage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImportPackage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ImportPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportPackage) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *ImportPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ImportPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ImportPackage) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetType

`func (o *ImportPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportPackage) SetType(v string)`

SetType sets Type field to given value.


### GetFoundBy

`func (o *ImportPackage) GetFoundBy() string`

GetFoundBy returns the FoundBy field if non-nil, zero value otherwise.

### GetFoundByOk

`func (o *ImportPackage) GetFoundByOk() (*string, bool)`

GetFoundByOk returns a tuple with the FoundBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundBy

`func (o *ImportPackage) SetFoundBy(v string)`

SetFoundBy sets FoundBy field to given value.

### HasFoundBy

`func (o *ImportPackage) HasFoundBy() bool`

HasFoundBy returns a boolean if a field has been set.

### GetLocations

`func (o *ImportPackage) GetLocations() []ImportPackageLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ImportPackage) GetLocationsOk() (*[]ImportPackageLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ImportPackage) SetLocations(v []ImportPackageLocation)`

SetLocations sets Locations field to given value.


### GetLicenses

`func (o *ImportPackage) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *ImportPackage) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *ImportPackage) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.


### GetLanguage

`func (o *ImportPackage) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ImportPackage) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ImportPackage) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetCpes

`func (o *ImportPackage) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *ImportPackage) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *ImportPackage) SetCpes(v []string)`

SetCpes sets Cpes field to given value.


### GetPurl

`func (o *ImportPackage) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *ImportPackage) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *ImportPackage) SetPurl(v string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *ImportPackage) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

### GetMetadataType

`func (o *ImportPackage) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *ImportPackage) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *ImportPackage) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.


### GetMetadata

`func (o *ImportPackage) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImportPackage) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImportPackage) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ImportPackage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


