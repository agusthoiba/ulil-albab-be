package handlers

import (
	"net/http"
	"strconv"
	"ulil-albab-be/src/project/services"

	"github.com/labstack/echo/v4"
)

type QuranHandler struct {
	service services.ServiceInt
}

func NewQuranHandler(service services.ServiceInt) *QuranHandler {
	return &QuranHandler{service: service}
}

func (qh *QuranHandler) GetSurah(c echo.Context) error {
	surahRepo, err := qh.service.GetSurah()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, surahRepo)
}

func (qh *QuranHandler) GetAyats(c echo.Context) error {
	surId, err := strconv.Atoi(c.Param("suraId"))
	if err != nil {
		return err
	}

	surahRepo, err := qh.service.GetAyatBySuratId(surId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, surahRepo)
}

func (qh *QuranHandler) GetAllAyats(c echo.Context) error {
	surahRepo, err := qh.service.GetAllAyat()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, surahRepo)
}

func (qh *QuranHandler) GetAll(c echo.Context) error {
	quranData, err := qh.service.GetAll()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, quranData)
}
