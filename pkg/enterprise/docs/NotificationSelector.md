# NotificationSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**ConfigurationUuid** | Pointer to **string** | UUID of the endpoint configuration bound to this selector | [optional] 
**Scope** | **string** | The scope to filter events. &#39;global&#39; scope encompasses all the events in the system, only the admin account can request this selector scope. &#39;account&#39; covers events scoped to a specific account. | 
**Event** | [**NotificationEventSelector**](NotificationEventSelector.md) |  | 

## Methods

### NewNotificationSelector

`func NewNotificationSelector(scope string, event NotificationEventSelector, ) *NotificationSelector`

NewNotificationSelector instantiates a new NotificationSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSelectorWithDefaults

`func NewNotificationSelectorWithDefaults() *NotificationSelector`

NewNotificationSelectorWithDefaults instantiates a new NotificationSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NotificationSelector) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NotificationSelector) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NotificationSelector) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NotificationSelector) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetConfigurationUuid

`func (o *NotificationSelector) GetConfigurationUuid() string`

GetConfigurationUuid returns the ConfigurationUuid field if non-nil, zero value otherwise.

### GetConfigurationUuidOk

`func (o *NotificationSelector) GetConfigurationUuidOk() (*string, bool)`

GetConfigurationUuidOk returns a tuple with the ConfigurationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationUuid

`func (o *NotificationSelector) SetConfigurationUuid(v string)`

SetConfigurationUuid sets ConfigurationUuid field to given value.

### HasConfigurationUuid

`func (o *NotificationSelector) HasConfigurationUuid() bool`

HasConfigurationUuid returns a boolean if a field has been set.

### GetScope

`func (o *NotificationSelector) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NotificationSelector) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NotificationSelector) SetScope(v string)`

SetScope sets Scope field to given value.


### GetEvent

`func (o *NotificationSelector) GetEvent() NotificationEventSelector`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NotificationSelector) GetEventOk() (*NotificationEventSelector, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NotificationSelector) SetEvent(v NotificationEventSelector)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


