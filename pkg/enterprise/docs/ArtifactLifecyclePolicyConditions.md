# ArtifactLifecyclePolicyConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** | The version of these policy conditions. | [optional] [readonly] 
**DaysSinceLastSeenInRuntime** | **int32** | If the image was reported via runtime inventory, it must be this many days old before it will be considered for processing. An integer value less than or equal to zero will prevent the deletion of the image if runtime inventory still has record of it. | 
**DaysSinceAnalyzed** | **int32** | An image analysis must be this many days old before it will be considered for processing. An integer value less than or equal to zero will cause this field to be ignored. | 
**ArtifactType** | **string** | The type of artifact that will be processed. | 

## Methods

### NewArtifactLifecyclePolicyConditions

`func NewArtifactLifecyclePolicyConditions(daysSinceLastSeenInRuntime int32, daysSinceAnalyzed int32, artifactType string, ) *ArtifactLifecyclePolicyConditions`

NewArtifactLifecyclePolicyConditions instantiates a new ArtifactLifecyclePolicyConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactLifecyclePolicyConditionsWithDefaults

`func NewArtifactLifecyclePolicyConditionsWithDefaults() *ArtifactLifecyclePolicyConditions`

NewArtifactLifecyclePolicyConditionsWithDefaults instantiates a new ArtifactLifecyclePolicyConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ArtifactLifecyclePolicyConditions) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArtifactLifecyclePolicyConditions) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArtifactLifecyclePolicyConditions) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ArtifactLifecyclePolicyConditions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDaysSinceLastSeenInRuntime

`func (o *ArtifactLifecyclePolicyConditions) GetDaysSinceLastSeenInRuntime() int32`

GetDaysSinceLastSeenInRuntime returns the DaysSinceLastSeenInRuntime field if non-nil, zero value otherwise.

### GetDaysSinceLastSeenInRuntimeOk

`func (o *ArtifactLifecyclePolicyConditions) GetDaysSinceLastSeenInRuntimeOk() (*int32, bool)`

GetDaysSinceLastSeenInRuntimeOk returns a tuple with the DaysSinceLastSeenInRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSinceLastSeenInRuntime

`func (o *ArtifactLifecyclePolicyConditions) SetDaysSinceLastSeenInRuntime(v int32)`

SetDaysSinceLastSeenInRuntime sets DaysSinceLastSeenInRuntime field to given value.


### GetDaysSinceAnalyzed

`func (o *ArtifactLifecyclePolicyConditions) GetDaysSinceAnalyzed() int32`

GetDaysSinceAnalyzed returns the DaysSinceAnalyzed field if non-nil, zero value otherwise.

### GetDaysSinceAnalyzedOk

`func (o *ArtifactLifecyclePolicyConditions) GetDaysSinceAnalyzedOk() (*int32, bool)`

GetDaysSinceAnalyzedOk returns a tuple with the DaysSinceAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSinceAnalyzed

`func (o *ArtifactLifecyclePolicyConditions) SetDaysSinceAnalyzed(v int32)`

SetDaysSinceAnalyzed sets DaysSinceAnalyzed field to given value.


### GetArtifactType

`func (o *ArtifactLifecyclePolicyConditions) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *ArtifactLifecyclePolicyConditions) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *ArtifactLifecyclePolicyConditions) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


