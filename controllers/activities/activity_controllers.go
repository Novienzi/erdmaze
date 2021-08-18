package Activities

import (
	"erdmaze/businesses/activities"
	"erdmaze/controllers/activities/request"
	"erdmaze/controllers/activities/response"
	"net/http"
	"strconv"

	controller "erdmaze/controllers"

	echo "github.com/labstack/echo/v4"
)

type ActivityController struct {
	activityUsecase activities.Usecase
}

func NewActivityController(cu activities.Usecase) *ActivityController {
	return &ActivityController{
		activityUsecase: cu,
	}
}

func (ctrl *ActivityController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.activityUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Activity{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *ActivityController) FindById(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	Activity, err := ctrl.activityUsecase.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, Activity)

}

func (ctrl *ActivityController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Activities{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.activityUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}
