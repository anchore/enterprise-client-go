/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type IntegrationsApi interface {

	/*
	DeleteIntegration Delete an integration instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param integrationUuid The integration instance's universally unique identifier
	@return ApiDeleteIntegrationRequest
	*/
	DeleteIntegration(ctx context.Context, integrationUuid string) ApiDeleteIntegrationRequest

	// DeleteIntegrationExecute executes the request
	DeleteIntegrationExecute(r ApiDeleteIntegrationRequest) (*http.Response, error)

	/*
	GetIntegrationById Get information about an integration instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param integrationUuid The integration instance's universally unique identifier
	@return ApiGetIntegrationByIdRequest
	*/
	GetIntegrationById(ctx context.Context, integrationUuid string) ApiGetIntegrationByIdRequest

	// GetIntegrationByIdExecute executes the request
	//  @return Integration
	GetIntegrationByIdExecute(r ApiGetIntegrationByIdRequest) (*Integration, *http.Response, error)

	/*
	HandleHealthReport Report health status for an integration

	An integration, such as an agent, will use this API to report its current health status. Such calls will typically be done on an interval, e.g., every 60 seconds.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param integrationUuid The integration's universally unique identifier
	@return ApiHandleHealthReportRequest
	*/
	HandleHealthReport(ctx context.Context, integrationUuid string) ApiHandleHealthReportRequest

	// HandleHealthReportExecute executes the request
	HandleHealthReportExecute(r ApiHandleHealthReportRequest) (*http.Response, error)

	/*
	ListIntegrations List known integration instances

	Lists all integrations that have been created (explicitly or indirectly via registration).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListIntegrationsRequest
	*/
	ListIntegrations(ctx context.Context) ApiListIntegrationsRequest

	// ListIntegrationsExecute executes the request
	//  @return IntegrationListResponse
	ListIntegrationsExecute(r ApiListIntegrationsRequest) (*IntegrationListResponse, *http.Response, error)

	/*
	RegisterIntegration Register an integration instance

	An integration instance, such as an agent, will use this API to register itself with Anchore. The integration registers using registration_id and registration_instance_id as temporary identifiers and is assigned its integration uuid upon successful registration. This method is idempotent in the sense that if an integration instance call it multiple time and uses the same registration_id and registration_instance_id, it will be assigned the same integration uuid.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRegisterIntegrationRequest
	*/
	RegisterIntegration(ctx context.Context) ApiRegisterIntegrationRequest

	// RegisterIntegrationExecute executes the request
	//  @return Integration
	RegisterIntegrationExecute(r ApiRegisterIntegrationRequest) (*Integration, *http.Response, error)
}

// IntegrationsApiService IntegrationsApi service
type IntegrationsApiService service

type ApiDeleteIntegrationRequest struct {
	ctx context.Context
	ApiService IntegrationsApi
	integrationUuid string
	force *bool
}

// Force deletion of the integration instance regardless of its state
func (r ApiDeleteIntegrationRequest) Force(force bool) ApiDeleteIntegrationRequest {
	r.force = &force
	return r
}

func (r ApiDeleteIntegrationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteIntegrationExecute(r)
}

/*
DeleteIntegration Delete an integration instance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param integrationUuid The integration instance's universally unique identifier
 @return ApiDeleteIntegrationRequest
*/
func (a *IntegrationsApiService) DeleteIntegration(ctx context.Context, integrationUuid string) ApiDeleteIntegrationRequest {
	return ApiDeleteIntegrationRequest{
		ApiService: a,
		ctx: ctx,
		integrationUuid: integrationUuid,
	}
}

// Execute executes the request
func (a *IntegrationsApiService) DeleteIntegrationExecute(r ApiDeleteIntegrationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsApiService.DeleteIntegration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integrations/{integration_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"integration_uuid"+"}", url.PathEscape(parameterToString(r.integrationUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.force != nil {
		localVarQueryParams.Add("force", parameterToString(*r.force, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetIntegrationByIdRequest struct {
	ctx context.Context
	ApiService IntegrationsApi
	integrationUuid string
}

func (r ApiGetIntegrationByIdRequest) Execute() (*Integration, *http.Response, error) {
	return r.ApiService.GetIntegrationByIdExecute(r)
}

/*
GetIntegrationById Get information about an integration instance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param integrationUuid The integration instance's universally unique identifier
 @return ApiGetIntegrationByIdRequest
*/
func (a *IntegrationsApiService) GetIntegrationById(ctx context.Context, integrationUuid string) ApiGetIntegrationByIdRequest {
	return ApiGetIntegrationByIdRequest{
		ApiService: a,
		ctx: ctx,
		integrationUuid: integrationUuid,
	}
}

// Execute executes the request
//  @return Integration
func (a *IntegrationsApiService) GetIntegrationByIdExecute(r ApiGetIntegrationByIdRequest) (*Integration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Integration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsApiService.GetIntegrationById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integrations/{integration_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"integration_uuid"+"}", url.PathEscape(parameterToString(r.integrationUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiHandleHealthReportRequest struct {
	ctx context.Context
	ApiService IntegrationsApi
	integrationUuid string
	healthReport *HealthReport
}

func (r ApiHandleHealthReportRequest) HealthReport(healthReport HealthReport) ApiHandleHealthReportRequest {
	r.healthReport = &healthReport
	return r
}

func (r ApiHandleHealthReportRequest) Execute() (*http.Response, error) {
	return r.ApiService.HandleHealthReportExecute(r)
}

/*
HandleHealthReport Report health status for an integration

An integration, such as an agent, will use this API to report its current health status. Such calls will typically be done on an interval, e.g., every 60 seconds.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param integrationUuid The integration's universally unique identifier
 @return ApiHandleHealthReportRequest
*/
func (a *IntegrationsApiService) HandleHealthReport(ctx context.Context, integrationUuid string) ApiHandleHealthReportRequest {
	return ApiHandleHealthReportRequest{
		ApiService: a,
		ctx: ctx,
		integrationUuid: integrationUuid,
	}
}

// Execute executes the request
func (a *IntegrationsApiService) HandleHealthReportExecute(r ApiHandleHealthReportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsApiService.HandleHealthReport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integrations/{integration_uuid}/health-report"
	localVarPath = strings.Replace(localVarPath, "{"+"integration_uuid"+"}", url.PathEscape(parameterToString(r.integrationUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.healthReport == nil {
		return nil, reportError("healthReport is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.healthReport
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiListIntegrationsRequest struct {
	ctx context.Context
	ApiService IntegrationsApi
	onlyDegraded *bool
}

// If true, limit listing to unhealthy or inactive integrations
func (r ApiListIntegrationsRequest) OnlyDegraded(onlyDegraded bool) ApiListIntegrationsRequest {
	r.onlyDegraded = &onlyDegraded
	return r
}

func (r ApiListIntegrationsRequest) Execute() (*IntegrationListResponse, *http.Response, error) {
	return r.ApiService.ListIntegrationsExecute(r)
}

/*
ListIntegrations List known integration instances

Lists all integrations that have been created (explicitly or indirectly via registration).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListIntegrationsRequest
*/
func (a *IntegrationsApiService) ListIntegrations(ctx context.Context) ApiListIntegrationsRequest {
	return ApiListIntegrationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IntegrationListResponse
func (a *IntegrationsApiService) ListIntegrationsExecute(r ApiListIntegrationsRequest) (*IntegrationListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IntegrationListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsApiService.ListIntegrations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integrations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.onlyDegraded != nil {
		localVarQueryParams.Add("only_degraded", parameterToString(*r.onlyDegraded, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRegisterIntegrationRequest struct {
	ctx context.Context
	ApiService IntegrationsApi
	integrationRegister *IntegrationRegister
}

func (r ApiRegisterIntegrationRequest) IntegrationRegister(integrationRegister IntegrationRegister) ApiRegisterIntegrationRequest {
	r.integrationRegister = &integrationRegister
	return r
}

func (r ApiRegisterIntegrationRequest) Execute() (*Integration, *http.Response, error) {
	return r.ApiService.RegisterIntegrationExecute(r)
}

/*
RegisterIntegration Register an integration instance

An integration instance, such as an agent, will use this API to register itself with Anchore. The integration registers using registration_id and registration_instance_id as temporary identifiers and is assigned its integration uuid upon successful registration. This method is idempotent in the sense that if an integration instance call it multiple time and uses the same registration_id and registration_instance_id, it will be assigned the same integration uuid.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRegisterIntegrationRequest
*/
func (a *IntegrationsApiService) RegisterIntegration(ctx context.Context) ApiRegisterIntegrationRequest {
	return ApiRegisterIntegrationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Integration
func (a *IntegrationsApiService) RegisterIntegrationExecute(r ApiRegisterIntegrationRequest) (*Integration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Integration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsApiService.RegisterIntegration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integrations/registration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.integrationRegister == nil {
		return localVarReturnValue, nil, reportError("integrationRegister is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.integrationRegister
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
