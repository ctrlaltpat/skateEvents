package database

import (
	"database/sql"

	"github.com/ctrlaltpat/skate-events/internal/models"
)

type Models struct {
	Users     models.UserModel
	Events    models.EventModel
	Attendees models.AttendeeModel
	Skates    models.SkateModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:     models.UserModel{DB: db},
		Events:    models.EventModel{DB: db},
		Attendees: models.AttendeeModel{DB: db},
		Skates:    models.SkateModel{DB: db},
	}
}
