package middlewares

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	log "github.com/sirupsen/logrus"

	"database/sql"

	_ "github.com/lib/pq"

	"ulil-albab-be/src/project/connectors"
	"ulil-albab-be/src/project/handlers"
	"ulil-albab-be/src/project/models"
	"ulil-albab-be/src/project/repositories"
	"ulil-albab-be/src/project/services"
)

// DBMiddleware adds the database connection to the context
func DBMiddleware(db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), "db", db)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func NewMiddleware(e *echo.Echo) error {
	e.Use(middleware.Logger())

	for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
		if pair[0] != "DB_SQL_PASSWORD" { 
        	log.Info(pair[0],":",pair[1])
		}
    }


	log.SetLevel(log.TraceLevel)

	dbPort, err := strconv.Atoi(os.Getenv("DB_SQL_PORT"))
	if err != nil {
		panic(err)
	}

	optionDb := models.OptionDb{
		Host:     os.Getenv("DB_SQL_HOST"),
		Port:     dbPort,
		User:     os.Getenv("DB_SQL_USER"),
		Password: os.Getenv("DB_SQL_PASSWORD"),
		DbName:   os.Getenv("DB_SQL_NAME"),
	}

	db, err := connectors.InitDB(optionDb)

	if err != nil {
		fmt.Println(err)
		return err
	}

	e.Use(DBMiddleware(db))

	ayahRepo := repositories.NewAyah(db)
	surahRepo := repositories.NewSurah(db)

	service := services.NewService(surahRepo, ayahRepo)

	quranHandler := handlers.NewQuranHandler(service)

	e.GET("/quran/surah", quranHandler.GetSurah)
	e.GET("/quran/ayat/:suraId", quranHandler.GetAyats)
	e.GET("/quran/ayat", quranHandler.GetAllAyats)
	e.GET("/quran", quranHandler.GetAll)

	return nil
}
