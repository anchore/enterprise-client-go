/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

<<<<<<< HEAD
API version: 2.0.0
=======
API version: 0.1.0
>>>>>>> 48fc108 (feat: updated the enterprise ref)
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type ArchivesApi interface {

	/*
	ArchiveImageAnalysis Method for ArchiveImageAnalysis

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiArchiveImageAnalysisRequest
	*/
	ArchiveImageAnalysis(ctx _context.Context) ApiArchiveImageAnalysisRequest

	// ArchiveImageAnalysisExecute executes the request
	//  @return []AnalysisArchiveAddResult
	ArchiveImageAnalysisExecute(r ApiArchiveImageAnalysisRequest) ([]AnalysisArchiveAddResult, *_nethttp.Response, error)

	/*
	CreateAnalysisArchiveRule Method for CreateAnalysisArchiveRule

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiCreateAnalysisArchiveRuleRequest
	*/
	CreateAnalysisArchiveRule(ctx _context.Context) ApiCreateAnalysisArchiveRuleRequest

	// CreateAnalysisArchiveRuleExecute executes the request
	//  @return AnalysisArchiveTransitionRule
	CreateAnalysisArchiveRuleExecute(r ApiCreateAnalysisArchiveRuleRequest) (AnalysisArchiveTransitionRule, *_nethttp.Response, error)

	/*
	DeleteAnalysisArchiveRule Method for DeleteAnalysisArchiveRule

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ruleId
	 @return ApiDeleteAnalysisArchiveRuleRequest
	*/
	DeleteAnalysisArchiveRule(ctx _context.Context, ruleId string) ApiDeleteAnalysisArchiveRuleRequest

	// DeleteAnalysisArchiveRuleExecute executes the request
	DeleteAnalysisArchiveRuleExecute(r ApiDeleteAnalysisArchiveRuleRequest) (*_nethttp.Response, error)

	/*
	DeleteArchivedAnalysis Method for DeleteArchivedAnalysis

	Performs a synchronous archive deletion

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param imageDigest
	 @return ApiDeleteArchivedAnalysisRequest
	*/
	DeleteArchivedAnalysis(ctx _context.Context, imageDigest string) ApiDeleteArchivedAnalysisRequest

	// DeleteArchivedAnalysisExecute executes the request
	DeleteArchivedAnalysisExecute(r ApiDeleteArchivedAnalysisRequest) (*_nethttp.Response, error)

	/*
	GetAnalysisArchiveRule Method for GetAnalysisArchiveRule

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ruleId
	 @return ApiGetAnalysisArchiveRuleRequest
	*/
	GetAnalysisArchiveRule(ctx _context.Context, ruleId string) ApiGetAnalysisArchiveRuleRequest

	// GetAnalysisArchiveRuleExecute executes the request
	//  @return AnalysisArchiveTransitionRule
	GetAnalysisArchiveRuleExecute(r ApiGetAnalysisArchiveRuleRequest) (AnalysisArchiveTransitionRule, *_nethttp.Response, error)

	/*
	GetArchivedAnalysis Method for GetArchivedAnalysis

	Returns the archive metadata record identifying the image and tags for the analysis in the archive.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param imageDigest The image digest to identify the image analysis
	 @return ApiGetArchivedAnalysisRequest
	*/
	GetArchivedAnalysis(ctx _context.Context, imageDigest string) ApiGetArchivedAnalysisRequest

	// GetArchivedAnalysisExecute executes the request
	//  @return ArchivedAnalysis
	GetArchivedAnalysisExecute(r ApiGetArchivedAnalysisRequest) (ArchivedAnalysis, *_nethttp.Response, error)

	/*
	ListAnalysisArchive Method for ListAnalysisArchive

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiListAnalysisArchiveRequest
	*/
	ListAnalysisArchive(ctx _context.Context) ApiListAnalysisArchiveRequest

	// ListAnalysisArchiveExecute executes the request
	//  @return []ArchivedAnalysis
	ListAnalysisArchiveExecute(r ApiListAnalysisArchiveRequest) ([]ArchivedAnalysis, *_nethttp.Response, error)

	/*
	ListAnalysisArchiveRules Method for ListAnalysisArchiveRules

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiListAnalysisArchiveRulesRequest
	*/
	ListAnalysisArchiveRules(ctx _context.Context) ApiListAnalysisArchiveRulesRequest

	// ListAnalysisArchiveRulesExecute executes the request
	//  @return []AnalysisArchiveTransitionRule
	ListAnalysisArchiveRulesExecute(r ApiListAnalysisArchiveRulesRequest) ([]AnalysisArchiveTransitionRule, *_nethttp.Response, error)

	/*
	ListArchives Method for ListArchives

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiListArchivesRequest
	*/
	ListArchives(ctx _context.Context) ApiListArchivesRequest

	// ListArchivesExecute executes the request
	//  @return ArchiveSummary
	ListArchivesExecute(r ApiListArchivesRequest) (ArchiveSummary, *_nethttp.Response, error)
}

// ArchivesApiService ArchivesApi service
type ArchivesApiService service

type ApiArchiveImageAnalysisRequest struct {
	ctx _context.Context
	ApiService ArchivesApi
	imageReferences *[]string
}

func (r ApiArchiveImageAnalysisRequest) ImageReferences(imageReferences []string) ApiArchiveImageAnalysisRequest {
	r.imageReferences = &imageReferences
	return r
}

func (r ApiArchiveImageAnalysisRequest) Execute() ([]AnalysisArchiveAddResult, *_nethttp.Response, error) {
	return r.ApiService.ArchiveImageAnalysisExecute(r)
}

/*
ArchiveImageAnalysis Method for ArchiveImageAnalysis

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiArchiveImageAnalysisRequest
*/
func (a *ArchivesApiService) ArchiveImageAnalysis(ctx _context.Context) ApiArchiveImageAnalysisRequest {
	return ApiArchiveImageAnalysisRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AnalysisArchiveAddResult
func (a *ArchivesApiService) ArchiveImageAnalysisExecute(r ApiArchiveImageAnalysisRequest) ([]AnalysisArchiveAddResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AnalysisArchiveAddResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesApiService.ArchiveImageAnalysis")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/images"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.imageReferences == nil {
		return localVarReturnValue, nil, reportError("imageReferences is required and must be specified")
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
	localVarPostBody = r.imageReferences
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateAnalysisArchiveRuleRequest struct {
	ctx _context.Context
	ApiService ArchivesApi
	rule *AnalysisArchiveTransitionRule
}

func (r ApiCreateAnalysisArchiveRuleRequest) Rule(rule AnalysisArchiveTransitionRule) ApiCreateAnalysisArchiveRuleRequest {
	r.rule = &rule
	return r
}

func (r ApiCreateAnalysisArchiveRuleRequest) Execute() (AnalysisArchiveTransitionRule, *_nethttp.Response, error) {
	return r.ApiService.CreateAnalysisArchiveRuleExecute(r)
}

/*
CreateAnalysisArchiveRule Method for CreateAnalysisArchiveRule

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateAnalysisArchiveRuleRequest
*/
func (a *ArchivesApiService) CreateAnalysisArchiveRule(ctx _context.Context) ApiCreateAnalysisArchiveRuleRequest {
	return ApiCreateAnalysisArchiveRuleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AnalysisArchiveTransitionRule
func (a *ArchivesApiService) CreateAnalysisArchiveRuleExecute(r ApiCreateAnalysisArchiveRuleRequest) (AnalysisArchiveTransitionRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AnalysisArchiveTransitionRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesApiService.CreateAnalysisArchiveRule")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.rule == nil {
		return localVarReturnValue, nil, reportError("rule is required and must be specified")
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
	localVarPostBody = r.rule
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteAnalysisArchiveRuleRequest struct {
	ctx _context.Context
	ApiService ArchivesApi
	ruleId string
}


func (r ApiDeleteAnalysisArchiveRuleRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteAnalysisArchiveRuleExecute(r)
}

/*
DeleteAnalysisArchiveRule Method for DeleteAnalysisArchiveRule

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ruleId
 @return ApiDeleteAnalysisArchiveRuleRequest
*/
func (a *ArchivesApiService) DeleteAnalysisArchiveRule(ctx _context.Context, ruleId string) ApiDeleteAnalysisArchiveRuleRequest {
	return ApiDeleteAnalysisArchiveRuleRequest{
		ApiService: a,
		ctx: ctx,
		ruleId: ruleId,
	}
}

// Execute executes the request
func (a *ArchivesApiService) DeleteAnalysisArchiveRuleExecute(r ApiDeleteAnalysisArchiveRuleRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesApiService.DeleteAnalysisArchiveRule")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/rules/{rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"rule_id"+"}", _neturl.PathEscape(parameterToString(r.ruleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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

type ApiDeleteArchivedAnalysisRequest struct {
	ctx _context.Context
	ApiService ArchivesApi
	imageDigest string
	force *bool
}

func (r ApiDeleteArchivedAnalysisRequest) Force(force bool) ApiDeleteArchivedAnalysisRequest {
	r.force = &force
	return r
}

func (r ApiDeleteArchivedAnalysisRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteArchivedAnalysisExecute(r)
}

/*
DeleteArchivedAnalysis Method for DeleteArchivedAnalysis

Performs a synchronous archive deletion

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageDigest
 @return ApiDeleteArchivedAnalysisRequest
*/
func (a *ArchivesApiService) DeleteArchivedAnalysis(ctx _context.Context, imageDigest string) ApiDeleteArchivedAnalysisRequest {
	return ApiDeleteArchivedAnalysisRequest{
		ApiService: a,
		ctx: ctx,
		imageDigest: imageDigest,
	}
}

// Execute executes the request
func (a *ArchivesApiService) DeleteArchivedAnalysisExecute(r ApiDeleteArchivedAnalysisRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesApiService.DeleteArchivedAnalysis")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/images/{image_digest}"
	localVarPath = strings.Replace(localVarPath, "{"+"image_digest"+"}", _neturl.PathEscape(parameterToString(r.imageDigest, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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

type ApiGetAnalysisArchiveRuleRequest struct {
	ctx _context.Context
	ApiService ArchivesApi
	ruleId string
}


func (r ApiGetAnalysisArchiveRuleRequest) Execute() (AnalysisArchiveTransitionRule, *_nethttp.Response, error) {
	return r.ApiService.GetAnalysisArchiveRuleExecute(r)
}

/*
GetAnalysisArchiveRule Method for GetAnalysisArchiveRule

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ruleId
 @return ApiGetAnalysisArchiveRuleRequest
*/
func (a *ArchivesApiService) GetAnalysisArchiveRule(ctx _context.Context, ruleId string) ApiGetAnalysisArchiveRuleRequest {
	return ApiGetAnalysisArchiveRuleRequest{
		ApiService: a,
		ctx: ctx,
		ruleId: ruleId,
	}
}

// Execute executes the request
//  @return AnalysisArchiveTransitionRule
func (a *ArchivesApiService) GetAnalysisArchiveRuleExecute(r ApiGetAnalysisArchiveRuleRequest) (AnalysisArchiveTransitionRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AnalysisArchiveTransitionRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesApiService.GetAnalysisArchiveRule")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/rules/{rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"rule_id"+"}", _neturl.PathEscape(parameterToString(r.ruleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetArchivedAnalysisRequest struct {
	ctx _context.Context
	ApiService ArchivesApi
	imageDigest string
}


func (r ApiGetArchivedAnalysisRequest) Execute() (ArchivedAnalysis, *_nethttp.Response, error) {
	return r.ApiService.GetArchivedAnalysisExecute(r)
}

/*
GetArchivedAnalysis Method for GetArchivedAnalysis

Returns the archive metadata record identifying the image and tags for the analysis in the archive.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageDigest The image digest to identify the image analysis
 @return ApiGetArchivedAnalysisRequest
*/
func (a *ArchivesApiService) GetArchivedAnalysis(ctx _context.Context, imageDigest string) ApiGetArchivedAnalysisRequest {
	return ApiGetArchivedAnalysisRequest{
		ApiService: a,
		ctx: ctx,
		imageDigest: imageDigest,
	}
}

// Execute executes the request
//  @return ArchivedAnalysis
func (a *ArchivesApiService) GetArchivedAnalysisExecute(r ApiGetArchivedAnalysisRequest) (ArchivedAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ArchivedAnalysis
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesApiService.GetArchivedAnalysis")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/images/{image_digest}"
	localVarPath = strings.Replace(localVarPath, "{"+"image_digest"+"}", _neturl.PathEscape(parameterToString(r.imageDigest, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAnalysisArchiveRequest struct {
	ctx _context.Context
	ApiService ArchivesApi
}


func (r ApiListAnalysisArchiveRequest) Execute() ([]ArchivedAnalysis, *_nethttp.Response, error) {
	return r.ApiService.ListAnalysisArchiveExecute(r)
}

/*
ListAnalysisArchive Method for ListAnalysisArchive

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAnalysisArchiveRequest
*/
func (a *ArchivesApiService) ListAnalysisArchive(ctx _context.Context) ApiListAnalysisArchiveRequest {
	return ApiListAnalysisArchiveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ArchivedAnalysis
func (a *ArchivesApiService) ListAnalysisArchiveExecute(r ApiListAnalysisArchiveRequest) ([]ArchivedAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ArchivedAnalysis
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesApiService.ListAnalysisArchive")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/images"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAnalysisArchiveRulesRequest struct {
	ctx _context.Context
	ApiService ArchivesApi
	systemGlobal *bool
}

// If true include system global rules (owned by admin) even for non-admin users. Defaults to true if not set. Can be set to false to exclude globals
func (r ApiListAnalysisArchiveRulesRequest) SystemGlobal(systemGlobal bool) ApiListAnalysisArchiveRulesRequest {
	r.systemGlobal = &systemGlobal
	return r
}

func (r ApiListAnalysisArchiveRulesRequest) Execute() ([]AnalysisArchiveTransitionRule, *_nethttp.Response, error) {
	return r.ApiService.ListAnalysisArchiveRulesExecute(r)
}

/*
ListAnalysisArchiveRules Method for ListAnalysisArchiveRules

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAnalysisArchiveRulesRequest
*/
func (a *ArchivesApiService) ListAnalysisArchiveRules(ctx _context.Context) ApiListAnalysisArchiveRulesRequest {
	return ApiListAnalysisArchiveRulesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AnalysisArchiveTransitionRule
func (a *ArchivesApiService) ListAnalysisArchiveRulesExecute(r ApiListAnalysisArchiveRulesRequest) ([]AnalysisArchiveTransitionRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AnalysisArchiveTransitionRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesApiService.ListAnalysisArchiveRules")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.systemGlobal != nil {
		localVarQueryParams.Add("system_global", parameterToString(*r.systemGlobal, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListArchivesRequest struct {
	ctx _context.Context
	ApiService ArchivesApi
}


func (r ApiListArchivesRequest) Execute() (ArchiveSummary, *_nethttp.Response, error) {
	return r.ApiService.ListArchivesExecute(r)
}

/*
ListArchives Method for ListArchives

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListArchivesRequest
*/
func (a *ArchivesApiService) ListArchives(ctx _context.Context) ApiListArchivesRequest {
	return ApiListArchivesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ArchiveSummary
func (a *ArchivesApiService) ListArchivesExecute(r ApiListArchivesRequest) (ArchiveSummary, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ArchiveSummary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivesApiService.ListArchives")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/archives"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
