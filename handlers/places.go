package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"tokyo2020sched/models"
)

var query_place_list string = `
SELECT
    *
FROM
    place
`

var query_place_detail string = `
SELECT
    *
FROM
    place
WHERE
    id = $1
`

// @Description 場所の一覧を取得する
// @Tags Places
// @Summary 場所一覧を取得
// @Produce json
// @Success 200 {array} models.Place
// @Router /places [get]
func (handler *Handler) PlaceList(c echo.Context) error {
	var places []models.Place

	_, err := handler.DbMap.Select(&places, query_place_list)
	if err != nil {
		// e.Logger.Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, places)
}

// @Description 指定した場所の詳細情報を取得する
// @Tags Places
// @Summary 場所の詳細情報を取得
// @Produce json
// @Param id path integer true "場所ID"
// @Success 200 {object} models.Place
// @Router /places/{id} [get]
func (handler *Handler) PlaceDetail(c echo.Context) error {
	id := c.Param("id")

	var place models.Place
	err := handler.DbMap.SelectOne(&place, query_place_detail, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "Not Found")
		}

		// e.Logger.Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, place)
}
