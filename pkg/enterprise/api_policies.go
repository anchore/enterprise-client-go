/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
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


type PoliciesApi interface {

	/*
	AddPolicy Add a new policy

	Adds a new policy to the system

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddPolicyRequest
	*/
	AddPolicy(ctx context.Context) ApiAddPolicyRequest

	// AddPolicyExecute executes the request
	//  @return PolicyRecord
<<<<<<< HEAD
	AddPolicyExecute(r ApiAddPolicyRequest) (*PolicyRecord, *http.Response, error)
=======
	AddPolicyExecute(r ApiAddPolicyRequest) (PolicyRecord, *_nethttp.Response, error)
>>>>>>> main

	/*
	DeletePolicy Delete policy

	Delete the specified policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param policyId
	@return ApiDeletePolicyRequest
	*/
	DeletePolicy(ctx context.Context, policyId string) ApiDeletePolicyRequest

	// DeletePolicyExecute executes the request
	DeletePolicyExecute(r ApiDeletePolicyRequest) (*http.Response, error)

	/*
	GetPolicy Get specific policy

	Get the policy content

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param policyId
	@return ApiGetPolicyRequest
	*/
	GetPolicy(ctx context.Context, policyId string) ApiGetPolicyRequest

	// GetPolicyExecute executes the request
	//  @return PolicyRecord
<<<<<<< HEAD
	GetPolicyExecute(r ApiGetPolicyRequest) (*PolicyRecord, *http.Response, error)
=======
	GetPolicyExecute(r ApiGetPolicyRequest) (PolicyRecord, *_nethttp.Response, error)
>>>>>>> main

	/*
	ListPolicies List policies

	List all saved policies

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListPoliciesRequest
	*/
	ListPolicies(ctx context.Context) ApiListPoliciesRequest

	// ListPoliciesExecute executes the request
	//  @return []PolicyRecord
<<<<<<< HEAD
	ListPoliciesExecute(r ApiListPoliciesRequest) ([]PolicyRecord, *http.Response, error)
=======
	ListPoliciesExecute(r ApiListPoliciesRequest) ([]PolicyRecord, *_nethttp.Response, error)
>>>>>>> main

	/*
	UpdatePolicy Update policy

	Update/replace and existing policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param policyId
	@return ApiUpdatePolicyRequest
	*/
	UpdatePolicy(ctx context.Context, policyId string) ApiUpdatePolicyRequest

	// UpdatePolicyExecute executes the request
	//  @return PolicyRecord
<<<<<<< HEAD
	UpdatePolicyExecute(r ApiUpdatePolicyRequest) (*PolicyRecord, *http.Response, error)
=======
	UpdatePolicyExecute(r ApiUpdatePolicyRequest) (PolicyRecord, *_nethttp.Response, error)
>>>>>>> main
}

// PoliciesApiService PoliciesApi service
type PoliciesApiService service

type ApiAddPolicyRequest struct {
	ctx context.Context
	ApiService PoliciesApi
	policy *Policy
	xAnchoreAccount *string
}

func (r ApiAddPolicyRequest) Policy(policy Policy) ApiAddPolicyRequest {
	r.policy = &policy
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiAddPolicyRequest) XAnchoreAccount(xAnchoreAccount string) ApiAddPolicyRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

<<<<<<< HEAD
func (r ApiAddPolicyRequest) Execute() (*PolicyRecord, *http.Response, error) {
=======
func (r ApiAddPolicyRequest) Execute() (PolicyRecord, *_nethttp.Response, error) {
>>>>>>> main
	return r.ApiService.AddPolicyExecute(r)
}

/*
AddPolicy Add a new policy

Adds a new policy to the system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddPolicyRequest
*/
func (a *PoliciesApiService) AddPolicy(ctx context.Context) ApiAddPolicyRequest {
	return ApiAddPolicyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PolicyRecord
<<<<<<< HEAD
func (a *PoliciesApiService) AddPolicyExecute(r ApiAddPolicyRequest) (*PolicyRecord, *http.Response, error) {
=======
func (a *PoliciesApiService) AddPolicyExecute(r ApiAddPolicyRequest) (PolicyRecord, *_nethttp.Response, error) {
>>>>>>> main
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
<<<<<<< HEAD
		formFiles            []formFile
		localVarReturnValue  *PolicyRecord
=======
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PolicyRecord
>>>>>>> main
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesApiService.AddPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies"

	localVarHeaderParams := make(map[string]string)
<<<<<<< HEAD
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
=======
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
>>>>>>> main
	if r.policy == nil {
		return localVarReturnValue, nil, reportError("policy is required and must be specified")
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
	localVarPostBody = r.policy
<<<<<<< HEAD
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
=======
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
>>>>>>> main
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

type ApiDeletePolicyRequest struct {
	ctx context.Context
	ApiService PoliciesApi
	policyId string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiDeletePolicyRequest) XAnchoreAccount(xAnchoreAccount string) ApiDeletePolicyRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiDeletePolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePolicyExecute(r)
}

/*
DeletePolicy Delete policy

Delete the specified policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId
 @return ApiDeletePolicyRequest
*/
func (a *PoliciesApiService) DeletePolicy(ctx context.Context, policyId string) ApiDeletePolicyRequest {
	return ApiDeletePolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
func (a *PoliciesApiService) DeletePolicyExecute(r ApiDeletePolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesApiService.DeletePolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{policy_id}"
<<<<<<< HEAD
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)
=======
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", _neturl.PathEscape(parameterToString(r.policyId, "")), -1)
>>>>>>> main

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
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

type ApiGetPolicyRequest struct {
	ctx context.Context
	ApiService PoliciesApi
	policyId string
	detail *bool
	xAnchoreAccount *string
}

// Include policy detail in the form of the full policy content for each entry
func (r ApiGetPolicyRequest) Detail(detail bool) ApiGetPolicyRequest {
	r.detail = &detail
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetPolicyRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetPolicyRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

<<<<<<< HEAD
func (r ApiGetPolicyRequest) Execute() (*PolicyRecord, *http.Response, error) {
=======
func (r ApiGetPolicyRequest) Execute() (PolicyRecord, *_nethttp.Response, error) {
>>>>>>> main
	return r.ApiService.GetPolicyExecute(r)
}

/*
GetPolicy Get specific policy

Get the policy content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId
 @return ApiGetPolicyRequest
*/
func (a *PoliciesApiService) GetPolicy(ctx context.Context, policyId string) ApiGetPolicyRequest {
	return ApiGetPolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return PolicyRecord
<<<<<<< HEAD
func (a *PoliciesApiService) GetPolicyExecute(r ApiGetPolicyRequest) (*PolicyRecord, *http.Response, error) {
=======
func (a *PoliciesApiService) GetPolicyExecute(r ApiGetPolicyRequest) (PolicyRecord, *_nethttp.Response, error) {
>>>>>>> main
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
<<<<<<< HEAD
		formFiles            []formFile
		localVarReturnValue  *PolicyRecord
=======
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PolicyRecord
>>>>>>> main
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesApiService.GetPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{policy_id}"
<<<<<<< HEAD
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)
=======
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", _neturl.PathEscape(parameterToString(r.policyId, "")), -1)
>>>>>>> main

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.detail != nil {
		localVarQueryParams.Add("detail", parameterToString(*r.detail, ""))
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

type ApiListPoliciesRequest struct {
	ctx context.Context
	ApiService PoliciesApi
	detail *bool
	xAnchoreAccount *string
}

// Include policy detail in the form of the full policy content for each entry
func (r ApiListPoliciesRequest) Detail(detail bool) ApiListPoliciesRequest {
	r.detail = &detail
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiListPoliciesRequest) XAnchoreAccount(xAnchoreAccount string) ApiListPoliciesRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

<<<<<<< HEAD
func (r ApiListPoliciesRequest) Execute() ([]PolicyRecord, *http.Response, error) {
=======
func (r ApiListPoliciesRequest) Execute() ([]PolicyRecord, *_nethttp.Response, error) {
>>>>>>> main
	return r.ApiService.ListPoliciesExecute(r)
}

/*
ListPolicies List policies

List all saved policies

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPoliciesRequest
*/
func (a *PoliciesApiService) ListPolicies(ctx context.Context) ApiListPoliciesRequest {
	return ApiListPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []PolicyRecord
<<<<<<< HEAD
func (a *PoliciesApiService) ListPoliciesExecute(r ApiListPoliciesRequest) ([]PolicyRecord, *http.Response, error) {
=======
func (a *PoliciesApiService) ListPoliciesExecute(r ApiListPoliciesRequest) ([]PolicyRecord, *_nethttp.Response, error) {
>>>>>>> main
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
<<<<<<< HEAD
		formFiles            []formFile
=======
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
>>>>>>> main
		localVarReturnValue  []PolicyRecord
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesApiService.ListPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.detail != nil {
		localVarQueryParams.Add("detail", parameterToString(*r.detail, ""))
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

type ApiUpdatePolicyRequest struct {
	ctx context.Context
	ApiService PoliciesApi
	policyId string
	policy *PolicyRecord
	active *bool
	xAnchoreAccount *string
}

func (r ApiUpdatePolicyRequest) Policy(policy PolicyRecord) ApiUpdatePolicyRequest {
	r.policy = &policy
	return r
}

// Mark policy as active
func (r ApiUpdatePolicyRequest) Active(active bool) ApiUpdatePolicyRequest {
	r.active = &active
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiUpdatePolicyRequest) XAnchoreAccount(xAnchoreAccount string) ApiUpdatePolicyRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

<<<<<<< HEAD
func (r ApiUpdatePolicyRequest) Execute() (*PolicyRecord, *http.Response, error) {
=======
func (r ApiUpdatePolicyRequest) Execute() (PolicyRecord, *_nethttp.Response, error) {
>>>>>>> main
	return r.ApiService.UpdatePolicyExecute(r)
}

/*
UpdatePolicy Update policy

Update/replace and existing policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId
 @return ApiUpdatePolicyRequest
*/
func (a *PoliciesApiService) UpdatePolicy(ctx context.Context, policyId string) ApiUpdatePolicyRequest {
	return ApiUpdatePolicyRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return PolicyRecord
<<<<<<< HEAD
func (a *PoliciesApiService) UpdatePolicyExecute(r ApiUpdatePolicyRequest) (*PolicyRecord, *http.Response, error) {
=======
func (a *PoliciesApiService) UpdatePolicyExecute(r ApiUpdatePolicyRequest) (PolicyRecord, *_nethttp.Response, error) {
>>>>>>> main
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
<<<<<<< HEAD
		formFiles            []formFile
		localVarReturnValue  *PolicyRecord
=======
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PolicyRecord
>>>>>>> main
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesApiService.UpdatePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{policy_id}"
<<<<<<< HEAD
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
=======
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", _neturl.PathEscape(parameterToString(r.policyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
>>>>>>> main
	if r.policy == nil {
		return localVarReturnValue, nil, reportError("policy is required and must be specified")
	}

	if r.active != nil {
		localVarQueryParams.Add("active", parameterToString(*r.active, ""))
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
	localVarPostBody = r.policy
<<<<<<< HEAD
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
=======
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
>>>>>>> main
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
