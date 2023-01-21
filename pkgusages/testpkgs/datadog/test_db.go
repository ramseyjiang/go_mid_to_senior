package datadog

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
)

func TestUserDBRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("expected not to have error")
	}

	mock.
		ExpectQuery("SELECT * FROM users WHERE lastname = ?").
		WithArgs("some last name").
		WillReturnError(errors.New("error"))

	repository := NewUserDBRepository(db)
	users, err := repository.FilterByLastname(context.Background(), "some last name")
	log.Println(users)
}
