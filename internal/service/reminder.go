package service

import (
	"testdk/internal"
	"testdk/internal/domain"
	"testdk/internal/repository"
)

type ReminderService struct {
	storage *repository.InMemmory
}

func NewReminderService(a *repository.InMemmory) *ReminderService {

	return &ReminderService{
		storage: a,
	}
}

func (s *ReminderService) ReminderCreate() *internal.Reminder {
	id := domain.CreateEntityId()
	reminderModel := internal.NewReminder(id)
	s.storage.Add(reminderModel)
	return reminderModel
}
