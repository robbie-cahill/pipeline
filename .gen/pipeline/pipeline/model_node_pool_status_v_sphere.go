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

type NodePoolStatusVSphere struct {

	Count int32 `json:"count,omitempty"`

	Ram int32 `json:"ram,omitempty"`

	Vcpu int32 `json:"vcpu,omitempty"`

	Template string `json:"template,omitempty"`

	ResourceSummary map[string]ResourceSummary `json:"resourceSummary,omitempty"`
}

// AssertNodePoolStatusVSphereRequired checks if the required fields are not zero-ed
func AssertNodePoolStatusVSphereRequired(obj NodePoolStatusVSphere) error {
	return nil
}

// AssertRecurseNodePoolStatusVSphereRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NodePoolStatusVSphere (e.g. [][]NodePoolStatusVSphere), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNodePoolStatusVSphereRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNodePoolStatusVSphere, ok := obj.(NodePoolStatusVSphere)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNodePoolStatusVSphereRequired(aNodePoolStatusVSphere)
	})
}
