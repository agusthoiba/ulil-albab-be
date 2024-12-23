package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
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

	surahs := append(surahs, *surah1, *surah2)

	mockServices := new(MockServices)

	h := NewQuranHandler(mockServices)

	mockServices.On("GetSurah").Return(surahs, nil).Once()

	surahJsonBytes, _ := json.Marshal(surahs)
	surahJsonStr := string(surahJsonBytes)
	surahJsonStr = surahJsonStr + "\n"

	// Assertions
	if assert.NoError(t, h.GetSurah(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, surahJsonStr, rec.Body.String())
	}
}

func TestQuranHandler_GetSurahError(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/quran/surah")

	mockServices := new(MockServices)

	h := NewQuranHandler(mockServices)

	mockServiceError := errors.New("Service error")

	mockServices.On("GetSurah").Return(surahs, mockServiceError).Once()

	// Assertions
	assert.Error(t, h.GetSurah(c))
}
func TestQuranHandler_GetAllAyats(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/quran/ayat")

	juzId := sql.NullInt64{
		Int64: 1,
		Valid: true,
	}

	var ayats []models.AyatResp

	ayat1 := &models.AyatResp{
		Id:       0,
		SuraId:   1,
		VerseID:  1,
		AyahText: "بِسْمِ اللّٰهِ الرَّحْمٰنِ الرَّحِيْمِ",
		IndoText: "Dengan nama Allah Yang Maha Pengasih, Maha Penyayang.",
		ReadText: "bismillahir-rahmanir-rahim",
		JuzId:    juzId,
	}

	ayat2 := &models.AyatResp{
		Id:       7,
		SuraId:   2,
		VerseID:  1,
		AyahText: "الۤمّۤ ۚ",
		IndoText: "Alif Lam Mim.",
		ReadText: "alif lam mim",
		JuzId:    juzId,
	}

	ayats = append(ayats, *ayat1, *ayat2)

	mockServices := new(MockServices)

	h := NewQuranHandler(mockServices)

	mockServices.On("GetAllAyat").Return(ayats, nil).Once()

	ayatJsonBytes, _ := json.Marshal(ayats)
	ayatJsonStr := string(ayatJsonBytes)
	ayatJsonStr = ayatJsonStr + "\n"

	// Assertions
	if assert.NoError(t, h.GetAllAyats(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, ayatJsonStr, rec.Body.String())
	}
}

func TestQuranHandler_GetAllAyatsError(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/quran/ayat")

	var ayats []models.AyatResp

	mockServices := new(MockServices)
	h := NewQuranHandler(mockServices)

	mockServiceError := errors.New("Service error")

	mockServices.On("GetAllAyat").Return(ayats, mockServiceError).Once()

	// Assertions
	assert.Error(t, h.GetAllAyats(c))
}

func TestQuranHandler_GetAyatBySurahId(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/quran/ayat/:suraId")
	c.SetParamNames("suraId")
	c.SetParamValues("1")

	juzId := sql.NullInt64{
		Int64: 1,
		Valid: true,
	}

	var ayats []models.AyatResp

	ayat1 := &models.AyatResp{
		Id:       0,
		SuraId:   1,
		VerseID:  1,
		AyahText: "بِسْمِ اللّٰهِ الرَّحْمٰنِ الرَّحِيْمِ",
		IndoText: "Dengan nama Allah Yang Maha Pengasih, Maha Penyayang.",
		ReadText: "bismillahir-rahmanir-rahim",
		JuzId:    juzId,
	}

	ayat2 := &models.AyatResp{
		Id:       7,
		SuraId:   1,
		VerseID:  1,
		AyahText: "الۤمّۤ ۚ",
		IndoText: "Alif Lam Mim.",
		ReadText: "alif lam mim",
		JuzId:    juzId,
	}

	ayats = append(ayats, *ayat1, *ayat2)

	mockServices := new(MockServices)

	h := NewQuranHandler(mockServices)

	mockServices.On("GetAyatBySuratId", 1).Return(ayats, nil).Once()

	ayatJsonBytes, _ := json.Marshal(ayats)
	ayatJsonStr := string(ayatJsonBytes)
	ayatJsonStr = ayatJsonStr + "\n"

	// Assertions
	if assert.NoError(t, h.GetAyats(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, ayatJsonStr, rec.Body.String())
	}
}

func TestQuranHandler_GetAyatBySurahIdError(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/quran/ayat/:suraId")
	c.SetParamNames("suraId")
	c.SetParamValues("1")

	var ayats []models.AyatResp

	mockServices := new(MockServices)

	h := NewQuranHandler(mockServices)

	mockServiceError := errors.New("Service error")

	mockServices.On("GetAyatBySuratId", 1).Return(ayats, mockServiceError).Once()

	// Assertions
	assert.Error(t, h.GetAyats(c))
}

func TestQuranHandler_GetAyatBySurahIdErrorBadInput(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/quran/ayat/:suraId")
	c.SetParamNames("suraId")
	c.SetParamValues("a")

	var ayats []models.AyatResp

	mockServices := new(MockServices)

	h := NewQuranHandler(mockServices)

	mockServiceError := errors.New("Atoi must be number")

	mockServices.On("GetAyatBySuratId", "a").Return(ayats, mockServiceError).Once()

	// Assertions
	assert.Error(t, h.GetAyats(c))
}
