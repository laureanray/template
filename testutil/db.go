package testutil

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func NewDBMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, error) {
	t.Helper()

	// TODO: Implement this!
}
