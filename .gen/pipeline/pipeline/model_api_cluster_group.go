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

type ApiClusterGroup struct {

	EnabledFeatures []string `json:"enabledFeatures,omitempty"`

	Id int32 `json:"id,omitempty"`

	Members []ApiMember `json:"members,omitempty"`

	Name string `json:"name,omitempty"`

	OrganizationId int32 `json:"organizationId,omitempty"`

	Uid string `json:"uid,omitempty"`
}

// AssertApiClusterGroupRequired checks if the required fields are not zero-ed
func AssertApiClusterGroupRequired(obj ApiClusterGroup) error {
	for _, el := range obj.Members {
		if err := AssertApiMemberRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseApiClusterGroupRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiClusterGroup (e.g. [][]ApiClusterGroup), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiClusterGroupRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiClusterGroup, ok := obj.(ApiClusterGroup)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiClusterGroupRequired(aApiClusterGroup)
	})
}
