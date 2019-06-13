package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/swaggo/echo-swagger"
	"gopkg.in/gorp.v2"

	_ "tokyo2020sched/docs"
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
	dbmap.AddTableWithName(models.OlympicSchedule{}, "olympic_schedule").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.Handler{DbMap: dbmap}

	return handler
}

// @title TOKYO2020 schedule API
// @version 0.1
// @description TOKYO2020 schedule API

// @host tokyo2020sched.herokuapp.com
// @BasePath /
// @schemes https
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}

	port := os.Getenv("PORT")

	handler := initDb()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.GET("/ping", handler.Ping)

	// 競技
	g := e.Group("/competitions")
	g.GET("", handler.CompetitionList)
	g.GET("/:id", handler.CompetitionDetail)

	// 種別等
	g = e.Group("/classifications")
	g.GET("", handler.ClassificationList)
	g.GET("/:id", handler.ClassificationDetail)

	// 場所
	g = e.Group("/places")
	g.GET("", handler.PlaceList)
	g.GET("/:id", handler.PlaceDetail)

	// スケジュール
	g = e.Group("/schedules/olympic")
	g.GET("", handler.ScheduleList)
	g.GET("/:id", handler.ScheduleDetail)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":" + port))
}
