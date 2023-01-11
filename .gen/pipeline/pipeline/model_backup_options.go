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

type BackupOptions struct {

	IncludedNamespaces []string `json:"includedNamespaces,omitempty"`

	IncludedResources []string `json:"includedResources,omitempty"`

	ExcludedNamespaces []string `json:"excludedNamespaces,omitempty"`

	ExcludedResources []string `json:"excludedResources,omitempty"`

	SnapshotVolumes bool `json:"snapshotVolumes,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`
}

// AssertBackupOptionsRequired checks if the required fields are not zero-ed
func AssertBackupOptionsRequired(obj BackupOptions) error {
	return nil
}

// AssertRecurseBackupOptionsRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of BackupOptions (e.g. [][]BackupOptions), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseBackupOptionsRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aBackupOptions, ok := obj.(BackupOptions)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertBackupOptionsRequired(aBackupOptions)
	})
}
