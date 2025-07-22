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
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "INSERT INTO events (owner_id, name, description, start_date, end_date, location, status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	err := repo.DB.QueryRowContext(ctx, query, event.OwnerId, event.Name, event.Description, event.StartDate, event.EndDate, event.Location, event.Status).Scan(&event.Id)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (repo EventRepository) GetAll(ctx context.Context) ([]models.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
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
		err := rows.Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.StartDate, &event.EndDate, &event.Location, &event.Status)
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
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "SELECT * FROM events WHERE id = $1"

	row := repo.DB.QueryRowContext(ctx, query, id)

	var event models.Event

	err := row.Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.StartDate, &event.EndDate, &event.Location, &event.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &event, nil
}

func (repo EventRepository) Update(ctx context.Context, id int, event *models.Event) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "UPDATE events SET name = $1, description = $2, start_date = $3, end_date = $4, location = $5, status = $6 WHERE id = $7"

	_, err := repo.DB.ExecContext(ctx, query, event.Name, event.Description, event.StartDate, event.EndDate, event.Location, event.Status, id)
	if err != nil {
		return nil, err
	}
	event.Id = id
	return event, nil
}

func (repo EventRepository) UpdateStatus(ctx context.Context, id int, status string) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	stmt := `UPDATE events SET status = $1 WHERE id = $2 RETURNING id, name, description, start_date, end_date, location, status`
	row := repo.DB.QueryRowContext(ctx, stmt, status, id)

	var event models.Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.StartDate, &event.EndDate, &event.Location, &event.Status)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (repo EventRepository) Delete(ctx context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "DELETE FROM events WHERE id = $1"

	res, err := repo.DB.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
