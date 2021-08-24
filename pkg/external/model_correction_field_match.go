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
// CorrectionFieldMatch Defines a particular field name and value to match for a Correction
type CorrectionFieldMatch struct {
	// The package field name to match
	FieldName string `json:"field_name"`
	// The package field value for the corresponding field_name above to match. If field_name corresponds to a list value, this will search the list
	FieldValue string `json:"field_value"`
}
