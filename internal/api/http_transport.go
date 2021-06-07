package api

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/generic"
	"github.com/go-kit/kit/tracing/opencensus"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/sbutakov/layout-internal-api/internal/api/gpay"
	"github.com/sbutakov/layout-internal-api/internal/api/teller"
)

type Option func(*configuration)

type configuration struct {
	latency metrics.Histogram
}

func Routes(logger log.Logger, options ...Option) http.Handler {
	config := &configuration{latency: generic.NewHistogram("nop", 1)}
	for _, option := range options {
		option(config)
	}

	serverOptions := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(level.Error(logger))),
		opencensus.HTTPServerTrace(),
	}

	r := mux.NewRouter()
	gpay.RegisterRoutes(r, config.latency, serverOptions...)
	teller.RegisterRoutes(r, config.latency, serverOptions...)
	return r
}

func WithHistogramMetric(histogram metrics.Histogram) Option {
	return func(c *configuration) {
		c.latency = histogram
	}
}
