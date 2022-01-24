package property

import (
	"fmt"
	"github.com/sewakantor/sw-be/businesses"
	"strconv"
	"strings"
)

type propertyService struct {
	propertyRepository Repository
}

func NewPropertyService(rep Repository) Service {
	return &propertyService{
		propertyRepository: rep,
	}
}

func (us *propertyService) AddComplex(data *Complex) (*Complex, error) {
	var res *Complex
	existedRecord, err := us.propertyRepository.GetComplexByName(data.Name)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			res, err = us.propertyRepository.StoreComplex(data)
		}
	}

	if existedRecord != nil {
		return nil, businesses.ErrComplexDuplicate
	}

	return res, nil
}

func (us *propertyService) DeleteComplex(ID string) error {
	complexID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return businesses.ErrInternalServer
	}

	_, err = us.propertyRepository.GetComplexByID(complexID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return businesses.ErrComplexNotFound
		}
	}

	err = us.propertyRepository.DeleteComplex(complexID)
	if err != nil {
		return businesses.ErrInternalServer
	}

	return nil
}

func (us *propertyService) GetAllComplex(name string) ([]Complex, error) {
	res, err := us.propertyRepository.GetAllComplex(name)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *propertyService) UpdateComplex(data *Complex, ID string) (*Complex, error) {
	complexID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	var res *Complex
	res, err = us.propertyRepository.UpdateComplex(data, complexID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, businesses.ErrComplexNotFound
		}
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *propertyService) AddBuilding(data *Building, ID string) (*Building, error) {
	var res *Building
	//existedRecord, err := us.propertyRepository.GetComplexByName(data.Name)
	complexID, err := strconv.ParseUint(ID, 10, 64)
	data.ComplexID = uint(complexID)
	res, err = us.propertyRepository.StoreBuilding(data)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign") {
			return nil, businesses.ErrComplexNotFound
		}
		return nil, businesses.ErrInternalServer
	}
	//if existedRecord != nil {
	//	return nil, businesses.ErrComplexDuplicate
	//}

	return res, nil
}

func (us *propertyService) GetAllBuilding(complexName string) ([]Building, error) {
	res, err := us.propertyRepository.GetAllBuilding(complexName)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	if res == nil {
		return nil, businesses.ErrBuildingNotFound
	}
	return res, nil
}

func (us *propertyService) GetRecommendedBuilding(limit string) ([]Building, error) {
	limitData, err := strconv.ParseUint(limit, 10, 64)
	res, err := us.propertyRepository.GetRecommendedBuilding(int(limitData))
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *propertyService) DeleteBuilding(ID string) error {
	complexID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return businesses.ErrInternalServer
	}

	_, err = us.propertyRepository.GetBuildingByID(uint(complexID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return businesses.ErrBuildingNotFound
		}
	}

	err = us.propertyRepository.DeleteBuilding(uint(complexID))
	if err != nil {
		return businesses.ErrInternalServer
	}

	return nil
}

func (us *propertyService) GetSingleBuilding(ID string) (*Building, error) {
	buildingID, err := strconv.ParseUint(ID, 10, 64)
	res, err := us.propertyRepository.GetSingleBuilding(uint(buildingID))
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	if res.ID == 0 {
		return nil, businesses.ErrBuildingNotFound
	}

	return res, nil
}

func (us *propertyService) UpdateBuilding(data *Building, ID string) (*Building, error) {
	buildingID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	var res *Building
	res, err = us.propertyRepository.UpdateBuilding(data, uint(buildingID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, businesses.ErrBuildingNotFound
		}
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *propertyService) AddReview(data *Review, buildingID string, usersID uint) (*Review, error) {
	buildID, err := strconv.ParseUint(buildingID, 10, 64)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	data.Status = 0
	data.UserID = usersID
	data.BuildingID = uint(buildID)
	res, err := us.propertyRepository.StoreReview(data)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign") {
			return nil, businesses.ErrBuildingNotFound
		}
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *propertyService) ApproveReview(ID string) (*Review, error) {
	complexID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		fmt.Println(err)
		return nil, businesses.ErrInternalServer
	}
	res, err := us.propertyRepository.ApproveReview(uint(complexID))

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, businesses.ErrReviewNotFound
		}
		fmt.Println(err)
		return nil, businesses.ErrInternalServer
	}
	return res, nil
}

func (us *propertyService) GetAllReview(buildingID, limit, isApprove string) ([]Review, error) {
	buildID, err := strconv.ParseUint(buildingID, 10, 64)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	fmt.Println(isApprove)
	fmt.Println(limit)
	var lmt uint64
	var isApproved bool
	if limit != "" && isApprove != "" {
		lmt, err = strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return nil, businesses.ErrInternalServer
		}
		isApproved, err = strconv.ParseBool(isApprove)
		fmt.Println(isApproved)
		if err != nil {
			return nil, businesses.ErrInternalServer
		}
	} else if isApprove != ""{
		isApproved, err = strconv.ParseBool(isApprove)
		if err != nil {
			return nil, businesses.ErrInternalServer
		}
	} else if limit != "" {
		lmt, err = strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return nil, businesses.ErrInternalServer
		}
	}
	res, err := us.propertyRepository.GetAllReview(uint(buildID), uint(lmt) ,isApproved)

	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}


