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
// SourceRepositoryMetadataMetadata struct for SourceRepositoryMetadataMetadata
type SourceRepositoryMetadataMetadata struct {
	Branch string `json:"branch,omitempty"`
	Tags []string `json:"tags,omitempty"`
}