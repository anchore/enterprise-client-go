# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | The unique id of the host on which the service is executing | [optional] 
**ServiceName** | Pointer to **string** | Registered service name | [optional] 
**BaseUrl** | Pointer to **string** | The url to reach the service, including port as needed | [optional] 
**StatusMessage** | Pointer to **string** | A state indicating the condition of the service. Normal operation is &#39;registered&#39; | [optional] 
**ServiceDetail** | Pointer to [**StatusResponse**](StatusResponse.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** | The version of the service as reported by the service implementation on registration | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *Service) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *Service) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *Service) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *Service) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetServiceName

`func (o *Service) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Service) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Service) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *Service) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetBaseUrl

`func (o *Service) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *Service) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *Service) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *Service) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Service) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Service) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Service) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Service) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetServiceDetail

`func (o *Service) GetServiceDetail() StatusResponse`

GetServiceDetail returns the ServiceDetail field if non-nil, zero value otherwise.

### GetServiceDetailOk

`func (o *Service) GetServiceDetailOk() (*StatusResponse, bool)`

GetServiceDetailOk returns a tuple with the ServiceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDetail

`func (o *Service) SetServiceDetail(v StatusResponse)`

SetServiceDetail sets ServiceDetail field to given value.

### HasServiceDetail

`func (o *Service) HasServiceDetail() bool`

HasServiceDetail returns a boolean if a field has been set.

### GetStatus

`func (o *Service) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Service) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Service) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Service) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *Service) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Service) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Service) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Service) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


