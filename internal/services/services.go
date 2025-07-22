package services

import "github.com/ctrlaltpat/skate-events/internal/repositories"

type Services struct {
	Event EventService
	User  UserService
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		Event: EventService{EventRepo: repos.Event, AttendeeRepo: repos.Attendee},
		User:  UserService{Repo: repos.User},
	}
}
