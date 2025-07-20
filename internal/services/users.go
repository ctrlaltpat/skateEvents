package services

import (
	"context"
	"log"

	"github.com/ctrlaltpat/skate-events/internal/models"
	"github.com/ctrlaltpat/skate-events/internal/repositories"
)

type UserService struct {
	Repo repositories.UserRepository
}

func (svc UserService) Register(ctx context.Context, user *models.User) (*models.User, error) {
	log.Println("Registering user")
	user, err := svc.Repo.Insert(ctx, user)
	if err != nil {
		log.Println("Error creating user in repo")
		return nil, err
	}
	return user, nil
}
