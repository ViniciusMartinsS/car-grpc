package main

import (
	"github.com/ViniciusMartinss/car-grpc/repository"
	"github.com/ViniciusMartinss/car-grpc/server"
	"log"
	"net"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	r := repository.NewCar(redisClient)
	s := grpc.NewServer()
	server.NewCar(s, r).Register()

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
