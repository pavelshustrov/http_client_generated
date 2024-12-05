package client

import (
	"time"

	"github.com/gojek/heimdall/v7"
)

type options struct {
	appName string

	timeout time.Duration
	retryFn Retriable
}

type Retriable interface {
	NextInterval(retry int) time.Duration
}

type Option func(*options)

func WithAppName(appName string) Option {
	return func(o *options) {
		o.appName = appName
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.timeout = timeout
	}
}

func WithRetryFn(fn Retriable) Option {
	return func(o *options) {
		o.retryFn = fn
	}
}

func NewExponentialBackoff(initialTimeout, maxTimeout time.Duration, exponentFactor float64, maximumJitterInterval time.Duration) Retriable {
	return heimdall.NewRetrier(heimdall.NewExponentialBackoff(initialTimeout, maxTimeout, exponentFactor, maximumJitterInterval))
}
