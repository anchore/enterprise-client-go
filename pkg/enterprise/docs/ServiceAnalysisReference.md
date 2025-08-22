# ServiceAnalysisReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | The unique id of the host on which the service is executing | [optional] 
**ServiceName** | Pointer to **string** | Registered service name | [optional] 
**AnalysisEngine** | Pointer to **string** | The generation of analysis engine that was used for the analysis | [optional] 
**EnterpriseVersion** | Pointer to **string** | The version of the Anchore Enterprise service that performed the analysis | [optional] 

## Methods

### NewServiceAnalysisReference

`func NewServiceAnalysisReference() *ServiceAnalysisReference`

NewServiceAnalysisReference instantiates a new ServiceAnalysisReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAnalysisReferenceWithDefaults

`func NewServiceAnalysisReferenceWithDefaults() *ServiceAnalysisReference`

NewServiceAnalysisReferenceWithDefaults instantiates a new ServiceAnalysisReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *ServiceAnalysisReference) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ServiceAnalysisReference) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ServiceAnalysisReference) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ServiceAnalysisReference) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceAnalysisReference) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceAnalysisReference) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceAnalysisReference) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceAnalysisReference) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetAnalysisEngine

`func (o *ServiceAnalysisReference) GetAnalysisEngine() string`

GetAnalysisEngine returns the AnalysisEngine field if non-nil, zero value otherwise.

### GetAnalysisEngineOk

`func (o *ServiceAnalysisReference) GetAnalysisEngineOk() (*string, bool)`

GetAnalysisEngineOk returns a tuple with the AnalysisEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisEngine

`func (o *ServiceAnalysisReference) SetAnalysisEngine(v string)`

SetAnalysisEngine sets AnalysisEngine field to given value.

### HasAnalysisEngine

`func (o *ServiceAnalysisReference) HasAnalysisEngine() bool`

HasAnalysisEngine returns a boolean if a field has been set.

### GetEnterpriseVersion

`func (o *ServiceAnalysisReference) GetEnterpriseVersion() string`

GetEnterpriseVersion returns the EnterpriseVersion field if non-nil, zero value otherwise.

### GetEnterpriseVersionOk

`func (o *ServiceAnalysisReference) GetEnterpriseVersionOk() (*string, bool)`

GetEnterpriseVersionOk returns a tuple with the EnterpriseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseVersion

`func (o *ServiceAnalysisReference) SetEnterpriseVersion(v string)`

SetEnterpriseVersion sets EnterpriseVersion field to given value.

### HasEnterpriseVersion

`func (o *ServiceAnalysisReference) HasEnterpriseVersion() bool`

HasEnterpriseVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


