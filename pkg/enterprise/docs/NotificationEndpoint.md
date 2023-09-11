# NotificationEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | the name of the endpoint | [optional] 
**Enabled** | Pointer to **bool** | Is the endpoint enabled for use in the system. Affects all usage, including system-level if set to false. | [optional] 
**LastUpdated** | Pointer to **time.Time** | The timestamp of the last change to the status | [optional] 

## Methods

### NewNotificationEndpoint

`func NewNotificationEndpoint() *NotificationEndpoint`

NewNotificationEndpoint instantiates a new NotificationEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEndpointWithDefaults

`func NewNotificationEndpointWithDefaults() *NotificationEndpoint`

NewNotificationEndpointWithDefaults instantiates a new NotificationEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *NotificationEndpoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationEndpoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationEndpoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NotificationEndpoint) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NotificationEndpoint) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NotificationEndpoint) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NotificationEndpoint) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NotificationEndpoint) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


