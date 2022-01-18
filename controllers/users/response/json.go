package response

import (
	"github.com/sewakantor/sw-be/businesses/users"
	"time"
)

type Users struct {
	ID    uint    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role	  string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginUsers struct {
	ID    	  uint    `json:"id"`
	Name      string    `json:"name"`
	Role	  string    `json:"role"`
	Token     string    `json:"token"`
}

func FromDomain(domain *users.Domain) *Users {
	return &Users{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Role:      domain.Role,
		CreatedAt: domain.CreatedAt,
	}
}

func LoginFromDomain(domain *users.Domain) *LoginUsers {
	return &LoginUsers{
		ID:     domain.ID,
		Name:   domain.Name,
		Role:   domain.Role,
		Token:  domain.Token,
	}
}