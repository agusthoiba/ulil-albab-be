package services

import (
	"sync"
	"testing"

	"ulil-albab-be/src/project/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSurahRepo struct {
	mock.Mock
}

func (m *MockSurahRepo) GetSurahList() ([]models.SurahResp, error) {
	args := m.Called()
	return args.Get(0).([]models.SurahResp), args.Error(1)

}

func (m *MockSurahRepo) GetSurahListRoutine(*sync.WaitGroup, chan []models.SurahResp) {
	//args := m.Called()
	return
}

type MockAyahRepo struct {
	mock.Mock
}

func (m *MockAyahRepo) GetAllAyat() ([]models.AyatResp, error) {
	args := m.Called()
	return args.Get(0).([]models.AyatResp), args.Error(1)

}

func TestService_GetQuranSurah(t *testing.T) {
	mockSurahRepo := new(MockSurahRepo)
	// mockAyahRepo := new(MockAyahRepo)

	service := NewService(mockSurahRepo, nil)

	surah1 := &models.SurahResp{
		Number:        1,
		NumberOfAyahs: 114,
		Name:          "Al-Fatihah",
		Translation:   "الفاتحة",
		Revelation:    "Mekah",
		Description:   "The first chapter of the Quran",
		Audio:         "al-fatihah.mp3",
	}

	surah2 := &models.SurahResp{
		Number:        2,
		NumberOfAyahs: 112,
		Name:          "Al-Baqarah",
		Translation:   "البقرة",
		Revelation:    "Mekah",
		Description:   "The second chapter of the Quran",
		Audio:         "al-baqarah.mp3",
	}

	var surahs []models.SurahResp

	surahs = append(surahs, *surah1, *surah2)

	t.Run("successful get surah", func(t *testing.T) {
		mockSurahRepo.On("GetSurahList").Return(surahs, nil).Once()

		_, err := service.GetSurah()

		assert.NoError(t, err)
		mockSurahRepo.AssertExpectations(t)
	})

}
