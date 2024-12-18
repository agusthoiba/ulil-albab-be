package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"ulil-albab-be/src/project/models"
	"ulil-albab-be/src/project/repositories"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"go.uber.org/mock/gomock"
)

var (
	surah1 = &models.SurahResp{
		Number:        1,
		NumberOfAyahs: 114,
		Name:          "Al-Fatihah",
		Translation:   "الفاتحة",
		Revelation:    "Mekah",
		Description:   "The first chapter of the Quran",
		Audio:         "al-fatihah.mp3",
	}

	surah2 = &models.SurahResp{
		Number:        2,
		NumberOfAyahs: 112,
		Name:          "Al-Baqarah",
		Translation:   "البقرة",
		Revelation:    "Mekah",
		Description:   "The second chapter of the Quran",
		Audio:         "al-baqarah.mp3",
	}

	surahs []models.SurahResp
)

func TestQuranHandler_GetSurah(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/quran/surah")
	//c.SetParamNames("email")
	//c.SetParamValues("jon@labstack.com")

	mockSurahRepo := repositories.NewMockSurahRepo(ctrl)
	mockAyahRepo := repositories.NewMockAyahRepo(ctrl)

	fmt.Println("mockAyahRepo: ", mockAyahRepo)

	fmt.Println("mockSurahRepo: ", mockSurahRepo)
	surahs := append(surahs, *surah1, *surah2)

	mockSurahRepo.EXPECT().GetSurahList().Return(surahs, nil)

	h := NewQuranHandler(mockSurahRepo, mockAyahRepo)

	// Assertions
	if assert.NoError(t, h.GetSurah(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, surahs, rec.Body.String())
	}

}
