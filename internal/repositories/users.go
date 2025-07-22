package repositories

import (
	"context"
	"database/sql"
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
		return nil, err
	}
	return user, nil
}

func (repo UserRepository) GetById(ctx context.Context, id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	stmt := `SELECT id, email, name FROM users WHERE id = $1`
	row := repo.DB.QueryRowContext(ctx, stmt, id)

	var user models.User
	err := row.Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	stmt := `SELECT id, email, name FROM users WHERE email = $1`
	row := repo.DB.QueryRowContext(ctx, stmt, email)

	var user models.User
	err := row.Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
