package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

const firebaseTokenKey = "firebaseToken"

var FirebaseAuth *auth.Client

func InitFirebase(credentialPath string) {
	credBytes, err := os.ReadFile(credentialPath)
	if err != nil {
		log.Fatalf("Error reading service account key: %v", err)
	}
	opt := option.WithCredentialsJSON(credBytes)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}

	FirebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error getting Auth client: %v", err)
	}

	log.Println("Firebase initialized successfully")
}

// FirebaseAuthMiddleware verifies the Firebase Bearer token on protected routes.
func FirebaseAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization format")
			}

			token, err := FirebaseAuth.VerifyIDToken(context.Background(), parts[1])
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
			}

			c.Set(firebaseTokenKey, token)
			return next(c)
		}
	}
}

// GetTokenFromContext retrieves the verified Firebase token from the Echo context.
func GetTokenFromContext(c echo.Context) *auth.Token {
	token, _ := c.Get(firebaseTokenKey).(*auth.Token)
	return token
}
