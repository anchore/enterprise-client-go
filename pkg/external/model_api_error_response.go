/*
 * Anchore Enterprise API Server
 *
 * This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.
 *
 * API version: 0.2.0
 * Contact: dev@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external
// ApiErrorResponse Generic HTTP API error response
type ApiErrorResponse struct {
	Code int32 `json:"code,omitempty"`
	ErrorType string `json:"error_type,omitempty"`
	Message string `json:"message,omitempty"`
	// Details structure for additional information about the error if available. Content and structure will be error specific.
	Detail interface{} `json:"detail,omitempty"`
}
