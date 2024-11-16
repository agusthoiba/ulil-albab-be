package connectors

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"

	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	//psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	os.Getenv("DB_SQL_HOST"), dbPort, os.Getenv("DB_SQL_USER"), os.Getenv("DB_SQL_PASSWORD"), os.Getenv("DB_SQL_NAME"))

	dbPort, err := strconv.Atoi(os.Getenv("DB_SQL_PORT"))
	if err != nil {
		panic(err)
	}

	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		os.Getenv("DB_SQL_USER"), os.Getenv("DB_SQL_PASSWORD"), os.Getenv("DB_SQL_HOST"), dbPort, os.Getenv("DB_SQL_NAME"))

	fmt.Println("psqlconn --", psqlconn)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to database")

	return db
}

// GetDB retrieves the database connection from the context
func GetDB(c echo.Context) *sql.DB {
	db, ok := c.Request().Context().Value("db").(*sql.DB)
	if !ok {
		return nil
	}
	return db
}
