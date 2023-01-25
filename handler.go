package main

import (
	"net/http"

	"github.com/chinathaip/catfact/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	s service.Service
}

func NewRouter(s service.Service) *Handler {
	return &Handler{s: s}
}

func (r *Handler) RegisterRouter() *echo.Echo {
	e := echo.New()

	e.GET("/", r.getCatFact)
	return e
}

func (r *Handler) getCatFact(c echo.Context) error {
	fact, err := r.s.GetCatFact(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, fact)
}
