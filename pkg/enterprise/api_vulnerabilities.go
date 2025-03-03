/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type VulnerabilitiesAPI interface {

	/*
	VulnerabilityScanSbom Return a vulnerability scan for the uploaded SBOM without storing the SBOM and without any side-effects in the system.

	Use this operation for checking sboms for vulnerabilities in cases where the sbom does not need to be stored for later re-scans or added to the managed set of SBOMs in Anchore. If you need to upload and save an SBOM use the "/import/*" API set instead.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiVulnerabilityScanSbomRequest
	*/
	VulnerabilityScanSbom(ctx context.Context) ApiVulnerabilityScanSbomRequest

	// VulnerabilityScanSbomExecute executes the request
	//  @return SBOMVulnerabilitiesResponse
	VulnerabilityScanSbomExecute(r ApiVulnerabilityScanSbomRequest) (*SBOMVulnerabilitiesResponse, *http.Response, error)
}

// VulnerabilitiesAPIService VulnerabilitiesAPI service
type VulnerabilitiesAPIService service

type ApiVulnerabilityScanSbomRequest struct {
	ctx context.Context
	ApiService VulnerabilitiesAPI
	sbom *interface{}
	xAnchoreAccount *string
	includeVulnDescription *bool
}

func (r ApiVulnerabilityScanSbomRequest) Sbom(sbom interface{}) ApiVulnerabilityScanSbomRequest {
	r.sbom = &sbom
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiVulnerabilityScanSbomRequest) XAnchoreAccount(xAnchoreAccount string) ApiVulnerabilityScanSbomRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiVulnerabilityScanSbomRequest) IncludeVulnDescription(includeVulnDescription bool) ApiVulnerabilityScanSbomRequest {
	r.includeVulnDescription = &includeVulnDescription
	return r
}

func (r ApiVulnerabilityScanSbomRequest) Execute() (*SBOMVulnerabilitiesResponse, *http.Response, error) {
	return r.ApiService.VulnerabilityScanSbomExecute(r)
}

/*
VulnerabilityScanSbom Return a vulnerability scan for the uploaded SBOM without storing the SBOM and without any side-effects in the system.

Use this operation for checking sboms for vulnerabilities in cases where the sbom does not need to be stored for later re-scans or added to the managed set of SBOMs in Anchore. If you need to upload and save an SBOM use the "/import/*" API set instead.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVulnerabilityScanSbomRequest
*/
func (a *VulnerabilitiesAPIService) VulnerabilityScanSbom(ctx context.Context) ApiVulnerabilityScanSbomRequest {
	return ApiVulnerabilityScanSbomRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SBOMVulnerabilitiesResponse
func (a *VulnerabilitiesAPIService) VulnerabilityScanSbomExecute(r ApiVulnerabilityScanSbomRequest) (*SBOMVulnerabilitiesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SBOMVulnerabilitiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesAPIService.VulnerabilityScanSbom")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vulnerability-scan"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sbom == nil {
		return localVarReturnValue, nil, reportError("sbom is required and must be specified")
	}

	if r.includeVulnDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_vuln_description", r.includeVulnDescription, "form", "")
	} else {
		var defaultValue bool = false
		r.includeVulnDescription = &defaultValue
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-anchore-account", r.xAnchoreAccount, "simple", "")
	}
	// body params
	localVarPostBody = r.sbom
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
