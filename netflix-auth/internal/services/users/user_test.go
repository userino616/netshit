package users

import (
	"errors"
	"netflix-auth/internal/fixtures"
	"netflix-auth/internal/models"
	"netflix-auth/internal/repository/users"
	"netflix-auth/pkg/hash"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/stretchr/testify/assert"
)

var (
	r users.UserRepository
	s Service

	existsUserID, _ = uuid.Parse("06ea7596-210d-11ec-a866-0242ac140003")
)

func init() {
	r = users.NewUserRepository(fixtures.GetDB())
	s = NewUserService(r, hash.NewHashService(fixtures.CFG))
}

func TestUserService_Create(t *testing.T) {
	fixtures.PrepareFixtures()

	tableCases := []struct {
		testName        string
		email, password string
		err             error
	}{
		{
			testName: "valid test",
			email:    "new@email.com",
			password: "test12345",
			err:      nil,
		},
		{
			testName: "email exists",
			email:    "exists@email.com",
			password: "test12345",
			err:      ErrUserAlreadyExists,
		},
		{
			testName: "short password",
			email:    "new@email.com",
			password: "1",
			err:      validation.ErrLengthTooShort,
		},
		{
			testName: "invalid email",
			email:    "newemail.com",
			password: "112313d",
			err:      is.ErrEmail,
		},
	}
	for _, tc := range tableCases {
		_, err := s.Create(models.User{
			Email:    tc.email,
			Password: tc.password,
		})
		if tc.err != nil && err != nil {
			assert.Equal(t, true, errors.As(err, &tc.err), tc.testName)
		}
	}
}

func TestUserService_GetByEmail(t *testing.T) {
	fixtures.PrepareFixtures()

	tableCases := []struct {
		testName string
		id       uuid.UUID
		email    string
		err      error
	}{
		{
			testName: "user exists",
			email:    "exists@email.com",
			id:       existsUserID,
			err:      nil,
		},
		{
			testName: "user not exists",
			email:    "notexists@email.com",
			err:      pg.ErrNoRows,
		},
	}
	for _, tc := range tableCases {
		user, err := s.GetByEmail(tc.email)
		if err != nil {
			assert.Equal(t, true, errors.Is(err, tc.err), tc.testName)
		} else {
			assert.Equal(t, tc.id, user.ID, tc.testName)
		}
	}
}

func TestUserService_GetByID(t *testing.T) {
	fixtures.PrepareFixtures()

	tableCases := []struct {
		testName string
		id       uuid.UUID
		err      error
	}{
		{
			testName: "user exists",
			id:       existsUserID,
			err:      nil,
		},
		{
			testName: "user not exists",
			id:       uuid.New(),
			err:      pg.ErrNoRows,
		},
	}
	for _, tc := range tableCases {
		user, err := s.GetByID(tc.id)
		if err != nil {
			assert.Equal(t, true, errors.Is(err, tc.err), tc.testName)
		} else {
			assert.Equal(t, tc.id, user.ID, tc.testName)
		}
	}
}
