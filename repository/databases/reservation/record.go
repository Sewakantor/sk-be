package reservations

import (
	"github.com/sewakantor/sw-be/businesses/reservation"
	"github.com/sewakantor/sw-be/repository/databases/property"
	"github.com/sewakantor/sw-be/repository/databases/users"
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	CustomerID   uint
	Customer     users.Users `gorm:"foreignKey:CustomerID"`
	UnitID       uint
	Units        property.Unit `gorm:"foreignKey:UnitID"`
	CheckInDate  time.Time
	CheckOutDate time.Time
	Status       string
	Price        int
}

func fromDomain(domain *reservation.Domain) *Reservation {
	return &Reservation{
		CustomerID: domain.CustomerID,
		UnitID: domain.UnitID,
		CheckInDate: domain.CheckInDate,
		CheckOutDate: domain.CheckOutDate,
		Status: domain.Status,
		Price: domain.Price,
	}
}

func toDomain(res *Reservation) *reservation.Domain {
	return &reservation.Domain{
		ID: res.ID,
		CustomerID: res.CustomerID,
		Unit: *property.ToDomainUnit(&res.Units),
		UnitID: res.UnitID,
		Customer: *users.ToDomainUser(&res.Customer),
		CheckInDate: res.CheckInDate,
		CheckOutDate: res.CheckOutDate,
		Status: "unconfirmed",
		Price: res.Price,
	}
}

func ToDomains(data []Reservation) []reservation.Domain {
	var res []reservation.Domain
	for _, s := range data {
		res = append(res, *toDomain(&s))
	}
	return res
}