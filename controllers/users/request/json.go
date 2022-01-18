package request

import "github.com/sewakantor/sw-be/businesses/users"

type UserRegistration struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,password"`
	Role      string `json:"role" validate:"required,role"`
	Company   string `json:"company" validate:"required"`
	Phone 	  int    `json:"phone" validate:"required"`
}

type Login struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func (rec *UserRegistration) UserRegistrationToDomain() *users.Domain{
	return &users.Domain{
		Name      :rec.Name,
		Email     :rec.Email,
		Password  :rec.Password,
		Role      :rec.Role,
		Phone	  :rec.Phone,
		Company   :rec.Company,
	}
}