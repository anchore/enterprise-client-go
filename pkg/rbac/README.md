# Go API client for rbac

Enterprise API for managing roles, permissions, and user mappings

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 5.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./rbac"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**AddIdp**](docs/DefaultApi.md#addidp) | **Post** /saml/idps | 
*DefaultApi* | [**AddRoleUser**](docs/DefaultApi.md#addroleuser) | **Post** /roles/{role_name}/members | Add a user to the role
*DefaultApi* | [**DeleteIdp**](docs/DefaultApi.md#deleteidp) | **Delete** /saml/idps/{name} | 
*DefaultApi* | [**DeleteRoleUser**](docs/DefaultApi.md#deleteroleuser) | **Delete** /roles/{role_name}/members | Remove a user from the role
*DefaultApi* | [**GetIdp**](docs/DefaultApi.md#getidp) | **Get** /saml/idps/{name} | 
*DefaultApi* | [**GetRole**](docs/DefaultApi.md#getrole) | **Get** /roles/{role_name} | Get detailed information about a specific role
*DefaultApi* | [**GetStatus**](docs/DefaultApi.md#getstatus) | **Get** /status | Service status
*DefaultApi* | [**HealthCheck**](docs/DefaultApi.md#healthcheck) | **Get** /health | 
*DefaultApi* | [**ListIdps**](docs/DefaultApi.md#listidps) | **Get** /saml/idps | 
*DefaultApi* | [**ListRoleMembers**](docs/DefaultApi.md#listrolemembers) | **Get** /roles/{role_name}/members | Returns a list of objects that have members in the role. The list is filtered by &#39;listRoleMembers&#39; access for the &#39;account&#39; element of each entry.
*DefaultApi* | [**ListRoles**](docs/DefaultApi.md#listroles) | **Get** /roles | List roles available in the system
*DefaultApi* | [**ListUserRoles**](docs/DefaultApi.md#listuserroles) | **Get** /users/{username}/roles | List the roles for which the requested user is a member
*DefaultApi* | [**MyRoles**](docs/DefaultApi.md#myroles) | **Get** /my_roles | List the roles for which the authenticated user is a member
*DefaultApi* | [**SamlLogin**](docs/DefaultApi.md#samllogin) | **Get** /saml/login/{idp_name} | 
*DefaultApi* | [**SamlSso**](docs/DefaultApi.md#samlsso) | **Post** /saml/sso/{idp_name} | 
*DefaultApi* | [**UpdateIdp**](docs/DefaultApi.md#updateidp) | **Put** /saml/idps/{name} | 
*DefaultApi* | [**VersionCheck**](docs/DefaultApi.md#versioncheck) | **Get** /version | 


## Documentation For Models

 - [AccountRole](docs/AccountRole.md)
 - [ApiErrorResponse](docs/ApiErrorResponse.md)
 - [Permission](docs/Permission.md)
 - [Role](docs/Role.md)
 - [RoleMember](docs/RoleMember.md)
 - [RoleMembership](docs/RoleMembership.md)
 - [RoleSummary](docs/RoleSummary.md)
 - [SamlConfiguration](docs/SamlConfiguration.md)
 - [ServiceVersion](docs/ServiceVersion.md)
 - [ServiceVersionApi](docs/ServiceVersionApi.md)
 - [ServiceVersionDb](docs/ServiceVersionDb.md)
 - [ServiceVersionEngine](docs/ServiceVersionEngine.md)
 - [ServiceVersionService](docs/ServiceVersionService.md)
 - [StatusResponse](docs/StatusResponse.md)
 - [TokenResponse](docs/TokenResponse.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

dev@anchore.com
