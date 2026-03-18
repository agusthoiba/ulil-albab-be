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
	Number        int            `db:"number" json:"number"`
	NumberOfAyahs int            `db:"numberofayahs" json:"numberOfAyahs"`
	Name          string         `db:"name" json:"name"`
	Translation   string         `db:"translation" json:"translation"`
	Revelation    string         `db:"revelation" json:"revelation"`
	Description   string         `db:"description" json:"description"`
	Audio         string         `db:"audio" json:"audio"`
	NameArab      sql.NullString `db:"name_arab" json:"nameArab"`
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

type User struct {
	ID          int    `json:"id" db:"id"`
	FirebaseUID string `json:"firebase_uid" db:"firebase_uid"`
	Email       string `json:"email" db:"email"`
	Name        string `json:"name" db:"name"`
	PhotoURL    string `json:"photo_url" db:"photo_url"`
	Provider    string `json:"provider" db:"provider"`
	FCMToken    string `json:"fcm_token" db:"fcm_token"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

type AuthRequest struct {
	FCMToken string `json:"fcmToken"`
}

type AuthResponse struct {
	Success bool   `json:"success"`
	User    *User  `json:"user"`
	Message string `json:"message"`
}
