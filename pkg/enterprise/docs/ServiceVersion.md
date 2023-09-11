# ServiceVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**NotificationServiceVersionService**](NotificationServiceVersionService.md) |  | [optional] 
**Api** | Pointer to [**NotificationServiceVersionApi**](NotificationServiceVersionApi.md) |  | [optional] 
**Db** | Pointer to [**NotificationServiceVersionDb**](NotificationServiceVersionDb.md) |  | [optional] 

## Methods

### NewServiceVersion

`func NewServiceVersion() *ServiceVersion`

NewServiceVersion instantiates a new ServiceVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVersionWithDefaults

`func NewServiceVersionWithDefaults() *ServiceVersion`

NewServiceVersionWithDefaults instantiates a new ServiceVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ServiceVersion) GetService() NotificationServiceVersionService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceVersion) GetServiceOk() (*NotificationServiceVersionService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceVersion) SetService(v NotificationServiceVersionService)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceVersion) HasService() bool`

HasService returns a boolean if a field has been set.

### GetApi

`func (o *ServiceVersion) GetApi() NotificationServiceVersionApi`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *ServiceVersion) GetApiOk() (*NotificationServiceVersionApi, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *ServiceVersion) SetApi(v NotificationServiceVersionApi)`

SetApi sets Api field to given value.

### HasApi

`func (o *ServiceVersion) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetDb

`func (o *ServiceVersion) GetDb() NotificationServiceVersionDb`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *ServiceVersion) GetDbOk() (*NotificationServiceVersionDb, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *ServiceVersion) SetDb(v NotificationServiceVersionDb)`

SetDb sets Db field to given value.

### HasDb

`func (o *ServiceVersion) HasDb() bool`

HasDb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


