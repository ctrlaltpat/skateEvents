package repositories

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/ctrlaltpat/skate-events/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (repo UserRepository) Insert(ctx context.Context, user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id`
	err := repo.DB.QueryRowContext(ctx, stmt, user.Email, user.Password, user.Name).Scan(&user.Id)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint") {
			return nil, errors.New("user email already exists")
		}
		return nil, err
	}
	return user, nil
}

func (repo UserRepository) getUser(ctx context.Context, query string, args ...any) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var user models.User
	err := repo.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo UserRepository) GetById(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	return repo.getUser(ctx, query, id)
}

func (repo UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = $1`
	return repo.getUser(ctx, query, email)
}
