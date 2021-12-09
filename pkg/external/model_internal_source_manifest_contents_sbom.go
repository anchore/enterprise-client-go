/*
 * Anchore Enterprise API Server
 *
 * This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.
 *
 * API version: 0.2.1
 * Contact: dev@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external
// InternalSourceManifestContentsSbom Sbom content reference
type InternalSourceManifestContentsSbom struct {
	ContentType string `json:"content_type,omitempty"`
	Digest string `json:"digest,omitempty"`
	Bucket string `json:"bucket,omitempty"`
	Key string `json:"key,omitempty"`
}
