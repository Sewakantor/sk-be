package property

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sewakantor/sw-be/app/middleware"
	"github.com/sewakantor/sw-be/businesses/property"
	"github.com/sewakantor/sw-be/controllers/property/request"
	"github.com/sewakantor/sw-be/controllers/property/response"
	"github.com/sewakantor/sw-be/helpers"
	"net/http"
	"strings"
)

type PropertyControllers struct {
	complexService property.Service
}

func NewPropertyController(uc property.Service) *PropertyControllers {
	return &PropertyControllers{
		complexService: uc,
	}
}

func (ctrl *PropertyControllers) AddComplex(c echo.Context) error {
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

	res, err := ctrl.complexService.AddComplex(req.AddComplexToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Successfully created a complex!",
			response.FromDomainComplex(res)))
}

func (ctrl *PropertyControllers) DeleteComplex(c echo.Context) error {
	id := c.Param("id")

	err := ctrl.complexService.DeleteComplex(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Not Found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully delete a complex!", map[string]string{"id": id}))
}

func (ctrl *PropertyControllers) GetAllComplex(c echo.Context) error {
	name := c.QueryParam("name")
	res, err := ctrl.complexService.GetAllComplex(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully get a complex!", res))
}

func (ctrl *PropertyControllers) UpdateComplex(c echo.Context) error {
	req := new(request.AddComplex)
	ID := c.Param("id")
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

	res, err := ctrl.complexService.UpdateComplex(req.AddComplexToDomain(), ID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Not Found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully updated a complex!",
			response.FromDomainComplex(res)))
}

func (ctrl *PropertyControllers) AddBuilding(c echo.Context) error {
	req := new(request.AddBuilding)
	id := c.Param("id")
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

	res, err := ctrl.complexService.AddBuilding(req.AddBuildingToDomain(), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Not Found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Successfully created a building!",
			response.FromDomainBuilding(res)))
}

func (ctrl *PropertyControllers) GetAllBuilding(c echo.Context) error {
	complexName := c.QueryParam("complex_name")
	res, err := ctrl.complexService.GetAllBuilding(complexName)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Not Found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully get all building!", res))
}

func (ctrl *PropertyControllers) GetRecommendedBuilding(c echo.Context) error {
	limit := c.QueryParam("limit")
	res, err := ctrl.complexService.GetRecommendedBuilding(limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully get recommended building!", res))
}

func (ctrl *PropertyControllers) DeleteBuilding(c echo.Context) error {
	id := c.Param("id")

	err := ctrl.complexService.DeleteBuilding(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Not Found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully delete a building!", map[string]string{"id": id}))
}

func (ctrl *PropertyControllers) GetSingleBuilding(c echo.Context) error {
	id := c.Param("id")

	res, err := ctrl.complexService.GetSingleBuilding(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Not Found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully get a data!", res))
}

func (ctrl *PropertyControllers) UpdateBuilding(c echo.Context) error {
	req := new(request.AddBuilding)
	ID := c.Param("id")
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

	res, err := ctrl.complexService.UpdateBuilding(req.AddBuildingToDomain(), ID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Not Found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully updated a building!",
			response.FromDomainBuilding(res)))
}

func (ctrl *PropertyControllers) AddReview(c echo.Context) error {
	req := new(request.AddReview)
	claim := middleware.GetUser(c)
	buildingID := c.Param("buildingID")
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
	res, err := ctrl.complexService.AddReview(req.AddReviewToDomain(), buildingID, claim.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Successfully created a review!",
			response.FromDomainReview(res)))
}

func (ctrl *PropertyControllers) ApproveReview(c echo.Context) error {
	ID := c.Param("id")
	res, err := ctrl.complexService.ApproveReview(ID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Not Found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully approve a review!",
			response.FromDomainReview(res)))
}

func (ctrl *PropertyControllers) GetAllReview(c echo.Context) error {
	ID := c.Param("id")
	limit := c.QueryParam("limit")
	isApprove := c.QueryParam("is_approve")
	res, err := ctrl.complexService.GetAllReview(ID, limit, isApprove)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	if res == nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Not Found",
				errors.New("reviews not found!"), helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully approve a review!",
			response.FromDomainReviewsSpecific(res)))
}

func (ctrl *PropertyControllers) AddUnit(c echo.Context) error {
	req := new(request.AddUnit)
	buildingID := c.Param("buildingID")
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
	res, err := ctrl.complexService.AddUnit(req.AddUnitToDomain(), buildingID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Successfully created a unit!",
			response.FromDomainUnit(res)))
}

func (ctrl *PropertyControllers) DeleteUnit(c echo.Context) error {
	id := c.Param("id")

	err := ctrl.complexService.DeleteUnit(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Not Found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully delete a unit!", map[string]string{"id": id}))
}

func (ctrl *PropertyControllers) GetAllUnit(c echo.Context) error {
	res, err := ctrl.complexService.GetAllUnit()
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Successfully get a complex!", response.FromDomainUnitsSpecific(res)))
}
