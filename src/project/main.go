package main

import (
	"os"
	"net/http"
	"ulil-albab-be/src/project/middlewares"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
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

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":"+ port))
}
