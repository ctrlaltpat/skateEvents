package handlers

import "github.com/ctrlaltpat/skate-events/internal/services"

type Handlers struct {
	User *UserHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		User: &UserHandler{Service: services.User},
	}
}
