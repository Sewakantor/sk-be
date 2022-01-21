package complex

import (
	"time"
)

type Domain struct {
	ID         uint
	Name       string
	Street     string
	City       string
	State      string
	Country    string
	PostalCode int
	Latitude   float64
	Longitude  float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	Add(data *Domain) (*Domain, error)
	Delete(ID string) error
}

type Repository interface {
	Store(data *Domain) (*Domain, error)
	Delete(ID uint64) error
	GetByID (ID uint64) (*Domain, error)
}
