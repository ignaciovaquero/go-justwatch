package justwatch

import (
	"fmt"
	"io"
	"net/http"
)

const (
	userAgent   string = "JustWatch Golang client (github.com/igvaquero18/go-justwatch)"
	contentType string = "application/json"

	// URL is the default API URL
	URL string = "https://apis.justwatch.com/content"
)

// Client is the client for JustWatch API
type Client struct {
	*Logger
	URL string
}

// ClientOptionFunc is a function that configures a JustWatch client.
// It is used in NewClient and in Option methods.
type ClientOptionFunc func(c *Client) error

// Option sets the options specified.
// It returns an option to restore the last arg's previous value.
func (c *Client) Option(opts ...ClientOptionFunc) error {
	for _, opt := range opts {
		if err := opt(j); err != nil {
			return err
		}
	}
	return nil
}

func Logger(logger *Logger) ClientOptionFunc {
	return func(c *Client) error {
		c.Logger = logger
		return nil
	}
}

func NewClient(opts ...ClientOptionFunc) (*Client, error) {
	c := &Client{

	}
}

func (c *Client) doReq(method, endpoint string, body io.Reader) (*http.Response, error) {
	c.Debugw(
		"executing request to the JustWatch api",
		"url", c.URL,
		"method", method,
		"endpoint", endpoint,
	)
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s/?access_token=%s", c.URL, endpoint, c.Token),
		body,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("User-Agent", userAgent)
	response, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
