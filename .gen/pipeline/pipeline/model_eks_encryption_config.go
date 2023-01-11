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

// EksEncryptionConfig - EKS encryption configuration object describing an encryption provider and the corresponding resources to encrypt. More information can be found at https://docs.aws.amazon.com/eks/latest/APIReference/API_EncryptionConfig.html.
type EksEncryptionConfig struct {

	Provider EksEncryptionConfigProvider `json:"provider,omitempty"`

	// Resource kinds to encrypt with the corresponding encryption provider.
	Resources []string `json:"resources,omitempty"`
}

// AssertEksEncryptionConfigRequired checks if the required fields are not zero-ed
func AssertEksEncryptionConfigRequired(obj EksEncryptionConfig) error {
	if err := AssertEksEncryptionConfigProviderRequired(obj.Provider); err != nil {
		return err
	}
	return nil
}

// AssertRecurseEksEncryptionConfigRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of EksEncryptionConfig (e.g. [][]EksEncryptionConfig), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseEksEncryptionConfigRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aEksEncryptionConfig, ok := obj.(EksEncryptionConfig)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertEksEncryptionConfigRequired(aEksEncryptionConfig)
	})
}
