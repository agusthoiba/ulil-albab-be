package repositories

import (
	"fmt"
	"ulil-albab-be/src/project/connectors"
	"ulil-albab-be/src/project/models"

	"github.com/labstack/echo/v4"
)

func GetSurahList(c echo.Context) ([]models.SurahResp, error) {
	db := connectors.GetDB(c)

	fmt.Println("db", db)

	rows, err := db.Query("SELECT * FROM surah")
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

func GetAyatBySuratId(c echo.Context, suraId int) ([]models.AyatResp, error) {
	db := connectors.GetDB(c)

	rows, err := db.Query(`SELECT * FROM quran_id WHERE sura_id=$1;`, suraId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var ayats []models.AyatResp
	for rows.Next() {
		var ayat models.AyatResp
		if err := rows.Scan(&ayat.Id, &ayat.SuraId, &ayat.AyahText, &ayat.IndoText,
			&ayat.ReadText, &ayat.JuzId, &ayat.VerseID); err != nil {
			return nil, err
		}
		ayats = append(ayats, ayat)
	}

	return ayats, nil
}

func GetAllAyat(c echo.Context) ([]models.AyatResp, error) {
	db := connectors.GetDB(c)

	rows, err := db.Query(`SELECT * FROM quran_id`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var ayats []models.AyatResp
	for rows.Next() {
		var ayat models.AyatResp
		if err := rows.Scan(&ayat.Id, &ayat.SuraId, &ayat.AyahText, &ayat.IndoText,
			&ayat.ReadText, &ayat.JuzId, &ayat.VerseID); err != nil {
			return nil, err
		}
		ayats = append(ayats, ayat)
	}

	return ayats, nil
}
