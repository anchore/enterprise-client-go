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
// ArtifactListResponse The response provided when querying for the artifacts on an application version
type ArtifactListResponse struct {
	AssociatedSourceArtifacts []AssociatedSourceArtifact `json:"associated_source_artifacts,omitempty"`
	AssociatedImageArtifacts []AssociatedImageArtifact `json:"associated_image_artifacts,omitempty"`
}