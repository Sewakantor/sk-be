package reservation

import (
	"github.com/sewakantor/sw-be/businesses"
	"github.com/sewakantor/sw-be/businesses/property"
	"github.com/sewakantor/sw-be/businesses/users"
	"strconv"
)

type reservationService struct {
	reservationRepository Repository
	userService           users.Service
	unitService           property.Service
}

func NewReservationService(rep Repository, us users.Service, unitserv property.Service) Service {
	return &reservationService{
		reservationRepository: rep,
		userService:           us,
		unitService:           unitserv,
	}
}

func (rs *reservationService) Reservation(data *Domain) (*Domain, error) {
	var err error

	err = rs.reservationRepository.GetByDate(data)
	if err == nil {
		return nil, businesses.ErrUnitReserved
	}
	res, err := rs.reservationRepository.Store(data)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (rs *reservationService) GetByByCustID(ID string) ([]Domain, error) {
	var err error
	buildID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	res, err := rs.reservationRepository.GetReservationByCustID(uint(buildID))
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

