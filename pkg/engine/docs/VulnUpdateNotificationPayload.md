# VulnUpdateNotificationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**SubscriptionKey** | Pointer to **string** |  | [optional] 
**SubscriptionType** | Pointer to **string** |  | [optional] 
**NotificationId** | Pointer to **string** |  | [optional] 
**DiffVulnerabilityResult** | Pointer to [**VulnDiffResult**](VulnDiffResult.md) |  | [optional] 
**ImageDigest** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **interface{}** | List of Corresponding Image Annotations | [optional] 

## Methods

### NewVulnUpdateNotificationPayload

`func NewVulnUpdateNotificationPayload() *VulnUpdateNotificationPayload`

NewVulnUpdateNotificationPayload instantiates a new VulnUpdateNotificationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnUpdateNotificationPayloadWithDefaults

`func NewVulnUpdateNotificationPayloadWithDefaults() *VulnUpdateNotificationPayload`

NewVulnUpdateNotificationPayloadWithDefaults instantiates a new VulnUpdateNotificationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *VulnUpdateNotificationPayload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VulnUpdateNotificationPayload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VulnUpdateNotificationPayload) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *VulnUpdateNotificationPayload) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSubscriptionKey

`func (o *VulnUpdateNotificationPayload) GetSubscriptionKey() string`

GetSubscriptionKey returns the SubscriptionKey field if non-nil, zero value otherwise.

### GetSubscriptionKeyOk

`func (o *VulnUpdateNotificationPayload) GetSubscriptionKeyOk() (*string, bool)`

GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionKey

`func (o *VulnUpdateNotificationPayload) SetSubscriptionKey(v string)`

SetSubscriptionKey sets SubscriptionKey field to given value.

### HasSubscriptionKey

`func (o *VulnUpdateNotificationPayload) HasSubscriptionKey() bool`

HasSubscriptionKey returns a boolean if a field has been set.

### GetSubscriptionType

`func (o *VulnUpdateNotificationPayload) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *VulnUpdateNotificationPayload) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *VulnUpdateNotificationPayload) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *VulnUpdateNotificationPayload) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetNotificationId

`func (o *VulnUpdateNotificationPayload) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *VulnUpdateNotificationPayload) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *VulnUpdateNotificationPayload) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *VulnUpdateNotificationPayload) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetDiffVulnerabilityResult

`func (o *VulnUpdateNotificationPayload) GetDiffVulnerabilityResult() VulnDiffResult`

GetDiffVulnerabilityResult returns the DiffVulnerabilityResult field if non-nil, zero value otherwise.

### GetDiffVulnerabilityResultOk

`func (o *VulnUpdateNotificationPayload) GetDiffVulnerabilityResultOk() (*VulnDiffResult, bool)`

GetDiffVulnerabilityResultOk returns a tuple with the DiffVulnerabilityResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffVulnerabilityResult

`func (o *VulnUpdateNotificationPayload) SetDiffVulnerabilityResult(v VulnDiffResult)`

SetDiffVulnerabilityResult sets DiffVulnerabilityResult field to given value.

### HasDiffVulnerabilityResult

`func (o *VulnUpdateNotificationPayload) HasDiffVulnerabilityResult() bool`

HasDiffVulnerabilityResult returns a boolean if a field has been set.

### GetImageDigest

`func (o *VulnUpdateNotificationPayload) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *VulnUpdateNotificationPayload) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *VulnUpdateNotificationPayload) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *VulnUpdateNotificationPayload) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetAnnotations

`func (o *VulnUpdateNotificationPayload) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *VulnUpdateNotificationPayload) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *VulnUpdateNotificationPayload) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *VulnUpdateNotificationPayload) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *VulnUpdateNotificationPayload) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *VulnUpdateNotificationPayload) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


