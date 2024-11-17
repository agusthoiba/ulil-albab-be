package main

import (
	"errors"
	"fmt"
	"net/http"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"

	"github.com/joho/godotenv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

type Surah struct {
	number        int
	numberOfAyahs int
	name          string
	translation   string
	revelation    string
	description   string
	audio         string
}

type SurahResp struct {
	Number        int    `json:"number"`
	NumberOfAyahs int    `json:"numberOfAyahs"`
	Name          string `json:"name"`
	Translation   string `json:"translation"`
	Revelation    string `json:"revelation"`
	Description   string `json:"description"`
	Audio         string `json:"audio"`
}

func getSurah(c echo.Context) error {
	// c.QueryParam("team")
	connStr := "postgres://postgres:p0stgr3@localhost/local_ulilalbab?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	fmt.Println("db", db)
	if err != nil {
		errors.New("error")
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM surah")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var surahs []SurahResp
	for rows.Next() {
		var surah SurahResp
		if err := rows.Scan(&surah.Number, &surah.NumberOfAyahs, &surah.Name,
			&surah.Translation, &surah.Revelation, &surah.Description, &surah.Audio); err != nil {
			panic(err)
		}

		surahs = append(surahs, surah)
		//fmt.Printf("Number: %d, Name: %s\n", surah.number, surah.name)
	}

	//fmt.Println(surahs)
	return c.JSON(http.StatusOK, surahs)
}

func mainMain() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	err := godotenv.Load(".env")
	if err != nil {
		e.Logger.Fatalf("Error loading .env file: %s", err)
	}

	e.GET("/surahs", getSurah)
	e.Logger.Fatal(e.Start(":1323"))
}
