package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/ctrlaltpat/skate-events/internal/models"
)

type AttendeeRepository struct {
	DB *sql.DB
}

func (repo AttendeeRepository) Insert(ctx context.Context, attendee *models.Attendee) (*models.Attendee, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "INSERT INTO attendees (event_id, user_id) VALUES ($1, $2) RETURNING id"
	err := repo.DB.QueryRowContext(ctx, query, attendee.EventId, attendee.UserId).Scan(&attendee.Id)
	if err != nil {
		return nil, err
	}

	return attendee, nil
}

// func (repo AttendeeRepository) GetAllAttendeesByEventId(ctx context.Context, eventId int) ([]models.Attendee, error) {
func (repo AttendeeRepository) GetAllAttendeesByEventId(ctx context.Context, eventId int) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// query := "SELECT * FROM attendees WHERE event_id = $1"
	query := `
		SELECT u.email, u.id, u.name
		FROM users u
		JOIN attendees a ON u.id = a.user_id
		WHERE a.event_id = $1
	`
	rows, err := repo.DB.QueryContext(ctx, query, eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// attendees := []models.Attendee{}
	attendees := []models.User{}
	for rows.Next() {
		// var attendee models.Attendee
		var attendee models.User
		// err := rows.Scan(&attendee.Id, &attendee.EventId, &attendee.UserId)
		err := rows.Scan(&attendee.Email, &attendee.Id, &attendee.Name)
		if err != nil {
			return nil, err
		}
		attendees = append(attendees, attendee)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return attendees, nil
}

func (repo AttendeeRepository) GetByUserIdAndEventId(ctx context.Context, userId int, eventId int) (*models.Attendee, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "SELECT * FROM attendees WHERE user_id = $1 AND event_id = $2"
	row := repo.DB.QueryRowContext(ctx, query, userId, eventId)

	var attendee models.Attendee
	err := row.Scan(&attendee.Id, &attendee.EventId, &attendee.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &attendee, nil
}
