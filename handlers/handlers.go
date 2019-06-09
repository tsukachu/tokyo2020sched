package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

type Handler struct {
	DbMap *gorp.DbMap
}

func (handler *Handler) Index(c echo.Context) error {
	return c.String(http.StatusOK, "hello, world")
}
