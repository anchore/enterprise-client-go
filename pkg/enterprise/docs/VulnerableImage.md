# VulnerableImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**ImageReference**](ImageReference.md) |  | [optional] 
**VulnerablePackages** | Pointer to [**[]VulnerablePackageReference**](VulnerablePackageReference.md) |  | [optional] 

## Methods

### NewVulnerableImage

`func NewVulnerableImage() *VulnerableImage`

NewVulnerableImage instantiates a new VulnerableImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnerableImageWithDefaults

`func NewVulnerableImageWithDefaults() *VulnerableImage`

NewVulnerableImageWithDefaults instantiates a new VulnerableImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *VulnerableImage) GetImage() ImageReference`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VulnerableImage) GetImageOk() (*ImageReference, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VulnerableImage) SetImage(v ImageReference)`

SetImage sets Image field to given value.

### HasImage

`func (o *VulnerableImage) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVulnerablePackages

`func (o *VulnerableImage) GetVulnerablePackages() []VulnerablePackageReference`

GetVulnerablePackages returns the VulnerablePackages field if non-nil, zero value otherwise.

### GetVulnerablePackagesOk

`func (o *VulnerableImage) GetVulnerablePackagesOk() (*[]VulnerablePackageReference, bool)`

GetVulnerablePackagesOk returns a tuple with the VulnerablePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerablePackages

`func (o *VulnerableImage) SetVulnerablePackages(v []VulnerablePackageReference)`

SetVulnerablePackages sets VulnerablePackages field to given value.

### HasVulnerablePackages

`func (o *VulnerableImage) HasVulnerablePackages() bool`

HasVulnerablePackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


