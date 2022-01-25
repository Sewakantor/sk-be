package property

import (
	"fmt"
	"github.com/sewakantor/sw-be/businesses/property"
	"gorm.io/gorm"
)

type propertyRepository struct {
	DB *gorm.DB
}

func NewRepoPostgres(db *gorm.DB) property.Repository {
	return &propertyRepository{
		DB: db,
	}
}

func (repo *propertyRepository) StoreComplex(data *property.Complex) (*property.Complex, error) {
	complex := fromDomainComplex(data)
	if err := repo.DB.Create(&complex).Error; err != nil {
		return nil, err
	}

	return toDomainComplex(complex), nil
}

func (repo *propertyRepository) DeleteComplex(ID uint64) error {
	var complex Complex
	if err := repo.DB.Debug().Where("id = ?", ID).Delete(&complex).Error; err != nil {
		return err
	}
	return nil
}

func (repo *propertyRepository) GetComplexByID(ID uint64) (*property.Complex, error) {
	var complex Complex
	if err := repo.DB.Where("id = ?", ID).First(&complex).Error; err != nil {
		return nil, err
	}
	result := toDomainComplex(&complex)
	return result, nil
}

func (repo *propertyRepository) GetAllComplex(name string) ([]property.Complex, error) {
	var complex []Complex
	if err := repo.DB.Where("name LIKE ?", "%"+name+"%").Find(&complex).Error; err != nil {
		return nil, err
	}
	fmt.Println(complex)
	return ToComplexesDomain(complex), nil
}

func (repo *propertyRepository) GetComplexByName(name string) (*property.Complex, error) {
	var complex Complex
	if err := repo.DB.Where("name = ?", name).First(&complex).Error; err != nil {
		return nil, err
	}
	result := toDomainComplex(&complex)
	return result, nil
}

func (repo *propertyRepository) UpdateComplex(data *property.Complex, ID uint64) (*property.Complex, error) {
	var complex Complex
	if err := repo.DB.Where("id = ?", ID).First(&complex).Error; err != nil {
		return nil, err
	}

	complex.Name = data.Name
	complex.Street = data.Street
	complex.City = data.City
	complex.State = data.State
	complex.Country = data.Country
	complex.PostalCode = data.PostalCode
	complex.Lng = data.Longitude
	complex.Lat = data.Latitude

	if err := repo.DB.Save(&complex).Error; err != nil {
		return nil, err
	}
	result := toDomainComplex(&complex)
	return result, nil
}

func (repo *propertyRepository) StoreBuilding(data *property.Building) (*property.Building, error) {
	building := fromDomainBuilding(data)
	if err := repo.DB.Debug().Create(&building).Error; err != nil {
		return nil, err
	}
	return toDomainBuilding(building), nil
}

func (repo *propertyRepository) GetAllBuilding(complexName string) ([]property.Building, error) {
	var building []Building
	if err := repo.DB.Joins("Complexes", repo.DB.Where(&Complex{Name: complexName})).Find(&building).Error; err != nil {
		return nil, err
	}
	return ToBuildingsDomain(building), nil
}

func (repo *propertyRepository) GetRecommendedBuilding(limit int) ([]property.Building, error) {
	var building []Building
	if limit == 0 {
		limit = 3
	}
	if err := repo.DB.Debug().Limit(limit).Joins("Complexes").Find(&building).Error; err != nil {
		return nil, err
	}
	return ToBuildingsDomain(building), nil
}

func (repo *propertyRepository) DeleteBuilding(ID uint) error {
	var building Building
	if err := repo.DB.Debug().Where("id = ?", ID).Delete(&building).Error; err != nil {
		return err
	}
	return nil
}

func (repo *propertyRepository) GetBuildingByID(ID uint) (*property.Building, error) {
	var building Building
	if err := repo.DB.Where("id = ?", ID).First(&building).Error; err != nil {
		return nil, err
	}
	result := toDomainBuilding(&building)
	return result, nil
}

func (repo *propertyRepository) GetSingleBuilding(ID uint) (*property.Building, error) {
	var building Building
	if err := repo.DB.Debug().Joins("Complexes").Where("buildings.id = ?", ID).Find(&building).Error; err != nil {
		return nil, err
	}
	result := toDomainBuilding(&building)
	return result, nil
}

func (repo *propertyRepository) UpdateBuilding(data *property.Building, ID uint) (*property.Building, error) {
	var building Building
	if err := repo.DB.Where("id = ?", ID).First(&building).Error; err != nil {
		return nil, err
	}

	building.Name = data.Name
	building.Price = data.Price
	building.Photo = data.Photo
	building.Year = data.Year
	building.FloorSurface = data.FloorSurface
	building.Floor = data.Floor
	building.TotalSurface = data.TotalSurface

	if err := repo.DB.Save(&building).Error; err != nil {
		return nil, err
	}
	result := toDomainBuilding(&building)
	return result, nil
}

func (repo *propertyRepository) StoreReview(data *property.Review) (*property.Review, error) {
	review := fromDomainReview(data)
	if err := repo.DB.Create(&review).Error; err != nil {
		return nil, err
	}

	return toDomainReview(review), nil
}

func (repo *propertyRepository) ApproveReview(ID uint) (*property.Review, error) {
	var review Review
	if err := repo.DB.Where("id = ?", ID).First(&review).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	review.Status = 1

	if err := repo.DB.Save(&review).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := toDomainReview(&review)
	return result, nil
}

func (repo *propertyRepository) GetAllReview(buildingID uint, limit uint, isApprove bool) ([]property.Review, error) {
	var review []Review
	var err error
	if isApprove {
		if limit == 0 {
			err = repo.DB.Joins("Users").Where("building_id = ? AND reviews.status = 1", buildingID).Find(&review).Error
		} else {
			err = repo.DB.Joins("Users").Limit(int(limit)).Where("building_id = ? AND reviews.status = 1", buildingID).Find(&review).Error
		}
	} else {
		if limit == 0 {
			err = repo.DB.Joins("Users").Where("building_id = ?", buildingID).Find(&review).Error
		} else {
			err = repo.DB.Joins("Users").Limit(int(limit)).Where("building_id = ?", buildingID).Find(&review).Error
		}
	}
	if err != nil {
		return nil, err
	}
	return ToReviewsDomain(review), nil
}
func (repo *propertyRepository) StoreUnit(data *property.Unit) (*property.Unit, error) {
	unit := fromDomainUnit(data)
	if err := repo.DB.Create(&unit).Error; err != nil {
		return nil, err
	}
	return toDomainUnit(unit), nil
}
