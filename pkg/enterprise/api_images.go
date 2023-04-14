/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.7.0
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

type ImagesApi interface {

	/*
	GetImageAncestors Return the list of ancestor images for the given image

	Returns list of ancestor images, which are the images that form the base layers of the image

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param imageDigest
	 @return ApiGetImageAncestorsRequest
	*/
	GetImageAncestors(ctx _context.Context, imageDigest string) ApiGetImageAncestorsRequest

	// GetImageAncestorsExecute executes the request
	//  @return []ImageAncestor
	GetImageAncestorsExecute(r ApiGetImageAncestorsRequest) ([]ImageAncestor, *_nethttp.Response, error)

	/*
	GetImagePolicyCheckByDigest Check policy evaluation status for image

	Get the policy evaluation for the given image

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param imageDigest
	 @return ApiGetImagePolicyCheckByDigestRequest
	*/
	GetImagePolicyCheckByDigest(ctx _context.Context, imageDigest string) ApiGetImagePolicyCheckByDigestRequest

	// GetImagePolicyCheckByDigestExecute executes the request
	//  @return []interface{}
	GetImagePolicyCheckByDigestExecute(r ApiGetImagePolicyCheckByDigestRequest) ([]interface{}, *_nethttp.Response, error)

	/*
	GetImageVulnerabilitiesByDigest Get vulnerabilities by type

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param imageDigest
	 @param vulnType
	 @return ApiGetImageVulnerabilitiesByDigestRequest
	*/
	GetImageVulnerabilitiesByDigest(ctx _context.Context, imageDigest string, vulnType string) ApiGetImageVulnerabilitiesByDigestRequest

	// GetImageVulnerabilitiesByDigestExecute executes the request
	//  @return EnterpriseVulnerabilityResponse
	GetImageVulnerabilitiesByDigestExecute(r ApiGetImageVulnerabilitiesByDigestRequest) (EnterpriseVulnerabilityResponse, *_nethttp.Response, error)
}

// ImagesApiService ImagesApi service
type ImagesApiService service

type ApiGetImageAncestorsRequest struct {
	ctx _context.Context
	ApiService ImagesApi
	imageDigest string
	xAnchoreAccount *string
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetImageAncestorsRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetImageAncestorsRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetImageAncestorsRequest) Execute() ([]ImageAncestor, *_nethttp.Response, error) {
	return r.ApiService.GetImageAncestorsExecute(r)
}

/*
GetImageAncestors Return the list of ancestor images for the given image

Returns list of ancestor images, which are the images that form the base layers of the image

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageDigest
 @return ApiGetImageAncestorsRequest
*/
func (a *ImagesApiService) GetImageAncestors(ctx _context.Context, imageDigest string) ApiGetImageAncestorsRequest {
	return ApiGetImageAncestorsRequest{
		ApiService: a,
		ctx: ctx,
		imageDigest: imageDigest,
	}
}

// Execute executes the request
//  @return []ImageAncestor
func (a *ImagesApiService) GetImageAncestorsExecute(r ApiGetImageAncestorsRequest) ([]ImageAncestor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ImageAncestor
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesApiService.GetImageAncestors")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/images/{image_digest}/ancestors"
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
	if r.xAnchoreAccount != nil {
		localVarHeaderParams["x-anchore-account"] = parameterToString(*r.xAnchoreAccount, "")
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

type ApiGetImagePolicyCheckByDigestRequest struct {
	ctx _context.Context
	ApiService ImagesApi
	imageDigest string
	tag *string
	policyId *string
	detail *bool
	history *bool
	interactive *bool
	baseDigest *string
	xAnchoreAccount *string
}

func (r ApiGetImagePolicyCheckByDigestRequest) Tag(tag string) ApiGetImagePolicyCheckByDigestRequest {
	r.tag = &tag
	return r
}
func (r ApiGetImagePolicyCheckByDigestRequest) PolicyId(policyId string) ApiGetImagePolicyCheckByDigestRequest {
	r.policyId = &policyId
	return r
}
func (r ApiGetImagePolicyCheckByDigestRequest) Detail(detail bool) ApiGetImagePolicyCheckByDigestRequest {
	r.detail = &detail
	return r
}
func (r ApiGetImagePolicyCheckByDigestRequest) History(history bool) ApiGetImagePolicyCheckByDigestRequest {
	r.history = &history
	return r
}
func (r ApiGetImagePolicyCheckByDigestRequest) Interactive(interactive bool) ApiGetImagePolicyCheckByDigestRequest {
	r.interactive = &interactive
	return r
}
// Digest of a base image. If specified the evaluation will indicate results inherited from the base image
func (r ApiGetImagePolicyCheckByDigestRequest) BaseDigest(baseDigest string) ApiGetImagePolicyCheckByDigestRequest {
	r.baseDigest = &baseDigest
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetImagePolicyCheckByDigestRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetImagePolicyCheckByDigestRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetImagePolicyCheckByDigestRequest) Execute() ([]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetImagePolicyCheckByDigestExecute(r)
}

/*
GetImagePolicyCheckByDigest Check policy evaluation status for image

Get the policy evaluation for the given image

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageDigest
 @return ApiGetImagePolicyCheckByDigestRequest
*/
func (a *ImagesApiService) GetImagePolicyCheckByDigest(ctx _context.Context, imageDigest string) ApiGetImagePolicyCheckByDigestRequest {
	return ApiGetImagePolicyCheckByDigestRequest{
		ApiService: a,
		ctx: ctx,
		imageDigest: imageDigest,
	}
}

// Execute executes the request
//  @return []interface{}
func (a *ImagesApiService) GetImagePolicyCheckByDigestExecute(r ApiGetImagePolicyCheckByDigestRequest) ([]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesApiService.GetImagePolicyCheckByDigest")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/images/{image_digest}/check"
	localVarPath = strings.Replace(localVarPath, "{"+"image_digest"+"}", _neturl.PathEscape(parameterToString(r.imageDigest, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.tag == nil {
		return localVarReturnValue, nil, reportError("tag is required and must be specified")
	}

	if r.policyId != nil {
		localVarQueryParams.Add("policy_id", parameterToString(*r.policyId, ""))
	}
	localVarQueryParams.Add("tag", parameterToString(*r.tag, ""))
	if r.detail != nil {
		localVarQueryParams.Add("detail", parameterToString(*r.detail, ""))
	}
	if r.history != nil {
		localVarQueryParams.Add("history", parameterToString(*r.history, ""))
	}
	if r.interactive != nil {
		localVarQueryParams.Add("interactive", parameterToString(*r.interactive, ""))
	}
	if r.baseDigest != nil {
		localVarQueryParams.Add("base_digest", parameterToString(*r.baseDigest, ""))
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

type ApiGetImageVulnerabilitiesByDigestRequest struct {
	ctx _context.Context
	ApiService ImagesApi
	imageDigest string
	vulnType string
	forceRefresh *bool
	vendorOnly *bool
	baseDigest *string
	xAnchoreAccount *string
}

func (r ApiGetImageVulnerabilitiesByDigestRequest) ForceRefresh(forceRefresh bool) ApiGetImageVulnerabilitiesByDigestRequest {
	r.forceRefresh = &forceRefresh
	return r
}
// Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data. When set to true, it will filter out all vulnerabilities where &#x60;will_not_fix&#x60; is False. If false all vulnerabilities are returned regardless of &#x60;will_not_fix&#x60;
func (r ApiGetImageVulnerabilitiesByDigestRequest) VendorOnly(vendorOnly bool) ApiGetImageVulnerabilitiesByDigestRequest {
	r.vendorOnly = &vendorOnly
	return r
}
// Digest of a base image. If specified the vulnerabilities will indicate inheritance from the base image
func (r ApiGetImageVulnerabilitiesByDigestRequest) BaseDigest(baseDigest string) ApiGetImageVulnerabilitiesByDigestRequest {
	r.baseDigest = &baseDigest
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiGetImageVulnerabilitiesByDigestRequest) XAnchoreAccount(xAnchoreAccount string) ApiGetImageVulnerabilitiesByDigestRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiGetImageVulnerabilitiesByDigestRequest) Execute() (EnterpriseVulnerabilityResponse, *_nethttp.Response, error) {
	return r.ApiService.GetImageVulnerabilitiesByDigestExecute(r)
}

/*
GetImageVulnerabilitiesByDigest Get vulnerabilities by type

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageDigest
 @param vulnType
 @return ApiGetImageVulnerabilitiesByDigestRequest
*/
func (a *ImagesApiService) GetImageVulnerabilitiesByDigest(ctx _context.Context, imageDigest string, vulnType string) ApiGetImageVulnerabilitiesByDigestRequest {
	return ApiGetImageVulnerabilitiesByDigestRequest{
		ApiService: a,
		ctx: ctx,
		imageDigest: imageDigest,
		vulnType: vulnType,
	}
}

// Execute executes the request
//  @return EnterpriseVulnerabilityResponse
func (a *ImagesApiService) GetImageVulnerabilitiesByDigestExecute(r ApiGetImageVulnerabilitiesByDigestRequest) (EnterpriseVulnerabilityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EnterpriseVulnerabilityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesApiService.GetImageVulnerabilitiesByDigest")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/images/{image_digest}/vuln/{vuln_type}"
	localVarPath = strings.Replace(localVarPath, "{"+"image_digest"+"}", _neturl.PathEscape(parameterToString(r.imageDigest, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vuln_type"+"}", _neturl.PathEscape(parameterToString(r.vulnType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.forceRefresh != nil {
		localVarQueryParams.Add("force_refresh", parameterToString(*r.forceRefresh, ""))
	}
	if r.vendorOnly != nil {
		localVarQueryParams.Add("vendor_only", parameterToString(*r.vendorOnly, ""))
	}
	if r.baseDigest != nil {
		localVarQueryParams.Add("base_digest", parameterToString(*r.baseDigest, ""))
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
