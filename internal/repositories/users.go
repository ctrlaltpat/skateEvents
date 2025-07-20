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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id`
	err := repo.DB.QueryRowContext(ctx, stmt, user.Email, user.Password, user.Name).Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
