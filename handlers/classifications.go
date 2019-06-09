package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo"

	"tokyo2020-sch-api/holders"
)

var query_classification_list string = `
SELECT
    classification.id,
    classification.name,
    classification.created_at,
    classification.updated_at,
    competition.id AS competition_id,
    competition.name AS competition_name,
    competition.created_at AS competition_created_at,
    competition.updated_at AS competition_updated_at
FROM
    classification
    INNER JOIN
        competition
    ON  (
            classification.competition_id = competition.id
        )
`

var query_classification_detail string = `
SELECT
    classification.id,
    classification.name,
    classification.created_at,
    classification.updated_at,
    competition.id AS competition_id,
    competition.name AS competition_name,
    competition.created_at AS competition_created_at,
    competition.updated_at AS competition_updated_at
FROM
    classification
    INNER JOIN
        competition
    ON  (
            classification.competition_id = competition.id
        )
WHERE
    classification.id = $1
`

func (handler *Handler) ClassificationList(c echo.Context) error {
	var classifications []struct {
		Id                   int64
		Name                 string
		CreatedAt            time.Time `db:"created_at"`
		UpdatedAt            time.Time `db:"updated_at"`
		CompetitionId        int64     `db:"competition_id"`
		CompetitionName      string    `db:"competition_name"`
		CompetitionCreatedAt time.Time `db:"competition_created_at"`
		CompetitionUpdatedAt time.Time `db:"competition_updated_at"`
	}

	_, err := handler.DbMap.Select(&classifications, query_classification_list)
	if err != nil {
		// e.Logger.Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	result := make([]holders.ClassificationWithCompetition, len(classifications))
	for i, v := range classifications {
		result[i] = holders.ClassificationWithCompetition{
			Id:   v.Id,
			Name: v.Name,
			Competition: holders.Competition{
				Id:        v.CompetitionId,
				Name:      v.CompetitionName,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func (handler *Handler) ClassificationDetail(c echo.Context) error {
	id := c.Param("id")

	var classification struct {
		Id                   int64
		Name                 string
		CreatedAt            time.Time `db:"created_at"`
		UpdatedAt            time.Time `db:"updated_at"`
		CompetitionId        int64     `db:"competition_id"`
		CompetitionName      string    `db:"competition_name"`
		CompetitionCreatedAt time.Time `db:"competition_created_at"`
		CompetitionUpdatedAt time.Time `db:"competition_updated_at"`
	}
	err := handler.DbMap.SelectOne(&classification, query_classification_detail, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "Not Found")
		}

		// e.Logger.Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, holders.ClassificationWithCompetition{
		Id:   classification.Id,
		Name: classification.Name,
		Competition: holders.Competition{
			Id:        classification.CompetitionId,
			Name:      classification.CompetitionName,
			CreatedAt: classification.CreatedAt,
			UpdatedAt: classification.UpdatedAt,
		},
		CreatedAt: classification.CreatedAt,
		UpdatedAt: classification.UpdatedAt,
	})
}
