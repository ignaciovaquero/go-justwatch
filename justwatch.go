package justwatch

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	userAgent   string = "JustWatch Golang client (github.com/igvaquero18/go-justwatch)"
	contentType string = "application/json"

	justWatchURL   string = "https://apis.justwatch.com/content"
	defaultCountry string = "ES"
)

// Client is the client for JustWatch API
type Client struct {
	Logger
	*http.Client
	URL     string
	Country string
	locale  string
}

// ClientOptionFunc is a function that configures a JustWatch client.
// It is used in NewClient and in Option methods.
type ClientOptionFunc func(c *Client) error

// Option sets the options specified.
// It returns an option to restore the last arg's previous value.
func (c *Client) Option(opts ...ClientOptionFunc) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	return nil
}

// SetLogger sets a logger for the client
func SetLogger(logger Logger) ClientOptionFunc {
	return func(c *Client) error {
		c.Logger = logger
		return nil
	}
}

// SetURL overrides the default JustWatch API URL
func SetURL(url string) ClientOptionFunc {
	return func(c *Client) error {
		c.URL = url
		return nil
	}
}

// NewClient creates a new JustWatch client
func NewClient(opts ...ClientOptionFunc) (*Client, error) {
	c := &Client{
		Logger: &defaultLogger{},
		URL:    justWatchURL,
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	const defaultLocale = "es_ES"
	resp, err := c.doReq(http.MethodGet, "locales/state", nil)
	if err != nil {
		return nil, err
	}
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return c, nil
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
		fmt.Sprintf("%s/%s", c.URL, endpoint),
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
