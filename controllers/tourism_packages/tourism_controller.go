package tourism_packages

import (
	tourismpackages "erdmaze/businesses/tourism_packages"
	controller "erdmaze/controllers"
	"erdmaze/controllers/tourism_packages/request"
	"erdmaze/controllers/tourism_packages/response"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type TourismPackagesController struct {
	tourismPackagesUseCase tourismpackages.Usecase
}

func NewTourismPackagesController(tourismUC tourismpackages.Usecase) *TourismPackagesController {
	return &TourismPackagesController{
		tourismPackagesUseCase: tourismUC,
	}
}

func (ctrl *TourismPackagesController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.TourismPackages{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.tourismPackagesUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *TourismPackagesController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	tourismName := c.QueryParam("name")
	locationName := c.QueryParam("location")
	activityName := c.QueryParam("activity")

	resp, err := ctrl.tourismPackagesUseCase.GetAll(ctx, tourismName, locationName, activityName)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.TourismPackages{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *TourismPackagesController) SelectAll(c echo.Context) error {
	ctx := c.Request().Context()
	page := c.QueryParam("page")
	offset := c.QueryParam("limit")

	var varPage response.Pagination

	p, err := strconv.Atoi(page)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	o, err := strconv.Atoi(offset)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, count, lastPage, err := ctrl.tourismPackagesUseCase.Fetch(ctx, p, o)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	varPage.CurrentPage = p
	varPage.LastPage = lastPage
	varPage.PerPage = o
	varPage.Total = count

	responseController := []response.TourismPackages{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponseFetch(c, responseController, varPage)
}

func (ctrl *TourismPackagesController) FindById(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	tourism, err := ctrl.tourismPackagesUseCase.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, tourism)

}
