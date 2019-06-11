package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v2"

	"tokyo2020sched/handlers"
	"tokyo2020sched/models"
)

func initDb() handlers.Handler {
	dsn := os.Getenv("DSN")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(models.Competition{}, "competition").SetKeys(true, "Id")
	dbmap.AddTableWithName(models.Classification{}, "classification").SetKeys(true, "Id")
	dbmap.AddTableWithName(models.Place{}, "place").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.Handler{DbMap: dbmap}

	return handler
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}

	port := os.Getenv("PORT")

	handler := initDb()

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

	// 競技
	g := e.Group("/competitions")
	g.GET("/", handler.CompetitionList)
	g.GET("/:id/", handler.CompetitionDetail)

	// 種別等
	g = e.Group("/classifications")
	g.GET("/", handler.ClassificationList)
	g.GET("/:id/", handler.ClassificationDetail)

	// 場所
	g = e.Group("/places")
	g.GET("/", handler.PlaceList)
	g.GET("/:id/", handler.PlaceDetail)

	e.Logger.Fatal(e.Start(":" + port))
}
