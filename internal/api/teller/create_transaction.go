package teller

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

type createTransactionResponse struct{}

func createTransactionServer(options ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(createTransactionEndpoint(), decodeJSONRequest, encodeJSONResponse, options...)
}

func createTransactionEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return createTransactionResponse{}, nil
	}
}
