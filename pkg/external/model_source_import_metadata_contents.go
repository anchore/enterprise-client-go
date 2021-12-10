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
// SourceImportMetadataContents Digest of content to use in the final import
type SourceImportMetadataContents struct {
	// Digest to use for the sbom
	Sbom string `json:"sbom"`
}
