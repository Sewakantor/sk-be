package response

import (
	"github.com/sewakantor/sw-be/businesses/complex"
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

func FromDomain(domain *complex.Domain) *Complex {
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
