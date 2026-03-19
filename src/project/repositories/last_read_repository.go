package repositories

import (
	"database/sql"
	"time"

	"ulil-albab-be/src/project/logger"
	"ulil-albab-be/src/project/models"
)

type LastReadRepo interface {
	EnsureTable() error
	GetByFirebaseUID(uid string) (models.LastReadResp, error)
	Upsert(req models.LastReadReq) (models.LastReadResp, error)
}

type LastReadRepository struct {
	db     *sql.DB
	logger *logger.LogClass
}

func NewLastReadRepository(db *sql.DB, logger *logger.LogClass) *LastReadRepository {
	return &LastReadRepository{
		db:     db,
		logger: logger,
	}
}

func (r *LastReadRepository) EnsureTable() error {
	_, err := r.db.Exec(`
CREATE TABLE IF NOT EXISTS user_last_read (
  id           SERIAL      PRIMARY KEY,
  firebase_uid TEXT        NOT NULL UNIQUE,
  user_id      INTEGER     NOT NULL REFERENCES users(id),
  sura_id      INTEGER     NOT NULL,
  verse_id     INTEGER     NOT NULL,
  updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);
`)
	return err
}

func (r *LastReadRepository) GetByFirebaseUID(uid string) (models.LastReadResp, error) {
	var resp models.LastReadResp
	var updatedAt time.Time

	err := r.db.QueryRow(
		`SELECT id, firebase_uid, user_id, sura_id, verse_id, updated_at
		 FROM user_last_read WHERE firebase_uid = $1`,
		uid,
	).Scan(&resp.ID, &resp.FirebaseUID, &resp.UserID, &resp.SuraID, &resp.VerseID, &updatedAt)
	if err != nil {
		return models.LastReadResp{}, err
	}

	resp.UpdatedAt = updatedAt
	return resp, nil
}

func (r *LastReadRepository) Upsert(req models.LastReadReq) (models.LastReadResp, error) {
	var resp models.LastReadResp
	var updatedAt time.Time

	err := r.db.QueryRow(
		`INSERT INTO user_last_read (firebase_uid, user_id, sura_id, verse_id)
		 VALUES ($1, $2, $3, $4)
		 ON CONFLICT (firebase_uid)
		 DO UPDATE SET user_id = EXCLUDED.user_id, sura_id = EXCLUDED.sura_id, verse_id = EXCLUDED.verse_id, updated_at = now()
		 RETURNING id, firebase_uid, user_id, sura_id, verse_id, updated_at`,
		req.FirebaseUID,
		req.UserID,
		req.SuraID,
		req.VerseID,
	).Scan(&resp.ID, &resp.FirebaseUID, &resp.UserID, &resp.SuraID, &resp.VerseID, &updatedAt)
	if err != nil {
		if r.logger != nil {
			r.logger.Log().Error(err)
		}
		return models.LastReadResp{}, err
	}

	resp.UpdatedAt = updatedAt
	return resp, nil
}
