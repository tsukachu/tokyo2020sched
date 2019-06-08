package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

type Competitions struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func initDb() *gorp.DbMap {
	dsn := os.Getenv("DSN")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(Competitions{}, "competitions").SetKeys(true, "Id")

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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, world")
	})

	e.GET("/competitions/", func(c echo.Context) error {
		var competitions []Competitions

		_, err := dbmap.Select(&competitions, "SELECT * FROM competitions")
		if err != nil {
			e.Logger.Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, competitions)
	})

	e.GET("/competitions/:id/", func(c echo.Context) error {
		id := c.Param("id")

		var competition Competitions
		err := dbmap.SelectOne(&competition, "SELECT * FROM competitions WHERE id = $1", id)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.String(http.StatusNotFound, "Not Found")
			}

			e.Logger.Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, competition)
	})

	e.Logger.Fatal(e.Start(":" + port))
}
