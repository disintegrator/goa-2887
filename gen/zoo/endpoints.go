// Code generated by goa v3.4.3, DO NOT EDIT.
//
// zoo endpoints
//
// Command:
// $ goa gen bugrepro/design

package zoo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "zoo" service endpoints.
type Endpoints struct {
	PetAnimal goa.Endpoint
}

// NewEndpoints wraps the methods of the "zoo" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		PetAnimal: NewPetAnimalEndpoint(s),
	}
}

// Use applies the given middleware to all the "zoo" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.PetAnimal = m(e.PetAnimal)
}

// NewPetAnimalEndpoint returns an endpoint function that calls the method
// "petAnimal" of service "zoo".
func NewPetAnimalEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PetAnimalPayload)
		return s.PetAnimal(ctx, p)
	}
}
