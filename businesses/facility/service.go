package facility

import (
	"github.com/sewakantor/sw-be/businesses"
	"strconv"
)

type facilityService struct {
	facilityRepository Repository
}

func NewPropertyService(rep Repository) Service {
	return &facilityService{
		facilityRepository: rep,
	}
}

func (us *facilityService) AddFacility(data *Domain) (*Domain, error) {
	var res *Domain
	res, err := us.facilityRepository.Store(data)
	if err != nil {
			return nil, businesses.ErrInternalServer
	}

	return res, nil
}


func (us *facilityService) UnitsByGeo(long, lat string) ([]Domain, error) {
	var res []Domain
	var err error
	if long != "" && lat != "" {
		longFloat, _ := strconv.ParseFloat(lat, 64)
		latFloat, _ := strconv.ParseFloat(long, 64)
		res, err = us.facilityRepository.GetFacilityByGeo(latFloat, longFloat)
		if err != nil {
			return nil, businesses.ErrInternalServer
		}
	}

	if res == nil {
		return nil, businesses.ErrBuildingNotFound
	}

	return res, nil
}
