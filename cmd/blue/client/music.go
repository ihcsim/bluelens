package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// CreateMusicPath computes a request path to the create action of music.
func CreateMusicPath() string {
	return fmt.Sprintf("/bluelens/music")
}

// CreateMusic makes a request to the create action endpoint of the music resource
func (c *Client) CreateMusic(ctx context.Context, path string, payload *Music) (*http.Response, error) {
	req, err := c.NewCreateMusicRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateMusicRequest create the request corresponding to the create action endpoint of the music resource.
func (c *Client) NewCreateMusicRequest(ctx context.Context, path string, payload *Music) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	if c.APIKeySigner != nil {
		c.APIKeySigner.Sign(req)
	}
	return req, nil
}

// ListMusicPath computes a request path to the list action of music.
func ListMusicPath() string {
	return fmt.Sprintf("/bluelens/music")
}

// List up to N music resources. N can be adjusted using the 'limit' and 'offset' parameters.
func (c *Client) ListMusic(ctx context.Context, path string, limit *int, offset *int) (*http.Response, error) {
	req, err := c.NewListMusicRequest(ctx, path, limit, offset)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListMusicRequest create the request corresponding to the list action endpoint of the music resource.
func (c *Client) NewListMusicRequest(ctx context.Context, path string, limit *int, offset *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if limit != nil {
		tmp10 := strconv.Itoa(*limit)
		values.Set("limit", tmp10)
	}
	if offset != nil {
		tmp11 := strconv.Itoa(*offset)
		values.Set("offset", tmp11)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.APIKeySigner != nil {
		c.APIKeySigner.Sign(req)
	}
	return req, nil
}

// ShowMusicPath computes a request path to the show action of music.
func ShowMusicPath(id string) string {
	return fmt.Sprintf("/bluelens/music/%v", id)
}

// Get a music resource with the given ID
func (c *Client) ShowMusic(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowMusicRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowMusicRequest create the request corresponding to the show action endpoint of the music resource.
func (c *Client) NewShowMusicRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.APIKeySigner != nil {
		c.APIKeySigner.Sign(req)
	}
	return req, nil
}
