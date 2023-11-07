# ArtifactLifeCyclePolicyConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** |  | [optional] 
**DaysSinceLastSeenInRuntime** | **int32** |  | 
**DaysSinceAnalyzed** | **int32** |  | 
**ArtifactType** | **string** |  | 

## Methods

### NewArtifactLifeCyclePolicyConditions

`func NewArtifactLifeCyclePolicyConditions(daysSinceLastSeenInRuntime int32, daysSinceAnalyzed int32, artifactType string, ) *ArtifactLifeCyclePolicyConditions`

NewArtifactLifeCyclePolicyConditions instantiates a new ArtifactLifeCyclePolicyConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactLifeCyclePolicyConditionsWithDefaults

`func NewArtifactLifeCyclePolicyConditionsWithDefaults() *ArtifactLifeCyclePolicyConditions`

NewArtifactLifeCyclePolicyConditionsWithDefaults instantiates a new ArtifactLifeCyclePolicyConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ArtifactLifeCyclePolicyConditions) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArtifactLifeCyclePolicyConditions) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArtifactLifeCyclePolicyConditions) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ArtifactLifeCyclePolicyConditions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDaysSinceLastSeenInRuntime

`func (o *ArtifactLifeCyclePolicyConditions) GetDaysSinceLastSeenInRuntime() int32`

GetDaysSinceLastSeenInRuntime returns the DaysSinceLastSeenInRuntime field if non-nil, zero value otherwise.

### GetDaysSinceLastSeenInRuntimeOk

`func (o *ArtifactLifeCyclePolicyConditions) GetDaysSinceLastSeenInRuntimeOk() (*int32, bool)`

GetDaysSinceLastSeenInRuntimeOk returns a tuple with the DaysSinceLastSeenInRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSinceLastSeenInRuntime

`func (o *ArtifactLifeCyclePolicyConditions) SetDaysSinceLastSeenInRuntime(v int32)`

SetDaysSinceLastSeenInRuntime sets DaysSinceLastSeenInRuntime field to given value.


### GetDaysSinceAnalyzed

`func (o *ArtifactLifeCyclePolicyConditions) GetDaysSinceAnalyzed() int32`

GetDaysSinceAnalyzed returns the DaysSinceAnalyzed field if non-nil, zero value otherwise.

### GetDaysSinceAnalyzedOk

`func (o *ArtifactLifeCyclePolicyConditions) GetDaysSinceAnalyzedOk() (*int32, bool)`

GetDaysSinceAnalyzedOk returns a tuple with the DaysSinceAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSinceAnalyzed

`func (o *ArtifactLifeCyclePolicyConditions) SetDaysSinceAnalyzed(v int32)`

SetDaysSinceAnalyzed sets DaysSinceAnalyzed field to given value.


### GetArtifactType

`func (o *ArtifactLifeCyclePolicyConditions) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *ArtifactLifeCyclePolicyConditions) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *ArtifactLifeCyclePolicyConditions) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


