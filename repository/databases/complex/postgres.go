package complex

import (
	"github.com/sewakantor/sw-be/businesses/complex"
	"gorm.io/gorm"
)

type complexRepository struct {
	DB *gorm.DB
}

func NewRepoPostgres(db *gorm.DB) complex.Repository {
	return &complexRepository{
		DB: db,
	}
}

func (repo *complexRepository) Store(data *complex.Domain) (*complex.Domain, error) {
	complex := fromDomain(data)
	if err := repo.DB.Create(&complex).Error; err != nil {
		return nil, err
	}

	return toDomain(complex), nil
}

func (repo *complexRepository) Delete(ID uint64) error {
	var complex Complex
	if err := repo.DB.Debug().Where("id = ?", ID).Delete(&complex).Error; err != nil {
		return err
	}
	return nil
}

func (repo *complexRepository) GetByID (ID uint64) (*complex.Domain, error) {
	var user Complex
	if err := repo.DB.Where("id = ?", ID).First(&user).Error; err != nil {
		return nil, err
	}
	result := toDomain(&user)
	return result, nil
}
