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
// VulnDiffResult The results of the comparing two vulnerability records during an update
type VulnDiffResult struct {
	Added []interface{} `json:"added,omitempty"`
	Updated []interface{} `json:"updated,omitempty"`
	Removed []interface{} `json:"removed,omitempty"`
}
