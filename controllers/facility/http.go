package facility

import (
	"github.com/labstack/echo/v4"
	"github.com/sewakantor/sw-be/businesses/facility"
	"github.com/sewakantor/sw-be/controllers/facility/request"
	"github.com/sewakantor/sw-be/controllers/facility/response"
	"github.com/sewakantor/sw-be/helpers"
	"net/http"
)

type FacilityControllers struct {
	facilityService facility.Service
}

func NewPropertyController(uc facility.Service) *FacilityControllers {
	return &FacilityControllers{
		facilityService: uc,
	}
}

func (ctrl *FacilityControllers) AddFacility(c echo.Context) error {
	req := new(request.Facility)
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

	res, err := ctrl.facilityService.AddFacility(req.FacilityToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Successfully created a facility!",
			res))
}

func (ctrl *FacilityControllers) GetFacility(c echo.Context) error {
	long := c.QueryParam("long")
	lat := c.QueryParam("lat")

	res, err := ctrl.facilityService.UnitsByGeo(long, lat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Success get detail unit!",
			response.FromDomainReviewsSpecific(res)))
}