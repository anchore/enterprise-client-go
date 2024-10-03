/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
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


type QueryApi interface {

	/*
	QueryImagesByPackage List of images containing given package

	Filterable query interface to search for images containing specified package

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQueryImagesByPackageRequest
	*/
	QueryImagesByPackage(ctx context.Context) ApiQueryImagesByPackageRequest

	// QueryImagesByPackageExecute executes the request
	//  @return PaginatedImageList
	QueryImagesByPackageExecute(r ApiQueryImagesByPackageRequest) (*PaginatedImageList, *http.Response, error)

	/*
	QueryVulnerabilities Listing information about given vulnerability

	List (w/filters) vulnerability records known by the system, with affected packages information if present

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQueryVulnerabilitiesRequest
	*/
	QueryVulnerabilities(ctx context.Context) ApiQueryVulnerabilitiesRequest

	// QueryVulnerabilitiesExecute executes the request
	//  @return PaginatedVulnerabilityList
	QueryVulnerabilitiesExecute(r ApiQueryVulnerabilitiesRequest) (*PaginatedVulnerabilityList, *http.Response, error)
}

// QueryApiService QueryApi service
type QueryApiService service

type ApiQueryImagesByPackageRequest struct {
	ctx context.Context
	ApiService QueryApi
	name *string
	packageType *string
	version *string
	page *string
	limit *int32
	xAnchoreAccount *string
}

// Name of package to search for (e.g. sed)
func (r ApiQueryImagesByPackageRequest) Name(name string) ApiQueryImagesByPackageRequest {
	r.name = &name
	return r
}

// Type of package to filter on (e.g. dpkg)
func (r ApiQueryImagesByPackageRequest) PackageType(packageType string) ApiQueryImagesByPackageRequest {
	r.packageType = &packageType
	return r
}

// Version of named package to filter on (e.g. 4.4-1)
func (r ApiQueryImagesByPackageRequest) Version(version string) ApiQueryImagesByPackageRequest {
	r.version = &version
	return r
}

// The page of results to fetch. Pages start at 1
func (r ApiQueryImagesByPackageRequest) Page(page string) ApiQueryImagesByPackageRequest {
	r.page = &page
	return r
}

// Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page
func (r ApiQueryImagesByPackageRequest) Limit(limit int32) ApiQueryImagesByPackageRequest {
	r.limit = &limit
	return r
}

// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiQueryImagesByPackageRequest) XAnchoreAccount(xAnchoreAccount string) ApiQueryImagesByPackageRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiQueryImagesByPackageRequest) Execute() (*PaginatedImageList, *http.Response, error) {
	return r.ApiService.QueryImagesByPackageExecute(r)
}

/*
QueryImagesByPackage List of images containing given package

Filterable query interface to search for images containing specified package

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryImagesByPackageRequest
*/
func (a *QueryApiService) QueryImagesByPackage(ctx context.Context) ApiQueryImagesByPackageRequest {
	return ApiQueryImagesByPackageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedImageList
func (a *QueryApiService) QueryImagesByPackageExecute(r ApiQueryImagesByPackageRequest) (*PaginatedImageList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedImageList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryApiService.QueryImagesByPackage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query/images/by-package"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	if r.packageType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "package_type", r.packageType, "form", "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-anchore-account", r.xAnchoreAccount, "simple", "")
	}
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiQueryVulnerabilitiesRequest struct {
	ctx context.Context
	ApiService QueryApi
	id *[]string
	affectedPackage *string
	affectedPackageVersion *string
	page *string
	limit *int32
	namespace *[]string
}

// The ID of the vulnerability (e.g. CVE-1999-0001)
func (r ApiQueryVulnerabilitiesRequest) Id(id []string) ApiQueryVulnerabilitiesRequest {
	r.id = &id
	return r
}

// Filter results by specified package name (e.g. sed)
func (r ApiQueryVulnerabilitiesRequest) AffectedPackage(affectedPackage string) ApiQueryVulnerabilitiesRequest {
	r.affectedPackage = &affectedPackage
	return r
}

// Filter results by specified package version (e.g. 4.4-1)
func (r ApiQueryVulnerabilitiesRequest) AffectedPackageVersion(affectedPackageVersion string) ApiQueryVulnerabilitiesRequest {
	r.affectedPackageVersion = &affectedPackageVersion
	return r
}

// The page of results to fetch. Pages start at 1
func (r ApiQueryVulnerabilitiesRequest) Page(page string) ApiQueryVulnerabilitiesRequest {
	r.page = &page
	return r
}

// Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page
func (r ApiQueryVulnerabilitiesRequest) Limit(limit int32) ApiQueryVulnerabilitiesRequest {
	r.limit = &limit
	return r
}

// Namespace(s) to filter vulnerability records by
func (r ApiQueryVulnerabilitiesRequest) Namespace(namespace []string) ApiQueryVulnerabilitiesRequest {
	r.namespace = &namespace
	return r
}

func (r ApiQueryVulnerabilitiesRequest) Execute() (*PaginatedVulnerabilityList, *http.Response, error) {
	return r.ApiService.QueryVulnerabilitiesExecute(r)
}

/*
QueryVulnerabilities Listing information about given vulnerability

List (w/filters) vulnerability records known by the system, with affected packages information if present

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryVulnerabilitiesRequest
*/
func (a *QueryApiService) QueryVulnerabilities(ctx context.Context) ApiQueryVulnerabilitiesRequest {
	return ApiQueryVulnerabilitiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedVulnerabilityList
func (a *QueryApiService) QueryVulnerabilitiesExecute(r ApiQueryVulnerabilitiesRequest) (*PaginatedVulnerabilityList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedVulnerabilityList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryApiService.QueryVulnerabilities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query/vulnerabilities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id == nil {
		return localVarReturnValue, nil, reportError("id is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "csv")
	if r.affectedPackage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "affected_package", r.affectedPackage, "form", "")
	}
	if r.affectedPackageVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "affected_package_version", r.affectedPackageVersion, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue string = "1"
		r.page = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "form", "csv")
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
		if localVarHTTPResponse.StatusCode == 400 {
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
