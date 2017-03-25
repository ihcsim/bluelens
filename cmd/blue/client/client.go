package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the bluelens service client.
type Client struct {
	*goaclient.Client
	APIKeySigner    goaclient.Signer
	BasicAuthSigner goaclient.Signer
	Encoder         *goa.HTTPEncoder
	Decoder         *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// SetAPIKeySigner sets the request signer for the APIKey security scheme.
func (c *Client) SetAPIKeySigner(signer goaclient.Signer) {
	c.APIKeySigner = signer
}

// SetBasicAuthSigner sets the request signer for the BasicAuth security scheme.
func (c *Client) SetBasicAuthSigner(signer goaclient.Signer) {
	c.BasicAuthSigner = signer
}
