package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// FollowsUserPath computes a request path to the follows action of user.
func FollowsUserPath(followerID int, followeeID string) string {
	return fmt.Sprintf("/bluelens/user/%v/follows/%v", followerID, followeeID)
}

// A user follows another user.
func (c *Client) FollowsUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFollowsUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFollowsUserRequest create the request corresponding to the follows action endpoint of the user resource.
func (c *Client) NewFollowsUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListensUserPath computes a request path to the listens action of user.
func ListensUserPath(userID string, musicID int) string {
	return fmt.Sprintf("/bluelens/user/%v/listen/%v", userID, musicID)
}

// A user listens to a music.
func (c *Client) ListensUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListensUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListensUserRequest create the request corresponding to the listens action endpoint of the user resource.
func (c *Client) NewListensUserRequest(ctx context.Context, path string) (*http.Request, error) {
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
