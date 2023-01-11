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

// NodePoolStatusInformation - Current status and its reason of the node pool.
type NodePoolStatusInformation struct {

	// Current status of the node pool.
	Status string `json:"status,omitempty"`

	// Details and reasoning about the status value.
	StatusMessage string `json:"statusMessage,omitempty"`
}

// AssertNodePoolStatusInformationRequired checks if the required fields are not zero-ed
func AssertNodePoolStatusInformationRequired(obj NodePoolStatusInformation) error {
	return nil
}

// AssertRecurseNodePoolStatusInformationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NodePoolStatusInformation (e.g. [][]NodePoolStatusInformation), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNodePoolStatusInformationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNodePoolStatusInformation, ok := obj.(NodePoolStatusInformation)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNodePoolStatusInformationRequired(aNodePoolStatusInformation)
	})
}
