package models

import "time"

type LastReadReq struct {
	FirebaseUID string `json:"firebaseUID"`
	UserID      int    `json:"userId"`
	SuraID      int    `json:"suraId"`
	VerseID     int    `json:"verseId"`
	AyahID      int    `json:"ayahId"`
}

type LastReadResp struct {
	ID          int       `json:"id"`
	FirebaseUID string    `json:"firebaseUID"`
	UserID      int       `json:"userId"`
	SuraID      int       `json:"suraId"`
	VerseID     int       `json:"verseId"`
	AyahID      int       `json:"ayahId"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
