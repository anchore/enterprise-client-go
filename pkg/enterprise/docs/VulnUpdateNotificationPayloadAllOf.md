# VulnUpdateNotificationPayloadAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiffVulnerabilityResult** | Pointer to [**VulnDiffResult**](VulnDiffResult.md) |  | [optional] 
**ImageDigest** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **interface{}** | List of Corresponding Image Annotations | [optional] 

## Methods

### NewVulnUpdateNotificationPayloadAllOf

`func NewVulnUpdateNotificationPayloadAllOf() *VulnUpdateNotificationPayloadAllOf`

NewVulnUpdateNotificationPayloadAllOf instantiates a new VulnUpdateNotificationPayloadAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnUpdateNotificationPayloadAllOfWithDefaults

`func NewVulnUpdateNotificationPayloadAllOfWithDefaults() *VulnUpdateNotificationPayloadAllOf`

NewVulnUpdateNotificationPayloadAllOfWithDefaults instantiates a new VulnUpdateNotificationPayloadAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiffVulnerabilityResult

`func (o *VulnUpdateNotificationPayloadAllOf) GetDiffVulnerabilityResult() VulnDiffResult`

GetDiffVulnerabilityResult returns the DiffVulnerabilityResult field if non-nil, zero value otherwise.

### GetDiffVulnerabilityResultOk

`func (o *VulnUpdateNotificationPayloadAllOf) GetDiffVulnerabilityResultOk() (*VulnDiffResult, bool)`

GetDiffVulnerabilityResultOk returns a tuple with the DiffVulnerabilityResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffVulnerabilityResult

`func (o *VulnUpdateNotificationPayloadAllOf) SetDiffVulnerabilityResult(v VulnDiffResult)`

SetDiffVulnerabilityResult sets DiffVulnerabilityResult field to given value.

### HasDiffVulnerabilityResult

`func (o *VulnUpdateNotificationPayloadAllOf) HasDiffVulnerabilityResult() bool`

HasDiffVulnerabilityResult returns a boolean if a field has been set.

### GetImageDigest

`func (o *VulnUpdateNotificationPayloadAllOf) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *VulnUpdateNotificationPayloadAllOf) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *VulnUpdateNotificationPayloadAllOf) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *VulnUpdateNotificationPayloadAllOf) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetAnnotations

`func (o *VulnUpdateNotificationPayloadAllOf) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *VulnUpdateNotificationPayloadAllOf) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *VulnUpdateNotificationPayloadAllOf) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *VulnUpdateNotificationPayloadAllOf) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *VulnUpdateNotificationPayloadAllOf) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *VulnUpdateNotificationPayloadAllOf) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


