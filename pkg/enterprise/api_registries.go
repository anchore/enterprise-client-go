/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.4.0
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


type RegistriesApi interface {

	/*
	CreateRegistry Add a new registry

	Adds a new registry to the system

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateRegistryRequest
	*/
	CreateRegistry(ctx context.Context) ApiCreateRegistryRequest

	// CreateRegistryExecute executes the request
	//  @return []RegistryConfiguration
	CreateRegistryExecute(r ApiCreateRegistryRequest) ([]RegistryConfiguration, *http.Response, error)

	/*
	DeleteRegistry Delete a registry configuration

	Delete a registry configuration record from the system. Does not remove any images.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param registry
	@return ApiDeleteRegistryRequest
	*/
	DeleteRegistry(ctx context.Context, registry string) ApiDeleteRegistryRequest

	// DeleteRegistryExecute executes the request
	DeleteRegistryExecute(r ApiDeleteRegistryRequest) (*http.Response, error)

	/*
	GetRegistry Get a specific registry configuration

	Get information on a specific registry

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param registry
	@return ApiGetRegistryRequest
	*/
	GetRegistry(ctx context.Context, registry string) ApiGetRegistryRequest

	// GetRegistryExecute executes the request
	//  @return []RegistryConfiguration
	GetRegistryExecute(r ApiGetRegistryRequest) ([]RegistryConfiguration, *http.Response, error)

	/*
	ListRegistries List configured registries

	List all configured registries the system can/will watch

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListRegistriesRequest
	*/
	ListRegistries(ctx context.Context) ApiListRegistriesRequest

	// ListRegistriesExecute executes the request
	//  @return []RegistryConfiguration
	ListRegistriesExecute(r ApiListRegistriesRequest) ([]RegistryConfiguration, *http.Response, error)

	/*
	UpdateRegistry Update/replace a registry configuration

	Replaces an existing registry record with the given record

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param registry
	@return ApiUpdateRegistryRequest
	*/
	UpdateRegistry(ctx context.Context, registry string) ApiUpdateRegistryRequest

	// UpdateRegistryExecute executes the request
	//  @return []RegistryConfiguration
	UpdateRegistryExecute(r ApiUpdateRegistryRequest) ([]RegistryConfiguration, *http.Response, error)
}

// RegistriesApiService RegistriesApi service
type RegistriesApiService service

type ApiCreateRegistryRequest struct {
	ctx context.Context
	ApiService RegistriesApi
	registryData *RegistryConfigurationRequest
	validate *bool
	xAnchoreAccount *string
}

func (r ApiCreateRegistryRequest) RegistryData(registryData RegistryConfigurationRequest) ApiCreateRegistryRequest {
	r.registryData = &registryData
	return r
}

// flag to determine whether or not to validate registry/credential at registry add time
func (r ApiCreateRegistryRequest) Validate(validate bool) ApiCreateRegistryRequest {
	r.validate = &validate
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiCreateRegistryRequest) XAnchoreAccount(xAnchoreAccount string) ApiCreateRegistryRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiCreateRegistryRequest) Execute() ([]RegistryConfiguration, *http.Response, error) {
	return r.ApiService.CreateRegistryExecute(r)
}

/*
CreateRegistry Add a new registry

Adds a new registry to the system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateRegistryRequest
*/
func (a *RegistriesApiService) CreateRegistry(ctx context.Context) ApiCreateRegistryRequest {
	return ApiCreateRegistryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []RegistryConfiguration
func (a *RegistriesApiService) CreateRegistryExecute(r ApiCreateRegistryRequest) ([]RegistryConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RegistryConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistriesApiService.CreateRegistry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.registryData == nil {
		return localVarReturnValue, nil, reportError("registryData is required and must be specified")
	}

	if r.validate != nil {
		localVarQueryParams.Add("validate", parameterToString(*r.validate, ""))
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
	if r.xAnchoreAccount != nil {
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
	}
	// body params
	localVarPostBody = r.registryData
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

type ApiDeleteRegistryRequest struct {
	ctx context.Context
	ApiService RegistriesApi
	registry string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiDeleteRegistryRequest) XAnchoreAccount(xAnchoreAccount string) ApiDeleteRegistryRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiDeleteRegistryRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRegistryExecute(r)
}

/*
DeleteRegistry Delete a registry configuration

Delete a registry configuration record from the system. Does not remove any images.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registry
 @return ApiDeleteRegistryRequest
*/
func (a *RegistriesApiService) DeleteRegistry(ctx context.Context, registry string) ApiDeleteRegistryRequest {
	return ApiDeleteRegistryRequest{
		ApiService: a,
		ctx: ctx,
		registry: registry,
	}
}

// Execute executes the request
func (a *RegistriesApiService) DeleteRegistryExecute(r ApiDeleteRegistryRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistriesApiService.DeleteRegistry")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registries/{registry}"
	localVarPath = strings.Replace(localVarPath, "{"+"registry"+"}", url.PathEscape(parameterToString(r.registry, "")), -1)

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
	if r.xAnchoreAccount != nil {
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
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

type ApiGetRegistryRequest struct {
	ctx context.Context
	ApiService RegistriesApi
	registry string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetRegistryRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetRegistryRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetRegistryRequest) Execute() ([]RegistryConfiguration, *http.Response, error) {
	return r.ApiService.GetRegistryExecute(r)
}

/*
GetRegistry Get a specific registry configuration

Get information on a specific registry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registry
 @return ApiGetRegistryRequest
*/
func (a *RegistriesApiService) GetRegistry(ctx context.Context, registry string) ApiGetRegistryRequest {
	return ApiGetRegistryRequest{
		ApiService: a,
		ctx: ctx,
		registry: registry,
	}
}

// Execute executes the request
//  @return []RegistryConfiguration
func (a *RegistriesApiService) GetRegistryExecute(r ApiGetRegistryRequest) ([]RegistryConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RegistryConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistriesApiService.GetRegistry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registries/{registry}"
	localVarPath = strings.Replace(localVarPath, "{"+"registry"+"}", url.PathEscape(parameterToString(r.registry, "")), -1)

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
	if r.xAnchoreAccount != nil {
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
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

type ApiListRegistriesRequest struct {
	ctx context.Context
	ApiService RegistriesApi
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiListRegistriesRequest) XAnchoreAccount(xAnchoreAccount string) ApiListRegistriesRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiListRegistriesRequest) Execute() ([]RegistryConfiguration, *http.Response, error) {
	return r.ApiService.ListRegistriesExecute(r)
}

/*
ListRegistries List configured registries

List all configured registries the system can/will watch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListRegistriesRequest
*/
func (a *RegistriesApiService) ListRegistries(ctx context.Context) ApiListRegistriesRequest {
	return ApiListRegistriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []RegistryConfiguration
func (a *RegistriesApiService) ListRegistriesExecute(r ApiListRegistriesRequest) ([]RegistryConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RegistryConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistriesApiService.ListRegistries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registries"

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
	if r.xAnchoreAccount != nil {
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
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

type ApiUpdateRegistryRequest struct {
	ctx context.Context
	ApiService RegistriesApi
	registry string
	registryData *RegistryConfigurationRequest
	validate *bool
	xAnchoreAccount *string
}

func (r ApiUpdateRegistryRequest) RegistryData(registryData RegistryConfigurationRequest) ApiUpdateRegistryRequest {
	r.registryData = &registryData
	return r
}

// flag to determine whether or not to validate registry/credential at registry update time
func (r ApiUpdateRegistryRequest) Validate(validate bool) ApiUpdateRegistryRequest {
	r.validate = &validate
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiUpdateRegistryRequest) XAnchoreAccount(xAnchoreAccount string) ApiUpdateRegistryRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiUpdateRegistryRequest) Execute() ([]RegistryConfiguration, *http.Response, error) {
	return r.ApiService.UpdateRegistryExecute(r)
}

/*
UpdateRegistry Update/replace a registry configuration

Replaces an existing registry record with the given record

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registry
 @return ApiUpdateRegistryRequest
*/
func (a *RegistriesApiService) UpdateRegistry(ctx context.Context, registry string) ApiUpdateRegistryRequest {
	return ApiUpdateRegistryRequest{
		ApiService: a,
		ctx: ctx,
		registry: registry,
	}
}

// Execute executes the request
//  @return []RegistryConfiguration
func (a *RegistriesApiService) UpdateRegistryExecute(r ApiUpdateRegistryRequest) ([]RegistryConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RegistryConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistriesApiService.UpdateRegistry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registries/{registry}"
	localVarPath = strings.Replace(localVarPath, "{"+"registry"+"}", url.PathEscape(parameterToString(r.registry, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.registryData == nil {
		return localVarReturnValue, nil, reportError("registryData is required and must be specified")
	}

	if r.validate != nil {
		localVarQueryParams.Add("validate", parameterToString(*r.validate, ""))
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
	if r.xAnchoreAccount != nil {
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
	}
	// body params
	localVarPostBody = r.registryData
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
