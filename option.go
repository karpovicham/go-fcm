package fcm

import (
	"errors"
)

// Option configurates SimpleClient with defined option.
type Option func(*SimpleClient) error

// WithEndpoint returns Option to configure FCM Endpoint.
func WithEndpoint(endpoint string) Option {
	return func(c *SimpleClient) error {
		if endpoint == "" {
			return errors.New("invalid endpoint")
		}
		c.endpoint = []byte(endpoint)
		return nil
	}
}

// WithHTTPClient returns Option to configure HTTP Client.
func WithHTTPClient(httpClient FastHttpClient) Option {
	return func(c *SimpleClient) error {
		c.client = httpClient
		return nil
	}
}
