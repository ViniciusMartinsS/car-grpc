package repository

import (
	"github.com/google/uuid"
)

type Content struct {
	UUID     string
	Brand    string
	Model    string
	FuelType string
	Year     int32
}

type CarRepository struct{}

func NewCar() CarRepository {
	return CarRepository{}
}

func (c CarRepository) Save(content Content) string {
	content.UUID = uuid.NewString()
	return content.UUID
}
