/*
 * Anchore Enterprise API Server
 *
 * This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.
 *
 * API version: 0.2.0
 * Contact: dev@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external
import (
	"time"
)
// Correction Defines a correction object for false positive management
type Correction struct {
	// Identifier for the correction
	Uuid string `json:"uuid,omitempty"`
	// Type of correction
	Type string `json:"type"`
	Description string `json:"description,omitempty"`
	Match CorrectionMatch `json:"match"`
	Replace []CorrectionFieldMatch `json:"replace"`
	// RFC 3339 formatted UTC timestamp when the correction was generated
	CreatedAt time.Time `json:"created_at,omitempty"`
	// RFC 3339 formatted UTC timestamp when the correction was last modified
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
