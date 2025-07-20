package services

import "github.com/ctrlaltpat/skate-events/internal/repositories"

type Services struct {
	User UserService
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		User: UserService{Repo: repos.User},
	}
}
