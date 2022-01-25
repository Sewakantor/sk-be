package reservation

import (
	"github.com/sewakantor/sw-be/businesses/property"
	"github.com/sewakantor/sw-be/businesses/users"
	"time"
)

type Domain struct {
	ID           uint
	CustomerID   uint
	Customer     users.Domain
	UnitID       uint
	Unit         property.Unit
	CheckInDate  time.Time
	CheckOutDate time.Time
	Status       string
	Price        int
}

type Service interface {
	Reservation(data *Domain) (*Domain, error)
	GetByByCustID(ID string) ([]Domain, error)
}

type Repository interface {
	Store(domain *Domain) (*Domain, error)
	GetByDate(domain *Domain) error
	GetReservationByCustID(ID uint) ([]Domain, error)
}
