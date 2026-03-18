package services

import (
	"context"
	"time"

	firebaseauth "firebase.google.com/go/v4/auth"

	"ulil-albab-be/src/project/models"
)

type AuthRepo interface {
	EnsureTable() error
	UpsertUser(user *models.User) (*models.User, error)
}

type AuthService interface {
	SyncUser(ctx context.Context, token *firebaseauth.Token) (*models.User, error)
}

type AuthServiceImpl struct {
	repo           AuthRepo
	firebaseClient *firebaseauth.Client
}

func NewAuthService(repo AuthRepo, firebaseClient *firebaseauth.Client) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo, firebaseClient: firebaseClient}
}

// SyncUser upserts the Firebase user into the local DB. Called transparently on every
// authenticated request — no explicit login endpoint needed.
func (s *AuthServiceImpl) SyncUser(ctx context.Context, token *firebaseauth.Token) (*models.User, error) {
	email, _ := token.Claims["email"].(string)
	name, _ := token.Claims["name"].(string)
	picture, _ := token.Claims["picture"].(string)
	provider := token.Firebase.SignInProvider

	// Enrich from Firebase user record when token claims are incomplete
	if userRecord, err := s.firebaseClient.GetUser(ctx, token.UID); err == nil {
		if email == "" {
			email = userRecord.Email
		}
		if name == "" {
			name = userRecord.DisplayName
		}
		if picture == "" {
			picture = userRecord.PhotoURL
		}
	}

	user := &models.User{
		FirebaseUID: token.UID,
		Email:       email,
		Name:        name,
		PhotoURL:    picture,
		Provider:    provider,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	return s.repo.UpsertUser(user)
}
