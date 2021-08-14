package routes

import (
	Activities "erdmaze/controllers/activities"
	"erdmaze/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
	ActivityController Activities.ActivityController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("v1/api/users")
	users.POST("/register", cl.UserController.Store)
	users.POST("/login", cl.UserController.Login)
	users.GET("/:id", cl.UserController.GetUserDetail, middleware.JWTWithConfig(cl.JWTMiddleware))

	//activities...
	activity := e.Group("v1/api/activities")
	activity.Use(middleware.JWTWithConfig(cl.JWTMiddleware))

	activity.GET("", cl.ActivityController.GetAll)
	activity.GET("/id/:id", cl.ActivityController.FindById)
	activity.POST("", cl.ActivityController.Store)
	activity.PUT("/id/:id", cl.ActivityController.Update)
	activity.DELETE("/id/:id", cl.ActivityController.Delete)

}
