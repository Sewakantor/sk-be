package users

import (
	"github.com/sewakantor/sw-be/businesses/users"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uint      `gorm:"primary_key:auto_increment"`
	Name     string    `gorm:"type:varchar(255)"`
	Email    string    `gorm:"uniqueIndex;type:varchar(255)"`
	Password string    `gorm:"->;<-;not null;type:varchar(255)" `
	Company  string    `gorm:"type:varchar(255)"`
	Role     string    `gorm:"type:varchar(255)"`
	Status   int       `gorm:"default:0;size:10"`
}

func ToDomain(rec *Users) *users.Domain {
	return &users.Domain{
		ID: 	   rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Company:   rec.Company,
		Role:      rec.Role,
		Status:    rec.Status,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func ToDomainUser(rec *Users) *users.Domain {
	return &users.Domain{
		ID: 	   rec.ID,
		Name:      rec.Name,
		Company:   rec.Company,
	}
}

func fromDomain(domain *users.Domain) *Users {
	return &Users{
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		Company:  domain.Company,
		Role:     domain.Role,
		Status:   domain.Status,
	}
}