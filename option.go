package fcm

import (
	"context"
	"fmt"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func applyOptions(c *SimpleClient, opts ...Option) error {
	for _, o := range opts {
		if err := o(c); err != nil {
			return fmt.Errorf("failed to apply option: %w", err)
		}
	}

	return nil
}

// Option configurates SimpleClient with defined option.
type Option func(*SimpleClient) error

// WithEndpoint returns Option to configure FCM Endpoint.
func WithEndpoint(endpoint string) Option {
	return func(c *SimpleClient) error {
		urlCfg, err := parseEndpoint(endpoint)
		if err != nil {
			return err
		}

		c.url = urlCfg
		return nil
	}
}

// WithHTTPClient returns Option to configure HTTP Client.
func WithHTTPClient(httpClient FastHTTPDoer) Option {
	return func(c *SimpleClient) error {
		c.client = httpClient
		return nil
	}
}

func WithCredentialsData(bb []byte) Option {
	return func(c *SimpleClient) error {
		if c.url.Endpoint == "" {
			return fmt.Errorf("endpoint is not set")
		}

		// loading service account json to grab the project id
		creds, err := google.CredentialsFromJSON(context.TODO(), bb)
		if err != nil {
			return fmt.Errorf("failed to load credentials from json: %w", err)
		}

		// this approach is used in google packages, so just reusing the logic
		audience := c.url.Endpoint
		tokenSource, err := google.JWTAccessTokenSourceFromJSON(bb, audience)
		if err != nil {
			return fmt.Errorf("failed to create token source from json: %w", err)
		}

		path := fmt.Sprintf("/v1/projects/%s/messages:send", creds.ProjectID)
		c.sendPath = []byte(path)
		c.tokenSource = tokenSource

		return nil
	}
}

type urlConfig struct {
	Endpoint string
	Scheme   []byte
	Host     []byte
}

func parseEndpoint(endpoint string) (urlConfig, error) {
	parsedUrl, err := url.Parse(endpoint)
	if err != nil {
		return urlConfig{}, fmt.Errorf("%q: invalid endpoint: %w", endpoint, err)
	}

	return urlConfig{
		Endpoint: endpoint,
		Scheme:   []byte(parsedUrl.Scheme),
		Host:     []byte(parsedUrl.Host),
	}, nil
}

const (
	// DefaultEndpoint contains endpoint URL of FCM service.
	// this constant value are used as audience value for the auth token
	// be careful in in case of changes
	DefaultEndpoint = "https://fcm.googleapis.com/"
)

type NoopTokenSource struct{}

func (n *NoopTokenSource) Token() (*oauth2.Token, error) {
	return nil, fmt.Errorf("noop token source")
}
