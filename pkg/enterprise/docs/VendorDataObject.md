# VendorDataObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Vendor Vulnerability ID | [optional] 
**Type** | Pointer to **string** | Indicates the classification of the CVSS | [optional] 
**Source** | Pointer to **string** | Identifies the organization or entity that generated or provided the CVSS score | [optional] 
**CvssV2** | Pointer to [**CVSSV2Scores**](CVSSV2Scores.md) |  | [optional] 
**CvssV3** | Pointer to [**CVSSV3Scores**](CVSSV3Scores.md) |  | [optional] 

## Methods

### NewVendorDataObject

`func NewVendorDataObject() *VendorDataObject`

NewVendorDataObject instantiates a new VendorDataObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorDataObjectWithDefaults

`func NewVendorDataObjectWithDefaults() *VendorDataObject`

NewVendorDataObjectWithDefaults instantiates a new VendorDataObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VendorDataObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VendorDataObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VendorDataObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VendorDataObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *VendorDataObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VendorDataObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VendorDataObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VendorDataObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSource

`func (o *VendorDataObject) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VendorDataObject) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VendorDataObject) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VendorDataObject) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCvssV2

`func (o *VendorDataObject) GetCvssV2() CVSSV2Scores`

GetCvssV2 returns the CvssV2 field if non-nil, zero value otherwise.

### GetCvssV2Ok

`func (o *VendorDataObject) GetCvssV2Ok() (*CVSSV2Scores, bool)`

GetCvssV2Ok returns a tuple with the CvssV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV2

`func (o *VendorDataObject) SetCvssV2(v CVSSV2Scores)`

SetCvssV2 sets CvssV2 field to given value.

### HasCvssV2

`func (o *VendorDataObject) HasCvssV2() bool`

HasCvssV2 returns a boolean if a field has been set.

### GetCvssV3

`func (o *VendorDataObject) GetCvssV3() CVSSV3Scores`

GetCvssV3 returns the CvssV3 field if non-nil, zero value otherwise.

### GetCvssV3Ok

`func (o *VendorDataObject) GetCvssV3Ok() (*CVSSV3Scores, bool)`

GetCvssV3Ok returns a tuple with the CvssV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV3

`func (o *VendorDataObject) SetCvssV3(v CVSSV3Scores)`

SetCvssV3 sets CvssV3 field to given value.

### HasCvssV3

`func (o *VendorDataObject) HasCvssV3() bool`

HasCvssV3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


