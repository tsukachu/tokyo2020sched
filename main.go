package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"

	"tokyo2020-sch-api/handlers"
	"tokyo2020-sch-api/holders"
)

func initDb() handlers.Handler {
	dsn := os.Getenv("DSN")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(holders.Competition{}, "competitions").SetKeys(true, "Id")
	dbmap.AddTableWithName(holders.Classification{}, "classifications").SetKeys(true, "Id")

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

	e.GET("/", handler.Index)

	// 競技
	g := e.Group("/competitions")
	g.GET("/", handler.CompetitionList)
	g.GET("/:id/", handler.CompetitionDetail)

	// 種別等
	g = e.Group("/classifications")
	g.GET("/", handler.ClassificationList)
	g.GET("/:id/", handler.ClassificationDetail)

	e.Logger.Fatal(e.Start(":" + port))
}
