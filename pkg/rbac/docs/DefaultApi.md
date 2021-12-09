# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIdp**](DefaultApi.md#AddIdp) | **Post** /saml/idps | 
[**AddRoleUser**](DefaultApi.md#AddRoleUser) | **Post** /roles/{rolename}/members | Add a user to the role
[**DeleteIdp**](DefaultApi.md#DeleteIdp) | **Delete** /saml/idps/{name} | 
[**DeleteRoleUser**](DefaultApi.md#DeleteRoleUser) | **Delete** /roles/{rolename}/members | Remove a user from the role
[**GetIdp**](DefaultApi.md#GetIdp) | **Get** /saml/idps/{name} | 
[**GetRole**](DefaultApi.md#GetRole) | **Get** /roles/{rolename} | Get detailed information about a specific role
[**GetStatus**](DefaultApi.md#GetStatus) | **Get** /status | Service status
[**HealthCheck**](DefaultApi.md#HealthCheck) | **Get** /health | 
[**ListIdps**](DefaultApi.md#ListIdps) | **Get** /saml/idps | 
[**ListRoleMembers**](DefaultApi.md#ListRoleMembers) | **Get** /roles/{rolename}/members | Returns a list of objects that have members in the role. The list is filtered by &#39;listRoleMembers&#39; access for the &#39;account&#39; element of each entry.
[**ListRoles**](DefaultApi.md#ListRoles) | **Get** /roles | List roles available in the system
[**ListUserRoles**](DefaultApi.md#ListUserRoles) | **Get** /users/{username}/roles | List the roles for which the requested user is a member
[**MyRoles**](DefaultApi.md#MyRoles) | **Get** /my_roles | List the roles for which the authenticated user is a member
[**SamlLogin**](DefaultApi.md#SamlLogin) | **Get** /saml/login/{idp_name} | 
[**SamlSso**](DefaultApi.md#SamlSso) | **Post** /saml/sso/{idp_name} | 
[**UpdateIdp**](DefaultApi.md#UpdateIdp) | **Put** /saml/idps/{name} | 
[**VersionCheck**](DefaultApi.md#VersionCheck) | **Get** /version | 



## AddIdp

> SamlConfiguration AddIdp(ctx, configuration)



Add a new Identity Provider to the system, with a specific name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configuration** | [**SamlConfiguration**](SamlConfiguration.md)|  | 

### Return type

[**SamlConfiguration**](SamlConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoleUser

> RoleMember AddRoleUser(ctx, rolename, member)

Add a user to the role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rolename** | **string**|  | 
**member** | [**RoleMember**](RoleMember.md)|  | 

### Return type

[**RoleMember**](RoleMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdp

> DeleteIdp(ctx, name)



Delete an idp configuration. Users will not longer be able to login from this idp

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleUser

> DeleteRoleUser(ctx, rolename, username, forAccount)

Remove a user from the role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rolename** | **string**|  | 
**username** | **string**| The username to remove the role for | 
**forAccount** | **string**| The account that the user has the role to be removed | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdp

> SamlConfiguration GetIdp(ctx, name)



Return the configuration for a named Identity Provider

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

### Return type

[**SamlConfiguration**](SamlConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> Role GetRole(ctx, rolename)

Get detailed information about a specific role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rolename** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> StatusResponse GetStatus(ctx, )

Service status

Get the API service status

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthCheck

> HealthCheck(ctx, )



Health check, returns 200 and no body if service is running

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdps

> []string ListIdps(ctx, )



List the names of configured Identity Providers for this anchore installation

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleMembers

> []RoleMember ListRoleMembers(ctx, rolename, optional)

Returns a list of objects that have members in the role. The list is filtered by 'listRoleMembers' access for the 'account' element of each entry.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rolename** | **string**|  | 
 **optional** | ***ListRoleMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRoleMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forAccount** | **optional.String**| Optional filter paramter to limit the set fo returned items to only those with matching account. Will return Access Denied if caller does not have permission to listRoleMembers for that account. | 

### Return type

[**[]RoleMember**](RoleMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> []RoleSummary ListRoles(ctx, )

List roles available in the system

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]RoleSummary**](RoleSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserRoles

> []RoleMembership ListUserRoles(ctx, username, optional)

List the roles for which the requested user is a member

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**|  | 
 **optional** | ***ListUserRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUserRolesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forAccount** | **optional.String**|  | 
 **role** | **optional.String**|  | 

### Return type

[**[]RoleMembership**](RoleMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyRoles

> []AccountRole MyRoles(ctx, )

List the roles for which the authenticated user is a member

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AccountRole**](AccountRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SamlLogin

> TokenResponse SamlLogin(ctx, idpName)



Initiate an SP-initiated login sequence for the Idp. The SP will respond with the SAML AuthN Request the client must send to the Idp URL

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpName** | **string**|  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SamlSso

> TokenResponse SamlSso(ctx, idpName)



Perform a login using a SAML assertion, no HTTP auth is required as the SAML assertion is considered the authenticating token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpName** | **string**|  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdp

> SamlConfiguration UpdateIdp(ctx, name, configuration)



Update an existing Identity Provider configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 
**configuration** | [**SamlConfiguration**](SamlConfiguration.md)|  | 

### Return type

[**SamlConfiguration**](SamlConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionCheck

> ServiceVersion VersionCheck(ctx, )



Returns the version object for the service, including db schema version info

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ServiceVersion**](ServiceVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

