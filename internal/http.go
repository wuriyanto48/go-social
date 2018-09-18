package internal

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// HTTPClient structure
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient function for intialize HTTPClient object
// Parameter, timeout in time.Duration
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}

// newRequest function for initalize http request,
// paramters, http method, uri path, body, and headers
func (c *HTTPClient) newRequest(ctx context.Context, method string, fullPath string, body io.Reader, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

// Execute function for call http request
func (c *HTTPClient) Execute(ctx context.Context, method, path string, body io.Reader, v interface{}, headers map[string]string) error {
	req, err := c.newRequest(ctx, method, path, body, headers)

	if err != nil {
		return err
	}

	res, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if v != nil {
		return json.NewDecoder(res.Body).Decode(v)
	}

	return nil
}
