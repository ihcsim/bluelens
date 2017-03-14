package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// GetMusicPath computes a request path to the get action of music.
func GetMusicPath(musicID string) string {
	return fmt.Sprintf("/bluelens/music/%v", musicID)
}

// Get a music resource with the given ID
func (c *Client) GetMusic(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetMusicRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetMusicRequest create the request corresponding to the get action endpoint of the music resource.
func (c *Client) NewGetMusicRequest(ctx context.Context, path string) (*http.Request, error) {
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
