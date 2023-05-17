# NotificationEndpointEnabledStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Is the endpoint enabled for use in the system. Affects all usage, including system-level if set to false. | [optional] 
**LastUpdated** | Pointer to **time.Time** | The timestamp of the last change to the status | [optional] 

## Methods

### NewNotificationEndpointEnabledStatus

`func NewNotificationEndpointEnabledStatus() *NotificationEndpointEnabledStatus`

NewNotificationEndpointEnabledStatus instantiates a new NotificationEndpointEnabledStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEndpointEnabledStatusWithDefaults

`func NewNotificationEndpointEnabledStatusWithDefaults() *NotificationEndpointEnabledStatus`

NewNotificationEndpointEnabledStatusWithDefaults instantiates a new NotificationEndpointEnabledStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NotificationEndpointEnabledStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationEndpointEnabledStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationEndpointEnabledStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NotificationEndpointEnabledStatus) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NotificationEndpointEnabledStatus) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NotificationEndpointEnabledStatus) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NotificationEndpointEnabledStatus) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NotificationEndpointEnabledStatus) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


