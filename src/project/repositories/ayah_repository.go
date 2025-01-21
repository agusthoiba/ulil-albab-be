package repositories

import (
	"database/sql"
	"sync"

	"ulil-albab-be/src/project/models"
	"ulil-albab-be/src/project/logger"

	_ "github.com/lib/pq"
)

type AyahRepository struct {
	db *sql.DB
	logger *logger.LogClass
}

type AyahRepo interface {
	GetAllAyat() ([]models.AyatResp, error)
	GetAyatBySuratId(int) ([]models.AyatResp, error)
	GetAllAyatRoutine(*sync.WaitGroup, chan []models.AyatResp)
}

// constructor
func NewAyah(db *sql.DB, logger *logger.LogClass) *AyahRepository {
	return &AyahRepository{
		db: db,
		logger: logger,
	}
}

func (ay *AyahRepository) GetAllAyat() ([]models.AyatResp, error) {
	rows, err := ay.db.Query("SELECT * FROM quran_id")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ayats []models.AyatResp

	for rows.Next() {
		var ayat models.AyatResp
		if err := rows.Scan(&ayat.Id, &ayat.SuraId, &ayat.VerseID, &ayat.AyahText, &ayat.IndoText,
			&ayat.ReadText, &ayat.JuzId); err != nil {
			return nil, err
		}
		ayats = append(ayats, ayat)
	}

	return ayats, nil
}

func (ay *AyahRepository) GetAyatBySuratId(suraId int) ([]models.AyatResp, error) {
	rows, err := ay.db.Query("SELECT * FROM quran_id WHERE sura_id = $1", suraId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ayats []models.AyatResp
	for rows.Next() {
		var ayat models.AyatResp
		if err := rows.Scan(&ayat.Id, &ayat.SuraId, &ayat.VerseID, &ayat.AyahText, &ayat.IndoText,
			&ayat.ReadText, &ayat.JuzId); err != nil {
			return nil, err
		}
		ayats = append(ayats, ayat)
	}

	return ayats, nil
}

func (ay *AyahRepository) GetAllAyatRoutine(wg *sync.WaitGroup, ch chan []models.AyatResp) {
	defer wg.Done()
	rows, err := ay.db.Query("SELECT * FROM quran_id")

	if err != nil {
		ay.logger.Log().Error(err)
	}

	defer rows.Close()

	var ayats []models.AyatResp

	for rows.Next() {
		var ayat models.AyatResp
		if err := rows.Scan(&ayat.Id, &ayat.SuraId, &ayat.VerseID, &ayat.AyahText, &ayat.IndoText,
			&ayat.ReadText, &ayat.JuzId); err != nil {
			ay.logger.Log().Error(err)
		}
		ayats = append(ayats, ayat)
	}

	ch <- ayats

	close(ch)
}
