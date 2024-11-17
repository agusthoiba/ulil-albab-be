package models

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
	Id       int    `json:"id"`
	SuraId   int    `json:"suraId"`
	AyahText string `json:"ayahText"`
	IndoText string `json:"indoText"`
	ReadText string `json:"ReadText"`
	JuzId    *int   `json:"juzId"`
	VerseID  int    `json:"verseID"`
}
