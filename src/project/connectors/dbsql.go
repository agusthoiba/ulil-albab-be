package connectors

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"database/sql"

	_ "github.com/lib/pq"

	"ulil-albab-be/src/project/models"
)

func InitDB(option models.OptionDb) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		option.User, option.Password, option.Host, option.Port, option.DbName)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		panic("Error connected to database")
	}

	log.Println("Successfully connected to database")

	return db, nil
}

// GetDB retrieves the database connection from the context
func GetDB(c echo.Context) *sql.DB {
	db, ok := c.Request().Context().Value("db").(*sql.DB)
	if !ok {
		return nil
	}
	return db
}
