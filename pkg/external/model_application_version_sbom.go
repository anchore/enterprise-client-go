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
import (
	"time"
)
// ApplicationVersionSbom A combined sbom for the artifacts associated with an application version
type ApplicationVersionSbom struct {
	Application Application `json:"application,omitempty"`
	ApplicationVersion ApplicationVersion `json:"application_version,omitempty"`
	// RFC 3339 formatted UTC timestamp when the application version sbom was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	SourceSboms []interface{} `json:"source_sboms,omitempty"`
	ImageSboms []interface{} `json:"image_sboms,omitempty"`
}