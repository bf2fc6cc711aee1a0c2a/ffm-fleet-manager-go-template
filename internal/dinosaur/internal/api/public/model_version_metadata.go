/*
 * Dinosaur Service Fleet Manager
 *
 * Dinosaur Service Fleet Manager is a Rest API to manage Dinosaur instances.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// VersionMetadata struct for VersionMetadata
type VersionMetadata struct {
	Id          string            `json:"id,omitempty"`
	Kind        string            `json:"kind,omitempty"`
	Href        string            `json:"href,omitempty"`
	Collections []ObjectReference `json:"collections,omitempty"`
}
