package complex

import (
	"github.com/sewakantor/sw-be/businesses"
	"strconv"
	"strings"
)

type complexService struct {
	complexRepository Repository
}

func NewComplexService(rep Repository) Service {
	return &complexService{
		complexRepository: rep,
	}
}

func (us *complexService) Add(data *Domain) (*Domain, error) {
	res, err := us.complexRepository.Store(data)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *complexService) Delete(ID string) error {
	complexID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return businesses.ErrInternalServer
	}

	_, err = us.complexRepository.GetByID(complexID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return businesses.ErrComplexNotFound
		}
	}

	err = us.complexRepository.Delete(complexID)
	if err != nil {
		return businesses.ErrInternalServer
	}

	return nil
}
