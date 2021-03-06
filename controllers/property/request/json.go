package request

import (
	"github.com/sewakantor/sw-be/businesses/property"
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

type AddBuilding struct {
	Name         string `json:"name" validate:"required"`
	Photo        string `json:"photo" validate:"required"`
	Year         int    `json:"year" validate:"required"`
	Floor        int    `json:"floor" validate:"required"`
	FloorSurface int    `json:"floor_surface" validate:"required"`
	TotalSurface int    `json:"total_surface" validate:"required"`
	Price        int    `json:"price" validate:"required"`
	Desc         string `json:"desc" validate:"required"`
}

type AddReview struct {
	Commend string `json:"comment" validate:"required"`
	Star    int    `json:"star" validate:"required"`
}

type AddUnit struct {
	Name     string `json:"name" validate:"required"`
	Surface  int    `json:"surface" validate:"required"`
	Capacity int    `json:"capacity" validate:"required"`
}

func (rec *AddComplex) AddComplexToDomain() *property.Complex {
	return &property.Complex{
		Name:       rec.Name,
		Street:     rec.Street,
		City:       rec.City,
		State:      rec.State,
		Country:    rec.Country,
		PostalCode: rec.PostalCode,
		Latitude:   rec.Latitude,
		Longitude:  rec.Longitude,
	}
}

func (rec *AddBuilding) AddBuildingToDomain() *property.Building {
	return &property.Building{
		Name:         rec.Name,
		Photo:        rec.Photo,
		Year:         rec.Year,
		Floor:        rec.Floor,
		FloorSurface: rec.FloorSurface,
		TotalSurface: rec.TotalSurface,
		Price:        rec.Price,
		Desc:         rec.Desc,
	}
}

func (rec *AddReview) AddReviewToDomain() *property.Review {
	return &property.Review{
		Commend: rec.Commend,
		Star:    rec.Star,
	}
}

func (rec *AddUnit) AddUnitToDomain() *property.Unit {
	return &property.Unit{
		Name:     rec.Name,
		Capacity: rec.Capacity,
		Surface:  rec.Surface,
	}
}
