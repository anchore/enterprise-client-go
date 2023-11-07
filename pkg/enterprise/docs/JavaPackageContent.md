# JavaPackageContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to **string** |  | [optional] 
**ImplementationVersion** | Pointer to **string** |  | [optional] 
**SpecificationVersion** | Pointer to **string** |  | [optional] 
**MavenVersion** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**Cpes** | Pointer to **[]string** | A list of Common Platform Enumerations that may uniquely identify the package | [optional] 
**Purl** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewJavaPackageContent

`func NewJavaPackageContent() *JavaPackageContent`

NewJavaPackageContent instantiates a new JavaPackageContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJavaPackageContentWithDefaults

`func NewJavaPackageContentWithDefaults() *JavaPackageContent`

NewJavaPackageContentWithDefaults instantiates a new JavaPackageContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *JavaPackageContent) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *JavaPackageContent) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *JavaPackageContent) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *JavaPackageContent) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetImplementationVersion

`func (o *JavaPackageContent) GetImplementationVersion() string`

GetImplementationVersion returns the ImplementationVersion field if non-nil, zero value otherwise.

### GetImplementationVersionOk

`func (o *JavaPackageContent) GetImplementationVersionOk() (*string, bool)`

GetImplementationVersionOk returns a tuple with the ImplementationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationVersion

`func (o *JavaPackageContent) SetImplementationVersion(v string)`

SetImplementationVersion sets ImplementationVersion field to given value.

### HasImplementationVersion

`func (o *JavaPackageContent) HasImplementationVersion() bool`

HasImplementationVersion returns a boolean if a field has been set.

### GetSpecificationVersion

`func (o *JavaPackageContent) GetSpecificationVersion() string`

GetSpecificationVersion returns the SpecificationVersion field if non-nil, zero value otherwise.

### GetSpecificationVersionOk

`func (o *JavaPackageContent) GetSpecificationVersionOk() (*string, bool)`

GetSpecificationVersionOk returns a tuple with the SpecificationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificationVersion

`func (o *JavaPackageContent) SetSpecificationVersion(v string)`

SetSpecificationVersion sets SpecificationVersion field to given value.

### HasSpecificationVersion

`func (o *JavaPackageContent) HasSpecificationVersion() bool`

HasSpecificationVersion returns a boolean if a field has been set.

### GetMavenVersion

`func (o *JavaPackageContent) GetMavenVersion() string`

GetMavenVersion returns the MavenVersion field if non-nil, zero value otherwise.

### GetMavenVersionOk

`func (o *JavaPackageContent) GetMavenVersionOk() (*string, bool)`

GetMavenVersionOk returns a tuple with the MavenVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMavenVersion

`func (o *JavaPackageContent) SetMavenVersion(v string)`

SetMavenVersion sets MavenVersion field to given value.

### HasMavenVersion

`func (o *JavaPackageContent) HasMavenVersion() bool`

HasMavenVersion returns a boolean if a field has been set.

### GetLocation

`func (o *JavaPackageContent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *JavaPackageContent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *JavaPackageContent) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *JavaPackageContent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetType

`func (o *JavaPackageContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JavaPackageContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JavaPackageContent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JavaPackageContent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrigin

`func (o *JavaPackageContent) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *JavaPackageContent) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *JavaPackageContent) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *JavaPackageContent) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetCpes

`func (o *JavaPackageContent) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *JavaPackageContent) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *JavaPackageContent) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *JavaPackageContent) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetPurl

`func (o *JavaPackageContent) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *JavaPackageContent) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *JavaPackageContent) SetPurl(v string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *JavaPackageContent) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

### GetVersion

`func (o *JavaPackageContent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *JavaPackageContent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *JavaPackageContent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *JavaPackageContent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


