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
// AssociatedImageArtifact Model for an associated image artifact. Composites the artifact with the association metadata
type AssociatedImageArtifact struct {
	ArtifactAssociationMetadata ArtifactAssociationMetadata `json:"artifact_association_metadata,omitempty"`
	Image ImageArtifact `json:"image,omitempty"`
}