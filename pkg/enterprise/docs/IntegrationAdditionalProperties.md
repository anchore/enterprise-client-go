# IntegrationAdditionalProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Short description of the integration instance | [optional] 
**StartedAt** | Pointer to **time.Time** | Timestamp when integration instance was started | [optional] 
**Namespaces** | Pointer to **[]string** | List of namespaces handled by the integration instance | [optional] 
**Configuration** | Pointer to **interface{}** | Configuration of the integration instance | [optional] 

## Methods

### NewIntegrationAdditionalProperties

`func NewIntegrationAdditionalProperties() *IntegrationAdditionalProperties`

NewIntegrationAdditionalProperties instantiates a new IntegrationAdditionalProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationAdditionalPropertiesWithDefaults

`func NewIntegrationAdditionalPropertiesWithDefaults() *IntegrationAdditionalProperties`

NewIntegrationAdditionalPropertiesWithDefaults instantiates a new IntegrationAdditionalProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IntegrationAdditionalProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationAdditionalProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationAdditionalProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationAdditionalProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartedAt

`func (o *IntegrationAdditionalProperties) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *IntegrationAdditionalProperties) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *IntegrationAdditionalProperties) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *IntegrationAdditionalProperties) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetNamespaces

`func (o *IntegrationAdditionalProperties) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *IntegrationAdditionalProperties) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *IntegrationAdditionalProperties) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *IntegrationAdditionalProperties) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetConfiguration

`func (o *IntegrationAdditionalProperties) GetConfiguration() interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IntegrationAdditionalProperties) GetConfigurationOk() (*interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IntegrationAdditionalProperties) SetConfiguration(v interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *IntegrationAdditionalProperties) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


