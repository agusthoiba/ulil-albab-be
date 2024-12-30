package models

import "database/sql"

type OptionDb struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type ErrorResp struct {
	Code         int    `json:"code"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type Surah struct {
	number        int
	numberOfAyahs int
	name          string
	translation   string
	revelation    string
	description   string
	audio         string
}

type SurahResp struct {
	Number        int    `json:"number"`
	NumberOfAyahs int    `json:"numberOfAyahs"`
	Name          string `json:"name"`
	Translation   string `json:"translation"`
	Revelation    string `json:"revelation"`
	Description   string `json:"description"`
	Audio         string `json:"audio"`
}

type AyatResp struct {
	Id       int           `db:"id" json:"id"`
	SuraId   int           `db:"sura_id" json:"suraId"`
	AyahText string        `db:"ayah_text" json:"ayahText"`
	IndoText string        `db:"indo_text"  json:"indoText"`
	ReadText string        `db:"read_text" json:"ReadText"`
	JuzId    sql.NullInt64 `db:"juz_id" json:"juzId"`
	VerseID  int           `db:"verse_id" json:"verseID"`
}

type QuranAllResp struct {
	Surahs []SurahResp `json:"surahs"`
	Ayahs  []AyatResp  `json:"ayahs"`
}
