package service

import "internal\repository\repository"

type ReminderService struct {
	storage *InMemmory
}

func NewReminderService(a *InMemmory) *ReminderService {

	return &ReminderService{
		storage: a,
	}
}

func (s *ReminderService) ReminderCreate(id string) {
	reminderModel := &Reminder{}
	s.storage.Add(reminderModel)
}
