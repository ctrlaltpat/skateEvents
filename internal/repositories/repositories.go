package repositories

import "database/sql"

type Repositories struct {
	Event EventRepository
	User  UserRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Event: EventRepository{DB: db},
		User:  UserRepository{DB: db},
	}
}
