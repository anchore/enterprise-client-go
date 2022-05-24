# NvdDataObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | NVD Vulnerability ID | [optional] 
**CvssV2** | Pointer to [**CVSSV2Scores**](CVSSV2Scores.md) |  | [optional] 
**CvssV3** | Pointer to [**CVSSV3Scores**](CVSSV3Scores.md) |  | [optional] 

## Methods

### NewNvdDataObject

`func NewNvdDataObject() *NvdDataObject`

NewNvdDataObject instantiates a new NvdDataObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNvdDataObjectWithDefaults

`func NewNvdDataObjectWithDefaults() *NvdDataObject`

NewNvdDataObjectWithDefaults instantiates a new NvdDataObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NvdDataObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NvdDataObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NvdDataObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NvdDataObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCvssV2

`func (o *NvdDataObject) GetCvssV2() CVSSV2Scores`

GetCvssV2 returns the CvssV2 field if non-nil, zero value otherwise.

### GetCvssV2Ok

`func (o *NvdDataObject) GetCvssV2Ok() (*CVSSV2Scores, bool)`

GetCvssV2Ok returns a tuple with the CvssV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV2

`func (o *NvdDataObject) SetCvssV2(v CVSSV2Scores)`

SetCvssV2 sets CvssV2 field to given value.

### HasCvssV2

`func (o *NvdDataObject) HasCvssV2() bool`

HasCvssV2 returns a boolean if a field has been set.

### GetCvssV3

`func (o *NvdDataObject) GetCvssV3() CVSSV3Scores`

GetCvssV3 returns the CvssV3 field if non-nil, zero value otherwise.

### GetCvssV3Ok

`func (o *NvdDataObject) GetCvssV3Ok() (*CVSSV3Scores, bool)`

GetCvssV3Ok returns a tuple with the CvssV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV3

`func (o *NvdDataObject) SetCvssV3(v CVSSV3Scores)`

SetCvssV3 sets CvssV3 field to given value.

### HasCvssV3

`func (o *NvdDataObject) HasCvssV3() bool`

HasCvssV3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


