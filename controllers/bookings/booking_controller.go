package bookings

import (
	"erdmaze/app/middleware"
	bookings "erdmaze/businesses/bookings"
	controller "erdmaze/controllers"
	"erdmaze/controllers/bookings/request"
	"erdmaze/controllers/bookings/response"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type BookingsController struct {
	bookingsUseCase bookings.Usecase
	JWTAuth         *middleware.ConfigJWT
}

func NewBookingsController(uc bookings.Usecase) *BookingsController {
	return &BookingsController{
		bookingsUseCase: uc,
	}
}

func (ctrl *BookingsController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := ctrl.JWTAuth.GetUser(c)
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

// func (ctrl *BookingsController) GetAll(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	resp, err := ctrl.bookingPackagesUseCase.GetByUserID(ctx, userID)

// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	responseController := []response.TourismPackages{}
// 	for _, value := range resp {
// 		responseController = append(responseController, response.FromDomain(value))
// 	}

// 	return controller.NewSuccessResponse(c, responseController)
// }

// func (ctrl *BookingsController) SelectAll(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	page := c.QueryParam("page")
// 	offset := c.QueryParam("limit")

// 	var varPage response.Pagination

// 	p, err := strconv.Atoi(page)
// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}

// 	o, err := strconv.Atoi(offset)
// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}

// 	resp, count, lastPage, err := ctrl.tourismPackagesUseCase.Fetch(ctx, p, o)
// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	varPage.CurrentPage = p
// 	varPage.LastPage = lastPage
// 	varPage.PerPage = o
// 	varPage.Total = count

// 	responseController := []response.TourismPackages{}
// 	for _, value := range resp {
// 		responseController = append(responseController, response.FromDomain(value))
// 	}

// 	return controller.NewSuccessResponseFetch(c, responseController, varPage)
// }

// func (ctrl *BookingsController) FindById(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}

// 	tourism, err := ctrl.tourismPackagesUseCase.GetByID(ctx, id)

// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}

// 	return controller.NewSuccessResponse(c, tourism)

// }
