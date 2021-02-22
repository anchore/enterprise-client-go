/*
 * Anchore Enterprise API Server
 *
 * This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.
 *
 * API version: 0.1.1
 * Contact: dev@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external
// ActionPlanResolution defines the trigger IDs and content of a resolution for an action plan
type ActionPlanResolution struct {
	TriggerIds []string `json:"trigger_ids,omitempty"`
	Content string `json:"content,omitempty"`
}
