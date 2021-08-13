package routes

import (
	"erdmaze/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.POST("/login", cl.UserController.Login)
	users.GET("/:id", cl.UserController.GetUserDetail, middleware.JWTWithConfig(cl.JWTMiddleware))

}
