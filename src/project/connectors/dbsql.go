package connectors

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"database/sql"

	_ "github.com/lib/pq"

	"ulil-albab-be/src/project/models"
)

/* var (
	SqlOpen = sql.Open
)*/

type (
	sqlOpener func(string, string) (*sql.DB, error)
)

func InitDB(option models.OptionDb, sqlOpen sqlOpener) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		option.User, option.Password, option.Host, option.Port, option.DbName)

	db, err := sqlOpen("postgres", psqlconn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		panic("Error connected to database")
	}

	fmt.Println("Successfully connected to database")

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
