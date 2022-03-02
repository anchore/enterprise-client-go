/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.20
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external
// PackageReference A record of a software item which is vulnerable or carries a fix for a vulnerability
type PackageReference struct {
	// Package name
	Name string `json:"name,omitempty"`
	// A version for the package. If null, then references all versions
	Version *string `json:"version,omitempty"`
	// Package type (e.g. package, rpm, deb, apk, jar, npm, gem, ...)
	Type string `json:"type,omitempty"`
	// Whether a vendor will or will not fix a vulnerabitlity
	WillNotFix bool `json:"will_not_fix,omitempty"`
}