package users

import (
	"time"
)

type Domain struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Company   string
	Role      string
	Phone	  int
	Status    int
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Registration(data *Domain) (*Domain, error)
	Activation(userID string) (*Domain, error)
	Login(email, password string) (*Domain, error)
}

type Repository interface {
	StoreNewUsers(data *Domain) (*Domain, error)
	GetByEmail(email string) (*Domain, error)
	GetByUserID(userID uint64) (*Domain, error)
	UpdateStatus(userID uint64) error
	GetByEmailAndPassword(email string) (*Domain, error)
}
