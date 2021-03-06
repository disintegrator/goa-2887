// Code generated by goa v3.4.3, DO NOT EDIT.
//
// zoo client HTTP transport
//
// Command:
// $ goa gen bugrepro/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the zoo service endpoint HTTP clients.
type Client struct {
	// PetAnimal Doer is the HTTP client used to make requests to the petAnimal
	// endpoint.
	PetAnimalDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the zoo service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		PetAnimalDoer:       doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// PetAnimal returns an endpoint that makes HTTP requests to the zoo service
// petAnimal server.
func (c *Client) PetAnimal() goa.Endpoint {
	var (
		encodeRequest  = EncodePetAnimalRequest(c.encoder)
		decodeResponse = DecodePetAnimalResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPetAnimalRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PetAnimalDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("zoo", "petAnimal", err)
		}
		return decodeResponse(resp)
	}
}
