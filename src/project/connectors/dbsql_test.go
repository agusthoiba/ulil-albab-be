package connectors

import (
	"database/sql"
	"errors"
	"testing"

	"ulil-albab-be/src/project/models"
)

type mockSqlDb struct {
	sql.DB
}

func (msd *mockSqlDb) Ping() error {
	return nil
}

func TestInitDB(t *testing.T) {
	mockError := errors.New("uh oh")
	mockSqlDb := &mockSqlDb{}

	subtests := []struct {
		name        string
		option      models.OptionDb
		sqlOpener   func(s string, s2 string) (db *sql.DB, err error)
		expectedErr error
	}{
		{
			name: "happy path",
			option: models.OptionDb{
				Host:     "a",
				Port:     19,
				User:     "u",
				Password: "p",
				DbName:   "db",
			},
			sqlOpener: func(s string, s2 string) (*sql.DB, error) {
				db := mockSqlDb
				//db.Ping()
				return &db.DB, nil
			},
		},
		{
			name: "error from sqlOpener",
			sqlOpener: func(s string, s2 string) (db *sql.DB, err error) {
				return nil, mockError
			},
			expectedErr: mockError,
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			_, err := InitDB(subtest.option, subtest.sqlOpener)
			if !errors.Is(err, subtest.expectedErr) {
				t.Errorf("expected error (%v), got error (%v)", subtest.expectedErr, err)
			}
		})
	}
}
