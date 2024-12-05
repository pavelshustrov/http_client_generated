package client

import (
	"net/http"
	"os"
	"time"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
)

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type HTTPClient struct {
	client doer
	opts   options
}

func (hc *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	//todo metrics
	return hc.client.Do(req)
}

func New(opts ...Option) *HTTPClient {
	o := options{
		timeout: 5 * time.Second,
		appName: os.Getenv("APP_NAME"),
		retryFn: heimdall.NewNoRetrier(),
	}

	for _, opt := range opts {
		opt(&o)
	}

	return newClient(o)
}

func newClient(o options) *HTTPClient {
	// todo wrap with metrics
	httpClient := httpclient.NewClient(
		httpclient.WithHTTPTimeout(o.timeout),
		httpclient.WithRetrier(o.retryFn),
	)

	return &HTTPClient{
		client: httpClient,
		opts:   o,
	}
}
