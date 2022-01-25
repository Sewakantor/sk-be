package response

import (
	"github.com/sewakantor/sw-be/businesses/reservation"
	"time"
)

type ReservationSpecific struct {
	ID   uint   `json:"id"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	Status       string    `json:"status"`
	Price        int       `json:"price"`
	Unit         Units     `json:"unit"`
}

type Units struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func FromDomainReservationSpecific(domain *reservation.Domain) *ReservationSpecific {
	return &ReservationSpecific{
		ID:        domain.ID,
		CheckInDate:   domain.CheckInDate,
		CheckOutDate:  domain.CheckOutDate,
		Status: domain.Status,
		Price: domain.Price,
		Unit: Units{ID: domain.Unit.ID, Name: domain.Unit.Name},
	}
}

func FromDomainUnitsSpecific(data []reservation.Domain) []ReservationSpecific {
	var res []ReservationSpecific
	for _, s := range data {
		res = append(res, *FromDomainReservationSpecific(&s))
	}
	return res
}
