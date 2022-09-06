package repository_test

import (
	"template/interface/repository"
	"template/testutil"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func setup(t *testing.T) (r repository.UserRepository, mock sqlmock.Sqlmock, teardown func()) {
	db, mock, _ := testutil.NewDBMock()

	r = repository.NewUserRepository(db)

	return r, mock, func() {
		db, _ := db.DB()
		db.Close()
	}
}
