/*
 * Anchore Enterprise RBAC API
 *
 * Enterprise API for managing roles, permissions, and user mappings
 *
 * API version: 0.0.1
 * Contact: dev@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rbac
// ApiErrorResponse Generic HTTP API error response
type ApiErrorResponse struct {
	Code int32 `json:"code,omitempty"`
	ErrorType string `json:"error_type,omitempty"`
	Message string `json:"message,omitempty"`
	// Details structure for additional information about the error if available. Content and structure will be error specific.
	Detail interface{} `json:"detail,omitempty"`
}
