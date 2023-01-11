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

type CreateClusterResponse202 struct {

	Name string `json:"name,omitempty"`

	Id int32 `json:"id,omitempty"`
}

// AssertCreateClusterResponse202Required checks if the required fields are not zero-ed
func AssertCreateClusterResponse202Required(obj CreateClusterResponse202) error {
	return nil
}

// AssertRecurseCreateClusterResponse202Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CreateClusterResponse202 (e.g. [][]CreateClusterResponse202), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCreateClusterResponse202Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCreateClusterResponse202, ok := obj.(CreateClusterResponse202)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCreateClusterResponse202Required(aCreateClusterResponse202)
	})
}
