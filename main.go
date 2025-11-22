package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/muh-hidayaat/go-grpc-ecommerce-be/internal/handler"
	"github.com/muh-hidayaat/go-grpc-ecommerce-be/pb/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	godotenv.Load()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	serviceHandler := handler.NewServiceHandler()

	serv := grpc.NewServer()

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
