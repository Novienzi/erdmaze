package locations

import (
	"erdmaze/businesses/locations"
	"erdmaze/controllers/locations/request"
	"erdmaze/controllers/locations/response"
	"net/http"
	"strconv"

	controller "erdmaze/controllers"

	echo "github.com/labstack/echo/v4"
)

type LocationController struct {
	locationUsecase locations.Usecase
}

func NewLocationController(cu locations.Usecase) *LocationController {
	return &LocationController{
		locationUsecase: cu,
	}
}

func (ctrl *LocationController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.locationUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Location{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *LocationController) FindById(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	Activity, err := ctrl.locationUsecase.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, Activity)

}

func (ctrl *LocationController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Locations{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.locationUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}
