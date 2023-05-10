# ServiceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | The unique id of the host on which the service is executing | [optional] 
**ServiceName** | Pointer to **string** | Registered service name | [optional] 

## Methods

### NewServiceReference

`func NewServiceReference() *ServiceReference`

NewServiceReference instantiates a new ServiceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceReferenceWithDefaults

`func NewServiceReferenceWithDefaults() *ServiceReference`

NewServiceReferenceWithDefaults instantiates a new ServiceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *ServiceReference) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ServiceReference) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ServiceReference) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ServiceReference) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceReference) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceReference) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceReference) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceReference) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


