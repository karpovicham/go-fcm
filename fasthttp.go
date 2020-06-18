package fcm

import (
	"context"

	"github.com/valyala/fasthttp"
)

// FastHTTPDoer defines methods used to perform http requests using fasthttp library.
type FastHTTPDoer interface {
	Do(ctx context.Context, req *fasthttp.Request, resp *fasthttp.Response) error
}

var _ FastHTTPDoer = (*FastHTTPAdapter)(nil)

type FastHTTPAdapter struct {
	Client *fasthttp.Client
}

func NewFastHTTPAdapter(c *fasthttp.Client) *FastHTTPAdapter {
	return &FastHTTPAdapter{c}
}

func (a *FastHTTPAdapter) Do(ctx context.Context, req *fasthttp.Request, resp *fasthttp.Response) error {
	if deadline, ok := ctx.Deadline(); ok {
		return a.Client.DoDeadline(req, resp, deadline)
	}

	return a.Client.Do(req, resp)
}

var DefaultHTTPAdapter = NewFastHTTPAdapter(&fasthttp.Client{})

var (
	contentTypeHeader   = []byte("Content-Type")
	contentTypeHeaderV  = []byte("application/json")
	authorizationHeader = []byte("Authorization")
)
