/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.7.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type QueryApi interface {

	/*
	QueryImagesByPackage List of images containing given package

	Filterable query interface to search for images containing specified package

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiQueryImagesByPackageRequest
	*/
	QueryImagesByPackage(ctx _context.Context) ApiQueryImagesByPackageRequest

	// QueryImagesByPackageExecute executes the request
	//  @return PaginatedImageList
	QueryImagesByPackageExecute(r ApiQueryImagesByPackageRequest) (PaginatedImageList, *_nethttp.Response, error)

	/*
	QueryImagesByVulnerability List images vulnerable to the specific vulnerability ID.

	Returns a listing of images and their respective packages vulnerable to the given vulnerability ID

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiQueryImagesByVulnerabilityRequest
	*/
	QueryImagesByVulnerability(ctx _context.Context) ApiQueryImagesByVulnerabilityRequest

	// QueryImagesByVulnerabilityExecute executes the request
	//  @return PaginatedVulnerableImageList
	QueryImagesByVulnerabilityExecute(r ApiQueryImagesByVulnerabilityRequest) (PaginatedVulnerableImageList, *_nethttp.Response, error)

	/*
	QueryVulnerabilities Listing information about given vulnerability

	List (w/filters) vulnerability records known by the system, with affected packages information if present

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiQueryVulnerabilitiesRequest
	*/
	QueryVulnerabilities(ctx _context.Context) ApiQueryVulnerabilitiesRequest

	// QueryVulnerabilitiesExecute executes the request
	//  @return PaginatedVulnerabilityList
	QueryVulnerabilitiesExecute(r ApiQueryVulnerabilitiesRequest) (PaginatedVulnerabilityList, *_nethttp.Response, error)
}

// QueryApiService QueryApi service
type QueryApiService service

type ApiQueryImagesByPackageRequest struct {
	ctx _context.Context
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

func (r ApiQueryImagesByPackageRequest) Execute() (PaginatedImageList, *_nethttp.Response, error) {
	return r.ApiService.QueryImagesByPackageExecute(r)
}

/*
QueryImagesByPackage List of images containing given package

Filterable query interface to search for images containing specified package

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryImagesByPackageRequest
*/
func (a *QueryApiService) QueryImagesByPackage(ctx _context.Context) ApiQueryImagesByPackageRequest {
	return ApiQueryImagesByPackageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedImageList
func (a *QueryApiService) QueryImagesByPackageExecute(r ApiQueryImagesByPackageRequest) (PaginatedImageList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedImageList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryApiService.QueryImagesByPackage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query/images/by_package"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}

	localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	if r.packageType != nil {
		localVarQueryParams.Add("package_type", parameterToString(*r.packageType, ""))
	}
	if r.version != nil {
		localVarQueryParams.Add("version", parameterToString(*r.version, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiQueryImagesByVulnerabilityRequest struct {
	ctx _context.Context
	ApiService QueryApi
	vulnerabilityId *string
	namespace *string
	affectedPackage *string
	severity *string
	vendorOnly *bool
	page *int32
	limit *int32
	xAnchoreAccount *string
}

// The ID of the vulnerability to search for within all images stored in anchore-engine (e.g. CVE-1999-0001)
func (r ApiQueryImagesByVulnerabilityRequest) VulnerabilityId(vulnerabilityId string) ApiQueryImagesByVulnerabilityRequest {
	r.vulnerabilityId = &vulnerabilityId
	return r
}
// Filter results to images within the given vulnerability namespace (e.g. debian:8, ubuntu:14.04)
func (r ApiQueryImagesByVulnerabilityRequest) Namespace(namespace string) ApiQueryImagesByVulnerabilityRequest {
	r.namespace = &namespace
	return r
}
// Filter results to images with vulnable packages with the given package name (e.g. libssl)
func (r ApiQueryImagesByVulnerabilityRequest) AffectedPackage(affectedPackage string) ApiQueryImagesByVulnerabilityRequest {
	r.affectedPackage = &affectedPackage
	return r
}
// Filter results to vulnerable package/vulnerability with the given severity
func (r ApiQueryImagesByVulnerabilityRequest) Severity(severity string) ApiQueryImagesByVulnerabilityRequest {
	r.severity = &severity
	return r
}
// Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data
func (r ApiQueryImagesByVulnerabilityRequest) VendorOnly(vendorOnly bool) ApiQueryImagesByVulnerabilityRequest {
	r.vendorOnly = &vendorOnly
	return r
}
// The page of results to fetch. Pages start at 1
func (r ApiQueryImagesByVulnerabilityRequest) Page(page int32) ApiQueryImagesByVulnerabilityRequest {
	r.page = &page
	return r
}
// Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page
func (r ApiQueryImagesByVulnerabilityRequest) Limit(limit int32) ApiQueryImagesByVulnerabilityRequest {
	r.limit = &limit
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiQueryImagesByVulnerabilityRequest) XAnchoreAccount(xAnchoreAccount string) ApiQueryImagesByVulnerabilityRequest {
	r.xAnchoreAccount = &xAnchoreAccount
	return r
}

func (r ApiQueryImagesByVulnerabilityRequest) Execute() (PaginatedVulnerableImageList, *_nethttp.Response, error) {
	return r.ApiService.QueryImagesByVulnerabilityExecute(r)
}

/*
QueryImagesByVulnerability List images vulnerable to the specific vulnerability ID.

Returns a listing of images and their respective packages vulnerable to the given vulnerability ID

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryImagesByVulnerabilityRequest
*/
func (a *QueryApiService) QueryImagesByVulnerability(ctx _context.Context) ApiQueryImagesByVulnerabilityRequest {
	return ApiQueryImagesByVulnerabilityRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedVulnerableImageList
func (a *QueryApiService) QueryImagesByVulnerabilityExecute(r ApiQueryImagesByVulnerabilityRequest) (PaginatedVulnerableImageList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedVulnerableImageList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryApiService.QueryImagesByVulnerability")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query/images/by_vulnerability"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.vulnerabilityId == nil {
		return localVarReturnValue, nil, reportError("vulnerabilityId is required and must be specified")
	}

	localVarQueryParams.Add("vulnerability_id", parameterToString(*r.vulnerabilityId, ""))
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.affectedPackage != nil {
		localVarQueryParams.Add("affected_package", parameterToString(*r.affectedPackage, ""))
	}
	if r.severity != nil {
		localVarQueryParams.Add("severity", parameterToString(*r.severity, ""))
	}
	if r.vendorOnly != nil {
		localVarQueryParams.Add("vendor_only", parameterToString(*r.vendorOnly, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiQueryVulnerabilitiesRequest struct {
	ctx _context.Context
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

func (r ApiQueryVulnerabilitiesRequest) Execute() (PaginatedVulnerabilityList, *_nethttp.Response, error) {
	return r.ApiService.QueryVulnerabilitiesExecute(r)
}

/*
QueryVulnerabilities Listing information about given vulnerability

List (w/filters) vulnerability records known by the system, with affected packages information if present

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryVulnerabilitiesRequest
*/
func (a *QueryApiService) QueryVulnerabilities(ctx _context.Context) ApiQueryVulnerabilitiesRequest {
	return ApiQueryVulnerabilitiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedVulnerabilityList
func (a *QueryApiService) QueryVulnerabilitiesExecute(r ApiQueryVulnerabilitiesRequest) (PaginatedVulnerabilityList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedVulnerabilityList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryApiService.QueryVulnerabilities")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query/vulnerabilities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.id == nil {
		return localVarReturnValue, nil, reportError("id is required and must be specified")
	}

	localVarQueryParams.Add("id", parameterToString(*r.id, "csv"))
	if r.affectedPackage != nil {
		localVarQueryParams.Add("affected_package", parameterToString(*r.affectedPackage, ""))
	}
	if r.affectedPackageVersion != nil {
		localVarQueryParams.Add("affected_package_version", parameterToString(*r.affectedPackageVersion, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, "csv"))
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
		if localVarHTTPResponse.StatusCode == 400 {
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
