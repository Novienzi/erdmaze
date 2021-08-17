package main

import (
	_userUsecase "erdmaze/businesses/users"
	_userController "erdmaze/controllers/users"
	_userRepo "erdmaze/drivers/databases/users"

	_activityUsecase "erdmaze/businesses/activities"
	_activityController "erdmaze/controllers/activities"
	_activityRepo "erdmaze/drivers/databases/activities"

	_locationUsecase "erdmaze/businesses/locations"
	_locationController "erdmaze/controllers/locations"
	_locationRepo "erdmaze/drivers/databases/locations"

	_tourismUsecase "erdmaze/businesses/tourism_packages"
	_tourismController "erdmaze/controllers/tourism_packages"
	_tourismRepo "erdmaze/drivers/databases/tourism_packages"

	_bookingUsecase "erdmaze/businesses/bookings"
	_bookingController "erdmaze/controllers/bookings"
	_bookingRepo "erdmaze/drivers/databases/bookings"

	_dbDriver "erdmaze/drivers/mysql"

	_middleware "erdmaze/app/middleware"
	_routes "erdmaze/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	userRepo := _userRepo.NewMySQLUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	activityRepo := _activityRepo.NewMySQLActivityRepository(db)
	activityUseCase := _activityUsecase.NewActivityUsecase(timeoutContext, activityRepo)
	activityCtrl := _activityController.NewActivityController(activityUseCase)

	locationRepo := _locationRepo.NewMySQLLocationRepository(db)
	locationUsacase := _locationUsecase.NewLocationUsecase(timeoutContext, locationRepo)
	locationCtrl := _locationController.NewLocationController(locationUsacase)

	tourismRepo := _tourismRepo.NewMySQLTourismPackagesRepository(db)
	tourismUsacase := _tourismUsecase.NewTourismPackagesUsecase(tourismRepo, activityUseCase, locationUsacase, timeoutContext)
	tourismCtrl := _tourismController.NewTourismPackagesController(tourismUsacase)

	bookingRepo := _bookingRepo.NewMySQLBookingsRepository(db)
	bookingUsacase := _bookingUsecase.NewBookingsUsecase(bookingRepo, timeoutContext)
	bookingCtrl := _bookingController.NewBookingsController(bookingUsacase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:             configJWT.Init(),
		UserController:            *userCtrl,
		ActivityController:        *activityCtrl,
		LocationController:        *locationCtrl,
		TourismPackagesController: *tourismCtrl,
		BookingsController:        *bookingCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
