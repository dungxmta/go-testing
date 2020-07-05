package mock

import (
	"context"
	"fmt"

	proto "github.com/micro/examples/helloworld/proto"
	"github.com/micro/go-micro/v2/client"
)

type mockGreeterService struct {
}

func (m *mockGreeterService) Call(ctx context.Context, req *proto.Request, opts ...client.CallOption) (*proto.Response, error) {
	return &proto.Response{
		Msg: fmt.Sprintf("Hi %v, Im response to u from mock", req.Name),
	}, nil
}

func NewGreeterService() proto.HelloworldService {
	return new(mockGreeterService)
}
