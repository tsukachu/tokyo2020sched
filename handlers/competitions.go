package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"tokyo2020sched/models"
)

var query_competition_list string = `
SELECT
    *
FROM
    competition
`

var query_competition_detail string = `
SELECT
    *
FROM
    competition
WHERE
    id = $1
`

// @Description 競技の一覧を取得する
// @Tags Competitions
// @Summary 競技一覧を取得
// @Produce json
// @Success 200 {array} models.Competition
// @Router /competitions [get]
func (handler *Handler) CompetitionList(c echo.Context) error {
	var competitions []models.Competition

	_, err := handler.DbMap.Select(&competitions, query_competition_list)
	if err != nil {
		// e.Logger.Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, competitions)
}

// @Description 指定した競技の詳細情報を取得する
// @Tags Competitions
// @Summary 競技の詳細情報を取得
// @Produce json
// @Param id path integer true "競技ID"
// @Success 200 {object} models.Competition
// @Router /competitions/{id} [get]
func (handler *Handler) CompetitionDetail(c echo.Context) error {
	id := c.Param("id")

	var competition models.Competition
	err := handler.DbMap.SelectOne(&competition, query_competition_detail, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "Not Found")
		}

		// e.Logger.Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, competition)
}
