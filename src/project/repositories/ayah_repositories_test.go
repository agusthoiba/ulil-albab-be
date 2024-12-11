package repositories

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"testing"
	"ulil-albab-be/src/project/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func (ayat *AyahRepository) assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected := ayat.EncodeToBytes(data)

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func (ayat *AyahRepository) EncodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		panic(err)
	}
	// fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

func TestAyahRepositories_GetAllAyat(t *testing.T) {
	db, qlmock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	juzId := sql.NullInt64{
		Int64: 1,
		Valid: true,
	}

	ayahRepo := &AyahRepository{db}

	rows := sqlmock.NewRows([]string{"id", "sura_id", "verse_id", "ayah_text", "indo_text", "read_text", "juz_id"}).
		AddRow(0, 1, 1, "بِسْمِ اللّٰهِ الرَّحْمٰنِ الرَّحِيْمِ", "Dengan nama Allah Yang Maha Pengasih, Maha Penyayang.", "bismillahir-rahmanir-rahim", juzId).
		AddRow(7, 2, 1, "الۤمّۤ ۚ", "Alif Lam Mim.", "alif lam mim", juzId)

	qlmock.ExpectQuery("^SELECT (.+) FROM quran_id$").WillReturnRows(rows)

	ayahListData, _ := ayahRepo.GetAllAyat()

	ayahListBytes := ayahRepo.EncodeToBytes(ayahListData)

	var data []models.AyatResp

	ayat1 := models.AyatResp{
		Id:       0,
		SuraId:   1,
		VerseID:  1,
		AyahText: "بِسْمِ اللّٰهِ الرَّحْمٰنِ الرَّحِيْمِ",
		IndoText: "Dengan nama Allah Yang Maha Pengasih, Maha Penyayang.",
		ReadText: "bismillahir-rahmanir-rahim",
		JuzId:    juzId,
	}

	ayat2 := models.AyatResp{
		Id:       7,
		SuraId:   2,
		VerseID:  1,
		AyahText: "الۤمّۤ ۚ",
		IndoText: "Alif Lam Mim.",
		ReadText: "alif lam mim",
		JuzId:    juzId,
	}

	data = append(data, ayat1, ayat2)

	ayahRepo.assertJSON(ayahListBytes, data, t)

	// we make sure that all expectations were met
	if err := qlmock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
