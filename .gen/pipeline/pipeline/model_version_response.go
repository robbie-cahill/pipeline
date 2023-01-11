/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pipeline

// VersionResponse - Pipeline build and deployment info
type VersionResponse struct {

	Version string `json:"version,omitempty"`

	CommitHash string `json:"commit_hash,omitempty"`

	BuildDate string `json:"build_date,omitempty"`

	GoVersion string `json:"go_version,omitempty"`

	Os string `json:"os,omitempty"`

	Arch string `json:"arch,omitempty"`

	Compiler string `json:"compiler,omitempty"`

	InstanceUuid string `json:"instance_uuid,omitempty"`
}

// AssertVersionResponseRequired checks if the required fields are not zero-ed
func AssertVersionResponseRequired(obj VersionResponse) error {
	return nil
}

// AssertRecurseVersionResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of VersionResponse (e.g. [][]VersionResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseVersionResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aVersionResponse, ok := obj.(VersionResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertVersionResponseRequired(aVersionResponse)
	})
}
