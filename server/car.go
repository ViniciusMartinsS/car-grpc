package server

import (
	"context"
	"errors"
	pb "github.com/ViniciusMartinss/car-grpc/proto"
	"github.com/ViniciusMartinss/car-grpc/repository"
	"time"

	"google.golang.org/grpc"
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

func (s *Server) Create(ctx context.Context, in *pb.CarCreateRequest) (*pb.CarCreateResponse, error) {
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

	uuid, err := s.r.Save(
		ctx,
		repository.Content{
			Brand:    in.Brand,
			Model:    in.Model,
			FuelType: in.FuelType,
			Year:     in.Year,
		},
	)
	if err != nil {
		return nil, errors.New("err to save on the db: " + err.Error())
	}

	return &pb.CarCreateResponse{
		Uuid: uuid,
	}, nil
}

func (s *Server) List(ctx context.Context, _ *pb.CarListRequest) (*pb.CarListResponse, error) {
	cars, err := s.r.List(ctx)
	if err != nil {
		return nil, errors.New("err to fetch from the db: " + err.Error())
	}

	rCars := make([]*pb.CarListResponseContent, 0)
	for _, c := range cars {
		rCars = append(
			rCars,
			&pb.CarListResponseContent{
				Uuid:     c.UUID,
				Brand:    c.Brand,
				Model:    c.Model,
				FuelType: c.FuelType,
				Year:     c.Year,
			},
		)
	}

	return &pb.CarListResponse{
		Status: true,
		Cars:   rCars,
	}, nil
}
