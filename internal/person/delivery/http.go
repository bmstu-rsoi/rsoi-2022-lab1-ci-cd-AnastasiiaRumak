package delivery

import (
	"net/http"
	//"errors"
	"fmt"
	"context"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/bmstu-rsoi/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/models"
)

const (
	locationValueFormat = "/api/v1/persons/%d"
)

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) Configure(e *echo.Echo) {
	e.POST("/api/v1/persons", h.CreatePerson())

	e.DELETE("/api/v1/persons/:id", h.DeletePerson())
	e.PATCH("/api/v1/persons/:id", h.UpdatePerson())
	e.GET("/api/v1/persons/:id", h.GetPersonID())
	e.GET("/api/v1/persons", h.GetAll())
	
}


type request struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Age     int64    `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}

type response struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Age     int64    `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}

func fromModel(m models.Person) response {
	return response{	
		ID:      m.ID,
		Name:    m.Name,
		Age:     m.Age,
		Address: m.Address,
		Work:    m.Work,
	}
}

func toModel(req request) models.Person {
	return models.Person{
		ID:      req.ID,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
		Work:    req.Work,
	}
}

func (h *Handler) CreatePerson() echo.HandlerFunc {

	return func(ctx echo.Context) error {
		//var req request//**
		var req request//**

		if err := ctx.Bind(&req); err != nil {
			//return err
			return ctx.JSON(http.StatusBadRequest, err)
		}
		id, err := h.usecase.CreatePerson(context.Background(), toModel(req))
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}		
		locationValue := fmt.Sprintf(locationValueFormat,id)
		ctx.Response().Header().Set("Location", locationValue)
//
		ctx.Response().Header().Set("Content-Type", "application/json")
//			
//fmt.
		return ctx.JSON(http.StatusCreated, nil)
	}
}

func (h *Handler) DeletePerson() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		//strID := ctx.Param("id")
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)		
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, nil)
		}

		err = h.usecase.DeletePerson(context.Background(), id)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		//locationValue := fmt.Sprintf(locationValueFormat,id)
		//ctx.Response().Header().Set("Location", locationValue)
//
		ctx.Response().Header().Set("Content-Type", "application/json")
//
		return ctx.JSON(http.StatusNoContent, nil)
	}
}

func (h *Handler) UpdatePerson ()  echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req request//**
		//req := &request{}
		if err := ctx.Bind(req); err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			
			return ctx.JSON(http.StatusBadRequest, nil)
		}
		req.ID = id

		update, err := h.usecase.UpdatePerson(context.Background(), toModel(req))
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
//
		ctx.Response().Header().Set("Content-Type", "application/json")
//		
		return ctx.JSON(http.StatusOK, fromModel(update))
	}
}

func (h *Handler) GetPersonID () echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, nil)
		}

		model, err := h.usecase.GetPersonID(context.Background(), id)
		if err != nil {

			return ctx.JSON(http.StatusInternalServerError, err)
		}
//
		ctx.Response().Header().Set("Content-Type", "application/json")
//
		return ctx.JSON(http.StatusOK, fromModel(model))
	}
}

func (h *Handler) GetAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		persons, err := h.usecase.GetAll(context.Background())
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		var response []response
		for _, p := range *persons {
			response = append(response, fromModel(p))
		}
//
		ctx.Response().Header().Set("Content-Type", "application/json")
//
		return ctx.JSON(http.StatusOK, response)
	}
}


/*
func (h *Handler) GetPersonID() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req request
		if err := ctx.Bind(req); err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	id, err := h.usecase.GetPersonID(ctx, toModel(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)

	}
	locationValue := fmt.Sprintf(locationValueFormat, id)
	ctx.Response().Header().Get("Location", locationValue)

	return ctx.JSON(http.StatusCreated, nil)
	}
}
*/
