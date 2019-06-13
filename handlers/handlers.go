package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/gorp.v2"
)

type Handler struct {
	DbMap *gorp.DbMap
}

// @Description 疎通を確認する
// @Tags Ping
// @Summary 疎通確認
// @Produce json
// @Success 200 {string} string
// @Router /ping [get]
func (handler *Handler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
