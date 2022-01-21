package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_middleware "github.com/sewakantor/sw-be/app/middleware"
	"github.com/sewakantor/sw-be/controllers/complex"
	"github.com/sewakantor/sw-be/controllers/users"
)

type ControllerList struct {
	UserController    users.UserController
	JWTMiddleware     middleware.JWTConfig
	ComplexController complex.ComplexControllers
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", cl.UserController.Registration)
	users.POST("/activation/:id", cl.UserController.Activation)

	auth := e.Group("auth")
	auth.POST("", cl.UserController.Login)

	complex := e.Group("complex")
	complex.POST("", cl.ComplexController.AddComplex, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation([]string{"supervisor", "superadmin"}))
	complex.DELETE("/:id", cl.ComplexController.DeleteComplex, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation([]string{"supervisor", "superadmin"}))
}
