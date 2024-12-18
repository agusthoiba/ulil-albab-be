package handlers

import (
	"net/http"
	"strconv"
	"ulil-albab-be/src/project/repositories"

	"github.com/labstack/echo/v4"
)

type QuranHandler struct {
	SurahRepository *repositories.SurahRepository
	AyahRepository  *repositories.AyahRepository
}

func NewQuranHandler(surahRepo *repositories.SurahRepository, ayahRepo *repositories.AyahRepository) *QuranHandler {
	return &QuranHandler{
		SurahRepository: surahRepo,
		AyahRepository:  ayahRepo,
	}
}

func (qh *QuranHandler) GetSurah(c echo.Context) error {
	surahRepo, err := qh.SurahRepository.GetSurahList()

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

	surahRepo, err := qh.AyahRepository.GetAyatBySuratId(surId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, surahRepo)
}

func (qh *QuranHandler) GetAllAyats(c echo.Context) error {
	surahRepo, err := qh.AyahRepository.GetAllAyat()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, surahRepo)
}
