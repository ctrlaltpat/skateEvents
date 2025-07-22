package repositories

import "database/sql"

type Repositories struct {
	Attendee AttendeeRepository
	Event    EventRepository
	User     UserRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Attendee: AttendeeRepository{DB: db},
		Event:    EventRepository{DB: db},
		User:     UserRepository{DB: db},
	}
}
