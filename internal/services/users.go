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
	user, err := svc.Repo.Insert(ctx, user)
	if err != nil {
		log.Println("Error creating user")
		return nil, err
	}
	return user, nil
}

func (svc UserService) GetById(ctx context.Context, id int) (*models.User, error) {
	user, err := svc.Repo.GetById(ctx, id)
	if err != nil {
		log.Println("Error getting user")
		return nil, err
	}
	return user, nil
}

func (svc UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := svc.Repo.GetByEmail(ctx, email)
	if err != nil {
		log.Println("Error getting user by email")
		return nil, err
	}
	return user, nil
}
