package server

import (
	"context"
	"errors"
	pb "github.com/ViniciusMartinss/car-grpc/proto"
	"github.com/ViniciusMartinss/car-grpc/repository"
	"google.golang.org/grpc"
	"time"
)

type Server struct {
	pb.UnimplementedCarServiceServer

	gs *grpc.Server
	r  repository.CarRepository
}

func NewCar(gs *grpc.Server, r repository.CarRepository) *Server {
	return &Server{
		gs: gs,
		r:  r,
	}
}

func (s *Server) Register() {
	pb.RegisterCarServiceServer(s.gs, s)
}

func (s *Server) Create(ctx context.Context, in *pb.CarRequest) (*pb.CarResponse, error) {
	if in == nil {
		return nil, errors.New("missing body")
	}
	if in.Brand == "" {
		return nil, errors.New("missing brand")
	}
	if in.Model == "" {
		return nil, errors.New("missing model")
	}
	if in.FuelType == "" {
		return nil, errors.New("missing fueltype")
	}
	if in.Year == 0 || (in.Year < 1886 || in.Year > int32(time.Now().Year()+1)) {
		return nil, errors.New("invalid year")
	}

	return &pb.CarResponse{
		Uuid: s.r.Save(
			repository.Content{
				Brand:    in.Brand,
				Model:    in.Model,
				FuelType: in.FuelType,
				Year:     in.Year,
			},
		),
	}, nil
}
