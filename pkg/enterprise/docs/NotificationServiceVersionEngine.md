# NotificationServiceVersionEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version of the installed engine library | [optional] 
**Db** | Pointer to **string** | Version of the installed engine db schema | [optional] 

## Methods

### NewNotificationServiceVersionEngine

`func NewNotificationServiceVersionEngine() *NotificationServiceVersionEngine`

NewNotificationServiceVersionEngine instantiates a new NotificationServiceVersionEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationServiceVersionEngineWithDefaults

`func NewNotificationServiceVersionEngineWithDefaults() *NotificationServiceVersionEngine`

NewNotificationServiceVersionEngineWithDefaults instantiates a new NotificationServiceVersionEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *NotificationServiceVersionEngine) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NotificationServiceVersionEngine) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NotificationServiceVersionEngine) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NotificationServiceVersionEngine) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDb

`func (o *NotificationServiceVersionEngine) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *NotificationServiceVersionEngine) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *NotificationServiceVersionEngine) SetDb(v string)`

SetDb sets Db field to given value.

### HasDb

`func (o *NotificationServiceVersionEngine) HasDb() bool`

HasDb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


