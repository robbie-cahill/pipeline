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

type UpdateNodePoolsAzure struct {

	Autoscaling bool `json:"autoscaling,omitempty"`

	Count int32 `json:"count"`

	MinCount int32 `json:"minCount,omitempty"`

	MaxCount int32 `json:"maxCount,omitempty"`

	InstanceType string `json:"instanceType,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`

	VnetSubnetID string `json:"vnetSubnetID,omitempty"`
}

// AssertUpdateNodePoolsAzureRequired checks if the required fields are not zero-ed
func AssertUpdateNodePoolsAzureRequired(obj UpdateNodePoolsAzure) error {
	elements := map[string]interface{}{
		"count": obj.Count,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUpdateNodePoolsAzureRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UpdateNodePoolsAzure (e.g. [][]UpdateNodePoolsAzure), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUpdateNodePoolsAzureRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUpdateNodePoolsAzure, ok := obj.(UpdateNodePoolsAzure)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUpdateNodePoolsAzureRequired(aUpdateNodePoolsAzure)
	})
}
