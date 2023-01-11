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

type CreatePkeClusterKubernetesOidc struct {

	Enabled bool `json:"enabled,omitempty"`
}

// AssertCreatePkeClusterKubernetesOidcRequired checks if the required fields are not zero-ed
func AssertCreatePkeClusterKubernetesOidcRequired(obj CreatePkeClusterKubernetesOidc) error {
	return nil
}

// AssertRecurseCreatePkeClusterKubernetesOidcRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CreatePkeClusterKubernetesOidc (e.g. [][]CreatePkeClusterKubernetesOidc), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCreatePkeClusterKubernetesOidcRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCreatePkeClusterKubernetesOidc, ok := obj.(CreatePkeClusterKubernetesOidc)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCreatePkeClusterKubernetesOidcRequired(aCreatePkeClusterKubernetesOidc)
	})
}
