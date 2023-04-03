# ServiceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostid** | Pointer to **string** | The unique id of the host on which the service is executing | [optional] 
**Servicename** | Pointer to **string** | Registered service name | [optional] 

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

### GetHostid

`func (o *ServiceReference) GetHostid() string`

GetHostid returns the Hostid field if non-nil, zero value otherwise.

### GetHostidOk

`func (o *ServiceReference) GetHostidOk() (*string, bool)`

GetHostidOk returns a tuple with the Hostid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostid

`func (o *ServiceReference) SetHostid(v string)`

SetHostid sets Hostid field to given value.

### HasHostid

`func (o *ServiceReference) HasHostid() bool`

HasHostid returns a boolean if a field has been set.

### GetServicename

`func (o *ServiceReference) GetServicename() string`

GetServicename returns the Servicename field if non-nil, zero value otherwise.

### GetServicenameOk

`func (o *ServiceReference) GetServicenameOk() (*string, bool)`

GetServicenameOk returns a tuple with the Servicename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicename

`func (o *ServiceReference) SetServicename(v string)`

SetServicename sets Servicename field to given value.

### HasServicename

`func (o *ServiceReference) HasServicename() bool`

HasServicename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


