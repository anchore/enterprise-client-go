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
// AnalysisArchiveAddResult The result of adding a single digest to the archive
type AnalysisArchiveAddResult struct {
	// The image digest requested to be added
	Digest string `json:"digest,omitempty"`
	// The status of the archive add operation. Typically either 'archived' or 'error'
	Status string `json:"status,omitempty"`
	// Details on the status, e.g. the error message
	Detail string `json:"detail,omitempty"`
}
