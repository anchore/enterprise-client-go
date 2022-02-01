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
// AssociatedSourceArtifact Model for an associated source artifact. Composite of the source artifact and its asssociation metadata
type AssociatedSourceArtifact struct {
	ArtifactAssociationMetadata ArtifactAssociationMetadata `json:"artifact_association_metadata,omitempty"`
	Source Source `json:"source,omitempty"`
}
