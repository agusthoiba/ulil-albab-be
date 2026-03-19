package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	firebaseauth "firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"

	"ulil-albab-be/src/project/models"
	"ulil-albab-be/src/project/services"
)

type LastReadHandler struct {
	service services.LastReadServiceInt
}

func NewLastReadHandler(service services.LastReadServiceInt) *LastReadHandler {
	return &LastReadHandler{service: service}
}

func authIdentity(c echo.Context) (uid string, userID int, err error) {
	token, _ := c.Get("firebaseToken").(*firebaseauth.Token)
	if token == nil {
		return "", 0, echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	user, _ := c.Get("syncedUser").(*models.User)
	if user == nil {
		return "", 0, echo.NewHTTPError(http.StatusUnauthorized, "user not synced")
	}
	return token.UID, user.ID, nil
}

// GET /quran/last-read
func (h *LastReadHandler) GetLastRead(c echo.Context) error {
	uid, userID, err := authIdentity(c)
	if err != nil {
		return err
	}

	lastReadData := models.LastReadResp{
		FirebaseUID: uid,
		UserID:      userID,
		SuraID:      1,
		VerseID:     1,
	}

	lastReadDb, err := h.service.GetLastRead(uid)
	if err == nil {
		lastReadData = lastReadDb
	} else if err != sql.ErrNoRows {
		return err
	}

	return c.JSON(http.StatusOK, lastReadData)
}

// PUT /quran/last-read
func (h *LastReadHandler) PutLastRead(c echo.Context) error {
	uid, userID, err := authIdentity(c)
	if err != nil {
		return err
	}

	var req models.LastReadReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid json body")
	}

	req.FirebaseUID = uid
	req.UserID = userID

	fmt.Printf("surahId: %v, %T, verseId: %v, %T", req.SuraID, req.SuraID, req.VerseID, req.VerseID)
	if req.SuraID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "suraId must be > 0")
	}
	if req.VerseID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "verseId must be > 0")
	}
	if req.AyahID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "ayahId must be > 0")
	}

	resp, err := h.service.SaveLastRead(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
