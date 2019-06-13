package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/guregu/null.v3/zero"

	"tokyo2020sched/models"
)

type OlympicScheduleForScan struct {
	Id                          int64
	CompetitionId               int64       `db:"competition_id"`
	CompetitionName             string      `db:"competition_name"`
	CompetitionCreatedAt        time.Time   `db:"competition_created_at"`
	CompetitionUpdatedAt        time.Time   `db:"competition_updated_at"`
	ClassificationId            zero.Int    `db:"classification_id"`
	ClassificationName          zero.String `db:"classification_name"`
	ClassificationCompetitionId zero.Int    `db:"classification_competition_id"`
	ClassificationCreatedAt     zero.Time   `db:"classification_created_at"`
	ClassificationUpdatedAt     zero.Time   `db:"classification_updated_at"`
	Title                       string
	Begin                       time.Time
	End                         time.Time
	PlaceId                     int64     `db:"place_id"`
	PlaceName                   string    `db:"place_name"`
	PlaceCreatedAt              time.Time `db:"place_created_at"`
	PlaceUpdatedAt              time.Time `db:"place_updated_at"`
	Content                     null.String
	CreatedAt                   time.Time `db:"created_at"`
	UpdatedAt                   time.Time `db:"updated_at"`
}

var query_schedule_list string = `
SELECT
    olympic_schedule.id,
    competition.id AS competition_id,
    competition.name AS competition_name,
    competition.created_at AS competition_created_at,
    competition.updated_at AS competition_updated_at,
    classification.id AS classification_id,
    classification.name AS classification_name,
    classification.competition_id AS classification_competition_id,
    classification.created_at AS classification_created_at,
    classification.updated_at AS classification_updated_at,
    olympic_schedule.title,
    olympic_schedule.begin,
    olympic_schedule.end,
    place.id AS place_id,
    place.name AS place_name,
    place.created_at AS place_created_at,
    place.updated_at AS place_updated_at,
    olympic_schedule.content,
    olympic_schedule.created_at,
    olympic_schedule.updated_at
FROM
    olympic_schedule
    INNER JOIN
        competition
    ON  (
            olympic_schedule.competition_id = competition.id
        )
    LEFT JOIN
        classification
    ON  (
            olympic_schedule.classification_id = classification.id
        )
    INNER JOIN
        place
    ON  (
            olympic_schedule.place_id = place.id
        )
`

var query_schedule_detail string = `
SELECT
    olympic_schedule.id,
    competition.id AS competition_id,
    competition.name AS competition_name,
    competition.created_at AS competition_created_at,
    competition.updated_at AS competition_updated_at,
    classification.id AS classification_id,
    classification.name AS classification_name,
    classification.competition_id AS classification_competition_id,
    classification.created_at AS classification_created_at,
    classification.updated_at AS classification_updated_at,
    olympic_schedule.title,
    olympic_schedule.begin,
    olympic_schedule.end,
    place.id AS place_id,
    place.name AS place_name,
    place.created_at AS place_created_at,
    place.updated_at AS place_updated_at,
    olympic_schedule.content,
    olympic_schedule.created_at,
    olympic_schedule.updated_at
FROM
    olympic_schedule
    INNER JOIN
        competition
    ON  (
            olympic_schedule.competition_id = competition.id
        )
    LEFT JOIN
        classification
    ON  (
            olympic_schedule.classification_id = classification.id
        )
    INNER JOIN
        place
    ON  (
            olympic_schedule.place_id = place.id
        )
WHERE
    olympic_schedule.id = $1
`

// @Description スケジュールの一覧を取得する
// @Tags OlympicSchedules
// @Summary スケジュール一覧を取得
// @Produce json
// @Success 200 {array} models.OlympicSchedule
// @Router /schedules/olympic [get]
func (handler *Handler) ScheduleList(c echo.Context) error {
	var schedules []OlympicScheduleForScan

	_, err := handler.DbMap.Select(&schedules, query_schedule_list)
	if err != nil {
		// e.Logger.Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	result := make([]models.OlympicSchedule, len(schedules))
	for i, v := range schedules {
		result[i] = models.OlympicSchedule{
			Id: v.Id,
			Competition: models.Competition{
				Id:        v.CompetitionId,
				Name:      v.CompetitionName,
				CreatedAt: v.CompetitionCreatedAt,
				UpdatedAt: v.CompetitionUpdatedAt,
			},
			Classification: models.Classification{
				Id:            v.ClassificationId.Int64,
				Name:          v.ClassificationName.String,
				CompetitionId: v.ClassificationCompetitionId.Int64,
				CreatedAt:     v.ClassificationCreatedAt.Time,
				UpdatedAt:     v.ClassificationUpdatedAt.Time,
			},
			Title: v.Title,
			Begin: v.Begin,
			End:   v.End,
			Place: models.Place{
				Id:        v.PlaceId,
				Name:      v.PlaceName,
				CreatedAt: v.PlaceCreatedAt,
				UpdatedAt: v.PlaceUpdatedAt,
			},
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, result)
}

// @Description 指定したスケジュールの詳細情報を取得する
// @Tags OlympicSchedules
// @Summary スケジュールの詳細情報を取得
// @Produce json
// @Param id path integer true "スケジュールID"
// @Success 200 {object} models.OlympicSchedule
// @Router /schedules/olympic/{id} [get]
func (handler *Handler) ScheduleDetail(c echo.Context) error {
	id := c.Param("id")

	var schedule OlympicScheduleForScan
	err := handler.DbMap.SelectOne(&schedule, query_schedule_detail, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "Not Found")
		}

		// e.Logger.Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.OlympicSchedule{
		Id: schedule.Id,
		Competition: models.Competition{
			Id:        schedule.CompetitionId,
			Name:      schedule.CompetitionName,
			CreatedAt: schedule.CompetitionCreatedAt,
			UpdatedAt: schedule.CompetitionUpdatedAt,
		},
		Classification: models.Classification{
			Id:            schedule.ClassificationId.Int64,
			Name:          schedule.ClassificationName.String,
			CompetitionId: schedule.ClassificationCompetitionId.Int64,
			CreatedAt:     schedule.ClassificationCreatedAt.Time,
			UpdatedAt:     schedule.ClassificationUpdatedAt.Time,
		},
		Title: schedule.Title,
		Begin: schedule.Begin,
		End:   schedule.End,
		Place: models.Place{
			Id:        schedule.PlaceId,
			Name:      schedule.PlaceName,
			CreatedAt: schedule.PlaceCreatedAt,
			UpdatedAt: schedule.PlaceUpdatedAt,
		},
		Content:   schedule.Content,
		CreatedAt: schedule.CreatedAt,
		UpdatedAt: schedule.UpdatedAt,
	})
}
