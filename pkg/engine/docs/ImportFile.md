# ImportFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Location** | **interface{}** |  | 
**Metadata** | **interface{}** |  | 
**Digests** | Pointer to [**[]ImportFileDigest**](ImportFileDigest.md) |  | [optional] 

## Methods

### NewImportFile

`func NewImportFile(id string, location interface{}, metadata interface{}, ) *ImportFile`

NewImportFile instantiates a new ImportFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportFileWithDefaults

`func NewImportFileWithDefaults() *ImportFile`

NewImportFileWithDefaults instantiates a new ImportFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportFile) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *ImportFile) GetLocation() interface{}`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ImportFile) GetLocationOk() (*interface{}, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ImportFile) SetLocation(v interface{})`

SetLocation sets Location field to given value.


### GetMetadata

`func (o *ImportFile) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImportFile) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImportFile) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.


### GetDigests

`func (o *ImportFile) GetDigests() []ImportFileDigest`

GetDigests returns the Digests field if non-nil, zero value otherwise.

### GetDigestsOk

`func (o *ImportFile) GetDigestsOk() (*[]ImportFileDigest, bool)`

GetDigestsOk returns a tuple with the Digests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigests

`func (o *ImportFile) SetDigests(v []ImportFileDigest)`

SetDigests sets Digests field to given value.

### HasDigests

`func (o *ImportFile) HasDigests() bool`

HasDigests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


