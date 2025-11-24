package grpcmiddleware

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			err = status.Error(codes.Internal, "Internal Server Error")
		}
	}()

	res, err := handler(ctx, req)
	if err != nil {
		log.Printf("Error when listening: %v", err)
	}
	return res, err
}
