package main

import (
	"github.com/ViniciusMartinss/car-grpc/repository"
	"github.com/ViniciusMartinss/car-grpc/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	r := repository.NewCar()
	s := grpc.NewServer()
	server.NewCar(s, r).Register()

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
