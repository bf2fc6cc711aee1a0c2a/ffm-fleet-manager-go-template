/*
 * Dinosaur Service Fleet Manager
 *
 * Dinosaur Service Fleet Manager APIs that are used by internal services e.g fleetshard operators.
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// ManagedDinosaurVersions struct for ManagedDinosaurVersions
type ManagedDinosaurVersions struct {
	Dinosaur         string `json:"dinosaur,omitempty"`
	DinosaurOperator string `json:"dinosaurOperator,omitempty"`
}
