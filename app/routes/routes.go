package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_middleware "github.com/sewakantor/sw-be/app/middleware"
	"github.com/sewakantor/sw-be/controllers/property"
	"github.com/sewakantor/sw-be/controllers/users"
)

type ControllerList struct {
	UserController    users.UserController
	JWTMiddleware     middleware.JWTConfig
	PropertyController property.PropertyControllers
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", cl.UserController.Registration)
	users.POST("/activation/:id", cl.UserController.Activation)

	auth := e.Group("auth")
	auth.POST("", cl.UserController.Login)

	property := e.Group("property")
	property.POST("/complex", cl.PropertyController.AddComplex, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation([]string{"supervisor", "superadmin"}))
	property.DELETE("/complex/:id", cl.PropertyController.DeleteComplex, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation([]string{"supervisor", "superadmin"}))
	property.GET("/complex", cl.PropertyController.GetAllComplex, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation([]string{"supervisor", "superadmin"}))
	property.PUT("/complex/:id", cl.PropertyController.UpdateComplex, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation([]string{"supervisor", "superadmin"}))

	property.POST("/building/:id/complex", cl.PropertyController.AddBuilding, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation([]string{"supervisor", "superadmin"}))
	property.GET("/building", cl.PropertyController.GetAllBuilding)
	property.GET("/building/recommend", cl.PropertyController.GetRecommendedBuilding)
	property.DELETE("/building/:id", cl.PropertyController.DeleteBuilding, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation([]string{"supervisor", "superadmin"}))
	property.GET("/building/:id", cl.PropertyController.GetSingleBuilding)
	property.PUT("/building/:id", cl.PropertyController.UpdateBuilding, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation([]string{"supervisor", "superadmin"}))
}
