package services

import (
	"context"
	"log"

	"github.com/ctrlaltpat/skate-events/internal/models"
	"github.com/ctrlaltpat/skate-events/internal/repositories"
)

type EventService struct {
	Repo repositories.EventRepository
}

func (svc EventService) GetAll(ctx context.Context) ([]models.Event, error) {
	events, err := svc.Repo.GetAll(ctx)
	if err != nil {
		log.Println("Error getting events")
		return nil, err
	}
	return events, nil
}

func (svc EventService) Get(ctx context.Context, id int) (*models.Event, error) {
	event, err := svc.Repo.Get(ctx, id)
	if err != nil {
		log.Println("Error getting an event")
		return nil, err
	}
	return event, nil
}

func (svc EventService) Create(ctx context.Context, event *models.Event) (*models.Event, error) {
	event, err := svc.Repo.Insert(ctx, event)
	if err != nil {
		log.Println("Error creating event")
		return nil, err
	}
	return event, nil
}

func (svc EventService) Update(ctx context.Context, id int, event *models.Event) (*models.Event, error) {
	event, err := svc.Repo.Update(ctx, id, event)
	if err != nil {
		log.Println("Error updating event")
		return nil, err
	}
	return event, nil
}

func (svc EventService) Delete(ctx context.Context, id int) error {
	err := svc.Repo.Delete(ctx, id)
	if err != nil {
		log.Println("Error deleting event")
		return err
	}
	return nil
}