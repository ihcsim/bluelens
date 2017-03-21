package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// CreateUserPath computes a request path to the create action of user.
func CreateUserPath() string {
	return fmt.Sprintf("/bluelens/user")
}

// CreateUser makes a request to the create action endpoint of the user resource
func (c *Client) CreateUser(ctx context.Context, path string, payload *User, id *string) (*http.Response, error) {
	req, err := c.NewCreateUserRequest(ctx, path, payload, id)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateUserRequest create the request corresponding to the create action endpoint of the user resource.
func (c *Client) NewCreateUserRequest(ctx context.Context, path string, payload *User, id *string) (*http.Request, error) {
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
	values := u.Query()
	if id != nil {
		values.Set("id", *id)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// FollowUserPath computes a request path to the follow action of user.
func FollowUserPath(id string, followeeID string) string {
	return fmt.Sprintf("/bluelens/user/%v/follows/%v", id, followeeID)
}

// Update a user's followees list with a new followee.
func (c *Client) FollowUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFollowUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFollowUserRequest create the request corresponding to the follow action endpoint of the user resource.
func (c *Client) NewFollowUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListUserPath computes a request path to the list action of user.
func ListUserPath() string {
	return fmt.Sprintf("/bluelens/user")
}

// List up to N user resources. N can be adjusted using the 'limit' and 'offset' parameters.
func (c *Client) ListUser(ctx context.Context, path string, id *string, limit *int, offset *int) (*http.Response, error) {
	req, err := c.NewListUserRequest(ctx, path, id, limit, offset)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListUserRequest create the request corresponding to the list action endpoint of the user resource.
func (c *Client) NewListUserRequest(ctx context.Context, path string, id *string, limit *int, offset *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if id != nil {
		values.Set("id", *id)
	}
	if limit != nil {
		tmp12 := strconv.Itoa(*limit)
		values.Set("limit", tmp12)
	}
	if offset != nil {
		tmp13 := strconv.Itoa(*offset)
		values.Set("offset", tmp13)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListenUserPath computes a request path to the listen action of user.
func ListenUserPath(id string, musicID string) string {
	return fmt.Sprintf("/bluelens/user/%v/listen/%v", id, musicID)
}

// Add a music to a user's history.
func (c *Client) ListenUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListenUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListenUserRequest create the request corresponding to the listen action endpoint of the user resource.
func (c *Client) NewListenUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowUserPath computes a request path to the show action of user.
func ShowUserPath(id string) string {
	return fmt.Sprintf("/bluelens/user/%v", id)
}

// Get a user resource with the given ID.
func (c *Client) ShowUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowUserRequest create the request corresponding to the show action endpoint of the user resource.
func (c *Client) NewShowUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
