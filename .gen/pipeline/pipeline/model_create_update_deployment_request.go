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

type CreateUpdateDeploymentRequest struct {

	Name string `json:"name"`

	// Version of the deployment. If not specified, the latest version is used.
	Version string `json:"version,omitempty"`

	// The chart content packaged by `helm package`. If specified chart version is ignored.
	Package string `json:"package,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	ReleaseName string `json:"releaseName,omitempty"`

	DryRun bool `json:"dryRun,omitempty"`

	// if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment are in a ready state before marking the release as successful
	Wait bool `json:"wait,omitempty"`

	// time in seconds to wait for any individual Kubernetes operation (default 300)
	Timeout int64 `json:"timeout,omitempty"`

	// Map of resources in the template where replicas should have a minimum on-demand percentage. Format: <kind.resourceName: min-percentage>
	OdPcts map[string]interface{} `json:"odPcts,omitempty"`

	ReuseValues bool `json:"reuseValues,omitempty"`

	// current values of the deployment
	Values map[string]interface{} `json:"values,omitempty"`
}

// AssertCreateUpdateDeploymentRequestRequired checks if the required fields are not zero-ed
func AssertCreateUpdateDeploymentRequestRequired(obj CreateUpdateDeploymentRequest) error {
	elements := map[string]interface{}{
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCreateUpdateDeploymentRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CreateUpdateDeploymentRequest (e.g. [][]CreateUpdateDeploymentRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCreateUpdateDeploymentRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCreateUpdateDeploymentRequest, ok := obj.(CreateUpdateDeploymentRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCreateUpdateDeploymentRequestRequired(aCreateUpdateDeploymentRequest)
	})
}
