# PackageReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Package name | [optional] 
**Version** | Pointer to **NullableString** | A version for the package. If null, then references all versions | [optional] 
**Type** | Pointer to **string** | Package type (e.g. package, rpm, deb, apk, jar, npm, gem, ...) | [optional] 
**WillNotFix** | Pointer to **bool** | Whether a vendor will or will not fix a vulnerabitlity | [optional] 

## Methods

### NewPackageReference

`func NewPackageReference() *PackageReference`

NewPackageReference instantiates a new PackageReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageReferenceWithDefaults

`func NewPackageReferenceWithDefaults() *PackageReference`

NewPackageReferenceWithDefaults instantiates a new PackageReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PackageReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackageReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *PackageReference) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackageReference) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackageReference) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackageReference) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PackageReference) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PackageReference) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetType

`func (o *PackageReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackageReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackageReference) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PackageReference) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWillNotFix

`func (o *PackageReference) GetWillNotFix() bool`

GetWillNotFix returns the WillNotFix field if non-nil, zero value otherwise.

### GetWillNotFixOk

`func (o *PackageReference) GetWillNotFixOk() (*bool, bool)`

GetWillNotFixOk returns a tuple with the WillNotFix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillNotFix

`func (o *PackageReference) SetWillNotFix(v bool)`

SetWillNotFix sets WillNotFix field to given value.

### HasWillNotFix

`func (o *PackageReference) HasWillNotFix() bool`

HasWillNotFix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


