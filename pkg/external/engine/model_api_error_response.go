/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.20
 * Contact: nurmi@anchore.com
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
