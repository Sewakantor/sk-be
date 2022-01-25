package request

import (
	"github.com/sewakantor/sw-be/businesses/reservation"
	"strconv"
	"time"
)

type Reservation struct {
	CustomerID   uint   `json:"user_id" validate:"required"`
	CheckInDate  string `json:"check_in_date" validate:"required"`
	CheckOutDate string `json:"check_out_date" validate:"required"`
	Price        int    `json:"price" validate:"required"`
}

func ToDomain(data *Reservation, claim string) *reservation.Domain {
	buildID, _ := strconv.ParseUint(claim, 10, 64)
	startDate, _ := time.Parse("2006-01-02", data.CheckInDate)
	endDate, _ := time.Parse("2006-01-02", data.CheckOutDate)
	return &reservation.Domain{
		CustomerID:   data.CustomerID,
		UnitID:       uint(buildID),
		CheckInDate:  startDate,
		CheckOutDate: endDate,
		Price:        data.Price,
	}
}
