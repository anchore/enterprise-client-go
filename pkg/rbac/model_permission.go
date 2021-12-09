/*
 * Anchore Enterprise RBAC API
 *
 * Enterprise API for managing roles, permissions, and user mappings
 *
 * API version: 0.0.1
 * Contact: dev@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rbac
// Permission A grant of specific action against a specific scope and target
type Permission struct {
	// The allowed action. e.g. getImage
	Action string `json:"action,omitempty"`
	// The target to which the action may be applied. Either a '*' for all or a specific target id
	Target string `json:"target,omitempty"`
}
