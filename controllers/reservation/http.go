package reservation

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sewakantor/sw-be/businesses/reservation"
	"github.com/sewakantor/sw-be/controllers/reservation/request"
	"github.com/sewakantor/sw-be/controllers/reservation/response"
	"github.com/sewakantor/sw-be/helpers"
	"net/http"
)

type ReservationController struct {
	reservationService  reservation.Service
}

func NewReservationController(uc reservation.Service) *ReservationController {
	return &ReservationController{
		reservationService:  uc,
	}
}

func (ctrl *ReservationController) Reservation(c echo.Context) error {
	id := c.Param("id")

	req := new(request.Reservation)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	res, err := ctrl.reservationService.Reservation(request.ToDomain(req, id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Success reservation hotel!",
			map[string]int{"price":res.Price}))
}

func (ctrl *ReservationController) GetReservation(c echo.Context) error {
	ID := c.Param("id")
	res, err := ctrl.reservationService.GetByByCustID(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	if res == nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Not Found",
				errors.New("reservation not found"), helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully get a reservation!", response.FromDomainUnitsSpecific(res)))
}
