package request

import (
	"github.com/sewakantor/sw-be/businesses/complex"
)

type AddComplex struct {
	Name       string  `json:"name" validate:"required"`
	Street     string  `json:"street" validate:"required"`
	City       string  `json:"city" validate:"required"`
	State      string  `json:"state" validate:"required"`
	Country    string  `json:"country" validate:"required"`
	PostalCode int     `json:"postal_code" validate:"required"`
	Latitude   float64 `json:"latitude" validate:"required"`
	Longitude  float64 `json:"longitude" validate:"required"`
}

func (rec *AddComplex) AddComplexToDomain() *complex.Domain {
	return &complex.Domain{
		Name:   rec.Name,
		Street: rec.Street,
		City: rec.City,
		State: rec.State,
		Country: rec.Country,
		PostalCode: rec.PostalCode,
		Latitude: rec.Latitude,
		Longitude: rec.Longitude,
	}
}
