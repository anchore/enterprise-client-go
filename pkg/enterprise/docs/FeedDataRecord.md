# FeedDataRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the data record | [optional] 
**Name** | Pointer to **string** | The name of the data record | [optional] 
**Dataset** | Pointer to **string** | The data set type of the record | [optional] 
**DataFormat** | Pointer to **string** | The format of the data record | [optional] 
**Checksum** | Pointer to **string** | The checksum of the data record | [optional] 
**BuiltAt** | Pointer to **time.Time** | The build timestamp of the data record | [optional] 
**CreatedAt** | Pointer to **time.Time** | The creation timestamp of the data record | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The last update timestamp of the data record | [optional] 

## Methods

### NewFeedDataRecord

`func NewFeedDataRecord() *FeedDataRecord`

NewFeedDataRecord instantiates a new FeedDataRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedDataRecordWithDefaults

`func NewFeedDataRecordWithDefaults() *FeedDataRecord`

NewFeedDataRecordWithDefaults instantiates a new FeedDataRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeedDataRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeedDataRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeedDataRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeedDataRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FeedDataRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeedDataRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeedDataRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeedDataRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataset

`func (o *FeedDataRecord) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *FeedDataRecord) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *FeedDataRecord) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *FeedDataRecord) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetDataFormat

`func (o *FeedDataRecord) GetDataFormat() string`

GetDataFormat returns the DataFormat field if non-nil, zero value otherwise.

### GetDataFormatOk

`func (o *FeedDataRecord) GetDataFormatOk() (*string, bool)`

GetDataFormatOk returns a tuple with the DataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFormat

`func (o *FeedDataRecord) SetDataFormat(v string)`

SetDataFormat sets DataFormat field to given value.

### HasDataFormat

`func (o *FeedDataRecord) HasDataFormat() bool`

HasDataFormat returns a boolean if a field has been set.

### GetChecksum

`func (o *FeedDataRecord) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *FeedDataRecord) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *FeedDataRecord) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *FeedDataRecord) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetBuiltAt

`func (o *FeedDataRecord) GetBuiltAt() time.Time`

GetBuiltAt returns the BuiltAt field if non-nil, zero value otherwise.

### GetBuiltAtOk

`func (o *FeedDataRecord) GetBuiltAtOk() (*time.Time, bool)`

GetBuiltAtOk returns a tuple with the BuiltAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltAt

`func (o *FeedDataRecord) SetBuiltAt(v time.Time)`

SetBuiltAt sets BuiltAt field to given value.

### HasBuiltAt

`func (o *FeedDataRecord) HasBuiltAt() bool`

HasBuiltAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FeedDataRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeedDataRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeedDataRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FeedDataRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FeedDataRecord) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FeedDataRecord) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FeedDataRecord) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FeedDataRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


