package routes

import (
	Activities "erdmaze/controllers/activities"
	"erdmaze/controllers/locations"
	"erdmaze/controllers/tourism_packages"
	"erdmaze/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware             middleware.JWTConfig
	UserController            users.UserController
	ActivityController        Activities.ActivityController
	LocationController        locations.LocationController
	TourismPackagesController tourism_packages.TourismPackagesController
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
	activity.GET("/:id", cl.ActivityController.FindById)
	activity.POST("", cl.ActivityController.Store)
	activity.PUT("/:id", cl.ActivityController.Update)
	activity.DELETE("/:id", cl.ActivityController.Delete)

	//locations...
	locations := e.Group("v1/api/locations")
	locations.Use(middleware.JWTWithConfig(cl.JWTMiddleware))

	locations.GET("", cl.LocationController.GetAll)
	locations.GET("/:id", cl.LocationController.FindById)
	locations.POST("", cl.LocationController.Store)
	locations.PUT("/:id", cl.LocationController.Update)
	locations.DELETE("/:id", cl.LocationController.Delete)

	//tourism_packages...
	tourismPackages := e.Group("v1/api/tourism")
	tourismPackages.Use(middleware.JWTWithConfig(cl.JWTMiddleware))

	tourismPackages.GET("/pagination", cl.TourismPackagesController.SelectAll)
	tourismPackages.GET("", cl.TourismPackagesController.GetAll)
	tourismPackages.GET("/:id", cl.TourismPackagesController.FindById)
	tourismPackages.POST("", cl.TourismPackagesController.Store)

}
