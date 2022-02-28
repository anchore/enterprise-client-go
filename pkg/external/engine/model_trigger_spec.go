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
// TriggerSpec Definition of a trigger and its parameters
type TriggerSpec struct {
	// Name of the trigger as it would appear in a policy document
	Name string `json:"name,omitempty"`
	// Trigger description for what it tests and when it will fire during evaluation
	Description string `json:"description,omitempty"`
	// State of the trigger
	State string `json:"state,omitempty"`
	// The name of another trigger that supercedes this on functionally if this is deprecated
	SupercededBy *string `json:"superceded_by,omitempty"`
	// The list of parameters that are valid for this trigger
	Parameters []TriggerParamSpec `json:"parameters,omitempty"`
}
