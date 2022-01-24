package facility

import (
	"database/sql"
	"github.com/sewakantor/sw-be/businesses/facility"
	"gorm.io/gorm"
)

type facilityRepository struct {
	DB *gorm.DB
}

func NewRepoPostgres(db *gorm.DB) facility.Repository {
	return &facilityRepository{
		DB: db,
	}
}

func (repo *facilityRepository) Store(data *facility.Domain) (*facility.Domain, error) {
	facility := fromDomainFacility(data)
	if err := repo.DB.Create(&facility).Error; err != nil {
		return nil, err
	}

	return toDomainFacility(facility), nil
}

func (repo *facilityRepository) GetFacilityByGeo(lat, long float64) ([]facility.Domain, error) {
	qe := `SELECT *,
       (
           (
                   6371.04 * ACOS(((COS(((PI() / 2) - RADIANS((90 - lat)))) *
                                    COS(PI() / 2 - RADIANS(90 - @lat)) *
                                    COS((RADIANS(lng) - RADIANS(@long))))
                   + (SIN(((PI() / 2) - RADIANS((90 - lat)))) *
                      SIN(((PI() / 2) - RADIANS(90 - @lat))))))
               )
           ) as distance
 FROM "facilities";`


	var res []Result
	err := repo.DB.Raw(qe, sql.Named("lat", lat), sql.Named("long", long)).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return resultsToDomain(res), nil
}