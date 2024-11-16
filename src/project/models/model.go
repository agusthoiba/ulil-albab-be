package models

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
