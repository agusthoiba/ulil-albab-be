package main

import (
	"os"
	"fmt"
	"net/http"
	"ulil-albab-be/src/project/middlewares"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	if _, err := os.Stat(".env"); err == nil {
		godotenv.Load()
	} else {
		fmt.Printf("File .env does not exist use os.Environment instead\n");
	}
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	err := middlewares.NewMiddleware(e)

	if err != nil {
		e.Logger.Fatal("Error middleware")
	}

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":"+ port))
}
