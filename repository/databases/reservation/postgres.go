package reservations

import (
	"github.com/sewakantor/sw-be/businesses/reservation"
	"gorm.io/gorm"
)

type repoUnit struct {
	DB *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) reservation.Repository {
	return &repoUnit{
		DB: db,
	}
}

func (ur *repoUnit) Store(domain *reservation.Domain) (*reservation.Domain, error) {
	res := fromDomain(domain)

	ur.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&res).Error; err != nil {
			return err
		}

		return nil
	})

	return toDomain(res), nil
}

func (ur *repoUnit) GetByDate(domain *reservation.Domain) error {
	var rev Reservation
	if err := ur.DB.Where("check_in_date >= ? AND check_out_date <= ? AND unit_id = ?",
		domain.CheckInDate, domain.CheckOutDate, domain.UnitID).First(&rev).Error; err != nil {
		return err
	}
	return nil
}

func (ur *repoUnit) GetReservationByCustID(ID uint) ([]reservation.Domain, error) {
	var resv []Reservation
	if err := ur.DB.Joins("Units").Where("customer_id = ?", ID).Find(&resv).Error; err != nil {
		return nil, err
	}
	result := ToDomains(resv)
	return result, nil
}
