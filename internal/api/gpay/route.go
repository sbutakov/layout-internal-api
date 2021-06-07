package gpay

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/metrics"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type gpayRequest func() interface{}

func RegisterRoutes(r *mux.Router, _ metrics.Histogram, options ...kithttp.ServerOption) {
	options = append(options, kithttp.ServerErrorEncoder(errorResponseEncode))
	r.Methods(http.MethodGet).Path("/v1/echo").Handler(echoServer(options...))
	r.Methods(http.MethodPost).Path("/v1/associateAccount").Handler(associationServer(options...))
}

func decodeEncryptedRequest(fn gpayRequest) kithttp.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		return fn(), nil
	}
}

func encodeResponse() kithttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, i interface{}) error {
		_, _ = w.Write([]byte(fmt.Sprintf("%T", i)))
		return nil
	}
}

func errorResponseEncode(context.Context, error, http.ResponseWriter) {}
