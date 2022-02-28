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
// PolicyEvalNotificationPayloadAllOf struct for PolicyEvalNotificationPayloadAllOf
type PolicyEvalNotificationPayloadAllOf struct {
	// The Current Policy Evaluation result
	CurrEval interface{} `json:"curr_eval,omitempty"`
	// The Previous Policy Evaluation result
	LastEval interface{} `json:"last_eval,omitempty"`
	// List of Corresponding Image Annotations
	Annotations *interface{} `json:"annotations,omitempty"`
}
