package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/muh-hidayaat/go-grpc-ecommerce-be/internal/handler"
	"github.com/muh-hidayaat/go-grpc-ecommerce-be/pb/service"
	"github.com/muh-hidayaat/go-grpc-ecommerce-be/pkg/database"
	"github.com/muh-hidayaat/go-grpc-ecommerce-be/pkg/grpcmiddleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	godotenv.Load()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("Connected to database")

	serviceHandler := handler.NewServiceHandler()

	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware,
		),
	)

	service.RegisterHelloWorldServer(serv, serviceHandler)

	if os.Getenv("ENVIRONMENT") == "dev" {
		reflection.Register(serv)
		log.Println("Reflection service registered")
	}

	log.Println("server is running on :50051 port.")
	if err := serv.Serve(lis); err != nil {
		log.Panicf("server is error: %v", err)
	}
}
