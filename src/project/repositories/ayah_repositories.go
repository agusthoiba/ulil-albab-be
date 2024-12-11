package repositories

import (
	"database/sql"
	"fmt"
	"ulil-albab-be/src/project/models"

	_ "github.com/lib/pq"
)

type AyahRepository struct {
	db *sql.DB
}

type AyahRepo interface {
	GetAllAyat() ([]models.AyatResp, error)
	GetAyatBySuratId(int) ([]models.AyatResp, error)
}

// constructor
func NewAyah(db *sql.DB) *AyahRepository {
	return &AyahRepository{
		db: db,
	}
}

func (ayat *AyahRepository) GetAllAyat() ([]models.AyatResp, error) {
	rows, err := ayat.db.Query("SELECT * FROM quran_id")

	if err != nil {
		fmt.Println(err)
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

func (ayat *AyahRepository) GetAyatBySuratId(suraId int) ([]models.AyatResp, error) {
	rows, err := ayat.db.Query(`SELECT * FROM quran_id WHERE sura_id = $1`, suraId)
	if err != nil {
		fmt.Println(err)
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
