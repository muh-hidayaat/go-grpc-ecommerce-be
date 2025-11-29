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
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}
	if validationErrors != nil {
		return &service.HelloResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	return &service.HelloResponse{
		Message: fmt.Sprintf("Hello %v", request.Name),
		Base:    utils.SuccessResponse("success"),
	}, nil
}

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
