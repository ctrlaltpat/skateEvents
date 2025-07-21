package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/ctrlaltpat/skate-events/internal/models"
)

type EventRepository struct {
	DB *sql.DB
}

func (repo EventRepository) Insert(ctx context.Context, event *models.Event) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO events (owner_id, name, description, date, location) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	err := repo.DB.QueryRowContext(ctx, query, event.OwnerId, event.Name, event.Description, event.Date, event.Location).Scan(&event.Id)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (repo EventRepository) GetAll(ctx context.Context) ([]models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM events"

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []models.Event{}

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (repo EventRepository) Get(ctx context.Context, id int) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM events WHERE id = $1"

	row := repo.DB.QueryRowContext(ctx, query, id)

	var event models.Event

	err := row.Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &event, nil
}

func (repo EventRepository) Update(ctx context.Context, id int, event *models.Event) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "UPDATE events SET name = $1, description = $2, date = $3, location = $4 WHERE id = $5"

	_, err := repo.DB.ExecContext(ctx, query, event.Name, event.Description, event.Date, event.Location, id)
	if err != nil {
		return nil, err
	}
	event.Id = id
	return event, nil
}

func (repo EventRepository) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM events WHERE id = $1"

	_, err := repo.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
