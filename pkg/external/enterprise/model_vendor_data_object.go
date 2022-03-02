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
// VendorDataObject struct for VendorDataObject
type VendorDataObject struct {
	// Vendor Vulnerability ID
	Id string `json:"id,omitempty"`
	CvssV2 Cvssv2Scores `json:"cvss_v2,omitempty"`
	CvssV3 Cvssv3Scores `json:"cvss_v3,omitempty"`
}