package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// RecommendRecommendationsPath computes a request path to the recommend action of recommendations.
func RecommendRecommendationsPath(userID string, limit int) string {
	return fmt.Sprintf("/bluelens/recommendations/%v/%v", userID, limit)
}

// Make music recommendations for a user.
func (c *Client) RecommendRecommendations(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRecommendRecommendationsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRecommendRecommendationsRequest create the request corresponding to the recommend action endpoint of the recommendations resource.
func (c *Client) NewRecommendRecommendationsRequest(ctx context.Context, path string) (*http.Request, error) {
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
