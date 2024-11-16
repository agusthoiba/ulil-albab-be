package handlers

import (
	"net/http"
	"ulil-albab-be/src/project/repositories"

	"github.com/labstack/echo/v4"
)

func GetSurah(c echo.Context) error {
	surahRepo, err := repositories.GetSurahList(c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, surahRepo)
}
