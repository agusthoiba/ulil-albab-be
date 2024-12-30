package services

import (
	"sync"
	"ulil-albab-be/src/project/models"
	"ulil-albab-be/src/project/repositories"
)

type Service struct {
	surahRepo repositories.SurahRepo
	ayahRepo  repositories.AyahRepo
}

type ServiceInt interface {
	GetSurah() ([]models.SurahResp, error)
	GetAllAyat() ([]models.AyatResp, error)
	GetAyatBySuratId(id int) ([]models.AyatResp, error)
	GetAll() (models.QuranAllResp, error)
}

func NewService(surahRepo repositories.SurahRepo, ayahRepo repositories.AyahRepo) *Service {
	return &Service{
		surahRepo: surahRepo,
		ayahRepo:  ayahRepo,
	}
}

func (s *Service) GetSurah() ([]models.SurahResp, error) {

	surahData, err := s.surahRepo.GetSurahList()
	return surahData, err
}

func (s *Service) GetAllAyat() ([]models.AyatResp, error) {

	ayahData, err := s.ayahRepo.GetAllAyat()
	return ayahData, err
}

func (s *Service) GetAyatBySuratId(id int) ([]models.AyatResp, error) {

	ayahData, err := s.ayahRepo.GetAyatBySuratId(id)
	return ayahData, err
}

func (s *Service) GetAll() (models.QuranAllResp, error) {
	// var err error

	var wg sync.WaitGroup

	surahChan := make(chan []models.SurahResp)
	ayahChan := make(chan []models.AyatResp)

	wg.Add(2)
	go s.surahRepo.GetSurahListRoutine(&wg, surahChan)
	go s.ayahRepo.GetAllAyatRoutine(&wg, ayahChan)

	surahData := <-surahChan
	ayahData := <-ayahChan

	wg.Wait()

	allData := models.QuranAllResp{
		Surahs: surahData,
		Ayahs:  ayahData,
	}

	return allData, nil
}
