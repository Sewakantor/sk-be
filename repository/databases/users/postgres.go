package users

import (
	"github.com/sewakantor/sw-be/businesses/users"
	"gorm.io/gorm"
)

type repoUsers struct {
	DB *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) users.Repository {
	return &repoUsers{
		DB: db,
	}
}

func (ru *repoUsers) StoreNewUsers (data *users.Domain) (*users.Domain, error) {
	user := fromDomain(data)
	if err := ru.DB.Create(&user); err.Error != nil {
		return nil, err.Error
	}
	result := toDomain(user)
	return result, nil
}

func (ru *repoUsers) GetByEmail (email string) (*users.Domain, error) {
	var user Users
	if err := ru.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	result := toDomain(&user)
	return result, nil
}

func (ru *repoUsers) UpdateStatus(userID uint64) error {
	var user Users
	if err := ru.DB.Model(&user).Where("id = ?", userID).Update("status", 1).Error; err != nil {
		return err
	}

	return nil
}

func (ru *repoUsers) GetByUserID(userID uint64) (*users.Domain, error) {
	var user Users
	if err := ru.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	result := toDomain(&user)
	return result, nil
}

func (ru *repoUsers) GetByEmailAndPassword(email string) (*users.Domain, error) {
	var user Users
	if err := ru.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	result := toDomain(&user)
	return result, nil
}
