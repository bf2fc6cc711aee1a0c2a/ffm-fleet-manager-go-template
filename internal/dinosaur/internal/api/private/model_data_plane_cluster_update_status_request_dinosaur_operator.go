/*
 * Dinosaur Service Fleet Manager
 *
 * Dinosaur Service Fleet Manager APIs that are used by internal services e.g fleetshard operators.
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// DataPlaneClusterUpdateStatusRequestDinosaurOperator struct for DataPlaneClusterUpdateStatusRequestDinosaurOperator
type DataPlaneClusterUpdateStatusRequestDinosaurOperator struct {
	Ready            bool     `json:"ready"`
	Version          string   `json:"version"`
	DinosaurVersions []string `json:"dinosaurVersions,omitempty"`
}
