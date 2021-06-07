package gpay

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

type associationRequest struct{}
type associationResponse struct{}

func associationServer(options ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(associationEndpoint(), decodeEncryptedRequest(func() interface{} {
		return associationRequest{}
	}), encodeResponse(), options...)
}

func associationEndpoint() endpoint.Endpoint {
	return func(context.Context, interface{}) (interface{}, error) {
		return associationResponse{}, nil
	}
}
