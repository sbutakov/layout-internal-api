package teller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/metrics"
	"github.com/gorilla/mux"

	kithttp "github.com/go-kit/kit/transport/http"
)

func RegisterRoutes(r *mux.Router, _ metrics.Histogram, options ...kithttp.ServerOption) {
	options = append(options, kithttp.ServerErrorEncoder(errorResponseEncode))
	r.Methods(http.MethodPost).Path("/v1/transactions").Handler(createTransactionServer(options...))
	r.Methods(http.MethodGet).Path("/v1/transactions").Handler(listTransactionServer(options...))
}

func decodeJSONRequest(context.Context, *http.Request) (request interface{}, err error) {
	return nil, nil
}

func encodeJSONResponse(_ context.Context, w http.ResponseWriter, i interface{}) error {
	_, _ = w.Write([]byte(fmt.Sprintf("%T", i)))
	return nil

}

func errorResponseEncode(context.Context, error, http.ResponseWriter) {}
