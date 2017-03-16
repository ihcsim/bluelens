package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// FollowUserPath computes a request path to the follow action of user.
func FollowUserPath(userID string, followeeID string) string {
	return fmt.Sprintf("/bluelens/user/%v/follows/%v", userID, followeeID)
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

// GetUserPath computes a request path to the get action of user.
func GetUserPath(userID string) string {
	return fmt.Sprintf("/bluelens/user/%v", userID)
}

// Get a user resource with the given ID
func (c *Client) GetUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetUserRequest create the request corresponding to the get action endpoint of the user resource.
func (c *Client) NewGetUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListenUserPath computes a request path to the listen action of user.
func ListenUserPath(userID string, musicID string) string {
	return fmt.Sprintf("/bluelens/user/%v/listen/%v", userID, musicID)
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
