/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.3.0
 * Contact: info@banzaicloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

type ListDeploymentsResponseInner struct {
	Name      string `json:"name,omitempty"`
	Chart     string `json:"chart,omitempty"`
	Version   int32  `json:"version,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	Status    string `json:"status,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}
