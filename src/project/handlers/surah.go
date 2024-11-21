package handlers

import (
	"net/http"
	"strconv"
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

func GetAyats(c echo.Context) error {
	surId, err := strconv.Atoi(c.Param("suraId"))
	if err != nil {
		return err
	}

	surahRepo, err := repositories.GetAyatBySuratId(c, surId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, surahRepo)
}

func GetAllAyats(c echo.Context) error {
	surahRepo, err := repositories.GetAllAyat(c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, surahRepo)
}
