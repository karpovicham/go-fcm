package fcm

import (
	"context"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

var (
	// DefaultEndpoint contains endpoint URL of FCM service.
	DefaultEndpoint = []byte("https://fcm.googleapis.com/fcm/send")

	// ErrInvalidAPIKey occurs if API key is not set.
	ErrInvalidAPIKey = errors.New("client API Key is invalid")

	contentTypeHeader   = []byte("Content-Type")
	contentTypeHeaderV  = []byte("application/json")
	authorizationHeader = []byte("Authorization")
)

// SimpleClient abstracts the interaction between the application server and the
// FCM server via HTTP protocol. The developer must obtain an API key from the
// Google APIs Console page and pass it to the `SimpleClient` so that it can
// perform authorized requests on the application server's behalf.
// To send a message to one or more devices use the SimpleClient's Send.
//
// If the `HTTP` field is nil, a zeroed http.SimpleClient will be allocated and used
// to send messages.
type Client interface {
	// Send sends a message to the FCM server without retrying in case of service
	// unavailability. A non-nil error is returned if a non-recoverable error
	// occurs (i.e. if the response status is not "200 OK").
	Send(ctx context.Context, msg *Message) (*Response, error)
}

type FastHttpClient interface {
	Do(req *fasthttp.Request, resp *fasthttp.Response) error
}

var _ FastHttpClient = &fasthttp.Client{}

type SimpleClient struct {
	apiKey   []byte
	client   FastHttpClient
	endpoint []byte
}

var _ Client = &SimpleClient{}

// NewClient creates new Firebase Cloud Messaging SimpleClient based on API key and
// with default endpoint and http client.
func NewClient(apiKey string, opts ...Option) (*SimpleClient, error) {
	if apiKey == "" {
		return nil, ErrInvalidAPIKey
	}
	c := &SimpleClient{
		apiKey:   []byte("key=" + apiKey),
		endpoint: DefaultEndpoint,
		client:   &fasthttp.Client{},
	}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

var (
	postBytes = []byte("POST")
)

func (c *SimpleClient) Send(ctx context.Context, msg *Message) (*Response, error) {
	if err := msg.Validate(); err != nil {
		return nil, err
	}

	data, err := msg.MarshalJSON()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal msg %#v", *msg)
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethodBytes(postBytes)
	req.Header.SetRequestURIBytes(c.endpoint)
	req.Header.AddBytesKV(contentTypeHeader, contentTypeHeaderV)
	req.Header.AddBytesKV(authorizationHeader, c.apiKey)
	req.SetBody(data)

	if err := c.client.Do(req, resp); err != nil {
		return nil, errors.Wrapf(err, "failed to do http request")
	}

	response := new(Response)
	response.StatusCode = resp.StatusCode()
	respBody := resp.Body()
	if err := response.UnmarshalJSON(respBody); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal response %s", string(respBody))
	}

	return response, nil
}
