package handler

import (
	"context"
	"fmt"

	"github.com/muh-hidayaat/go-grpc-ecommerce-be/internal/utils"
	"github.com/muh-hidayaat/go-grpc-ecommerce-be/pb/service"
)

type serviceHandler struct {
	service.UnimplementedHelloWorldServer
}

func (sh *serviceHandler) HelloWorld(ctx context.Context, request *service.HelloRequest) (*service.HelloResponse, error) {
	return &service.HelloResponse{
		Message: fmt.Sprintf("Hello %v", request.Name),
		Base:    utils.SuccessResponse("success"),
	}, nil
}

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
