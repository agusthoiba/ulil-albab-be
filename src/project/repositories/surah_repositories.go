package repositories

import (
	"database/sql"
	"ulil-albab-be/src/project/models"
)

type SurahRepository struct {
	db *sql.DB
}

type SurahRepo interface {
	GetSurahList() ([]models.AyatResp, error)
}

func NewSurah(db *sql.DB) *SurahRepository {
	return &SurahRepository{
		db: db,
	}
}

func (sr *SurahRepository) GetSurahList() ([]models.SurahResp, error) {
	rows, err := sr.db.Query("SELECT * FROM surah")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var surahs []models.SurahResp
	for rows.Next() {
		var surah models.SurahResp
		if err := rows.Scan(&surah.Number, &surah.NumberOfAyahs, &surah.Name, &surah.Translation,
			&surah.Revelation, &surah.Description, &surah.Audio); err != nil {
			return nil, err
		}

		surahs = append(surahs, surah)
	}

	return surahs, nil
}
