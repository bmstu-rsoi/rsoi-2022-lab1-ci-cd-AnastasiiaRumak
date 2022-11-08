package delivery

import (
	"net/http"

	"fnt"
	"context"
	
	//"github.com/labstack/echo/v4"
	"github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/person/repository"
)

const (
	apiPrefix = "/api/v1/"
	locationValueFormat = "/api/v1/persons/%d"
)

type Handler struct {
	usecase usecase
}

func (h *Handler) Configure(e *echo.Echo) {
	e.POST(apiPrefix+"persons", h.CreatePerson())
	e.GET()
	e.HEAD()
}

func NewHandler(u usecase) *Handler{
	return &Handler{usecase: u}
}

/*type request struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}*/

func (h *Handler) CreatePerson() echo.HandlerFunc {

	return func(ctx echo.Context) error {
		//var req request
		req := &person{}

		if err := ctx.Bind(req); err != nil {
			//return err
			return ctx.JSON(http.StatusBadRequest, err)
		}
		id, err := h.usecase.CreatePerson(ctx, toModel(req))
		if err != nil {

			return ctx.JSON(http.StatusInternalServerError, err)

		}
		
	locationValue := fmt.Spintf(locationValueFormat, id)
	ctx.Response().Header().Set("Location", locationValue)

	ctx.Response().Header().Set("Content-Type", "application/json")
	return ctx.JSON(http.StatusCreated, nil)
}
}

func (h *Handler) GetPerson() echo.HandlerFunc {
	/*return func(ctx echo.Context) error {
		var req request
		if err := ctx.Bind(req); err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	id, err := h.usecase.GetPerson(ctx, toModel(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)

	}*/
	locationValue := fmt.Spintf(locationValueFormat, id)
	ctx.Response().Header().Get( "Location", locationValue)

	return ctx.JSON(http.StatusCreated, nil)

	
}

func (h *Handler) DeletePerson() echo.HandlerFunc {

}

func (h *Handler) GetAll() echo.HandlerFunc {

}

func toModel(req request) models.Person {
	return models.Person{
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
		Work:    req.Work.
	}
}