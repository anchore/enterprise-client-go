# InventoryClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialType** | Pointer to **string** |  | [optional] 
**Namespaces** | Pointer to **[]string** |  | [optional] 
**ClusterServer** | Pointer to **string** | FQDN of the cluster API server | [optional] 
**ClusterCertificate** | Pointer to **string** | Base64 Encoded Public Certificate for the cluster | [optional] 
**ClientCertificate** | Pointer to **string** | Base64 Encoded Public Certificate for the client. Not required if credential_type &#x3D;&#x3D; token | [optional] 
**Credential** | Pointer to **string** | Base64 Encoded credential for the client | [optional] 

## Methods

### NewInventoryClusterConfig

`func NewInventoryClusterConfig() *InventoryClusterConfig`

NewInventoryClusterConfig instantiates a new InventoryClusterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryClusterConfigWithDefaults

`func NewInventoryClusterConfigWithDefaults() *InventoryClusterConfig`

NewInventoryClusterConfigWithDefaults instantiates a new InventoryClusterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialType

`func (o *InventoryClusterConfig) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *InventoryClusterConfig) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *InventoryClusterConfig) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *InventoryClusterConfig) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetNamespaces

`func (o *InventoryClusterConfig) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *InventoryClusterConfig) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *InventoryClusterConfig) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *InventoryClusterConfig) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetClusterServer

`func (o *InventoryClusterConfig) GetClusterServer() string`

GetClusterServer returns the ClusterServer field if non-nil, zero value otherwise.

### GetClusterServerOk

`func (o *InventoryClusterConfig) GetClusterServerOk() (*string, bool)`

GetClusterServerOk returns a tuple with the ClusterServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterServer

`func (o *InventoryClusterConfig) SetClusterServer(v string)`

SetClusterServer sets ClusterServer field to given value.

### HasClusterServer

`func (o *InventoryClusterConfig) HasClusterServer() bool`

HasClusterServer returns a boolean if a field has been set.

### GetClusterCertificate

`func (o *InventoryClusterConfig) GetClusterCertificate() string`

GetClusterCertificate returns the ClusterCertificate field if non-nil, zero value otherwise.

### GetClusterCertificateOk

`func (o *InventoryClusterConfig) GetClusterCertificateOk() (*string, bool)`

GetClusterCertificateOk returns a tuple with the ClusterCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCertificate

`func (o *InventoryClusterConfig) SetClusterCertificate(v string)`

SetClusterCertificate sets ClusterCertificate field to given value.

### HasClusterCertificate

`func (o *InventoryClusterConfig) HasClusterCertificate() bool`

HasClusterCertificate returns a boolean if a field has been set.

### GetClientCertificate

`func (o *InventoryClusterConfig) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *InventoryClusterConfig) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *InventoryClusterConfig) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *InventoryClusterConfig) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetCredential

`func (o *InventoryClusterConfig) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *InventoryClusterConfig) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *InventoryClusterConfig) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *InventoryClusterConfig) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


