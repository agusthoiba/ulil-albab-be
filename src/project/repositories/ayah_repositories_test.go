package repositories

import (
	"bytes"
	"database/sql"
	"errors"
	"testing"
	"ulil-albab-be/src/project/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"ulil-albab-be/src/project/utils"
)

func (ayat *AyahRepository) assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected := utils.EncodeToBytes(data)

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
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

	ayahListBytes := utils.EncodeToBytes(ayahListData)

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

func TestAyahRepositories_GetAllAyatError(t *testing.T) {
	db, qlmock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ayahRepo := &AyahRepository{db}

	mockError := errors.New("Some error")
	qlmock.ExpectQuery("^SELECT (.+) FROM quran_id$").WillReturnError(mockError)

	_, err = ayahRepo.GetAllAyat()
	assert.Error(t, mockError, err)
}

func TestAyahRepositories_GetAyatBySuraId(t *testing.T) {
	db, qlmock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
		AddRow(7, 1, 1, "الۤمّۤ ۚ", "Alif Lam Mim.", "alif lam mim", juzId)

	qlmock.ExpectQuery("SELECT * FROM quran_id WHERE sura_id = $1").WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

	ayahListData, _ := ayahRepo.GetAyatBySuratId(1)

	ayahListBytes := utils.EncodeToBytes(ayahListData)

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
		SuraId:   1,
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

func TestAyahRepositories_GetAyatBySuraIdError(t *testing.T) {
	db, qlmock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ayahRepo := &AyahRepository{db}

	mockError := errors.New("Some error")

	qlmock.ExpectQuery("SELECT * FROM quran_id WHERE sura_id = $1").WithArgs(sqlmock.AnyArg()).WillReturnError(mockError)

	_, err = ayahRepo.GetAyatBySuratId(1)

	assert.Error(t, mockError, err)
	// we make sure that all expectations were met
	if err := qlmock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
