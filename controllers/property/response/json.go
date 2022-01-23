package response

import (
	"github.com/sewakantor/sw-be/businesses/property"
	"time"
)

type Complex struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Street     string    `json:"street"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	Country    string    `json:"country"`
	PostalCode int       `json:"postal-code"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	CreatedAt  time.Time `json:"created_at"`
}

type Building struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Photo        string    `json:"photo"`
	Year         int       `json:"year"`
	Floor        int       `json:"floor"`
	FloorSurface int       `json:"floor_surface"`
	TotalSurface int       `json:"total_surface"`
	Price        int       `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
}

func FromDomainComplex(domain *property.Complex) *Complex {
	return &Complex{
		ID:         domain.ID,
		Name:       domain.Name,
		Street:     domain.Street,
		City:       domain.City,
		State:      domain.State,
		Country:    domain.Country,
		PostalCode: domain.PostalCode,
		Latitude:   domain.Latitude,
		Longitude:  domain.Longitude,
		CreatedAt:  domain.CreatedAt,
	}
}

func FromDomainBuilding(domain *property.Building) *Building {
	return &Building{
		ID:           domain.ID,
		Name:         domain.Name,
		Photo:        domain.Photo,
		Year:         domain.Year,
		Floor:        domain.Floor,
		FloorSurface: domain.FloorSurface,
		TotalSurface: domain.TotalSurface,
		Price:        domain.Price,
		CreatedAt:    domain.CreatedAt,
	}
}
