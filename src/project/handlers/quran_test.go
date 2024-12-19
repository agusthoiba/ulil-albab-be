package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"ulil-albab-be/src/project/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

type MockServices struct {
	mock.Mock
}

func (m *MockServices) GetSurah() ([]models.SurahResp, error) {
	args := m.Called()
	return args.Get(0).([]models.SurahResp), args.Error(1)

}

func (m *MockServices) GetAllAyat() ([]models.AyatResp, error) {
	args := m.Called()
	return args.Get(0).([]models.AyatResp), args.Error(1)
}

func (m *MockServices) GetAyatBySuratId(id int) ([]models.AyatResp, error) {
	args := m.Called(id)
	return args.Get(0).([]models.AyatResp), args.Error(1)
}

func TestQuranHandler_GetSurah(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/quran/surah")
	//c.SetParamNames("email")
	//c.SetParamValues("jon@labstack.com")

	surahs := append(surahs, *surah1, *surah2)

	mockServices := new(MockServices)

	h := NewQuranHandler(mockServices)

	mockServices.On("GetSurah").Return(surahs, nil).Once()

	// Assertions
	if assert.NoError(t, h.GetSurah(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.Equal(t, surahs, rec.Body.String())
	}
}
