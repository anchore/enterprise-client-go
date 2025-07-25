# PackageEPSS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**Epss** | Pointer to **float32** |  | [optional] 
**Percentile** | Pointer to **float32** |  | [optional] 

## Methods

### NewPackageEPSS

`func NewPackageEPSS() *PackageEPSS`

NewPackageEPSS instantiates a new PackageEPSS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageEPSSWithDefaults

`func NewPackageEPSSWithDefaults() *PackageEPSS`

NewPackageEPSSWithDefaults instantiates a new PackageEPSS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *PackageEPSS) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *PackageEPSS) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *PackageEPSS) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *PackageEPSS) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDate

`func (o *PackageEPSS) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PackageEPSS) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PackageEPSS) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *PackageEPSS) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEpss

`func (o *PackageEPSS) GetEpss() float32`

GetEpss returns the Epss field if non-nil, zero value otherwise.

### GetEpssOk

`func (o *PackageEPSS) GetEpssOk() (*float32, bool)`

GetEpssOk returns a tuple with the Epss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpss

`func (o *PackageEPSS) SetEpss(v float32)`

SetEpss sets Epss field to given value.

### HasEpss

`func (o *PackageEPSS) HasEpss() bool`

HasEpss returns a boolean if a field has been set.

### GetPercentile

`func (o *PackageEPSS) GetPercentile() float32`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *PackageEPSS) GetPercentileOk() (*float32, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *PackageEPSS) SetPercentile(v float32)`

SetPercentile sets Percentile field to given value.

### HasPercentile

`func (o *PackageEPSS) HasPercentile() bool`

HasPercentile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


