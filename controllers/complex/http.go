package complex

import (
	"github.com/labstack/echo/v4"
	"github.com/sewakantor/sw-be/businesses/complex"
	"github.com/sewakantor/sw-be/controllers/complex/request"
	"github.com/sewakantor/sw-be/controllers/complex/response"
	"github.com/sewakantor/sw-be/helpers"
	"net/http"
	"strings"
)

type ComplexControllers struct {
	complexService complex.Service
}

func NewComplexController(uc complex.Service) *ComplexControllers {
	return &ComplexControllers{
		complexService: uc,
	}
}

func (ctrl *ComplexControllers) AddComplex(c echo.Context) error {
	req := new(request.AddComplex)
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

	res, err := ctrl.complexService.Add(req.AddComplexToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Successfully created a complex!",
			response.FromDomain(res)))
}

func (ctrl *ComplexControllers) DeleteComplex(c echo.Context) error {
	id := c.Param("id")

	err := ctrl.complexService.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Complex Not Found!",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Successfully created a complex!",""))
}
