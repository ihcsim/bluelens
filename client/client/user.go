package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// FollowUserPath computes a request path to the follow action of user.
func FollowUserPath(followerID int, followeeID string) string {
	return fmt.Sprintf("/bluelens/user/%v/follows/%v", followerID, followeeID)
}

// Add a user to another user's followees list.
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

// ListenUserPath computes a request path to the listen action of user.
func ListenUserPath(userID string, musicID int) string {
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
