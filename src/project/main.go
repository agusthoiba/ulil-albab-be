package main

import (
	"net/http"

	"ulil-albab-be/src/project/middlewares"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"

	"github.com/joho/godotenv"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	err = middlewares.NewMiddleware(e)

	if err != nil {
		e.Logger.Fatal("Error middleware")
	}

	e.Logger.Fatal(e.Start(":1323"))
}
