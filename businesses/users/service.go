package users

import (
	"fmt"
	"github.com/sewakantor/sw-be/app/middleware"
	"github.com/sewakantor/sw-be/businesses"
	"github.com/sewakantor/sw-be/helpers"
	"strconv"
	"strings"
)

type userService struct {
	userRepository Repository
	jwtAuth        *middleware.ConfigJWT
}

func NewUserService(rep Repository, jwt *middleware.ConfigJWT) Service {
	return &userService{
		userRepository: rep,
		jwtAuth:        jwt,
	}
}

func (us *userService) Registration(userDomain *Domain) (*Domain, error) {
	existedUser, err := us.userRepository.GetByEmail(userDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedUser != nil {
		return nil, businesses.ErrEmailDuplicate
	}
	userDomain.Password, err = helpers.HashPassword(userDomain.Password)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	res, err := us.userRepository.StoreNewUsers(userDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	err = helpers.SendEmail(res.Email, "register", res.Name, "activation/"+strconv.FormatUint(uint64(res.ID), 10))
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	return res, nil
}

func (us *userService) Activation(userID string) (*Domain, error) {
	ID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	res, err := us.userRepository.GetByUserID(ID)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		} else {
			return nil, businesses.ErrAccountNotFound
		}
	}

	if res.Status == 1 {
		return nil, businesses.ErrAccountActivated
	}

	if err = us.userRepository.UpdateStatus(ID); err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *userService) Login(email, password string) (*Domain, error) {

	res, err := us.userRepository.GetByEmailAndPassword(email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
		return nil, businesses.ErrAccountNotFound
	}

	if !helpers.ValidateHash(password, res.Password) {
		return nil, businesses.ErrInvalidCredential
	}


	if res.Status != 1 {
		return nil, businesses.ErrAccountUnactivated
	}
	res.Token = us.jwtAuth.GenerateToken(res.ID, res.Role)
	return res, nil
}
