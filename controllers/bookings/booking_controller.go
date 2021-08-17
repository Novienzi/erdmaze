package bookings

import (
	"erdmaze/app/middleware"
	bookings "erdmaze/businesses/bookings"
	controller "erdmaze/controllers"
	"erdmaze/controllers/bookings/request"
	"erdmaze/controllers/bookings/response"
	"errors"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type BookingsController struct {
	bookingsUseCase bookings.Usecase
	jwtAuth         *middleware.ConfigJWT
}

func NewBookingsController(uc bookings.Usecase) *BookingsController {
	return &BookingsController{
		bookingsUseCase: uc,
	}
}

func (ctrl *BookingsController) GetByUserID(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := strconv.Atoi(c.Param("user_id"))

	resp, err := ctrl.bookingsUseCase.GetByUserID(ctx, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Bookings{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *BookingsController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	booking, err := ctrl.bookingsUseCase.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, booking)

}

func (ctrl *BookingsController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := ctrl.jwtAuth.GetUser(c)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	req := request.Bookings{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.UserID = user.ID

	resp, err := ctrl.bookingsUseCase.Store(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *BookingsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Bookings{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	resp, err := ctrl.bookingsUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}
