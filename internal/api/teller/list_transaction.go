package teller

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

type listTransactionResponse struct{}

func listTransactionServer(options ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(listTransactionEndpoint(), decodeJSONRequest, encodeJSONResponse, options...)
}

func listTransactionEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return listTransactionResponse{}, nil
	}
}
