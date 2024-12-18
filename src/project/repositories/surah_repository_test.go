package repositories

import (
	"bytes"
	"encoding/gob"
	"testing"
	"ulil-albab-be/src/project/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func (sr *SurahRepository) assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected := sr.EncodeToBytes(data)

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func (sr *SurahRepository) EncodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		panic(err)
	}
	// fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

func TestSurahRepositories_GetSurahList(t *testing.T) {
	db, qlmock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	surRepo := &SurahRepository{db}

	rows := sqlmock.NewRows([]string{"number", "numberOfAyahs", "name", "translation", "revelation", "description", "audio"}).
		AddRow(1, 7, "Al-Fatihah", "Pembukaan", "Mekah", "Pembukaan Desc", "alfatihah.mp3").
		AddRow(2, 286, "Al-Baqarah", "Sapi", "Madaniyah", "Surat yang terpanjang", "albaqarah.mp3")

	qlmock.ExpectQuery("^SELECT (.+) FROM surah$").WillReturnRows(rows)

	surahListData, _ := surRepo.GetSurahList()

	surahListBytes := surRepo.EncodeToBytes(surahListData)

	var data []models.SurahResp

	surah1 := models.SurahResp{
		Number:        1,
		NumberOfAyahs: 7,
		Name:          "Al-Fatihah",
		Translation:   "Pembukaan",
		Revelation:    "Mekah",
		Description:   "Pembukaan Desc",
		Audio:         "alfatihah.mp3",
	}

	surah2 := models.SurahResp{
		Number:        2,
		NumberOfAyahs: 286,
		Name:          "Al-Baqarah",
		Translation:   "Sapi",
		Revelation:    "Madaniyah",
		Description:   "Surat yang terpanjang",
		Audio:         "albaqarah.mp3",
	}

	data = append(data, surah1)
	data = append(data, surah2)

	surRepo.assertJSON(surahListBytes, data, t)

	// we make sure that all expectations were met
	if err := qlmock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
