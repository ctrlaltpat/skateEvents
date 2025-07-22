package handlers

import "github.com/ctrlaltpat/skate-events/internal/services"

type Handlers struct {
	Event *EventHandler
	User  *UserHandler
}

func NewHandlers(services *services.Services, jwtSecret string) *Handlers {
	return &Handlers{
		Event: &EventHandler{EventService: services.Event, UserService: services.User},
		User:  &UserHandler{Service: services.User, JWTSecret: jwtSecret},
	}
}
