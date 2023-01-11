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

type SecretKeyValueAmazon struct {

	AWS_ACCESS_KEY_ID string `json:"AWS_ACCESS_KEY_ID"`

	AWS_SECRET_ACCESS_KEY string `json:"AWS_SECRET_ACCESS_KEY"`
}

// AssertSecretKeyValueAmazonRequired checks if the required fields are not zero-ed
func AssertSecretKeyValueAmazonRequired(obj SecretKeyValueAmazon) error {
	elements := map[string]interface{}{
		"AWS_ACCESS_KEY_ID": obj.AWS_ACCESS_KEY_ID,
		"AWS_SECRET_ACCESS_KEY": obj.AWS_SECRET_ACCESS_KEY,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseSecretKeyValueAmazonRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SecretKeyValueAmazon (e.g. [][]SecretKeyValueAmazon), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSecretKeyValueAmazonRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSecretKeyValueAmazon, ok := obj.(SecretKeyValueAmazon)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSecretKeyValueAmazonRequired(aSecretKeyValueAmazon)
	})
}
