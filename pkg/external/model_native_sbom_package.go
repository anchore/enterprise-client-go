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
// NativeSbomPackage struct for NativeSbomPackage
type NativeSbomPackage struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	Version string `json:"version"`
	Type string `json:"type"`
	FoundBy string `json:"foundBy,omitempty"`
	Locations []NativeSbomPackageLocation `json:"locations"`
	Licenses []string `json:"licenses"`
	Language string `json:"language"`
	Cpes []string `json:"cpes"`
	Purl string `json:"purl,omitempty"`
	MetadataType *string `json:"metadataType,omitempty"`
	Metadata *interface{} `json:"metadata,omitempty"`
}
