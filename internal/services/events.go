package services

import (
	"context"
	"log"

	"github.com/ctrlaltpat/skate-events/internal/models"
	"github.com/ctrlaltpat/skate-events/internal/repositories"
)

type EventService struct {
	EventRepo    repositories.EventRepository
	AttendeeRepo repositories.AttendeeRepository
}

func (svc EventService) GetAll(ctx context.Context) ([]models.Event, error) {
	events, err := svc.EventRepo.GetAll(ctx)
	if err != nil {
		log.Println("Error getting events")
		return nil, err
	}
	return events, nil
}

func (svc EventService) Get(ctx context.Context, id int) (*models.Event, error) {
	event, err := svc.EventRepo.Get(ctx, id)
	if err != nil {
		log.Println("Error getting an event")
		return nil, err
	}
	return event, nil
}

func (svc EventService) Create(ctx context.Context, event *models.Event) (*models.Event, error) {
	event, err := svc.EventRepo.Insert(ctx, event)
	if err != nil {
		log.Println("Error creating event")
		return nil, err
	}
	return event, nil
}

func (svc EventService) Update(ctx context.Context, id int, event *models.Event) (*models.Event, error) {
	event, err := svc.EventRepo.Update(ctx, id, event)
	if err != nil {
		log.Println("Error updating event")
		return nil, err
	}
	return event, nil
}

func (svc EventService) UpdateStatus(ctx context.Context, id int, status string) (*models.Event, error) {
    event, err := svc.EventRepo.UpdateStatus(ctx, id, status)
    if err != nil {
        log.Println("Error updating event status")
        return nil, err
    }
    return event, nil
}

func (svc EventService) Delete(ctx context.Context, id int) (int64, error) {
	rows, err := svc.EventRepo.Delete(ctx, id)
	if err != nil {
		log.Println("Error deleting event")
		return rows, err
	}
	return rows, nil
}

func (svc EventService) IsAlreadyAttending(ctx context.Context, userId int, eventId int) (bool, error) {
	attendee, err := svc.AttendeeRepo.GetByUserIdAndEventId(ctx, userId, eventId)
	if err != nil {
		log.Println("Error checking attendee existence")
		return attendee != nil, err
	}
	return attendee != nil, nil
}

func (svc EventService) AddAttendee(ctx context.Context, eventId int, userId int) (*models.Attendee, error) {
	attendee := &models.Attendee{
		EventId: eventId,
		UserId:  userId,
	}
	attendee, err := svc.AttendeeRepo.Insert(ctx, attendee)
	if err != nil {
		log.Println("Error adding attendee")
		return nil, err
	}
	return attendee, nil
}

func (svc EventService) GetAttendees(ctx context.Context, eventId int) ([]models.User, error) {
	attendees, err := svc.AttendeeRepo.GetAllAttendeesByEventId(ctx, eventId)
	if err != nil {
		log.Println("Error retrieving attendees for this event")
		return nil, err
	}
	return attendees, nil
}
