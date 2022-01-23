package users

import (
	echo "github.com/labstack/echo/v4"
	"github.com/sewakantor/sw-be/businesses/users"
	"github.com/sewakantor/sw-be/controllers/users/request"
	"github.com/sewakantor/sw-be/controllers/users/response"
	"github.com/sewakantor/sw-be/helpers"
	"net/http"
	"strings"
)

type UserController struct {
	userService  users.Service
}

func NewUserController(uc users.Service) *UserController {
	return &UserController{
		userService:  uc,
	}
}

func (ctrl *UserController) Registration(c echo.Context) error {
	req := new(request.UserRegistration)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	res, err := ctrl.userService.Registration(req.UserRegistrationToDomain())
	if err != nil {
		if strings.Contains(err.Error(), "taken") {
			return c.JSON(http.StatusConflict,
				helpers.BuildErrorResponse("Conflict Data",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Successfully created an account, please check your email to activate!",
			response.FromDomain(res)))
}

func (ctrl *UserController) Activation(c echo.Context) error {
	ID := c.Param("id")

	res, err := ctrl.userService.Activation(ID)
	if err != nil {
		if strings.Contains(err.Error(), "activated") {
			return c.JSON(http.StatusConflict,
				helpers.BuildErrorResponse("Conflict Data",
					err, helpers.EmptyObj{}))
		} else if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Data Not found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully activation an account!",
			response.FromDomain(res)))
}

func (ctrl *UserController) Login(c echo.Context) error {
	req := new(request.Login)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Not Found",
				err, helpers.EmptyObj{}))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	res, err := ctrl.userService.Login(req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "not match") {
			return c.JSON(http.StatusUnauthorized,
				helpers.BuildErrorResponse("Wrong credentials",
					err, helpers.EmptyObj{}))
		} else if strings.Contains(err.Error(), "not been activated") {
			return c.JSON(http.StatusForbidden,
				helpers.BuildErrorResponse("Forbidden",
					err, helpers.EmptyObj{}))
		} else if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Not Found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully login an account!",
			response.LoginFromDomain(res)))
}