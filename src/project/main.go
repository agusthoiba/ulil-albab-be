package main

import (
	"context"
	"net/http"

	"ulil-albab-be/src/project/connectors"
	"ulil-albab-be/src/project/handlers"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"

	"github.com/joho/godotenv"
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

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	db := connectors.InitDB()
	//e.Use(middleware.Logger())

	e.Use(DBMiddleware(db))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/quran/surah", handlers.GetSurah)
	e.GET("/quran/ayat/:suraId", handlers.GetAyats)
	e.Logger.Fatal(e.Start(":1323"))
}
