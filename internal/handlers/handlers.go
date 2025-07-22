package handlers

import "github.com/ctrlaltpat/skate-events/internal/services"

type Handlers struct {
	Event *EventHandler
	User  *UserHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		Event: &EventHandler{EventService: services.Event, UserService: services.User},
		User:  &UserHandler{Service: services.User},
	}
}
