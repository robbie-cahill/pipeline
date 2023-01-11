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

type AmazonAutoScalingGroupSize struct {

	Min int32 `json:"min"`

	Max int32 `json:"max"`
}

// AssertAmazonAutoScalingGroupSizeRequired checks if the required fields are not zero-ed
func AssertAmazonAutoScalingGroupSizeRequired(obj AmazonAutoScalingGroupSize) error {
	elements := map[string]interface{}{
		"min": obj.Min,
		"max": obj.Max,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseAmazonAutoScalingGroupSizeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AmazonAutoScalingGroupSize (e.g. [][]AmazonAutoScalingGroupSize), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAmazonAutoScalingGroupSizeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAmazonAutoScalingGroupSize, ok := obj.(AmazonAutoScalingGroupSize)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAmazonAutoScalingGroupSizeRequired(aAmazonAutoScalingGroupSize)
	})
}
