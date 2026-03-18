package repositories

import (
	"database/sql"

	"ulil-albab-be/src/project/logger"
	"ulil-albab-be/src/project/models"
)

type UserRepository struct {
	db     *sql.DB
	logger *logger.LogClass
}

func NewUserRepository(db *sql.DB, logger *logger.LogClass) *UserRepository {
	return &UserRepository{db: db, logger: logger}
}

func (r *UserRepository) EnsureTable() error {
	_, err := r.db.Exec(`
CREATE TABLE IF NOT EXISTS public.users (
  id           SERIAL PRIMARY KEY,
  firebase_uid TEXT        NOT NULL UNIQUE,
  email        TEXT        NOT NULL,
  name         TEXT        NOT NULL DEFAULT '',
  photo_url    TEXT        NOT NULL DEFAULT '',
  provider     TEXT        NOT NULL DEFAULT '',
  fcm_token    TEXT        NOT NULL DEFAULT '',
  created_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);
`)
	return err
}

func (r *UserRepository) UpsertUser(user *models.User) (*models.User, error) {
	var saved models.User

	err := r.db.QueryRow(`
INSERT INTO public.users (firebase_uid, email, name, photo_url, provider, fcm_token)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (firebase_uid)
DO UPDATE SET
  email     = EXCLUDED.email,
  name      = EXCLUDED.name,
  photo_url = EXCLUDED.photo_url,
  provider  = EXCLUDED.provider,
  fcm_token = EXCLUDED.fcm_token
RETURNING id, firebase_uid, email, name, photo_url, provider, fcm_token, created_at`,
		user.FirebaseUID,
		user.Email,
		user.Name,
		user.PhotoURL,
		user.Provider,
		user.FCMToken,
	).Scan(
		&saved.ID,
		&saved.FirebaseUID,
		&saved.Email,
		&saved.Name,
		&saved.PhotoURL,
		&saved.Provider,
		&saved.FCMToken,
		&saved.CreatedAt,
	)
	if err != nil {
		if r.logger != nil {
			r.logger.Log().Error(err)
		}
		return nil, err
	}

	return &saved, nil
}
