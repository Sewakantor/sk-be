package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sewakantor/sw-be/app/config"
	_middleware "github.com/sewakantor/sw-be/app/middleware"
	"github.com/sewakantor/sw-be/app/routes"
	_propertyService "github.com/sewakantor/sw-be/businesses/property"
	_usersService "github.com/sewakantor/sw-be/businesses/users"
	_facilityService "github.com/sewakantor/sw-be/businesses/facility"
	_propertyController "github.com/sewakantor/sw-be/controllers/property"
	_usersController "github.com/sewakantor/sw-be/controllers/users"
	_facilityController "github.com/sewakantor/sw-be/controllers/facility"
	"github.com/sewakantor/sw-be/helpers"
	_propertyRepo "github.com/sewakantor/sw-be/repository/databases/property"
	_usersRepo "github.com/sewakantor/sw-be/repository/databases/users"
	_facilityRepo "github.com/sewakantor/sw-be/repository/databases/facility"
	"log"
	"os"
	"strconv"
)

func main() {
	var (
		db = config.SetupDatabaseConnection()
	)

	port := os.Getenv("PORT")
	timeJWT, _ := strconv.Atoi(os.Getenv("JWT_TOKEN_AGE"))
	secretToken := os.Getenv("SECRET_TOKEN_KEY")
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       secretToken,
		ExpiresDuration: timeJWT,
	}

	e := echo.New()
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(_middleware.LoggerConfig()))

	userRepo := _usersRepo.NewRepoMySQL(db)
	userService := _usersService.NewUserService(userRepo, &configJWT)
	userCtrl := _usersController.NewUserController(userService)

	propertyRepo := _propertyRepo.NewRepoPostgres(db)
	propertyService := _propertyService.NewPropertyService(propertyRepo)
	propertyCtrl := _propertyController.NewPropertyController(propertyService)

	facilityRepo := _facilityRepo.NewRepoPostgres(db)
	facilityService := _facilityService.NewPropertyService(facilityRepo)
	facilityCtrl := _facilityController.NewPropertyController(facilityService)

	routesInit := routes.ControllerList{
		JWTMiddleware:  configJWT.Init(),
		UserController: *userCtrl,
		PropertyController: *propertyCtrl,
		FacilityController: *facilityCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":" + port))
}
