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

type GetPkeCommandsResponse struct {

	AdditionalPropertiesField string `json:"additionalProperties,omitempty"`
}

// AssertGetPkeCommandsResponseRequired checks if the required fields are not zero-ed
func AssertGetPkeCommandsResponseRequired(obj GetPkeCommandsResponse) error {
	return nil
}

// AssertRecurseGetPkeCommandsResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetPkeCommandsResponse (e.g. [][]GetPkeCommandsResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetPkeCommandsResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetPkeCommandsResponse, ok := obj.(GetPkeCommandsResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetPkeCommandsResponseRequired(aGetPkeCommandsResponse)
	})
}
