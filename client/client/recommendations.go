package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// ListRecommendationsPath computes a request path to the list action of recommendations.
func ListRecommendationsPath(userID string, maxCount int) string {
	return fmt.Sprintf("/bluelens/recommendations/%v/%v", userID, maxCount)
}

// List all the music recommendations for a user.
func (c *Client) ListRecommendations(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListRecommendationsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListRecommendationsRequest create the request corresponding to the list action endpoint of the recommendations resource.
func (c *Client) NewListRecommendationsRequest(ctx context.Context, path string) (*http.Request, error) {
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
