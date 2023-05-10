# VulnerablePackageReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Package name | [optional] 
**Version** | Pointer to **NullableString** | A version for the package. If null, then references all versions | [optional] 
**Type** | Pointer to **string** | Package type (e.g. package, rpm, deb, apk, jar, npm, gem, ...) | [optional] 
**Severity** | Pointer to **string** | Severity of vulnerability affecting package | [optional] 
**Namespace** | Pointer to **string** | Vulnerability namespace of affected package | [optional] 

## Methods

### NewVulnerablePackageReference

`func NewVulnerablePackageReference() *VulnerablePackageReference`

NewVulnerablePackageReference instantiates a new VulnerablePackageReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnerablePackageReferenceWithDefaults

`func NewVulnerablePackageReferenceWithDefaults() *VulnerablePackageReference`

NewVulnerablePackageReferenceWithDefaults instantiates a new VulnerablePackageReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VulnerablePackageReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VulnerablePackageReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VulnerablePackageReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VulnerablePackageReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *VulnerablePackageReference) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VulnerablePackageReference) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VulnerablePackageReference) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VulnerablePackageReference) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *VulnerablePackageReference) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *VulnerablePackageReference) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetType

`func (o *VulnerablePackageReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VulnerablePackageReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VulnerablePackageReference) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VulnerablePackageReference) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSeverity

`func (o *VulnerablePackageReference) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *VulnerablePackageReference) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *VulnerablePackageReference) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *VulnerablePackageReference) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetNamespace

`func (o *VulnerablePackageReference) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *VulnerablePackageReference) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *VulnerablePackageReference) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *VulnerablePackageReference) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


