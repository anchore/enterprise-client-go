# VulnerableImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**ImageReference**](ImageReference.md) |  | [optional] 
**AffectedPackages** | Pointer to [**[]VulnerablePackageReference**](VulnerablePackageReference.md) |  | [optional] 

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

### GetAffectedPackages

`func (o *VulnerableImage) GetAffectedPackages() []VulnerablePackageReference`

GetAffectedPackages returns the AffectedPackages field if non-nil, zero value otherwise.

### GetAffectedPackagesOk

`func (o *VulnerableImage) GetAffectedPackagesOk() (*[]VulnerablePackageReference, bool)`

GetAffectedPackagesOk returns a tuple with the AffectedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPackages

`func (o *VulnerableImage) SetAffectedPackages(v []VulnerablePackageReference)`

SetAffectedPackages sets AffectedPackages field to given value.

### HasAffectedPackages

`func (o *VulnerableImage) HasAffectedPackages() bool`

HasAffectedPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


