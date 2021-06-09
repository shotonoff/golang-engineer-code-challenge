package network

import (
	"fmt"
	"net/http"

	"github.com/shotonoff/golang-engineer-code-challenge/internal/app/httpclient"
	"github.com/shotonoff/golang-engineer-code-challenge/internal/app/metric"
)

const (
	// P2PNetwork is used for p2p network
	P2PNetwork = "p2p"
	// HostedNetwork is used for hosted network
	HostedNetwork = "hosted"
)

var httpHeaders = map[string]string{
	"User-Agent":   "dash/client",
	"Content-Type": "application/json",
}

// NewHTTPClient returns a new network http client
func NewHTTPClient(metrics metric.Persister, network string) (*http.Client, error) {
	var costFn metric.ComputeCostFunc
	switch network {
	case P2PNetwork:
		costFn = metric.ComputeP2PRequestCost
	case HostedNetwork:
		costFn = metric.ComputeHostedRequestCost
	default:
		return nil, fmt.Errorf("given unknown network %q", network)
	}
	client := httpclient.New(
		httpclient.WithHeaders(httpHeaders),
		metric.WithMetricsMiddleware(costFn, metrics, network),
	)
	return client, nil
}
