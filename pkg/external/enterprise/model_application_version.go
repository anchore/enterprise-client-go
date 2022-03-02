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
// ApplicationVersion A representation of an SLDC application
type ApplicationVersion struct {
	// The id of the application version
	ApplicationVersionId string `json:"application_version_id,omitempty"`
	// The id of the application
	ApplicationId string `json:"application_id,omitempty"`
	// The name of the application
	VersionName string `json:"version_name"`
	// RFC 3339 formatted UTC timestamp when the application was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// RFC 3339 formatted UTC timestamp when the application was last updated
	LastUpdated time.Time `json:"last_updated,omitempty"`
}