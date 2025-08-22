# PackageIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the package to which this scope applies | 
**Version** | **string** | The version of the package to which this scope applies | 
**Type** | **string** | The type of the package to which this scope applies | 

## Methods

### NewPackageIdentifier

`func NewPackageIdentifier(name string, version string, type_ string, ) *PackageIdentifier`

NewPackageIdentifier instantiates a new PackageIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageIdentifierWithDefaults

`func NewPackageIdentifierWithDefaults() *PackageIdentifier`

NewPackageIdentifierWithDefaults instantiates a new PackageIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PackageIdentifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageIdentifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageIdentifier) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *PackageIdentifier) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackageIdentifier) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackageIdentifier) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetType

`func (o *PackageIdentifier) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackageIdentifier) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackageIdentifier) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


