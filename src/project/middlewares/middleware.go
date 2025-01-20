package middlewares

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"database/sql"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"

	logging "ulil-albab-be/src/project/logger"
	"ulil-albab-be/src/project/connectors"
	"ulil-albab-be/src/project/handlers"
	"ulil-albab-be/src/project/models"
	"ulil-albab-be/src/project/repositories"
	"ulil-albab-be/src/project/services"
)

// DBMiddleware adds the database connection to the context
func DBMiddleware(db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), "db", db)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func NewMiddleware(e *echo.Echo) error {
	skipper := func(c echo.Context) bool {
		// Skip health check endpoint
		return c.Request().URL.Path == "/health"
	}

	logger := logging.NewInitiateLogger()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogError:    true,
		LogMethod: true,
		LogRemoteIP: true,
		LogRequestID: true,
		Skipper: skipper,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				logs := logger.Log().WithFields(logrus.Fields{
					"method": v.Method,
					"uri":   v.URI,
					"headers": c.Request().Header,
					"body": c.Request().Body,
					"status": v.Status,
					"latency": v.Latency,
					"ip": v.RemoteIP,
				})
			
				if v.Error!= nil {
                    logs.WithError(v.Error).Error("error")
                } else {
                    logs.Info("request")
                }

        		return nil
    	},
	}))

	for _, env := range os.Environ() {
		if !strings.HasPrefix(env, "DB_SQL_PASSWORD") {
			logger.Log().Info(env)
		}
    }


	dbPort, err := strconv.Atoi(os.Getenv("DB_SQL_PORT"))
	if err != nil {
		panic(err)
	}

	optionDb := models.OptionDb{
		Host:     os.Getenv("DB_SQL_HOST"),
		Port:     dbPort,
		User:     os.Getenv("DB_SQL_USER"),
		Password: os.Getenv("DB_SQL_PASSWORD"),
		DbName:   os.Getenv("DB_SQL_NAME"),
	}

	db, err := connectors.InitDB(optionDb, logger)

	if err != nil {
		logger.Log().Error(err)
		return err
	}

	e.Use(DBMiddleware(db))

	ayahRepo := repositories.NewAyah(db)
	surahRepo := repositories.NewSurah(db)

	service := services.NewService(surahRepo, ayahRepo)

	quranHandler := handlers.NewQuranHandler(service)

	e.GET("/quran/surah", quranHandler.GetSurah)
	e.GET("/quran/ayat/:suraId", quranHandler.GetAyats)
	e.GET("/quran/ayat", quranHandler.GetAllAyats)
	e.GET("/quran", quranHandler.GetAll)

	return nil
}
