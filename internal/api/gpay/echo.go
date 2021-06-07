package gpay

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

type echoRequest struct{}
type echoResponse struct{}

func echoServer(options ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(echoEndpoint(), decodeEncryptedRequest(func() interface{} {
		return echoRequest{}
	}), encodeResponse(), options...)
}

func echoEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return echoResponse{}, nil
	}
}
