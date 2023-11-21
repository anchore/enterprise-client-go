# ContentJAVAPackageResponseContentInner

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
<<<<<<< HEAD:pkg/enterprise/docs/ContentJAVAPackageResponseContentInner.md
**Version** | Pointer to **string** |  | [optional] 
=======
>>>>>>> main:pkg/enterprise/docs/ContentJAVAPackageResponseContent.md

## Methods

### NewContentJAVAPackageResponseContentInner

`func NewContentJAVAPackageResponseContentInner() *ContentJAVAPackageResponseContentInner`

NewContentJAVAPackageResponseContentInner instantiates a new ContentJAVAPackageResponseContentInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentJAVAPackageResponseContentInnerWithDefaults

`func NewContentJAVAPackageResponseContentInnerWithDefaults() *ContentJAVAPackageResponseContentInner`

NewContentJAVAPackageResponseContentInnerWithDefaults instantiates a new ContentJAVAPackageResponseContentInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *ContentJAVAPackageResponseContentInner) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ContentJAVAPackageResponseContentInner) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ContentJAVAPackageResponseContentInner) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *ContentJAVAPackageResponseContentInner) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetImplementationVersion

`func (o *ContentJAVAPackageResponseContentInner) GetImplementationVersion() string`

GetImplementationVersion returns the ImplementationVersion field if non-nil, zero value otherwise.

### GetImplementationVersionOk

`func (o *ContentJAVAPackageResponseContentInner) GetImplementationVersionOk() (*string, bool)`

GetImplementationVersionOk returns a tuple with the ImplementationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationVersion

`func (o *ContentJAVAPackageResponseContentInner) SetImplementationVersion(v string)`

SetImplementationVersion sets ImplementationVersion field to given value.

### HasImplementationVersion

`func (o *ContentJAVAPackageResponseContentInner) HasImplementationVersion() bool`

HasImplementationVersion returns a boolean if a field has been set.

### GetSpecificationVersion

`func (o *ContentJAVAPackageResponseContentInner) GetSpecificationVersion() string`

GetSpecificationVersion returns the SpecificationVersion field if non-nil, zero value otherwise.

### GetSpecificationVersionOk

`func (o *ContentJAVAPackageResponseContentInner) GetSpecificationVersionOk() (*string, bool)`

GetSpecificationVersionOk returns a tuple with the SpecificationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificationVersion

`func (o *ContentJAVAPackageResponseContentInner) SetSpecificationVersion(v string)`

SetSpecificationVersion sets SpecificationVersion field to given value.

### HasSpecificationVersion

`func (o *ContentJAVAPackageResponseContentInner) HasSpecificationVersion() bool`

HasSpecificationVersion returns a boolean if a field has been set.

### GetMavenVersion

`func (o *ContentJAVAPackageResponseContentInner) GetMavenVersion() string`

GetMavenVersion returns the MavenVersion field if non-nil, zero value otherwise.

### GetMavenVersionOk

`func (o *ContentJAVAPackageResponseContentInner) GetMavenVersionOk() (*string, bool)`

GetMavenVersionOk returns a tuple with the MavenVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMavenVersion

`func (o *ContentJAVAPackageResponseContentInner) SetMavenVersion(v string)`

SetMavenVersion sets MavenVersion field to given value.

### HasMavenVersion

`func (o *ContentJAVAPackageResponseContentInner) HasMavenVersion() bool`

HasMavenVersion returns a boolean if a field has been set.

### GetLocation

`func (o *ContentJAVAPackageResponseContentInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ContentJAVAPackageResponseContentInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ContentJAVAPackageResponseContentInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ContentJAVAPackageResponseContentInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetType

`func (o *ContentJAVAPackageResponseContentInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentJAVAPackageResponseContentInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentJAVAPackageResponseContentInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentJAVAPackageResponseContentInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrigin

`func (o *ContentJAVAPackageResponseContentInner) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ContentJAVAPackageResponseContentInner) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ContentJAVAPackageResponseContentInner) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ContentJAVAPackageResponseContentInner) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetCpes

`func (o *ContentJAVAPackageResponseContentInner) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *ContentJAVAPackageResponseContentInner) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *ContentJAVAPackageResponseContentInner) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *ContentJAVAPackageResponseContentInner) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetPurl

<<<<<<< HEAD:pkg/enterprise/docs/ContentJAVAPackageResponseContentInner.md
`func (o *ContentJAVAPackageResponseContentInner) GetPurl() string`
=======
`func (o *ContentJAVAPackageResponseContent) GetPurl() string`
>>>>>>> main:pkg/enterprise/docs/ContentJAVAPackageResponseContent.md

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

<<<<<<< HEAD:pkg/enterprise/docs/ContentJAVAPackageResponseContentInner.md
`func (o *ContentJAVAPackageResponseContentInner) GetPurlOk() (*string, bool)`
=======
`func (o *ContentJAVAPackageResponseContent) GetPurlOk() (*string, bool)`
>>>>>>> main:pkg/enterprise/docs/ContentJAVAPackageResponseContent.md

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

<<<<<<< HEAD:pkg/enterprise/docs/ContentJAVAPackageResponseContentInner.md
`func (o *ContentJAVAPackageResponseContentInner) SetPurl(v string)`
=======
`func (o *ContentJAVAPackageResponseContent) SetPurl(v string)`
>>>>>>> main:pkg/enterprise/docs/ContentJAVAPackageResponseContent.md

SetPurl sets Purl field to given value.

### HasPurl

<<<<<<< HEAD:pkg/enterprise/docs/ContentJAVAPackageResponseContentInner.md
`func (o *ContentJAVAPackageResponseContentInner) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

### GetVersion

`func (o *ContentJAVAPackageResponseContentInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContentJAVAPackageResponseContentInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContentJAVAPackageResponseContentInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ContentJAVAPackageResponseContentInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

=======
`func (o *ContentJAVAPackageResponseContent) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

>>>>>>> main:pkg/enterprise/docs/ContentJAVAPackageResponseContent.md

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


