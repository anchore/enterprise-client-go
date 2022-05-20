# SystemStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceStates** | Pointer to [**[]Service**](Service.md) | A list of service objects | [optional] 

## Methods

### NewSystemStatusResponse

`func NewSystemStatusResponse() *SystemStatusResponse`

NewSystemStatusResponse instantiates a new SystemStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemStatusResponseWithDefaults

`func NewSystemStatusResponseWithDefaults() *SystemStatusResponse`

NewSystemStatusResponseWithDefaults instantiates a new SystemStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceStates

`func (o *SystemStatusResponse) GetServiceStates() []Service`

GetServiceStates returns the ServiceStates field if non-nil, zero value otherwise.

### GetServiceStatesOk

`func (o *SystemStatusResponse) GetServiceStatesOk() (*[]Service, bool)`

GetServiceStatesOk returns a tuple with the ServiceStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStates

`func (o *SystemStatusResponse) SetServiceStates(v []Service)`

SetServiceStates sets ServiceStates field to given value.

### HasServiceStates

`func (o *SystemStatusResponse) HasServiceStates() bool`

HasServiceStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


