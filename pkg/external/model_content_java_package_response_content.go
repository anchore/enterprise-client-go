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
// ContentJavaPackageResponseContent struct for ContentJavaPackageResponseContent
type ContentJavaPackageResponseContent struct {
	Package string `json:"package,omitempty"`
	ImplementationVersion string `json:"implementation-version,omitempty"`
	SpecificationVersion string `json:"specification-version,omitempty"`
	MavenVersion string `json:"maven-version,omitempty"`
	Location string `json:"location,omitempty"`
	Type string `json:"type,omitempty"`
	Origin string `json:"origin,omitempty"`
	// A list of Common Platform Enumerations that may uniquely identify the package
	Cpes []string `json:"cpes,omitempty"`
}
