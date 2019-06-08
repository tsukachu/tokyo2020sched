package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

type Competition struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Classification struct {
	Id            int64     `db:"id"`
	Name          string    `db:"name"`
	CompetitionId int64     `db:"competition_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type JoinedClassification struct {
	Id                   int64
	Name                 string
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
	CompetitionId        int64     `db:"co_id"`
	CompetitionName      string    `db:"co_name"`
	CompetitionCreatedAt time.Time `db:"co_created_at"`
	CompetitionUpdatedAt time.Time `db:"co_updated_at"`
}

type ClassificationWithCompetition struct {
	Id          int64
	Name        string
	Competition Competition
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func initDb() *gorp.DbMap {
	dsn := os.Getenv("DSN")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(Competition{}, "competitions").SetKeys(true, "Id")
	dbmap.AddTableWithName(Classification{}, "classifications").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}

	return dbmap
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}

	port := os.Getenv("PORT")

	dbmap := initDb()

	e := echo.New()

	e.Use(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, world")
	})

	e.GET("/competitions/", func(c echo.Context) error {
		var competitions []Competition

		_, err := dbmap.Select(&competitions, "SELECT * FROM competition")
		if err != nil {
			e.Logger.Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, competitions)
	})

	e.GET("/competitions/:id/", func(c echo.Context) error {
		id := c.Param("id")

		var competition Competition
		err := dbmap.SelectOne(&competition, "SELECT * FROM competition WHERE id = $1", id)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.String(http.StatusNotFound, "Not Found")
			}

			e.Logger.Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, competition)
	})

	e.GET("/classifications/", func(c echo.Context) error {
		var classifications []JoinedClassification

		_, err := dbmap.Select(&classifications, "SELECT cl.id, cl.name, cl.created_at, cl.updated_at, co.id AS co_id, co.name AS co_name, co.created_at AS co_created_at, co.updated_at AS co_updated_at FROM classification AS cl INNER JOIN competition AS co ON (cl.competition_id = co.id)")
		if err != nil {
			e.Logger.Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}

		result := make([]ClassificationWithCompetition, len(classifications))
		for i, v := range classifications {
			result[i] = ClassificationWithCompetition{
				Id:   v.Id,
				Name: v.Name,
				Competition: Competition{
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
	})

	e.GET("/classifications/:id/", func(c echo.Context) error {
		id := c.Param("id")

		var classification JoinedClassification
		err := dbmap.SelectOne(&classification, "SELECT cl.id, cl.name, cl.created_at, cl.updated_at, co.id AS co_id, co.name AS co_name, co.created_at AS co_created_at, co.updated_at AS co_updated_at FROM classification AS cl INNER JOIN competition AS co ON (cl.competition_id = co.id) WHERE cl.id = $1", id)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.String(http.StatusNotFound, "Not Found")
			}

			e.Logger.Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, ClassificationWithCompetition{
			Id:   classification.Id,
			Name: classification.Name,
			Competition: Competition{
				Id:        classification.CompetitionId,
				Name:      classification.CompetitionName,
				CreatedAt: classification.CreatedAt,
				UpdatedAt: classification.UpdatedAt,
			},
			CreatedAt: classification.CreatedAt,
			UpdatedAt: classification.UpdatedAt,
		})
	})

	e.Logger.Fatal(e.Start(":" + port))
}
