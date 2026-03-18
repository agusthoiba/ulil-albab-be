package models

import "time"

type LastReadReq struct {
	FirebaseUID string `json:"firebaseUID"`
	UserID      int    `json:"userId"`
	SuraID      int    `json:"suraId"`
	VerseID     int    `json:"verseId"`
}

type LastReadResp struct {
	ID          int       `json:"id"`
	FirebaseUID string    `json:"firebaseUID"`
	UserID      int       `json:"userId"`
	SuraID      int       `json:"suraId"`
	VerseID     int       `json:"verseId"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
