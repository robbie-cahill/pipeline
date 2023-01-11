//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package projectdriver

import (
	"context"
	"errors"
	"github.com/banzaicloud/pipeline/internal/app/pipeline/cloud/google/project"
	"github.com/go-kit/kit/endpoint"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
)

// endpointError identifies an error that should be returned as an endpoint error.
type endpointError interface {
	EndpointError() bool
}

// serviceError identifies an error that should be returned as a service error.
type serviceError interface {
	ServiceError() bool
}

// Endpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	ListProjects endpoint.Endpoint
}

// MakeEndpoints returns a(n) Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeEndpoints(service project.Service, middleware ...endpoint.Middleware) Endpoints {
	mw := kitxendpoint.Combine(middleware...)

	return Endpoints{ListProjects: kitxendpoint.OperationNameMiddleware("cloud/google/project.ListProjects")(mw(MakeListProjectsEndpoint(service)))}
}

// ListProjectsRequest is a request struct for ListProjects endpoint.
type ListProjectsRequest struct {
	SecretID string
}

// ListProjectsResponse is a response struct for ListProjects endpoint.
type ListProjectsResponse struct {
	Projects []project.Project
	Err      error
}

func (r ListProjectsResponse) Failed() error {
	return r.Err
}

// MakeListProjectsEndpoint returns an endpoint for the matching method of the underlying service.
func MakeListProjectsEndpoint(service project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListProjectsRequest)

		projects, err := service.ListProjects(ctx, req.SecretID)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return ListProjectsResponse{
					Err:      err,
					Projects: projects,
				}, nil
			}

			return ListProjectsResponse{
				Err:      err,
				Projects: projects,
			}, err
		}

		return ListProjectsResponse{Projects: projects}, nil
	}
}
