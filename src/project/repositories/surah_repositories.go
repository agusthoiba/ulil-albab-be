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

		fmt.Println("surah", surah)
		surahs = append(surahs, surah)
	}

	return surahs, nil

}
