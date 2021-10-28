package users

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"netflix-auth/internal/models"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	GetByID(id uuid.UUID) (models.User, error)
	GetByEmail(email string) (models.User, error)
}

type userRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(id uuid.UUID) (user models.User, err error) {
	err = r.db.
		Model(&user).
		Where("id = ?", id).
		Select()

	return user, err
}

func (r *userRepository) GetByEmail(email string) (user models.User, err error) {
	err = r.db.
		Model(&user).
		Where("email = ?", email).
		Select()

	return user, err
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	_, err := r.db.
		Model(&user).
		Insert()

	return user, err
}
