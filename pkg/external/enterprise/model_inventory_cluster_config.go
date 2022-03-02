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
// InventoryClusterConfig Cluster specific configuration
type InventoryClusterConfig struct {
	CredentialType string `json:"credential_type,omitempty"`
	Namespaces []string `json:"namespaces,omitempty"`
	// FQDN of the cluster API server
	ClusterServer string `json:"cluster_server,omitempty"`
	// Base64 Encoded Public Certificate for the cluster
	ClusterCertificate string `json:"cluster_certificate,omitempty"`
	// Base64 Encoded Public Certificate for the client. Not required if credential_type == token
	ClientCertificate string `json:"client_certificate,omitempty"`
	// Base64 Encoded credential for the client
	Credential string `json:"credential,omitempty"`
}