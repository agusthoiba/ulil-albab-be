package repositories

import (
	"database/sql"
	"fmt"
	"sync"
	"ulil-albab-be/src/project/models"
)

type SurahRepository struct {
	db *sql.DB
}

type SurahRepo interface {
	GetSurahList() ([]models.SurahResp, error)
	GetSurahListRoutine(*sync.WaitGroup, chan []models.SurahResp)
}

func NewSurah(db *sql.DB) *SurahRepository {
	return &SurahRepository{
		db: db,
	}
}

func (sr *SurahRepository) GetSurahList() ([]models.SurahResp, error) {
	rows, err := sr.db.Query("SELECT * FROM surah ORDER BY number")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var surahs []models.SurahResp
	for rows.Next() {
		var surah models.SurahResp
		if err := rows.Scan(&surah.Number, &surah.NumberOfAyahs, &surah.Name, &surah.Translation,
			&surah.Revelation, &surah.Description, &surah.Audio, &surah.NameArab); err != nil {
			return nil, err
		}

		surahs = append(surahs, surah)
	}

	return surahs, nil
}

func (sr *SurahRepository) GetSurahListRoutine(wg *sync.WaitGroup, ch chan []models.SurahResp) {
	defer wg.Done()
	rows, err := sr.db.Query("SELECT * FROM surah ORDER BY number")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var surahs []models.SurahResp
	for rows.Next() {
		var surah models.SurahResp
		if err := rows.Scan(&surah.Number, &surah.NumberOfAyahs, &surah.Name, &surah.Translation,
			&surah.Revelation, &surah.Description, &surah.Audio, &surah.NameArab); err != nil {
			fmt.Println(err)
		}

		surahs = append(surahs, surah)
	}

	ch <- surahs
	close(ch)
}
